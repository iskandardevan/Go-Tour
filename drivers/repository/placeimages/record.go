package placeimages

import (
	"go-tour/businesses/placeimages"
	"go-tour/drivers/repository/places"
	"time"

	"gorm.io/gorm"
)

type PlaceImage struct {
	Id          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Place_ID    uint
	Place       places.Place `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Img_url     string
}

func (p *PlaceImage) ToDomain() placeimages.Domain {
	return placeimages.Domain{
		Id:          p.Id, 
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		Place_ID:    p.Place_ID,
		Place:       p.Place.ToDomain(),
		Img_url:     p.Img_url,
	}
}


func FromDomain(domain placeimages.Domain) PlaceImage {
	return PlaceImage{
		Id:          domain.Id,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		Place_ID:    domain.Place_ID,
		Place:       places.FromDomain(domain.Place),
		Img_url:     domain.Img_url,
	}
}

func GetAll(data []PlaceImage) []placeimages.Domain {
	res := []placeimages.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res
}