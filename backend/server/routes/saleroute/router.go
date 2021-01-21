package saleroute

import (
	"github.com/labstack/echo/v4"
)

// SaleRouter holds the sale handlers
type SaleRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new SaleRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *SaleRouter {
	return &SaleRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of sale requests
func (r *SaleRouter) RegisterRoutes() {
	r.router.POST("/sale/", r.ctrl.handleCreateSale)
}
