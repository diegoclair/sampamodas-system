package productroute

import (
	"github.com/labstack/echo/v4"
)

// ProductRouter holds the product handlers
type ProductRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new ProductRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *ProductRouter {
	return &ProductRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of product requests
func (r *ProductRouter) RegisterRoutes() {
	r.router.POST("/product/", r.ctrl.handleCreateProduct)
	r.router.GET("/products/", r.ctrl.handleGetProducts)
	r.router.GET("/product/:id/", r.ctrl.handleGetProductByID)
}
