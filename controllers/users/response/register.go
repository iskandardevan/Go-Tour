package response

import (
	"go-tour/businesses/users"
)


func FromUserRegister(domain users.Domain) UserResponse {
	return UserResponse{
		Id			:domain.Id,
		CreatedAt	:domain.CreatedAt,
		UpdatedAt	:domain.UpdatedAt,
		DeletedAt	:domain.DeletedAt,
		Name		:domain.Name,
		Email		:domain.Email, 
	}
}