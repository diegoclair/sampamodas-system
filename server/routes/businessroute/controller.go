package businessroute

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/server/viewmodel"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds business handler functions
type Controller struct {
	businessService contract.BusinessService
	mapper          mapper.Mapper
}

//NewController to handle requests
func NewController(businessService contract.BusinessService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			businessService: businessService,
			mapper:          mapper,
		}
	})
	return instance
}

func (s *Controller) handleCreateBusiness(c echo.Context) error {

	input := viewmodel.Business{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateBusiness.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	business := entity.Business{}

	mapErr := s.mapper.From(input).To(&business)
	if mapErr != nil {
		restErr := resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(mapErr))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	restErr := s.businessService.CreateBusiness(business)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Controller) handleGetBusinesses(c echo.Context) error {

	businesses, restErr := s.businessService.GetBusinesses()
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := []viewmodel.Business{}
	err := s.mapper.From(businesses).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleGetBusinessByID(c echo.Context) error {

	businessID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := resterrors.NewBadRequestError("id parameter is invalid")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	business, restErr := s.businessService.GetBusinessByID(int64(businessID))
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := viewmodel.Business{}
	err = s.mapper.From(business).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleGetBusinessByCompanyID(c echo.Context) error {

	companyID, err := strconv.Atoi(c.Param("company_id"))
	if err != nil {
		restErr := resterrors.NewBadRequestError("company_id parameter is invalid")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	business, restErr := s.businessService.GetBusinessesByCompanyID(int64(companyID))
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := []viewmodel.Business{}
	err = s.mapper.From(business).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}
