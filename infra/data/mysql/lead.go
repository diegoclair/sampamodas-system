package mysql

import (
	"github.com/diegoclair/go_utils-lib/v2/mysqlutils"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type leadRepo struct {
	db connection
}

func newLeadRepo(db connection) *leadRepo {
	return &leadRepo{
		db: db,
	}
}

func (s *leadRepo) GetLeadByPhoneNumber(phoneNumber string) (lead entity.Lead, err error) {

	query := `
		SELECT
			tl.lead_id,
			tl.lead_uuid,
			tl.name,
			tl.email,
			tl.phone_number,
			tl.instagram

		FROM 	tab_lead 		tl
		WHERE  	tl.phone_number = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return lead, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(phoneNumber)
	if err != nil {
		return lead, mysqlutils.HandleMySQLError(err)
	}

	err = row.Scan(
		&lead.ID,
		&lead.UUID,
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

func (s *leadRepo) GetLeadAddressByLeadID(leadID int64) (addresses []entity.LeadAddress, err error) {

	query := `
		SELECT
			tla.lead_address_id,
			tla.lead_id,
			tl.lead_uuid,
			tla.address_type,
			tla.street,
			tla.number,
			tla.neighborhood,
			tla.complement,
			tla.city,
			tla.federative_unit,
			tla.zip_code

		FROM 	tab_lead_address tla
		
		INNER JOIN 	tab_lead 	tl
		ON		tl.lead_id	= tla.lead_id

		WHERE  	tla.lead_id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return addresses, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(leadID)
	if err != nil {
		return addresses, mysqlutils.HandleMySQLError(err)
	}

	var address entity.LeadAddress
	for rows.Next() {
		err = rows.Scan(
			&address.ID,
			&address.LeadID,
			&address.LeadUUID,
			&address.AddressType,
			&address.Street,
			&address.Number,
			&address.Neighborhood,
			&address.Complement,
			&address.City,
			&address.FederativeUnit,
			&address.ZipCode,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		addresses = append(addresses, address)
	}

	return addresses, nil
}

func (s *leadRepo) CreateLead(lead entity.Lead) (leadID int64, err error) {

	query := `
		INSERT INTO tab_lead (
			lead_uuid,
			name,
			email,
			phone_number,
			instagram
		) 
		VALUES	
			(?, ?, ?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return leadID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		lead.UUID,
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

func (s *leadRepo) CreateLeadAddress(leadAddress entity.LeadAddress) error {

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
		return mysqlutils.HandleMySQLError(err)
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
		leadAddress.FederativeUnit,
		leadAddress.ZipCode,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
