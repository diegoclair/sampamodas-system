package businessroute

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

	input := viewmodel.CreateBusiness{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateBusiness.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	business := entity.Business{}

	mapErr := s.mapper.From(input).To(&business)
	if mapErr != nil {
		errorMessage := "Error to do the mapper: "
		logger.Error("handleCreateBusiness - "+errorMessage, mapErr)
		restErr := resterrors.NewInternalServerError(errorMessage + fmt.Sprint(mapErr))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	restErr := s.businessService.CreateBusiness(business)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
}
