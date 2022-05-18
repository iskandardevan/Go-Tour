package placeimages

import (
	"context"
	"errors"
	"go-tour/businesses/placeimages"

	"gorm.io/gorm"
)

type placeimageRepo struct {
	DB *gorm.DB
}

func NewPlaceImageRepo(DB *gorm.DB) *placeimageRepo {
	return &placeimageRepo{DB: DB}
}



func (Repo *placeimageRepo) Add(ctx context.Context, domain placeimages.Domain) (placeimages.Domain, error) {
	placeimage := PlaceImage{ 
		Place_ID: domain.Place_ID,
		Img_url: domain.Img_url,
	}
	err := Repo.DB.Preload("Author").Preload("Place").Create(&placeimage)
	if err.Error != nil {
		return placeimages.Domain{}, err.Error
	}
	return domain, nil
}

func (Repo *placeimageRepo) GetAll(ctx context.Context) ([]placeimages.Domain, error) {
	var placeimage []PlaceImage
	err := Repo.DB.Find(&placeimage)
	if err.Error != nil {
		return []placeimages.Domain{}, err.Error
	}
	return GetAll(placeimage), nil
}
	
func (Repo *placeimageRepo) GetByID(id uint, ctx context.Context) (placeimages.Domain, error) {
	var placeimage PlaceImage
	err := Repo.DB.Preload("Author").Find(&placeimage,"id=?", id)
	if err.Error != nil {
		return placeimages.Domain{}, err.Error
	}
	return placeimage.ToDomain(), nil
}

func (Repo *placeimageRepo) Update(id uint, ctx context.Context, domain placeimages.Domain) (placeimages.Domain, error) {
	place := FromDomain(domain)
	if Repo.DB.Save(&place).Error != nil {
		return placeimages.Domain{}, errors.New("error updating place")
	}
	return place.ToDomain(), nil
}

func (Repo *placeimageRepo) Delete(id uint, ctx context.Context) error {
	placeimage := PlaceImage{}
	err := Repo.DB.Where("id = ?", id).Delete(&placeimage)
	if err.Error != nil {
		return err.Error
	}
	return nil
}