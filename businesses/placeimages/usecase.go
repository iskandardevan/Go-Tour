package placeimages

import (
	"context"
	"errors" 
	"time"
 
)

type PlaceImageUseCase struct {
	repo PlaceImageRepoInterface
	ctx time.Duration
}

func NewUseCase(placeImageRepo PlaceImageRepoInterface, ctx time.Duration ) *PlaceImageUseCase {
	return &PlaceImageUseCase{
		repo: 		placeImageRepo,
		ctx:		ctx,
	}
}

func (usecase *PlaceImageUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Img_url == "" {
		return Domain{}, errors.New("image belum di isi")
	}
	
	link, err := usecase.repo.Upload(domain.Img_url)
	
	if err != nil {
		return Domain{}, errors.New("gagal upload image")
	}
	domain.Img_url = link
	placeImage, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return placeImage, nil
}


func (usecase *PlaceImageUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	placeimages, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return placeimages, nil
}

func (usecase *PlaceImageUseCase) GetByID(id uint, ctx context.Context) (Domain, error) {
	placeimage, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada placeimage dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return placeimage, nil
}

func (usecase *PlaceImageUseCase) Update(id uint, ctx context.Context, domain Domain) (Domain, error) {
	domain.Id = (id)
	placeimage, err := usecase.repo.Update(id, ctx, domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada placeimage dengan ID tersebut")
	}
	return placeimage, nil
}

func (usecase *PlaceImageUseCase) Delete(id uint, ctx context.Context)error {
	err := usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada place dengan ID tersebut")
	}
	return nil
}