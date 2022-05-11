package response

import (
	"go-tour/businesses/places"
	"time"
	respUser "go-tour/controllers/users/response"
	"gorm.io/gorm"
)

type PlaceResponse struct {
	Id        	uint           `json:"id"`
	CreatedAt 	time.Time      `json:"createdAt"`
	UpdatedAt 	time.Time      `json:"updatedAt"`
	DeletedAt 	gorm.DeletedAt `json:"deletedAt"` 
	Author_ID   uint           `json:"author_id"`
	Author		respUser.UserResponse
	Name      	string         `json:"name"` 
	Location  	string         `json:"location"`
	Description	string         `json:"description"`
	Rating		float64        `json:"rating"`

}


func FromDomain(domain places.Domain) PlaceResponse {
	return PlaceResponse{
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		Author		:respUser.FromDomain(domain.Author),
		Name		:domain.Name,
		Location	:domain.Location,
		Description	:domain.Description,
		Rating		:domain.Rating,
	}
}

func GetAllPlaces(domain []places.Domain) []PlaceResponse {
	var GetAllPlaces []PlaceResponse
	for _, val := range domain{
		GetAllPlaces = append(GetAllPlaces, FromDomain(val))
	}
	return GetAllPlaces 
}