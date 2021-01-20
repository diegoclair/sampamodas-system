package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

// PingService holds a ping service operations
type PingService interface {
}

// BusinessService holds a ping service operations
type BusinessService interface {
	CreateBusiness(company entity.Business) resterrors.RestErr
	GetBusinesses() (businesses []entity.Business, restErr resterrors.RestErr)
	GetBusinessByID(businessID int64) (businesses entity.Business, restErr resterrors.RestErr)
	GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr)
}

// CompanyService holds a ping service operations
type CompanyService interface {
	CreateCompany(company entity.Company) resterrors.RestErr
	GetCompanies() (companies []entity.Company, restErr resterrors.RestErr)
	GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr)
}

// LeadService holds a lead service operations
type LeadService interface {
	GetLeadAddress(leadID int64) (address []entity.Address, err resterrors.RestErr)
	CreateSale(sale entity.Sale) (saleNumber string, err resterrors.RestErr)
	GetLeadSalesSummary(leadID int64) (summary []entity.SaleSummary, err resterrors.RestErr)
}
