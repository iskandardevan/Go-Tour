package request

import "go-tour/businesses/placeimages"

type PlaceImageRequest struct {
	Place_ID   uint  `json:"place_id"`
	Img_url 	string  `json:"img_url"`
}

func (req *PlaceImageRequest) ToDomain() *placeimages.Domain {
	return &placeimages.Domain{
		Place_ID:    req.Place_ID,
		Img_url:     req.Img_url,
	}
}