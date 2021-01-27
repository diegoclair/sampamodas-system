package mysql

import (
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type productRepo struct {
	db connection
}

// newProductRepo returns a instance of dbrepo
func newProductRepo(db connection) *productRepo {
	return &productRepo{
		db: db,
	}
}

var queryProduct string = `
	SELECT
		tp.id,
		tp.name,
		tp.cost,
		tp.price,
		tp.business_id,
		gender.id,
		gender.name,
		brand.id,
		brand.name

	FROM 	tab_product 		tp

	INNER JOIN 	tab_gender 		gender
			ON 		gender.id		= tp.gender_id

	INNER JOIN 	tab_brand 		brand
		ON 		brand.id		= tp.brand_id
`

func (s *productRepo) CreateProduct(product entity.Product) (productID int64, restErr resterrors.RestErr) {

	query := `
		INSERT INTO tab_product (
			name,
			cost,
			price,
			gender_id,
			brand_id,
			business_id
		) 
		VALUES	
			(?, ?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return productID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		product.Name,
		product.Cost,
		product.Price,
		product.Gender.ID,
		product.Brand.ID,
		product.BusinessID,
	)
	if err != nil {
		return productID, mysqlutils.HandleMySQLError(err)
	}

	productID, err = result.LastInsertId()
	if err != nil {
		return productID, mysqlutils.HandleMySQLError(err)
	}

	return productID, nil
}

func (s *productRepo) GetProducts() (products []entity.Product, restErr resterrors.RestErr) {

	query := queryProduct

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return products, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return products, mysqlutils.HandleMySQLError(err)
	}

	var product entity.Product
	for rows.Next() {
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Cost,
			&product.Price,
			&product.BusinessID,
			&product.Gender.ID,
			&product.Gender.Name,
			&product.Brand.ID,
			&product.Brand.Name,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (s *productRepo) GetProductByID(productID int64) (product entity.Product, restErr resterrors.RestErr) {

	query := queryProduct + `
		WHERE  	tp.id 			= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return product, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(productID)
	if err != nil {
		return product, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&product.ID,
		&product.Name,
		&product.Cost,
		&product.Price,
		&product.BusinessID,
		&product.Gender.ID,
		&product.Gender.Name,
		&product.Brand.ID,
		&product.Brand.Name,
	)
	if err != nil {
		return product, mysqlutils.HandleMySQLError(err)
	}

	return product, nil
}

func (s *productRepo) GetProductIDByProductStockID(producStockID int64) (productID int64, restErr resterrors.RestErr) {

	query := `
		SELECT 	tps.product_id 
		FROM 	tab_product_stock 	tps
		WHERE 	tps.id 				= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return productID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(
		producStockID,
	)

	err = row.Scan(
		&productID,
	)
	if err != nil {
		return productID, mysqlutils.HandleMySQLError(err)
	}

	return productID, nil
}

func (s *productRepo) GetStockProductByProductID(productID int64) (productsStock []entity.ProductStock, restErr resterrors.RestErr) {

	query := `
		SELECT
			tps.id,
			color.id,
			color.name,
			tps.size,
			tps.quantity

		FROM 	tab_product_stock	tps

		INNER JOIN 	tab_color 		color
			ON 		color.id		= tps.color_id


		WHERE  	tps.product_id 			= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return productsStock, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(productID)
	if err != nil {
		return productsStock, mysqlutils.HandleMySQLError(err)
	}

	var productStock entity.ProductStock
	for rows.Next() {
		err = rows.Scan(
			&productStock.ID,
			&productStock.Color.ID,
			&productStock.Color.Name,
			&productStock.Size,
			&productStock.Quantity,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		productsStock = append(productsStock, productStock)
	}

	return productsStock, nil
}

func (s *productRepo) CreateProductStock(productID int64, productStock entity.ProductStock) resterrors.RestErr {

	query := `
		INSERT INTO tab_product_stock (
			product_id,
			color_id,
			size,
			quantity
		) 
		VALUES	
			(?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		productID,
		productStock.Color.ID,
		productStock.Size,
		productStock.Quantity,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *productRepo) CreateBrand(brandName string) (brandID int64, restErr resterrors.RestErr) {

	query := `
		INSERT INTO tab_brand (
			name
		) 
		VALUES	
			(?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return brandID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		brandName,
	)
	if err != nil {
		return brandID, mysqlutils.HandleMySQLError(err)
	}

	brandID, err = result.LastInsertId()
	if err != nil {
		return brandID, mysqlutils.HandleMySQLError(err)
	}

	return brandID, nil
}

func (s *productRepo) GetBrandByName(brandName string) (brandID int64, restErr resterrors.RestErr) {

	query := `
		SELECT
			tb.id

		FROM 	tab_brand 	tb

		WHERE  	tb.name 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return brandID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(brandName)
	if err != nil {
		return brandID, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&brandID,
	)
	if err != nil {
		return brandID, mysqlutils.HandleMySQLError(err)
	}

	return brandID, nil
}

func (s *productRepo) CreateColor(colorName string) (colorID int64, restErr resterrors.RestErr) {

	query := `
		INSERT INTO tab_color (
			name
		) 
		VALUES	
			(?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return colorID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		colorName,
	)
	if err != nil {
		return colorID, mysqlutils.HandleMySQLError(err)
	}

	colorID, err = result.LastInsertId()
	if err != nil {
		return colorID, mysqlutils.HandleMySQLError(err)
	}

	return colorID, nil
}

func (s *productRepo) GetColorByName(colorName string) (colorID int64, restErr resterrors.RestErr) {

	query := `
		SELECT
			tc.id

		FROM 	tab_color 	tc

		WHERE  	tc.name 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return colorID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(colorName)
	if err != nil {
		return colorID, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&colorID,
	)
	if err != nil {
		return colorID, mysqlutils.HandleMySQLError(err)
	}

	return colorID, nil
}

func (s *productRepo) CreateGender(genderName string) (genderID int64, restErr resterrors.RestErr) {

	query := `
		INSERT INTO tab_gender (
			name
		) 
		VALUES	
			(?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return genderID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		genderName,
	)
	if err != nil {
		return genderID, mysqlutils.HandleMySQLError(err)
	}

	genderID, err = result.LastInsertId()
	if err != nil {
		return genderID, mysqlutils.HandleMySQLError(err)
	}

	return genderID, nil
}

func (s *productRepo) GetGenderByName(genderName string) (genderID int64, restErr resterrors.RestErr) {

	query := `
		SELECT
			tg.id

		FROM 	tab_gender 	tg

		WHERE  	tg.name 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return genderID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(genderName)
	if err != nil {
		return genderID, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&genderID,
	)
	if err != nil {
		return genderID, mysqlutils.HandleMySQLError(err)
	}

	return genderID, nil
}
