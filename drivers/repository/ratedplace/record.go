package ratedplace

import (
	"go-tour/businesses/ratedplace" 
	"go-tour/drivers/repository/places"
	"go-tour/drivers/repository/users"
	"time"

	"gorm.io/gorm"
)

type RatedPlace struct {
	Id          uint 			`gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt 	`gorm:"index"`
	User_ID 	uint
	User 		users.User 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Place_ID    uint
	Place       places.Place 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Rating	   	float64
}

func (p *RatedPlace) ToDomain() ratedplace.Domain {
	return ratedplace.Domain{
		Id:          p.Id, 
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		User_ID:     p.User_ID,
		User: 		 p.User.ToDomain(),	
		Place_ID:    p.Place_ID,
		Place:       p.Place.ToDomain(),
		Rating:      p.Rating,
	}
}


func FromDomain(domain ratedplace.Domain) RatedPlace {
	return RatedPlace{
		Id:          domain.Id,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		User_ID:     domain.User_ID,
		User:        users.FromDomain(domain.User),
		Place_ID:    domain.Place_ID,
		Place:       places.FromDomain(domain.Place), 
		Rating: 	 domain.Rating,
	}
}

func GetAll(data []RatedPlace) []ratedplace.Domain {
	res := []ratedplace.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res
}