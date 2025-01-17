// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	reflect "reflect"
	time "time"

	core "github.com/Max425/wbschool_exam_L2/tree/main/develop/dev11/pkg/model/core"
	gomock "github.com/golang/mock/gomock"
)

// MockEvent is a mock of Event interface.
type MockEvent struct {
	ctrl     *gomock.Controller
	recorder *MockEventMockRecorder
}

// MockEventMockRecorder is the mock recorder for MockEvent.
type MockEventMockRecorder struct {
	mock *MockEvent
}

// NewMockEvent creates a new mock instance.
func NewMockEvent(ctrl *gomock.Controller) *MockEvent {
	mock := &MockEvent{ctrl: ctrl}
	mock.recorder = &MockEventMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvent) EXPECT() *MockEventMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockEvent) CreateEvent(event *core.Event) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", event)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockEventMockRecorder) CreateEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockEvent)(nil).CreateEvent), event)
}

// DeleteEvent mocks base method.
func (m *MockEvent) DeleteEvent(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockEventMockRecorder) DeleteEvent(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockEvent)(nil).DeleteEvent), id)
}

// GetEventsForTime mocks base method.
func (m *MockEvent) GetEventsForTime(date time.Duration) ([]core.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEventsForTime", date)
	ret0, _ := ret[0].([]core.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEventsForTime indicates an expected call of GetEventsForTime.
func (mr *MockEventMockRecorder) GetEventsForTime(date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEventsForTime", reflect.TypeOf((*MockEvent)(nil).GetEventsForTime), date)
}

// UpdateEvent mocks base method.
func (m *MockEvent) UpdateEvent(event *core.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", event)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockEventMockRecorder) UpdateEvent(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockEvent)(nil).UpdateEvent), event)
}
