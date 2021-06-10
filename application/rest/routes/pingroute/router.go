package pingroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute = "/"
)

type PingRouter struct {
	ctrl      *Controller
	routeName string
}

// NewRouter returns a new PingRouter instance
func NewRouter(ctrl *Controller, routeName string) *PingRouter {
	return &PingRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *PingRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := appGroup.Group(r.routeName)
	router.GET(rootRoute, r.ctrl.handlePing)
}
