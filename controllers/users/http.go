package users

import (
	"go-tour/businesses/users"
	"go-tour/controllers"
	"go-tour/controllers/users/request"
	"go-tour/controllers/users/response"
	"go-tour/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase users.UserUsecaseInterface
}

func NewUserController(UserUseCase users.UserUsecaseInterface) *UserController{
	return &UserController{
		userUseCase: UserUseCase,
	}
}

func (userController *UserController) RegisterUser (c echo.Context) error {
	req := request.RegisterUserRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data users.Domain
	data, err = userController.userUseCase.RegisterUser(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))

}



func (userController *UserController) LoginUser (c echo.Context) error {
	var login users.Domain
	var err error
	var token string
	ctx := c.Request().Context()

	req := request.UserLoginRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	login, token, err = userController.userUseCase.LoginUser(req.Email, req.Password, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.UserLogin(login, token))
}

func (userController *UserController) CheckingUser (c echo.Context) error {
	var data users.Domain
	var err error
	ctx := c.Request().Context()

	req := request.UserLoginRequest{}
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err = userController.userUseCase.CheckingUser(req.Email, req.Password, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))
}


func (userController *UserController) GetByEmail (c echo.Context) error {
	ctx := c.Request().Context()
	req := request.GetByEmailRequest{} 
	err := c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := userController.userUseCase.GetByEmail(req.Email, ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))
}

func (userController *UserController) GetByID (c echo.Context) error{
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := userController.userUseCase.GetByID(uint(Convint), req )
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))
}

func (userController *UserController) GetAllUsers (c echo.Context) error {
	req := c.Request().Context()
	user, err := userController.userUseCase.GetAllUsers(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllUsers(user))

}

func (userController *UserController) UpdateUserByID (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.RegisterUserRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := userController.userUseCase.UpdateUserByID(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromUserRegister(data))

}


func (userController *UserController) DeleteUserByID (c echo.Context) error{
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = userController.userUseCase.DeleteUserByID(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}