package v500

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles the version 0 state which could be V4 or V5.
//
// V4 state has "script_name", V5 state has "script". The union PriorSchema
// includes both fields. Detection via raw tftypes map determines which path.
//
// Paths:
//   - V4 state: script_name present → rename to script
//   - V5 state: script present → strip V4-only fields, pass through
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	isV4, err := detectV4State(req.State.Raw)
	if err != nil {
		resp.Diagnostics.AddError("Failed to detect state format", fmt.Sprintf("Could not determine V4 vs V5 state: %s", err))
		return
	}

	if isV4 {
		tflog.Info(ctx, "Detected V4 workers_route state (script_name present), performing transform")
		upgradeFromV4(ctx, req, resp)
	} else {
		tflog.Info(ctx, "Detected V5 workers_route state, passing through")
		upgradeFromV5(ctx, req, resp)
	}
}

// detectV4State checks if "script_name" attribute is present and non-null.
func detectV4State(raw tftypes.Value) (bool, error) {
	var rawState map[string]tftypes.Value
	if err := raw.As(&rawState); err != nil {
		return false, fmt.Errorf("failed to read raw state as object: %w", err)
	}

	scriptNameVal, has := rawState["script_name"]
	if has && scriptNameVal.IsKnown() && !scriptNameVal.IsNull() {
		return true, nil
	}

	return false, nil
}

// upgradeFromV4 extracts V4 fields from raw tftypes map and transforms to V5.
// Can't use req.State.Get() because the union schema has "script" which SourceWorkerRouteModel lacks.
func upgradeFromV4(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var rawState map[string]tftypes.Value
	if err := req.State.Raw.As(&rawState); err != nil {
		resp.Diagnostics.AddError("Failed to read V4 state", fmt.Sprintf("Could not read raw state as map: %s", err))
		return
	}

	v4State := SourceWorkerRouteModel{
		ID:         extractString(rawState, "id"),
		ZoneID:     extractString(rawState, "zone_id"),
		Pattern:    extractString(rawState, "pattern"),
		ScriptName: extractString(rawState, "script_name"),
	}

	v5State, diags := Transform(ctx, v4State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, v5State)...)
	tflog.Info(ctx, "V4→V5 workers_route state upgrade completed")
}

// upgradeFromV5 strips V4-only fields (script_name) from the union-typed raw state.
func upgradeFromV5(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	var rawMap map[string]tftypes.Value
	if err := req.State.Raw.As(&rawMap); err != nil {
		resp.Diagnostics.AddError("Failed to read V5 state", fmt.Sprintf("Could not read raw state as map: %s", err))
		return
	}

	delete(rawMap, "script_name")

	cleanAttrTypes := make(map[string]tftypes.Type, len(rawMap))
	for k, v := range rawMap {
		cleanAttrTypes[k] = v.Type()
	}

	resp.State.Raw = tftypes.NewValue(tftypes.Object{AttributeTypes: cleanAttrTypes}, rawMap)
}

func extractString(m map[string]tftypes.Value, key string) types.String {
	val, ok := m[key]
	if !ok || val.IsNull() || !val.IsKnown() {
		return types.StringNull()
	}
	var s string
	if err := val.As(&s); err != nil {
		return types.StringNull()
	}
	return types.StringValue(s)
}

// UpgradeFromV1 handles version 1 state — always a no-op.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading workers_route state from version=1 to current (no-op)")
	resp.State.Raw = req.State.Raw
}
