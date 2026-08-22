// Unit tests for the signing key list response handling, which has to select
// this resource's key because the API exposes no per-key GET endpoint.
package stream_key

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

const (
	firstKeyID  = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	secondKeyID = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
)

// listResponse is the shape GET /accounts/{account_id}/stream/keys returns: an
// array carrying only id and created, with no jwk, pem or key_id.
const listResponse = `{
  "success": true,
  "errors": [],
  "messages": [],
  "result": [
    { "id": "` + firstKeyID + `", "created": "2026-01-01T00:00:00.000000Z" },
    { "id": "` + secondKeyID + `", "created": "2026-01-01T00:01:00.000000Z" }
  ]
}`

// jwkFor builds a signing key in the format the API returns it: base64 of a JSON
// object whose kid is the key id.
func jwkFor(t *testing.T, kid string) string {
	t.Helper()
	raw, err := json.Marshal(map[string]string{"kid": kid, "kty": "RSA", "alg": "RS256"})
	if err != nil {
		t.Fatalf("failed to build jwk fixture: %v", err)
	}
	return base64.StdEncoding.EncodeToString(raw)
}

// assertCreated compares created as an instant. timetypes.RFC3339 renders the
// value in its own normalized form, so comparing the raw strings would assert on
// formatting rather than on the timestamp that was read.
func assertCreated(t *testing.T, got timetypes.RFC3339, want string) {
	t.Helper()
	if got.IsNull() {
		t.Fatal("created is null, want the value from the list response")
	}

	gotTime, diags := got.ValueRFC3339Time()
	if diags.HasError() {
		t.Fatalf("failed to read created: %v", diags)
	}
	wantTime, err := time.Parse(time.RFC3339, want)
	if err != nil {
		t.Fatalf("invalid expectation %q: %v", want, err)
	}
	if !gotTime.Equal(wantTime) {
		t.Errorf("created = %s, want %s", gotTime, wantTime)
	}
}

// stateFor builds the prior state of a resource bound to the given key id.
func stateFor(t *testing.T, id string) *StreamKeyModel {
	t.Helper()
	return &StreamKeyModel{
		ID:        types.StringValue(id),
		AccountID: types.StringValue("account-under-test"),
		Jwk:       types.StringValue(jwkFor(t, id)),
		Pem:       types.StringValue("pem-of-" + id),
	}
}

func TestRefreshStreamKeyFromList_RestoresRefreshedFields(t *testing.T) {
	data := stateFor(t, firstKeyID)
	data.Created = timetypes.NewRFC3339Null()

	found, err := refreshStreamKeyFromList(data, firstKeyID, []byte(listResponse))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found {
		t.Fatal("expected the key to be found in the list response")
	}

	if got := data.ID.ValueString(); got != firstKeyID {
		t.Errorf("id = %q, want %q", got, firstKeyID)
	}
	assertCreated(t, data.Created, "2026-01-01T00:00:00.000000Z")
}

func TestRefreshStreamKeyFromList_SelectsTheMatchingElement(t *testing.T) {
	data := stateFor(t, secondKeyID)

	found, err := refreshStreamKeyFromList(data, secondKeyID, []byte(listResponse))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found {
		t.Fatal("expected the key to be found in the list response")
	}

	if got := data.ID.ValueString(); got != secondKeyID {
		t.Errorf("id = %q, want %q", got, secondKeyID)
	}
	assertCreated(t, data.Created, "2026-01-01T00:01:00.000000Z")
}

// The list endpoint reports neither the secret material nor key_id, so a refresh
// must not blank out what state already holds for them.
func TestRefreshStreamKeyFromList_KeepsFieldsAbsentFromTheResponse(t *testing.T) {
	data := stateFor(t, firstKeyID)
	data.KeyID = types.StringValue("key-id-from-create")
	jwk, pem := data.Jwk, data.Pem

	if _, err := refreshStreamKeyFromList(data, firstKeyID, []byte(listResponse)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if data.Jwk != jwk {
		t.Errorf("jwk = %v, want it preserved from state", data.Jwk)
	}
	if data.Pem != pem {
		t.Errorf("pem = %v, want it preserved from state", data.Pem)
	}
	if got := data.KeyID.ValueString(); got != "key-id-from-create" {
		t.Errorf("key_id = %q, want it preserved from state", got)
	}
	if got := data.AccountID.ValueString(); got != "account-under-test" {
		t.Errorf("account_id = %q, want it preserved from state", got)
	}
}

func TestRefreshStreamKeyFromList_ReportsMissingKey(t *testing.T) {
	data := stateFor(t, "cccccccccccccccccccccccccccccccc")

	found, err := refreshStreamKeyFromList(data, "cccccccccccccccccccccccccccccccc", []byte(listResponse))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found {
		t.Error("expected the key to be reported as missing")
	}
}

func TestRefreshStreamKeyFromList_EmptyList(t *testing.T) {
	data := stateFor(t, firstKeyID)

	found, err := refreshStreamKeyFromList(data, firstKeyID, []byte(`{"success":true,"result":[]}`))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found {
		t.Error("expected the key to be reported as missing")
	}
}

// An unreadable response must not be mistaken for "the key is gone": the caller
// removes the resource from state when no element matches.
func TestRefreshStreamKeyFromList_UnreadableBodyIsAnError(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{"not json", `not json`},
		{"truncated json", `{"success":true,"result":[{"id":"`},
		{"no result array", `{"success":true,"errors":[]}`},
		// An explicit null is treated the same as an absent array: it says
		// nothing about which keys the account has.
		{"null result", `{"success":true,"result":null}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := stateFor(t, firstKeyID)

			found, err := refreshStreamKeyFromList(data, firstKeyID, []byte(tt.body))
			if err == nil {
				t.Error("expected an error for a response that cannot be read as the key list")
			}
			if found {
				t.Error("found = true, want false")
			}
		})
	}
}

func TestResolveStreamKeyID(t *testing.T) {
	tests := []struct {
		name string
		data *StreamKeyModel
		want string
	}{
		{
			name: "id in state",
			data: stateFor(t, firstKeyID),
			want: firstKeyID,
		},
		{
			// State written by a version that wiped id still carries the jwk.
			name: "id recovered from the retained jwk",
			data: &StreamKeyModel{ID: types.StringNull(), Jwk: types.StringValue(jwkFor(t, secondKeyID))},
			want: secondKeyID,
		},
		{
			name: "id preferred over the jwk",
			data: &StreamKeyModel{ID: types.StringValue(firstKeyID), Jwk: types.StringValue(jwkFor(t, secondKeyID))},
			want: firstKeyID,
		},
		{
			name: "nothing to correlate on",
			data: &StreamKeyModel{ID: types.StringNull(), Jwk: types.StringNull()},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveStreamKeyID(tt.data); got != tt.want {
				t.Errorf("resolveStreamKeyID() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestKeyIDFromJWK(t *testing.T) {
	payload := `{"kid":"` + firstKeyID + `","kty":"RSA"}`

	tests := []struct {
		name string
		jwk  string
		want string
	}{
		{"padded base64", base64.StdEncoding.EncodeToString([]byte(payload)), firstKeyID},
		{"unpadded base64", base64.RawStdEncoding.EncodeToString([]byte(payload)), firstKeyID},
		{"plain json", payload, firstKeyID},
		{"empty", "", ""},
		{"not base64 and not json", "@@@not-a-jwk@@@", ""},
		{"json without kid", base64.StdEncoding.EncodeToString([]byte(`{"kty":"RSA"}`)), ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyIDFromJWK(tt.jwk); got != tt.want {
				t.Errorf("keyIDFromJWK() = %q, want %q", got, tt.want)
			}
		})
	}
}
