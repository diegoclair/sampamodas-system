package errors

import (
	"github.com/diegoclair/sampamodas-system/backend/domain"
)

//SQLResultIsEmpty - Check if the error is because there are no sql rows or
//no records find with given parameters
func SQLResultIsEmpty(err string) bool {
	noRowsIdx := domain.NoSQLRowsRE.FindStringIndex(err)
	if len(noRowsIdx) > 0 {
		return true
	}

	noRecorsIdx := domain.NoRecordsFindRE.FindStringIndex(err)
	if len(noRecorsIdx) > 0 {
		return true
	}
	return false
}
