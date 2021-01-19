package data

import (
	"github.com/diegoclair/sampamodas-system/backend/data/mysql"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
