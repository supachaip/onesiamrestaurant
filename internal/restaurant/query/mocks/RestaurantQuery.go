// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "one-siam-restaurant/internal/restaurant/dto"

	mock "github.com/stretchr/testify/mock"
)

// RestaurantQuery is an autogenerated mock type for the RestaurantQuery type
type RestaurantQuery struct {
	mock.Mock
}

type RestaurantQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *RestaurantQuery) EXPECT() *RestaurantQuery_Expecter {
	return &RestaurantQuery_Expecter{mock: &_m.Mock}
}

// CancelReservation provides a mock function with given fields: ctx, bookingID
func (_m *RestaurantQuery) CancelReservation(ctx context.Context, bookingID string) dto.CancelReservationInfo {
	ret := _m.Called(ctx, bookingID)

	if len(ret) == 0 {
		panic("no return value specified for CancelReservation")
	}

	var r0 dto.CancelReservationInfo
	if rf, ok := ret.Get(0).(func(context.Context, string) dto.CancelReservationInfo); ok {
		r0 = rf(ctx, bookingID)
	} else {
		r0 = ret.Get(0).(dto.CancelReservationInfo)
	}

	return r0
}

// RestaurantQuery_CancelReservation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CancelReservation'
type RestaurantQuery_CancelReservation_Call struct {
	*mock.Call
}

// CancelReservation is a helper method to define mock.On call
//   - ctx context.Context
//   - bookingID string
func (_e *RestaurantQuery_Expecter) CancelReservation(ctx interface{}, bookingID interface{}) *RestaurantQuery_CancelReservation_Call {
	return &RestaurantQuery_CancelReservation_Call{Call: _e.mock.On("CancelReservation", ctx, bookingID)}
}

func (_c *RestaurantQuery_CancelReservation_Call) Run(run func(ctx context.Context, bookingID string)) *RestaurantQuery_CancelReservation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RestaurantQuery_CancelReservation_Call) Return(_a0 dto.CancelReservationInfo) *RestaurantQuery_CancelReservation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RestaurantQuery_CancelReservation_Call) RunAndReturn(run func(context.Context, string) dto.CancelReservationInfo) *RestaurantQuery_CancelReservation_Call {
	_c.Call.Return(run)
	return _c
}

// InitializeRestaurant provides a mock function with given fields: ctx, tables
func (_m *RestaurantQuery) InitializeRestaurant(ctx context.Context, tables int) bool {
	ret := _m.Called(ctx, tables)

	if len(ret) == 0 {
		panic("no return value specified for InitializeRestaurant")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, int) bool); ok {
		r0 = rf(ctx, tables)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RestaurantQuery_InitializeRestaurant_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitializeRestaurant'
type RestaurantQuery_InitializeRestaurant_Call struct {
	*mock.Call
}

// InitializeRestaurant is a helper method to define mock.On call
//   - ctx context.Context
//   - tables int
func (_e *RestaurantQuery_Expecter) InitializeRestaurant(ctx interface{}, tables interface{}) *RestaurantQuery_InitializeRestaurant_Call {
	return &RestaurantQuery_InitializeRestaurant_Call{Call: _e.mock.On("InitializeRestaurant", ctx, tables)}
}

func (_c *RestaurantQuery_InitializeRestaurant_Call) Run(run func(ctx context.Context, tables int)) *RestaurantQuery_InitializeRestaurant_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *RestaurantQuery_InitializeRestaurant_Call) Return(_a0 bool) *RestaurantQuery_InitializeRestaurant_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RestaurantQuery_InitializeRestaurant_Call) RunAndReturn(run func(context.Context, int) bool) *RestaurantQuery_InitializeRestaurant_Call {
	_c.Call.Return(run)
	return _c
}

// IsInitialzed provides a mock function with given fields:
func (_m *RestaurantQuery) IsInitialzed() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsInitialzed")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RestaurantQuery_IsInitialzed_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsInitialzed'
type RestaurantQuery_IsInitialzed_Call struct {
	*mock.Call
}

// IsInitialzed is a helper method to define mock.On call
func (_e *RestaurantQuery_Expecter) IsInitialzed() *RestaurantQuery_IsInitialzed_Call {
	return &RestaurantQuery_IsInitialzed_Call{Call: _e.mock.On("IsInitialzed")}
}

func (_c *RestaurantQuery_IsInitialzed_Call) Run(run func()) *RestaurantQuery_IsInitialzed_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RestaurantQuery_IsInitialzed_Call) Return(_a0 bool) *RestaurantQuery_IsInitialzed_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RestaurantQuery_IsInitialzed_Call) RunAndReturn(run func() bool) *RestaurantQuery_IsInitialzed_Call {
	_c.Call.Return(run)
	return _c
}

// IsReserved provides a mock function with given fields: ctx, bookingID
func (_m *RestaurantQuery) IsReserved(ctx context.Context, bookingID string) bool {
	ret := _m.Called(ctx, bookingID)

	if len(ret) == 0 {
		panic("no return value specified for IsReserved")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, bookingID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RestaurantQuery_IsReserved_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsReserved'
type RestaurantQuery_IsReserved_Call struct {
	*mock.Call
}

// IsReserved is a helper method to define mock.On call
//   - ctx context.Context
//   - bookingID string
func (_e *RestaurantQuery_Expecter) IsReserved(ctx interface{}, bookingID interface{}) *RestaurantQuery_IsReserved_Call {
	return &RestaurantQuery_IsReserved_Call{Call: _e.mock.On("IsReserved", ctx, bookingID)}
}

func (_c *RestaurantQuery_IsReserved_Call) Run(run func(ctx context.Context, bookingID string)) *RestaurantQuery_IsReserved_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *RestaurantQuery_IsReserved_Call) Return(_a0 bool) *RestaurantQuery_IsReserved_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RestaurantQuery_IsReserved_Call) RunAndReturn(run func(context.Context, string) bool) *RestaurantQuery_IsReserved_Call {
	_c.Call.Return(run)
	return _c
}

// ReserveTable provides a mock function with given fields: ctx, numCustomers
func (_m *RestaurantQuery) ReserveTable(ctx context.Context, numCustomers int) dto.ReserveInfo {
	ret := _m.Called(ctx, numCustomers)

	if len(ret) == 0 {
		panic("no return value specified for ReserveTable")
	}

	var r0 dto.ReserveInfo
	if rf, ok := ret.Get(0).(func(context.Context, int) dto.ReserveInfo); ok {
		r0 = rf(ctx, numCustomers)
	} else {
		r0 = ret.Get(0).(dto.ReserveInfo)
	}

	return r0
}

// RestaurantQuery_ReserveTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReserveTable'
type RestaurantQuery_ReserveTable_Call struct {
	*mock.Call
}

// ReserveTable is a helper method to define mock.On call
//   - ctx context.Context
//   - numCustomers int
func (_e *RestaurantQuery_Expecter) ReserveTable(ctx interface{}, numCustomers interface{}) *RestaurantQuery_ReserveTable_Call {
	return &RestaurantQuery_ReserveTable_Call{Call: _e.mock.On("ReserveTable", ctx, numCustomers)}
}

func (_c *RestaurantQuery_ReserveTable_Call) Run(run func(ctx context.Context, numCustomers int)) *RestaurantQuery_ReserveTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *RestaurantQuery_ReserveTable_Call) Return(_a0 dto.ReserveInfo) *RestaurantQuery_ReserveTable_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RestaurantQuery_ReserveTable_Call) RunAndReturn(run func(context.Context, int) dto.ReserveInfo) *RestaurantQuery_ReserveTable_Call {
	_c.Call.Return(run)
	return _c
}

// NewRestaurantQuery creates a new instance of RestaurantQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRestaurantQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *RestaurantQuery {
	mock := &RestaurantQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}