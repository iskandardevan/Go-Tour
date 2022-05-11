package places

import (
	"context"
	"errors"
	"go-tour/businesses/places"

	"gorm.io/gorm"
)

type placeRepo struct {
	DB *gorm.DB
}

func NewPlaceRepo(DB *gorm.DB) *placeRepo {
	return &placeRepo{DB: DB}
}

func (Repo *placeRepo) Add(ctx context.Context, domain places.Domain) (places.Domain, error) {
	place := Place{
		Author_ID: domain.Author_ID,
		Name: domain.Name,
		Location: domain.Location,
		Description: domain.Description,
		Rating: domain.Rating,
	}
	err := Repo.DB.Create(&place)
	if err.Error != nil {
		return places.Domain{}, err.Error
	}
	return domain, nil
}

func (Repo *placeRepo) GetAll(ctx context.Context) ([]places.Domain, error) {
	var place []Place
	err := Repo.DB.Preload("Author").Find(&place)
	if err.Error != nil {
		return []places.Domain{}, err.Error
	}
	return GetAll(place), nil
}
	
func (Repo *placeRepo) GetByID(id uint, ctx context.Context) (places.Domain, error) {
	var place Place
	err := Repo.DB.Preload("Author").Find(&place,"id=?", id)
	if err.Error != nil {
		return places.Domain{}, err.Error
	}
	return place.ToDomain(), nil
}

func (Repo *placeRepo) Update(id uint, ctx context.Context, domain places.Domain) (places.Domain, error) {
	place := FromDomain(domain)
	if Repo.DB.Save(&place).Error != nil {
		return places.Domain{}, errors.New("error updating place")
	}
	return place.ToDomain(), nil
}

func (Repo *placeRepo) Delete(id uint, ctx context.Context) error {
	place := Place{}
	err := Repo.DB.Where("id = ?", id).Delete(&place)
	if err.Error != nil {
		return err.Error
	}
	return nil
}