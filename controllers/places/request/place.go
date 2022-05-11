package request

import "go-tour/businesses/places"

type PlaceRequest struct {
	Author_ID   uint  `json:"author_id"`
	Name        string  `json:"name"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

func (req *PlaceRequest) ToDomain() *places.Domain {
	return &places.Domain{
		Author_ID:   req.Author_ID,
		Name:        req.Name,
		Location:    req.Location,
		Description: req.Description,
		Rating:      req.Rating,
	}
}