package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"strconv"
)

var (
	db *sql.DB
)

// Connect the database for use.
func Connect(host string, port int, schema, user, password string) (err error) {

	// sql.Open does not ever actually return an error.
	db, _ = sql.Open(
		"mysql",
		user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+schema,
	)
	if err = db.Ping(); nil != err {
		return
	}

	err = upgrade()
	return
}

// Execute is used to run a query without regard for the result.
func Execute(qry string, args ...interface{}) (err error) {
	_, err = db.Exec(qry, args...)
	return
}

// Query is used to retrieve a result from a single query execution.
func Query(qry string, args ...interface{}) (result *Result, err error) {
	var rows *sql.Rows
	rows, err = db.Query(qry, args...)
	if nil != err {
		return
	}

	result = &Result{
		rows: rows,
	}
	return
}

// IsNull returns true if the passed ID matches a NULL database ID.
func IsNull(id string) bool {
	return "00000000000000000000000000000000" == id
}
