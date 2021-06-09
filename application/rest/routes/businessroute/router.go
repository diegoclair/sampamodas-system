package businessroute

import (
	"github.com/labstack/echo/v4"
)

// BusinessRouter holds the business handlers
type BusinessRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new BusinessRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *BusinessRouter {
	return &BusinessRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of business requests
func (r *BusinessRouter) RegisterRoutes() {
	r.router.POST("/business/", r.ctrl.handleCreateBusiness)
	r.router.GET("/businesses/", r.ctrl.handleGetBusinesses)
	r.router.GET("/business/:uuid/", r.ctrl.handleGetBusinessByID)
	r.router.GET("/businesses/:company_uuid/", r.ctrl.handleGetBusinessByCompanyID)
}
