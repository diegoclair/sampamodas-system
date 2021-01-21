package saleroute

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
	saleService contract.SaleService
	mapper      mapper.Mapper
}

//NewController to handle requests
func NewController(saleService contract.SaleService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			saleService: saleService,
			mapper:      mapper,
		}
	})
	return instance
}

func (s *Controller) handleCreateSale(c echo.Context) error {

	input := viewmodel.Sale{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateSale.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	sale := entity.Sale{}

	err = s.mapper.From(input).To(&sale)
	if err != nil {
		restErr := resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	saleID, restErr := s.saleService.CreateSale(sale)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := viewmodel.CreateSaleResponse{}
	response.SaleID = saleID

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleSaleProduct(c echo.Context) error {

	input := viewmodel.SaleProduct{}

	err := c.Bind(&input)
	if err != nil {
		logger.Error("handleCreateLead.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	saleProduct := entity.SaleProduct{}

	err = s.mapper.From(input).To(&saleProduct)
	if err != nil {
		restErr := resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	restErr := s.saleService.CreateSaleProduct(saleProduct)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
}
