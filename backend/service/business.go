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

func (s *businessService) GetLeadAddress(leadID int64) (address []entity.Address, err resterrors.RestErr) {
	return s.svc.db.Lead().GetLeadAddress(leadID)
}

func (s *businessService) GetLeadSalesSummary(leadID int64) (summary []entity.SaleSummary, err resterrors.RestErr) {

	return s.svc.db.Lead().GetSaleSummary(leadID)
}

func (s *businessService) CreateBusiness(business entity.Business) resterrors.RestErr {

	err := s.svc.db.Business().CreateBusiness(business)
	if err != nil {
		return err
	}

	return nil
}
