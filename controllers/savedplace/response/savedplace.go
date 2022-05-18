package response

import (
	"go-tour/businesses/savedplace"
	respUser "go-tour/controllers/users/response"
	respPlace "go-tour/controllers/places/response"
	"time"

	"gorm.io/gorm"
)

type SavedPlaceResponse struct {
	Id        	uint           `json:"id"`
	CreatedAt 	time.Time      `json:"createdAt"`
	UpdatedAt 	time.Time      `json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt `json:"deletedAt"` 
	User_ID		uint           `json:"user_id"`
	User 		respUser.UserResponse `json:"user"`
	Place_ID    uint           `json:"place_id"`
	Place		respPlace.PlaceResponse 

}


func FromDomain(domain savedplace.Domain) SavedPlaceResponse {
	return SavedPlaceResponse{ 
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		User_ID		:domain.User_ID,
		User		:respUser.FromDomain(domain.User),
		Place_ID	:domain.Place_ID,
		Place		:respPlace.FromDomain(domain.Place),
		
	}
}

func GetAllSavedPlaces(domain []savedplace.Domain) []SavedPlaceResponse {
	var GetAllSavedPlaces []SavedPlaceResponse
	for _, val := range domain{
		GetAllSavedPlaces = append(GetAllSavedPlaces, FromDomain(val))
	}
	return GetAllSavedPlaces 
}