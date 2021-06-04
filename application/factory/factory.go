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

		cfg := config.GetConfigEnvironment()
		svc := service.New(data, cfg)
		svm := service.NewServiceManager()
		mapper := mapper.New()

		instance = &Services{
			Cfg:             cfg,
			Mapper:          mapper,
			BusinessService: svm.BusinessService(svc),
			CompanyService:  svm.CompanyService(svc),
			LeadService:     svm.LeadService(svc),
			ProductService:  svm.ProductService(svc),
			SaleService:     svm.SaleService(svc, instance.ProductService),
		}

	})

	return instance
}
