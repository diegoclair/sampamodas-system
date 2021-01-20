package mysql

import (
	"database/sql"

	"github.com/diegoclair/sampamodas-system/backend/domain/contract"
)

// TxManager is the MySQL transaction manager
type TxManager struct {
	tx         *sql.Tx
	committed  bool
	rolledback bool
}

func newTransaction(tx *sql.Tx) *TxManager {
	instance := &TxManager{tx: tx}
	return instance
}

// Commit persists changes to database
func (t *TxManager) Commit() error {
	err := t.tx.Commit()
	if err != nil {
		return err
	}

	t.committed = true

	return nil
}

// Rollback discards changes on database
func (t *TxManager) Rollback() error {
	if t != nil && !t.committed && !t.rolledback {
		err := t.tx.Rollback()
		if err != nil {
			return err
		}

		t.rolledback = true
	}

	return nil
}

// GetDBTransaction returns the transaction instance reference
func (t *TxManager) GetDBTransaction() *sql.Tx {
	return t.tx
}

//Business returns the company set
func (t *TxManager) Business() contract.BusinessRepo {
	return newBusinessRepo(t.tx)
}

//Company returns the company set
func (t *TxManager) Company() contract.CompanyRepo {
	return newCompanyRepo(t.tx)
}

//Lead returns the company set
func (t *TxManager) Lead() contract.LeadRepo {
	return newLeadRepo(t.tx)
}

//Product returns the company set
func (t *TxManager) Product() contract.ProductRepo {
	return newProductRepo(t.tx)
}
