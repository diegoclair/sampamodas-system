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

func (s *leadRepo) GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, restErr resterrors.RestErr) {

	query := `
		SELECT
			tl.id,
			tl.name,
			tl.email,
			tl.phone_number,
			tl.instagram

		FROM 	tab_lead 		tl
		WHERE  	tl.phone_number = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return lead, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	row := stmt.QueryRow(phoneNumber)
	if err != nil {
		return lead, resterrors.NewInternalServerError("Database error")
	}

	err = row.Scan(
		&lead.ID,
		&lead.Name,
		&lead.Email,
		&lead.PhoneNumber,
		&lead.Instagram,
	)
	if err != nil {
		return lead, mysqlutils.HandleMySQLError(err)
	}

	return lead, nil
}

func (s *leadRepo) GetLeadAddressByID(leadID int64) (addresses []entity.LeadAddress, restErr resterrors.RestErr) {

	query := `
		SELECT
			tla.id,
			tla.address_type,
			tla.street,
			tla.number,
			tla.neighborhood,
			tla.complement,
			tla.city,
			tla.federative_unit,
			tla.zip_code

		FROM 	tab_lead_address tla

		WHERE  	tla.lead_id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return addresses, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(leadID)
	if err != nil {
		return addresses, resterrors.NewInternalServerError("Database error")
	}

	var address entity.LeadAddress
	for rows.Next() {
		err = rows.Scan(
			&address.ID,
			&address.AddressType,
			&address.Street,
			&address.Number,
			&address.Neighborhood,
			&address.Complement,
			&address.City,
			&address.FederativeUInit,
			&address.ZipCode,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}
