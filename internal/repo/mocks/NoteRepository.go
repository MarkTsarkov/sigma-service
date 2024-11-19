// Code generated by mockery v2.48.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/marktsarkov/sigma-service/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// NoteRepository is an autogenerated mock type for the NoteRepository type
type NoteRepository struct {
	mock.Mock
}

type NoteRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *NoteRepository) EXPECT() *NoteRepository_Expecter {
	return &NoteRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, note
func (_m *NoteRepository) Create(ctx context.Context, note *entity.Note) (int64, error) {
	ret := _m.Called(ctx, note)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) (int64, error)); ok {
		return rf(ctx, note)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Note) int64); ok {
		r0 = rf(ctx, note)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Note) error); ok {
		r1 = rf(ctx, note)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type NoteRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - note *entity.Note
func (_e *NoteRepository_Expecter) Create(ctx interface{}, note interface{}) *NoteRepository_Create_Call {
	return &NoteRepository_Create_Call{Call: _e.mock.On("Create", ctx, note)}
}

func (_c *NoteRepository_Create_Call) Run(run func(ctx context.Context, note *entity.Note)) *NoteRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entity.Note))
	})
	return _c
}

func (_c *NoteRepository_Create_Call) Return(_a0 int64, _a1 error) *NoteRepository_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteRepository_Create_Call) RunAndReturn(run func(context.Context, *entity.Note) (int64, error)) *NoteRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *NoteRepository) GetById(ctx context.Context, id int64) (*entity.Note, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *entity.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*entity.Note, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *entity.Note); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Note)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NoteRepository_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type NoteRepository_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *NoteRepository_Expecter) GetById(ctx interface{}, id interface{}) *NoteRepository_GetById_Call {
	return &NoteRepository_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *NoteRepository_GetById_Call) Run(run func(ctx context.Context, id int64)) *NoteRepository_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *NoteRepository_GetById_Call) Return(_a0 *entity.Note, _a1 error) *NoteRepository_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NoteRepository_GetById_Call) RunAndReturn(run func(context.Context, int64) (*entity.Note, error)) *NoteRepository_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// NewNoteRepository creates a new instance of NoteRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNoteRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *NoteRepository {
	mock := &NoteRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
