// main package is the entry point into the sample application. The sample
// application provides the user the ability to store and manage versioned files
// in object storage. This is a small subset of a significantly more complex
// application suite that includes a web server, CLI client and web client
// (written in Dart).
//
// Functionality of this application includes:
//   - Authentication
//   - Uploading files as objects
//   - Uploading new revisions of files
//   - Listing available files and revisions
//   - Deletion of files and revisions
//
// NOTE: Set up MariaDB 10.x as follows:
//   CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci
//   CREATE USER 'sample'@'%' IDENTIFIED BY 'sample';
//   GRANT ALL PRIVILEGES ON sample.* TO 'sample'@'%';
//   FLUSH PRIVILEGES;
//
// In /etc/mysql/my.conf, add:
//   log_bin_trust_function_creators = 1
package main

import (
	"fmt"
	"os"

	"github.com/halverneus/sample/cli"
)

func main() {
	if err := cli.Parse(); nil != err {
		fmt.Println(err)
		os.Exit(-1)
	}
}
