package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
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

func (s *businessService) GetBusinessByID(businessID int64) (business entity.Business, err error) {

	business, err = s.svc.dm.MySQL().Business().GetBusinessByID(businessID)
	if err != nil {
		logger.Error("businessService.GetBusinesses.GetBusinessByID: ", err)
		return business, err
	}

	return business, nil
}

func (s *businessService) GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, err error) {
	businesses, err = s.svc.dm.MySQL().Business().GetBusinessesByCompanyID(companyID)
	if err != nil {
		logger.Error("businessService.GetBusinesses.GetBusinessesByCompanyID: ", err)
		return businesses, err
	}

	return businesses, nil
}

func (s *businessService) CreateBusiness(business entity.Business) error {

	err := s.svc.dm.MySQL().Business().CreateBusiness(business)
	if err != nil {
		logger.Error("businessService.GetBusinesses.CreateBusiness: ", err)
		return err
	}

	return nil
}
