package companyroute

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/viewmodel"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds company handler functions
type Controller struct {
	companyService contract.CompanyService
	mapper         mapper.Mapper
}

//NewController to handle requests
func NewController(companyService contract.CompanyService, mapper mapper.Mapper) *Controller {
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
		logger.Error("handleCreateCompany.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	company := entity.Company{}

	err = s.mapper.From(input).To(&company)
	if err != nil {
		restErr := resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	restErr := s.companyService.CreateCompany(company)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Controller) handleGetCompanies(c echo.Context) error {

	companies, restErr := s.companyService.GetCompanies()
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := []viewmodel.Company{}
	err := s.mapper.From(companies).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleGetCompanyByID(c echo.Context) error {

	companyID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := resterrors.NewBadRequestError("id parameter is invalid")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	company, restErr := s.companyService.GetCompanyByID(int64(companyID))
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := viewmodel.Company{}
	err = s.mapper.From(company).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}
