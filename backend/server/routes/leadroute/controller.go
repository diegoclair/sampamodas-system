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

func (s *Controller) handleGetLeadAddress(c *gin.Context) {

	leadID, parseErr := strconv.Atoi(c.Param("lead_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("lead_id parameter is invalid")
		c.JSON(err.StatusCode(), err)
		return
	}

	leadAddresses, err := s.leadService.GetLeadAddress(int64(leadID))
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	response := []viewmodel.Address{}
	mapErr := s.mapper.From(leadAddresses).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s *Controller) handleCreateNewSale(c *gin.Context) {

	leadID, parseErr := strconv.Atoi(c.Param("lead_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("lead_id parameter is invalid")
		c.JSON(err.StatusCode(), err)
		return
	}

	input := viewmodel.Sale{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		logger.Error("handleCreateNewSale", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.StatusCode(), restErr)
		return
	}

	sale := entity.Sale{}

	mapErr := s.mapper.From(input).To(&sale)
	if mapErr != nil {
		restErr := resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		c.JSON(restErr.StatusCode(), restErr)
		return
	}

	sale.LeadID = int64(leadID)

	saleNumber, createErr := s.leadService.CreateSale(sale)
	if createErr != nil {
		c.JSON(createErr.StatusCode(), createErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"saleNumber": saleNumber})
}

func (s *Controller) handleGetSaleSummary(c *gin.Context) {

	leadID, parseErr := strconv.Atoi(c.Param("lead_id"))
	if parseErr != nil {
		err := resterrors.NewBadRequestError("lead_id parameter is invalid")
		c.JSON(err.StatusCode(), err)
		return
	}

	saleSummary, err := s.leadService.GetLeadSalesSummary(int64(leadID))
	if err != nil {
		c.JSON(err.StatusCode(), err)
		return
	}

	response := []viewmodel.SaleSummary{}
	mapErr := s.mapper.From(saleSummary).To(&response)
	if mapErr != nil {
		err = resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(mapErr))
		c.JSON(err.StatusCode(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}
