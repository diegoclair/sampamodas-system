package businessroute

import (
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routeutils"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/viewmodel"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/domain/service"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds business handler functions
type Controller struct {
	businessService service.BusinessService
	mapper          mapper.Mapper
}

//NewController to handle requests
func NewController(businessService service.BusinessService, mapper mapper.Mapper) *Controller {
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
		return routeutils.HandleAPIError(c, err)
	}

	business := entity.Business{}

	err = s.mapper.From(input).To(&business)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = s.businessService.CreateBusiness(business)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}

func (s *Controller) handleGetBusinesses(c echo.Context) error {

	businesses, err := s.businessService.GetBusinesses()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := []viewmodel.Business{}
	err = s.mapper.From(businesses).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleGetBusinessByID(c echo.Context) error {

	businessUUID, err := routeutils.GetAndValidateStringParam(c, "uuid", true)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	business, err := s.businessService.GetBusinessByUUID(businessUUID)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.Business{}
	err = s.mapper.From(business).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleGetBusinessByCompanyID(c echo.Context) error {

	companyUUID, err := routeutils.GetAndValidateStringParam(c, "company_uuid", true)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	business, err := s.businessService.GetBusinessesByCompanyUUID(companyUUID)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := []viewmodel.Business{}
	err = s.mapper.From(business).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}
