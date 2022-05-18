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
	rate := RatedPlace{}
	err := Repo.DB.Where("user_id = ? AND place_id = ?", domain.User_ID, domain.Place_ID).First(&rate).Error
	// if errors.Is(err, gorm.ErrRecordNotFound) {
	// 	rate.Rating = domain.Rating
	// 	rate.User_ID = domain.User_ID
	// 	rate.Place_ID = domain.Place_ID
	// } else {
	// 	rate.Rating = (rate.Rating + domain.Rating)
	// } 
	if err != nil &&  !errors.Is(err, gorm.ErrRecordNotFound){
		return ratedplace.Domain{}, err
	}
	
	if err == nil{
		return ratedplace.Domain{}, errors.New("place already rated")
	}

	rate.Rating = domain.Rating
	rate.User_ID = domain.User_ID
	rate.Place_ID = domain.Place_ID
	
	err = Repo.DB.Save(&rate).Error
	if err != nil {
		return ratedplace.Domain{}, err
	}
	
	return domain, nil
}

// count users rated places
func (Repo *ratedplaceRepo) Count(id uint,ctx context.Context) (int64, error) {
	var count int64
	err := Repo.DB.Model(&RatedPlace{}).Where("place_id = ?", id).Count(&count)
	if err.Error != nil {
		return 0, err.Error
	}
	return count, nil
}

// sum of all ratings by place
func (Repo *ratedplaceRepo) GetRating(id uint,ctx context.Context) (float64, error) {
	var rate float64
	err := Repo.DB.Raw("SELECT SUM(rating) FROM rated_places WHERE place_id = ?", id).Row().Scan(&rate)
	if err != nil {
		return 0, err
	}
	return rate, nil
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