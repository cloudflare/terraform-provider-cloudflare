// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultEnvelope struct {
	Result RulesetModel `json:"result,computed"`
}

type RulesetModel struct {
	ID          types.String          `tfsdk:"id" json:"id,computed"`
	AccountID   types.String          `tfsdk:"account_id" path:"account_id"`
	ZoneID      types.String          `tfsdk:"zone_id" path:"zone_id"`
	Kind        types.String          `tfsdk:"kind" json:"kind"`
	Name        types.String          `tfsdk:"name" json:"name"`
	Phase       types.String          `tfsdk:"phase" json:"phase"`
	Rules       *[]*RulesetRulesModel `tfsdk:"rules" json:"rules"`
	Description types.String          `tfsdk:"description" json:"description"`
}

type RulesetRulesModel struct {
	LastUpdated      types.String                       `tfsdk:"last_updated" json:"last_updated,computed"`
	Version          types.String                       `tfsdk:"version" json:"version,computed"`
	ID               types.String                       `tfsdk:"id" json:"id"`
	Action           types.String                       `tfsdk:"action" json:"action"`
	ActionParameters *RulesetRulesActionParametersModel `tfsdk:"action_parameters" json:"action_parameters"`
	Categories       types.String                       `tfsdk:"categories" json:"categories,computed"`
	Description      types.String                       `tfsdk:"description" json:"description,computed"`
	Enabled          types.Bool                         `tfsdk:"enabled" json:"enabled,computed"`
	Expression       types.String                       `tfsdk:"expression" json:"expression"`
	Logging          *RulesetRulesLoggingModel          `tfsdk:"logging" json:"logging"`
	Ref              types.String                       `tfsdk:"ref" json:"ref"`
}

type RulesetRulesActionParametersModel struct {
	Response *RulesetRulesActionParametersResponseModel `tfsdk:"response" json:"response"`
}

type RulesetRulesActionParametersResponseModel struct {
	Content     types.String `tfsdk:"content" json:"content"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code"`
}

type RulesetRulesLoggingModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}
