package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source v4 state to target v5 state.
// Handles BOTH v4 resource names (they share identical schemas):
//   - cloudflare_access_organization
//   - cloudflare_zero_trust_access_organization
func Transform(ctx context.Context, source SourceCloudflareAccessOrganizationModel) (*TargetZeroTrustOrganizationModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetZeroTrustOrganizationModel{
		// Direct pass-through fields
		AuthDomain:                     source.AuthDomain,
		Name:                           source.Name,
		SessionDuration:                source.SessionDuration,
		UserSeatExpirationInactiveTime: source.UserSeatExpirationInactiveTime,
		WARPAuthSessionDuration:        source.WARPAuthSessionDuration,
		UIReadOnlyToggleReason:         source.UIReadOnlyToggleReason,
	}

	// Account ID / Zone ID mutual exclusivity
	// v5 requires exactly one to be set. Ensure the other is null.
	if !source.AccountID.IsNull() && source.AccountID.ValueString() != "" {
		target.AccountID = source.AccountID
		target.ZoneID = types.StringNull()
	} else if !source.ZoneID.IsNull() && source.ZoneID.ValueString() != "" {
		target.ZoneID = source.ZoneID
		target.AccountID = types.StringNull()
	} else {
		// Both null - preserve as-is (shouldn't happen in valid state, but handle gracefully)
		target.AccountID = source.AccountID
		target.ZoneID = source.ZoneID
	}

	// Boolean fields with defaults
	// v5 schema has explicit defaults (false). Add them if missing to prevent diffs.
	target.AllowAuthenticateViaWARP = ensureBooleanDefault(source.AllowAuthenticateViaWARP)
	target.AutoRedirectToIdentity = ensureBooleanDefault(source.AutoRedirectToIdentity)
	target.IsUIReadOnly = ensureBooleanDefault(source.IsUIReadOnly)

	// login_design: Convert from array (MaxItems:1) to object (SingleNestedAttribute)
	// SDK v2 stored this as []SourceLoginDesignModel
	// Framework expects *TargetLoginDesignModel
	if len(source.LoginDesign) > 0 {
		target.LoginDesign = convertLoginDesignArrayToObject(source.LoginDesign[0])
	}
	// If empty array, leave as nil (field not set)

	// custom_pages: Convert from array (MaxItems:1) to object (SingleNestedAttribute)
	// SDK v2 stored this as []SourceCustomPagesModel
	// Framework expects *TargetCustomPagesModel
	if len(source.CustomPages) > 0 {
		target.CustomPages = convertCustomPagesArrayToObject(source.CustomPages[0])
	}
	// If empty array, leave as nil (field not set)

	// New v5 fields not present in v4:
	// - deny_unmatched_requests: Leave as null
	// - deny_unmatched_requests_exempted_zone_names: Leave as nil
	// - mfa_config: Leave as nil
	// - mfa_configuration_allowed: Leave as null
	// - mfa_required_for_all_apps: Leave as null
	// These will be refreshed from API or remain unset

	return target, diags
}

// ensureBooleanDefault returns false if the value is null or unknown.
// Preserves existing true/false values.
func ensureBooleanDefault(field types.Bool) types.Bool {
	if field.IsNull() || field.IsUnknown() {
		return types.BoolValue(false)
	}
	return field
}

// convertLoginDesignArrayToObject converts the v4 login_design array element to v5 object.
func convertLoginDesignArrayToObject(source SourceLoginDesignModel) *TargetLoginDesignModel {
	return &TargetLoginDesignModel{
		BackgroundColor: source.BackgroundColor,
		TextColor:       source.TextColor,
		LogoPath:        source.LogoPath,
		HeaderText:      source.HeaderText,
		FooterText:      source.FooterText,
	}
}

// convertCustomPagesArrayToObject converts the v4 custom_pages array element to v5 object.
func convertCustomPagesArrayToObject(source SourceCustomPagesModel) *TargetCustomPagesModel {
	return &TargetCustomPagesModel{
		Forbidden:      source.Forbidden,
		IdentityDenied: source.IdentityDenied,
	}
}
