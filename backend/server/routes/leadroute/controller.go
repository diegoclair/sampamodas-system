package leadroute

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/server/viewmodel"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller is a interface to interact with services
type Controller struct {
	leadService contract.LeadService
	mapper      mapper.Mapper
}

//NewController to handle requests
func NewController(leadService contract.LeadService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			leadService: leadService,
			mapper:      mapper,
		}
	})
	return instance
}

func (s *Controller) handleGetLeadByPhoneNumber(c echo.Context) error {

	phoneNumber := c.Param("phone_number")

	lead, err := s.leadService.GetLeadByPhoneNumber(phoneNumber)
	if err != nil {
		return c.JSON(err.StatusCode(), err)
	}

	response := viewmodel.Lead{}
	mapErr := s.mapper.From(lead).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(mapErr))
		return c.JSON(err.StatusCode(), err)
	}

	return c.JSON(http.StatusOK, response)
}
