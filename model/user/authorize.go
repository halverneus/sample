package user

import (
	"fmt"

	"github.com/halverneus/sample/common/db"
	"github.com/halverneus/sample/sql/usersql"
)

// Authorize a user to access the API.
func Authorize(user, password string) (id string, err error) {
	var result *db.Result
	if result, err = db.Query(usersql.Authorize, user, user, password); nil != err {
		return
	}
	defer result.Free()

	if !result.ScanNextRow(&id) {
		err = fmt.Errorf("User %s not found", user)
	}
	return
}
