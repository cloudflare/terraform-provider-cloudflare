package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (v0 — earliest v5 releases with map-based resources)
// ============================================================================

// SourceMetaV0 represents the permission group meta object in v0 state.
type SourceMetaV0 struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// SourcePermissionGroupV0 represents a permission group in v0 state.
// Includes computed fields (meta, name) that are removed in v500.
type SourcePermissionGroupV0 struct {
	ID   types.String  `tfsdk:"id"`
	Meta *SourceMetaV0 `tfsdk:"meta"`
	Name types.String  `tfsdk:"name"`
}

// SourcePolicyV0 represents a policy in v0 state.
// Key differences from v500:
// - ID field exists (computed, removed in v500)
// - Resources is map[string]types.String (converted to JSON string in v500)
// - PermissionGroups includes meta + name (removed in v500)
type SourcePolicyV0 struct {
	ID               types.String              `tfsdk:"id"`
	Effect           types.String              `tfsdk:"effect"`
	PermissionGroups []SourcePermissionGroupV0 `tfsdk:"permission_groups"`
	Resources        map[string]types.String   `tfsdk:"resources"`
}

// SourceConditionRequestIPV0 represents the request_ip condition in v0 state.
type SourceConditionRequestIPV0 struct {
	In    []types.String `tfsdk:"in"`
	NotIn []types.String `tfsdk:"not_in"`
}

// SourceConditionV0 represents the condition object in v0 state.
type SourceConditionV0 struct {
	RequestIP *SourceConditionRequestIPV0 `tfsdk:"request_ip"`
}

// SourceAccountTokenModelV0 represents the full account_token state in v0.
type SourceAccountTokenModelV0 struct {
	AccountID  types.String       `tfsdk:"account_id"`
	ID         types.String       `tfsdk:"id"`
	IssuedOn   timetypes.RFC3339  `tfsdk:"issued_on"`
	ModifiedOn timetypes.RFC3339  `tfsdk:"modified_on"`
	Name       types.String       `tfsdk:"name"`
	Policies   []SourcePolicyV0   `tfsdk:"policies"`
	Status     types.String       `tfsdk:"status"`
	Value      types.String       `tfsdk:"value"`
	NotBefore  timetypes.RFC3339  `tfsdk:"not_before"`
	ExpiresOn  timetypes.RFC3339  `tfsdk:"expires_on"`
	Condition  *SourceConditionV0 `tfsdk:"condition"`
	LastUsedOn timetypes.RFC3339  `tfsdk:"last_used_on"`
}

// ============================================================================
// Target Models (v500 — current schema)
// ============================================================================

// TargetPermissionGroupV500 represents a permission group in v500 state.
// Only contains id (meta and name are removed).
type TargetPermissionGroupV500 struct {
	ID types.String `tfsdk:"id"`
}

// TargetPolicyV500 represents a policy in v500 state.
// Key differences from v0:
// - No ID field
// - Resources is types.String (JSON-encoded)
// - PermissionGroups only has id
type TargetPolicyV500 struct {
	Effect           types.String                `tfsdk:"effect"`
	PermissionGroups []TargetPermissionGroupV500 `tfsdk:"permission_groups"`
	Resources        types.String                `tfsdk:"resources"`
}

// TargetAccountTokenModelV500 represents the full account_token state in v500.
// Used as the output model for the v0→v500 transform.
type TargetAccountTokenModelV500 struct {
	AccountID  types.String       `tfsdk:"account_id"`
	ID         types.String       `tfsdk:"id"`
	IssuedOn   timetypes.RFC3339  `tfsdk:"issued_on"`
	ModifiedOn timetypes.RFC3339  `tfsdk:"modified_on"`
	Name       types.String       `tfsdk:"name"`
	Policies   []TargetPolicyV500 `tfsdk:"policies"`
	Status     types.String       `tfsdk:"status"`
	Value      types.String       `tfsdk:"value"`
	NotBefore  timetypes.RFC3339  `tfsdk:"not_before"`
	ExpiresOn  timetypes.RFC3339  `tfsdk:"expires_on"`
	Condition  *SourceConditionV0 `tfsdk:"condition"`
	LastUsedOn timetypes.RFC3339  `tfsdk:"last_used_on"`
}
