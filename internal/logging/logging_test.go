package logging

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestRedactSensitiveBodyFields(t *testing.T) {
	t.Parallel()

	body := []byte(`{"name":"server","client_secret":"secret-value","nested":{"auth_credentials":"credential-value"}}`)
	redacted := redactSensitiveBodyFields(body, "auth_credentials", "client_secret")

	var got map[string]any
	if err := json.Unmarshal(redacted, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if got["name"] != "server" {
		t.Fatalf("name = %v, want server", got["name"])
	}
	if got["client_secret"] != "[redacted]" {
		t.Fatalf("client_secret = %v, want [redacted]", got["client_secret"])
	}
	nested, ok := got["nested"].(map[string]any)
	if !ok {
		t.Fatalf("nested = %T, want map", got["nested"])
	}
	if nested["auth_credentials"] != "[redacted]" {
		t.Fatalf("auth_credentials = %v, want [redacted]", nested["auth_credentials"])
	}
}

func TestRedactSensitiveBodyFieldsPreservesBodyWithoutConfiguredFields(t *testing.T) {
	t.Parallel()

	body := []byte("{ \"client_secret\": \"secret-value\" }\n")
	if got := redactSensitiveBodyFields(body); !bytes.Equal(got, body) {
		t.Fatalf("body = %q, want unchanged %q", got, body)
	}
}

func TestRedactSensitiveBodyFieldsPreservesNonJSON(t *testing.T) {
	t.Parallel()

	body := []byte("not-json")
	if got := string(redactSensitiveBodyFields(body, "client_secret")); got != string(body) {
		t.Fatalf("redacted body = %q, want %q", got, body)
	}
}
