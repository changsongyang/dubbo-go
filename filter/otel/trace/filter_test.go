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

package trace

import (
	"context"
	"reflect"
	"testing"
)

import (
	"github.com/golang/mock/gomock"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/protocol/base"
	"dubbo.apache.org/dubbo-go/v3/protocol/result"
)

type fields struct {
	Propagators    propagation.TextMapPropagator
	TracerProvider trace.TracerProvider
}
type args struct {
	ctx        context.Context
	result     result.Result
	invoker    base.Invoker
	protocol   base.Invocation
	invocation base.Invocation
}

// MockInvocation is a mock of Invocation interface
type MockInvocation struct {
	ctrl     *gomock.Controller
	recorder *MockInvocationMockRecorder
}

// MockInvocationMockRecorder is the mock recorder for MockInvocation
type MockInvocationMockRecorder struct {
	mock *MockInvocation
}

// NewMockInvocation creates a new mock instance
func NewMockInvocation(ctrl *gomock.Controller) *MockInvocation {
	mock := &MockInvocation{ctrl: ctrl}
	mock.recorder = &MockInvocationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvocation) EXPECT() *MockInvocationMockRecorder {
	return m.recorder
}

// MethodName mocks base method
func (m *MockInvocation) MethodName() string {
	ret := m.ctrl.Call(m, "MethodName")
	ret0, _ := ret[0].(string)
	return ret0
}

// MethodName indicates an expected call of MethodName
func (mr *MockInvocationMockRecorder) MethodName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MethodName", reflect.TypeOf((*MockInvocation)(nil).MethodName))
}

// ActualMethodName mocks base method
func (m *MockInvocation) ActualMethodName() string {
	ret := m.ctrl.Call(m, "ActualMethodName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ActualMethodName indicates an expected call of ActualMethodName
func (mr *MockInvocationMockRecorder) ActualMethodName() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActualMethodName", reflect.TypeOf((*MockInvocation)(nil).ActualMethodName))
}

// ParameterTypeNames mocks base method
func (m *MockInvocation) ParameterTypeNames() []string {
	ret := m.ctrl.Call(m, "ParameterTypeNames")
	ret0, _ := ret[0].([]string)
	return ret0
}

// ParameterTypeNames indicates an expected call of ParameterTypeNames
func (mr *MockInvocationMockRecorder) ParameterTypeNames() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParameterTypeNames", reflect.TypeOf((*MockInvocation)(nil).ParameterTypeNames))
}

// ParameterTypes mocks base method
func (m *MockInvocation) ParameterTypes() []reflect.Type {
	ret := m.ctrl.Call(m, "ParameterTypes")
	ret0, _ := ret[0].([]reflect.Type)
	return ret0
}

// ParameterTypes indicates an expected call of ParameterTypes
func (mr *MockInvocationMockRecorder) ParameterTypes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParameterTypes", reflect.TypeOf((*MockInvocation)(nil).ParameterTypes))
}

// ParameterValues mocks base method
func (m *MockInvocation) ParameterValues() []reflect.Value {
	ret := m.ctrl.Call(m, "ParameterValues")
	ret0, _ := ret[0].([]reflect.Value)
	return ret0
}

// ParameterValues indicates an expected call of ParameterValues
func (mr *MockInvocationMockRecorder) ParameterValues() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParameterValues", reflect.TypeOf((*MockInvocation)(nil).ParameterValues))
}

func (m *MockInvocation) ParameterRawValues() []any {
	ret := m.ctrl.Call(m, "ParameterRawValues")
	ret0, _ := ret[0].([]any)
	return ret0
}

// ParameterValues indicates an expected call of ParameterValues
func (mr *MockInvocationMockRecorder) ParameterRawValues() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParameterRawValues", reflect.TypeOf((*MockInvocation)(nil).ParameterRawValues))
}

// Arguments mocks base method
func (m *MockInvocation) Arguments() []any {
	ret := m.ctrl.Call(m, "Arguments")
	ret0, _ := ret[0].([]any)
	return ret0
}

// Arguments indicates an expected call of Arguments
func (mr *MockInvocationMockRecorder) Arguments() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Arguments", reflect.TypeOf((*MockInvocation)(nil).Arguments))
}

// Reply mocks base method
func (m *MockInvocation) Reply() any {
	ret := m.ctrl.Call(m, "Reply")
	ret0 := ret[0]
	return ret0
}

// Reply indicates an expected call of Reply
func (mr *MockInvocationMockRecorder) Reply() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reply", reflect.TypeOf((*MockInvocation)(nil).Reply))
}

// Invoker mocks base method
func (m *MockInvocation) Invoker() base.Invoker {
	ret := m.ctrl.Call(m, "Invoker")
	ret0, _ := ret[0].(base.Invoker)
	return ret0
}

// Invoker indicates an expected call of Invoker
func (mr *MockInvocationMockRecorder) Invoker() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoker", reflect.TypeOf((*MockInvocation)(nil).Invoker))
}

// IsGenericInvocation mocks base method
func (m *MockInvocation) IsGenericInvocation() bool {
	ret := m.ctrl.Call(m, "IsGenericInvocation")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsGenericInvocation indicates an expected call of IsGenericInvocation
func (mr *MockInvocationMockRecorder) IsGenericInvocation() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsGenericInvocation", reflect.TypeOf((*MockInvocation)(nil).IsGenericInvocation))
}

// Attachments mocks base method
func (m *MockInvocation) Attachments() map[string]any {
	ret := m.ctrl.Call(m, "Attachments")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// Attachments indicates an expected call of Attachments
func (mr *MockInvocationMockRecorder) Attachments() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attachments", reflect.TypeOf((*MockInvocation)(nil).Attachments))
}

// SetAttachment mocks base method
func (m *MockInvocation) SetAttachment(key string, value any) {
	m.ctrl.Call(m, "SetAttachment", key, value)
}

// SetAttachment indicates an expected call of SetAttachment
func (mr *MockInvocationMockRecorder) SetAttachment(key, value any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAttachment", reflect.TypeOf((*MockInvocation)(nil).SetAttachment), key, value)
}

// GetAttachment mocks base method
func (m *MockInvocation) GetAttachment(key string) (string, bool) {
	ret := m.ctrl.Call(m, "GetAttachment", key)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetAttachment indicates an expected call of GetAttachment
func (mr *MockInvocationMockRecorder) GetAttachment(key any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachment", reflect.TypeOf((*MockInvocation)(nil).GetAttachment), key)
}

// GetAttachmentInterface mocks base method
func (m *MockInvocation) GetAttachmentInterface(arg0 string) any {
	ret := m.ctrl.Call(m, "GetAttachmentInterface", arg0)
	ret0 := ret[0]
	return ret0
}

// GetAttachmentInterface indicates an expected call of GetAttachmentInterface
func (mr *MockInvocationMockRecorder) GetAttachmentInterface(arg0 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentInterface", reflect.TypeOf((*MockInvocation)(nil).GetAttachmentInterface), arg0)
}

// GetAttachmentWithDefaultValue mocks base method
func (m *MockInvocation) GetAttachmentWithDefaultValue(key, defaultValue string) string {
	ret := m.ctrl.Call(m, "GetAttachmentWithDefaultValue", key, defaultValue)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetAttachmentWithDefaultValue indicates an expected call of GetAttachmentWithDefaultValue
func (mr *MockInvocationMockRecorder) GetAttachmentWithDefaultValue(key, defaultValue any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentWithDefaultValue", reflect.TypeOf((*MockInvocation)(nil).GetAttachmentWithDefaultValue), key, defaultValue)
}

// GetAttachmentAsContext mocks base method
func (m *MockInvocation) GetAttachmentAsContext() context.Context {
	ret := m.ctrl.Call(m, "GetAttachmentAsContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

func (m *MockInvocation) MergeAttachmentFromContext(ctx context.Context) {
	m.ctrl.Call(m, "MergeAttachmentFromContext", reflect.TypeOf((*MockInvocation)(nil).GetAttachmentWithDefaultValue), ctx)
}

// GetAttachmentAsContext indicates an expected call of GetAttachmentAsContext
func (mr *MockInvocationMockRecorder) GetAttachmentAsContext() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentAsContext", reflect.TypeOf((*MockInvocation)(nil).GetAttachmentAsContext))
}

// Attributes mocks base method
func (m *MockInvocation) Attributes() map[string]any {
	ret := m.ctrl.Call(m, "Attributes")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// Attributes indicates an expected call of Attributes
func (mr *MockInvocationMockRecorder) Attributes() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attributes", reflect.TypeOf((*MockInvocation)(nil).Attributes))
}

// SetAttribute mocks base method
func (m *MockInvocation) SetAttribute(key string, value any) {
	m.ctrl.Call(m, "SetAttribute", key, value)
}

// SetAttribute indicates an expected call of SetAttribute
func (mr *MockInvocationMockRecorder) SetAttribute(key, value any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAttribute", reflect.TypeOf((*MockInvocation)(nil).SetAttribute), key, value)
}

// GetAttribute mocks base method
func (m *MockInvocation) GetAttribute(key string) (any, bool) {
	ret := m.ctrl.Call(m, "GetAttribute", key)
	ret0 := ret[0]
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetAttribute indicates an expected call of GetAttribute
func (mr *MockInvocationMockRecorder) GetAttribute(key any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttribute", reflect.TypeOf((*MockInvocation)(nil).GetAttribute), key)
}

// GetAttributeWithDefaultValue mocks base method
func (m *MockInvocation) GetAttributeWithDefaultValue(key string, defaultValue any) any {
	ret := m.ctrl.Call(m, "GetAttributeWithDefaultValue", key, defaultValue)
	ret0 := ret[0]
	return ret0
}

// GetAttributeWithDefaultValue indicates an expected call of GetAttributeWithDefaultValue
func (mr *MockInvocationMockRecorder) GetAttributeWithDefaultValue(key, defaultValue any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttributeWithDefaultValue", reflect.TypeOf((*MockInvocation)(nil).GetAttributeWithDefaultValue), key, defaultValue)
}

// MockResult is a mock of Result interface
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *MockResultMockRecorder
}

// MockResultMockRecorder is the mock recorder for MockResult
type MockResultMockRecorder struct {
	mock *MockResult
}

// NewMockResult creates a new mock instance
func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &MockResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResult) EXPECT() *MockResultMockRecorder {
	return m.recorder
}

// SetError mocks base method
func (m *MockResult) SetError(arg0 error) {
	m.ctrl.Call(m, "SetError", arg0)
}

// SetError indicates an expected call of SetError
func (mr *MockResultMockRecorder) SetError(arg0 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetError", reflect.TypeOf((*MockResult)(nil).SetError), arg0)
}

// Error mocks base method
func (m *MockResult) Error() error {
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error
func (mr *MockResultMockRecorder) Error() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockResult)(nil).Error))
}

// SetResult mocks base method
func (m *MockResult) SetResult(arg0 any) {
	m.ctrl.Call(m, "SetResult", arg0)
}

// SetResult indicates an expected call of SetResult
func (mr *MockResultMockRecorder) SetResult(arg0 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetResult", reflect.TypeOf((*MockResult)(nil).SetResult), arg0)
}

// Result mocks base method
func (m *MockResult) Result() any {
	ret := m.ctrl.Call(m, "Result")
	ret0 := ret[0]
	return ret0
}

// Result indicates an expected call of Result
func (mr *MockResultMockRecorder) Result() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Result", reflect.TypeOf((*MockResult)(nil).Result))
}

// SetAttachments mocks base method
func (m *MockResult) SetAttachments(arg0 map[string]any) {
	m.ctrl.Call(m, "SetAttachments", arg0)
}

// SetAttachments indicates an expected call of SetAttachments
func (mr *MockResultMockRecorder) SetAttachments(arg0 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAttachments", reflect.TypeOf((*MockResult)(nil).SetAttachments), arg0)
}

// Attachments mocks base method
func (m *MockResult) Attachments() map[string]any {
	ret := m.ctrl.Call(m, "Attachments")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// Attachments indicates an expected call of Attachments
func (mr *MockResultMockRecorder) Attachments() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attachments", reflect.TypeOf((*MockResult)(nil).Attachments))
}

// AddAttachment mocks base method
func (m *MockResult) AddAttachment(arg0 string, arg1 any) {
	m.ctrl.Call(m, "AddAttachment", arg0, arg1)
}

// AddAttachment indicates an expected call of AddAttachment
func (mr *MockResultMockRecorder) AddAttachment(arg0, arg1 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAttachment", reflect.TypeOf((*MockResult)(nil).AddAttachment), arg0, arg1)
}

// Attachment mocks base method
func (m *MockResult) Attachment(arg0 string, arg1 any) any {
	ret := m.ctrl.Call(m, "Attachment", arg0, arg1)
	ret0 := ret[0]
	return ret0
}

// Attachment indicates an expected call of Attachment
func (mr *MockResultMockRecorder) Attachment(arg0, arg1 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Attachment", reflect.TypeOf((*MockResult)(nil).Attachment), arg0, arg1)
}

// MockInvoker is a mock of Invoker interface
type MockInvoker struct {
	ctrl     *gomock.Controller
	recorder *MockInvokerMockRecorder
}

// MockInvokerMockRecorder is the mock recorder for MockInvoker
type MockInvokerMockRecorder struct {
	mock *MockInvoker
}

// NewMockInvoker creates a new mock instance
func NewMockInvoker(ctrl *gomock.Controller) *MockInvoker {
	mock := &MockInvoker{ctrl: ctrl}
	mock.recorder = &MockInvokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInvoker) EXPECT() *MockInvokerMockRecorder {
	return m.recorder
}

// GetURL mocks base method
func (m *MockInvoker) GetURL() *common.URL {
	ret := m.ctrl.Call(m, "GetURL")
	ret0, _ := ret[0].(*common.URL)
	return ret0
}

// GetURL indicates an expected call of GetURL
func (mr *MockInvokerMockRecorder) GetURL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockInvoker)(nil).GetURL))
}

// IsAvailable mocks base method
func (m *MockInvoker) IsAvailable() bool {
	ret := m.ctrl.Call(m, "IsAvailable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsAvailable indicates an expected call of IsAvailable
func (mr *MockInvokerMockRecorder) IsAvailable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAvailable", reflect.TypeOf((*MockInvoker)(nil).IsAvailable))
}

// Destroy mocks base method
func (m *MockInvoker) Destroy() {
	m.ctrl.Call(m, "Destroy")
}

// Destroy indicates an expected call of Destroy
func (mr *MockInvokerMockRecorder) Destroy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInvoker)(nil).Destroy))
}

// Invoke mocks base method
func (m *MockInvoker) Invoke(arg0 context.Context, arg1 base.Invocation) result.Result {
	ret := m.ctrl.Call(m, "Invoke", arg0, arg1)
	ret0, _ := ret[0].(result.Result)
	return ret0
}

// Invoke indicates an expected call of Invoke
func (mr *MockInvokerMockRecorder) Invoke(arg0, arg1 any) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockInvoker)(nil).Invoke), arg0, arg1)
}

func getFields() fields {
	return fields{
		Propagators:    otel.GetTextMapPropagator(),
		TracerProvider: otel.GetTracerProvider(),
	}
}

func Test_otelServerFilter_OnResponse(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		args   args
		want   result.Result
	}{
		{
			name:   "test",
			fields: getFields(),
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &otelServerFilter{
				Propagators:    tt.fields.Propagators,
				TracerProvider: tt.fields.TracerProvider,
			}
			if got := f.OnResponse(tt.args.ctx, tt.args.result, tt.args.invoker, tt.args.protocol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_otelClientFilter_OnResponse(t *testing.T) {
	tests := []struct {
		name   string
		fields fields
		args   args
		want   result.Result
	}{
		{
			name:   "test",
			fields: getFields(),
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &otelClientFilter{
				Propagators:    tt.fields.Propagators,
				TracerProvider: tt.fields.TracerProvider,
			}
			if got := f.OnResponse(tt.args.ctx, tt.args.result, tt.args.invoker, tt.args.protocol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OnResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_otelServerFilter_Invoke(t *testing.T) {
	ctrl := gomock.NewController(t)

	res := NewMockResult(ctrl)
	res.EXPECT().Error().Return(nil).AnyTimes()

	invoker := NewMockInvoker(ctrl)
	invoker.EXPECT().GetURL().Return(&common.URL{}).AnyTimes()
	invoker.EXPECT().Invoke(gomock.Any(), gomock.Any()).Return(res).AnyTimes()

	invocation := NewMockInvocation(ctrl)
	invocation.EXPECT().ActualMethodName().Return("oteldubbogo").AnyTimes()
	invocation.EXPECT().MethodName().Return("otel").AnyTimes()
	invocation.EXPECT().SetAttachment(gomock.Any(), gomock.Any()).Return().AnyTimes()
	invocation.EXPECT().Attachments().Return(map[string]any{}).AnyTimes()

	tests := []struct {
		name   string
		fields fields
		args   args
		want   result.Result
	}{
		{
			name:   "test",
			fields: getFields(),
			args: args{
				ctx:        context.Background(),
				invoker:    invoker,
				invocation: invocation,
			},
			want: res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &otelServerFilter{
				Propagators:    tt.fields.Propagators,
				TracerProvider: tt.fields.TracerProvider,
			}
			if got := f.Invoke(tt.args.ctx, tt.args.invoker, tt.args.invocation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invoke() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_otelClientFilter_Invoke(t *testing.T) {
	ctrl := gomock.NewController(t)

	res := NewMockResult(ctrl)
	res.EXPECT().Error().Return(nil).AnyTimes()

	invoker := NewMockInvoker(ctrl)
	invoker.EXPECT().GetURL().Return(&common.URL{}).AnyTimes()
	invoker.EXPECT().Invoke(gomock.Any(), gomock.Any()).Return(res).AnyTimes()

	invocation := NewMockInvocation(ctrl)
	invocation.EXPECT().ActualMethodName().Return("oteldubbogo").AnyTimes()
	invocation.EXPECT().MethodName().Return("otel").AnyTimes()
	invocation.EXPECT().SetAttachment(gomock.Any(), gomock.Any()).Return().AnyTimes()
	invocation.EXPECT().Attachments().Return(map[string]any{}).AnyTimes()

	tests := []struct {
		name   string
		fields fields
		args   args
		want   result.Result
	}{
		{
			name:   "test",
			fields: getFields(),
			args: args{
				ctx:        context.Background(),
				invoker:    invoker,
				invocation: invocation,
			},
			want: res,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &otelClientFilter{
				Propagators:    tt.fields.Propagators,
				TracerProvider: tt.fields.TracerProvider,
			}
			if got := f.Invoke(tt.args.ctx, tt.args.invoker, tt.args.invocation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invoke() = %v, want %v", got, tt.want)
			}
		})
	}
}
