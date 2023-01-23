package test

import (
	"bytes"
	"userms/api/v1/routes"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_default_endpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/", nil)

	if err != nil {
		t.Error("No handler for endpoint: \"/api/v1/\"")
	}

	res := httptest.NewRecorder()
	routes.Init_Service().ServeHTTP(res, req)

	data := map[string]string{
		"information":"user management api of infind",
		"version":"0.0.1",
	}

	out,_ := json.Marshal(data)

	if !bytes.Equal(out, res.Body.Bytes()) {
		t.Errorf(
			"Retrieved invalid body: %v\nStatus code: %v", 
			res.Body, res.Code,
		)
	}
}