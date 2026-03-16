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
// PriorSchema is nil because v4 and v5 have incompatible schemas:
// - v4 state: context_awareness is an ARRAY (ListNestedAttribute)
// - v5.16.0 state: context_awareness is an OBJECT (SingleNestedAttribute)
//
// Detection strategy: Parse raw JSON and check if context_awareness is array (v4) or object (v5).
func UpgradeFromVersion0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_dlp_predefined_profile state from version=0 (detecting v4 vs v5.16.0 format)")

	// Detect v4 vs v5 format by inspecting raw JSON
	isV4, err := detectV4State(req)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect state format",
			fmt.Sprintf("Could not determine v4 vs v5.16.0 state format: %s", err))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 state format, performing transformation")
		upgradeFromV4Internal(ctx, req, resp)
	} else {
		tflog.Info(ctx, "Detected v5.16.0+ state format, no-op upgrade")
		unmarshalV5StateToResponse(ctx, req.RawState, resp)
	}
}

// detectV4State checks if the state is v4 format.
// v4: has "type" field (required in v4, removed in v5)
// v5: has "profile_id" field (required in v5, not in v4)
//
// Detection is deterministic: v4 always has "type", v5 always has "profile_id".
func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			// Check for v4-only field: "type" (required in v4, removed in v5)
			if _, ok := rawJSON["type"]; ok {
				return true, nil // v4 format - has v4-only "type" field
			}

			// Check for v5-only fields: "profile_id" (required), "entries" (computed)
			if _, ok := rawJSON["profile_id"]; ok {
				return false, nil // v5 format - has v5-only "profile_id" field
			}
			if _, ok := rawJSON["entries"]; ok {
				return false, nil // v5 format - has v5-only "entries" field
			}

			// Fallback: Check if context_awareness is an array (v4) or object (v5)
			if contextAwareness, ok := rawJSON["context_awareness"]; ok && contextAwareness != nil {
				if _, isArray := contextAwareness.([]interface{}); isArray {
					return true, nil // v4 format - context_awareness is array
				}
				if _, isMap := contextAwareness.(map[string]interface{}); isMap {
					return false, nil // v5 format - context_awareness is object
				}
			}
			// Should not reach here - "type" (v4) or "profile_id" (v5) should always exist
			return false, fmt.Errorf("could not determine state version: neither 'type' (v4) nor 'profile_id' (v5) found")
		}
	}
	return false, fmt.Errorf("no raw state JSON available")
}

// upgradeFromV4Internal performs the actual v4 → v5 transformation.
func upgradeFromV4Internal(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_dlp_predefined_profile state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 source model
	var v4State SourceCloudflareDLPProfileModel
	diags := unmarshalV4State(ctx, req.RawState, &v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Failed to parse v4 state for zero_trust_dlp_predefined_profile",
			map[string]interface{}{
				"diagnostics": resp.Diagnostics.Errors(),
			})
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
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully for zero_trust_dlp_predefined_profile")
}

// unmarshalV4State parses raw state using v4 schema.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceCloudflareDLPProfileModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceCloudflareDLPProfileSchema()
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

// UpgradeFromV0 is an alias for UpgradeFromVersion0 for backward compatibility.
// Deprecated: Use UpgradeFromVersion0 instead.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	UpgradeFromVersion0(ctx, req, resp)
}

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
// This handler is only triggered when TF_MIG_TEST=1 (GetSchemaVersion returns 500).
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_dlp_predefined_profile state from version=1 to version=500 (no-op)")

	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed for zero_trust_dlp_predefined_profile")
}
