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
	r.router.GET("/lead/:lead_id/address", r.ctrl.handleGetLeadAddress)
	r.router.POST("/lead/:lead_id/sale", r.ctrl.handleCreateNewSale)
	r.router.GET("/lead/:lead_id/sale_summary", r.ctrl.handleGetSaleSummary)
}
