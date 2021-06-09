package productroute

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

//Controller holds product handler functions
type Controller struct {
	productService service.ProductService
	mapper         mapper.Mapper
}

//NewController to handle requests
func NewController(productService service.ProductService, mapper mapper.Mapper) *Controller {
	once.Do(func() {
		instance = &Controller{
			productService: productService,
			mapper:         mapper,
		}
	})
	return instance
}

func (s *Controller) handleCreateProduct(c echo.Context) error {

	input := viewmodel.Product{}

	err := c.Bind(&input)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	product := entity.Product{}

	err = s.mapper.From(input).To(&product)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	product.Brand.Name = input.BrandName
	product.Gender.Name = input.GenderName
	for i := range input.ProductStock {
		product.ProductStock[i].Color.Name = input.ProductStock[i].Color
	}

	err = s.productService.CreateProduct(product)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseCreated(c)
}

func (s *Controller) handleGetProducts(c echo.Context) error {

	products, err := s.productService.GetProducts()
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := []viewmodel.Product{}
	err = s.mapper.From(products).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}

func (s *Controller) handleGetProductByID(c echo.Context) error {

	productID, err := routeutils.GetAndValidateInt64Param(c, "id", true)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	product, err := s.productService.GetProductByID(productID)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	response := viewmodel.Product{}
	err = s.mapper.From(product).To(&response)
	if err != nil {
		return routeutils.HandleAPIError(c, err)
	}

	return routeutils.ResponseAPIOK(c, response)
}
