package users

import (
	"go-tour/businesses/users"
	"time"

	"gorm.io/gorm"
)


type User struct {
	Id 			uint 			`gorm:"primaryKey"`
	Email		string			`gorm:"unique"`
	Name 		string 		 
	Password	string 
	Token		string
	CreatedAt 	time.Time   
	UpdatedAt 	time.Time   
	DeletedAt	gorm.DeletedAt 	`gorm:"index"`
}

func (user *User) ToDomain() users.Domain {
	return users.Domain{
		Id 			:user.Id,
		Email		:user.Email,
		Name 		:user.Name, 
		Password	:user.Password, 
		Token		:user.Token,
		CreatedAt 	:user.CreatedAt,
		UpdatedAt 	:user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) User  {
	return User{
		Id 			:domain.Id,
		Email		:domain.Email,
		Name 		:domain.Name, 
		Password	:domain.Password, 
		Token		:domain.Token,
		CreatedAt 	:domain.CreatedAt,
		UpdatedAt 	:domain.UpdatedAt,
	}

}

func GetAllUsers(data []User) []users.Domain{
	res := []users.Domain{}
	for _, val := range data{
		res = append(res, val.ToDomain())
	}
	return res
}