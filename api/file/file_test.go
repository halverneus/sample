package file

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/halverneus/sample/cli/upgrade"
	"github.com/halverneus/sample/common/db"
	"github.com/halverneus/sample/common/web"
	"github.com/halverneus/sample/settings"
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

var (
	testFilePath = "/api/file/test/file.txt"
	testFile     = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
)

func TestAll(t *testing.T) {
	settings.Get.Storage.Folder = "testfolder/"

	// Setup for testing.
	if err := db.Connect(host, port, schema, user, password); nil != err {
		t.Errorf("Failed to connect to database with: %v", err)
	}
	if err := upgrade.All(); nil != err {
		t.Errorf("Database failed to upgrade with: %v", err)
	}

	testPut(t)
	testGet(t)
	testDelete(t)

	if err := os.RemoveAll("testfolder/"); nil != err {
		t.Errorf("Failed to delete folder: %v", err)
	}
}

////////////////////////////////////////////////////////////////////////////////
// TEST PUT
////////////////////////////////////////////////////////////////////////////////
func testPut(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := web.NewContext(w, r)
		ctx.User = "UnitTest"
		ctx.UserID = "01010101010101010101010101010101"

		Route(ctx)
	}))
	defer srv.Close()

	testPutGood(t, srv)
	testPutBad1(t, srv)
	testPutBad2(t, srv)
}

func testPutGood(t *testing.T, srv *httptest.Server) {
	buff := bytes.NewBuffer(testFile)
	req, err := http.NewRequest("PUT", srv.URL+testFilePath, buff)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusOK), resp.Status)
	}
}

func testPutBad1(t *testing.T, srv *httptest.Server) {
	buff := bytes.NewBuffer(testFile)
	req, err := http.NewRequest("PUT", srv.URL+"/api/file/", buff)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusBadRequest != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusBadRequest), resp.Status)
	}
}

func testPutBad2(t *testing.T, srv *httptest.Server) {
	settings.Get.Storage.Folder = "/"
	buff := bytes.NewBuffer(testFile)
	req, err := http.NewRequest("PUT", srv.URL+testFilePath, buff)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusInternalServerError != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusInternalServerError), resp.Status)
	}
	settings.Get.Storage.Folder = "testfolder/"
}

////////////////////////////////////////////////////////////////////////////////
// TEST GET
////////////////////////////////////////////////////////////////////////////////
func testGet(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := web.NewContext(w, r)
		ctx.User = "UnitTest"
		ctx.UserID = "01010101010101010101010101010101"

		Route(ctx)
	}))
	defer srv.Close()

	testGetGood(t, srv)
	testGetBad1(t, srv)
	testGetBad2(t, srv)
}

func testGetGood(t *testing.T, srv *httptest.Server) {
	req, err := http.NewRequest("GET", srv.URL+testFilePath, nil)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusOK), resp.Status)
	}

	contents := make([]byte, len(testFile)+5)
	if count, err := resp.Body.Read(contents); nil != err && io.EOF != err {
		t.Errorf("Received error while copying response: %v", err)
	} else if count != len(testFile) {
		t.Errorf("Expected %d bytes in reply, got %d bytes in: %s", len(testFile), count, string(contents))
	} else {
		if 0 != bytes.Compare(testFile, contents[:count]) {
			t.Errorf("Files do no match; EXPECTED: %s GOT: %s", string(testFile), string(contents))
		}
	}
}

func testGetBad1(t *testing.T, srv *httptest.Server) {
	req, err := http.NewRequest("GET", srv.URL+"/api/file/test/other.txt", nil)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusNotFound != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusNotFound), resp.Status)
	}
}

func testGetBad2(t *testing.T, srv *httptest.Server) {
	req, err := http.NewRequest("GET", srv.URL+"/api/file/", nil)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusBadRequest != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusBadRequest), resp.Status)
	}
}

////////////////////////////////////////////////////////////////////////////////
// TEST DEL
////////////////////////////////////////////////////////////////////////////////
func testDelete(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := web.NewContext(w, r)
		ctx.User = "UnitTest"
		ctx.UserID = "01010101010101010101010101010101"

		Route(ctx)
	}))
	defer srv.Close()

	testDeleteGood(t, srv)
	testDeleteBad1(t, srv)
	testDeleteBad2(t, srv)
}

func testDeleteGood(t *testing.T, srv *httptest.Server) {
	req, err := http.NewRequest("DELETE", srv.URL+testFilePath, nil)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusOK), resp.Status)
	}
}

func testDeleteBad1(t *testing.T, srv *httptest.Server) {
	req, err := http.NewRequest("DELETE", srv.URL+testFilePath, nil)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusNotFound != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusNotFound), resp.Status)
	}
}

func testDeleteBad2(t *testing.T, srv *httptest.Server) {
	req, err := http.NewRequest("DELETE", srv.URL+"/api/file/", nil)
	if nil != err {
		t.Errorf("Failed to initialized a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if http.StatusBadRequest != resp.StatusCode {
		t.Errorf("Unexpected status code; expected: %s returned: %s", http.StatusText(http.StatusBadRequest), resp.Status)
	}
}
