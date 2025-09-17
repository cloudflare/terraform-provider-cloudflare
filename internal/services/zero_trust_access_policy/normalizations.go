package zero_trust_access_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

type SetValueInterface interface {
	IsNull() bool
	Elements() []attr.Value
}

func normalizeEmptyAndNullNestedObjectSet[T SetValueInterface](data *T, stateData T) {
	if (*data).IsNull() && stateData.IsNull() {
		return
	}
	if (!(*data).IsNull() && len((*data).Elements()) != 0) || (!stateData.IsNull() && len(stateData.Elements()) != 0) {
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
func normalizeReadZeroTrustAccessPolicyAPIData(ctx context.Context, data, sourceData *ZeroTrustAccessPolicyModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	// For rule fields, if they were null in the source (plan/state) and the API returned empty sets,
	// convert them back to null to prevent "inconsistent result after apply" errors
	if sourceData.Include.IsNull() && !data.Include.IsNull() && len(data.Include.Elements()) == 0 {
		data.Include = customfield.NullObjectSet[ZeroTrustAccessPolicyIncludeModel](ctx)
	}
	if sourceData.Require.IsNull() && !data.Require.IsNull() && len(data.Require.Elements()) == 0 {
		data.Require = customfield.NullObjectSet[ZeroTrustAccessPolicyRequireModel](ctx)
	}
	if sourceData.Exclude.IsNull() && !data.Exclude.IsNull() && len(data.Exclude.Elements()) == 0 {
		data.Exclude = customfield.NullObjectSet[ZeroTrustAccessPolicyExcludeModel](ctx)
	}
	
	// For non-empty rule fields, use standard normalization
	if !sourceData.Include.IsNull() || (!data.Include.IsNull() && len(data.Include.Elements()) > 0) {
		normalizeEmptyAndNullNestedObjectSet(&data.Include, sourceData.Include)
	}
	if !sourceData.Require.IsNull() || (!data.Require.IsNull() && len(data.Require.Elements()) > 0) {
		normalizeEmptyAndNullNestedObjectSet(&data.Require, sourceData.Require)
	}
	if !sourceData.Exclude.IsNull() || (!data.Exclude.IsNull() && len(data.Exclude.Elements()) > 0) {
		normalizeEmptyAndNullNestedObjectSet(&data.Exclude, sourceData.Exclude)
	}
	
	// For other fields, use the original normalization logic
	normalizeEmptyAndNullSlice(&data.ApprovalGroups, sourceData.ApprovalGroups)
	normalizeFalseAndNullBool(&data.PurposeJustificationRequired, sourceData.PurposeJustificationRequired)
	normalizeFalseAndNullBool(&data.ApprovalRequired, sourceData.ApprovalRequired)
	normalizeFalseAndNullBool(&data.IsolationRequired, sourceData.IsolationRequired)

	return diags
}

// Specialized normalization for import operations where API omits false boolean values
func normalizeImportZeroTrustAccessPolicyAPIData(ctx context.Context, data *ZeroTrustAccessPolicyModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	if !data.Include.IsNull() && len(data.Include.Elements()) == 0 {
		data.Include = customfield.NullObjectSet[ZeroTrustAccessPolicyIncludeModel](ctx)
	}

	if !data.Require.IsNull() && len(data.Require.Elements()) == 0 {
		data.Require = customfield.NullObjectSet[ZeroTrustAccessPolicyRequireModel](ctx)
	}

	if !data.Exclude.IsNull() && len(data.Exclude.Elements()) == 0 {
		data.Exclude = customfield.NullObjectSet[ZeroTrustAccessPolicyExcludeModel](ctx)
	}

	return diags
}
