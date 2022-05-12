package ratedplace

import (
	"context"
	"go-tour/businesses/places"
	"go-tour/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint		`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	User_ID		uint
	User 		users.Domain
	Place_ID    uint 
	Place		places.Domain
	Rating 		float64
}

type RatedPlaceUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	Update(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

type RatedPlaceRepoInterface interface { 
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	Update(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

