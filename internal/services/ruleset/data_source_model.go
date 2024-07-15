// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultDataSourceEnvelope struct {
	Result RulesetDataSourceModel `json:"result,computed"`
}

type RulesetResultListDataSourceEnvelope struct {
	Result *[]*RulesetDataSourceModel `json:"result,computed"`
}

type RulesetDataSourceModel struct {
	RulesetID   types.String                    `tfsdk:"ruleset_id" path:"ruleset_id"`
	AccountID   types.String                    `tfsdk:"account_id" path:"account_id"`
	ZoneID      types.String                    `tfsdk:"zone_id" path:"zone_id"`
	ID          types.String                    `tfsdk:"id" json:"id,computed"`
	Kind        types.String                    `tfsdk:"kind" json:"kind,computed"`
	LastUpdated timetypes.RFC3339               `tfsdk:"last_updated" json:"last_updated,computed"`
	Name        types.String                    `tfsdk:"name" json:"name,computed"`
	Phase       types.String                    `tfsdk:"phase" json:"phase,computed"`
	Rules       *[]*RulesetRulesDataSourceModel `tfsdk:"rules" json:"rules"`
	Version     types.String                    `tfsdk:"version" json:"version,computed"`
	Description types.String                    `tfsdk:"description" json:"description,computed"`
}

type RulesetRulesDataSourceModel struct {
	LastUpdated      timetypes.RFC3339                            `tfsdk:"last_updated" json:"last_updated,computed"`
	Version          types.String                                 `tfsdk:"version" json:"version,computed"`
	ID               types.String                                 `tfsdk:"id" json:"id"`
	Action           types.String                                 `tfsdk:"action" json:"action"`
	ActionParameters *RulesetRulesActionParametersDataSourceModel `tfsdk:"action_parameters" json:"action_parameters"`
	Categories       *[]types.String                              `tfsdk:"categories" json:"categories,computed"`
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
	Content     types.String `tfsdk:"content" json:"content,computed"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
}

type RulesetRulesLoggingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}
