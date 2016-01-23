package cli

import (
	"errors"
	"fmt"

	"github.com/halverneus/sample/cli/config"
	"github.com/halverneus/sample/cli/help"
	"github.com/halverneus/sample/cli/run"
	"github.com/halverneus/sample/cli/upgrade"
	"github.com/halverneus/sample/common/db"
	"github.com/halverneus/sample/settings"
)

// Parse command-line arguments and route to appropriate functions.
func Parse() (err error) {
	args := Args{}.Parse()

	switch {

	// HELP DOCUMENTATION
	case args.Matches("help"):
		fmt.Print(help.Overall)
	case args.Matches("config", "help"):
		fmt.Print(help.ConfigNoun)
	case args.Matches("config", "user", "help"):
		fmt.Print(help.ConfigUser)

	// OPERATIONS
	case args.Matches("*", "config", "user", "add", "*", "*"):
		err = WithSettings(args[0], config.AddUser, args[4], args[5])
	case args.Matches("*", "run"):
		err = WithSettings(args[0], run.Service)
	case args.Matches("*", "upgrade", "all"):
		err = WithSettings(args[0], upgrade.All)
	case args.Matches("*", "upgrade", "database"):
		err = WithSettings(args[0], upgrade.Database)

	// UNKNOWN COMMAND
	default:
		err = errors.New("Unknown arguments provided; try 'sample help'")
	}

	return
}

// WithSettings run the requested operation.
func WithSettings(
	settingsFile string,
	operation func(...string) error,
	args ...string,
) error {

	// Parse settings.
	if err := settings.Parse(settingsFile); nil != err {
		return err
	}

	// Connect to database.
	if err := db.Connect(
		settings.Get.Database.Host,
		settings.Get.Database.Port,
		settings.Get.Database.Name,
		settings.Get.Database.User,
		settings.Get.Database.Password,
	); nil != err {
		return err
	}

	return operation(args...)
}
