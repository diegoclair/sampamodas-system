package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

// PingService holds a ping service operations
type PingService interface {
}

// BusinessService holds a business service operations
type BusinessService interface {
	CreateBusiness(company entity.Business) resterrors.RestErr
	GetBusinesses() (businesses []entity.Business, restErr resterrors.RestErr)
	GetBusinessByID(businessID int64) (businesses entity.Business, restErr resterrors.RestErr)
	GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr)
}

// CompanyService holds a company service operations
type CompanyService interface {
	CreateCompany(company entity.Company) resterrors.RestErr
	GetCompanies() (companies []entity.Company, restErr resterrors.RestErr)
	GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr)
}

// LeadService holds a lead service operations
type LeadService interface {
	GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, err resterrors.RestErr)
}

// ProductService holds a product service operations
type ProductService interface {
	CreateProduct(product entity.Product) resterrors.RestErr
	GetProducts() (products []entity.Product, restErr resterrors.RestErr)
	GetProductByID(productID int64) (product entity.Product, restErr resterrors.RestErr)
}
