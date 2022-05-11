package placeimages

import (
	"go-tour/businesses/placeimages"
	"go-tour/controllers"
	"go-tour/controllers/placeimages/request"
	"go-tour/controllers/placeimages/response"
	"go-tour/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PlaceImageController struct {
	placeimageUseCase placeimages.PlaceImageUsecaseInterface
} 

func NewPlaceImageController(placeimageUseCase placeimages.PlaceImageUsecaseInterface) *PlaceImageController {
	return &PlaceImageController{
		placeimageUseCase: placeimageUseCase,
	}
}

func (controller *PlaceImageController) Add(c echo.Context) (error) {
	req := request.PlaceImageRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data placeimages.Domain
	data, err = controller.placeimageUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *PlaceImageController) GetAll(c echo.Context) (error) {
	req := c.Request().Context()
	placeimage, err := controller.placeimageUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllPlaceImages(placeimage))
}

func (controller *PlaceImageController) GetByID(c echo.Context) (error) {
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := controller.placeimageUseCase.GetByID(uint(Convint), req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *PlaceImageController) Update(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.PlaceImageRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := controller.placeimageUseCase.Update(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *PlaceImageController) Delete(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = controller.placeimageUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}
