package mysql

import (
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
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

func (s *companyRepo) CreateCompany(company entity.Company) resterrors.RestErr {

	query := `
		INSERT INTO tab_company (
			document_number,
			legal_name,
			commercial_name
		) 
		VALUES	
			(?, ?, ?);
		`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		company.DocumentNumber,
		company.LegalName,
		company.CommercialName,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}

func (s *companyRepo) GetCompanies() (companies []entity.Company, restErr resterrors.RestErr) {

	query := `
		SELECT
			tc.id,
			tc.document_number,
			tc.legal_name,
			tc.commercial_name

		FROM 	tab_company 	tc
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return companies, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return companies, resterrors.NewInternalServerError("Database error")
	}

	var company entity.Company
	for rows.Next() {
		err = rows.Scan(
			&company.ID,
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

func (s *companyRepo) GetCompanyByID(companyID int64) (company entity.Company, restErr resterrors.RestErr) {

	query := `
		SELECT
			tc.id,
			tc.document_number,
			tc.legal_name,
			tc.commercial_name

		FROM 	tab_company 	tc
		WHERE  	tc.id 			= ?
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return company, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(companyID)
	if err != nil {
		return company, resterrors.NewInternalServerError("Database error")
	}

	err = result.Scan(
		&company.ID,
		&company.DocumentNumber,
		&company.LegalName,
		&company.CommercialName,
	)
	if err != nil {
		return company, mysqlutils.HandleMySQLError(err)
	}

	return company, nil
}
