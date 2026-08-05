package v500_test

import (
	"context"
	"testing"

	v500 "github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config/migration/v500"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config"
)

// --- DetectV4State tests ---

func TestDetectV4State_V4State(t *testing.T) {
	t.Parallel()
	// v4 state: no created_on field
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{
			JSON: []byte(`{
				"id": "abc123",
				"account_id": "acct-456",
				"name": "my-hyperdrive",
				"origin": {
					"database": "mydb",
					"host": "db.example.com",
					"port": 5432,
					"scheme": "postgres",
					"user": "admin",
					"password": "secret"
				},
				"caching": {
					"disabled": false
				}
			}`),
		},
	}
	if !v500.DetectV4State(req) {
		t.Error("expected DetectV4State to return true for v4 state (no created_on)")
	}
}

func TestDetectV4State_V5State(t *testing.T) {
	t.Parallel()
	// v5 state: has created_on field
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{
			JSON: []byte(`{
				"id": "abc123",
				"account_id": "acct-456",
				"name": "my-hyperdrive",
				"origin": {
					"database": "mydb",
					"host": "db.example.com",
					"port": 5432,
					"scheme": "postgres",
					"user": "admin",
					"password": "secret",
					"service_id": null
				},
				"caching": {
					"disabled": false
				},
				"created_on": "2024-01-01T00:00:00Z",
				"modified_on": "2024-01-01T00:00:00Z"
			}`),
		},
	}
	if v500.DetectV4State(req) {
		t.Error("expected DetectV4State to return false for v5 state (has created_on)")
	}
}

func TestDetectV4State_NilRawState(t *testing.T) {
	t.Parallel()
	req := resource.UpgradeStateRequest{RawState: nil}
	if v500.DetectV4State(req) {
		t.Error("expected DetectV4State to return false for nil RawState")
	}
}

func TestDetectV4State_EmptyJSON(t *testing.T) {
	t.Parallel()
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: []byte{}},
	}
	if v500.DetectV4State(req) {
		t.Error("expected DetectV4State to return false for empty JSON")
	}
}

func TestDetectV4State_MalformedJSON(t *testing.T) {
	t.Parallel()
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: []byte(`{not valid json}`)},
	}
	if v500.DetectV4State(req) {
		t.Error("expected DetectV4State to return false for malformed JSON")
	}
}

func TestDetectV4State_EmptyObject(t *testing.T) {
	t.Parallel()
	// Empty JSON object has no created_on, so should be detected as v4
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: []byte(`{}`)},
	}
	if !v500.DetectV4State(req) {
		t.Error("expected DetectV4State to return true for empty object (no created_on)")
	}
}

// --- UpgradeFromV0 handler tests ---

func TestUpgradeFromV0_V4RawState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Set V5TargetSchema for the handler (normally done by parent init())
	origSchema := v500.V5TargetSchema
	v500.V5TargetSchema = hyperdrive_config.ResourceSchema
	defer func() { v500.V5TargetSchema = origSchema }()

	// Construct v4 raw state JSON (no created_on, no service_id, no mtls, no origin_connection_limit)
	v4JSON := []byte(`{
		"id": "hd-test-123",
		"account_id": "acct-456",
		"name": "my-hyperdrive",
		"origin": {
			"database": "mydb",
			"host": "db.example.com",
			"port": 5432,
			"scheme": "postgres",
			"user": "admin",
			"password": "secret",
			"access_client_id": null,
			"access_client_secret": null
		},
		"caching": {
			"disabled": false,
			"max_age": null,
			"stale_while_revalidate": null
		}
	}`)

	targetSchema := hyperdrive_config.ResourceSchema(ctx)
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: v4JSON},
	}
	resp := &resource.UpgradeStateResponse{
		State: tfsdk.State{
			Schema: targetSchema,
			Raw:    tftypes.NewValue(targetSchema.Type().TerraformType(ctx), nil),
		},
	}

	v500.UpgradeFromV0(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("UpgradeFromV0 returned errors for v4 state: %v", resp.Diagnostics.Errors())
	}

	// Verify the state was set (non-nil Raw)
	if resp.State.Raw.IsNull() {
		t.Fatal("expected resp.State.Raw to be non-null after v4 upgrade")
	}

	// Extract and verify key fields from the upgraded state
	var target v500.TargetHyperdriveConfigModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &target)...)
	if resp.Diagnostics.HasError() {
		t.Fatalf("failed to read upgraded state: %v", resp.Diagnostics.Errors())
	}

	assertStringValue(t, "id", target.ID, "hd-test-123")
	assertStringValue(t, "account_id", target.AccountID, "acct-456")
	assertStringValue(t, "name", target.Name, "my-hyperdrive")

	if target.Origin == nil {
		t.Fatal("expected origin to be non-nil")
	}
	assertStringValue(t, "origin.database", target.Origin.Database, "mydb")
	assertStringValue(t, "origin.host", target.Origin.Host, "db.example.com")
	assertNull(t, "origin.service_id", target.Origin.ServiceID)

	if target.Caching == nil {
		t.Fatal("expected caching to be non-nil")
	}
	assertBoolValue(t, "caching.disabled", target.Caching.Disabled, false)

	assertNull(t, "origin_connection_limit", target.OriginConnectionLimit)
	if target.MTLS != nil {
		t.Error("expected mtls to be nil")
	}
	assertRFC3339Null(t, "created_on", target.CreatedOn)
	assertRFC3339Null(t, "modified_on", target.ModifiedOn)
}

func TestUpgradeFromV0_V5RawState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// Set V5TargetSchema for the handler
	origSchema := v500.V5TargetSchema
	v500.V5TargetSchema = hyperdrive_config.ResourceSchema
	defer func() { v500.V5TargetSchema = origSchema }()

	// Construct v5 raw state JSON (has created_on, service_id, mtls, origin_connection_limit)
	v5JSON := []byte(`{
		"id": "hd-test-789",
		"account_id": "acct-456",
		"name": "my-v5-hyperdrive",
		"origin": {
			"database": "mydb",
			"host": "db.example.com",
			"port": 5432,
			"scheme": "postgres",
			"user": "admin",
			"password": "secret",
			"access_client_id": null,
			"access_client_secret": null,
			"service_id": null
		},
		"origin_connection_limit": null,
		"caching": {
			"disabled": false,
			"max_age": null,
			"stale_while_revalidate": null
		},
		"mtls": null,
		"created_on": "2024-06-15T10:30:00Z",
		"modified_on": "2024-06-15T10:30:00Z"
	}`)

	targetSchema := hyperdrive_config.ResourceSchema(ctx)
	req := resource.UpgradeStateRequest{
		RawState: &tfprotov6.RawState{JSON: v5JSON},
	}
	resp := &resource.UpgradeStateResponse{
		State: tfsdk.State{
			Schema: targetSchema,
			Raw:    tftypes.NewValue(targetSchema.Type().TerraformType(ctx), nil),
		},
	}

	v500.UpgradeFromV0(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("UpgradeFromV0 returned errors for v5 state: %v", resp.Diagnostics.Errors())
	}

	// Verify the state was passed through (non-null Raw)
	if resp.State.Raw.IsNull() {
		t.Fatal("expected resp.State.Raw to be non-null after v5 passthrough")
	}
}
