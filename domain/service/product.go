package service

import (
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
	"github.com/diegoclair/sampamodas-system/backend/util/errors"
	"github.com/diegoclair/sampamodas-system/backend/util/format"
)

type productService struct {
	svc *Service
}

//newProductService return a new instance of the service
func newProductService(svc *Service) contract.ProductService {
	return &productService{
		svc: svc,
	}
}

func (s *productService) GetProducts() (products []entity.Product, err resterrors.RestErr) {

	products, err = s.svc.dm.MySQL().Product().GetProducts()
	if err != nil {
		logger.Error("productService.GetProducts.GetProducts", err)
		return products, err
	}

	for i := range products {
		products[i].ProductStock, err = s.svc.dm.MySQL().Product().GetStockProductByProductID(products[i].ID)
		if err != nil {
			logger.Error("productService.GetProducts.GetStockProductByProductID", err)
			return products, err
		}
	}

	return
}

func (s *productService) GetProductByID(productID int64) (product entity.Product, restErr resterrors.RestErr) {

	product, restErr = s.svc.dm.MySQL().Product().GetProductByID(productID)
	if restErr != nil {
		logger.Error("productService.GetProductByID.GetProductByID", restErr)
		return product, restErr
	}
	return product, nil
}

func (s *productService) GetProductByProductStockID(productStockID int64) (product entity.Product, restErr resterrors.RestErr) {

	productID, restErr := s.svc.dm.MySQL().Product().GetProductIDByProductStockID(productStockID)
	if restErr != nil {
		logger.Error("productService.GetProductByProductStockID.GetProductIDByProductStockID", restErr)
		return product, restErr
	}

	product, restErr = s.svc.dm.MySQL().Product().GetProductByID(productID)
	if restErr != nil {
		logger.Error("productService.GetProductByProductStockID.GetProductByID", restErr)
		return product, restErr
	}

	product.ProductStock, restErr = s.svc.dm.MySQL().Product().GetStockProductByProductID(productID)
	if restErr != nil {
		logger.Error("productService.GetProducts.GetStockProductByProductID", restErr)
		return product, restErr
	}

	return
}

func (s *productService) CreateProduct(product entity.Product) (err resterrors.RestErr) {

	product.Brand.ID, err = s.getBrandIDByName(product.Brand.Name)
	if err != nil {
		return err
	}

	product.Gender.ID, err = s.getGenderIDByName(product.Gender.Name)
	if err != nil {
		return err
	}

	tx, txErr := s.svc.dm.MySQL().Begin()
	if txErr != nil {
		logger.Error("productService.CreateProduct.Begin: ", txErr)
		return resterrors.NewInternalServerError("Database transaction error")
	}
	defer tx.Rollback()

	format.FirstLetterUpperCase(&product.Name)

	product.ID, err = tx.Product().CreateProduct(product)
	if err != nil {
		logger.Error("productService.CreateProduct.CreateProduct: ", err)
		return err
	}

	for i := range product.ProductStock {

		format.FirstLetterUpperCase(&product.ProductStock[i].Size)

		product.ProductStock[i].Color.ID, err = s.getColorIDByName(product.ProductStock[i].Color.Name)
		if err != nil {
			return err
		}

		productStockID, err := tx.Product().CreateProductStock(product.ID, product.ProductStock[i])
		if err != nil {
			logger.Error("productService.CreateProduct.CreateProductStock: ", err)
			return err
		}

		err = s.registerStockInput(productStockID, product.ProductStock[i].InputQuantity, tx)
		if err != nil {
			return err
		}

		err = s.addStockAvailableQuantity(productStockID, product.ProductStock[i].InputQuantity, tx)
		if err != nil {
			return err
		}

	}

	txErr = tx.Commit()
	if txErr != nil {
		logger.Error("productService.CreateProduct.Commit: ", txErr)
		return resterrors.NewInternalServerError("Database transaction commit error")
	}

	return nil
}

func (s *productService) addStockAvailableQuantity(productStockID, quantity int64, tx contract.MysqlTransaction) (restErr resterrors.RestErr) {

	actualAvailableQuantity, restErr := s.svc.dm.MySQL().Product().GetAvailableQuantityByProductStockID(productStockID)
	if restErr != nil && !errors.SQLResultIsEmpty(restErr.Message()) {
		logger.Error("productService.addStockAvailableQuantity.GetAvailableQuantityByProductStockID: ", restErr)
		return restErr
	}

	availableQuantity := actualAvailableQuantity + quantity

	restErr = tx.Product().UpdateAvailableQuantityByProductStockID(productStockID, availableQuantity)
	if restErr != nil {
		logger.Error("productService.addStockAvailableQuantity.UpdateAvailableQuantityByProductStockID: ", restErr)
		return restErr
	}

	return nil
}

func (s *productService) registerStockInput(productStockID, quantity int64, tx contract.MysqlTransaction) (restErr resterrors.RestErr) {

	restErr = tx.Product().RegisterStockInput(productStockID, quantity)
	if restErr != nil {
		logger.Error("productService.registerStockInput.RegisterStockInput: ", restErr)
		return restErr
	}

	return nil
}

func (s *productService) getBrandIDByName(brandName string) (brandID int64, err resterrors.RestErr) {

	format.FirstLetterUpperCase(&brandName)
	brandID, err = s.svc.dm.MySQL().Product().GetBrandByName(brandName)
	if err != nil && !errors.SQLResultIsEmpty(err.Message()) {
		logger.Error("getColorIDByName.GetBrandByName: ", err)
		return brandID, err
	}

	if brandID > 0 {
		return brandID, nil
	}

	brandID, err = s.svc.dm.MySQL().Product().CreateBrand(brandName)
	if err != nil {
		logger.Error("getColorIDByName.CreateBrand: ", err)
		return brandID, err
	}

	return brandID, nil
}

func (s *productService) getColorIDByName(colorName string) (colorID int64, err resterrors.RestErr) {

	format.FirstLetterUpperCase(&colorName)
	colorID, err = s.svc.dm.MySQL().Product().GetColorByName(colorName)
	if err != nil && !errors.SQLResultIsEmpty(err.Message()) {
		logger.Error("getColorIDByName.GetColorByName: ", err)
		return colorID, err
	}

	if colorID > 0 {
		return colorID, nil
	}

	colorID, err = s.svc.dm.MySQL().Product().CreateColor(colorName)
	if err != nil {
		logger.Error("getColorIDByName.CreateColor: ", err)
		return colorID, err
	}

	return colorID, nil
}

func (s *productService) getGenderIDByName(genderName string) (genderID int64, err resterrors.RestErr) {

	format.FirstLetterUpperCase(&genderName)
	genderID, err = s.svc.dm.MySQL().Product().GetGenderByName(genderName)
	if err != nil && !errors.SQLResultIsEmpty(err.Message()) {
		logger.Error("getColorIDByName.GetGenderByName: ", err)
		return genderID, err
	}

	if genderID > 0 {
		return genderID, nil
	}

	genderID, err = s.svc.dm.MySQL().Product().CreateGender(genderName)
	if err != nil {
		logger.Error("getColorIDByName.CreateGender: ", err)
		return genderID, err
	}

	return genderID, nil
}
