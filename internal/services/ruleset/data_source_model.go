// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ruleset

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/rulesets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type RulesetResultDataSourceEnvelope struct {
	Result RulesetDataSourceModel `json:"result,computed"`
}

type RulesetResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[RulesetDataSourceModel] `json:"result,computed"`
}

type RulesetDataSourceModel struct {
	AccountID   types.String                     `tfsdk:"account_id" path:"account_id"`
	RulesetID   types.String                     `tfsdk:"ruleset_id" path:"ruleset_id"`
	ZoneID      types.String                     `tfsdk:"zone_id" path:"zone_id"`
	Rules       *[]*RulesetRulesDataSourceModel  `tfsdk:"rules" json:"rules"`
	Description types.String                     `tfsdk:"description" json:"description,computed"`
	ID          types.String                     `tfsdk:"id" json:"id,computed"`
	Kind        types.String                     `tfsdk:"kind" json:"kind,computed"`
	LastUpdated timetypes.RFC3339                `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Name        types.String                     `tfsdk:"name" json:"name,computed"`
	Phase       types.String                     `tfsdk:"phase" json:"phase,computed"`
	Version     types.String                     `tfsdk:"version" json:"version,computed"`
	Filter      *RulesetFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *RulesetDataSourceModel) toReadParams(_ context.Context) (params rulesets.RulesetGetParams, diags diag.Diagnostics) {
	params = rulesets.RulesetGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *RulesetDataSourceModel) toListParams(_ context.Context) (params rulesets.RulesetListParams, diags diag.Diagnostics) {
	params = rulesets.RulesetListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type RulesetRulesDataSourceModel struct {
	LastUpdated      timetypes.RFC3339                                                     `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Version          types.String                                                          `tfsdk:"version" json:"version,computed"`
	ID               types.String                                                          `tfsdk:"id" json:"id,computed"`
	Action           types.String                                                          `tfsdk:"action" json:"action,computed"`
	ActionParameters customfield.NestedObject[RulesetRulesActionParametersDataSourceModel] `tfsdk:"action_parameters" json:"action_parameters,computed"`
	Categories       customfield.List[types.String]                                        `tfsdk:"categories" json:"categories,computed"`
	Description      types.String                                                          `tfsdk:"description" json:"description,computed"`
	Enabled          types.Bool                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Expression       types.String                                                          `tfsdk:"expression" json:"expression,computed"`
	Logging          customfield.NestedObject[RulesetRulesLoggingDataSourceModel]          `tfsdk:"logging" json:"logging,computed"`
	Ref              types.String                                                          `tfsdk:"ref" json:"ref,computed"`
}

type RulesetRulesActionParametersDataSourceModel struct {
	Response customfield.NestedObject[RulesetRulesActionParametersResponseDataSourceModel] `tfsdk:"response" json:"response,computed"`
}

type RulesetRulesActionParametersResponseDataSourceModel struct {
	Content     types.String `tfsdk:"content" json:"content,computed"`
	ContentType types.String `tfsdk:"content_type" json:"content_type,computed"`
	StatusCode  types.Int64  `tfsdk:"status_code" json:"status_code,computed"`
}

type RulesetRulesLoggingDataSourceModel struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled,computed"`
}

type RulesetFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
