package zero_trust_access_policy

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type IsNull interface {
	IsNull() bool
}

func normalizeEmptyAndNullSlice[T any](data **[]T, stateData *[]T) {
	if (*data != nil && len(**data) != 0) || (stateData != nil && len(*stateData) != 0) {
		return
	}
	*data = stateData
}

// Normalizing function to ensure consistency between the state/plan and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustAccessPolicyAPIData(ctx context.Context, data, sourceData *ZeroTrustAccessPolicyModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	normalizeEmptyAndNullSlice(&data.Include, sourceData.Include)
	normalizeEmptyAndNullSlice(&data.Require, sourceData.Require)
	normalizeEmptyAndNullSlice(&data.Exclude, sourceData.Exclude)

	return diags
}
