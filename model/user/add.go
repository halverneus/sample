package user

import (
	"fmt"

	"github.com/halverneus/sample/common/db"
	"github.com/halverneus/sample/sql/usersql"
)

// Add a new user.
func Add(user, password string) (err error) {
	var result *db.Result
	if result, err = db.Query(usersql.Add, user, password); nil != err {
		return
	}
	defer result.Free()

	var id string
	if !result.ScanNextRow(&id) || db.IsNull(id) {
		err = fmt.Errorf("Failed to insert user %s", user)
	}
	return
}
