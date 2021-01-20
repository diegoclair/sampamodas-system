package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type companyService struct {
	svc *Service
}

//newCompanyService return a new instance of the service
func newCompanyService(svc *Service) contract.CompanyService {
	return &companyService{
		svc: svc,
	}
}

func (s *companyService) GetCompanies() (companies []entity.Company, restErr resterrors.RestErr) {
	return s.svc.db.Company().GetCompanies()
}

func (s *companyService) GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr) {
	return s.svc.db.Company().GetCompanyByID(companyID)
}

func (s *companyService) CreateCompany(company entity.Company) resterrors.RestErr {
	return s.svc.db.Company().CreateCompany(company)
}
