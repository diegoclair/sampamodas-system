package companyroute

import (
	"fmt"
	"net/http"
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

//Controller is a interface to interact with services
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

	input := viewmodel.CreateCompany{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateCompany.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	company := entity.Company{}

	mapErr := s.mapper.From(input).To(&company)
	if mapErr != nil {
		errorMessage := "Error to do the mapper: "
		logger.Error("handleCreateCompany - "+errorMessage, mapErr)
		restErr := resterrors.NewInternalServerError(errorMessage + fmt.Sprint(mapErr))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	restErr := s.companyService.CreateCompany(company)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
}
