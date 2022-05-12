package savedplace

import (
	"go-tour/businesses/savedplace"
	"go-tour/drivers/repository/places"
	"go-tour/drivers/repository/users"
	"time"

	"gorm.io/gorm"
)

type SavedPlace struct {
	Id          uint 			`gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt 	`gorm:"index"`
	User_ID 	uint
	User 		users.User 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Place_ID    uint
	Place       places.Place 	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (p *SavedPlace) ToDomain() savedplace.Domain {
	return savedplace.Domain{
		Id:          p.Id, 
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		User_ID:     p.User_ID,
		User: 		 p.User.ToDomain(),	
		Place_ID:    p.Place_ID,
		Place:       p.Place.ToDomain(),
	}
}


func FromDomain(domain savedplace.Domain) SavedPlace {
	return SavedPlace{
		Id:          domain.Id,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		User_ID:     domain.User_ID,
		User:        users.FromDomain(domain.User),
		Place_ID:    domain.Place_ID,
		Place:       places.FromDomain(domain.Place), 
	}
}

func GetAll(data []SavedPlace) []savedplace.Domain {
	res := []savedplace.Domain{}
	for _, val := range data {
		res = append(res, val.ToDomain())
	}
	return res
}