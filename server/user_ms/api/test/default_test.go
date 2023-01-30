package test

import (
	"bytes"
	"userms/api/response"
	"userms/api/v1"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var BaseUrl string = "/api/v1"

func Test_default_endpoint(t *testing.T) {
	req, err := http.NewRequest("GET", BaseUrl, nil)

	if err != nil {
		t.Error("No handler for endpoint: \"/api/v1/\"")
	}

	res := httptest.NewRecorder()

	v1.InitService().ServeHTTP(res, req)

	body := map[string]string{
		"information":"infind user service",
		"version":"0.2.0",
	}
	out,_ := json.Marshal(response.Ok("", body))

	if !bytes.Equal(out, res.Body.Bytes()) {
		t.Errorf(
			"Retrieved invalid body: %v\nStatus code: %v", 
			res.Body, res.Code,
		)
	}
}