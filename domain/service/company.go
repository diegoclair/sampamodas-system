package service

import (
	"github.com/diegoclair/go_utils-lib/v2/logger"
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

	logger.Info("GetCompanies: Process Started")
	defer logger.Info("GetCompanies: Process Finished")

	companies, err = s.svc.dm.MySQL().Company().GetCompanies()
	if err != nil {
		logger.Error("GetCompanies.GetCompanies: ", err)
		return companies, err
	}

	return companies, nil
}

func (s *companyService) GetCompanyByUUID(companyUUID string) (company entity.Company, err error) {

	logger.Info("GetCompanyByUUID: Process Started")
	defer logger.Info("GetCompanyByUUID: Process Finished")

	company, err = s.svc.dm.MySQL().Company().GetCompanyByUUID(companyUUID)
	if err != nil {
		logger.Error("GetCompanyByUUID.GetCompanyByUUID: ", err)
		return company, err
	}

	return company, nil
}

func (s *companyService) CreateCompany(company entity.Company) error {

	logger.Info("CreateCompany: Process Started")
	defer logger.Info("CreateCompany: Process Finished")

	company.UUID = uuid.NewV4().String()
	err := s.svc.dm.MySQL().Company().CreateCompany(company)
	if err != nil {
		logger.Error("CreateCompany.CreateCompany: ", err)
		return err
	}

	return nil
}
