package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	hello_server "github.com/duanyespindola/reverse-proxy-in-go/internal/hello-server"
)

func TestHelloServerMainIntegration(t *testing.T) {
	// Arrange
	serverNumber := 3
	expectedBody := fmt.Sprintf("Hello world from server %d", serverNumber)

	// Create the handler using our internal package's logic
	handler := hello_server.HelloServerHandler(serverNumber)

	// httptest.NewServer starts a real, local HTTP server on a
	// randomly chosen available port. It's perfect for integration tests.
	testServer := httptest.NewServer(handler)
	defer testServer.Close() // Ensure the server is shut down after the test

	// Act
	// Make a real HTTP GET request to the test server's URL
	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatalf("failed to make GET request to test server: %v", err)
	}
	defer resp.Body.Close()

	// Assert
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if string(body) != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, string(body))
	}
}
