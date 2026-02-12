// Package v500 contains the migration logic for cloudflare_access_rule from v4 to v5.
package v500

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceV4AccessRuleModel represents the legacy resource state from v4.x provider.
// Schema version: 1 (v4 schema version after v0→v1 migration)
// Resource type: cloudflare_access_rule
//
// Key differences from v5:
// - Configuration is stored as array (TypeList MaxItems:1)
// - No created_on, modified_on timestamps
// - No allowed_modes, scope computed fields
type SourceV4AccessRuleModel struct {
	// Resource identifier (implicit in SDKv2 but present in state)
	ID types.String `tfsdk:"id"`

	// Identifiers
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`

	// Required fields
	Mode          types.String                      `tfsdk:"mode"`
	Configuration []SourceV4ConfigurationModel      `tfsdk:"configuration"` // Array! (TypeList MaxItems:1)

	// Optional fields
	Notes types.String `tfsdk:"notes"`
}

// SourceV4ConfigurationModel represents the nested configuration structure from v4.x provider.
// This was stored as an array with MaxItems:1 due to SDK v2 TypeList behavior.
type SourceV4ConfigurationModel struct {
	Target types.String `tfsdk:"target"` // Required in v4
	Value  types.String `tfsdk:"value"`  // Required in v4
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetAccessRuleModel represents the current resource state from v5.x+ provider.
// Schema version: 500 (when TF_MIG_TEST=1, otherwise 1)
// Resource type: cloudflare_access_rule
//
// Key differences from v4:
// - Configuration is stored as object pointer (SingleNestedAttribute)
// - Includes id field
// - Includes created_on, modified_on timestamps
// - Includes allowed_modes, scope computed fields
//
// Note: This matches the model in the parent package's model.go file.
// We duplicate it here to keep the migration package self-contained.
type TargetAccessRuleModel struct {
	// Identifiers
	ID        types.String `tfsdk:"id"`
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`

	// Required fields
	Mode          types.String                  `tfsdk:"mode"`
	Configuration *TargetConfigurationModel     `tfsdk:"configuration"` // Pointer! (SingleNestedAttribute)

	// Optional/Computed fields
	Notes types.String `tfsdk:"notes"` // computed_optional (default: "")

	// Computed fields (new in v5)
	CreatedOn    timetypes.RFC3339                              `tfsdk:"created_on"`
	ModifiedOn   timetypes.RFC3339                              `tfsdk:"modified_on"`
	AllowedModes customfield.List[types.String]                 `tfsdk:"allowed_modes"`
	Scope        customfield.NestedObject[TargetScopeModel]     `tfsdk:"scope"`
}

// TargetConfigurationModel represents the nested configuration structure from v5.x+ provider.
// This is stored as a single object (not an array) in v5 state.
type TargetConfigurationModel struct {
	Target types.String `tfsdk:"target"` // Optional in v5
	Value  types.String `tfsdk:"value"`  // Optional in v5
}

// TargetScopeModel represents the scope computed field from v5.x+ provider.
// This is a new field that doesn't exist in v4.
type TargetScopeModel struct {
	ID    types.String `tfsdk:"id"`
	Email types.String `tfsdk:"email"`
	Type  types.String `tfsdk:"type"`
}
