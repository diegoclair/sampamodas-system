package companyroute

import (
	"github.com/labstack/echo/v4"
)

// ComapanyRouter holds the lead handlers
type ComapanyRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new ComapanyRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *ComapanyRouter {
	return &ComapanyRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of lead requests
func (r *ComapanyRouter) RegisterRoutes() {
	r.router.POST("/company/", r.ctrl.handleCreateCompany)
}
