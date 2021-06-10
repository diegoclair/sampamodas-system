package companyroute

import (
	"github.com/labstack/echo/v4"
)

const (
	rootRoute     = "/"
	companyByUUID = "/:uuid/"
)

type CompanyRouter struct {
	ctrl      *Controller
	routeName string
}

func NewRouter(ctrl *Controller, routeName string) *CompanyRouter {
	return &CompanyRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

func (r *CompanyRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := privateGroup.Group(r.routeName)
	router.POST(rootRoute, r.ctrl.handleCreateCompany)
	router.GET(rootRoute, r.ctrl.handleGetCompanies)
	router.GET(companyByUUID, r.ctrl.handleGetCompanyByUUID)
}
