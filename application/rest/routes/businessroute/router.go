package businessroute

import "github.com/labstack/echo/v4"

const (
	rootRoute             = "/"
	businessByUUID        = "/:uuid/"
	businessByCompanyUUID = "/:company_uuid/"
)

type BusinessRouter struct {
	ctrl      *Controller
	routeName string
}

func NewRouter(ctrl *Controller, routeName string) *BusinessRouter {
	return &BusinessRouter{
		ctrl:      ctrl,
		routeName: routeName,
	}
}

func (r *BusinessRouter) RegisterRoutes(appGroup, privateGroup *echo.Group) {
	router := privateGroup.Group(r.routeName)
	router.POST(rootRoute, r.ctrl.handleCreateBusiness)
	router.GET(rootRoute, r.ctrl.handleGetBusinesses)
	router.GET(businessByUUID, r.ctrl.handleGetBusinessByID)
	router.GET(businessByCompanyUUID, r.ctrl.handleGetBusinessByCompanyID)
}
