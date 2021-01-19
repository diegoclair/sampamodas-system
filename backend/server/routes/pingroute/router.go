package pingroute

import (
	"github.com/labstack/echo/v4"
)

// PingRouter holds the ping handlers
type PingRouter struct {
	ctrl   *Controller
	router *echo.Echo
}

// NewRouter returns a new PingRouter instance
func NewRouter(ctrl *Controller, router *echo.Echo) *PingRouter {
	return &PingRouter{
		ctrl:   ctrl,
		router: router,
	}
}

//RegisterRoutes is a routers map of ping requests
func (r *PingRouter) RegisterRoutes() {
	r.router.GET("/ping/", r.ctrl.handlePing)
}
