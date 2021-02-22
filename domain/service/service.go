package service

import (
	"github.com/diegoclair/sampamodas-system/backend/contract"
)

// Service holds the domain service repositories
type Service struct {
	dm contract.DataManager
}

// New returns a new domain Service instance
func New(dm contract.DataManager) *Service {
	svc := new(Service)
	svc.dm = dm

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	LeadService(svc *Service) contract.LeadService
	CompanyService(svc *Service) contract.CompanyService
	BusinessService(svc *Service) contract.BusinessService
	ProductService(svc *Service) contract.ProductService
	SaleService(svc *Service, productService contract.ProductService) contract.SaleService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) BusinessService(svc *Service) contract.BusinessService {
	return newBusinessService(svc)
}

func (s *serviceManager) CompanyService(svc *Service) contract.CompanyService {
	return newCompanyService(svc)
}

func (s *serviceManager) LeadService(svc *Service) contract.LeadService {
	return newLeadService(svc)
}

func (s *serviceManager) ProductService(svc *Service) contract.ProductService {
	return newProductService(svc)
}

func (s *serviceManager) SaleService(svc *Service, productService contract.ProductService) contract.SaleService {
	return newSaleService(svc, productService)
}
