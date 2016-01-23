package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TRequest struct {
	Value string
}

func TestContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := NewContext(w, r)
		req := &TRequest{}
		if err := ctx.FromJSON(req); nil != err {
			t.Errorf("Failed to parse JSON: %v", err)
		}
		if req.Value != "My Value" {
			t.Errorf("Value does not match: %s", req.Value)
		}
		req.Value = "New Value"
		if err := ctx.Reply().Status(http.StatusCreated).With(req).Do(); nil != err {
			t.Errorf("Failed to reply with message: %v", err)
		}
	}))
	defer srv.Close()

	val := &TRequest{
		Value: "My Value",
	}
	raw, err := json.Marshal(val)
	if nil != err {
		t.Errorf("Failed to convert to JSON: %v", err)
	}
	buff := bytes.NewBuffer(raw)
	req, err := http.NewRequest("POST", srv.URL+"/hello", buff)
	if nil != err {
		t.Errorf("Failed to initialize a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	rVal := &TRequest{}
	if http.StatusCreated != resp.StatusCode {
		t.Error("Unexpected status code returned")
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(rVal); nil != err {
		t.Errorf("Failed to parse JSON message: %v", err)
	}

	if rVal.Value != "New Value" {
		t.Errorf("Unexpected value returned: %s", rVal.Value)
	}
}

func TestBadJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		ctx := NewContext(w, r)

		resp := &struct {
			Map map[int]string
		}{
			Map: make(map[int]string),
		}
		resp.Map[10] = "Kill it!"

		if err := ctx.Reply().Status(http.StatusCreated).With(resp).Do(); nil == err {
			t.Error("Expected JSON parsing failure, but none occurred")
		}
	}))
	defer srv.Close()

	raw := []byte("Test")
	buff := bytes.NewBuffer(raw)
	req, err := http.NewRequest("POST", srv.URL+"/hello", buff)
	if nil != err {
		t.Errorf("Failed to initialize a new request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if nil != err {
		t.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()
}
