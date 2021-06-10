package productroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute   = "/"
	productByID = "/:id/"
)

type ProductRouter struct {
	ctrl      *Controller
	routeName string
}

func NewRouter(ctrl *Controller, routeName string) *ProductRouter {
	return &ProductRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

func (r *ProductRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := privateGroup.Group(r.routeName)
	router.POST(rootRoute, r.ctrl.handleCreateProduct)
	router.GET(rootRoute, r.ctrl.handleGetProducts)
	router.GET(productByID, r.ctrl.handleGetProductByID)
}
