package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/businessroute"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/companyroute"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/leadroute"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/pingroute"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/productroute"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/saleroute"
	"github.com/diegoclair/sampamodas-system/backend/service"
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

//InitServer to initialize the server
func InitServer(svc *service.Service) *echo.Echo {
	mapper := mapper.New()
	svm := service.NewServiceManager()
	srv := echo.New()

	businessService := svm.BusinessService(svc)
	companyService := svm.CompanyService(svc)
	leadService := svm.LeadService(svc)
	productService := svm.ProductService(svc)
	saleService := svm.SaleService(svc, productService)

	srv.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	return setupRoutes(srv, &controller{
		pingController:     pingroute.NewController(),
		businessController: businessroute.NewController(businessService, mapper),
		companyController:  companyroute.NewController(companyService, mapper),
		leadController:     leadroute.NewController(leadService, mapper),
		productController:  productroute.NewController(productService, mapper),
		saleController:     saleroute.NewController(saleService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *echo.Echo, s *controller) *echo.Echo {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	businessroute.NewRouter(s.businessController, srv).RegisterRoutes()
	companyroute.NewRouter(s.companyController, srv).RegisterRoutes()
	leadroute.NewRouter(s.leadController, srv).RegisterRoutes()
	productroute.NewRouter(s.productController, srv).RegisterRoutes()
	saleroute.NewRouter(s.saleController, srv).RegisterRoutes()

	return srv
}