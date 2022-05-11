package main

import (
	_middleware "go-tour/app/middlewares"

	userUseCase "go-tour/businesses/users"
	userController "go-tour/controllers/users"
	userRepo "go-tour/drivers/repository/users"

	placeUseCase "go-tour/businesses/places"
	placeController "go-tour/controllers/places"
	placeRepo "go-tour/drivers/repository/places"

	placeimageUseCase "go-tour/businesses/placeimages"
	placeimageController "go-tour/controllers/placeimages"
	placeimageRepo "go-tour/drivers/repository/placeimages"

	"go-tour/app/routes"
	"go-tour/drivers/mysql"

	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&userRepo.User{})
	DB.AutoMigrate(&placeRepo.Place{})
	DB.AutoMigrate(&placeimageRepo.PlaceImage{})
}



func main() {
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString("JWT.secretKey"),
		ExpiresDuration: viper.GetInt("JWT.expired_time"),
	}

	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	_middleware.Log(e)
	e.Use(middleware.CORS()) 
	userRepoInterface := userRepo.NewUserRepo(DB)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &configJWT)
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)
	
	placeRepoInterface := placeRepo.NewPlaceRepo(DB)
	placeUseCaseInterface := placeUseCase.NewUseCase(placeRepoInterface, timeoutContext)
	placeUseControllerInterface := placeController.NewPlaceController(placeUseCaseInterface)
	
	placeimageRepoInterface := placeimageRepo.NewPlaceImageRepo(DB)
	placeimageUseCaseInterface := placeimageUseCase.NewUseCase(placeimageRepoInterface, timeoutContext)
	placeimageUseControllerInterface := placeimageController.NewPlaceImageController(placeimageUseCaseInterface)

	routesInit := routes.RouteControllerList{
		UserController: *userUseControllerInterface,
		PlaceController: *placeUseControllerInterface,
		PlaceImageController: *placeimageUseControllerInterface,
		JWTMiddleware: configJWT.Init(),
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}