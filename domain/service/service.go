package service

import (
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/util/config"
)

// Service holds the domain service repositories
type Service struct {
	dm  contract.DataManager
	cfg *config.EnvironmentVariables
}

// New returns a new domain Service instance
func New(dm contract.DataManager, cfg *config.EnvironmentVariables) *Service {
	svc := new(Service)
	svc.dm = dm
	svc.cfg = cfg

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	LeadService(svc *Service) LeadService
	CompanyService(svc *Service) CompanyService
	BusinessService(svc *Service) BusinessService
	ProductService(svc *Service) ProductService
	SaleService(svc *Service, productService ProductService) SaleService
}

// PingService holds a ping service operations
type PingService interface {
}

// BusinessService holds a business service operations
type BusinessService interface {
	CreateBusiness(company entity.Business) error
	GetBusinesses() (businesses []entity.Business, err error)
	GetBusinessByUUID(businessUUID string) (businesses entity.Business, err error)
	GetBusinessesByCompanyUUID(companyUUID string) (businesses []entity.Business, err error)
}

// CompanyService holds a company service operations
type CompanyService interface {
	CreateCompany(company entity.Company) error
	GetCompanies() (companies []entity.Company, err error)
	GetCompanyByUUID(companyUUID string) (company entity.Company, err error)
}

// LeadService holds a lead service operations
type LeadService interface {
	CreateLead(lead entity.Lead) (leadID int64, err error)
	CreateLeadAddress(leadAddress entity.LeadAddress) error
	GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, err error)
}

// ProductService holds a product service operations
type ProductService interface {
	CreateProduct(product entity.Product) error
	GetProducts() (products []entity.Product, err error)
	GetProductByID(productID int64) (product entity.Product, err error)
	GetProductByProductStockID(productStockID int64) (product entity.Product, err error)
}

// SaleService holds a sale service operations
type SaleService interface {
	CreateSale(sale entity.Sale) (saleID int64, err error)
	CreateSaleProduct(saleProduct entity.SaleProduct) error
	GetSales() (sales []entity.Sale, err error)
	GetSaleByID(saleID int64) (sale entity.Sale, err error)
}

type serviceManager struct {
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) BusinessService(svc *Service) BusinessService {
	return newBusinessService(svc)
}

func (s *serviceManager) CompanyService(svc *Service) CompanyService {
	return newCompanyService(svc)
}

func (s *serviceManager) LeadService(svc *Service) LeadService {
	return newLeadService(svc)
}

func (s *serviceManager) ProductService(svc *Service) ProductService {
	return newProductService(svc)
}

func (s *serviceManager) SaleService(svc *Service, productService ProductService) SaleService {
	return newSaleService(svc, productService)
}
