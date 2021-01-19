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

func (s *companyService) GetLeadAddress(leadID int64) (address []entity.Address, err resterrors.RestErr) {
	return s.svc.db.Lead().GetLeadAddress(leadID)
}

func (s *companyService) GetLeadSalesSummary(leadID int64) (summary []entity.SaleSummary, err resterrors.RestErr) {

	return s.svc.db.Lead().GetSaleSummary(leadID)
}

func (s *companyService) CreateCompany(company entity.Company) resterrors.RestErr {

	err := s.svc.db.Company().CreateCompany(company)
	if err != nil {
		return err
	}

	return nil
}
