package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4 SDKv2) state to target (current v5 Plugin Framework) state.
//
// Key transformations:
//   - zone → name (field rename)
//   - account_id (flat string) → account.id (nested object)
//   - jump_start → dropped (removed in v5)
//   - plan (string) → null (v5 is computed-only nested object; API repopulates)
//   - meta (TypeMap[Bool]) → null (incompatible type in v5; API repopulates)
//   - vanity_name_servers / name_servers: types.List → customfield.List[types.String]
//   - All new v5 computed fields → null (API repopulates on first read)
func Transform(ctx context.Context, source SourceCloudflareZoneModel) (*TargetZoneModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.Zone.IsNull() || source.Zone.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone is required for zone migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zone migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Step 2: Initialize target with direct copies and simple renames
	target := &TargetZoneModel{
		ID:              source.ID,
		Name:            source.Zone, // zone → name (renamed)
		Paused:          source.Paused,
		Type:            source.Type,
		Status:          source.Status,
		VerificationKey: source.VerificationKey,
	}

	// Step 3: account_id (flat string) → account.id (nested object)
	target.Account = &TargetZoneAccountModel{
		ID: source.AccountID,
	}

	// Step 4: vanity_name_servers — types.List → customfield.List[types.String]
	if !source.VanityNameServers.IsNull() && !source.VanityNameServers.IsUnknown() {
		vanityList, listDiags := convertListToCustomfieldList(ctx, source.VanityNameServers)
		diags.Append(listDiags...)
		if !diags.HasError() {
			target.VanityNameServers = vanityList
		}
	} else {
		target.VanityNameServers = customfield.NullList[types.String](ctx)
	}

	// Step 5: name_servers — types.List → customfield.List[types.String]
	if !source.NameServers.IsNull() && !source.NameServers.IsUnknown() {
		nameServersList, listDiags := convertListToCustomfieldList(ctx, source.NameServers)
		diags.Append(listDiags...)
		if !diags.HasError() {
			target.NameServers = nameServersList
		}
	} else {
		target.NameServers = customfield.NullList[types.String](ctx)
	}

	// Step 6: New v5 computed list fields — null (API repopulates)
	target.OriginalNameServers = customfield.NullList[types.String](ctx)
	target.Permissions = customfield.NullList[types.String](ctx)

	// Step 7: Computed nested object fields — null (API repopulates)
	// - plan: v4 string is dropped; v5 is a computed-only nested object
	// - meta: v4 TypeMap[Bool] is incompatible; v5 is a computed-only nested object
	target.Meta = customfield.NullObject[TargetZoneMetaModel](ctx)
	target.Owner = customfield.NullObject[TargetZoneOwnerModel](ctx)
	target.Plan = customfield.NullObject[TargetZonePlanModel](ctx)
	target.Tenant = customfield.NullObject[TargetZoneTenantModel](ctx)
	target.TenantUnit = customfield.NullObject[TargetZoneTenantUnitModel](ctx)

	// Step 8: New v5 timestamp fields — null (API repopulates)
	target.ActivatedOn = timetypes.NewRFC3339Null()
	target.CreatedOn = timetypes.NewRFC3339Null()
	target.ModifiedOn = timetypes.NewRFC3339Null()

	// Step 9: New v5 scalar computed fields — null (API repopulates)
	target.CNAMESuffix = types.StringNull()
	target.DevelopmentMode = types.Float64Null()
	target.OriginalDnshost = types.StringNull()
	target.OriginalRegistrar = types.StringNull()

	// NOTE: jump_start is intentionally not copied (removed in v5)
	// NOTE: plan string is intentionally not copied (v5 plan is computed-only nested object)

	return target, diags
}

// convertListToCustomfieldList converts a types.List (SDKv2-style) to customfield.List[types.String].
// Extracts to []string first to avoid attr.Value type-wrapping issues.
func convertListToCustomfieldList(ctx context.Context, list types.List) (customfield.List[types.String], diag.Diagnostics) {
	// Extract to []string first (avoids attr.Value wrapping issues)
	var rawStrings []string
	diags := list.ElementsAs(ctx, &rawStrings, false)
	if diags.HasError() {
		return customfield.NullList[types.String](ctx), diags
	}

	// Convert []string to []types.String
	result := make([]types.String, 0, len(rawStrings))
	for _, s := range rawStrings {
		result = append(result, types.StringValue(s))
	}

	cfList, listDiags := customfield.NewList[types.String](ctx, result)
	diags.Append(listDiags...)
	return cfList, diags
}
