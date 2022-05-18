package request

import ( 
	"go-tour/businesses/savedplace"
)

type SavedPlaceRequest struct {
	User_ID 	uint  `json:"user_id"`
	Place_ID   	uint  `json:"place_id"`
	
}

func (req *SavedPlaceRequest) ToDomain() *savedplace.Domain {
	return &savedplace.Domain{
		User_ID: 	 req.User_ID,
		Place_ID:    req.Place_ID,
		
	}
}