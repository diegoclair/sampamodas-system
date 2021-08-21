package factory

import (
	"log"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/sampamodas-system/backend/domain/service"
	"github.com/diegoclair/sampamodas-system/backend/infra/data"
	"github.com/diegoclair/sampamodas-system/backend/util/config"
)

type Services struct {
	Cfg             *config.EnvironmentVariables
	Mapper          mapper.Mapper
	BusinessService service.BusinessService
	CompanyService  service.CompanyService
	LeadService     service.LeadService
	ProductService  service.ProductService
	SaleService     service.SaleService
	AuthService     service.AuthService
}

var (
	instance *Services
	once     sync.Once
)

//GetDomainServices to get instace of all services
func GetDomainServices() *Services {

	once.Do(func() {

		data, err := data.Connect()
		if err != nil {
			log.Fatalf("Error to connect data repositories: %v", err)
		}

		cfg := config.GetConfigEnvironment()
		svc := service.New(data, cfg)
		svm := service.NewServiceManager()
		mapper := mapper.New()
		instance = &Services{}
		instance.Cfg = cfg
		instance.Mapper = mapper
		instance.AuthService = svm.AuthService(svc)
		instance.BusinessService = svm.BusinessService(svc)
		instance.CompanyService = svm.CompanyService(svc)
		instance.LeadService = svm.LeadService(svc)
		instance.ProductService = svm.ProductService(svc)
		instance.SaleService = svm.SaleService(svc, instance.ProductService)

	})

	return instance
}
