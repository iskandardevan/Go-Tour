package mocks

import (
	"context"
	"go-tour/businesses/users"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) GetByID (id uint, ctx context.Context) (users.Domain, error){
	ret := _m.Called(ctx, id)
	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context)users.Domain); ok{
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context)error); ok{
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}
