package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
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

func (s *leadService) GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, restErr resterrors.RestErr) {

	lead, restErr = s.svc.db.Lead().GetLeadByPhoneNumber(phoneNumber)
	if restErr != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadByPhoneNumber: ", restErr)
		return lead, restErr
	}

	lead.LeadAddress, restErr = s.svc.db.Lead().GetLeadAddressByID(lead.ID)
	if restErr != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadAddressByID: ", restErr)
		return lead, restErr
	}

	return lead, nil
}
