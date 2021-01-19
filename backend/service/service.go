package service

import (
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
)

// Service holds the domain service repositories
type Service struct {
	db contract.RepoManager
}

// New returns a new domain Service instance
func New(db contract.RepoManager) *Service {
	svc := new(Service)
	svc.db = db

	return svc
}

//Manager defines the services aggregator interface
type Manager interface {
	LeadService(svc *Service) contract.LeadService
	CompanyService(svc *Service) contract.CompanyService
}

type serviceManager struct {
	svc *Service
}

// NewServiceManager return a service manager instance
func NewServiceManager() Manager {
	return &serviceManager{}
}

func (s *serviceManager) LeadService(svc *Service) contract.LeadService {
	return newLeadService(svc)
}

func (s *serviceManager) CompanyService(svc *Service) contract.CompanyService {
	return newCompanyService(svc)
}
