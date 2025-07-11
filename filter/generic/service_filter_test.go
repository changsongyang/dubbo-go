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

package generic

import (
	"context"
	"fmt"
	"net/url"
	"testing"
)

import (
	hessian "github.com/apache/dubbo-go-hessian2"

	"github.com/golang/mock/gomock"

	perrors "github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
	"dubbo.apache.org/dubbo-go/v3/filter/generic/generalizer"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	"dubbo.apache.org/dubbo-go/v3/protocol/mock"
	"dubbo.apache.org/dubbo-go/v3/protocol/result"
)

type MockHelloService struct{}

func (s *MockHelloService) Hello(who string) (string, error) {
	return fmt.Sprintf("hello, %s", who), nil
}

func (s *MockHelloService) JavaClassName() string {
	return "org.apache.dubbo.hello"
}

func (s *MockHelloService) Reference() string {
	return "org.apache.dubbo.test"
}

func (s *MockHelloService) HelloPB(req *generalizer.RequestType) (*generalizer.ResponseType, error) {
	if req.GetId() == 1 {
		return &generalizer.ResponseType{
			Code:    200,
			Id:      1,
			Name:    "xavierniu",
			Message: "Nice to meet you",
		}, nil
	}
	return nil, perrors.Errorf("people not found")
}

func TestServiceFilter_Invoke(t *testing.T) {
	filter := &genericServiceFilter{}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockInvoker := mock.NewMockInvoker(ctrl)

	// methodName is not "$invoke"
	inv1 := invocation.NewRPCInvocation("test", nil, nil)
	mockInvoker.EXPECT().Invoke(gomock.Any(), gomock.Eq(inv1))
	_ = filter.Invoke(context.Background(), mockInvoker, inv1)
	// arguments are nil
	inv2 := invocation.NewRPCInvocation(constant.Generic, nil, nil)
	mockInvoker.EXPECT().Invoke(gomock.Any(), gomock.Eq(inv2))
	_ = filter.Invoke(context.Background(), mockInvoker, inv2)
	// the number of arguments is not 3
	inv3 := invocation.NewRPCInvocation(constant.Generic, []any{"hello"}, nil)
	mockInvoker.EXPECT().Invoke(gomock.Any(), gomock.Eq(inv3))
	_ = filter.Invoke(context.Background(), mockInvoker, inv3)

	// hello service
	service := &MockHelloService{}
	// invoke URL
	ivkUrl := common.NewURLWithOptions(
		common.WithProtocol("test"),
		common.WithParams(url.Values{}),
		common.WithParamsValue(constant.InterfaceKey, service.Reference()),
		common.WithParamsValue(constant.GenericKey, constant.GenericSerializationDefault))
	// registry RPC service
	_, err := common.ServiceMap.Register(ivkUrl.GetParam(constant.InterfaceKey, ""),
		ivkUrl.Protocol,
		"",
		"",
		service)
	assert.Nil(t, err)

	// mock
	mockInvoker.EXPECT().GetURL().Return(ivkUrl).Times(3)

	// invoke a method without errors using default generalization
	invocation4 := invocation.NewRPCInvocation(constant.Generic,
		[]any{
			"Hello",
			[]string{"java.lang.String"},
			[]hessian.Object{"world"},
		}, map[string]any{
			constant.GenericKey: "true",
		})
	// invoke a non-existed method
	invocation5 := invocation.NewRPCInvocation(constant.Generic,
		[]any{
			"hello11",
			[]string{"java.lang.String"},
			[]hessian.Object{"world"},
		}, map[string]any{
			constant.GenericKey: "true",
		})
	// invoke a method with incorrect arguments
	invocation6 := invocation.NewRPCInvocation(constant.Generic,
		[]any{
			"Hello",
			[]string{"java.lang.String", "java.lang.String"},
			[]hessian.Object{"world", "haha"},
		}, map[string]any{
			constant.GenericKey: "true",
		})
	// invoke a method without errors using protobuf-json generalization
	//invocation7 := invocation.NewRPCInvocation(constant.Generic,
	//	[]any{
	//		"HelloPB",
	//		[]string{},
	//		[]hessian.Object{"{\"id\":1}"},
	//	}, map[string]any{
	//		constant.GenericKey: constant.GenericSerializationProtobuf,
	//	})

	mockInvoker.EXPECT().Invoke(gomock.Any(), gomock.All(
		gomock.Not(inv1),
		gomock.Not(inv2),
		gomock.Not(inv3),
	)).DoAndReturn(
		func(ctx context.Context, invocation base.Invocation) result.Result {
			switch invocation.MethodName() {
			case "Hello":
				who := invocation.Arguments()[0].(string)
				res, _ := service.Hello(who)
				return &result.RPCResult{
					Rest: res,
				}
			case "HelloPB":
				req := invocation.Arguments()[0].(*generalizer.RequestType)
				res, _ := service.HelloPB(req)
				return &result.RPCResult{
					Rest: res,
				}
			default:
				panic("this branch shouldn't be reached")
			}
		}).AnyTimes()

	result := filter.Invoke(context.Background(), mockInvoker, invocation4)
	assert.Nil(t, result.Error())
	assert.Equal(t, "hello, world", result.Result())

	result = filter.Invoke(context.Background(), mockInvoker, invocation5)
	assert.Equal(t,
		fmt.Sprintf("\"hello11\" method is not found, service key: %s", ivkUrl.ServiceKey()),
		fmt.Sprintf("%v", result.Error()))

	result = filter.Invoke(context.Background(), mockInvoker, invocation6)
	assert.Equal(t,
		"the number of args(=2) is not matched with \"Hello\" method",
		fmt.Sprintf("%v", result.Error()))

	//result = filter.Invoke(context.Background(), mockInvoker, invocation7)
	//assert.Equal(t, int64(200), result.Result().(*generalizer.ResponseType).GetCode())
	//assert.Equal(t, int64(1), result.Result().(*generalizer.ResponseType).GetId())
	//assert.Equal(t, "xavierniu", result.Result().(*generalizer.ResponseType).GetName())
	//assert.Equal(t, "Nice to meet you", result.Result().(*generalizer.ResponseType).GetMessage())

}

func TestServiceFilter_OnResponse(t *testing.T) {
	filter := &genericServiceFilter{}

	// invoke a method without errors
	invocation1 := invocation.NewRPCInvocation(constant.Generic,
		[]any{
			"hello",
			[]any{"java.lang.String"},
			[]any{"world"},
		}, map[string]any{
			constant.GenericKey: "true",
		})

	rpcResult := &result.RPCResult{
		Rest: "result",
	}

	result := filter.OnResponse(context.Background(), rpcResult, nil, invocation1)
	assert.Equal(t, "result", result.Result())
}
