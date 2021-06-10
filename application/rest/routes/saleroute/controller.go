package saleroute

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

type Controller struct {
	saleService service.SaleService
	mapper      mapper.Mapper
}

func NewController(saleService service.SaleService, mapper mapper.Mapper) *Controller {
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
		return routeutils.HandleAPIError(c, err)
	}

	sale := entity.Sale{}

	err = s.mapper.From(input).To(&sale)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	saleID, err := s.saleService.CreateSale(sale)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.CreateSaleResponse{}
	response.SaleID = saleID

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleSaleProduct(c echo.Context) error {

	input := viewmodel.SaleProduct{}

	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	saleProduct := entity.SaleProduct{}

	err = s.mapper.From(input).To(&saleProduct)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	err = s.saleService.CreateSaleProduct(saleProduct)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}
