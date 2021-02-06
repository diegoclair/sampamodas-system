package data

import (
	"github.com/diegoclair/sampamodas-system/backend/data/mysql"
	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
)

// Connect returns a instace of mysql db
func Connect() (contract.DataManager, error) {
	return mysql.Instance()
}
