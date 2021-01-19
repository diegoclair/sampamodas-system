package mysql

import (
	"database/sql"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type companyRepo struct {
	db *sql.DB
}

// newCompanyRepo returns a instance of dbrepo
func newCompanyRepo(db *sql.DB) *companyRepo {
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
		logger.Error("companyRepo.CreateCompany", err)
		return resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		company.DocumentNumber,
		company.LegalName,
		company.CommercialName,
	)
	if err != nil {
		logger.Error("companyRepo.CreateCompany", err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
