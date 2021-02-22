package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/util/format"
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

	lead, restErr = s.svc.dm.MySQL().Lead().GetLeadByPhoneNumber(phoneNumber)
	if restErr != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadByPhoneNumber: ", restErr)
		return lead, restErr
	}

	lead.LeadAddress, restErr = s.svc.dm.MySQL().Lead().GetLeadAddressByLeadID(lead.ID)
	if restErr != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadAddressByLeadID: ", restErr)
		return lead, restErr
	}

	return lead, nil
}

func (s *leadService) CreateLead(lead entity.Lead) (leadID int64, restErr resterrors.RestErr) {

	format.FirstLetterUpperCase(&lead.Name)
	leadID, restErr = s.svc.dm.MySQL().Lead().CreateLead(lead)
	if restErr != nil {
		logger.Error("leadService.CreateLead.CreateLead: ", restErr)
		return leadID, restErr
	}

	return leadID, nil
}

func (s *leadService) CreateLeadAddress(leadAddress entity.LeadAddress) resterrors.RestErr {

	format.FirstLetterUpperCase(&leadAddress.Street)
	format.FirstLetterUpperCase(&leadAddress.Neighborhood)
	format.FirstLetterUpperCase(&leadAddress.City)
	format.ToUpperCase(&leadAddress.FederativeUnit)
	restErr := s.svc.dm.MySQL().Lead().CreateLeadAddress(leadAddress)
	if restErr != nil {
		logger.Error("leadService.CreateLeadAddress.CreateLeadAddress: ", restErr)
		return restErr
	}

	return nil
}
