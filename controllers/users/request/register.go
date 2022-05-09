package request

import "go-tour/businesses/users"

type RegisterUserRequest struct {
	Email    	string 	`json:"email"`
	Name     	string 	`json:"name"` 
	Password 	string 	`json:"password"` 
}

func (User *RegisterUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email	:User.Email,
		Name    :User.Name, 
		Password:User.Password, 
	}
}