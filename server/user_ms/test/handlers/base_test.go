package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"user_ms/pkg/response"
	"user_ms/test/utility"

	"github.com/stretchr/testify/assert"
)

func Test_Default_Endpoint(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, test.BaseUrl, nil)
	writer := httptest.NewRecorder()
	test.GetServer().ServeHTTP(writer, req)

	body := map[string]interface{}{
		"information": "infind user service",
		"version":     "0.2.0",
	}
	expected := response.Custom(body, http.StatusOK, "")

	var res response.Base
	json.Unmarshal(writer.Body.Bytes(), &res)

	assert.Equal(t, expected, res)
}

func Test_Health(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, test.BaseUrl+"/health", nil)
	writer := httptest.NewRecorder()
	test.GetServer().ServeHTTP(writer, req)
	
	var res map[string]string
	json.Unmarshal(writer.Body.Bytes(), &res)

	expected := map[string]string{"health":"ok"}

	assert.Equal(t, expected, res)
}

func Test_NonExistent_Endpoint(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, test.BaseUrl+"/wrong", nil)
	writer := httptest.NewRecorder()
	test.GetServer().ServeHTTP(writer, req)
	
	assert.Equal(t, http.StatusNotFound, writer.Code)
}
