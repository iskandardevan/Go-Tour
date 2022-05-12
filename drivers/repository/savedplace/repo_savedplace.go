package savedplace

import (
	"context"
	"errors"
	"go-tour/businesses/savedplace"

	"gorm.io/gorm"
)

type savedplaceRepo struct {
	DB *gorm.DB
}

func NewSavedPlaceRepo(DB *gorm.DB) *savedplaceRepo {
	return &savedplaceRepo{DB: DB}
}

func (Repo *savedplaceRepo) Add(ctx context.Context, domain savedplace.Domain) (savedplace.Domain, error) {
	savedplaces := SavedPlace{  
		User_ID: domain.User_ID,
		Place_ID: domain.Place_ID,
	}
	err := Repo.DB.Create(&savedplaces)
	if err.Error != nil {
		return savedplace.Domain{}, err.Error
	}
	return domain, nil
}

func (Repo *savedplaceRepo) GetAll(ctx context.Context) ([]savedplace.Domain, error) {
	var savedplaces []SavedPlace
	err := Repo.DB.Preload("User").Preload("Place").Find(&savedplaces)
	if err.Error != nil {
		return []savedplace.Domain{}, err.Error
	}
	return GetAll(savedplaces), nil
}
	
func (Repo *savedplaceRepo) GetByID(id uint, ctx context.Context) (savedplace.Domain, error) {
	var savedplaces SavedPlace
	err := Repo.DB.Preload("Author").Find(&savedplaces,"id=?", id)
	if err.Error != nil {
		return savedplace.Domain{}, err.Error
	}
	return savedplaces.ToDomain(), nil
}

func (Repo *savedplaceRepo) Update(id uint, ctx context.Context, domain savedplace.Domain) (savedplace.Domain, error) {
	place := FromDomain(domain)
	if Repo.DB.Save(&place).Error != nil {
		return savedplace.Domain{}, errors.New("error updating place")
	}
	return place.ToDomain(), nil
}

func (Repo *savedplaceRepo) Delete(id uint, ctx context.Context) error {
	savedplace := SavedPlace{}
	err := Repo.DB.Where("id = ?", id).Delete(&savedplace)
	if err.Error != nil {
		return err.Error
	}
	return nil
}