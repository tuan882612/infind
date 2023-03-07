package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"user_ms/test/utility"
	
	"github.com/stretchr/testify/assert"
)

func Test_Get_Existing_User(t *testing.T) {
	query := "?username=tuan882612"
	req, _ := http.NewRequest(
		http.MethodGet, 
		test.BaseUrl+"/user/find"+query, 
		nil,
	)
	writer := httptest.NewRecorder()
	test.GetServer().ServeHTTP(writer, req)

	assert.Equal(t, http.StatusOK, writer.Code)
}

func Test_Get_NonExistent_User(t *testing.T) {
	query := "?username=tu"
	req, _ := http.NewRequest(
		http.MethodGet, 
		test.BaseUrl+"/user/find"+query, 
		nil,
	)
	writer := httptest.NewRecorder()
	test.GetServer().ServeHTTP(writer, req)

	assert.Equal(t, http.StatusNotFound, writer.Code)
}

func Test_Update_Invalid_User(t *testing.T) {
	data := map[string]string{
		"test": "invalid sample",
	}
	body, _ := json.Marshal(data)

	req, _  := http.NewRequest(
		http.MethodPut,
		test.BaseUrl+"/user/update",
		bytes.NewBuffer(body),
	)

	writer := httptest.NewRecorder()
	test.GetServer().ServeHTTP(writer, req)

	assert.Equal(t, http.StatusBadRequest, writer.Code)
}