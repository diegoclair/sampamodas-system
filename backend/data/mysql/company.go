package mysql

import (
	"database/sql"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
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

func (s *companyRepo) GetCompanyIDByUUID(uuid string) (int64, resterrors.RestErr) {

	query := `
		SELECT 	tcp.id

		FROM 	tab_company_partners 	tcp
		WHERE  	tcp.uuid = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		logger.Error("GetCompanyIDByUUID", err)
		return 0, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(uuid)
	var companyID int64
	err = result.Scan(
		&companyID,
	)

	if err != nil {
		logger.Error("GetCompanyIDByUUID", err)
		return 0, mysqlutils.HandleMySQLError(err)
	}

	return companyID, nil
}
