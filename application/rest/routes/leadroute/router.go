package leadroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute         = "/"
	createAddress     = "/address/"
	leadByPhoneNumber = "/:phone_number/"
)

type LeadRouter struct {
	ctrl      *Controller
	routeName string
}

func NewRouter(ctrl *Controller, routeName string) *LeadRouter {
	return &LeadRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

func (r *LeadRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := privateGroup.Group(r.routeName)
	router.POST(rootRoute, r.ctrl.handleCreateLead)
	router.POST(createAddress, r.ctrl.handleCreateLeadAddress)
	router.GET(leadByPhoneNumber, r.ctrl.handleGetLeadByPhoneNumber)
}
