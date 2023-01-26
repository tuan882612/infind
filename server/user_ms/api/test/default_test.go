package test

import (
	"bytes"
	"userms/api/v1/model"
	"userms/api/v1/routes"

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

	routes.Init_Service().ServeHTTP(res, req)

	data := map[string]string{
		"information":"infind user service",
		"version":"0.0.1",
	}

	out,_ := json.Marshal(
		model.DefaultResponse{
			Code: http.StatusOK,
			Message: "",
			Body: data,
		},
	)

	if !bytes.Equal(out, res.Body.Bytes()) {
		t.Errorf(
			"Retrieved invalid body: %v\nStatus code: %v", 
			res.Body, res.Code,
		)
	}
}