package mysql

import (
	"github.com/diegoclair/go_utils-lib/v2/mysqlutils"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type businessRepo struct {
	db connection
}

// newBusinessRepo returns a instance of dbrepo
func newBusinessRepo(db connection) *businessRepo {
	return &businessRepo{
		db: db,
	}
}

func (s *businessRepo) CreateBusiness(business entity.Business) error {

	query := `
		INSERT INTO tab_business (
			business_uuid,
			company_id,
			name
		) 
		VALUES	
			(?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		business.UUID,
		business.CompanyID,
		business.Name,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *businessRepo) GetBusinesses() (businesses []entity.Business, err error) {

	query := `
		SELECT
			tb.business_id,
			tb.business_uuid,
			tb.company_id,
			tb.name

		FROM 	tab_business 	tb
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return businesses, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return businesses, mysqlutils.HandleMySQLError(err)
	}

	for rows.Next() {
		business := entity.Business{}
		err = rows.Scan(
			&business.ID,
			&business.UUID,
			&business.CompanyID,
			&business.Name,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		businesses = append(businesses, business)
	}

	return businesses, nil
}

func (s *businessRepo) GetBusinessByUUID(businessUUID string) (business entity.Business, err error) {

	query := `
		SELECT
			tb.business_id,
			tb.business_uuid,
			tb.company_id,
			tb.name

		FROM 	tab_business 		tb
		WHERE  	tb.business_uuid 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return business, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(businessUUID)
	if err != nil {
		return business, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&business.ID,
		&business.UUID,
		&business.CompanyID,
		&business.Name,
	)
	if err != nil {
		return business, mysqlutils.HandleMySQLError(err)
	}

	return business, nil
}

func (s *businessRepo) GetBusinessesByCompanyUUID(companyUUID string) (businesses []entity.Business, err error) {

	query := `
		SELECT
			tb.business_id,
			tb.business_uuid,
			tb.company_id,
			tb.name

		FROM 	tab_business	tb

		INNER JOIN tab_company 	tc
		ON 		tb.company_id 	= 	tc.company_id
		
		WHERE 	tc.company_uuid	= 	?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return businesses, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(companyUUID)
	if err != nil {
		return businesses, mysqlutils.HandleMySQLError(err)
	}

	for rows.Next() {
		business := entity.Business{}
		err = rows.Scan(
			&business.ID,
			&business.UUID,
			&business.CompanyID,
			&business.Name,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		businesses = append(businesses, business)
	}

	return businesses, nil
}
