package contract

import (
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

//MySQLRepo defines the repository aggregator interface
type MySQLRepo interface {
	Begin() (MysqlTransaction, error)
	Business() BusinessRepo
	Company() CompanyRepo
	Lead() LeadRepo
	Product() ProductRepo
	Sale() SaleRepo
}

// MysqlTransaction holds the methods that manipulates the main data, from within a transaction.
type MysqlTransaction interface {
	MySQLRepo
	Rollback() error
	Commit() error
}

// BusinessRepo defines the data set for business repo
type BusinessRepo interface {
	CreateBusiness(company entity.Business) error
	GetBusinesses() (businesses []entity.Business, err error)
	GetBusinessByUUID(businessUUID string) (businesses entity.Business, err error)
	GetBusinessesByCompanyUUID(companyUUID string) (businesses []entity.Business, err error)
}

// CompanyRepo defines the data set for company repo
type CompanyRepo interface {
	CreateCompany(company entity.Company) error
	GetCompanies() (companies []entity.Company, err error)
	GetCompanyByUUID(companyUUID string) (company entity.Company, err error)
	GetCompanyIDByUUID(companyUUID string) (companyID int64, err error)
}

// LeadRepo defines the data set for lead
type LeadRepo interface {
	GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, err error)
	GetLeadAddressByLeadID(leadID int64) (addresses []entity.LeadAddress, err error)
	CreateLead(lead entity.Lead) (leadID int64, err error)
	CreateLeadAddress(leadAddress entity.LeadAddress) error
}

// ProductRepo defines the data set for product repo
type ProductRepo interface {
	CreateProduct(product entity.Product) (productID int64, err error)
	GetProducts() (products []entity.Product, err error)
	GetProductByID(productID int64) (product entity.Product, err error)
	GetProductIDByProductStockID(producStockID int64) (productID int64, err error)
	RegisterStockInput(productStockID, quantity int64) error

	GetAvailableQuantityByProductStockID(productStockID int64) (availableQuantity int64, err error)
	UpdateAvailableQuantityByProductStockID(productStockID, quantity int64) error

	CreateProductStock(productID int64, product entity.ProductStock) (productStockID int64, err error)
	GetStockProductByProductID(productID int64) (product []entity.ProductStock, err error)

	CreateBrand(brandName string) (brandID int64, err error)
	GetBrandByName(brandName string) (brandID int64, err error)

	CreateColor(colorName string) (colorID int64, err error)
	GetColorByName(colorName string) (colorID int64, err error)

	CreateGender(genderName string) (genderID int64, err error)
	GetGenderByName(genderName string) (genderID int64, err error)
}

// SaleRepo defines the data set for sale
type SaleRepo interface {
	CreateSale(sale entity.Sale) (saleID int64, err error)
	CreateSaleProduct(saleProduct entity.SaleProduct) error
	UpdateSaleTotalPrice(saleID int64, totalPrice float64) error
}
