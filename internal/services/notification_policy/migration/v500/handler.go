// File generated for v4 to v5 state migration

package v500

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v5 (version=500).
//
// PriorSchema is nil in migrations.go — v4 SDKv2 state encoding is incompatible with
// Plugin Framework schema types (ListNestedBlock → ListNestedAttribute encoding mismatch).
// We unmarshal raw JSON directly using the source schema type.
//
// Major transformations:
// - filters: MaxItems:1 list → SingleNestedAttribute object
// - Three integration Sets → single mechanisms nested object
// - Integration items: drop "name" field, keep only "id"
// - Filter fields: Set → List conversion (~35 fields)
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading notification_policy state from v4 SDKv2 provider (schema_version=0)")

	// Unmarshal raw state using the v4 source schema type system
	sourceSchema := SourceCloudflareNotificationPolicySchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(sourceType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v4 notification_policy state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err))
		return
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	var v4State SourceCloudflareNotificationPolicyModel
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
// This handler is only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500).
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading notification_policy state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
