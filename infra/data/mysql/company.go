package mysql

import (
	"github.com/diegoclair/go_utils-lib/v2/mysqlutils"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type companyRepo struct {
	db connection
}

// newCompanyRepo returns a instance of dbrepo
func newCompanyRepo(db connection) *companyRepo {
	return &companyRepo{
		db: db,
	}
}

func (s *companyRepo) CreateCompany(company entity.Company) error {

	query := `
		INSERT INTO tab_company (
			company_uuid,
			document_number,
			legal_name,
			commercial_name
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
		company.UUID,
		company.DocumentNumber,
		company.LegalName,
		company.CommercialName,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *companyRepo) GetCompanies() (companies []entity.Company, err error) {

	query := `
		SELECT
			tc.company_id,
			tc.company_uuid,
			tc.document_number,
			tc.legal_name,
			tc.commercial_name

		FROM 	tab_company 	tc
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return companies, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return companies, mysqlutils.HandleMySQLError(err)
	}

	for rows.Next() {
		company := entity.Company{}
		err = rows.Scan(
			&company.ID,
			&company.UUID,
			&company.DocumentNumber,
			&company.LegalName,
			&company.CommercialName,
		)
		if err != nil {
			return nil, mysqlutils.HandleMySQLError(err)
		}
		companies = append(companies, company)
	}

	return companies, nil
}

func (s *companyRepo) GetCompanyByUUID(companyUUID string) (company entity.Company, err error) {

	query := `
		SELECT
			tc.company_id,
			tc.company_uuid,
			tc.document_number,
			tc.legal_name,
			tc.commercial_name

		FROM 	tab_company 	tc
		WHERE  	tc.company_uuid 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return company, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(companyUUID)
	if err != nil {
		return company, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&company.ID,
		&company.UUID,
		&company.DocumentNumber,
		&company.LegalName,
		&company.CommercialName,
	)
	if err != nil {
		return company, mysqlutils.HandleMySQLError(err)
	}

	return company, nil
}

func (s *companyRepo) GetCompanyIDByUUID(companyUUID string) (companyID int64, err error) {

	query := `
		SELECT
			tc.company_id

		FROM 	tab_company 	tc
		WHERE  	tc.company_uuid 	= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return companyID, mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(companyUUID)
	if err != nil {
		return companyID, mysqlutils.HandleMySQLError(err)
	}

	err = result.Scan(
		&companyID,
	)
	if err != nil {
		return companyID, mysqlutils.HandleMySQLError(err)
	}

	return companyID, nil
}
