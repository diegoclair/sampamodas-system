package leadroute

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/server/viewmodel"

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds lead handler functions
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

func (s *Controller) handleCreateLead(c echo.Context) error {

	input := viewmodel.Lead{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateLead.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	lead := entity.Lead{}

	err = s.mapper.From(input).To(&lead)
	if err != nil {
		restErr := resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	leadID, restErr := s.leadService.CreateLead(lead)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := viewmodel.CreateLeadResponse{}
	response.LeadID = leadID

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleCreateLeadAddress(c echo.Context) error {

	input := viewmodel.LeadAddress{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateLead.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	leadAddress := entity.LeadAddress{}

	err = s.mapper.From(input).To(&leadAddress)
	if err != nil {
		restErr := resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	restErr := s.leadService.CreateLeadAddress(leadAddress)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
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
