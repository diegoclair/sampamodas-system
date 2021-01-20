package service

import (
	"regexp"
	"strings"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
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

	products, err = s.svc.db.Product().GetProducts()
	if err != nil {
		return products, err
	}

	for i := range products {
		products[i].ProductStock, err = s.svc.db.Product().GetStockProductByID(products[i].ID)
		if err != nil {
			return products, err
		}
	}

	return
}

func (s *productService) GetProductByID(productID int64) (product entity.Product, restErr resterrors.RestErr) {
	return s.svc.db.Product().GetProductByID(productID)
}

func (s *productService) CreateProduct(product entity.Product) (err resterrors.RestErr) {

	s.firstLetterUppercase(&product.Name)

	product.Brand.ID, err = s.getBrandIDByName(product.Brand.Name)
	if err != nil {
		return err
	}

	product.Gender.ID, err = s.getGenderIDByName(product.Gender.Name)
	if err != nil {
		return err
	}

	tx, txErr := s.svc.db.Begin()
	if txErr != nil {
		logger.Error("productService.CreateProduct.Begin: ", err)
		return resterrors.NewInternalServerError("Database trasaction error")
	}
	defer tx.Rollback()

	product.ID, err = tx.Product().CreateProduct(product)
	if err != nil {
		logger.Error("productService.CreateProduct.CreateProduct: ", err)
		return err
	}

	for i := range product.ProductStock {

		s.firstLetterUppercase(&product.ProductStock[i].Size)

		product.ProductStock[i].Color.ID, err = s.getColorIDByName(product.ProductStock[i].Color.Name)
		if err != nil {
			return err
		}

		err = tx.Product().CreateProductStock(product.ID, product.ProductStock[i])
		if err != nil {
			logger.Error("productService.CreateProduct.CreateProductStock: ", err)
			return err
		}
	}

	txErr = tx.Commit()
	if txErr != nil {
		logger.Error("productService.CreateProduct.Commit: ", txErr)
		return resterrors.NewInternalServerError("Database trasaction error")
	}

	return nil
}

var noSQLRowsRE = regexp.MustCompile(domain.NoSQLRows)

func (s *productService) getBrandIDByName(brandName string) (brandID int64, err resterrors.RestErr) {

	s.firstLetterUppercase(&brandName)
	brandID, err = s.svc.db.Product().GetBrandByName(brandName)
	if err != nil {
		noRowsIdx := noSQLRowsRE.FindStringIndex(err.Error())
		if len(noRowsIdx) > 0 {
			logger.Error("getColorIDByName.GetBrandByName: ", err)
			return brandID, err
		}
	}

	if brandID > 0 {
		return brandID, nil
	}

	brandID, err = s.svc.db.Product().CreateBrand(brandName)
	if err != nil {
		logger.Error("getColorIDByName.CreateBrand: ", err)
		return brandID, err
	}

	return brandID, nil
}

func (s *productService) getColorIDByName(colorName string) (colorID int64, err resterrors.RestErr) {

	s.firstLetterUppercase(&colorName)
	colorID, err = s.svc.db.Product().GetColorByName(colorName)
	if err != nil {
		noRowsIdx := noSQLRowsRE.FindStringIndex(err.Error())
		if len(noRowsIdx) > 0 {
			logger.Error("getColorIDByName.GetColorByName: ", err)
			return colorID, err
		}
	}

	if colorID > 0 {
		return colorID, nil
	}

	colorID, err = s.svc.db.Product().CreateColor(colorName)
	if err != nil {
		logger.Error("getColorIDByName.CreateColor: ", err)
		return colorID, err
	}

	return colorID, nil
}

func (s *productService) getGenderIDByName(genderName string) (genderID int64, err resterrors.RestErr) {

	s.firstLetterUppercase(&genderName)
	genderID, err = s.svc.db.Product().GetGenderByName(genderName)
	if err != nil {
		noRowsIdx := noSQLRowsRE.FindStringIndex(err.Error())
		if len(noRowsIdx) > 0 {
			logger.Error("getColorIDByName.GetGenderByName: ", err)
			return genderID, err
		}
	}

	if genderID > 0 {
		return genderID, nil
	}

	genderID, err = s.svc.db.Product().CreateGender(genderName)
	if err != nil {
		logger.Error("getColorIDByName.CreateGender: ", err)
		return genderID, err
	}

	return genderID, nil
}

func (s *productService) firstLetterUppercase(str *string) {
	*str = strings.Title(strings.ToLower(*str))
}
