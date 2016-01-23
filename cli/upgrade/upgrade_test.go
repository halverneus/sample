package upgrade

import (
	"testing"

	"github.com/halverneus/sample/common/db"
	"github.com/halverneus/sample/sql/usersql"
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

func TestDatabaseUpgrade(t *testing.T) {
	// Happy path...
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Unable to connect to database: %v", err)
	}
	db.Execute("DROP SCHEMA sample")
	if err := db.Execute("CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci"); nil != err {
		t.Errorf("Unable to create database: %v", err)
	}
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Unable to reconnect to database: %v", err)
	}

	if err := Database(); nil != err {
		t.Errorf("Database upgrade failed: %v", err)
	}

	// Error path...
	db.Execute("DROP SCHEMA sample")
	if err := db.Execute("CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci"); nil != err {
		t.Errorf("Unable to create database: %v", err)
	}
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Unable to reconnect to database: %v", err)
	}

	if err := db.Execute(usersql.CreateEncryptionFunction); nil != err {
		t.Errorf("Unable to add pre-existing function: %v", err)
	}

	if err := Database(); nil == err {
		t.Error("Database upgrade succeeded when it should have failed")
	}

	// Restore database...
	db.Execute("DROP SCHEMA sample")
	if err := db.Execute("CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci"); nil != err {
		t.Errorf("Unable to create database: %v", err)
	}
}

func TestAllUpgrade(t *testing.T) {
	// Happy path...
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Unable to connect to database: %v", err)
	}
	db.Execute("DROP SCHEMA sample")
	if err := db.Execute("CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci"); nil != err {
		t.Errorf("Unable to create database: %v", err)
	}
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Unable to reconnect to database: %v", err)
	}

	if err := All(); nil != err {
		t.Errorf("All upgrade failed: %v", err)
	}

	// Error path...
	db.Execute("DROP SCHEMA sample")
	if err := db.Execute("CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci"); nil != err {
		t.Errorf("Unable to create database: %v", err)
	}
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Unable to reconnect to database: %v", err)
	}

	if err := db.Execute(usersql.CreateEncryptionFunction); nil != err {
		t.Errorf("Unable to add pre-existing function: %v", err)
	}

	if err := All(); nil == err {
		t.Error("All upgrade succeeded when it should have failed")
	}

	// Restore database...
	db.Execute("DROP SCHEMA sample")
	if err := db.Execute("CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci"); nil != err {
		t.Errorf("Unable to create database: %v", err)
	}
}
