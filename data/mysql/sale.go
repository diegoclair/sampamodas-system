package mysql

import (
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type saleRepo struct {
	db connection
}

func newSaleRepo(db connection) *saleRepo {
	return &saleRepo{
		db: db,
	}
}

func (s *saleRepo) CreateSale(sale entity.Sale) (saleID int64, restErr resterrors.RestErr) {
	query := `
		INSERT INTO tab_sale (
			lead_id, 
			freight, 
			payment_method_id, 
			send_method_id
		) 
		VALUES (
			?,
			(CASE WHEN ? > 0 THEN ? ELSE 0.00 END),
			?,
			?
		);
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return saleID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		sale.LeadID,
		sale.Freight,
		sale.Freight,
		sale.PaymentMethodID,
		sale.SendMethodID,
	)
	if err != nil {
		return saleID, mysqlutils.HandleMySQLError(err)
	}

	saleID, err = result.LastInsertId()
	if err != nil {
		return saleID, mysqlutils.HandleMySQLError(err)
	}

	return saleID, nil
}

func (s *saleRepo) CreateSaleProduct(saleProduct entity.SaleProduct) resterrors.RestErr {
	query := `
		INSERT INTO tab_sale_product (
			sale_id, 
			product_stock_id, 
			quantity, 
			price
		) 
		VALUES (
			?,
			?,
			?,
			?
		);
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		saleProduct.SaleID,
		saleProduct.ProductStockID,
		saleProduct.Quantity,
		saleProduct.Price,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *saleRepo) UpdateSaleTotalPrice(saleID int64, totalPrice float64) resterrors.RestErr {
	query := `
		UPDATE 	tab_sale
			SET total_price	= 	?

		WHERE 	id			=	?;
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		saleID,
		totalPrice,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
