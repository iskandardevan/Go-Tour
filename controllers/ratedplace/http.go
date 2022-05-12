package ratedplace

import (
	"go-tour/businesses/ratedplace"
	"go-tour/controllers"
	"go-tour/controllers/ratedplace/request"
	"go-tour/controllers/ratedplace/response"
	"go-tour/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RatedPlaceController struct {
	ratedplaceUseCase ratedplace.RatedPlaceUsecaseInterface
} 

func NewRatedPlaceController(ratedplaceUseCase ratedplace.RatedPlaceUsecaseInterface) *RatedPlaceController {
	return &RatedPlaceController{
		ratedplaceUseCase: ratedplaceUseCase,
	}
}

func (controller *RatedPlaceController) Add(c echo.Context) (error) {
	req := request.RatedPlaceRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data ratedplace.Domain
	data, err = controller.ratedplaceUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *RatedPlaceController) GetAll(c echo.Context) (error) {
	req := c.Request().Context()
	ratedplace, err := controller.ratedplaceUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllRatedPlaces(ratedplace))
}

func (controller *RatedPlaceController) GetByID(c echo.Context) (error) {
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := controller.ratedplaceUseCase.GetByID(uint(Convint), req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *RatedPlaceController) Update(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.RatedPlaceRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := controller.ratedplaceUseCase.Update(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *RatedPlaceController) Delete(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = controller.ratedplaceUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}
