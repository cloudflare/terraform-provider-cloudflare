package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceAPITokenModel represents the v4 SDKv2 state structure.
// Schema version: 0 (SDKv2 default)
// Resource type: cloudflare_api_token
//
// Key differences from v5:
//   - "policy" field (not "policies")
//   - permission_groups as Set of strings (not objects)
//   - resources as map (not JSON string)
//   - condition/request_ip as arrays (not single nested)
//   - timestamps as types.String (not timetypes.RFC3339)
type SourceAPITokenModel struct {
	ID         types.String           `tfsdk:"id"`
	Name       types.String           `tfsdk:"name"`
	Policy     []SourcePolicyModel    `tfsdk:"policy"`
	Status     types.String           `tfsdk:"status"`
	Value      types.String           `tfsdk:"value"`
	IssuedOn   types.String           `tfsdk:"issued_on"`
	ModifiedOn types.String           `tfsdk:"modified_on"`
	ExpiresOn  types.String           `tfsdk:"expires_on"`
	NotBefore  types.String           `tfsdk:"not_before"`
	Condition  []SourceConditionModel `tfsdk:"condition"` // array, MaxItems:1
}

// SourcePolicyModel represents a v4 policy block.
type SourcePolicyModel struct {
	ID               types.String            `tfsdk:"id"`                // computed, removed in v5
	Effect           types.String            `tfsdk:"effect"`
	PermissionGroups types.Set               `tfsdk:"permission_groups"` // Set of strings
	Resources        map[string]types.String `tfsdk:"resources"`         // Map of strings
}

// SourceConditionModel represents a v4 condition block (stored as array).
type SourceConditionModel struct {
	RequestIP []SourceRequestIPModel `tfsdk:"request_ip"` // array, MaxItems:1
}

// SourceRequestIPModel represents a v4 request_ip block (stored as array).
type SourceRequestIPModel struct {
	In    types.List `tfsdk:"in"`
	NotIn types.List `tfsdk:"not_in"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAPITokenModel represents the current v5 state structure.
// Schema version: 500
// Resource type: cloudflare_api_token
//
// Note: This must match the APITokenModel in the parent package's model.go file.
type TargetAPITokenModel struct {
	ID         types.String                `tfsdk:"id"`
	Name       types.String                `tfsdk:"name"`
	Policies   *[]*TargetPolicyModel       `tfsdk:"policies"`
	ExpiresOn  timetypes.RFC3339           `tfsdk:"expires_on"`
	NotBefore  timetypes.RFC3339           `tfsdk:"not_before"`
	Condition  *TargetConditionModel       `tfsdk:"condition"`
	Status     types.String                `tfsdk:"status"`
	IssuedOn   timetypes.RFC3339           `tfsdk:"issued_on"`
	LastUsedOn timetypes.RFC3339           `tfsdk:"last_used_on"`
	ModifiedOn timetypes.RFC3339           `tfsdk:"modified_on"`
	Value      types.String                `tfsdk:"value"`
}

// TargetPolicyModel represents a v5 policy in the policies set.
type TargetPolicyModel struct {
	Effect           types.String                    `tfsdk:"effect"`
	PermissionGroups *[]*TargetPermissionGroupModel  `tfsdk:"permission_groups"`
	Resources        types.String                    `tfsdk:"resources"` // JSON string
}

// TargetPermissionGroupModel represents a v5 permission group object.
type TargetPermissionGroupModel struct {
	ID types.String `tfsdk:"id"`
}

// TargetConditionModel represents a v5 condition (single nested).
type TargetConditionModel struct {
	RequestIP *TargetRequestIPModel `tfsdk:"request_ip"`
}

// TargetRequestIPModel represents a v5 request_ip (single nested).
type TargetRequestIPModel struct {
	In    *[]types.String `tfsdk:"in"`
	NotIn *[]types.String `tfsdk:"not_in"`
}
