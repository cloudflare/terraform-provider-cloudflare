package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UpgradeFromV0 handles state upgrades from v4 (schema_version=0) to v5.
// This is the main v4→v5 migration path for cloudflare_list_item.
// Key transformations:
//   - hostname: List (MaxItems:1) → SingleNestedAttribute
//   - redirect: List (MaxItems:1) → SingleNestedAttribute
//   - redirect boolean fields: "enabled"/"disabled" strings → true/false booleans
func UpgradeFromV0(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_list_item state from v4 (schema_version=0)")

	var sourceState SourceListItemModel
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	targetState, diags := Transform(ctx, sourceState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, targetState)...)

	tflog.Info(ctx, "State upgrade for cloudflare_list_item from v4 completed successfully")
}

// UpgradeFromV1 handles state upgrades from v4.52.5 framework (schema_version=1) to v500.
// In v4.52.5, hostname and redirect are ListNestedBlocks (lists with max 1 element).
// This upgrade converts them to SingleNestedAttributes (objects).
// Redirect boolean fields are already bools in v1 (no string conversion needed).
//
// IMPORTANT: This must only be called when req.State was populated with the v4.52.5 PriorSchema.
// When called from UpgradeFromV1Ambiguous (where PriorSchema is nil), use
// upgradeFromV1ViaRawState instead.
func UpgradeFromV1(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_list_item state from v4.52.5 framework (schema_version=1)")

	var sourceState SourceListItemV1Model
	resp.Diagnostics.Append(req.State.Get(ctx, &sourceState)...)
	if resp.Diagnostics.HasError() {
		return
	}

	target := &TargetListItemModel{
		AccountID: sourceState.AccountID,
		ListID:    sourceState.ListID,
		ID:        sourceState.ID,
		IP:        sourceState.IP,
		ASN:       sourceState.ASN,
		Comment:   sourceState.Comment,
	}

	// Hostname: List (max 1) → SingleNestedAttribute
	if len(sourceState.Hostname) > 0 {
		target.Hostname = customfield.NewObjectMust[TargetHostnameModel](ctx, &TargetHostnameModel{
			URLHostname: sourceState.Hostname[0].URLHostname,
		})
	} else {
		target.Hostname = customfield.NullObject[TargetHostnameModel](ctx)
	}

	// Redirect: List (max 1) → SingleNestedAttribute (booleans are already bools)
	if len(sourceState.Redirect) > 0 {
		src := sourceState.Redirect[0]
		target.Redirect = customfield.NewObjectMust[TargetRedirectModel](ctx, &TargetRedirectModel{
			SourceURL:           ensureSourceURLHasPath(src.SourceURL),
			TargetURL:           src.TargetURL,
			StatusCode:          src.StatusCode,
			IncludeSubdomains:   src.IncludeSubdomains,
			SubpathMatching:     src.SubpathMatching,
			PreserveQueryString: src.PreserveQueryString,
			PreservePathSuffix:  src.PreservePathSuffix,
		})
	} else {
		target.Redirect = customfield.NullObject[TargetRedirectModel](ctx)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, target)...)
	tflog.Info(ctx, "State upgrade for cloudflare_list_item from v1 completed successfully")
}

// UpgradeFromV1Ambiguous handles schema_version=1 state that is AMBIGUOUS between
// v4.52.5 Plugin Framework format and v5 production format.
//
// Both v4.52.5 and v5 production (v5.0-v5.18 with GetSchemaVersion(1,500)) stored state
// at schema_version=1, but with incompatible formats:
//   - v4.52.5: hostname/redirect stored as ARRAY (ListNestedBlock)
//   - v5: hostname/redirect stored as OBJECT (SingleNestedAttribute)
//
// migrations.go registers this with PriorSchema=nil so the framework skips pre-decoding
// req.State entirely. Both paths operate exclusively on req.RawState.JSON.
//
// Detection: inspect hostname/redirect in raw JSON.
// If array -> v4.52.5 format -> transform via raw state.
// If object (or absent) -> v5 format -> no-op re-decode with target schema.
func UpgradeFromV1Ambiguous(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	tflog.Info(ctx, "Upgrading cloudflare_list_item state from schema_version=1 (detecting v4.52.5 vs v5 format)")

	isV4, err := detectV4ListItemState(req)
	if err != nil {
		// Cannot determine format — default to no-op to avoid data loss.
		tflog.Warn(ctx, "Could not detect list_item state format, defaulting to no-op", map[string]interface{}{
			"error": err.Error(),
		})
		// Fall through to no-op path
	}

	if isV4 {
		tflog.Info(ctx, "Detected v4.52.5 list_item format (hostname/redirect is array), performing transformation via raw state")
		upgradeFromV1ViaRawState(ctx, req, resp)
		return
	}

	// v5 production state: re-decode raw JSON with the target schema.
	// req.State is nil here (PriorSchema=nil), so we must use req.RawState directly.
	tflog.Info(ctx, "Detected v5 list_item format (hostname/redirect is object), performing no-op version bump via raw state")
	targetType := resp.State.Schema.Type().TerraformType(ctx)
	rawValue, unmarshalErr := req.RawState.Unmarshal(targetType)
	if unmarshalErr != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v5 list_item state",
			"Could not parse raw state as v5 format: "+unmarshalErr.Error(),
		)
		return
	}
	resp.State.Raw = rawValue
	tflog.Info(ctx, "State version bump from 1 to 500 completed")
}

// upgradeFromV1ViaRawState performs the v4.52.5->v5 transformation by parsing req.RawState
// directly with the v4.52.5 schema. This is necessary when the upgrader was registered with
// PriorSchema=nil (so req.State is nil and cannot be used to parse v4.52.5 arrays).
func upgradeFromV1ViaRawState(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
	if req.RawState == nil {
		resp.Diagnostics.AddError("Missing raw state", "RawState was nil during v4.52.5 list_item upgrade")
		return
	}

	// Parse raw state using the v4.52.5 schema
	v1Schema := SourceListItemV1Schema()
	v1Type := v1Schema.Type().TerraformType(ctx)

	rawValue, err := req.RawState.Unmarshal(v1Type)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to unmarshal v4.52.5 list_item state",
			fmt.Sprintf("Could not parse raw state as v4.52.5 format: %s", err),
		)
		return
	}

	// Build a synthetic req with the v4.52.5-typed state so UpgradeFromV1 can call req.State.Get
	v1State := &tfsdk.State{Raw: rawValue, Schema: v1Schema}
	syntheticReq := resource.UpgradeStateRequest{
		RawState: req.RawState,
		State:    v1State,
	}

	UpgradeFromV1(ctx, syntheticReq, resp)
}

// detectV4ListItemState returns true if the raw state is in v4.52.5 format.
// v4.52.5 format: hostname/redirect are JSON arrays [] (ListNestedBlock).
// v5 format: hostname/redirect are JSON objects {} (SingleNestedAttribute) or null.
func detectV4ListItemState(req resource.UpgradeStateRequest) (bool, error) {
	if req.RawState == nil || len(req.RawState.JSON) == 0 {
		return false, nil
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(req.RawState.JSON, &raw); err != nil {
		return false, err
	}

	// Check hostname field: array = v4.52.5, object = v5
	if hostname, ok := raw["hostname"]; ok && hostname != nil {
		switch hostname.(type) {
		case []interface{}:
			return true, nil // v4.52.5 ListNestedBlock
		case map[string]interface{}:
			return false, nil // v5 SingleNestedAttribute
		}
	}

	// Check redirect field: array = v4.52.5, object = v5
	if redirect, ok := raw["redirect"]; ok && redirect != nil {
		switch redirect.(type) {
		case []interface{}:
			return true, nil // v4.52.5 ListNestedBlock
		case map[string]interface{}:
			return false, nil // v5 SingleNestedAttribute
		}
	}

	// Neither hostname nor redirect present (e.g., IP-only or ASN-only list item).
	// Cannot determine format, but no transformation needed either — assume v5.
	return false, nil
}
