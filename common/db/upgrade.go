package db

import (
	"fmt"
)

const (
	dbSpace = "db"
)

var (
	// steps is an order-specific set of upgrade steps that need to occur to make
	// the database current.
	// currentVersion is used to track if all upgrade steps have been completed.
	steps          = []func() error{s0}
	currentVersion = len(steps)
)

// upgrade the database.
func upgrade() (err error) {
	version := GetVersion(dbSpace)

	if version > currentVersion {
		return fmt.Errorf("Unknown schema version for db: %d", version)
	}

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
	return version, SetVersion(dbSpace, version)
}

// s0 creates the initial schema version table.
func s0() error {
	return Execute(sqlVersionCreate)
}
