package mysql

import (
	"database/sql"

	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type businessRepo struct {
	db *sql.DB
}

// newBusinessRepo returns a instance of dbrepo
func newBusinessRepo(db *sql.DB) *businessRepo {
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
		logger.Error("businessRepo.CreateBusiness", err)
		return resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		business.CompanyID,
		business.Name,
	)
	if err != nil {
		logger.Error("businessRepo.CreateBusiness", err)
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
