package places

import (
	"go-tour/businesses/places"
	"go-tour/controllers"
	"go-tour/controllers/places/request"
	"go-tour/controllers/places/response"
	"go-tour/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PlaceController struct {
	placeUseCase places.PlaceUsecaseInterface
}

func NewPlaceController(placeUseCase places.PlaceUsecaseInterface) *PlaceController {
	return &PlaceController{
		placeUseCase: placeUseCase,
	}
}

func (controller *PlaceController) Add(c echo.Context) (error) {
	req := request.PlaceRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data places.Domain
	data, err = controller.placeUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *PlaceController) GetAll(c echo.Context) (error) {
	req := c.Request().Context()
	place, err := controller.placeUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllPlaces(place))
}

func (controller *PlaceController) GetByID(c echo.Context) (error) {
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := controller.placeUseCase.GetByID(uint(Convint), req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *PlaceController) Update(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.PlaceRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := controller.placeUseCase.Update(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *PlaceController) Delete(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = controller.placeUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil)
}
