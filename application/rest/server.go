package rest

import (
	"fmt"
	"os"

	"github.com/diegoclair/go_utils-lib/logger"
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

type controller struct {
	pingController     *pingroute.Controller
	businessController *businessroute.Controller
	companyController  *companyroute.Controller
	leadController     *leadroute.Controller
	productController  *productroute.Controller
	saleController     *saleroute.Controller
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

	return registerRoutes(srv, &controller{
		pingController:     pingroute.NewController(),
		businessController: businessroute.NewController(factory.BusinessService, factory.Mapper),
		companyController:  companyroute.NewController(factory.CompanyService, factory.Mapper),
		leadController:     leadroute.NewController(factory.LeadService, factory.Mapper),
		productController:  productroute.NewController(factory.ProductService, factory.Mapper),
		saleController:     saleroute.NewController(factory.SaleService, factory.Mapper),
	})
}

//registerRoutes - Register and instantiate routes
func registerRoutes(srv *echo.Echo, s *controller) *echo.Echo {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	businessroute.NewRouter(s.businessController, srv).RegisterRoutes()
	companyroute.NewRouter(s.companyController, srv).RegisterRoutes()
	leadroute.NewRouter(s.leadController, srv).RegisterRoutes()
	productroute.NewRouter(s.productController, srv).RegisterRoutes()
	saleroute.NewRouter(s.saleController, srv).RegisterRoutes()

	return srv
}
