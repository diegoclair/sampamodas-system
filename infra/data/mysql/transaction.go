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

//Business returns the business transaction set
func (t *mysqlTransaction) Business() contract.BusinessRepo {
	return newBusinessRepo(t.tx)
}

//Company returns the company transaction set
func (t *mysqlTransaction) Company() contract.CompanyRepo {
	return newCompanyRepo(t.tx)
}

//Lead returns the lead transaction set
func (t *mysqlTransaction) Lead() contract.LeadRepo {
	return newLeadRepo(t.tx)
}

//Product returns the product transaction set
func (t *mysqlTransaction) Product() contract.ProductRepo {
	return newProductRepo(t.tx)
}

//Sale returns the sale transaction set
func (t *mysqlTransaction) Sale() contract.SaleRepo {
	return newSaleRepo(t.tx)
}
