package mysql

import (
	"database/sql"
	"log"

	"github.com/diegoclair/sampamodas-system/backend/contract"
)

// mysqlTransaction is the MySQL transaction manager
type mysqlTransaction struct {
	tx         *sql.Tx
	committed  bool
	rolledback bool
}

func newTransaction(tx *sql.Tx) *mysqlTransaction {
	instance := &mysqlTransaction{tx: tx}
	return instance
}

// Begin starts a transaction
func (t *mysqlTransaction) Begin() (contract.MysqlTransaction, error) {
	return &mysqlTransaction{
		tx: t.tx,
	}, nil
}

func (t *mysqlTransaction) MySQL() contract.MySQLRepo {
	mysqlRepo, err := Instance()
	if err != nil {
		log.Fatalf("Error to start mysql instance: %v", err)
	}
	return mysqlRepo
}

// Commit persists changes to database
func (t *mysqlTransaction) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	t.committed = true

	return nil
}

// Rollback discards changes on database
func (t *mysqlTransaction) Rollback() error {
	if t != nil && !t.committed && !t.rolledback {
		err := t.tx.Rollback()
		if err != nil {
			return err
		}

		t.rolledback = true
	}

	return nil
}

func (t *mysqlTransaction) Business() contract.BusinessRepo {
	return newBusinessRepo(t.tx)
}

func (t *mysqlTransaction) Company() contract.CompanyRepo {
	return newCompanyRepo(t.tx)
}

func (t *mysqlTransaction) Lead() contract.LeadRepo {
	return newLeadRepo(t.tx)
}

func (t *mysqlTransaction) Product() contract.ProductRepo {
	return newProductRepo(t.tx)
}

func (t *mysqlTransaction) Sale() contract.SaleRepo {
	return newSaleRepo(t.tx)
}

func (t *mysqlTransaction) User() contract.UserRepo {
	return newUserRepo(t.tx)
}
