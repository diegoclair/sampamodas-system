package rest

import (
	"fmt"
	"os"

	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/diegoclair/sampamodas-system/backend/application/factory"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routes/businessroute"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routes/companyroute"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routes/leadroute"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routes/pingroute"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routes/productroute"
	"github.com/diegoclair/sampamodas-system/backend/application/rest/routes/saleroute"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// IRouter interface for routers
type IRouter interface {
	RegisterRoutes(appGroup, privateGroup *echo.Group)
}

// Router holds application's routers
type Router struct {
	routers []IRouter
}

//StartRestServer starts the restServer
func StartRestServer() {
	server := initServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	logger.Info("About to start the application...")

	if err := server.Start(fmt.Sprintf(":%s", port)); err != nil {
		panic(err)
	}
}

//initServer to initialize the server
func initServer() *echo.Echo {

	factory := factory.GetDomainServices()

	srv := echo.New()
	srv.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	pingController := pingroute.NewController()
	businessController := businessroute.NewController(factory.BusinessService, factory.Mapper)
	companyController := companyroute.NewController(factory.CompanyService, factory.Mapper)
	leadController := leadroute.NewController(factory.LeadService, factory.Mapper)
	productController := productroute.NewController(factory.ProductService, factory.Mapper)
	saleController := saleroute.NewController(factory.SaleService, factory.Mapper)

	pingRoute := pingroute.NewRouter(pingController, "ping")
	businessRoute := businessroute.NewRouter(businessController, "businessess")
	companyRoute := companyroute.NewRouter(companyController, "companies")
	leadRoute := leadroute.NewRouter(leadController, "leads")
	productRoute := productroute.NewRouter(productController, "products")
	saleRoute := saleroute.NewRouter(saleController, "sales")

	appRouter := &Router{}
	appRouter.addRouters(pingRoute)
	appRouter.addRouters(businessRoute)
	appRouter.addRouters(companyRoute)
	appRouter.addRouters(leadRoute)
	appRouter.addRouters(productRoute)
	appRouter.addRouters(saleRoute)

	return appRouter.registerAppRouters(srv)
}

func (r *Router) addRouters(router IRouter) {
	r.routers = append(r.routers, router)
}

func (r *Router) registerAppRouters(srv *echo.Echo) *echo.Echo {

	appGroup := srv.Group("/")
	privateGroup := appGroup.Group("") //future create middleware to check logged user

	for _, appRouter := range r.routers {
		appRouter.RegisterRoutes(appGroup, privateGroup)
	}

	return srv
}
