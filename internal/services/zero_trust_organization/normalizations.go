package zero_trust_organization

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

func normalizeEmptyAndNullObject[T comparable](data **T, stateData *T) {
	var zeroValue T
	if (*data != nil && **data != zeroValue) || (stateData != nil && *stateData != zeroValue) {
		return
	}
	*data = stateData
}

func normalizeFalseAndNullBool(data *basetypes.BoolValue, stateData basetypes.BoolValue) {
	if data.ValueBool() || stateData.ValueBool() {
		return
	}

	if stateData.IsUnknown() {
		if data.IsNull() || data.IsUnknown() {
			*data = types.BoolValue(false)
		}
		return
	}

	*data = stateData
}

// normalizeEmptyAndNullList laxes the equality between a null list and an empty
// list. It is generic over the element type so it can be used for both
// []types.String and []types.Int64 attributes.
func normalizeEmptyAndNullList[T any](data **[]T, stateData *[]T) {
	if (data != nil && *data != nil && len(**data) > 0) || (stateData != nil && len(*stateData) > 0) {
		return
	}
	*data = stateData
}

// normalizeUnknownCustomList preserves the prior state value for a
// customfield.List-typed attribute when the freshly decoded API data resolves
// to null or unknown, unless the prior state is itself unknown (e.g. on
// Create, where there is no prior state to fall back on).
func normalizeUnknownCustomList[T attr.Value](data *customfield.List[T], stateData customfield.List[T]) {
	if data.IsNull() || data.IsUnknown() {
		if !stateData.IsUnknown() {
			*data = stateData
		}
	}
}

func normalizeEmptyAndNullString(data *basetypes.StringValue, stateData basetypes.StringValue) {
	// If data is unknown or null/empty, preserve state value (unless state is also unknown)
	if data.IsUnknown() || data.IsNull() || data.ValueString() == "" {
		if !stateData.IsUnknown() {
			*data = stateData
		} else {
			// If both are unknown, set to null to satisfy Terraform's requirement
			*data = types.StringNull()
		}
		return
	}
	// If data has a non-empty value, keep it
}

// Normalizing function to ensure consistency between the state/plan and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustOrganizationAPIData(_ context.Context, data, sourceData *ZeroTrustOrganizationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	normalizeFalseAndNullBool(&data.AutoRedirectToIdentity, sourceData.AutoRedirectToIdentity)
	normalizeFalseAndNullBool(&data.AllowAuthenticateViaWARP, sourceData.AllowAuthenticateViaWARP)
	normalizeFalseAndNullBool(&data.IsUIReadOnly, sourceData.IsUIReadOnly)
	normalizeFalseAndNullBool(&data.DenyUnmatchedRequests, sourceData.DenyUnmatchedRequests)
	normalizeFalseAndNullBool(&data.MfaRequiredForAllApps, sourceData.MfaRequiredForAllApps)
	normalizeFalseAndNullBool(&data.MfaConfigurationAllowed, sourceData.MfaConfigurationAllowed)
	normalizeEmptyAndNullObject(&data.LoginDesign, sourceData.LoginDesign)
	// service_token_inactivity is not Computed, but the API can populate
	// action/inactivity_threshold_days with server-side defaults even when
	// enabled=false, so the generic all-fields-zero check in
	// normalizeEmptyAndNullObject never collapses it. Treat enabled=false as
	// the disabled/null signal instead.
	if data.ServiceTokenInactivity != nil && !data.ServiceTokenInactivity.Enabled.ValueBool() &&
		(sourceData.ServiceTokenInactivity == nil || !sourceData.ServiceTokenInactivity.Enabled.ValueBool()) {
		data.ServiceTokenInactivity = sourceData.ServiceTokenInactivity
	}
	normalizeUnknownCustomList(&data.TrustedAccounts, sourceData.TrustedAccounts)
	normalizeEmptyAndNullList(&data.DenyUnmatchedRequestsExemptedZoneNames, sourceData.DenyUnmatchedRequestsExemptedZoneNames)
	normalizeEmptyAndNullString(&data.UIReadOnlyToggleReason, sourceData.UIReadOnlyToggleReason)

	// The API serializes mfa_piv_key_requirements.ssh_key_type and .ssh_key_size with
	// `omitempty`, so an explicitly configured empty list round-trips as an absent field
	// and decodes back as null. Lax that equality to avoid a perpetual diff.
	if data.MfaSSHPivKeyRequirements != nil && sourceData.MfaSSHPivKeyRequirements != nil {
		normalizeEmptyAndNullList(&data.MfaSSHPivKeyRequirements.SSHKeyType, sourceData.MfaSSHPivKeyRequirements.SSHKeyType)
		normalizeEmptyAndNullList(&data.MfaSSHPivKeyRequirements.SSHKeySize, sourceData.MfaSSHPivKeyRequirements.SSHKeySize)
	}

	return diags
}

func normalizeImportZeroTrustOrganizationAPIData(_ context.Context, data *ZeroTrustOrganizationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	if data.AutoRedirectToIdentity.IsNull() {
		data.AutoRedirectToIdentity = types.BoolValue(false)
	}

	if data.MfaRequiredForAllApps.IsNull() {
		data.MfaRequiredForAllApps = types.BoolValue(false)
	}

	// Set LoginDesign to nil if all fields are empty/null
	if data.LoginDesign != nil {
		allEmpty := true
		if !data.LoginDesign.BackgroundColor.IsNull() && data.LoginDesign.BackgroundColor.ValueString() != "" {
			allEmpty = false
		}
		if !data.LoginDesign.FooterText.IsNull() && data.LoginDesign.FooterText.ValueString() != "" {
			allEmpty = false
		}
		if !data.LoginDesign.HeaderText.IsNull() && data.LoginDesign.HeaderText.ValueString() != "" {
			allEmpty = false
		}
		if !data.LoginDesign.LogoPath.IsNull() && data.LoginDesign.LogoPath.ValueString() != "" {
			allEmpty = false
		}
		if !data.LoginDesign.TextColor.IsNull() && data.LoginDesign.TextColor.ValueString() != "" {
			allEmpty = false
		}

		if allEmpty {
			data.LoginDesign = nil
		}
	}

	return diags
}
