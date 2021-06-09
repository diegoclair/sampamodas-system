package companyroute

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

//Controller holds company handler functions
type Controller struct {
	companyService service.CompanyService
	mapper         mapper.Mapper
}

//NewController to handle requests
func NewController(companyService service.CompanyService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			companyService: companyService,
			mapper:         mapper,
		}
	})
	return instance
}

func (s *Controller) handleCreateCompany(c echo.Context) error {

	input := viewmodel.Company{}

	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	company := entity.Company{}

	err = s.mapper.From(input).To(&company)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = s.companyService.CreateCompany(company)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}

func (s *Controller) handleGetCompanies(c echo.Context) error {

	companies, err := s.companyService.GetCompanies()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := []viewmodel.Company{}
	err = s.mapper.From(companies).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleGetCompanyByID(c echo.Context) error {

	companyID, err := routeutils.GetAndValidateInt64Param(c, "id", true)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	company, err := s.companyService.GetCompanyByID(int64(companyID))
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.Company{}
	err = s.mapper.From(company).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}
