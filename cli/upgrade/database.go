package upgrade

import (
	"github.com/halverneus/sample/cli/upgrade/users"
)

// Database upgrade function. Inputs ignored.
func Database(...string) (err error) {
	if err = users.Upgrade(); nil != err {
		return
	}
	return
}
