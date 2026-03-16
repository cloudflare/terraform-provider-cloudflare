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

// UpgradeNoOp handles state upgrades within the v5 series (schema_version=1+).
// This is a no-op upgrade since the schema is compatible - just copy state through.
func UpgradeNoOp(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading pages_project state (no-op)")
	// No-op upgrade: schema is compatible, just copy raw state through
	resp.State.Raw = req.State.Raw
}

// UpgradeFromVersion0 handles state upgrades from schema_version=0 to version=500.
//
// IMPORTANT: Both v4 SDKv2 provider AND v5.16.0 (dormant) have schema_version=0.
// PriorSchema is nil because v4 and v5 have incompatible schemas:
// - v4 state: build_config, source are ARRAYS (ListNestedAttribute)
// - v5.16.0 state: these fields are OBJECTS (SingleNestedAttribute)
//
// Detection strategy: Parse raw JSON and check if build_config is array (v4) or object (v5).
func UpgradeFromVersion0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading pages_project state from version=0 (detecting v4 vs v5.16.0 format)")

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
// v4: has build_config or source as array [], has "secrets" field (v4-only)
// v5: has build_config or source as object {}, has "uses_functions" (computed, v5-only)
//
// Detection is deterministic: v5 always has "uses_functions" key in state (computed field).
func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState != nil && len(req.RawState.JSON) > 0 {
		var rawJSON map[string]interface{}
		if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err == nil {
			// Check for v5-only computed fields (key exists even if value is null)
			// "uses_functions", "framework", "latest_deployment" are v5-only
			v5OnlyFields := []string{"uses_functions", "framework", "latest_deployment", "canonical_deployment"}
			for _, field := range v5OnlyFields {
				if _, ok := rawJSON[field]; ok {
					return false, nil // v5 format - has v5-only field key
				}
			}

			// Check for v4-only field: "secrets" (removed in v5)
			if _, ok := rawJSON["secrets"]; ok {
				return true, nil // v4 format - has v4-only "secrets" field
			}

			// Check if build_config is an array (v4) or object (v5)
			if buildConfig, ok := rawJSON["build_config"]; ok && buildConfig != nil {
				if _, isArray := buildConfig.([]interface{}); isArray {
					return true, nil // v4 format - build_config is array
				}
				if _, isMap := buildConfig.(map[string]interface{}); isMap {
					return false, nil // v5 format - build_config is object
				}
			}
			// Check source as secondary indicator
			if source, ok := rawJSON["source"]; ok && source != nil {
				if _, isArray := source.([]interface{}); isArray {
					return true, nil // v4 format - source is array
				}
				if _, isMap := source.(map[string]interface{}); isMap {
					return false, nil // v5 format - source is object
				}
			}
			// Should not reach here - one of the above checks should match
			return false, fmt.Errorf("could not determine state version: no distinguishing fields found")
		}
	}
	return false, fmt.Errorf("no raw state JSON available")
}

// upgradeFromV4Internal performs the actual v4 → v5 transformation.
func upgradeFromV4Internal(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading pages_project state from v4 SDKv2 provider (schema_version=0)")

	// Parse v4 state using v4 source model
	var v4State SourcePagesProjectModelV0
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

// unmarshalV4State parses raw state using v4 schema.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourcePagesProjectModelV0) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourcePagesProjectSchemaV0(ctx)
	sourceType := sourceSchema.Type().TerraformType(ctx)

	rawValue, err := rawState.Unmarshal(sourceType)
	if err != nil {
		diags.AddError("Failed to unmarshal v4 state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err))
		return diags
	}

	state := tfsdk.State{Raw: rawValue, Schema: *sourceSchema}
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
