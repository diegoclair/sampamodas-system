package mysql

import (
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type leadRepo struct {
	db connection
}

// newLeadRepo returns a instance of dbrepo
func newLeadRepo(db connection) *leadRepo {
	return &leadRepo{
		db: db,
	}
}

func (s *leadRepo) GetLeadAddress(leadID int64) (addresses []entity.Address, restErr resterrors.RestErr) {

	query := `
		SELECT
		tua.id,
		tua.street,
		tua.number,
		tua.complement,
		tua.zip_code,
		tua.city,
		tua.federative_unit

		FROM 	tab_lead_address 	tua
		WHERE  	tua.lead_id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return addresses, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(leadID)
	if err != nil {
		return addresses, resterrors.NewInternalServerError("Database error")
	}

	var address entity.Address
	for rows.Next() {
		err = rows.Scan(
			&address.ID,
			&address.Street,
			&address.Number,
			&address.Complement,
			&address.ZipCode,
			&address.City,
			&address.FederativeUnit,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (s *leadRepo) CreateSale(sale entity.Sale) resterrors.RestErr {

	query := `
		INSERT INTO tab_sale (
			lead_id,
			company_id,
			price,
			freight,
			qrcode_id,
			product_id,
			payment_type,
			address_id
		) 
		VALUES	
			(?, ?, ?, ?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		sale.LeadID,
		sale.CompanyID,
		sale.Price,
		sale.Freight,
		sale.ProductID,
		sale.PaymentMethodID,
		sale.AddressID,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *leadRepo) GetSaleSummary(leadID int64) (summary []entity.SaleSummary, restErr resterrors.RestErr) {

	query := `
		SELECT
			tcp.name,
			o.price,
			o.freight

		FROM 	tab_sale 	o

		INNER JOIN tab_company_partners tcp
			ON o.company_id = tcp.id

		WHERE  	o.lead_id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return summary, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(leadID)
	if err != nil {
		return summary, resterrors.NewInternalServerError("Database error")
	}

	var saleSummary entity.SaleSummary
	for rows.Next() {
		err = rows.Scan(
			&saleSummary.CompanyName,
			&saleSummary.Price,
			&saleSummary.Freight,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		summary = append(summary, saleSummary)
	}

	return summary, nil
}
