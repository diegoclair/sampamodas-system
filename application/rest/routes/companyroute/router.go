package companyroute

import (
	"github.com/labstack/echo/v4"
)

// CompanyRouter holds the company handlers
type CompanyRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new CompanyRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *CompanyRouter {
	return &CompanyRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of company requests
func (r *CompanyRouter) RegisterRoutes() {
	r.router.POST("/company/", r.ctrl.handleCreateCompany)
	r.router.GET("/companies/", r.ctrl.handleGetCompanies)
	r.router.GET("/company/:uuid/", r.ctrl.handleGetCompanyByID)
}
