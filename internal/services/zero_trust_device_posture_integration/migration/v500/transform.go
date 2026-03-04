package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (legacy v4) state to target (current v5) state.
// This function is shared by both UpgradeFromV4 and MoveState handlers.
//
// Transformations performed:
// - Direct copy: id, account_id, name, type (all strings, no changes)
// - interval: Provide default "24h" if missing (optional→required)
// - config: Transform from array [{...}] to pointer object {...} (TypeList MaxItems:1 → SingleNested)
// - config: Provide empty object if missing (optional→required)
// - identifier: Drop field (removed in v5)
func Transform(ctx context.Context, source SourceCloudflareDevicePostureIntegrationModel) (*TargetZeroTrustDevicePostureIntegrationModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for zero_trust_device_posture_integration migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.Name.IsNull() || source.Name.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"name is required for zero_trust_device_posture_integration migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.Type.IsNull() || source.Type.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"type is required for zero_trust_device_posture_integration migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Initialize target with direct copies (simple string fields)
	target := &TargetZeroTrustDevicePostureIntegrationModel{
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,
		Type:      source.Type,
	}

	// Handle interval field: Optional in v4 → Required in v5
	// Provide default "24h" if missing
	if !source.Interval.IsNull() && !source.Interval.IsUnknown() && source.Interval.ValueString() != "" {
		target.Interval = source.Interval
	} else {
		// Default value for v5 requirement
		target.Interval = types.StringValue("24h")
	}

	// Handle config field: TypeList MaxItems:1 (array) → SingleNested (pointer)
	// Also: Optional in v4 → Required in v5
	target.Config = transformConfig(source.Config)

	// identifier field is intentionally NOT copied (removed in v5)

	return target, diags
}

// transformConfig converts v4 config array to v5 config pointer object.
//
// v4: TypeList MaxItems:1 stored as [{...}] (array with single element)
// v5: SingleNestedAttribute stored as {...} (pointer to object)
//
// If v4 has no config (or empty array), provide empty object for v5 requirement.
//
// IMPORTANT: Converts empty strings to null for optional fields (SDK v2 → Framework behavior)
func transformConfig(sourceConfig []SourceConfigModel) *TargetConfigModel {
	// If v4 has config data, extract array[0] and transform to pointer
	if len(sourceConfig) > 0 {
		return &TargetConfigModel{
			APIURL:             convertEmptyStringToNull(sourceConfig[0].APIURL),
			AuthURL:            convertEmptyStringToNull(sourceConfig[0].AuthURL),
			ClientID:           convertEmptyStringToNull(sourceConfig[0].ClientID),
			ClientSecret:       convertEmptyStringToNull(sourceConfig[0].ClientSecret),
			CustomerID:         convertEmptyStringToNull(sourceConfig[0].CustomerID),
			ClientKey:          convertEmptyStringToNull(sourceConfig[0].ClientKey),
			AccessClientID:     convertEmptyStringToNull(sourceConfig[0].AccessClientID),
			AccessClientSecret: convertEmptyStringToNull(sourceConfig[0].AccessClientSecret),
		}
	}

	// v4 had no config or empty array
	// Provide empty config object for v5 requirement
	return &TargetConfigModel{}
}

// convertEmptyStringToNull converts empty strings to null for optional fields.
// This handles the SDK v2 → Framework migration where optional fields change from "" to null.
func convertEmptyStringToNull(value types.String) types.String {
	if value.IsNull() || value.IsUnknown() {
		return value
	}
	if value.ValueString() == "" {
		return types.StringNull()
	}
	return value
}
