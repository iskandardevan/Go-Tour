package placeimages

import (
	"context"
	"go-tour/businesses/places"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint		`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Place_ID    uint 
	Place		places.Domain
	Img_url 	string 
}

type PlaceImageUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	Update(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

type PlaceImageRepoInterface interface { 
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	Update(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

