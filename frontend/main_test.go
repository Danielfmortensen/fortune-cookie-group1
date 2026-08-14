package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthzHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "healthy"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestApiAll(t *testing.T) {
	// override the global variables to a guaranteed broken endpoint
	originalDns := BACKEND_DNS
	originalPort := BACKEND_PORT
	BACKEND_DNS = "127.0.0.1"
	BACKEND_PORT = "9999"

	// Defer resetting them so we don't break other potential tests
	defer func() {
		BACKEND_DNS = originalDns
		BACKEND_PORT = originalPort
	}()

	req, err := http.NewRequest("GET", "/api/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiAllHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadGateway)
	}
}

func TestApiRandom(t *testing.T) {
	// override the global variables to a guaranteed broken endpoint
	originalDns := BACKEND_DNS
	originalPort := BACKEND_PORT
	BACKEND_DNS = "127.0.0.1"
	BACKEND_PORT = "9999"

	// Defer resetting them so we don't break other potential tests
	defer func() {
		BACKEND_DNS = originalDns
		BACKEND_PORT = originalPort
	}()

	req, err := http.NewRequest("GET", "/api/random", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiRandomHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadGateway {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadGateway)
	}
}
