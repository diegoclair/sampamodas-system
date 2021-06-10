package mysql

import (
	"github.com/diegoclair/go_utils-lib/v2/mysqlutils"
	"github.com/diegoclair/sampamodas-system/backend/domain/entity"
)

type userRepo struct {
	db connection
}

func newUserRepo(db connection) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (s *userRepo) CreateUser(user entity.User) (err error) {
	query := `
		INSERT INTO tab_user (
			user_uuid,
			name,
			email,
			password,
			company_id
		) 
		VALUES (?, ?, ?, ?, NULL);
	`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.UUID,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return mysqlutils.HandleMySQLError(err)
	}

	return nil
}
