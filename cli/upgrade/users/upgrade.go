package users

import (
	"github.com/halverneus/sample/common/db"
	"github.com/halverneus/sample/sql/usersql"
)

const (
	// dbSpace is the database namespace identifier for tracking updates.
	dbSpace = "user"
)

var (
	// steps is an order-specific set of upgrade steps that need to occur to make
	// the database current.
	// currentVersion is used to track if all upgrade steps have been completed.
	steps          = []func() error{s0, s1, s2, s3, s4}
	currentVersion = len(steps)
)

// Upgrade the user-specific portions of the database.
func Upgrade() (err error) {
	version := db.GetVersion(dbSpace)

	// Run all upgrade steps until the version of the database is current.
	for currentVersion != version {
		if version, err = record(version, steps[version]()); nil != err {
			return
		}
	}

	return
}

// record the incremented version if the change was successfully applied.
func record(version int, err error) (int, error) {
	if nil != err {
		return version, err
	}
	version++
	return version, db.SetVersion(dbSpace, version)
}

// s0 creates the initial user table.
func s0() error {
	return db.Execute(usersql.CreateTable)
}

// s1 adds the password encryption function.
func s1() error {
	return db.Execute(usersql.CreateEncryptionFunction)
}

// s2 adds the ID generator function.
func s2() error {
	return db.Execute(usersql.CreateIDFunction)
}

// s3 adds the function used to generate 'null' IDs.
func s3() error {
	return db.Execute(usersql.CreateNullFunction)
}

// s4 adds the function used to add a new user.
func s4() error {
	return db.Execute(usersql.AddFunction)
}
