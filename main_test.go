package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if response["message"] != "Welcome to Demo App" {
		t.Errorf("unexpected message: got %v", response["message"])
	}
}

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response HealthResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if response.Status != "healthy" {
		t.Errorf("unexpected status: got %v want healthy", response.Status)
	}

	if response.Version != version {
		t.Errorf("unexpected version: got %v want %v", response.Version, version)
	}
}

func TestReadyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(readyHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if response["status"] != "ready" {
		t.Errorf("unexpected status: got %v want ready", response["status"])
	}
}

func TestInfoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(infoHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response InfoResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response: %v", err)
	}

	if response.Service != "demo-app" {
		t.Errorf("unexpected service: got %v want demo-app", response.Service)
	}

	if response.Version != version {
		t.Errorf("unexpected version: got %v want %v", response.Version, version)
	}
}

func BenchmarkHealthHandler(b *testing.B) {
	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)

	for i := 0; i < b.N; i++ {
		handler.ServeHTTP(rr, req)
	}
}
