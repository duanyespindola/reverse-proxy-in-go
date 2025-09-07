package hello_server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServerHandlerReturnsCorrectMessage(t *testing.T) {
	// Arrange
	serverNumber := 5 // Testing with a specific server number
	expectedBody := fmt.Sprintf("Hello world from server %d", serverNumber)

	// Create a dummy HTTP request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// Create a ResponseRecorder to capture the handler's output
	rr := httptest.NewRecorder()

	// Act
	// Call the handler function directly, treating it as an http.Handler
	handler := HelloServerHandler(serverNumber) // Our function will return an http.HandlerFunc
	handler.ServeHTTP(rr, req)

	// Assert
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatalf("could not read response body: %v", err)
	}
	if string(body) != expectedBody {
		t.Errorf("handler returned unexpected body: got %q want %q",
			string(body), expectedBody)
	}
}
