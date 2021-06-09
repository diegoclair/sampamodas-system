package service

import (
	"fmt"

	"github.com/diegoclair/go_utils-lib/v2/logger"
	"github.com/diegoclair/go_utils-lib/v2/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/util/errors"
)

type saleService struct {
	svc            *Service
	productService ProductService
}

func newSaleService(svc *Service, productService ProductService) SaleService {
	return &saleService{
		svc:            svc,
		productService: productService,
	}
}

func (s *saleService) CreateSale(sale entity.Sale) (saleID int64, err error) {

	logger.Info("CreateSale: Process Started")
	defer logger.Info("CreateSale: Process Finished")

	tx, err := s.svc.dm.MySQL().Begin()
	if err != nil {
		logger.Error("CreateSale.Begin: ", err)
		return saleID, resterrors.NewInternalServerError("Database transaction error", err)
	}
	defer tx.Rollback()

	saleID, err = tx.Sale().CreateSale(sale)
	if err != nil {
		logger.Error("CreateSale.CreateSale", err)
		return saleID, err
	}

	var totalPrice float64
	for i := range sale.SaleProduct {

		product, err := s.productService.GetProductByProductStockID(sale.SaleProduct[i].ProductStockID)
		if err != nil && errors.SQLResultIsEmpty(err.Error()) {
			logger.Error("ProductStockID is invalid", err)
			return saleID, resterrors.NewBadRequestError("O ID do estoque do produto é inválido, contate o administrador.", err)

		}
		if err != nil {
			logger.Error("CreateSale.GetProductByProductStockID", err)
			return saleID, err
		}

		err = s.removeStockAvailableQuantity(sale.SaleProduct[i].ProductStockID, sale.SaleProduct[i].Quantity, tx)
		if err != nil {
			return saleID, err
		}
		sale.SaleProduct[i].Price = product.Price
		sale.SaleProduct[i].SaleID = saleID

		err = tx.Sale().CreateSaleProduct(sale.SaleProduct[i])
		if err != nil {
			logger.Error("CreateSale.CreateSaleProduct", err)
			return saleID, err
		}

		totalPrice += product.Price
	}

	err = tx.Sale().UpdateSaleTotalPrice(saleID, totalPrice)
	if err != nil {
		logger.Error("CreateSale.UpdateSaleTotalPrice", err)
		return saleID, err
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("CreateSale.Commit: ", err)
		return saleID, resterrors.NewInternalServerError("Database transaction commit error", err)
	}

	return saleID, nil
}

func (s *saleService) removeStockAvailableQuantity(productStockID, quantity int64, tx contract.MysqlTransaction) error {

	logger.Info("removeStockAvailableQuantity: Process Started")
	defer logger.Info("removeStockAvailableQuantity: Process Finished")

	actualAvailableQuantity, err := s.svc.dm.MySQL().Product().GetAvailableQuantityByProductStockID(productStockID)
	if err != nil && !errors.SQLResultIsEmpty(err.Error()) {
		logger.Error("removeStockAvailableQuantity.GetAvailableQuantityByProductStockID: ", err)
		return err
	}

	if quantity > actualAvailableQuantity {
		logger.Error(fmt.Sprintf("The sale quantity is bigger than the stock quantity: saleQuantity: %v - availableQuantity: %v", quantity, actualAvailableQuantity), nil)
		return resterrors.NewBadRequestError("A quantidade de venda do produto não pode ser superior à quantidade disponível no estoque.", nil)
	}
	availableQuantity := actualAvailableQuantity - quantity

	err = tx.Product().UpdateAvailableQuantityByProductStockID(productStockID, availableQuantity)
	if err != nil {
		logger.Error("removeStockAvailableQuantity.UpdateAvailableQuantityByProductStockID: ", err)
		return err
	}

	return nil
}

func (s *saleService) CreateSaleProduct(saleProduct entity.SaleProduct) error {

	logger.Info("CreateSaleProduct: Process Started")
	defer logger.Info("CreateSaleProduct: Process Finished")

	return nil
}

func (s *saleService) GetSales() (sales []entity.Sale, err error) {

	logger.Info("GetSales: Process Started")
	defer logger.Info("GetSales: Process Finished")

	return sales, nil
}

func (s *saleService) GetSaleByID(saleID int64) (sale entity.Sale, err error) {

	logger.Info("GetSaleByID: Process Started")
	defer logger.Info("GetSaleByID: Process Finished")

	return sale, nil
}
