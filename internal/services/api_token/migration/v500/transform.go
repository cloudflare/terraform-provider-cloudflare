package v500

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Transform converts source (v4 SDKv2) state to target (v5 Plugin Framework) state.
// This function handles all v4→v5 conversions:
//   - policy → policies (rename + restructure)
//   - permission_groups: strings → objects with id
//   - resources: map → JSON string
//   - condition/request_ip: arrays → single nested objects
//   - timestamps: types.String → timetypes.RFC3339
//   - last_used_on: initialized as null (not present in v4)
//   - policy.id: dropped (removed from v5 schema)
func Transform(ctx context.Context, source SourceAPITokenModel) (*TargetAPITokenModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetAPITokenModel{
		ID:     source.ID,
		Name:   source.Name,
		Status: source.Status,
		Value:  source.Value,
	}

	// Convert timestamps: types.String → timetypes.RFC3339
	target.IssuedOn = convertToRFC3339(source.IssuedOn, &diags)
	target.ModifiedOn = convertToRFC3339(source.ModifiedOn, &diags)
	target.ExpiresOn = convertToRFC3339(source.ExpiresOn, &diags)
	target.NotBefore = convertToRFC3339(source.NotBefore, &diags)
	// last_used_on does not exist in v4 state
	target.LastUsedOn = timetypes.NewRFC3339Null()

	if diags.HasError() {
		return nil, diags
	}

	// Convert policies: policy[] → policies (rename + restructure)
	policies := make([]*TargetPolicyModel, 0, len(source.Policy))
	for _, srcPolicy := range source.Policy {
		policy, d := transformPolicy(ctx, srcPolicy)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		policies = append(policies, policy)
	}
	target.Policies = &policies

	// Convert condition: array[0] → pointer to object
	if len(source.Condition) > 0 {
		cond, d := transformCondition(ctx, source.Condition[0])
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		target.Condition = cond
	}
	// If condition array is empty or absent, target.Condition remains nil

	return target, diags
}

// transformPolicy converts a v4 policy to v5 format.
// - permission_groups: Set<string> → *[]*{id: string}
// - resources: map[string]string → JSON string
// - policy.id: dropped (not in v5 schema)
func transformPolicy(ctx context.Context, src SourcePolicyModel) (*TargetPolicyModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Convert permission_groups: Set<string> → *[]*TargetPermissionGroupModel
	var pgStrings []types.String
	diags.Append(src.PermissionGroups.ElementsAs(ctx, &pgStrings, false)...)
	if diags.HasError() {
		return nil, diags
	}
	permGroups := make([]*TargetPermissionGroupModel, 0, len(pgStrings))
	for _, pgStr := range pgStrings {
		permGroups = append(permGroups, &TargetPermissionGroupModel{ID: pgStr})
	}

	// Convert resources: map[string]string → JSON string
	resourcesMap := make(map[string]string)
	for k, v := range src.Resources {
		if !v.IsNull() && !v.IsUnknown() {
			resourcesMap[k] = v.ValueString()
		}
	}
	resourcesJSON, err := json.Marshal(resourcesMap)
	if err != nil {
		diags.AddError("State Upgrade Error",
			fmt.Sprintf("Failed to convert resources to JSON: %v", err))
		return nil, diags
	}

	return &TargetPolicyModel{
		Effect:           src.Effect,
		PermissionGroups: &permGroups,
		Resources:        types.StringValue(string(resourcesJSON)),
	}, diags
}

// transformCondition converts a v4 condition (from array element) to v5 single nested object.
func transformCondition(ctx context.Context, src SourceConditionModel) (*TargetConditionModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(src.RequestIP) == 0 {
		return &TargetConditionModel{}, diags
	}

	reqIP := src.RequestIP[0]
	targetReqIP := &TargetRequestIPModel{}

	// Convert In: types.List → *[]types.String
	if !reqIP.In.IsNull() && !reqIP.In.IsUnknown() {
		var inStrings []types.String
		diags.Append(reqIP.In.ElementsAs(ctx, &inStrings, false)...)
		if diags.HasError() {
			return nil, diags
		}
		targetReqIP.In = &inStrings
	}

	// Convert NotIn: types.List → *[]types.String
	if !reqIP.NotIn.IsNull() && !reqIP.NotIn.IsUnknown() {
		var notInStrings []types.String
		diags.Append(reqIP.NotIn.ElementsAs(ctx, &notInStrings, false)...)
		if diags.HasError() {
			return nil, diags
		}
		targetReqIP.NotIn = &notInStrings
	}

	return &TargetConditionModel{
		RequestIP: targetReqIP,
	}, diags
}

// convertToRFC3339 converts a types.String to timetypes.RFC3339.
func convertToRFC3339(s types.String, diags *diag.Diagnostics) timetypes.RFC3339 {
	if s.IsNull() || s.IsUnknown() || s.ValueString() == "" {
		return timetypes.NewRFC3339Null()
	}
	val, d := timetypes.NewRFC3339Value(s.ValueString())
	diags.Append(d...)
	return val
}
