package response

import "go-tour/businesses/users"

func UserLogin(domain users.Domain, token string) JWTResponse{
	Response := UserResponse{
		Id:        domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		Name:      domain.Name,
		Email:     domain.Email, 
	}

	TokenResponse := JWTResponse{}
	TokenResponse.Token = token
	TokenResponse.User = Response
	return TokenResponse

}

