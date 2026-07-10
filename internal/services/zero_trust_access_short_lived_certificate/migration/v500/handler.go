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

// UpgradeFromV0 handles state upgrades from schema_version=0 to version=500.
//
// IMPORTANT: Both v4 SDKv2 provider AND v5.16.0 (dormant) have schema_version=0.
// v4 state: uses the attribute name "application_id" (Required).
// v5.16.0 state: uses the attribute name "app_id" (Required).
//
// PriorSchema is nil (see migrations.go) because these two schemas are
// incompatible for direct plugin-framework unmarshal. Detection is done by
// inspecting raw JSON.
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_short_lived_certificate state from version=0 (detecting v4 vs v5.16.0 format)")

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

// detectV4State inspects raw JSON to determine v4 vs v5.16.0 format.
// v4 uses "application_id" (Required); v5 renamed it to "app_id".
func detectV4State(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		// Fallback: no raw state — assume v4 (safer to attempt transformation).
		return true, nil
	}
	var rawJSON map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &rawJSON); err != nil {
		return false, err
	}
	if _, ok := rawJSON["app_id"]; ok {
		return false, nil // v5 format
	}
	if _, ok := rawJSON["application_id"]; ok {
		return true, nil // v4 format
	}
	// Neither key present — default to v5 pass-through so an empty/minimal
	// state doesn't fail loudly.
	return false, nil
}

// upgradeFromV4Internal performs the v4 → v5 transformation using the v4 schema.
func upgradeFromV4Internal(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var v4State SourceAccessCACertificateModel
	diags := unmarshalV4State(ctx, req.RawState, &v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, transformDiags := Transform(ctx, v4State)
	resp.Diagnostics.Append(transformDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v5 completed successfully")
}

// unmarshalV4State parses raw state using the v4 source schema.
func unmarshalV4State(ctx context.Context, rawState *tfprotov6.RawState, target *SourceAccessCACertificateModel) diag.Diagnostics {
	var diags diag.Diagnostics

	sourceSchema := SourceAccessCACertificateSchema()
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
// Used for the v5.16.0 → v5 pass-through path where no field renames are needed.
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

// UpgradeFromV1 handles state upgrades from v5 Plugin Framework provider (version=1) to v5 (version=500).
//
// This is a no-op upgrade since the schema is compatible - just bumps the version.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_access_short_lived_certificate state from version=1 to version=500 (no-op)")

	resp.State.Raw = req.State.Raw

	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}
