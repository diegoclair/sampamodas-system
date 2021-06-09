package leadroute

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

//Controller holds lead handler functions
type Controller struct {
	leadService service.LeadService
	mapper      mapper.Mapper
}

//NewController to handle requests
func NewController(leadService service.LeadService, mapper mapper.Mapper) *Controller {
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
		return routeutils.HandleAPIError(c, err)
	}

	err = input.Validate()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	lead := entity.Lead{}

	err = s.mapper.From(input).To(&lead)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	leadID, err := s.leadService.CreateLead(lead)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.CreateLeadResponse{
		LeadID: leadID,
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleCreateLeadAddress(c echo.Context) error {

	input := viewmodel.LeadAddress{}

	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = input.Validate()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	leadAddress := entity.LeadAddress{}

	err = s.mapper.From(input).To(&leadAddress)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = s.leadService.CreateLeadAddress(leadAddress)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}

func (s *Controller) handleGetLeadByPhoneNumber(c echo.Context) error {

	phoneNumber, err := routeutils.GetAndValidateStringParam(c, "phone_number", true)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	lead, err := s.leadService.GetLeadByPhoneNumber(phoneNumber)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.Lead{}
	err = s.mapper.From(lead).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}
