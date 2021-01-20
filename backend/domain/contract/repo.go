package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Company() CompanyRepo
	Lead() LeadRepo
	Business() BusinessRepo
}

// BusinessRepo defines the data set for qrcode
type BusinessRepo interface {
	CreateBusiness(company entity.Business) resterrors.RestErr
	GetBusinesses() (businesses []entity.Business, restErr resterrors.RestErr)
	GetBusinessByID(businessID int64) (businesses entity.Business, restErr resterrors.RestErr)
	GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr)
}

// CompanyRepo defines the data set for qrcode
type CompanyRepo interface {
	CreateCompany(company entity.Company) resterrors.RestErr
	GetCompanies() (companies []entity.Company, restErr resterrors.RestErr)
	GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr)
}

// LeadRepo defines the data set for lead
type LeadRepo interface {
	GetLeadAddress(leadID int64) (address []entity.Address, err resterrors.RestErr)
	CreateSale(sale entity.Sale) (err resterrors.RestErr)
	GetSaleSummary(leadID int64) (summary []entity.SaleSummary, restErr resterrors.RestErr)
}
