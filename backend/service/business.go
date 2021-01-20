package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type businessService struct {
	svc *Service
}

//newBusinessService return a new instance of the service
func newBusinessService(svc *Service) contract.BusinessService {
	return &businessService{
		svc: svc,
	}
}

func (s *businessService) GetBusinesses() (businesses []entity.Business, restErr resterrors.RestErr) {

	businesses, restErr = s.svc.db.Business().GetBusinesses()
	if restErr != nil {
		logger.Error("businessService.GetBusinesses.GetBusinesses: ", restErr)
		return businesses, restErr
	}

	return businesses, nil
}

func (s *businessService) GetBusinessByID(businessID int64) (business entity.Business, restErr resterrors.RestErr) {

	business, restErr = s.svc.db.Business().GetBusinessByID(businessID)
	if restErr != nil {
		logger.Error("businessService.GetBusinesses.GetBusinessByID: ", restErr)
		return business, restErr
	}

	return business, nil
}

func (s *businessService) GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr) {
	businesses, restErr = s.svc.db.Business().GetBusinessesByCompanyID(companyID)
	if restErr != nil {
		logger.Error("businessService.GetBusinesses.GetBusinessesByCompanyID: ", restErr)
		return businesses, restErr
	}

	return businesses, nil
}

func (s *businessService) CreateBusiness(business entity.Business) resterrors.RestErr {

	restErr := s.svc.db.Business().CreateBusiness(business)
	if restErr != nil {
		logger.Error("businessService.GetBusinesses.CreateBusiness: ", restErr)
		return restErr
	}

	return nil
}
