package test

import (
	"encoding/json"
	"userms/api/response"
	"userms/api/v1/server"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var BaseUrl string = "/api/v1"

func Test_Default_Endpoint(t *testing.T) {
	req, _ := http.NewRequest(
		http.MethodGet, 
		BaseUrl, 
		nil,
	)

	res := httptest.NewRecorder()
	v1.InitService().ServeHTTP(res, req)

	body := map[string]interface{}{
		"information": "infind user service",
		"version":     "0.2.0",
	}
	expected := response.Custom(body, http.StatusOK, "")

	data := response.Base{}
	json.Unmarshal(res.Body.Bytes(), &data)

	assert.Equal(t, expected, data)
}

func Test_NonExistent_Endpoint(t *testing.T) {
	req, _ := http.NewRequest(
		http.MethodGet, 
		BaseUrl+"wrong", 
		nil,
	)

	res := httptest.NewRecorder()
	v1.InitService().ServeHTTP(res, req)
	
	assert.Equal(t, http.StatusNotFound, res.Code)
}