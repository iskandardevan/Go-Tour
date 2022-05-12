package savedplace

import (
	"go-tour/businesses/savedplace"
	"go-tour/controllers"
	"go-tour/controllers/savedplace/request"
	"go-tour/controllers/savedplace/response"
	"go-tour/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SavedPlaceController struct {
	savedplaceUseCase savedplace.SavedPlaceUsecaseInterface
} 

func NewSavedPlaceController(savedplaceUseCase savedplace.SavedPlaceUsecaseInterface) *SavedPlaceController {
	return &SavedPlaceController{
		savedplaceUseCase: savedplaceUseCase,
	}
}

func (controller *SavedPlaceController) Add(c echo.Context) (error) {
	req := request.SavedPlaceRequest{}
	var err error
	err = c.Bind(&req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()
	var data savedplace.Domain
	data, err = controller.savedplaceUseCase.Add(ctx, *req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data), "Place added to wishlist")
}

func (controller *SavedPlaceController) GetAll(c echo.Context) (error) {
	req := c.Request().Context()
	savedplace, err := controller.savedplaceUseCase.GetAll(req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.GetAllSavedPlaces(savedplace))
}

func (controller *SavedPlaceController) GetByID(c echo.Context) (error) {
	req := c.Request().Context()
	id := c.Param("id")
	Convint, errConvint := strconv.Atoi(id)
	if errConvint != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errConvint)
	}
	data, err := controller.savedplaceUseCase.GetByID(uint(Convint), req)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *SavedPlaceController) Update(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req := request.SavedPlaceRequest{}
	err = c.Bind(&req)
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	data, err := controller.savedplaceUseCase.Update(convID, ctx, *req.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (controller *SavedPlaceController) Delete(c echo.Context) (error) {
	id := c.Param("id")
	convID, err := helpers.StringToUint(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	ctx := c.Request().Context()
	err = controller.savedplaceUseCase.Delete(convID, ctx)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccesResponse(c, nil, "Place removed from wishlist")
}
