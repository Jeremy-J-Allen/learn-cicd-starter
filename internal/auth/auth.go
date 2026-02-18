package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected error when Authorization header is missing")
	}
}

func TestGetAPIKey_MalformedHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer abc123")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Fatal("expected error for malformed authorization header")
	}
}

func TestGetAPIKey_ValidHeader(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if key != "my-secret-key" {
		t.Fatalf("expected key 'my-secret-key', got '%s'", key)
	}
}
