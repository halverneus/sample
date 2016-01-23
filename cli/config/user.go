package config

import (
	"github.com/halverneus/sample/model/user"
)

// AddUser to permit access. Accepts (username, password).
func AddUser(args ...string) error {
	username := args[0]
	password := args[1]

	return user.Add(username, password)
}
