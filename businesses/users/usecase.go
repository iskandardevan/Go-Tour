package users

import (
	"context"
	"errors"
	"log"
	"time"

	"go-tour/app/middlewares"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo UserRepoInterface
	ctx time.Duration 
	JWTAuth *middlewares.ConfigJWT
}



func NewUseCase(userRepo UserRepoInterface, ctx time.Duration, JWTAuth2 *middlewares.ConfigJWT) *UserUseCase {
	return &UserUseCase{
		repo: 		userRepo,
		ctx:		ctx,
		JWTAuth: 	JWTAuth2,
	}
}

func (usecase *UserUseCase) RegisterUser(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email belum di isi")
	}
	if domain.Name == "" {
		return Domain{}, errors.New("nama belum di isi")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("password belum di isi")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.DefaultCost)
	if err != nil {
		return Domain{}, err
	}
	domain.Password = string(hashedPassword)

	user, err := usecase.repo.RegisterUser(ctx, &domain)
	if err != nil {
		return Domain{}, err
	}
	
	return user, nil
}

func (usecase *UserUseCase) LoginUser(email string, password string, ctx context.Context) (Domain, string, error){
	
	if email == "" {
		return Domain{},"", errors.New("email belum di isi")
	}
	if password == "" {
		return Domain{},"", errors.New("password belum di isi")
	
	}

	user, err := usecase.repo.GetEmail(ctx, email)
	if err != nil {
		return Domain{},"", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return Domain{},"", err
	}

	token, errToken := usecase.JWTAuth.GenerateTokenJWT(user.Id, user.Email, user.Name)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, "", errors.New("token kosong")
	}
	return user, token, nil
}

func (usecase *UserUseCase) CheckingUser(email string, password string, ctx context.Context) (Domain, error){
	if email == "" {
		return Domain{}, errors.New("email belum di isi")
	}
	if password == "" {
		return Domain{}, errors.New("password belum di isi")
	
	}
	user, err := usecase.repo.CheckUser(email, ctx)
	if err != nil {
		return Domain{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return Domain{}, err
	}
	return user,  nil
}


func (usecase *UserUseCase) GetByEmail(email string, ctx context.Context) (Domain, error){
	if email == "" {
		return Domain{}, errors.New("email belum di isi")
	}
	user, err := usecase.repo.GetByEmail(email, ctx)
	if err != nil {
		return Domain{}, err
	}
	return user,  nil
}

func (usecase *UserUseCase) GetByID(id uint, ctx context.Context) (Domain, error){
	user, err := usecase.repo.GetByID(id, ctx)
	if err != nil {
		return Domain{}, err
	}
	if id == 0 {
		return Domain{}, errors.New("ID harus diisi")
	}
	return user, nil
}

func (usecase *UserUseCase) GetAllUsers(ctx context.Context) ([]Domain, error){
	user, err := usecase.repo.GetAllUsers(ctx)
	if err != nil {
		return []Domain{}, errors.New("tidak ada pengguna")
	}
	return user, nil
}

func (usecase *UserUseCase) UpdateUserByID(id uint, ctx context.Context, domain Domain) (Domain, error){
	domain.Id = (id)
	user, err := usecase.repo.UpdateUserByID(id, ctx , domain)
	if err != nil {
		return Domain{}, errors.New("tidak ada user dengan ID tersebut")
	}
	return user, nil
}


func (usecase *UserUseCase) DeleteUserByID(id uint, ctx context.Context) error{
	err := usecase.repo.DeleteUserByID(id, ctx)
	if err != nil {
		return errors.New("tidak ada user dengan ID tersebut")
	}
	return nil
}