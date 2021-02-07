package productroute

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

	"github.com/labstack/echo/v4"
)

var (
	instance *Controller
	once     sync.Once
)

//Controller holds product handler functions
type Controller struct {
	productService contract.ProductService
	mapper         mapper.Mapper
}

//NewController to handle requests
func NewController(productService contract.ProductService, mapper mapper.Mapper) *Controller {
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
		logger.Error("handleCreateProduct.c.Bind: ", err)
		restErr := resterrors.NewBadRequestError("Invalid json body")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	product := entity.Product{}

	err = s.mapper.From(input).To(&product)
	if err != nil {
		restErr := resterrors.NewInternalServerError("Error to do the mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	product.Brand.Name = input.BrandName
	product.Gender.Name = input.GenderName
	for i := range input.ProductStock {
		product.ProductStock[i].Color.Name = input.ProductStock[i].Color
	}

	restErr := s.productService.CreateProduct(product)
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Controller) handleGetProducts(c echo.Context) error {

	products, restErr := s.productService.GetProducts()
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := []viewmodel.Product{}
	err := s.mapper.From(products).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}

func (s *Controller) handleGetProductByID(c echo.Context) error {

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		restErr := resterrors.NewBadRequestError("id parameter is invalid")
		return c.JSON(restErr.StatusCode(), restErr)
	}

	product, restErr := s.productService.GetProductByID(int64(productID))
	if restErr != nil {
		return c.JSON(restErr.StatusCode(), restErr)
	}

	response := viewmodel.Product{}
	err = s.mapper.From(product).To(&response)
	if err != nil {
		restErr = resterrors.NewInternalServerError("Error to mapper: " + fmt.Sprint(err))
		return c.JSON(restErr.StatusCode(), restErr)
	}

	return c.JSON(http.StatusOK, response)
}
