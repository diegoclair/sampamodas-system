package leadroute

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
	"github.com/gin-gonic/gin"

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

func (s *Controller) handleGetLeadAddress(c echo.Context) error {

	leadID, parseErr := strconv.Atoi(c.Param("lead_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("lead_id parameter is invalid")
		return c.JSON(err.StatusCode(), err)
	}

	leadAddresses, err := s.leadService.GetLeadAddress(int64(leadID))
	if err != nil {
		return c.JSON(err.StatusCode(), err)
	}

	response := []viewmodel.Address{}
	mapErr := s.mapper.From(leadAddresses).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		return c.JSON(err.StatusCode(), err)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleCreateNewSale(c echo.Context) error {

	leadID, parseErr := strconv.Atoi(c.Param("lead_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("lead_id parameter is invalid")
		return c.JSON(err.StatusCode(), err)
	}

	input := viewmodel.Sale{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateNewSale", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	sale := entity.Sale{}

	mapErr := s.mapper.From(input).To(&sale)
	if mapErr != nil {
		restErr := resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	sale.LeadID = int64(leadID)

	saleNumber, restErr := s.leadService.CreateSale(sale)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, gin.H{"saleNumber": saleNumber})
}

func (s *Controller) handleGetSaleSummary(c echo.Context) error {

	leadID, parseErr := strconv.Atoi(c.Param("lead_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("lead_id parameter is invalid")
		return c.JSON(err.StatusCode(), err)
	}

	saleSummary, err := s.leadService.GetLeadSalesSummary(int64(leadID))
	if err != nil {
		return c.JSON(err.StatusCode(), err)
	}

	response := []viewmodel.SaleSummary{}
	mapErr := s.mapper.From(saleSummary).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		return c.JSON(err.StatusCode(), err)
	}

	return c.JSON(http.StatusOK, response)
}
