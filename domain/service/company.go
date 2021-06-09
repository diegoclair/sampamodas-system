package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/twinj/uuid"
)

type companyService struct {
	svc *Service
}

//newCompanyService return a new instance of the service
func newCompanyService(svc *Service) CompanyService {
	return &companyService{
		svc: svc,
	}
}

func (s *companyService) GetCompanies() (companies []entity.Company, err error) {

	companies, err = s.svc.dm.MySQL().Company().GetCompanies()
	if err != nil {
		logger.Error("companyService.GetCompanies.GetCompanies: ", err)
		return companies, err
	}

	return companies, nil
}

func (s *companyService) GetCompanyByUUID(companyUUID string) (company entity.Company, err error) {

	company, err = s.svc.dm.MySQL().Company().GetCompanyByUUID(companyUUID)
	if err != nil {
		logger.Error("companyService.GetCompanies.GetCompanyByUUID: ", err)
		return company, err
	}

	return company, nil
}

func (s *companyService) CreateCompany(company entity.Company) error {

	company.UUID = uuid.NewV4().String()
	err := s.svc.dm.MySQL().Company().CreateCompany(company)
	if err != nil {
		logger.Error("companyService.GetCompanies.CreateCompany: ", err)
		return err
	}

	return nil
}
