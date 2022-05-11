package places

import (
	"go-tour/businesses/places"
	"go-tour/drivers/repository/users"
	"time"

	"gorm.io/gorm"
)

type Place struct {
	Id          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Author_ID   uint
	Author      users.User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string
	Location    string
	Description string
	Rating      float64
}

func (p *Place) ToDomain() places.Domain {
	return places.Domain{
		Id:          p.Id, 
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Author_ID:   p.Author_ID,
		Author: 	 p.Author.ToDomain(),
		Name:        p.Name,
		Location:    p.Location,
		Description: p.Description,
		Rating:      p.Rating,
	}
}

func FromDomain(domain places.Domain) Place {
	return Place{
		Id:          domain.Id, 
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		Author_ID:   domain.Author_ID,
		Author: 	 users.FromDomain(domain.Author),
		Name:        domain.Name,
		Location:    domain.Location,
		Description: domain.Description,
		Rating:      domain.Rating,
	}
}

func GetAll(data []Place) []places.Domain {
	res := []places.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res
}