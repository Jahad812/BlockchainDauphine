// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	mock "github.com/stretchr/testify/mock"

	vrfkey "github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/vrfkey"
)

// VRF is an autogenerated mock type for the VRF type
type VRF struct {
	mock.Mock
}

type VRF_Expecter struct {
	mock *mock.Mock
}

func (_m *VRF) EXPECT() *VRF_Expecter {
	return &VRF_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: ctx, key
func (_m *VRF) Add(ctx context.Context, key vrfkey.KeyV2) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, vrfkey.KeyV2) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VRF_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type VRF_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - ctx context.Context
//   - key vrfkey.KeyV2
func (_e *VRF_Expecter) Add(ctx interface{}, key interface{}) *VRF_Add_Call {
	return &VRF_Add_Call{Call: _e.mock.On("Add", ctx, key)}
}

func (_c *VRF_Add_Call) Run(run func(ctx context.Context, key vrfkey.KeyV2)) *VRF_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(vrfkey.KeyV2))
	})
	return _c
}

func (_c *VRF_Add_Call) Return(_a0 error) *VRF_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *VRF_Add_Call) RunAndReturn(run func(context.Context, vrfkey.KeyV2) error) *VRF_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: ctx
func (_m *VRF) Create(ctx context.Context) (vrfkey.KeyV2, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (vrfkey.KeyV2, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) vrfkey.KeyV2); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type VRF_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
func (_e *VRF_Expecter) Create(ctx interface{}) *VRF_Create_Call {
	return &VRF_Create_Call{Call: _e.mock.On("Create", ctx)}
}

func (_c *VRF_Create_Call) Run(run func(ctx context.Context)) *VRF_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *VRF_Create_Call) Return(_a0 vrfkey.KeyV2, _a1 error) *VRF_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_Create_Call) RunAndReturn(run func(context.Context) (vrfkey.KeyV2, error)) *VRF_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, id
func (_m *VRF) Delete(ctx context.Context, id string) (vrfkey.KeyV2, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (vrfkey.KeyV2, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) vrfkey.KeyV2); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type VRF_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *VRF_Expecter) Delete(ctx interface{}, id interface{}) *VRF_Delete_Call {
	return &VRF_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *VRF_Delete_Call) Run(run func(ctx context.Context, id string)) *VRF_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *VRF_Delete_Call) Return(_a0 vrfkey.KeyV2, _a1 error) *VRF_Delete_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_Delete_Call) RunAndReturn(run func(context.Context, string) (vrfkey.KeyV2, error)) *VRF_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Export provides a mock function with given fields: id, password
func (_m *VRF) Export(id string, password string) ([]byte, error) {
	ret := _m.Called(id, password)

	if len(ret) == 0 {
		panic("no return value specified for Export")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]byte, error)); ok {
		return rf(id, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(id, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(id, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_Export_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Export'
type VRF_Export_Call struct {
	*mock.Call
}

// Export is a helper method to define mock.On call
//   - id string
//   - password string
func (_e *VRF_Expecter) Export(id interface{}, password interface{}) *VRF_Export_Call {
	return &VRF_Export_Call{Call: _e.mock.On("Export", id, password)}
}

func (_c *VRF_Export_Call) Run(run func(id string, password string)) *VRF_Export_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *VRF_Export_Call) Return(_a0 []byte, _a1 error) *VRF_Export_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_Export_Call) RunAndReturn(run func(string, string) ([]byte, error)) *VRF_Export_Call {
	_c.Call.Return(run)
	return _c
}

// GenerateProof provides a mock function with given fields: id, seed
func (_m *VRF) GenerateProof(id string, seed *big.Int) (vrfkey.Proof, error) {
	ret := _m.Called(id, seed)

	if len(ret) == 0 {
		panic("no return value specified for GenerateProof")
	}

	var r0 vrfkey.Proof
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *big.Int) (vrfkey.Proof, error)); ok {
		return rf(id, seed)
	}
	if rf, ok := ret.Get(0).(func(string, *big.Int) vrfkey.Proof); ok {
		r0 = rf(id, seed)
	} else {
		r0 = ret.Get(0).(vrfkey.Proof)
	}

	if rf, ok := ret.Get(1).(func(string, *big.Int) error); ok {
		r1 = rf(id, seed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_GenerateProof_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateProof'
type VRF_GenerateProof_Call struct {
	*mock.Call
}

// GenerateProof is a helper method to define mock.On call
//   - id string
//   - seed *big.Int
func (_e *VRF_Expecter) GenerateProof(id interface{}, seed interface{}) *VRF_GenerateProof_Call {
	return &VRF_GenerateProof_Call{Call: _e.mock.On("GenerateProof", id, seed)}
}

func (_c *VRF_GenerateProof_Call) Run(run func(id string, seed *big.Int)) *VRF_GenerateProof_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*big.Int))
	})
	return _c
}

func (_c *VRF_GenerateProof_Call) Return(_a0 vrfkey.Proof, _a1 error) *VRF_GenerateProof_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_GenerateProof_Call) RunAndReturn(run func(string, *big.Int) (vrfkey.Proof, error)) *VRF_GenerateProof_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: id
func (_m *VRF) Get(id string) (vrfkey.KeyV2, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (vrfkey.KeyV2, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) vrfkey.KeyV2); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type VRF_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - id string
func (_e *VRF_Expecter) Get(id interface{}) *VRF_Get_Call {
	return &VRF_Get_Call{Call: _e.mock.On("Get", id)}
}

func (_c *VRF_Get_Call) Run(run func(id string)) *VRF_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *VRF_Get_Call) Return(_a0 vrfkey.KeyV2, _a1 error) *VRF_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_Get_Call) RunAndReturn(run func(string) (vrfkey.KeyV2, error)) *VRF_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with no fields
func (_m *VRF) GetAll() ([]vrfkey.KeyV2, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]vrfkey.KeyV2, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []vrfkey.KeyV2); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]vrfkey.KeyV2)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type VRF_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *VRF_Expecter) GetAll() *VRF_GetAll_Call {
	return &VRF_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *VRF_GetAll_Call) Run(run func()) *VRF_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *VRF_GetAll_Call) Return(_a0 []vrfkey.KeyV2, _a1 error) *VRF_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_GetAll_Call) RunAndReturn(run func() ([]vrfkey.KeyV2, error)) *VRF_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// Import provides a mock function with given fields: ctx, keyJSON, password
func (_m *VRF) Import(ctx context.Context, keyJSON []byte, password string) (vrfkey.KeyV2, error) {
	ret := _m.Called(ctx, keyJSON, password)

	if len(ret) == 0 {
		panic("no return value specified for Import")
	}

	var r0 vrfkey.KeyV2
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) (vrfkey.KeyV2, error)); ok {
		return rf(ctx, keyJSON, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []byte, string) vrfkey.KeyV2); ok {
		r0 = rf(ctx, keyJSON, password)
	} else {
		r0 = ret.Get(0).(vrfkey.KeyV2)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []byte, string) error); ok {
		r1 = rf(ctx, keyJSON, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VRF_Import_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Import'
type VRF_Import_Call struct {
	*mock.Call
}

// Import is a helper method to define mock.On call
//   - ctx context.Context
//   - keyJSON []byte
//   - password string
func (_e *VRF_Expecter) Import(ctx interface{}, keyJSON interface{}, password interface{}) *VRF_Import_Call {
	return &VRF_Import_Call{Call: _e.mock.On("Import", ctx, keyJSON, password)}
}

func (_c *VRF_Import_Call) Run(run func(ctx context.Context, keyJSON []byte, password string)) *VRF_Import_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]byte), args[2].(string))
	})
	return _c
}

func (_c *VRF_Import_Call) Return(_a0 vrfkey.KeyV2, _a1 error) *VRF_Import_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *VRF_Import_Call) RunAndReturn(run func(context.Context, []byte, string) (vrfkey.KeyV2, error)) *VRF_Import_Call {
	_c.Call.Return(run)
	return _c
}

// NewVRF creates a new instance of VRF. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVRF(t interface {
	mock.TestingT
	Cleanup(func())
}) *VRF {
	mock := &VRF{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
