package ruleset_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijsoncustom"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// RulesetRuleModel represents a single rule within a ruleset with resource-specific fields
type RulesetRuleModel struct {
	// Embed the existing rule model to reuse all rule fields
	ruleset.RulesetRulesModel

	RulesetID types.String `tfsdk:"ruleset_id"`
	AccountID types.String `tfsdk:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id"`

	Position *PositionModel `tfsdk:"position" json:"position,omitempty"`
}

func (m RulesetRuleModel) MarshalJSON() (data []byte, err error) {
	return apijsoncustom.MarshalRoot(m)
}

func (m RulesetRuleModel) MarshalJSONForUpdate(state RulesetRuleModel) (data []byte, err error) {
	return apijsoncustom.MarshalForUpdate(m, state)
}

type PositionModel struct {
	Index  *int    `tfsdk:"index"  json:"index,omitempty"`
	Before *string `tfsdk:"before" json:"before,omitempty"`
	After  *string `tfsdk:"after"  json:"after,omitempty"`
}
