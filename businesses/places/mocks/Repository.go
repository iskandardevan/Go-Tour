package mocks

import (
	"context"
	"go-tour/businesses/places"

	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) Add(ctx context.Context, domain places.Domain) (places.Domain, error) {
	ret := _m.Called(ctx, domain)
	var r0 places.Domain
	if rf, ok := ret.Get(0).(func(context.Context, places.Domain) places.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(places.Domain)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, places.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) GetByID (id uint, ctx context.Context) (places.Domain, error){
	ret := _m.Called(id, ctx)
	var r0 places.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context)places.Domain); ok{
		r0 = rf(id, ctx)
	} else {
		r0 = ret.Get(0).(places.Domain)
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context)error); ok{
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}


func (_m *Repository) Delete (id uint, ctx context.Context ) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}



func (_m *Repository) GetAll(ctx context.Context) ([]places.Domain, error) {
    ret := _m.Called(ctx)

    var r0 []places.Domain
    if rf, ok := ret.Get(0).(func(context.Context) []places.Domain); ok {
        r0 = rf(ctx)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).([]places.Domain)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(context.Context) error); ok {
        r1 = rf(ctx)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

func (_m *Repository) Update(id uint, ctx context.Context, domain places.Domain) (places.Domain, error) {
	ret := _m.Called(ctx, id, domain)
	var r0 places.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context, places.Domain) places.Domain); ok {
		r0 = rf(id, ctx, domain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(places.Domain)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context, places.Domain) error); ok {
		r1 = rf(id, ctx, domain)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) GetRating(id uint, ctx context.Context) (places.Domain, error) {
	ret := _m.Called(ctx, id)
	var r0 places.Domain
	if rf, ok := ret.Get(0).(func(uint, context.Context) places.Domain); ok {
		r0 = rf(id, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(places.Domain)
		}
	}
	var r1 error
	if rf, ok := ret.Get(1).(func(uint, context.Context) error); ok {
		r1 = rf(id, ctx)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

func (_m *Repository) GiveRate(id uint, ctx context.Context, domain places.Domain) (places.Domain, error) {
	ret := _m.Called(ctx, id, domain)
	var r0 error
	if rf, ok := ret.Get(0).(func(uint, context.Context, places.Domain) error); ok {
		r0 = rf(id, ctx, domain)
	} else {
		r0 = ret.Error(0)
	}
	var r1 places.Domain
	if rf, ok := ret.Get(1).(func(uint, context.Context, places.Domain) places.Domain); ok {
		r1 = rf(id, ctx, domain)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(places.Domain)
		}
	}
	return r1, r0
}
