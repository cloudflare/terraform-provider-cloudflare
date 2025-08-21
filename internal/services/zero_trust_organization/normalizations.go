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

// Normalizing function to ensure consistency between the state/plan and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustOrganizationAPIData(_ context.Context, data, sourceData *ZeroTrustOrganizationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	normalizeFalseAndNullBool(&data.AutoRedirectToIdentity, sourceData.AutoRedirectToIdentity)
	normalizeFalseAndNullBool(&data.AllowAuthenticateViaWARP, sourceData.AllowAuthenticateViaWARP)
	normalizeFalseAndNullBool(&data.IsUIReadOnly, sourceData.IsUIReadOnly)
	normalizeEmptyAndNullObject(&data.LoginDesign, sourceData.LoginDesign)

	return diags
}

func normalizeImportZeroTrustOrganizationAPIData(_ context.Context, data *ZeroTrustOrganizationModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	if data.AutoRedirectToIdentity.IsNull() {
		data.AutoRedirectToIdentity = types.BoolValue(false)
	}

	var empty ZeroTrustOrganizationLoginDesignModel
	if data.LoginDesign != nil || *data.LoginDesign == empty {
		data.LoginDesign = nil
	}

	return diags
}
