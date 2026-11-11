package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles schema_version=0 state. v4 SDKv2 wrote version 0, and
// so did v5 releases before the zone schema version was bumped to 1, so the
// state's shape decides which path runs. v5Schema is the current v5 schema.
func UpgradeFromV0(v5Schema schema.Schema) func(context.Context, resource.UpgradeStateRequest, *resource.UpgradeStateResponse) {
	return func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
		if req.RawState == nil || len(req.RawState.JSON) == 0 {
			resp.Diagnostics.AddError("Missing raw state", "RawState was nil or empty for schema version 0 migration")
			return
		}

		if !IsV5State(req.RawState) {
			UpgradeFromV4(ctx, req, resp)
			return
		}

		tflog.Info(ctx, "Detected v5 zone state at schema_version=0, performing no-op version bump")
		rawValue, err := req.RawState.Unmarshal(v5Schema.Type().TerraformType(ctx))
		if err != nil {
			resp.Diagnostics.AddError("Failed to unmarshal v5 zone state",
				fmt.Sprintf("Could not parse raw state as v5 format: %s", err))
			return
		}
		resp.State.Raw = rawValue
	}
}

// IsV5State reports whether raw state has the v5 shape. v5 nests the account
// under "account"; v4 has a flat "account_id" and names the zone "zone".
func IsV5State(raw *tfprotov6.RawState) bool {
	if raw == nil || len(raw.JSON) == 0 {
		return false
	}
	var attrs map[string]json.RawMessage
	if err := json.Unmarshal(raw.JSON, &attrs); err != nil {
		return false
	}
	_, hasAccount := attrs["account"]
	_, hasZone := attrs["zone"]
	return hasAccount && !hasZone
}

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// PriorSchema is nil in migrations.go — v4 SDKv2 state encoding is incompatible with
// Plugin Framework schema types (e.g. plan stored as string "enterprise" but v5 schema
// expects SingleNestedAttribute object). We unmarshal raw JSON directly.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zone state from v4 SDKv2 provider (schema_version=0)")

	// Unmarshal raw state using the v4 source schema type system
	sourceSchema := SourceCloudflareZoneSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(sourceType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v4 zone state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err))
		return
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	var v4State SourceCloudflareZoneModel
	resp.Diagnostics.Append(state.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zone state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
