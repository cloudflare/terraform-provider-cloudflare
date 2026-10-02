package api_shield_operation

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// TestUpgradeAPIShieldOperationV500ToV501 verifies that the 500→501 state
// upgrader correctly strips the query-parameter attributes (feature,
// with_schemas) that were incorrectly added to the resource schema in v5.26.0,
// while preserving all other fields.
func TestUpgradeAPIShieldOperationV500ToV501(t *testing.T) {
	ctx := context.Background()

	r := &APIShieldOperationResource{}
	upgrader, ok := r.UpgradeState(ctx)[500]
	if !ok {
		t.Fatal("no upgrader registered for schema version 500")
	}

	// Represents a v5.26.0 state that has feature and with_schemas stored.
	// with_schemas=false (from the Default:false that was present in v5.26.0),
	// feature=null (optional, rarely set by users).
	rawState := &tfprotov6.RawState{JSON: []byte(`{
		"id":           "op-abc123",
		"operation_id": "op-abc123",
		"zone_id":      "zone-xyz",
		"feature":      null,
		"with_schemas": false,
		"endpoint":     "/api/v1/users",
		"host":         "api.example.com",
		"method":       "GET",
		"last_updated": null,
		"features":     null,
		"schemas":      null
	}`)}

	targetSchema := ResourceSchema(ctx)
	req := resource.UpgradeStateRequest{RawState: rawState}
	resp := resource.UpgradeStateResponse{State: tfsdk.State{Schema: targetSchema}}

	upgrader.StateUpgrader(ctx, req, &resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("upgrade returned diagnostics: %v", resp.Diagnostics)
	}

	// Decode the upgraded state into the v501 model.
	// If feature or with_schemas were not stripped, this Get would fail because
	// the target schema no longer declares those attributes.
	var out APIShieldOperationModel
	if diags := resp.State.Get(ctx, &out); diags.HasError() {
		t.Fatalf("decode upgraded state: %v", diags)
	}

	if got := out.ID.ValueString(); got != "op-abc123" {
		t.Errorf("id: got %q, want %q", got, "op-abc123")
	}
	if got := out.OperationID.ValueString(); got != "op-abc123" {
		t.Errorf("operation_id: got %q, want %q", got, "op-abc123")
	}
	if got := out.ZoneID.ValueString(); got != "zone-xyz" {
		t.Errorf("zone_id: got %q, want %q", got, "zone-xyz")
	}
	if got := out.Endpoint.ValueString(); got != "/api/v1/users" {
		t.Errorf("endpoint: got %q, want %q", got, "/api/v1/users")
	}
	if got := out.Host.ValueString(); got != "api.example.com" {
		t.Errorf("host: got %q, want %q", got, "api.example.com")
	}
	if got := out.Method.ValueString(); got != "GET" {
		t.Errorf("method: got %q, want %q", got, "GET")
	}
}
