package ratedplace

import (
	"context"
	"errors"
	"log"
	"time"

	"go-tour/businesses/places"
)

type RatedPlaceUseCase struct {
	repo RatedPlaceRepoInterface
	repoPlace places.PlaceRepoInterface
	ctx time.Duration
}

func NewUseCase(RatedplaceRepo RatedPlaceRepoInterface, repoPlace places.PlaceRepoInterface,  ctx time.Duration ) *RatedPlaceUseCase {
	return &RatedPlaceUseCase{
		repo: 		RatedplaceRepo,
		repoPlace: 	repoPlace,
		ctx:		ctx,
	}
}

func (usecase *RatedPlaceUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	ratedplace, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	latestRating, err := usecase.repoPlace.GetRating(domain.Place_ID, ctx)

	if err != nil {
		return Domain{}, err
	}

	totalUser, err := usecase.repo.Count(domain.Place_ID, ctx)
	if err != nil {
		return Domain{}, err
	}
	sumRating, err := usecase.repo.GetRating(domain.Place_ID, ctx)
	if err != nil {
		return Domain{}, err
	}
	log.Println("isi latestRating before", latestRating.Rating, domain.Rating , sumRating, totalUser)
	if latestRating.Rating == 0 {
		latestRating.Rating = domain.Rating
	} else {
		latestRating.Rating = sumRating / float64(totalUser)
	}

	log.Println("isi latestRating",latestRating.Rating, domain.Rating , sumRating, totalUser)

	_, err = usecase.repoPlace.GiveRate(domain.Place_ID, ctx, latestRating)
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