package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// V5TargetSchema is set by parent package init() to provide the target schema.
// This avoids circular imports between the migration package and parent package.
var V5TargetSchema func(context.Context) schema.Schema

// UpgradeFromVersion0 handles state upgrades from schema_version=0 to version=500.
//
// IMPORTANT: Both v4 SDKv2 provider AND v5.16.0 (dormant) have schema_version=0.
// PriorSchema is nil because v4 and v5 have incompatible schemas for data:
// - v4 state: data is an ARRAY (ListNestedBlock, needs transformation)
// - v5.16.0 state: data is an OBJECT (SingleNestedAttribute, no-op)
//
// Detection strategy: Parse raw JSON and check if data is array or object.
func UpgradeFromVersion0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_dex_test state from version=0 (detecting v4 vs v5.16.0 format)")

	// Detect v4 vs v5 format by inspecting raw JSON
	isV4, err := detectV4State(req)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect state format",
			fmt.Sprintf("Could not determine v4 vs v5.16.0 state format: %s", err))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 state format (data is array), performing transformation")
		upgradeFromV4Internal(ctx, req, resp)
	} else {
		tflog.Info(ctx, "Detected v5.16.0+ state format (data is object), no-op upgrade")
		// PriorSchema is nil, so req.State is not populated.
		// Unmarshal RawState using v5 target schema.
		unmarshalV5StateToResponse(ctx, req.RawState, resp)
	}
}

// detectV4State checks if the state is v4 format by inspecting data field.
// v4: data is an array ([]interface{})
// v5: data is an object (map[string]interface{})
func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	// Try JSON-based detection first (most reliable)
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			if dataRaw, ok := rawJSON["data"]; ok {
				switch dataRaw.(type) {
				case []interface{}:
					// v4 format: data is an array
					return true, nil
				case map[string]interface{}:
					// v5 format: data is an object
					return false, nil
				}
			}
			// No data field - treat as v5 (shouldn't happen for valid state)
			return false, nil
		}
	}

	// Fallback: assume v5 if we can't detect
	return false, nil
}

// upgradeFromV4Internal performs the actual v4 → v5 transformation.
// Uses RawState to unmarshal with v4 schema since req.State is not populated.
func upgradeFromV4Internal(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_dex_test state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 schema and model
	var v4State SourceCloudflareDeviceDexTestModel
	diags := unmarshalV4State(ctx, req.RawState, &v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform v4 → v5
	v5State, transformDiags := Transform(ctx, v4State)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Write transformed state
	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// unmarshalV4State parses raw state using v4 schema into SourceCloudflareDeviceDexTestModel.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceCloudflareDeviceDexTestModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceCloudflareDeviceDexTestSchema()
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError("Failed to unmarshal v4 state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err))
		return diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: sourceSchema}
	diags.Append(state.Get(ctx, target)...)
	return diags
}

// unmarshalV5StateToResponse unmarshals v5 raw state and sets it on the response.
// Used when PriorSchema is nil and we need to copy v5 state through.
func unmarshalV5StateToResponse(ctx context.Context, rawState *tfprotov6.RawState, resp *resource.UpgradeStateResponse) {
	if V5TargetSchema == nil {
		resp.Diagnostics.AddError("Migration configuration error",
			"V5TargetSchema not set. Ensure parent package init() sets v500.V5TargetSchema.")
		return
	}

	targetSchema := V5TargetSchema(ctx)
	targetType := targetSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(targetType)
	if err != nil {
		resp.Diagnostics.AddError("Failed to unmarshal v5 state",
			fmt.Sprintf("Could not parse raw state as v5 format: %s", err))
		return
	}

	resp.State.Raw = rawValue
}

// UpgradeFromV5 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500).
//
// The version=1 state is the "dormant" v5 state (before migrations are activated).
// When TF_MIG_TEST=1, the schema version becomes 500, triggering this upgrade to bump the version.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_dex_test state from version=1 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// UpgradeFromV4 is an alias for UpgradeFromVersion0 for backward compatibility.
// Deprecated: Use UpgradeFromVersion0 instead.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	UpgradeFromVersion0(ctx, req, resp)
}
