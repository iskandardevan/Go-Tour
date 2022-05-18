package places_test

import (
	"context"
	"errors"
	"go-tour/businesses/places"
	"go-tour/businesses/places/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var placeRepository = mocks.Repository{Mock: mock.Mock{}}
var placeService places.PlaceUsecaseInterface
var placeDomain places.Domain

var allplaceDomain []places.Domain

func setup() {
	placeService = places.NewUseCase(&placeRepository, time.Hour*24)
	placeDomain = places.Domain{
		Author_ID: 1,
		Name:       "Place 1",
		Description: "Place 1 description",
		Location:    "Place 1 location",
		Rating: 5,
	}
}

func TestAdd(t *testing.T){
	setup()
	placeRepository.On("Add", mock.Anything, placeDomain).Return(placeDomain, nil)
	placeRepository.On(mock.Anything, mock.Anything).Return(placeDomain, nil)
	t.Run("Add place success", func(t *testing.T) {
		placeDomain, err := placeService.Add(context.Background(), placeDomain)
		assert.Nil(t, err)
		assert.Equal(t, placeDomain, placeDomain)
	})
	t.Run("No Name", func(t *testing.T) {
		placeRepository.On("Add", mock.Anything, placeDomain).Return(placeDomain, errors.New("nama belum di isi")).Once()
		place, err := placeService.Add(context.Background(), places.Domain{
			Name: "",
		})
		assert.Error(t, err)
		assert.NotNil(t, place)
		})
	t.Run("No Location", func(t *testing.T) {
		placeRepository.On("Add", mock.Anything, placeDomain).Return(placeDomain, errors.New("lokasi belum di isi")).Once()
		place, err := placeService.Add(context.Background(), places.Domain{
			Location: "",
		})
		assert.Error(t, err)
		assert.NotNil(t, place)
		}) 
	t.Run("No Description", func(t *testing.T) {
		placeRepository.On("Add", mock.Anything, placeDomain).Return(placeDomain, errors.New("deskripsi belum di isi")).Once()
		place, err := placeService.Add(context.Background(), places.Domain{
			Description: "",
		})
		assert.Error(t, err)
		assert.NotNil(t, place)
		}) 
	t.Run("Add place error", func(t *testing.T) {
		placeDomain, err := placeService.Add(context.Background(), placeDomain)
		assert.Nil(t, err)
		assert.Equal(t, placeDomain, placeDomain)
	})
}

func TestGetAll(t *testing.T){
	setup()
	placeRepository.On("GetAll", mock.Anything).Return(allplaceDomain, nil)
	t.Run("Get all place success", func(t *testing.T) {
		place, err := placeService.GetAll(context.Background())
		assert.NoError(t, err)
        assert.Nil(t, place)
        assert.Equal(t, len(place), len(allplaceDomain))
	})
	t.Run("Get all place error", func(t *testing.T) {
		place, err := placeService.GetAll(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, place, []places.Domain(nil))
	})
}

func TestGetByID(t *testing.T){
	setup()
	t.Run("Get place by id success", func(t *testing.T) {
		placeRepository.On("GetByID", mock.AnythingOfType("uint"), mock.Anything).Return(placeDomain, nil).Once()
		place, err := placeService.GetByID(placeDomain.Id, context.Background())
		assert.Error(t, err) 
		assert.NotNil(t, place)
	})
	t.Run("No ID", func(t *testing.T) {
		placeRepository.On("GetByID", mock.AnythingOfType("uint"), mock.Anything).Return(placeDomain, errors.New("ID harus diisi")).Once()
		place, err := placeService.GetByID(0, context.Background())
		assert.Error(t, err)
		assert.NotNil(t, place)
		}) 
	t.Run("Nil", func(t *testing.T) {
		placeRepository.On("GetByID", mock.AnythingOfType("uint"), mock.Anything).Return(placeDomain, errors.New("tidak ada place dengan ID tersebut")).Once()
		place, err := placeService.GetByID(0, context.Background())
		assert.Error(t, err)
		assert.NotNil(t, place)
		assert.Equal(t, place, places.Domain{})
		}) 
}

// func TestUpdate(t *testing.T){
// 	t.Run("Update place success", func(t *testing.T) {
// 		setup()
// 		placeRepository.On("Update", mock.Anything, placeDomain).Return(placeDomain, nil).Once()
// 		placeRepository.On(mock.Anything, mock.Anything).Return(placeDomain, nil)
// 		place, err := placeService.Update(1, context.Background(), placeDomain)
// 		assert.Nil(t, err)
// 		assert.Equal(t, place, placeDomain)
// 	})
// }