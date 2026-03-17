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

// UpgradeFromVersion1 handles state upgrades from schema_version=1 to version=500.
//
// IMPORTANT: Both v4 SDKv2 provider AND v5.18.0+ (dormant) have schema_version=1.
// PriorSchema is nil because v4 and v5 have incompatible schemas for configuration.
// We must detect the format at runtime by inspecting the raw state:
// - v4 state: configuration is an ARRAY (needs transformation)
// - v5.18.0+ state: configuration is an OBJECT (no-op)
//
// Detection strategy: Parse raw JSON and check if configuration is array or object.
func UpgradeFromVersion1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading access_rule state from version=1 (detecting v4 vs v5.18.0 format)")

	// Detect v4 vs v5 format by inspecting raw JSON
	isV4, err := detectV4State(req)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect state format",
			fmt.Sprintf("Could not determine v4 vs v5.18.0 state format: %s", err))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 state format (configuration is array), performing transformation")
		upgradeFromV4Internal(ctx, req, resp)
	} else {
		tflog.Info(ctx, "Detected v5.18.0+ state format (configuration is object), no-op upgrade")
		// PriorSchema is nil, so req.State is not populated.
		// Unmarshal RawState using v5 target schema.
		unmarshalV5StateToResponse(ctx, req.RawState, resp)
	}
}

// detectV4State checks if the state is v4 format by inspecting configuration field.
// v4: configuration is an array ([]interface{})
// v5: configuration is an object (map[string]interface{})
func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	// Try JSON-based detection first (most reliable)
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			if configRaw, ok := rawJSON["configuration"]; ok {
				switch configRaw.(type) {
				case []interface{}:
					// v4 format: configuration is an array
					return true, nil
				case map[string]interface{}:
					// v5 format: configuration is an object
					return false, nil
				}
			}
			// No configuration field - treat as v5 (shouldn't happen for valid state)
			return false, nil
		}
	}

	// Fallback: assume v5 if we can't detect
	return false, nil
}

// upgradeFromV4Internal performs the actual v4 → v5 transformation.
// Uses RawState to unmarshal with v4 schema since req.State uses v5 schema.
func upgradeFromV4Internal(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	// Parse v4 state using v4 schema and model
	var v4State SourceV4AccessRuleModel
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

// unmarshalV4State parses raw state using v4 schema into SourceV4AccessRuleModel.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceV4AccessRuleModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceV4AccessRuleSchema(ctx)
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

// UpgradeFromV5 handles state upgrades from v5.0.0-v5.12.x (version=0) to v5 (version=500).
//
// This is a no-op upgrade since the schema is already in v5 format.
func UpgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading access_rule state from version=0 to version=500 (no-op)")

	// CRITICAL: For no-op upgrades, copy raw state directly
	// This preserves all state data without any transformation
	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 0 to 500 completed")
}
