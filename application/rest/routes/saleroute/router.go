package saleroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute = "/"
)

type SaleRouter struct {
	ctrl      *Controller
	routeName string
}

// NewRouter returns a new SaleRouter instance
func NewRouter(ctrl *Controller, routeName string) *SaleRouter {
	return &SaleRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

//RegisterRoutes is a routers map of sale requests
func (r *SaleRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := privateGroup.Group(r.routeName)
	router.POST(rootRoute, r.ctrl.handleCreateSale)
}
