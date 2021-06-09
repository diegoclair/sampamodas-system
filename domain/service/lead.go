package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/twinj/uuid"
)

type leadService struct {
	svc *Service
}

//newLeadService return a new instance of the service
func newLeadService(svc *Service) LeadService {
	return &leadService{
		svc: svc,
	}
}

func (s *leadService) GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, err error) {

	//the lead should be by business, so a company cant acess other leads ..

	lead, err = s.svc.dm.MySQL().Lead().GetLeadByPhoneNumber(phoneNumber)
	if err != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadByPhoneNumber: ", err)
		return lead, err
	}

	lead.LeadAddress, err = s.svc.dm.MySQL().Lead().GetLeadAddressByLeadID(lead.ID)
	if err != nil {
		logger.Error("leadService.GetLeadByPhoneNumber.GetLeadAddressByLeadID: ", err)
		return lead, err
	}

	return lead, nil
}

func (s *leadService) CreateLead(lead entity.Lead) (leadID int64, err error) {

	lead.UUID = uuid.NewV4().String()
	leadID, err = s.svc.dm.MySQL().Lead().CreateLead(lead)
	if err != nil {
		logger.Error("leadService.CreateLead.CreateLead: ", err)
		return leadID, err
	}

	return leadID, nil
}

func (s *leadService) CreateLeadAddress(leadAddress entity.LeadAddress) error {

	err := s.svc.dm.MySQL().Lead().CreateLeadAddress(leadAddress)
	if err != nil {
		logger.Error("leadService.CreateLeadAddress.CreateLeadAddress: ", err)
		return err
	}

	return nil
}
