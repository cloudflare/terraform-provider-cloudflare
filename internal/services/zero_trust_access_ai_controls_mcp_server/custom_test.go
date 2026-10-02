package zero_trust_access_ai_controls_mcp_server

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMarshalJSONForUpdateOmitsUnchangedCredentials(t *testing.T) {
	t.Parallel()

	state := ZeroTrustAccessAIControlsMcpServerModel{
		ID:              types.StringValue("manual-server"),
		AccountID:       types.StringValue("account-id"),
		AuthType:        types.StringValue("oauth"),
		Hostname:        types.StringValue("https://mcp.example.com/mcp"),
		Name:            types.StringValue("Manual server"),
		AuthCredentials: types.StringValue(`{"auth_mode":"manual"}`),
		ClientSecret:    types.StringValue("unchanged-secret"),
	}
	plan := state
	plan.Name = types.StringValue("Renamed server")

	body, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if _, ok := got["client_secret"]; ok {
		t.Fatal("unchanged client_secret was serialized")
	}
	if _, ok := got["auth_credentials"]; ok {
		t.Fatal("unchanged auth_credentials was serialized")
	}
	if got["name"] != "Renamed server" {
		t.Fatalf("name = %v, want Renamed server", got["name"])
	}
}

func TestMarshalJSONForUpdateIncludesChangedSecret(t *testing.T) {
	t.Parallel()

	state := ZeroTrustAccessAIControlsMcpServerModel{
		ClientSecret: types.StringValue("old-secret"),
	}
	plan := state
	plan.ClientSecret = types.StringValue("new-secret")

	body, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if got["client_secret"] != "new-secret" {
		t.Fatalf("client_secret = %v, want new-secret", got["client_secret"])
	}
}

func TestMarshalJSONForUpdateIncludesChangedAuthCredentialsWithoutSecret(t *testing.T) {
	t.Parallel()

	state := ZeroTrustAccessAIControlsMcpServerModel{
		AuthCredentials: types.StringValue(`{"auth_mode":"manual","registration_info":{"client_id":"old-client"}}`),
		ClientSecret:    types.StringValue("unchanged-secret"),
	}
	plan := state
	plan.AuthCredentials = types.StringValue(`{"auth_mode":"manual","registration_info":{"client_id":"new-client"}}`)

	body, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if got["auth_credentials"] != plan.AuthCredentials.ValueString() {
		t.Fatalf("auth_credentials = %v, want changed configuration", got["auth_credentials"])
	}
	if _, ok := got["client_secret"]; ok {
		t.Fatal("unchanged client_secret was serialized")
	}
}

func TestMarshalJSONForUpdateOmitsUnmanagedCredentials(t *testing.T) {
	t.Parallel()

	state := ZeroTrustAccessAIControlsMcpServerModel{
		AuthCredentials: types.StringValue(`{"auth_mode":"manual"}`),
		ClientSecret:    types.StringValue("managed-secret"),
		Description:     types.StringValue("old description"),
	}
	plan := state
	plan.AuthCredentials = types.StringNull()
	plan.ClientSecret = types.StringNull()
	plan.Description = types.StringNull()

	body, err := plan.MarshalJSONForUpdate(state)
	if err != nil {
		t.Fatalf("MarshalJSONForUpdate() error = %v", err)
	}

	var got map[string]any
	if err := json.Unmarshal(body, &got); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if _, ok := got["client_secret"]; ok {
		t.Fatal("unmanaged client_secret was serialized")
	}
	if _, ok := got["auth_credentials"]; ok {
		t.Fatal("unmanaged auth_credentials was serialized")
	}
	if description, ok := got["description"]; !ok || description != nil {
		t.Fatalf("description = %v, want explicit null", description)
	}
}
