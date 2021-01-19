package server

import (
	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/leadroute"
	"github.com/diegoclair/sampamodas-system/backend/server/routes/pingroute"
	"github.com/diegoclair/sampamodas-system/backend/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type controller struct {
	pingController *pingroute.Controller
	leadController *leadroute.Controller
}

//InitServer to initialize the server
func InitServer(svc *service.Service) *gin.Engine {
	mapper := mapper.New()
	svm := service.NewServiceManager()
	srv := gin.Default()

	leadService := svm.LeadService(svc)

	srv.Use(cors.Default())

	return setupRoutes(srv, &controller{
		pingController: pingroute.NewController(),
		leadController: leadroute.NewController(leadService, mapper),
	})
}

//setupRoutes - Register and instantiate the routes
func setupRoutes(srv *gin.Engine, s *controller) *gin.Engine {

	pingroute.NewRouter(s.pingController, srv).RegisterRoutes()
	leadroute.NewRouter(s.leadController, srv).RegisterRoutes()

	return srv
}
