package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceSnippetRulesModel represents the legacy cloudflare_snippet_rules state from v4.x provider.
// Schema version: 1
// Resource type: cloudflare_snippet_rules
//
// In v4, rules were defined as repeated blocks. In state, they are stored as a JSON array.
// Field names are identical between v4 and v5, but v4 did not have computed fields (id, last_updated).
type SourceSnippetRulesModel struct {
	ZoneID types.String            `tfsdk:"zone_id"`
	Rules  []SourceSnippetRuleModel `tfsdk:"rules"`
}

// SourceSnippetRuleModel represents a single rule in the v4 state.
type SourceSnippetRuleModel struct {
	Expression  types.String `tfsdk:"expression"`
	SnippetName types.String `tfsdk:"snippet_name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Description types.String `tfsdk:"description"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetSnippetRulesModel represents the current cloudflare_snippet_rules state from v5.x+ provider.
// Schema version: 500
// Resource type: cloudflare_snippet_rules
//
// Types must match the v5 schema exactly so that resp.State.Set() succeeds:
// - Rules uses *[]*TargetSnippetRuleModel (matches SnippetRulesModel.Rules)
// - last_updated uses timetypes.RFC3339 (matches schema CustomType)
type TargetSnippetRulesModel struct {
	ZoneID types.String              `tfsdk:"zone_id"`
	Rules  *[]*TargetSnippetRuleModel `tfsdk:"rules"`
}

// TargetSnippetRuleModel represents a single rule in the v5 state.
type TargetSnippetRuleModel struct {
	ID          types.String      `tfsdk:"id"`
	Expression  types.String      `tfsdk:"expression"`
	LastUpdated timetypes.RFC3339 `tfsdk:"last_updated"`
	SnippetName types.String      `tfsdk:"snippet_name"`
	Description types.String      `tfsdk:"description"`
	Enabled     types.Bool        `tfsdk:"enabled"`
}
