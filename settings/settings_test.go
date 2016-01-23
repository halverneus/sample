package settings

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	testGoodFile = `
database:
  host: asdf
  port: 123
  name: qwer
  user: rewq
  password: fdsa
  max-connections: 42

storage:
  folder: /path/to/folder
`

	testBadFile = `
database:
	host: asdf
	port: 123
	name: qwer
	user: rewq
	password: fdsa
	max-connections: 42

storage:
	folder: /path/to/folder
`
)

func TestParse(t *testing.T) {
	defer os.Remove("file.yaml")

	err := ioutil.WriteFile(
		"file.yaml",
		[]byte(testGoodFile),
		0666,
	)
	if nil != err {
		t.Errorf("Failed to write good settings: %v", err)
	}

	if err = Parse("file.yaml"); nil != err {
		t.Errorf("Failed to parse good settings: %v", err)
	}

	d := &Get.Database
	if d.Host != "asdf" {
		t.Error("Database: Host did not match")
	}
	if d.Port != 123 {
		t.Error("Database: Port did not match")
	}
	if d.Name != "qwer" {
		t.Error("Database: Name did not match")
	}
	if d.User != "rewq" {
		t.Error("Database: User did not match")
	}
	if d.Password != "fdsa" {
		t.Error("Database: Password did not match")
	}
	if d.MaxConnections != 42 {
		t.Error("Database: Max Connections did not match")
	}
	if Get.Storage.Folder != "/path/to/folder" {
		t.Error("Storage: Folder did not match")
	}

	os.Remove("file.yaml")

	err = ioutil.WriteFile(
		"file.yaml",
		[]byte(testBadFile),
		0666,
	)
	if nil != err {
		t.Errorf("Failed to write bad settings: %v", err)
	}

	if err = Parse("file.yaml"); nil == err {
		t.Error("Bad settings were parsed and should have failed")
	}

	os.Remove("file.yaml")

	if err = Parse("file.yaml"); nil == err {
		t.Error("Non-existant file parsed and should have failed")
	}
}
