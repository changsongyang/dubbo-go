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

package sentinel

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

import (
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/flow"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/assert"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	"dubbo.apache.org/dubbo-go/v3/protocol/result"
)

func TestSentinelFilter_QPS(t *testing.T) {
	url, err := common.NewURL("dubbo://127.0.0.1:20000/UserProvider?anyhost=true&" +
		"version=1.0.0&group=myGroup&" +
		"application=BDTService&category=providers&default.timeout=10000&dubbo=dubbo-provider-golang-1.0.0&" +
		"environment=dev&interface=com.ikurento.user.UserProvider&ip=192.168.56.1&methods=GetUser%2C&" +
		"module=dubbogo+user-info+server&org=ikurento.com&owner=ZX&pid=1447&revision=0.0.1&" +
		"side=provider&timeout=3000&timestamp=1556509797245&bean.name=UserProvider")
	assert.NoError(t, err)
	mockInvoker := base.NewBaseInvoker(url)
	interfaceResourceName, _ := getResourceName(mockInvoker,
		invocation.NewRPCInvocation("hello", []any{"OK"}, make(map[string]any)), "prefix_")
	mockInvocation := invocation.NewRPCInvocation("hello", []any{"OK"}, make(map[string]any))

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource: interfaceResourceName,
			// MetricType:             flow.QPS,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              100,
			RelationStrategy:       flow.CurrentResource,
		},
	})
	assert.NoError(t, err)

	wg := sync.WaitGroup{}
	wg.Add(10)
	f := &sentinelProviderFilter{}
	pass := int64(0)
	block := int64(0)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 30; j++ {
				result := f.Invoke(context.TODO(), mockInvoker, mockInvocation)
				if result.Error() == nil {
					atomic.AddInt64(&pass, 1)
				} else {
					atomic.AddInt64(&block, 1)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// todo sentinel can't assure the passed count is 100, sometimes is 101
	assert.True(t, atomic.LoadInt64(&pass) <= 105 && atomic.LoadInt64(&pass) >= 95)
	assert.True(t, atomic.LoadInt64(&block) <= 205 && atomic.LoadInt64(&block) >= 195)
}

type ErrInvoker struct {
	*base.BaseInvoker
}

func (ei *ErrInvoker) Invoke(context context.Context, invocation base.Invocation) result.Result {
	invoke := ei.BaseInvoker.Invoke(context, invocation)
	invoke.SetError(errors.New("error"))
	return invoke
}

type stateChangeTestListener struct {
	OnTransformToOpenChan chan struct{}
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot any) {
	s.OnTransformToOpenChan <- struct{}{}
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
}

func TestSentinelFilter_ErrorCount(t *testing.T) {
	url, err := common.NewURL("dubbo://127.0.0.1:20000/UserProvider?anyhost=true&" +
		"version=1.0.0&group=myGroup&" +
		"application=BDTService&category=providers&default.timeout=10000&dubbo=dubbo-provider-golang-1.0.0&" +
		"environment=dev&interface=com.test.user.UserProvider&ip=192.168.56.1&methods=GetUser%2C&" +
		"module=dubbogo+user-info+server&org=test.com&owner=ZX&pid=1447&revision=0.0.1&" +
		"side=provider&timeout=3000&timestamp=1556509797245&bean.name=UserProvider")
	assert.NoError(t, err)
	mockInvoker := &ErrInvoker{base.NewBaseInvoker(url)}
	_, methodResourceName := getResourceName(mockInvoker,
		invocation.NewRPCInvocation("hi", []any{"OK"}, make(map[string]any)), DefaultProviderPrefix)
	mockInvocation := invocation.NewRPCInvocation("hi", []any{"OK"}, make(map[string]any))

	// Register a state change listener so that we could observe the state change of the internal circuit breaker.
	listener := &stateChangeTestListener{}
	listener.OnTransformToOpenChan = make(chan struct{}, 1)
	circuitbreaker.RegisterStateChangeListeners(listener)
	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=0.9s, recoveryTimeout=3s, maxErrorCount=50
		{
			Resource:                     methodResourceName,
			Strategy:                     circuitbreaker.ErrorCount,
			RetryTimeoutMs:               3000,
			MinRequestAmount:             10,
			StatIntervalMs:               900,
			StatSlidingWindowBucketCount: 10,
			Threshold:                    50,
		},
	})
	assert.NoError(t, err)

	f := &sentinelProviderFilter{}
	for i := 0; i < 50; i++ {
		result := f.Invoke(context.TODO(), mockInvoker, mockInvocation)
		assert.Error(t, result.Error())
	}
	select {
	case <-time.After(time.Second):
		t.Error()
	case <-listener.OnTransformToOpenChan:
	}

}

func TestConsumerFilter_Invoke(t *testing.T) {
	f := &sentinelConsumerFilter{}
	url, err := common.NewURL("dubbo://127.0.0.1:20000/UserProvider?anyhost=true&" +
		"application=BDTService&category=providers&default.timeout=10000&dubbo=dubbo-provider-golang-1.0.0&" +
		"environment=dev&interface=com.ikurento.user.UserProvider&ip=192.168.56.1&methods=GetUser%2C&" +
		"module=dubbogo+user-info+server&org=ikurento.com&owner=ZX&pid=1447&revision=0.0.1&" +
		"side=provider&timeout=3000&timestamp=1556509797245&bean.name=UserProvider")
	assert.NoError(t, err)
	mockInvoker := base.NewBaseInvoker(url)
	mockInvocation := invocation.NewRPCInvocation("hello", []any{"OK"}, make(map[string]any))
	result := f.Invoke(context.TODO(), mockInvoker, mockInvocation)
	assert.NoError(t, result.Error())
}

func TestProviderFilter_Invoke(t *testing.T) {
	f := &sentinelProviderFilter{}
	url, err := common.NewURL("dubbo://127.0.0.1:20000/UserProvider?anyhost=true&" +
		"application=BDTService&category=providers&default.timeout=10000&dubbo=dubbo-provider-golang-1.0.0&" +
		"environment=dev&interface=com.ikurento.user.UserProvider&ip=192.168.56.1&methods=GetUser%2C&" +
		"module=dubbogo+user-info+server&org=ikurento.com&owner=ZX&pid=1447&revision=0.0.1&" +
		"side=provider&timeout=3000&timestamp=1556509797245&bean.name=UserProvider")
	assert.NoError(t, err)
	mockInvoker := base.NewBaseInvoker(url)
	mockInvocation := invocation.NewRPCInvocation("hello", []any{"OK"}, make(map[string]any))
	result := f.Invoke(context.TODO(), mockInvoker, mockInvocation)
	assert.NoError(t, result.Error())
}

func TestGetResourceName(t *testing.T) {
	url, err := common.NewURL("dubbo://127.0.0.1:20000/UserProvider?anyhost=true&" +
		"version=1.0.0&group=myGroup&" +
		"application=BDTService&category=providers&default.timeout=10000&dubbo=dubbo-provider-golang-1.0.0&" +
		"environment=dev&interface=com.ikurento.user.UserProvider&ip=192.168.56.1&methods=GetUser%2C&" +
		"module=dubbogo+user-info+server&org=ikurento.com&owner=ZX&pid=1447&revision=0.0.1&" +
		"side=provider&timeout=3000&timestamp=1556509797245&bean.name=UserProvider")
	assert.NoError(t, err)
	mockInvoker := base.NewBaseInvoker(url)
	interfaceResourceName, methodResourceName := getResourceName(mockInvoker,
		invocation.NewRPCInvocation("hello", []any{"OK"}, make(map[string]any)), "prefix_")
	assert.Equal(t, "com.ikurento.user.UserProvider:myGroup:1.0.0", interfaceResourceName)
	assert.Equal(t, "prefix_com.ikurento.user.UserProvider:myGroup:1.0.0:hello()", methodResourceName)
}
