package routes

import (
	"go-tour/controllers/placeimages"
	"go-tour/controllers/places"
	"go-tour/controllers/ratedplace"
	"go-tour/controllers/savedplace"
	"go-tour/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController 
	PlaceController places.PlaceController
	PlaceImageController placeimages.PlaceImageController
	SavedPlaceController savedplace.SavedPlaceController
	RatedPlaceController ratedplace.RatedPlaceController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	jwt := middleware.JWTWithConfig(ctrl.JWTMiddleware)
	
	e.DELETE("/user/:id", ctrl.UserController.DeleteUserByID, jwt)
	e.POST("register", ctrl.UserController.RegisterUser)
	e.POST("login", ctrl.UserController.LoginUser) 
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.POST("user", ctrl.UserController.GetByEmail)
	e.GET("users", ctrl.UserController.GetAllUsers)
	e.PUT("user/:id", ctrl.UserController.UpdateUserByID, jwt)

	e.POST("/place", ctrl.PlaceController.Add)
	e.GET("/places", ctrl.PlaceController.GetAll)
	e.GET("/place/:id", ctrl.PlaceController.GetByID)
	e.PUT("/place/:id", ctrl.PlaceController.Update)
	e.DELETE("/place/:id", ctrl.PlaceController.Delete)

	e.POST("/placeimage", ctrl.PlaceImageController.Add)
	e.GET("/placeimages", ctrl.PlaceImageController.GetAll)
	e.GET("/placeimage/:id", ctrl.PlaceImageController.GetByID)
	e.PUT("/placeimage/:id", ctrl.PlaceImageController.Update)
	e.DELETE("/placeimage/:id", ctrl.PlaceImageController.Delete)

	e.POST("/savedplace", ctrl.SavedPlaceController.Add)
	e.GET("/savedplaces", ctrl.SavedPlaceController.GetAll)
	e.GET("/savedplace/:id", ctrl.SavedPlaceController.GetByID)
	e.PUT("/savedplace/:id", ctrl.SavedPlaceController.Update)
	e.DELETE("/savedplace/:id", ctrl.SavedPlaceController.Delete)
	
	e.POST("/ratedplace", ctrl.RatedPlaceController.Add)
	e.GET("/ratedplaces", ctrl.RatedPlaceController.GetAll)
	e.GET("/ratedplace/:id", ctrl.RatedPlaceController.GetByID)
	e.PUT("/ratedplace/:id", ctrl.RatedPlaceController.Update)
	e.DELETE("/ratedplace/:id", ctrl.RatedPlaceController.Delete)

}