[![Go Report Card](https://goreportcard.com/badge/github.com/halverneus/sample)](https://goreportcard.com/report/github.com/halverneus/sample)

# Sample
A simplified Object Storage sample service written in Go.

## What it is?
Sample is exactly what it sounds like: a sample. After creating several other applications, I decided to throw together a little sample application that shows some of the architectural and coding lessons I have learned. Anyone interested in contributing is welcome to place a pull request. Please remember, this is intended to be extremely simplified so that beginners can have a starting point with easy-to-follow code. The code leaves out command-line parsing, settings validation, session management, file locking (parallel read/write on a single file WILL break), etc...

## How to get started
Steps to follow to get started developing:
* Install MariaDB 10.x
* Add the following to your /etc/mysql/my.conf file:
```
log_bin_trust_function_creators = 1
```
* In MariaDB, run:
```
CREATE SCHEMA sample DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci;
CREATE USER 'sample'@'%' IDENTIFIED BY 'sample';
GRANT ALL PRIVILEGES ON sample.* TO 'sample'@'%';
FLUSH PRIVILEGES;
```
* Pull source code for Sample:
```
go get github.com/halverneus/sample
```
* Navigate to the Sample directory:
```
cd $GOPATH/src/github.com/halverneus/sample
```
* Use the following output to create a configuration file:
```
go run sample.go help
```
* Upgrade the storage, replacing 'settings.yaml' with the path to your configuration file:
```
go run sample.go settings.yaml upgrade all
```
* Add an initial authorized user, replacing 'username' and 'password' as appropriate:
```
go run sample.go settings.yaml config user add username password
```
* Start the Sample service:
```
go run sample.go settings.yaml run
```

## Using the service
The following are a series of commands that can be executed with the 'curl' utility to upload, download and delete files. Be sure to replace 'username' and 'password' with the credentials created above. Replace '127.0.0.1:8080' with the appropriate IP address and port for configuration of your service. 'myfolder/my.file' and 'my.file' should be replaced with the remote file location or local file location, respectively.

Uploading a file:
```
curl --user username:password --upload-file my.file http://127.0.0.1:8080/api/file/myfolder/my.file
```
Downloading a file:
```
curl --user username:password http://127.0.0.1:8080/api/file/myfolder/my.file > my.file
```
Deleting a file:
```
curl --user username:password -X "DELETE" http://127.0.0.1:8080/api/file/myfolder/my.file
```

## Code layout
Quick code layout explanation:
* api -> Everything in this folder relates to the URL address. For example, api/file/get refers to a HTTP GET request to http(s)://{host}/api/file*
* cli -> Command-line interface where folder layout relates to command ordering. For example, cli/config/user.go refers to a CLI call of: 'sample config user ...'
* common/db -> Interface for database access.
* common/web -> Interface for reading from and replying to clients.
* model -> Simplified calls permitting reusable data manipulations.
* router -> Handles authenticated wrapping and routing of API calls.
* sql -> SQL string constants.
* sample.go -> Entry point for the service.
