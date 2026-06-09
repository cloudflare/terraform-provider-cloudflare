// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_flag

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/flagship"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FlagshipFlagsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[FlagshipFlagsResultDataSourceModel] `json:"result,computed"`
}

type FlagshipFlagsDataSourceModel struct {
	AccountID types.String                                                     `tfsdk:"account_id" path:"account_id,required"`
	AppID     types.String                                                     `tfsdk:"app_id" path:"app_id,required"`
	Limit     types.String                                                     `tfsdk:"limit" query:"limit,optional"`
	MaxItems  types.Int64                                                      `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[FlagshipFlagsResultDataSourceModel] `tfsdk:"result"`
}

func (m *FlagshipFlagsDataSourceModel) toListParams(_ context.Context) (params flagship.AppFlagListParams, diags diag.Diagnostics) {
	params = flagship.AppFlagListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	if !m.Limit.IsNull() {
		params.Limit = cloudflare.F(m.Limit.ValueString())
	}

	return
}

type FlagshipFlagsResultDataSourceModel struct {
	DefaultVariation types.String                                                    `tfsdk:"default_variation" json:"default_variation,computed"`
	Enabled          types.Bool                                                      `tfsdk:"enabled" json:"enabled,computed"`
	Key              types.String                                                    `tfsdk:"key" json:"key,computed"`
	Rules            customfield.NestedObjectList[FlagshipFlagsRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
	Variations       customfield.Map[types.String]                                   `tfsdk:"variations" json:"variations,computed"`
	Description      types.String                                                    `tfsdk:"description" json:"description,computed"`
	Type             types.String                                                    `tfsdk:"type" json:"type,computed"`
	UpdatedAt        types.String                                                    `tfsdk:"updated_at" json:"updated_at,computed"`
	UpdatedBy        types.String                                                    `tfsdk:"updated_by" json:"updated_by,computed"`
}

type FlagshipFlagsRulesDataSourceModel struct {
	Conditions     customfield.NestedObjectList[FlagshipFlagsRulesConditionsDataSourceModel] `tfsdk:"conditions" json:"conditions,computed"`
	Priority       types.Int64                                                               `tfsdk:"priority" json:"priority,computed"`
	ServeVariation types.String                                                              `tfsdk:"serve_variation" json:"serve_variation,computed"`
	Rollout        customfield.NestedObject[FlagshipFlagsRulesRolloutDataSourceModel]        `tfsdk:"rollout" json:"rollout,computed"`
}

type FlagshipFlagsRulesConditionsDataSourceModel struct {
	Attribute       types.String                                                                     `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                     `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                             `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagsRulesConditionsClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                     `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagsRulesConditionsClausesDataSourceModel struct {
	Attribute       types.String                                                                            `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                            `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                    `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagsRulesConditionsClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                            `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagsRulesConditionsClausesClausesDataSourceModel struct {
	Attribute       types.String                                                                                   `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                                   `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                           `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagsRulesConditionsClausesClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                                   `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagsRulesConditionsClausesClausesClausesDataSourceModel struct {
	Attribute       types.String                                                                                          `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                                          `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                                  `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagsRulesConditionsClausesClausesClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                                          `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagsRulesConditionsClausesClausesClausesClausesDataSourceModel struct {
	Attribute       types.String                                                                                                 `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                                                 `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                                         `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagsRulesConditionsClausesClausesClausesClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                                                 `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagsRulesConditionsClausesClausesClausesClausesClausesDataSourceModel struct {
	Attribute       types.String                   `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                   `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized           `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.List[types.String] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                   `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagsRulesRolloutDataSourceModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
	Attribute  types.String  `tfsdk:"attribute" json:"attribute,computed"`
}
