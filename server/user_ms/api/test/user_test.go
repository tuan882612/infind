package test

import (
	"bytes"
	"encoding/json"
	"userms/api/v1"

	// "userms/api/v1/model"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Get_Existing_User(t *testing.T) {
	query := "?username=tuan882612"
	req, _ := http.NewRequest(
		http.MethodGet, 
		BaseUrl+"/user/find"+query, 
		nil,
	)
	res := httptest.NewRecorder()
	v1.InitService().ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func Test_Get_NonExistent_User(t *testing.T) {
	query := "?username=tu"
	req, _ := http.NewRequest(
		http.MethodGet, 
		BaseUrl+"/user/find"+query, 
		nil,
	)
	res := httptest.NewRecorder()
	v1.InitService().ServeHTTP(res, req)

	assert.Equal(t, http.StatusNotFound, res.Code)
}

func Test_Invalid_User(t *testing.T) {
	data := map[string]string{
		"test": "invalid sample",
	}
	body, _ := json.Marshal(data)

	req, _  := http.NewRequest(
		http.MethodPut,
		BaseUrl+"/user/update",
		bytes.NewBuffer(body),
	)

	res := httptest.NewRecorder()
	v1.InitService().ServeHTTP(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}