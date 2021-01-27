package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
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

	companies, restErr = s.svc.db.Company().GetCompanies()
	if restErr != nil {
		logger.Error("companyService.GetCompanies.GetCompanies: ", restErr)
		return companies, restErr
	}

	return companies, nil
}

func (s *companyService) GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr) {

	company, restErr = s.svc.db.Company().GetCompanyByID(companyID)
	if restErr != nil {
		logger.Error("companyService.GetCompanies.GetCompanyByID: ", restErr)
		return company, restErr
	}

	return company, nil
}

func (s *companyService) CreateCompany(company entity.Company) resterrors.RestErr {

	restErr := s.svc.db.Company().CreateCompany(company)
	if restErr != nil {
		logger.Error("companyService.GetCompanies.CreateCompany: ", restErr)
		return restErr
	}

	return nil
}
