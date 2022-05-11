package places

import (
	"context"
	"errors" 
	"time"
 
)

type PlaceUseCase struct {
	repo PlaceRepoInterface
	ctx time.Duration  
}

func NewUseCase(placeRepo PlaceRepoInterface, ctx time.Duration ) *PlaceUseCase {
	return &PlaceUseCase{
		repo: 		placeRepo,
		ctx:		ctx,
	}
}

func (usecase *PlaceUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("nama belum di isi")
	}
	if domain.Location == "" {
		return Domain{}, errors.New("lokasi belum di isi")
	}
	if domain.Description == "" {
		return Domain{}, errors.New("deskripsi belum di isi")
	} 
	place, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return place, nil
}

func (usecase *PlaceUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	places, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return places, nil
}

func (usecase *PlaceUseCase) GetByID(id uint, ctx context.Context) (Domain, error) {
	place, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada place dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return place, nil
}

func (usecase *PlaceUseCase) Update(id uint, ctx context.Context, domain Domain) (Domain, error) {
	domain.Id = (id)
	place, err := usecase.repo.Update(id, ctx, domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada place dengan ID tersebut")
	}
	return place, nil
}

func (usecase *PlaceUseCase) Delete(id uint, ctx context.Context)error {
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada place dengan ID tersebut")
	}
	return nil
}