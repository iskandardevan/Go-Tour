package response

import (
	"go-tour/businesses/ratedplace"
	respPlace "go-tour/controllers/places/response"
	respUser "go-tour/controllers/users/response"
	"time"

	"gorm.io/gorm"
)

type RatedPlaceResponse struct {
	Id        	uint           `json:"id"`
	CreatedAt 	time.Time      `json:"createdAt"`
	UpdatedAt 	time.Time      `json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt `json:"deletedAt"` 
	User_ID		uint           `json:"user_id"`
	User 		respUser.UserResponse `json:"user"`
	Place_ID    uint           `json:"place_id"`
	Place		respPlace.PlaceResponse 
	Rating	   	float64        `json:"rating"`


}


func FromDomain(domain ratedplace.Domain) RatedPlaceResponse {
	return RatedPlaceResponse{ 
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		User_ID		:domain.User_ID,
		User		:respUser.FromDomain(domain.User),
		Place_ID	:domain.Place_ID,
		Place		:respPlace.FromDomain(domain.Place),
		Rating		:domain.Rating,
		
	}
}

func GetAllRatedPlaces(domain []ratedplace.Domain) []RatedPlaceResponse {
	var GetAllRatedPlaces []RatedPlaceResponse
	for _, val := range domain{
		GetAllRatedPlaces = append(GetAllRatedPlaces, FromDomain(val))
	}
	return GetAllRatedPlaces 
}