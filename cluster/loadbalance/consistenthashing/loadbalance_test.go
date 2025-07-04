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

package consistenthashing

import (
	"fmt"
	"testing"
)

import (
	"github.com/stretchr/testify/suite"
)

import (
	"dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
)

const (
	ip       = "192.168.1.0"
	port8080 = 8080
	port8081 = 8081

	url8080Short = "dubbo://192.168.1.0:8080"
	url8081Short = "dubbo://192.168.1.0:8081"
	url20000     = "dubbo://192.168.1.0:20000/org.apache.demo.HelloService?methods.echo.hash.arguments=0,1"
	url8080      = "dubbo://192.168.1.0:8080/org.apache.demo.HelloService?methods.echo.hash.arguments=0,1"
	url8081      = "dubbo://192.168.1.0:8081/org.apache.demo.HelloService?methods.echo.hash.arguments=0,1"
	url8082      = "dubbo://192.168.1.0:8082/org.apache.demo.HelloService?methods.echo.hash.arguments=0,1"
)

func TestConsistentHashSelectorSuite(t *testing.T) {
	suite.Run(t, new(consistentHashSelectorSuite))
}

type consistentHashSelectorSuite struct {
	suite.Suite
	selector *selector
}

func (s *consistentHashSelectorSuite) SetupTest() {
	var invokers []base.Invoker
	url, _ := common.NewURL(url20000)
	invokers = append(invokers, base.NewBaseInvoker(url))
	s.selector = newSelector(invokers, "echo", 999944)
}

func (s *consistentHashSelectorSuite) TestToKey() {
	result := s.selector.toKey([]any{"username", "age"})
	s.Equal(result, "usernameage")
}

func (s *consistentHashSelectorSuite) TestSelectForKey() {
	url1, _ := common.NewURL(url8080Short)
	url2, _ := common.NewURL(url8081Short)
	s.selector.virtualInvokers = make(map[uint32]base.Invoker)
	s.selector.virtualInvokers[99874] = base.NewBaseInvoker(url1)
	s.selector.virtualInvokers[9999945] = base.NewBaseInvoker(url2)
	s.selector.keys = []uint32{99874, 9999945}
	result := s.selector.selectForKey(9999944)
	s.Equal(result.GetURL().String(), url8081Short+"?")
}

func TestConsistentHashLoadBalanceSuite(t *testing.T) {
	suite.Run(t, new(consistentHashLoadBalanceSuite))
}

type consistentHashLoadBalanceSuite struct {
	suite.Suite
	url1     *common.URL
	url2     *common.URL
	url3     *common.URL
	invokers []base.Invoker
	invoker1 base.Invoker
	invoker2 base.Invoker
	invoker3 base.Invoker
	lb       loadbalance.LoadBalance
}

func (s *consistentHashLoadBalanceSuite) SetupTest() {
	var err error
	s.url1, err = common.NewURL(url8080)
	s.NoError(err)
	s.url2, err = common.NewURL(url8081)
	s.NoError(err)
	s.url3, err = common.NewURL(url8082)
	s.NoError(err)

	s.invoker1 = base.NewBaseInvoker(s.url1)
	s.invoker2 = base.NewBaseInvoker(s.url2)
	s.invoker3 = base.NewBaseInvoker(s.url3)

	s.invokers = append(s.invokers, s.invoker1, s.invoker2, s.invoker3)
	s.lb = newConshashLoadBalance()
}

func (s *consistentHashLoadBalanceSuite) TestSelect() {
	args := []any{"name", "password", "age"}
	invoker := s.lb.Select(s.invokers, invocation.NewRPCInvocation("echo", args, nil))
	s.Equal(fmt.Sprintf("%s:%d", ip, port8081), invoker.GetURL().Location)

	args = []any{"ok", "abc"}
	invoker = s.lb.Select(s.invokers, invocation.NewRPCInvocation("echo", args, nil))
	s.Equal(fmt.Sprintf("%s:%d", ip, port8080), invoker.GetURL().Location)
}
