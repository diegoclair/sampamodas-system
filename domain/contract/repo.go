package contract

import (
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

//RepoManager defines the repository aggregator interface
type RepoManager interface {
	Business() BusinessRepo
	Company() CompanyRepo
	Lead() LeadRepo
	Product() ProductRepo
	Sale() SaleRepo
}

// BusinessRepo defines the data set for business repo
type BusinessRepo interface {
	CreateBusiness(company entity.Business) resterrors.RestErr
	GetBusinesses() (businesses []entity.Business, restErr resterrors.RestErr)
	GetBusinessByID(businessID int64) (businesses entity.Business, restErr resterrors.RestErr)
	GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr)
}

// CompanyRepo defines the data set for company repo
type CompanyRepo interface {
	CreateCompany(company entity.Company) resterrors.RestErr
	GetCompanies() (companies []entity.Company, restErr resterrors.RestErr)
	GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr)
}

// LeadRepo defines the data set for lead
type LeadRepo interface {
	GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, restErr resterrors.RestErr)
	GetLeadAddressByLeadID(leadID int64) (addresses []entity.LeadAddress, restErr resterrors.RestErr)
	CreateLead(lead entity.Lead) (leadID int64, restErr resterrors.RestErr)
	CreateLeadAddress(leadAddress entity.LeadAddress) resterrors.RestErr
}

// ProductRepo defines the data set for product repo
type ProductRepo interface {
	CreateProduct(product entity.Product) (productID int64, restErr resterrors.RestErr)
	GetProducts() (products []entity.Product, restErr resterrors.RestErr)
	GetProductByID(productID int64) (product entity.Product, restErr resterrors.RestErr)
	GetProductIDByProductStockID(producStockID int64) (productID int64, restErr resterrors.RestErr)
	RegisterStockInput(productStockID, quantity int64) resterrors.RestErr

	GetAvailableQuantityByProductStockID(productStockID int64) (availableQuantity int64, restErr resterrors.RestErr)
	UpdateAvailableQuantityByProductStockID(productStockID, quantity int64) resterrors.RestErr

	CreateProductStock(productID int64, product entity.ProductStock) (productStockID int64, restErr resterrors.RestErr)
	GetStockProductByProductID(productID int64) (product []entity.ProductStock, restErr resterrors.RestErr)

	CreateBrand(brandName string) (brandID int64, restErr resterrors.RestErr)
	GetBrandByName(brandName string) (brandID int64, restErr resterrors.RestErr)

	CreateColor(colorName string) (colorID int64, restErr resterrors.RestErr)
	GetColorByName(colorName string) (colorID int64, restErr resterrors.RestErr)

	CreateGender(genderName string) (genderID int64, restErr resterrors.RestErr)
	GetGenderByName(genderName string) (genderID int64, restErr resterrors.RestErr)
}

// SaleRepo defines the data set for sale
type SaleRepo interface {
	CreateSale(sale entity.Sale) (saleID int64, restErr resterrors.RestErr)
	CreateSaleProduct(saleProduct entity.SaleProduct) resterrors.RestErr
	UpdateSaleTotalPrice(saleID int64, totalPrice float64) resterrors.RestErr
}
