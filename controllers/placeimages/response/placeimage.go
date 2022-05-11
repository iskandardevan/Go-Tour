package response

import (
	"go-tour/businesses/placeimages" 
	// respPlace "go-tour/controllers/places/response"
	"time"

	"gorm.io/gorm"
)

type PlaceImageResponse struct {
	Id        	uint           `json:"id"`
	CreatedAt 	time.Time      `json:"createdAt"`
	UpdatedAt 	time.Time      `json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt `json:"deletedAt"` 
	Place_ID    uint           `json:"place_id"`
	// Place		respPlace.PlaceResponse
	Img_url 	string

}


func FromDomain(domain placeimages.Domain) PlaceImageResponse {
	return PlaceImageResponse{ 
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		Place_ID	:domain.Place_ID,
		// Place		:respPlace.FromDomain(domain.Place),
		Img_url		:domain.Img_url,
		
	}
}

func GetAllPlaceImages(domain []placeimages.Domain) []PlaceImageResponse {
	var GetAllPlaceImages []PlaceImageResponse
	for _, val := range domain{
		GetAllPlaceImages = append(GetAllPlaceImages, FromDomain(val))
	}
	return GetAllPlaceImages 
}