package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Company() CompanyRepo
	Lead() LeadRepo
}

// CompanyRepo defines the data set for qrcode
type CompanyRepo interface {
	CreateCompany(company entity.Company) resterrors.RestErr
}

// LeadRepo defines the data set for lead
type LeadRepo interface {
	GetLeadAddress(leadID int64) (address []entity.Address, err resterrors.RestErr)
	CreateSale(sale entity.Sale) (err resterrors.RestErr)
	GetSaleSummary(leadID int64) (summary []entity.SaleSummary, restErr resterrors.RestErr)
}
