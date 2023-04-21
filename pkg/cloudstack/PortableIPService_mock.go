//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: ./cloudstack/PortableIPService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPortableIPServiceIface is a mock of PortableIPServiceIface interface.
type MockPortableIPServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockPortableIPServiceIfaceMockRecorder
}

// MockPortableIPServiceIfaceMockRecorder is the mock recorder for MockPortableIPServiceIface.
type MockPortableIPServiceIfaceMockRecorder struct {
	mock *MockPortableIPServiceIface
}

// NewMockPortableIPServiceIface creates a new mock instance.
func NewMockPortableIPServiceIface(ctrl *gomock.Controller) *MockPortableIPServiceIface {
	mock := &MockPortableIPServiceIface{ctrl: ctrl}
	mock.recorder = &MockPortableIPServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPortableIPServiceIface) EXPECT() *MockPortableIPServiceIfaceMockRecorder {
	return m.recorder
}

// CreatePortableIpRange mocks base method.
func (m *MockPortableIPServiceIface) CreatePortableIpRange(p *CreatePortableIpRangeParams) (*CreatePortableIpRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePortableIpRange", p)
	ret0, _ := ret[0].(*CreatePortableIpRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePortableIpRange indicates an expected call of CreatePortableIpRange.
func (mr *MockPortableIPServiceIfaceMockRecorder) CreatePortableIpRange(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePortableIpRange", reflect.TypeOf((*MockPortableIPServiceIface)(nil).CreatePortableIpRange), p)
}

// DeletePortableIpRange mocks base method.
func (m *MockPortableIPServiceIface) DeletePortableIpRange(p *DeletePortableIpRangeParams) (*DeletePortableIpRangeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePortableIpRange", p)
	ret0, _ := ret[0].(*DeletePortableIpRangeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePortableIpRange indicates an expected call of DeletePortableIpRange.
func (mr *MockPortableIPServiceIfaceMockRecorder) DeletePortableIpRange(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePortableIpRange", reflect.TypeOf((*MockPortableIPServiceIface)(nil).DeletePortableIpRange), p)
}

// GetPortableIpRangeByID mocks base method.
func (m *MockPortableIPServiceIface) GetPortableIpRangeByID(id string, opts ...OptionFunc) (*PortableIpRange, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPortableIpRangeByID", varargs...)
	ret0, _ := ret[0].(*PortableIpRange)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPortableIpRangeByID indicates an expected call of GetPortableIpRangeByID.
func (mr *MockPortableIPServiceIfaceMockRecorder) GetPortableIpRangeByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPortableIpRangeByID", reflect.TypeOf((*MockPortableIPServiceIface)(nil).GetPortableIpRangeByID), varargs...)
}

// ListPortableIpRanges mocks base method.
func (m *MockPortableIPServiceIface) ListPortableIpRanges(p *ListPortableIpRangesParams) (*ListPortableIpRangesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPortableIpRanges", p)
	ret0, _ := ret[0].(*ListPortableIpRangesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPortableIpRanges indicates an expected call of ListPortableIpRanges.
func (mr *MockPortableIPServiceIfaceMockRecorder) ListPortableIpRanges(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPortableIpRanges", reflect.TypeOf((*MockPortableIPServiceIface)(nil).ListPortableIpRanges), p)
}

// NewCreatePortableIpRangeParams mocks base method.
func (m *MockPortableIPServiceIface) NewCreatePortableIpRangeParams(endip, gateway, netmask string, regionid int, startip string) *CreatePortableIpRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreatePortableIpRangeParams", endip, gateway, netmask, regionid, startip)
	ret0, _ := ret[0].(*CreatePortableIpRangeParams)
	return ret0
}

// NewCreatePortableIpRangeParams indicates an expected call of NewCreatePortableIpRangeParams.
func (mr *MockPortableIPServiceIfaceMockRecorder) NewCreatePortableIpRangeParams(endip, gateway, netmask, regionid, startip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreatePortableIpRangeParams", reflect.TypeOf((*MockPortableIPServiceIface)(nil).NewCreatePortableIpRangeParams), endip, gateway, netmask, regionid, startip)
}

// NewDeletePortableIpRangeParams mocks base method.
func (m *MockPortableIPServiceIface) NewDeletePortableIpRangeParams(id string) *DeletePortableIpRangeParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeletePortableIpRangeParams", id)
	ret0, _ := ret[0].(*DeletePortableIpRangeParams)
	return ret0
}

// NewDeletePortableIpRangeParams indicates an expected call of NewDeletePortableIpRangeParams.
func (mr *MockPortableIPServiceIfaceMockRecorder) NewDeletePortableIpRangeParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeletePortableIpRangeParams", reflect.TypeOf((*MockPortableIPServiceIface)(nil).NewDeletePortableIpRangeParams), id)
}

// NewListPortableIpRangesParams mocks base method.
func (m *MockPortableIPServiceIface) NewListPortableIpRangesParams() *ListPortableIpRangesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPortableIpRangesParams")
	ret0, _ := ret[0].(*ListPortableIpRangesParams)
	return ret0
}

// NewListPortableIpRangesParams indicates an expected call of NewListPortableIpRangesParams.
func (mr *MockPortableIPServiceIfaceMockRecorder) NewListPortableIpRangesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPortableIpRangesParams", reflect.TypeOf((*MockPortableIPServiceIface)(nil).NewListPortableIpRangesParams))
}
