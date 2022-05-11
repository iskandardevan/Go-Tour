package places

import (
	"context"
	"go-tour/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        	uint		`gorm:"primaryKey"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
	Author_ID   uint 
	Author		users.Domain
	Name 		string
	Location  	string
	Description     	string
	Rating float64
}

type PlaceUsecaseInterface interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	Update(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

type PlaceRepoInterface interface { 
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(id uint, ctx context.Context) (Domain, error)
	Update(id uint, ctx context.Context, domain Domain) (Domain, error)
	Delete(id uint, ctx context.Context)error
}

