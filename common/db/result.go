package db

import (
	"database/sql"
)

// Result contains the result of a single query.
type Result struct {
	rows *sql.Rows
}

// ScanNextRow places the SQL query results into the passed destinations and
// returns true on success.
func (result *Result) ScanNextRow(destinations ...interface{}) bool {
	if nil == result.rows || !result.rows.Next() {
		return false
	}

	return nil == result.rows.Scan(destinations...)
}

// Free the result to prevent connection leaks.
func (result *Result) Free() {
	if nil != result.rows {
		result.rows.Close()
		result.rows = nil
	}
}
