package db

import (
	"testing"
)

// NOTE: Testing requires a schema that can be accessed with the following
// settings prior to execution:
var (
	host     = "127.0.0.1"
	port     = 3306
	schema   = "sample"
	user     = "sample"
	password = "sample"
)

func TestConnect(t *testing.T) {
	if err := Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Failed to connect to database with: %v", err)
	}
	if err := Connect(host, port, schema, user, "Bad Password"); nil == err {
		t.Error("Database connected when connection should have failed")
	}
}

func TestUpgrade(t *testing.T) {
	if err := Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Failed to connect to database with: %v", err)
	}

	// Assign bad version.
	if err := Execute("UPDATE dbversion SET dbversion=2000 WHERE dbkey='db'"); nil != err {
		t.Errorf("Unable to set preconditions on bad version: %v", err)
	}
	if err := Connect(host, port, schema, user, password); nil == err {
		t.Errorf("Connection succeeded with bad version")
	}

	// Remove version tracker.
	if err := Execute("DROP TABLE dbversion"); nil != err {
		t.Log("Unable to delete dbversion table")
	}

	// Reconnect to schema.
	if err := Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Failed to connect to database with: %v", err)
	}

	// Break upgrade. TODO: Not possible with current upgrade path.
	// if err := Connect(host, port, schema, user, password); nil != err {
	// 	t.Errorf("Failed to connect to database with: %v", err)
	// }
}

func TestQuery(t *testing.T) {
	result, err := Query("SELECT * FROM dbversion")
	if nil != err {
		t.Errorf("Error while running query without args: %v", err)
	}
	result.Free()

	result, err = Query("SELECT dbversion FROM dbversion WHERE dbkey='ace'")
	if nil != err {
		t.Errorf("Error while running basic query: %v", err)
	}
	var version int
	if result.ScanNextRow(&version) {
		t.Error("Value returned unexpectedly")
	}
}

func TestIsNull(t *testing.T) {
	if !IsNull("00000000000000000000000000000000") {
		t.Error("NULL value not marked as null")
	}
	if IsNull("00000000000000000000000000000001") {
		t.Error("Non-NULL value markeds as null")
	}
}
