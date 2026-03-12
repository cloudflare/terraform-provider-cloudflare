package zero_trust_access_policy

import (
	"context"
	"reflect"
	"strings"

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
	diags.Append(pruneEmptyConditionSelectors(ctx, data)...)

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

	// Normalize clipboard arrays inside connection_rules.rdp — the API omits empty arrays
	// due to omitempty, so [] in state vs null from API should not cause drift.
	if data.ConnectionRules != nil && sourceData.ConnectionRules != nil {
		if data.ConnectionRules.RDP != nil && sourceData.ConnectionRules.RDP != nil {
			normalizeEmptyAndNullSlice(&data.ConnectionRules.RDP.AllowedClipboardLocalToRemoteFormats, sourceData.ConnectionRules.RDP.AllowedClipboardLocalToRemoteFormats)
			normalizeEmptyAndNullSlice(&data.ConnectionRules.RDP.AllowedClipboardRemoteToLocalFormats, sourceData.ConnectionRules.RDP.AllowedClipboardRemoteToLocalFormats)
		}
	}

	return diags
}

func pruneEmptyConditionSelectors(ctx context.Context, data *ZeroTrustAccessPolicyModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	diags.Append(pruneEmptyIncludeSelectors(ctx, &data.Include)...)
	diags.Append(pruneEmptyExcludeSelectors(ctx, &data.Exclude)...)
	diags.Append(pruneEmptyRequireSelectors(ctx, &data.Require)...)
	return diags
}

func pruneEmptyIncludeSelectors(ctx context.Context, set *customfield.NestedObjectSet[ZeroTrustAccessPolicyIncludeModel]) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	if set.IsNull() || set.IsUnknown() {
		return diags
	}

	conditions, d := set.AsStructSliceT(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	filtered := make([]ZeroTrustAccessPolicyIncludeModel, 0, len(conditions))
	for _, c := range conditions {
		if c.Email != nil && isNullUnknownOrEmptyString(c.Email.Email) {
			c.Email = nil
		}
		if c.EmailDomain != nil && isNullUnknownOrEmptyString(c.EmailDomain.Domain) {
			c.EmailDomain = nil
		}
		if c.IP != nil && isNullUnknownOrEmptyString(c.IP.IP) {
			c.IP = nil
		}
		if c.Geo != nil && isNullUnknownOrEmptyString(c.Geo.CountryCode) {
			c.Geo = nil
		}

		if reflect.DeepEqual(c, ZeroTrustAccessPolicyIncludeModel{}) {
			continue
		}
		filtered = append(filtered, c)
	}

	next, d := customfield.NewObjectSet(ctx, filtered)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	*set = next
	return diags
}

func pruneEmptyExcludeSelectors(ctx context.Context, set *customfield.NestedObjectSet[ZeroTrustAccessPolicyExcludeModel]) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	if set.IsNull() || set.IsUnknown() {
		return diags
	}

	conditions, d := set.AsStructSliceT(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	filtered := make([]ZeroTrustAccessPolicyExcludeModel, 0, len(conditions))
	for _, c := range conditions {
		if c.Email != nil && isNullUnknownOrEmptyString(c.Email.Email) {
			c.Email = nil
		}
		if c.EmailDomain != nil && isNullUnknownOrEmptyString(c.EmailDomain.Domain) {
			c.EmailDomain = nil
		}
		if c.IP != nil && isNullUnknownOrEmptyString(c.IP.IP) {
			c.IP = nil
		}
		if c.Geo != nil && isNullUnknownOrEmptyString(c.Geo.CountryCode) {
			c.Geo = nil
		}

		if reflect.DeepEqual(c, ZeroTrustAccessPolicyExcludeModel{}) {
			continue
		}
		filtered = append(filtered, c)
	}

	next, d := customfield.NewObjectSet(ctx, filtered)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	*set = next
	return diags
}

func pruneEmptyRequireSelectors(ctx context.Context, set *customfield.NestedObjectSet[ZeroTrustAccessPolicyRequireModel]) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)
	if set.IsNull() || set.IsUnknown() {
		return diags
	}

	conditions, d := set.AsStructSliceT(ctx)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}

	filtered := make([]ZeroTrustAccessPolicyRequireModel, 0, len(conditions))
	for _, c := range conditions {
		if c.Email != nil && isNullUnknownOrEmptyString(c.Email.Email) {
			c.Email = nil
		}
		if c.EmailDomain != nil && isNullUnknownOrEmptyString(c.EmailDomain.Domain) {
			c.EmailDomain = nil
		}
		if c.IP != nil && isNullUnknownOrEmptyString(c.IP.IP) {
			c.IP = nil
		}
		if c.Geo != nil && isNullUnknownOrEmptyString(c.Geo.CountryCode) {
			c.Geo = nil
		}

		if reflect.DeepEqual(c, ZeroTrustAccessPolicyRequireModel{}) {
			continue
		}
		filtered = append(filtered, c)
	}

	next, d := customfield.NewObjectSet(ctx, filtered)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	*set = next
	return diags
}

func isNullUnknownOrEmptyString(v basetypes.StringValue) bool {
	if v.IsNull() || v.IsUnknown() {
		return true
	}
	return strings.TrimSpace(v.ValueString()) == ""
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
