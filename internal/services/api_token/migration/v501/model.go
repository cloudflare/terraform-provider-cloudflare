package v501

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// APITokenModelV500 represents the full api_token state in v500.
// Uses Go slices (not pointers) for compatibility with Set-based schema deserialization.
type APITokenModelV500 struct {
	ID         types.String      `tfsdk:"id"`
	IssuedOn   timetypes.RFC3339 `tfsdk:"issued_on"`
	ModifiedOn timetypes.RFC3339 `tfsdk:"modified_on"`
	Name       types.String      `tfsdk:"name"`
	Policies   []PolicyV500      `tfsdk:"policies"`
	Status     types.String      `tfsdk:"status"`
	Value      types.String      `tfsdk:"value"`
	NotBefore  timetypes.RFC3339 `tfsdk:"not_before"`
	ExpiresOn  timetypes.RFC3339 `tfsdk:"expires_on"`
	Condition  *ConditionV500    `tfsdk:"condition"`
	LastUsedOn timetypes.RFC3339 `tfsdk:"last_used_on"`
}

// PolicyV500 represents a policy in v500 state.
type PolicyV500 struct {
	Effect           types.String          `tfsdk:"effect"`
	PermissionGroups []PermissionGroupV500 `tfsdk:"permission_groups"`
	Resources        types.String          `tfsdk:"resources"`
}

// PermissionGroupV500 represents a permission group in v500 state.
type PermissionGroupV500 struct {
	ID types.String `tfsdk:"id"`
}

// ConditionV500 represents the condition object in v500 state.
type ConditionV500 struct {
	RequestIP *ConditionRequestIPV500 `tfsdk:"request_ip"`
}

// ConditionRequestIPV500 represents the request_ip condition in v500 state.
type ConditionRequestIPV500 struct {
	In    []types.String `tfsdk:"in"`
	NotIn []types.String `tfsdk:"not_in"`
}
