package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from schema_version=0 to version=500.
//
// This handles TWO scenarios:
// 1. v4 SDKv2 state (has "key" field) - needs transformation
// 2. Early v5 state (has "key_name" field but version=0) - just needs version bump
//
// We detect which format by attempting to parse as v5 first. If key_name is present
// and valid, it's v5 state. Otherwise, we parse as v4 and transform.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading workers_kv state from schema_version=0")

	// Try to parse as v5 state first (PriorSchema is target schema)
	var v5State TargetWorkersKVModel
	diags := req.State.Get(ctx, &v5State)

	tflog.Debug(ctx, "Attempted to parse as v5 state", map[string]interface{}{
		"has_errors":        diags.HasError(),
		"key_name_is_null":  v5State.KeyName.IsNull(),
		"key_name_unknown":  v5State.KeyName.IsUnknown(),
		"key_name_value":    v5State.KeyName.ValueString(),
	})

	// Check if this is valid v5 state by seeing if key_name was populated
	if !diags.HasError() && !v5State.KeyName.IsNull() && !v5State.KeyName.IsUnknown() {
		// This is early v5 state with version=0 - just bump the version
		tflog.Info(ctx, "Detected early v5 state (has key_name), performing no-op version bump")
		resp.State.Raw = req.State.Raw
		return
	}

	// This is v4 state - we need to parse it from raw JSON since PriorSchema won't match
	tflog.Info(ctx, "Detected v4 state (missing key_name), performing transformation from raw state")

	// Parse v4 state using the v4 schema
	var v4State SourceCloudflareWorkersKVModel
	v4Diags := unmarshalV4State(ctx, req.RawState, &v4State)
	resp.Diagnostics.Append(v4Diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5Transformed, transformDiags := Transform(ctx, v4State)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5Transformed)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// unmarshalV4State unmarshals raw state into a v4 model using the v4 schema.
// This is needed when the PriorSchema is v5 but we need to parse v4 state.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceCloudflareWorkersKVModel) diag.Diagnostics {
	var diags diag.Diagnostics

	// Get the v4 schema
	sourceSchema := SourceCloudflareWorkersKVSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	// Unmarshal the raw state into a tftypes.Value
	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError(
			"Failed to unmarshal v4 state",
			"Could not parse raw state as v4 format: "+err.Error(),
		)
		return diags
	}

	// Create a tfsdk.State wrapper to use the Get method
	state := tfsdk.State{
		Raw:    rawValue,
		Schema: sourceSchema,
	}

	// Use the tfsdk.State.Get method to populate the model
	diags.Append(state.Get(ctx, target)...)
	return diags
}

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading workers_kv state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
