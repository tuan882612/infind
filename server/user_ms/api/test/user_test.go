package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"userms/api/v1"
)

func Test_GetUser(t *testing.T) {
	query := "?username=tuan882612"
	req, err := http.NewRequest("GET", BaseUrl+"/user/find"+query, nil)

	if err != nil {
		t.Errorf("No handler for the endpoint: \""+BaseUrl+query+"\"")
	}

	res := httptest.NewRecorder()

	v1.InitService().ServeHTTP(res, req)

}