package domain

import "regexp"

//NoSQLRowsRE - to check if sql error is because that there are no rows
var NoSQLRowsRE = regexp.MustCompile(noSQLRows)

//NoRecordsFindRE - to check if sql error is because that there are no records find with the parameters
var NoRecordsFindRE = regexp.MustCompile(noRecordsFind)

const noSQLRows string = "no rows in result set"
const noRecordsFind string = "No records find"
