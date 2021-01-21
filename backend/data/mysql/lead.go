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

func (s *leadRepo) GetLeadAddressByLeadID(leadID int64) (addresses []entity.LeadAddress, restErr resterrors.RestErr) {

	query := `
		SELECT
			tla.id,
			tla.lead_id,
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
			&address.LeadID,
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

func (s *leadRepo) CreateLead(lead entity.Lead) (leadID int64, restErr resterrors.RestErr) {

	query := `
		INSERT INTO tab_lead (
			name,
			email,
			phone_number,
			instagram
		) 
		VALUES	
			(?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return leadID, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		lead.Name,
		lead.Email,
		lead.PhoneNumber,
		lead.Instagram,
	)
	if err != nil {
		return leadID, mysqlutils.HandleMySQLError(err)
	}

	leadID, err = result.LastInsertId()
	if err != nil {
		return leadID, mysqlutils.HandleMySQLError(err)
	}

	return leadID, nil
}

func (s *leadRepo) CreateLeadAddress(leadAddress entity.LeadAddress) resterrors.RestErr {

	query := `
		INSERT INTO tab_lead_address (
			lead_id,
			address_type,
			street,
			number,
			neighborhood,
			complement,
			city,
			federative_unit,
			zip_code
		) 
		VALUES	
			(?, ?, ?, ?, ?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		leadAddress.LeadID,
		leadAddress.AddressType,
		leadAddress.Street,
		leadAddress.Number,
		leadAddress.Neighborhood,
		leadAddress.Complement,
		leadAddress.City,
		leadAddress.FederativeUInit,
		leadAddress.ZipCode,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
