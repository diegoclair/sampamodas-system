package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type leadService struct {
	svc *Service
}

//newLeadService return a new instance of the service
func newLeadService(svc *Service) contract.LeadService {
	return &leadService{
		svc: svc,
	}
}

func (s *leadService) GetLeadAddress(leadID int64) (address []entity.Address, err resterrors.RestErr) {
	return s.svc.db.Lead().GetLeadAddress(leadID)
}

func (s *leadService) GetLeadSalesSummary(leadID int64) (summary []entity.SaleSummary, err resterrors.RestErr) {

	return s.svc.db.Lead().GetSaleSummary(leadID)
}

func (s *leadService) CreateSale(sale entity.Sale) (saleNumber string, err resterrors.RestErr) {

	err = s.svc.db.Lead().CreateSale(sale)
	if err != nil {
		return saleNumber, err
	}

	return saleNumber, nil
}
