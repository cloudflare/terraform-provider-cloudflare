package zero_trust_organization

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
	*data = stateData
}

func normalizeEmptyAndNullList(data **[]types.String, stateData *[]types.String) {
	if (data != nil && *data != nil && len(**data) > 0) || (stateData != nil && len(*stateData) > 0) {
		return
	}
	*data = stateData
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
	normalizeEmptyAndNullObject(&data.LoginDesign, sourceData.LoginDesign)
	normalizeEmptyAndNullList(&data.DenyUnmatchedRequestsExemptedZoneNames, sourceData.DenyUnmatchedRequestsExemptedZoneNames)
	normalizeEmptyAndNullString(&data.UIReadOnlyToggleReason, sourceData.UIReadOnlyToggleReason)

	return diags
}

func normalizeImportZeroTrustOrganizationAPIData(_ context.Context, data *ZeroTrustOrganizationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	if data.AutoRedirectToIdentity.IsNull() {
		data.AutoRedirectToIdentity = types.BoolValue(false)
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
