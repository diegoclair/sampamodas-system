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

	lead.LeadAddress, restErr = s.svc.db.Lead().GetLeadAddressByLeadID(lead.ID)
	if restErr != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadAddressByLeadID: ", restErr)
		return lead, restErr
	}

	return lead, nil
}

func (s *leadService) CreateLead(lead entity.Lead) (leadID int64, restErr resterrors.RestErr) {

	leadID, restErr = s.svc.db.Lead().CreateLead(lead)
	if restErr != nil {
		logger.Error("leadService.CreateLead.CreateLead: ", restErr)
		return leadID, restErr
	}

	return leadID, nil
}

func (s *leadService) CreateLeadAddress(leadAddress entity.LeadAddress) resterrors.RestErr {

	restErr := s.svc.db.Lead().CreateLeadAddress(leadAddress)
	if restErr != nil {
		logger.Error("leadService.CreateLeadAddress.CreateLeadAddress: ", restErr)
		return restErr
	}

	return nil
}
