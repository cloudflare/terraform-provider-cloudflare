// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultDataSourceEnvelope struct {
	Result RulesetDataSourceModel `json:"result,computed"`
}

type RulesetResultListDataSourceEnvelope struct {
	Result *[]*RulesetDataSourceModel `json:"result,computed"`
}

type RulesetDataSourceModel struct {
	RulesetID   types.String                     `tfsdk:"ruleset_id" path:"ruleset_id"`
	AccountID   types.String                     `tfsdk:"account_id" path:"account_id"`
	ZoneID      types.String                     `tfsdk:"zone_id" path:"zone_id"`
	ID          types.String                     `tfsdk:"id" json:"id"`
	Kind        types.String                     `tfsdk:"kind" json:"kind"`
	LastUpdated types.String                     `tfsdk:"last_updated" json:"last_updated"`
	Name        types.String                     `tfsdk:"name" json:"name"`
	Phase       types.String                     `tfsdk:"phase" json:"phase"`
	Rules       *[]*RulesetRulesDataSourceModel  `tfsdk:"rules" json:"rules"`
	Version     types.String                     `tfsdk:"version" json:"version"`
	Description types.String                     `tfsdk:"description" json:"description"`
	FindOneBy   *RulesetFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type RulesetRulesDataSourceModel struct {
	LastUpdated      types.String                                 `tfsdk:"last_updated" json:"last_updated,computed"`
	Version          types.String                                 `tfsdk:"version" json:"version,computed"`
	ID               types.String                                 `tfsdk:"id" json:"id"`
	Action           types.String                                 `tfsdk:"action" json:"action"`
	ActionParameters *RulesetRulesActionParametersDataSourceModel `tfsdk:"action_parameters" json:"action_parameters"`
	Categories       types.String                                 `tfsdk:"categories" json:"categories,computed"`
	Description      types.String                                 `tfsdk:"description" json:"description,computed"`
	Enabled          types.Bool                                   `tfsdk:"enabled" json:"enabled,computed"`
	Expression       types.String                                 `tfsdk:"expression" json:"expression"`
	Logging          *RulesetRulesLoggingDataSourceModel          `tfsdk:"logging" json:"logging"`
	Ref              types.String                                 `tfsdk:"ref" json:"ref"`
}

type RulesetRulesActionParametersDataSourceModel struct {
	Response *RulesetRulesActionParametersResponseDataSourceModel `tfsdk:"response" json:"response"`
}

type RulesetRulesActionParametersResponseDataSourceModel struct {
	Content     types.String `tfsdk:"content" json:"content"`
	ContentType types.String `tfsdk:"content_type" json:"content_type"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code"`
}

type RulesetRulesLoggingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type RulesetFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
