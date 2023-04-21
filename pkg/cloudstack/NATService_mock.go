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
// Source: ./cloudstack/NATService.go

// Package cloudstack is a generated GoMock package.
package cloudstack

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNATServiceIface is a mock of NATServiceIface interface.
type MockNATServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockNATServiceIfaceMockRecorder
}

// MockNATServiceIfaceMockRecorder is the mock recorder for MockNATServiceIface.
type MockNATServiceIfaceMockRecorder struct {
	mock *MockNATServiceIface
}

// NewMockNATServiceIface creates a new mock instance.
func NewMockNATServiceIface(ctrl *gomock.Controller) *MockNATServiceIface {
	mock := &MockNATServiceIface{ctrl: ctrl}
	mock.recorder = &MockNATServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNATServiceIface) EXPECT() *MockNATServiceIfaceMockRecorder {
	return m.recorder
}

// CreateIpForwardingRule mocks base method.
func (m *MockNATServiceIface) CreateIpForwardingRule(p *CreateIpForwardingRuleParams) (*CreateIpForwardingRuleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIpForwardingRule", p)
	ret0, _ := ret[0].(*CreateIpForwardingRuleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIpForwardingRule indicates an expected call of CreateIpForwardingRule.
func (mr *MockNATServiceIfaceMockRecorder) CreateIpForwardingRule(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIpForwardingRule", reflect.TypeOf((*MockNATServiceIface)(nil).CreateIpForwardingRule), p)
}

// DeleteIpForwardingRule mocks base method.
func (m *MockNATServiceIface) DeleteIpForwardingRule(p *DeleteIpForwardingRuleParams) (*DeleteIpForwardingRuleResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIpForwardingRule", p)
	ret0, _ := ret[0].(*DeleteIpForwardingRuleResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteIpForwardingRule indicates an expected call of DeleteIpForwardingRule.
func (mr *MockNATServiceIfaceMockRecorder) DeleteIpForwardingRule(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIpForwardingRule", reflect.TypeOf((*MockNATServiceIface)(nil).DeleteIpForwardingRule), p)
}

// DisableStaticNat mocks base method.
func (m *MockNATServiceIface) DisableStaticNat(p *DisableStaticNatParams) (*DisableStaticNatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableStaticNat", p)
	ret0, _ := ret[0].(*DisableStaticNatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableStaticNat indicates an expected call of DisableStaticNat.
func (mr *MockNATServiceIfaceMockRecorder) DisableStaticNat(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableStaticNat", reflect.TypeOf((*MockNATServiceIface)(nil).DisableStaticNat), p)
}

// EnableStaticNat mocks base method.
func (m *MockNATServiceIface) EnableStaticNat(p *EnableStaticNatParams) (*EnableStaticNatResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableStaticNat", p)
	ret0, _ := ret[0].(*EnableStaticNatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableStaticNat indicates an expected call of EnableStaticNat.
func (mr *MockNATServiceIfaceMockRecorder) EnableStaticNat(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableStaticNat", reflect.TypeOf((*MockNATServiceIface)(nil).EnableStaticNat), p)
}

// GetIpForwardingRuleByID mocks base method.
func (m *MockNATServiceIface) GetIpForwardingRuleByID(id string, opts ...OptionFunc) (*IpForwardingRule, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetIpForwardingRuleByID", varargs...)
	ret0, _ := ret[0].(*IpForwardingRule)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetIpForwardingRuleByID indicates an expected call of GetIpForwardingRuleByID.
func (mr *MockNATServiceIfaceMockRecorder) GetIpForwardingRuleByID(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIpForwardingRuleByID", reflect.TypeOf((*MockNATServiceIface)(nil).GetIpForwardingRuleByID), varargs...)
}

// ListIpForwardingRules mocks base method.
func (m *MockNATServiceIface) ListIpForwardingRules(p *ListIpForwardingRulesParams) (*ListIpForwardingRulesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListIpForwardingRules", p)
	ret0, _ := ret[0].(*ListIpForwardingRulesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListIpForwardingRules indicates an expected call of ListIpForwardingRules.
func (mr *MockNATServiceIfaceMockRecorder) ListIpForwardingRules(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListIpForwardingRules", reflect.TypeOf((*MockNATServiceIface)(nil).ListIpForwardingRules), p)
}

// NewCreateIpForwardingRuleParams mocks base method.
func (m *MockNATServiceIface) NewCreateIpForwardingRuleParams(ipaddressid, protocol string, startport int) *CreateIpForwardingRuleParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewCreateIpForwardingRuleParams", ipaddressid, protocol, startport)
	ret0, _ := ret[0].(*CreateIpForwardingRuleParams)
	return ret0
}

// NewCreateIpForwardingRuleParams indicates an expected call of NewCreateIpForwardingRuleParams.
func (mr *MockNATServiceIfaceMockRecorder) NewCreateIpForwardingRuleParams(ipaddressid, protocol, startport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewCreateIpForwardingRuleParams", reflect.TypeOf((*MockNATServiceIface)(nil).NewCreateIpForwardingRuleParams), ipaddressid, protocol, startport)
}

// NewDeleteIpForwardingRuleParams mocks base method.
func (m *MockNATServiceIface) NewDeleteIpForwardingRuleParams(id string) *DeleteIpForwardingRuleParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDeleteIpForwardingRuleParams", id)
	ret0, _ := ret[0].(*DeleteIpForwardingRuleParams)
	return ret0
}

// NewDeleteIpForwardingRuleParams indicates an expected call of NewDeleteIpForwardingRuleParams.
func (mr *MockNATServiceIfaceMockRecorder) NewDeleteIpForwardingRuleParams(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDeleteIpForwardingRuleParams", reflect.TypeOf((*MockNATServiceIface)(nil).NewDeleteIpForwardingRuleParams), id)
}

// NewDisableStaticNatParams mocks base method.
func (m *MockNATServiceIface) NewDisableStaticNatParams(ipaddressid string) *DisableStaticNatParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDisableStaticNatParams", ipaddressid)
	ret0, _ := ret[0].(*DisableStaticNatParams)
	return ret0
}

// NewDisableStaticNatParams indicates an expected call of NewDisableStaticNatParams.
func (mr *MockNATServiceIfaceMockRecorder) NewDisableStaticNatParams(ipaddressid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDisableStaticNatParams", reflect.TypeOf((*MockNATServiceIface)(nil).NewDisableStaticNatParams), ipaddressid)
}

// NewEnableStaticNatParams mocks base method.
func (m *MockNATServiceIface) NewEnableStaticNatParams(ipaddressid, virtualmachineid string) *EnableStaticNatParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEnableStaticNatParams", ipaddressid, virtualmachineid)
	ret0, _ := ret[0].(*EnableStaticNatParams)
	return ret0
}

// NewEnableStaticNatParams indicates an expected call of NewEnableStaticNatParams.
func (mr *MockNATServiceIfaceMockRecorder) NewEnableStaticNatParams(ipaddressid, virtualmachineid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEnableStaticNatParams", reflect.TypeOf((*MockNATServiceIface)(nil).NewEnableStaticNatParams), ipaddressid, virtualmachineid)
}

// NewListIpForwardingRulesParams mocks base method.
func (m *MockNATServiceIface) NewListIpForwardingRulesParams() *ListIpForwardingRulesParams {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListIpForwardingRulesParams")
	ret0, _ := ret[0].(*ListIpForwardingRulesParams)
	return ret0
}

// NewListIpForwardingRulesParams indicates an expected call of NewListIpForwardingRulesParams.
func (mr *MockNATServiceIfaceMockRecorder) NewListIpForwardingRulesParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListIpForwardingRulesParams", reflect.TypeOf((*MockNATServiceIface)(nil).NewListIpForwardingRulesParams))
}
