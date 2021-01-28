package mysql

import (
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
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

func (s *businessRepo) CreateBusiness(business entity.Business) resterrors.RestErr {

	query := `
		INSERT INTO tab_business (
			company_id,
			name
		) 
		VALUES	
			(?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		business.CompanyID,
		business.Name,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *businessRepo) GetBusinesses() (businesses []entity.Business, restErr resterrors.RestErr) {

	query := `
		SELECT
			tb.business_id,
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

	var business entity.Business
	for rows.Next() {
		err = rows.Scan(
			&business.ID,
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

func (s *businessRepo) GetBusinessByID(businessID int64) (business entity.Business, restErr resterrors.RestErr) {

	query := `
		SELECT
			tb.business_id,
			tb.company_id,
			tb.name

		FROM 	tab_business 	tb
		WHERE  	tb.business_id 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return business, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(businessID)
	if err != nil {
		return business, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&business.ID,
		&business.CompanyID,
		&business.Name,
	)
	if err != nil {
		return business, mysqlutils.HandleMySQLError(err)
	}

	return business, nil
}

func (s *businessRepo) GetBusinessesByCompanyID(companyID int64) (businesses []entity.Business, restErr resterrors.RestErr) {

	query := `
		SELECT
			tb.business_id,
			tb.company_id,
			tb.name

		FROM 	tab_business	tb
		WHERE 	tb.company_id	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return businesses, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(companyID)
	if err != nil {
		return businesses, mysqlutils.HandleMySQLError(err)
	}

	var business entity.Business
	for rows.Next() {
		err = rows.Scan(
			&business.ID,
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
