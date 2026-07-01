package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGatewayAwareHandlerStripsConfiguredPrefix(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/apps", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/apps" {
			t.Fatalf("path was not stripped: %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	handler := gatewayAwareHandler(mux, "/app/fnos-apps-store")
	req := httptest.NewRequest(http.MethodGet, "/app/fnos-apps-store/api/apps", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestGatewayAwareHandlerMapsExactPrefixToRoot(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			t.Fatalf("path was not mapped to root: %q", r.URL.Path)
		}
		w.WriteHeader(http.StatusNoContent)
	})

	handler := gatewayAwareHandler(mux, "/app/fnos-apps-store")
	req := httptest.NewRequest(http.MethodGet, "/app/fnos-apps-store", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d", rec.Code)
	}
}

func TestGatewayAwareHandlerKeepsDirectPortPath(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/apps", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	handler := gatewayAwareHandler(mux, "/app/fnos-apps-store")
	req := httptest.NewRequest(http.MethodGet, "/api/apps", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status = %d", rec.Code)
	}
}
