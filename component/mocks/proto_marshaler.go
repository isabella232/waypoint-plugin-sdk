// Code generated by mockery v1.1.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
)

// ProtoMarshaler is an autogenerated mock type for the ProtoMarshaler type
type ProtoMarshaler struct {
	mock.Mock
}

// Proto provides a mock function with given fields:
func (_m *ProtoMarshaler) Proto() protoiface.MessageV1 {
	ret := _m.Called()

	var r0 protoiface.MessageV1
	if rf, ok := ret.Get(0).(func() protoiface.MessageV1); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protoiface.MessageV1)
		}
	}

	return r0
}
