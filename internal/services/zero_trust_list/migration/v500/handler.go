package v500

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0Ambiguous handles schema_version=0 state that is AMBIGUOUS between
// v4 SDKv2 format (cloudflare_teams_list) and v5 production format.
//
// Both v4 and some early/beta v5 builds wrote cloudflare_zero_trust_list state at
// schema_version=0, with incompatible representations of "items":
//   - v4: "items" is a list of strings, plus a separate "items_with_description"
//     list of {value, description} objects.
//   - v5: "items" is a set of {value, description} objects; no
//     "items_with_description"; carries computed fields (list_count, created_at).
//
// migrations.go registers this with PriorSchema=nil so the framework skips
// pre-decoding req.State. Pre-decoding v5 object-format items against the v4
// string schema fails with "AttributeName(\"items\").ElementKeyInt(0):
// unsupported type json.Delim sent as tftypes.String". Both paths therefore
// operate on req.RawState.JSON.
//
// Detection: presence of "items_with_description" (or string-typed "items"
// elements) identifies v4; otherwise the state is already v5 and only needs a
// version bump.
func UpgradeFromV0Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_zero_trust_list state from schema_version=0 (detecting v4 vs v5 format)")

	isV4, err := detectV4ZeroTrustListState(req)
	if err != nil {
		tflog.Warn(ctx, "Could not detect zero_trust_list state format, defaulting to v5 no-op", map[string]interface{}{
			"error": err.Error(),
		})
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4 zero_trust_list format, performing transformation via raw state")
		upgradeFromV4ViaRawState(ctx, req, resp)
		return
	}

	// v5 format: re-decode raw JSON with the target schema. req.State is nil
	// (PriorSchema=nil), so use req.RawState directly for a no-op version bump.
	tflog.Info(ctx, "Detected v5 zero_trust_list format, performing no-op version bump via raw state")
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during zero_trust_list v0 upgrade")
		return
	}
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, unmarshalErr := req.RawState.Unmarshal(targetType)
	if unmarshalErr != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 zero_trust_list state",
			"Could not parse raw state as v5 format: "+unmarshalErr.Error(),
		)
		return
	}
	resp.State.Raw = rawValue
	tflog.Info(ctx, "zero_trust_list state version bump from 0 to 500 completed")
}

// upgradeFromV4ViaRawState parses req.RawState with the v4 schema (since
// PriorSchema is nil and req.State is therefore unavailable) and delegates to
// UpgradeFromV4.
func upgradeFromV4ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4 zero_trust_list upgrade")
		return
	}

	v4Schema := SourceTeamsListSchema()
	v4Type := v4Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v4Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4 zero_trust_list state",
			"Could not parse raw state as v4 format: "+err.Error(),
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

// detectV4ZeroTrustListState returns true if the raw schema_version=0 state is in
// the v4 cloudflare_teams_list format. v4 carries an "items_with_description"
// field and/or stores "items" as a list of strings; v5 stores "items" as a list
// of {value, description} objects and never has "items_with_description".
func detectV4ZeroTrustListState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}

	// v4-only field.
	if _, ok := raw["items_with_description"]; ok {
		return true, nil
	}

	// Fall back to inspecting the first "items" element: a JSON string => v4,
	// a JSON object => v5.
	if itemsRaw, ok := raw["items"]; ok {
		var elems []json.RawMessage
		if err := json.Unmarshal(itemsRaw, &elems); err == nil && len(elems) > 0 {
			first := bytes.TrimSpace(elems[0])
			if len(first) > 0 && first[0] == '"' {
				return true, nil
			}
		}
	}

	return false, nil
}

// UpgradeFromV4 handles v4 state at schema_version=0.
// Merges items + items_with_description into unified items set.
func UpgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_list state from v4 (schema_version=0)")

	var v4State SourceTeamsListModel
	resp.Diagnostics.Append(req.State.Get(ctx, &v4State)...)
	if resp.Diagnostics.HasError() {
		return
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "V4→V5 zero_trust_list state upgrade completed")
}

// UpgradeFromV5V1 handles v5 state at version=1 — no-op.
func UpgradeFromV5V1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading zero_trust_list state from version=1 to current (no-op)")
	resp.State.Raw = req.State.Raw
}
