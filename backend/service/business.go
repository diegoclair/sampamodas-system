package service

import (
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
	return s.svc.db.Business().GetBusinesses()
}

func (s *businessService) GetBusinessByID(businessID int64) (business entity.Business, restErr resterrors.RestErr) {
	return s.svc.db.Business().GetBusinessByID(businessID)
}

func (s *businessService) GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr) {
	return s.svc.db.Business().GetBusinessesByCompanyID(companyID)
}

func (s *businessService) CreateBusiness(business entity.Business) resterrors.RestErr {
	return s.svc.db.Business().CreateBusiness(business)
}
