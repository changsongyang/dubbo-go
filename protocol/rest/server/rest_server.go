/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strconv"
)

import (
	"github.com/dubbogo/gost/log/logger"

	perrors "github.com/pkg/errors"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	rest_config "dubbo.apache.org/dubbo-go/v3/protocol/rest/config"
)

const parseParameterErrorStr = "an error occurred while parsing parameters on the server"

// RestServer user can implement this server interface
type RestServer interface {
	// Start rest server
	Start(url *common.URL)
	// Deploy a http api
	Deploy(restMethodConfig *rest_config.RestMethodConfig, routeFunc func(request RestServerRequest, response RestServerResponse))
	// UnDeploy a http api
	UnDeploy(restMethodConfig *rest_config.RestMethodConfig)
	// Destroy rest server
	Destroy()
}

// RestServerRequest interface
type RestServerRequest interface {
	// RawRequest get the Ptr of http.Request
	RawRequest() *http.Request
	// PathParameter get the path parameter by name
	PathParameter(name string) string
	// PathParameters get the map of the path parameters
	PathParameters() map[string]string
	// QueryParameter get the query parameter by name
	QueryParameter(name string) string
	// QueryParameters get the map of query parameters
	QueryParameters(name string) []string
	// BodyParameter get the body parameter of name
	BodyParameter(name string) (string, error)
	// HeaderParameter get the header parameter of name
	HeaderParameter(name string) string
	// ReadEntity checks the Accept header and reads the content into the entityPointer.
	ReadEntity(entityPointer any) error
}

// RestServerResponse interface
type RestServerResponse interface {
	http.ResponseWriter
	// WriteError writes the http status and the error string on the response. err can be nil.
	// Return an error if writing was not successful.
	WriteError(httpStatus int, err error) (writeErr error)
	// WriteEntity marshals the value using the representation denoted by the Accept Header.
	WriteEntity(value any) error
}

// GetRouteFunc is a route function will be invoked by http server
func GetRouteFunc(invoker base.Invoker, methodConfig *rest_config.RestMethodConfig) func(req RestServerRequest, resp RestServerResponse) {
	return func(req RestServerRequest, resp RestServerResponse) {
		var (
			err  error
			args []any
		)
		svc := common.ServiceMap.GetServiceByServiceKey(invoker.GetURL().Protocol, invoker.GetURL().ServiceKey())
		// get method
		method := svc.Method()[methodConfig.MethodName]
		argsTypes := method.ArgsType()
		replyType := method.ReplyType()
		// two ways to prepare arguments
		// if method like this 'func1(req []any, rsp *User) error'
		// we don't have arguments type
		if (len(argsTypes) == 1 || len(argsTypes) == 2 && replyType == nil) &&
			argsTypes[0].String() == "[]interface {}" {
			args, err = getArgsInterfaceFromRequest(req, methodConfig)
		} else {
			args, err = getArgsFromRequest(req, argsTypes, methodConfig)
		}
		if err != nil {
			logger.Errorf("[Go Restful] parsing http parameters error:%v", err)
			err = resp.WriteError(http.StatusInternalServerError, errors.New(parseParameterErrorStr))
			if err != nil {
				logger.Errorf("[Go Restful] WriteErrorString error:%v", err)
			}
		}
		result := invoker.Invoke(context.Background(), invocation.NewRPCInvocation(methodConfig.MethodName, args, make(map[string]any)))
		if result.Error() != nil {
			err = resp.WriteError(http.StatusInternalServerError, result.Error())
			if err != nil {
				logger.Errorf("[Go Restful] WriteError error:%v", err)
			}
			return
		}
		err = resp.WriteEntity(result.Result())
		if err != nil {
			logger.Errorf("[Go Restful] WriteEntity error:%v", err)
		}
	}
}

// getArgsInterfaceFromRequest when service function like GetUser(req []any, rsp *User) error
// use this method to get arguments
func getArgsInterfaceFromRequest(req RestServerRequest, methodConfig *rest_config.RestMethodConfig) ([]any, error) {
	argsMap := make(map[int]any, 8)
	maxKey := 0
	for k, v := range methodConfig.PathParamsMap {
		if maxKey < k {
			maxKey = k
		}
		argsMap[k] = req.PathParameter(v)
	}
	for k, v := range methodConfig.QueryParamsMap {
		if maxKey < k {
			maxKey = k
		}
		params := req.QueryParameters(v)
		if len(params) == 1 {
			argsMap[k] = params[0]
		} else {
			argsMap[k] = params
		}
	}
	for k, v := range methodConfig.HeadersMap {
		if maxKey < k {
			maxKey = k
		}
		argsMap[k] = req.HeaderParameter(v)
	}
	if methodConfig.Body >= 0 {
		if maxKey < methodConfig.Body {
			maxKey = methodConfig.Body
		}
		m := make(map[string]any)
		// TODO read as a slice
		if err := req.ReadEntity(&m); err != nil {
			return nil, perrors.Errorf("[Go restful] Read body entity as map[string]any error:%v", err)
		}
		argsMap[methodConfig.Body] = m
	}
	args := make([]any, maxKey+1)
	for k, v := range argsMap {
		if k >= 0 {
			args[k] = v
		}
	}
	return args, nil
}

// getArgsFromRequest get arguments from server.RestServerRequest
func getArgsFromRequest(req RestServerRequest, argsTypes []reflect.Type, methodConfig *rest_config.RestMethodConfig) ([]any, error) {
	argsLength := len(argsTypes)
	args := make([]any, argsLength)
	for i, t := range argsTypes {
		args[i] = reflect.Zero(t).Interface()
	}
	if err := assembleArgsFromPathParams(methodConfig, argsLength, argsTypes, req, args); err != nil {
		return nil, err
	}
	if err := assembleArgsFromQueryParams(methodConfig, argsLength, argsTypes, req, args); err != nil {
		return nil, err
	}
	if err := assembleArgsFromBody(methodConfig, argsTypes, req, args); err != nil {
		return nil, err
	}
	if err := assembleArgsFromHeaders(methodConfig, req, argsLength, argsTypes, args); err != nil {
		return nil, err
	}
	return args, nil
}

// assembleArgsFromHeaders assemble arguments from headers
func assembleArgsFromHeaders(methodConfig *rest_config.RestMethodConfig, req RestServerRequest, argsLength int, argsTypes []reflect.Type, args []any) error {
	for k, v := range methodConfig.HeadersMap {
		param := req.HeaderParameter(v)
		if k < 0 || k >= argsLength {
			return perrors.Errorf("[Go restful] Header param parse error, the index %v args of method:%v doesn't exist", k, methodConfig.MethodName)
		}
		t := argsTypes[k]
		if t.Kind() == reflect.Ptr {
			t = t.Elem()
		}
		if t.Kind() == reflect.String {
			args[k] = param
		} else {
			return perrors.Errorf("[Go restful] Header param parse error, the index %v args's type isn't string", k)
		}
	}
	return nil
}

// assembleArgsFromBody assemble arguments from body
func assembleArgsFromBody(methodConfig *rest_config.RestMethodConfig, argsTypes []reflect.Type, req RestServerRequest, args []any) error {
	if methodConfig.Body >= 0 && methodConfig.Body < len(argsTypes) {
		t := argsTypes[methodConfig.Body]
		kind := t.Kind()
		if kind == reflect.Ptr {
			t = t.Elem()
		}
		var ni any
		if t.String() == "[]interface {}" {
			ni = make([]map[string]any, 0)
		} else if t.String() == "interface {}" {
			ni = make(map[string]any)
		} else {
			n := reflect.New(t)
			if n.CanInterface() {
				ni = n.Interface()
			}
		}
		if err := req.ReadEntity(&ni); err != nil {
			return perrors.Errorf("[Go restful] Read body entity error, error is %v", perrors.WithStack(err))
		}
		args[methodConfig.Body] = ni
	}
	return nil
}

// assembleArgsFromQueryParams assemble arguments from query params
func assembleArgsFromQueryParams(methodConfig *rest_config.RestMethodConfig, argsLength int, argsTypes []reflect.Type, req RestServerRequest, args []any) error {
	var (
		err   error
		param any
		i64   int64
	)
	for k, v := range methodConfig.QueryParamsMap {
		if k < 0 || k >= argsLength {
			return perrors.Errorf("[Go restful] Query param parse error, the index %v args of method:%v doesn't exist", k, methodConfig.MethodName)
		}
		t := argsTypes[k]
		kind := t.Kind()
		if kind == reflect.Ptr {
			t = t.Elem()
			kind = t.Kind()
		}
		switch kind {
		case reflect.Slice:
			param = req.QueryParameters(v)
		case reflect.String:
			param = req.QueryParameter(v)
		case reflect.Int:
			param, err = strconv.Atoi(req.QueryParameter(v))
		case reflect.Int32:
			i64, err = strconv.ParseInt(req.QueryParameter(v), 10, 32)
			if err == nil {
				param = int32(i64)
			}
		case reflect.Int64:
			param, err = strconv.ParseInt(req.QueryParameter(v), 10, 64)
		default:
			return perrors.Errorf("[Go restful] Query param parse error, the index %v args's type isn't int or string or slice", k)
		}

		if err != nil {
			return perrors.Errorf("[Go restful] Query param parse error, error:%v", perrors.WithStack(err))
		}
		args[k] = param
	}
	return nil
}

// assembleArgsFromPathParams assemble arguments from path params
func assembleArgsFromPathParams(methodConfig *rest_config.RestMethodConfig, argsLength int, argsTypes []reflect.Type, req RestServerRequest, args []any) error {
	var (
		err   error
		param any
		i64   int64
	)
	for k, v := range methodConfig.PathParamsMap {
		if k < 0 || k >= argsLength {
			return perrors.Errorf("[Go restful] Path param parse error, the index %v args of method:%v doesn't exist", k, methodConfig.MethodName)
		}
		t := argsTypes[k]
		kind := t.Kind()
		if kind == reflect.Ptr {
			t = t.Elem()
			kind = t.Kind()
		}

		switch kind {
		case reflect.Int:
			param, err = strconv.Atoi(req.PathParameter(v))
		case reflect.Int32:
			i64, err = strconv.ParseInt(req.PathParameter(v), 10, 32)
			if err == nil {
				param = int32(i64)
			}
		case reflect.Int64:
			param, err = strconv.ParseInt(req.PathParameter(v), 10, 64)
		case reflect.String:
			param = req.PathParameter(v)
		default:
			return perrors.Errorf("[Go restful] Path param parse error, the index %v args's type isn't int or string", k)
		}

		if err != nil {
			return perrors.Errorf("[Go restful] Path param parse error, error is %v", perrors.WithStack(err))
		}
		args[k] = param
	}
	return nil
}
