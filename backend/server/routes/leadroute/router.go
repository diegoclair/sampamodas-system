package leadroute

import (
	"github.com/gin-gonic/gin"
)

// LeadRouter holds the lead handlers
type LeadRouter struct {
	ctrl   *Controller
	router *gin.Engine
}

// NewRouter returns a new LeadRouter instance
func NewRouter(ctrl *Controller, router *gin.Engine) *LeadRouter {
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
