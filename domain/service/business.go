package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/twinj/uuid"
)

type businessService struct {
	svc *Service
}

//newBusinessService return a new instance of the service
func newBusinessService(svc *Service) BusinessService {
	return &businessService{
		svc: svc,
	}
}

func (s *businessService) GetBusinesses() (businesses []entity.Business, err error) {

	businesses, err = s.svc.dm.MySQL().Business().GetBusinesses()
	if err != nil {
		logger.Error("businessService.GetBusinesses.GetBusinesses: ", err)
		return businesses, err
	}

	return businesses, nil
}

func (s *businessService) GetBusinessByUUID(businessUUID string) (business entity.Business, err error) {

	business, err = s.svc.dm.MySQL().Business().GetBusinessByUUID(businessUUID)
	if err != nil {
		logger.Error("businessService.GetBusinessByUUID.GetBusinessByUUID: ", err)
		return business, err
	}

	return business, nil
}

func (s *businessService) GetBusinessesByCompanyUUID(companyUUID string) (businesses []entity.Business, err error) {
	businesses, err = s.svc.dm.MySQL().Business().GetBusinessesByCompanyUUID(companyUUID)
	if err != nil {
		logger.Error("businessService.GetBusinessesByCompanyUUID.GetBusinessesByCompanyUUID: ", err)
		return businesses, err
	}

	return businesses, nil
}

func (s *businessService) CreateBusiness(business entity.Business) error {

	business.UUID = uuid.NewV4().String()
	err := s.svc.dm.MySQL().Business().CreateBusiness(business)
	if err != nil {
		logger.Error("businessService.CreateBusiness.CreateBusiness: ", err)
		return err
	}

	return nil
}
