package request

import (
	"go-tour/businesses/ratedplace"
)

type RatedPlaceRequest struct {
	User_ID 	uint  `json:"user_id"`
	Place_ID   	uint  `json:"place_id"`
	Rating	   	float64 `json:"rating"`
	
}

func (req *RatedPlaceRequest) ToDomain() *ratedplace.Domain {
	return &ratedplace.Domain{
		User_ID: 	 req.User_ID,
		Place_ID:    req.Place_ID,
		Rating: 	 req.Rating,
		
	}
}