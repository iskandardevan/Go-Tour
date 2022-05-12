package savedplace 

import (
	"context"
	"errors" 
	"time"
 
)

type SavedPlaceUseCase struct {
	repo SavedPlaceRepoInterface
	ctx time.Duration
}

func NewUseCase(savedplaceRepo SavedPlaceRepoInterface, ctx time.Duration ) *SavedPlaceUseCase {
	return &SavedPlaceUseCase{
		repo: 		savedplaceRepo,
		ctx:		ctx,
	}
}

func (usecase *SavedPlaceUseCase) Add(ctx context.Context, domain Domain) (Domain, error) {
	savedplace, err := usecase.repo.Add(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return savedplace, nil
}

func (usecase *SavedPlaceUseCase) GetAll(ctx context.Context) ([]Domain, error) {
	savedplaces, err := usecase.repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return savedplaces, nil
}

func (usecase *SavedPlaceUseCase) GetByID(id uint, ctx context.Context) (Domain, error) {
	savedplace, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, errors.New("tidak ada savedplace dengan ID tersebut")
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return savedplace, nil
}

func (usecase *SavedPlaceUseCase) Update(id uint, ctx context.Context, domain Domain) (Domain, error) {
	domain.Id = (id)
	savedplace, err := usecase.repo.Update(id, ctx, domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada savedplace dengan ID tersebut")
	}
	return savedplace, nil
}

func (usecase *SavedPlaceUseCase) Delete(id uint, ctx context.Context)( error ) {
	err :=  usecase.repo.Delete(id, ctx)
	if err != nil {
		return errors.New("tidak ada SavedPlace dengan ID tersebut")
	}
	return nil
}