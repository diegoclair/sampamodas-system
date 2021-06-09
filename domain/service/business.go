package service

import (
	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/util/errors"
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

	logger.Info("GetBusinesses: Process Started")
	defer logger.Info("GetBusinesses: Process Finished")

	businesses, err = s.svc.dm.MySQL().Business().GetBusinesses()
	if err != nil {
		logger.Error("GetBusinesses.GetBusinesses: ", err)
		return businesses, err
	}

	return businesses, nil
}

func (s *businessService) GetBusinessByUUID(businessUUID string) (business entity.Business, err error) {

	logger.Info("GetBusinessByUUID: Process Started")
	defer logger.Info("GetBusinessByUUID: Process Finished")

	business, err = s.svc.dm.MySQL().Business().GetBusinessByUUID(businessUUID)
	if err != nil {
		logger.Error("GetBusinessByUUID.GetBusinessByUUID: ", err)
		return business, err
	}

	return business, nil
}

func (s *businessService) GetBusinessesByCompanyUUID(companyUUID string) (businesses []entity.Business, err error) {

	logger.Info("GetBusinessesByCompanyUUID: Process Started")
	defer logger.Info("GetBusinessesByCompanyUUID: Process Finished")

	businesses, err = s.svc.dm.MySQL().Business().GetBusinessesByCompanyUUID(companyUUID)
	if err != nil {
		logger.Error("GetBusinessesByCompanyUUID.GetBusinessesByCompanyUUID: ", err)
		return businesses, err
	}

	return businesses, nil
}

func (s *businessService) CreateBusiness(business entity.Business) error {

	logger.Info("CreateBusiness: Process Started")
	defer logger.Info("CreateBusiness: Process Finished")

	companyID, err := s.svc.dm.MySQL().Company().GetCompanyIDByUUID(business.CompanyUUID)
	if err != nil && !errors.SQLResultIsEmpty(err.Error()) {
		logger.Error("CreateBusiness.GetCompanyIDByUUID: ", err)
		return err
	} else if err != nil && errors.SQLResultIsEmpty(err.Error()) {
		logger.Error("CreateBusiness.GetCompanyIDByUUID: company uuid don't exists", nil)
		return resterrors.NewBadRequestError("Invalid company UUID received", nil)
	}

	business.UUID = uuid.NewV4().String()
	business.CompanyID = companyID

	err = s.svc.dm.MySQL().Business().CreateBusiness(business)
	if err != nil {
		logger.Error("CreateBusiness.CreateBusiness: ", err)
		return err
	}

	return nil
}
