package factory

import (
	"log"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/service"
	"github.com/diegoclair/sampamodas-system/backend/infra/data"
	"github.com/diegoclair/sampamodas-system/backend/util/config"
)

//Services is the factory to all serrvices
type Services struct {
	Cfg             *config.EnvironmentVariables
	Mapper          mapper.Mapper
	BusinessService contract.BusinessService
	CompanyService  contract.CompanyService
	LeadService     contract.LeadService
	ProductService  contract.ProductService
	SaleService     contract.SaleService
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

		svc := service.New(data)

		mapper := mapper.New()
		svm := service.NewServiceManager()

		instance = &Services{}
		instance.Cfg = config.GetConfigEnvironment()
		instance.Mapper = mapper
		instance.BusinessService = svm.BusinessService(svc)
		instance.CompanyService = svm.CompanyService(svc)
		instance.LeadService = svm.LeadService(svc)
		instance.ProductService = svm.ProductService(svc)
		instance.SaleService = svm.SaleService(svc, instance.ProductService)
	})

	return instance
}
