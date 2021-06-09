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
func newProductService(svc *Service) ProductService {
	return &productService{
		svc: svc,
	}
}

func (s *productService) GetProducts() (products []entity.Product, err error) {

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

func (s *productService) GetProductByID(productID int64) (product entity.Product, err error) {

	product, err = s.svc.dm.MySQL().Product().GetProductByID(productID)
	if err != nil {
		logger.Error("productService.GetProductByID.GetProductByID", err)
		return product, err
	}
	return product, nil
}

func (s *productService) GetProductByProductStockID(productStockID int64) (product entity.Product, err error) {

	productID, err := s.svc.dm.MySQL().Product().GetProductIDByProductStockID(productStockID)
	if err != nil {
		logger.Error("productService.GetProductByProductStockID.GetProductIDByProductStockID", err)
		return product, err
	}

	product, err = s.svc.dm.MySQL().Product().GetProductByID(productID)
	if err != nil {
		logger.Error("productService.GetProductByProductStockID.GetProductByID", err)
		return product, err
	}

	product.ProductStock, err = s.svc.dm.MySQL().Product().GetStockProductByProductID(productID)
	if err != nil {
		logger.Error("productService.GetProducts.GetStockProductByProductID", err)
		return product, err
	}

	return
}

func (s *productService) CreateProduct(product entity.Product) (err error) {

	product.Brand.ID, err = s.getBrandIDByName(product.Brand.Name)
	if err != nil {
		return err
	}

	product.Gender.ID, err = s.getGenderIDByName(product.Gender.Name)
	if err != nil {
		return err
	}

	tx, err := s.svc.dm.MySQL().Begin()
	if err != nil {
		logger.Error("productService.CreateProduct.Begin: Database transaction error: ", err)
		return err
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

	err = tx.Commit()
	if err != nil {
		logger.Error("productService.CreateProduct.Commit: ", err)
		return resterrors.NewInternalServerError("Database transaction commit error")
	}

	return nil
}

func (s *productService) addStockAvailableQuantity(productStockID, quantity int64, tx contract.MysqlTransaction) (err error) {

	actualAvailableQuantity, err := s.svc.dm.MySQL().Product().GetAvailableQuantityByProductStockID(productStockID)
	if err != nil && !errors.SQLResultIsEmpty(err.Error()) {
		logger.Error("productService.addStockAvailableQuantity.GetAvailableQuantityByProductStockID: ", err)
		return err
	}

	availableQuantity := actualAvailableQuantity + quantity

	err = tx.Product().UpdateAvailableQuantityByProductStockID(productStockID, availableQuantity)
	if err != nil {
		logger.Error("productService.addStockAvailableQuantity.UpdateAvailableQuantityByProductStockID: ", err)
		return err
	}

	return nil
}

func (s *productService) registerStockInput(productStockID, quantity int64, tx contract.MysqlTransaction) (err error) {

	err = tx.Product().RegisterStockInput(productStockID, quantity)
	if err != nil {
		logger.Error("productService.registerStockInput.RegisterStockInput: ", err)
		return err
	}

	return nil
}

func (s *productService) getBrandIDByName(brandName string) (brandID int64, err error) {

	format.FirstLetterUpperCase(&brandName)
	brandID, err = s.svc.dm.MySQL().Product().GetBrandByName(brandName)
	if err != nil && !errors.SQLResultIsEmpty(err.Error()) {
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

func (s *productService) getColorIDByName(colorName string) (colorID int64, err error) {

	format.FirstLetterUpperCase(&colorName)
	colorID, err = s.svc.dm.MySQL().Product().GetColorByName(colorName)
	if err != nil && !errors.SQLResultIsEmpty(err.Error()) {
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

func (s *productService) getGenderIDByName(genderName string) (genderID int64, err error) {

	format.FirstLetterUpperCase(&genderName)
	genderID, err = s.svc.dm.MySQL().Product().GetGenderByName(genderName)
	if err != nil && !errors.SQLResultIsEmpty(err.Error()) {
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
