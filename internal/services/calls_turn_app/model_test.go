package calls_turn_app

import (
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// POST /accounts/{account_id}/calls/turn_keys returns the bearer token as
// `secret` (not `key`), and only on create. See issue #7412.
const createResponse = `{
  "success": true, "errors": [], "messages": [],
  "result": {
    "created": "2026-11-13T12:20:11Z",
    "modified": "2026-11-13T12:20:11Z",
    "name": "",
    "secret": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
    "uid": "9164c6bc78f7fce1cae7a3efc9dd8318"
  }
}`

// PUT and GET responses do not carry the token.
const updateResponse = `{
  "success": true, "errors": [], "messages": [],
  "result": {
    "created": "2026-11-13T12:20:11Z",
    "modified": "2026-11-14T08:00:00Z",
    "name": "renamed",
    "uid": "9164c6bc78f7fce1cae7a3efc9dd8318"
  }
}`

func TestCreateResponseStoresSecretAsKey(t *testing.T) {
	env := CallsTURNAppResultEnvelope{CallsTURNAppModel{
		AccountID: types.StringValue("acc"),
		Key:       types.StringUnknown(),
	}}
	if err := apijson.UnmarshalComputed([]byte(createResponse), &env); err != nil {
		t.Fatal(err)
	}
	got := env.Result.Key.ValueString()
	if got != "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef" {
		t.Fatalf("key = %q, want the `secret` from the create response", got)
	}
	if env.Result.UID.ValueString() != "9164c6bc78f7fce1cae7a3efc9dd8318" {
		t.Fatalf("uid = %q", env.Result.UID.ValueString())
	}
}

func TestUpdateKeepsKeyFromState(t *testing.T) {
	state := &CallsTURNAppModel{Key: types.StringValue("token-from-create")}
	env := CallsTURNAppResultEnvelope{CallsTURNAppModel{
		AccountID: types.StringValue("acc"),
		Key:       types.StringUnknown(),
	}}
	if err := apijson.UnmarshalComputed([]byte(updateResponse), &env); err != nil {
		t.Fatal(err)
	}
	data := &env.Result
	preserveCallsTURNKey(data, state)
	if data.Key.ValueString() != "token-from-create" {
		t.Fatalf("key = %q (null=%v unknown=%v), want the token kept from state",
			data.Key.ValueString(), data.Key.IsNull(), data.Key.IsUnknown())
	}
	if data.Name.ValueString() != "renamed" {
		t.Fatalf("name = %q", data.Name.ValueString())
	}
}
