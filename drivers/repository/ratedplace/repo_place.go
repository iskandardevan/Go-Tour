package ratedplace

import (
	"context"
	"errors"
	"go-tour/businesses/ratedplace" 

	"gorm.io/gorm"
)

type ratedplaceRepo struct {
	DB *gorm.DB
}

func NewRatedPlaceRepo(DB *gorm.DB) *ratedplaceRepo {
	return &ratedplaceRepo{DB: DB}
}

func (Repo *ratedplaceRepo) Add(ctx context.Context, domain ratedplace.Domain) (ratedplace.Domain, error) {
	ratedplaces := RatedPlace{  
		User_ID: domain.User_ID,
		Place_ID: domain.Place_ID,
	}
	err := Repo.DB.Create(&ratedplaces)
	if err.Error != nil {
		return ratedplace.Domain{}, err.Error
	}
	return domain, nil
}

func (Repo *ratedplaceRepo) GetAll(ctx context.Context) ([]ratedplace.Domain, error) {
	var ratedplaces []RatedPlace
	err := Repo.DB.Preload("User").Preload("Place").Find(&ratedplaces)
	if err.Error != nil {
		return []ratedplace.Domain{}, err.Error
	}
	return GetAll(ratedplaces), nil
}
	
func (Repo *ratedplaceRepo) GetByID(id uint, ctx context.Context) (ratedplace.Domain, error) {
	var ratedplaces RatedPlace
	err := Repo.DB.Preload("User").Find(&ratedplaces,"id=?", id)
	if err.Error != nil {
		return ratedplace.Domain{}, err.Error
	}
	return ratedplaces.ToDomain(), nil
}

func (Repo *ratedplaceRepo) Update(id uint, ctx context.Context, domain ratedplace.Domain) (ratedplace.Domain, error) {
	place := FromDomain(domain)
	if Repo.DB.Save(&place).Error != nil {
		return ratedplace.Domain{}, errors.New("error updating place")
	}
	return place.ToDomain(), nil
}

func (Repo *ratedplaceRepo) Delete(id uint, ctx context.Context) error {
	ratedplace := RatedPlace{}
	err := Repo.DB.Where("id = ?", id).Delete(&ratedplace)
	if err.Error != nil {
		return err.Error
	}
	return nil
}