package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type saleService struct {
	svc            *Service
	productService contract.ProductService
}

func newSaleService(svc *Service, productService contract.ProductService) contract.SaleService {
	return &saleService{
		svc:            svc,
		productService: productService,
	}
}

func (s *saleService) CreateSale(sale entity.Sale) (saleID int64, restErr resterrors.RestErr) {

	tx, txErr := s.svc.db.Begin()
	if txErr != nil {
		logger.Error("saleService.CreateSale.Begin: ", txErr)
		return saleID, resterrors.NewInternalServerError("Database transaction error")
	}
	defer tx.Rollback()

	saleID, restErr = tx.Sale().CreateSale(sale)
	if restErr != nil {
		logger.Error("saleService.CreateSale.CreateSale", restErr)
		return saleID, restErr
	}

	var totalPrice float64
	for i := range sale.SaleProduct {

		product, restErr := s.productService.GetProductByProductStockID(sale.SaleProduct[i].ProductStockID)
		if restErr != nil {
			noRecorsIdx := domain.NoRecordsFindRE.FindStringIndex(restErr.Error())
			if len(noRecorsIdx) > 0 {
				logger.Error("ProductStockID is invalid", restErr)
				return saleID, resterrors.NewBadRequestError("O ID do estoque do produto é inválido, contate o administrador.")
			}
			logger.Error("saleService.CreateSale.GetProductByProductStockID", restErr)
			return saleID, restErr
		}

		sale.SaleProduct[i].Price = product.Price
		sale.SaleProduct[i].SaleID = saleID

		restErr = tx.Sale().CreateSaleProduct(sale.SaleProduct[i])
		if restErr != nil {
			logger.Error("saleService.CreateSale.CreateSaleProduct", restErr)
			return saleID, restErr
		}

		totalPrice += product.Price
	}

	restErr = tx.Sale().UpdateSaleTotalPrice(saleID, totalPrice)
	if restErr != nil {
		logger.Error("saleService.CreateSale.UpdateSaleTotalPrice", restErr)
		return saleID, restErr
	}

	txErr = tx.Commit()
	if txErr != nil {
		logger.Error("saleService.CreateSale.Commit: ", txErr)
		return saleID, resterrors.NewInternalServerError("Database transaction commit error")
	}

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
