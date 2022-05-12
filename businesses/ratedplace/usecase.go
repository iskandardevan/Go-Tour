package ratedplace 

import (
	"context"
	"errors" 
	"time"
 
)

type RatedPlaceUseCase struct {
	repo RatedPlaceRepoInterface
	ctx time.Duration
}

func NewUseCase(RatedplaceRepo RatedPlaceRepoInterface, ctx time.Duration ) *RatedPlaceUseCase {
	return &RatedPlaceUseCase{
		repo: 		RatedplaceRepo,
		ctx:		ctx,
	}
}

func (usecase *RatedPlaceUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	ratedplace, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return ratedplace, nil
}

func (usecase *RatedPlaceUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	ratedplace, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return ratedplace, nil
}

func (usecase *RatedPlaceUseCase) GetByID(id uint, ctx context.Context) (Domain, error) {
	ratedplace, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada ratedplace dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return ratedplace, nil
}

func (usecase *RatedPlaceUseCase) Update(id uint, ctx context.Context, domain Domain) (Domain, error) {
	domain.Id = (id)
	ratedplace, err := usecase.repo.Update(id, ctx, domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada ratedplace dengan ID tersebut")
	}
	return ratedplace, nil
}

func (usecase *RatedPlaceUseCase) Delete(id uint, ctx context.Context)error {
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada ratedplace dengan ID tersebut")
	}
	return nil
}