package service

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type saleService struct {
	svc *Service
}

func newSaleService(svc *Service) contract.SaleService {
	return &saleService{
		svc: svc,
	}
}

func (s *saleService) CreateSale(sale entity.Sale) (saleID int64, restErr resterrors.RestErr) {
	return saleID, nil
}

func (s *saleService) CreateSaleProduct(saleProduct entity.SaleProduct) resterrors.RestErr {
	return nil
}

func (s *saleService) GetSales() (sales []entity.Sale, restErr resterrors.RestErr) {
	return sales, nil
}

func (s *saleService) GetSaleByID(saleID int64) (sale entity.Sale, restErr resterrors.RestErr) {
	return sale, nil
}
