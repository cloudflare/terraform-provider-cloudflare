package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV4 handles state upgrades from v4 SDKv2 provider (schema_version=0) to v500.
//
// This performs a full transformation from v4 → v5 format:
//   - email_address → email (rename)
//   - role_ids → roles (rename + type conversion)
//   - policies: initialized as null (not in v4)
//   - user: initialized as null (not in v4)
//
// IMPORTANT: This must only be called when req.State was populated with the v4 PriorSchema.
// When called from UpgradeFromV0Ambiguous (where PriorSchema is nil), use
// upgradeFromV4ViaRawState instead.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_member state from v4 SDKv2 provider (schema_version=0)")

	var v4State SourceV4AccountMemberModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := TransformV4toV500(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "State upgrade from v4 to v500 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v5 stepping stone (schema_version=1) to v500.
//
// This handles the schema changes between v5.16 and current v5:
//   - policies: ListNestedAttribute → SetNestedAttribute, remove 'id' field
//   - permission_groups: ListNestedAttribute → SetNestedAttribute
//   - resource_groups: ListNestedAttribute → SetNestedAttribute
//   - roles: ListAttribute → SetAttribute
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_member state from version=1 (v5.16 stepping stone) to version=500")

	var v513State SourceV513AccountMemberModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v513State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v500State, diags := TransformV513toV500(ctx, v513State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v500State)...)
	tflog.Info(ctx, "State upgrade from v5.16 stepping stone to v500 completed successfully")
}

// UpgradeFromV0Ambiguous handles schema_version=0 state that is AMBIGUOUS between
// v4 SDKv2 format and v5.0-v5.15 production format.
//
// Both v4 and v5.0-v5.15 stored state at schema_version=0, but with incompatible formats:
//   - v4: has "email_address" and "role_ids" fields (no policies, no user)
//   - v5.0-v5.15: has "email" and "roles" fields (plus policies, user)
//
// migrations.go registers this with PriorSchema=nil so the framework skips pre-decoding
// req.State entirely. Both paths operate exclusively on req.RawState.JSON.
//
// Detection: presence of "email_address" key in raw JSON unambiguously identifies v4 state.
func UpgradeFromV0Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading account_member state from schema_version=0 (detecting v4 vs v5 format)")

	isV4, err := detectV4AccountMemberState(req)
	if err != nil {
		tflog.Warn(ctx, "Could not detect account_member state format, defaulting to v5 no-op", map[string]interface{}{
			"error": err.Error(),
		})
		// Fall through to v5 path (safer default — avoids destructive transform on v5 state)
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 account_member format (email_address present), performing transformation via raw state")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5.0-v5.15 production state: same field names as current v5 but List types
	// instead of Set types, and policies may have 'id' field.
	// Reuse the v5.13→v500 transform path since the format is compatible.
	tflog.Info(ctx, "Detected v5.0-v5.15 account_member format (email present), performing List→Set upgrade via raw state")
	upgradeFromV5EarlyViaRawState(ctx, req, resp)
}

// upgradeFromV4ViaRawState performs the v4→v500 transformation by parsing req.RawState
// directly with the v4 schema. This is necessary when the upgrader was registered with
// PriorSchema=nil (so req.State is nil and cannot be used).
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4 account_member upgrade")
		return
	}

	v4Schema := SourceV4Schema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 account_member state",
			fmt.Sprintf("Could not parse raw state as v4 format: %s", err),
		)
		return
	}

	syntheticState := &tfsdk.State{Raw: rawValue, Schema: v4Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    syntheticState,
	}

	UpgradeFromV4(ctx, syntheticReq, resp)
}

// upgradeFromV5EarlyViaRawState performs the v5.0-v5.15→v500 transformation by parsing
// req.RawState with the v5.13 stepping stone schema. This handles users who skipped
// the v5.16+ stepping stone release and jumped directly from v5.0-v5.15 to v5.19+.
//
// The v5.0-v5.15 state format has the same field names as v5.16+ but uses List types
// and may include policies.id. The SourceV513Schema and TransformV513toV500 handle
// this format correctly (List→Set conversion, policies.id removal).
func upgradeFromV5EarlyViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v5.0-v5.15 account_member upgrade")
		return
	}

	v513Schema := SourceV513Schema()
	v513Type := v513Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v513Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 account_member state",
			fmt.Sprintf("Could not parse raw state as v5.0-v5.15 format: %s", err),
		)
		return
	}

	syntheticState := &tfsdk.State{Raw: rawValue, Schema: v513Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    syntheticState,
	}

	UpgradeFromV1(ctx, syntheticReq, resp)
}

// detectV4AccountMemberState returns true if the raw state is in v4 SDKv2 format.
// v4 format has "email_address" (renamed to "email" in v5).
func detectV4AccountMemberState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}

	_, hasV4Field := raw["email_address"]
	return hasV4Field, nil
}
