// Code generated by mockery v2.51.0. DO NOT EDIT.

package dao

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/goharbor/harbor/src/pkg/repository/model"

	q "github.com/goharbor/harbor/src/lib/q"
)

// DAO is an autogenerated mock type for the DAO type
type DAO struct {
	mock.Mock
}

// AddPullCount provides a mock function with given fields: ctx, id, count
func (_m *DAO) AddPullCount(ctx context.Context, id int64, count uint64) error {
	ret := _m.Called(ctx, id, count)

	if len(ret) == 0 {
		panic("no return value specified for AddPullCount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, uint64) error); ok {
		r0 = rf(ctx, id, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Count provides a mock function with given fields: ctx, query
func (_m *DAO) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, repository
func (_m *DAO) Create(ctx context.Context, repository *model.RepoRecord) (int64, error) {
	ret := _m.Called(ctx, repository)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.RepoRecord) (int64, error)); ok {
		return rf(ctx, repository)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.RepoRecord) int64); ok {
		r0 = rf(ctx, repository)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.RepoRecord) error); ok {
		r1 = rf(ctx, repository)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DAO) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, id
func (_m *DAO) Get(ctx context.Context, id int64) (*model.RepoRecord, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.RepoRecord
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*model.RepoRecord, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.RepoRecord); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.RepoRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *DAO) List(ctx context.Context, query *q.Query) ([]*model.RepoRecord, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*model.RepoRecord
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*model.RepoRecord, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*model.RepoRecord); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RepoRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NonEmptyRepos provides a mock function with given fields: ctx
func (_m *DAO) NonEmptyRepos(ctx context.Context) ([]*model.RepoRecord, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for NonEmptyRepos")
	}

	var r0 []*model.RepoRecord
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.RepoRecord, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.RepoRecord); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.RepoRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, repository, props
func (_m *DAO) Update(ctx context.Context, repository *model.RepoRecord, props ...string) error {
	_va := make([]interface{}, len(props))
	for _i := range props {
		_va[_i] = props[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, repository)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.RepoRecord, ...string) error); ok {
		r0 = rf(ctx, repository, props...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDAO creates a new instance of DAO. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDAO(t interface {
	mock.TestingT
	Cleanup(func())
}) *DAO {
	mock := &DAO{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
