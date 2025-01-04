package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.init()
	a.handleRoutes()
	a.run()

	m.Run()
}

func TestGetProducts(t *testing.T) {
	request, _ := http.NewRequest("GET", "/products", nil)
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Errorf("Test Failed")
	}
}
