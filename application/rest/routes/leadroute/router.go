package leadroute

import (
	"github.com/labstack/echo/v4"
)

// LeadRouter holds the lead handlers
type LeadRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new LeadRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *LeadRouter {
	return &LeadRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of lead requests
func (r *LeadRouter) RegisterRoutes() {
	r.router.POST("/lead/", r.ctrl.handleCreateLead)
	r.router.POST("/lead/address/", r.ctrl.handleCreateLeadAddress)
	r.router.GET("/lead/:phone_number/", r.ctrl.handleGetLeadByPhoneNumber)
}
