package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// V5TargetSchema is set by parent package init() to provide the v5 target schema.
// This avoids circular imports between the migration package and parent package.
var V5TargetSchema func(context.Context) schema.Schema

// UpgradeFromV0 handles state upgrades from schema_version=0 to version=500.
//
// IMPORTANT: Both v4 framework provider AND early v5 provider have schema_version=0.
// PriorSchema is nil because v4 and v5 have incompatible schemas for caching
// (v4 uses types.Object, v5 uses pointer struct with additional fields).
// We detect the format at runtime by inspecting the raw state:
//   - v4 state: no created_on field (needs transformation)
//   - v5 state: has created_on field (no-op passthrough)
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading hyperdrive_config state from schema_version=0 (detecting v4 vs v5 format)")

	if DetectV4State(req) {
		tflog.Info(ctx, "Detected v4 hyperdrive_config format (no created_on), performing transformation")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5 state: already in correct format, just unmarshal and pass through
	tflog.Info(ctx, "Detected v5 hyperdrive_config format (created_on present), no-op upgrade")
	unmarshalV5StateToResponse(ctx, req, resp)
}

// DetectV4State checks if the state is v4 format by inspecting raw JSON.
// v4 state does not have created_on (added in v5).
// v5 state has created_on as a computed field.
// Returns false (assume v5) if detection fails -- safer than applying a
// destructive transform to state that's already in v5 format.
func DetectV4State(req resource.UpgradeStateRequest) bool {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false
	}
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err != nil {
		return false
	}
	// created_on is a v5-only computed field, never present in v4 state
	_, hasCreatedOn := rawJSON["created_on"]
	return !hasCreatedOn
}

// upgradeFromV4ViaRawState performs the v4 -> v500 transformation by parsing
// req.RawState with the v4 schema. This is necessary when PriorSchema is nil.
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4 hyperdrive_config upgrade")
		return
	}

	// Parse raw state using v4 source schema
	v4Schema := SourceCloudflareHyperdriveConfigSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 hyperdrive_config state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err),
		)
		return
	}

	// Build a synthetic State from the raw value so Transform can read it
	syntheticState := tfsdk.State{Raw: rawValue, Schema: v4Schema}
	var sourceState SourceCloudflareHyperdriveConfigModel
	resp.Diagnostics.Append(syntheticState.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Transform to target
	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the upgraded state
	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade from v4 hyperdrive_config completed successfully")
}

// unmarshalV5StateToResponse unmarshals v5 raw state using the target schema
// and sets it on the response. Used when PriorSchema is nil (so req.State is
// not populated) and the state is already in v5 format.
func unmarshalV5StateToResponse(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if V5TargetSchema == nil {
		resp.Diagnostics.AddError("Migration configuration error",
			"V5TargetSchema not set. Ensure parent package init() sets v500.V5TargetSchema.")
		return
	}
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v5 hyperdrive_config upgrade")
		return
	}

	targetSchema := V5TargetSchema(ctx)
	targetType := targetSchema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(targetType)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 hyperdrive_config state",
			fmt.Sprintf("Could not parse raw state as v5 format: %s", err),
		)
		return
	}

	resp.State.Raw = rawValue
}
