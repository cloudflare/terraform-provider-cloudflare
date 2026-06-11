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

type FlagshipFlagResultDataSourceEnvelope struct {
	Result FlagshipFlagDataSourceModel `json:"result,computed"`
}

type FlagshipFlagDataSourceModel struct {
	AccountID        types.String                                                   `tfsdk:"account_id" path:"account_id,required"`
	AppID            types.String                                                   `tfsdk:"app_id" path:"app_id,required"`
	FlagKey          types.String                                                   `tfsdk:"flag_key" path:"flag_key,required"`
	DefaultVariation types.String                                                   `tfsdk:"default_variation" json:"default_variation,computed"`
	Description      types.String                                                   `tfsdk:"description" json:"description,computed"`
	Enabled          types.Bool                                                     `tfsdk:"enabled" json:"enabled,computed"`
	Key              types.String                                                   `tfsdk:"key" json:"key,computed"`
	Type             types.String                                                   `tfsdk:"type" json:"type,computed"`
	UpdatedAt        types.String                                                   `tfsdk:"updated_at" json:"updated_at,computed"`
	UpdatedBy        types.String                                                   `tfsdk:"updated_by" json:"updated_by,computed"`
	Variations       customfield.Map[types.String]                                  `tfsdk:"variations" json:"variations,computed"`
	Rules            customfield.NestedObjectList[FlagshipFlagRulesDataSourceModel] `tfsdk:"rules" json:"rules,computed"`
}

func (m *FlagshipFlagDataSourceModel) toReadParams(_ context.Context) (params flagship.AppFlagGetParams, diags diag.Diagnostics) {
	params = flagship.AppFlagGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type FlagshipFlagRulesDataSourceModel struct {
	Conditions     customfield.NestedObjectList[FlagshipFlagRulesConditionsDataSourceModel] `tfsdk:"conditions" json:"conditions,computed"`
	Priority       types.Int64                                                              `tfsdk:"priority" json:"priority,computed"`
	ServeVariation types.String                                                             `tfsdk:"serve_variation" json:"serve_variation,computed"`
	Rollout        customfield.NestedObject[FlagshipFlagRulesRolloutDataSourceModel]        `tfsdk:"rollout" json:"rollout,computed"`
}

type FlagshipFlagRulesConditionsDataSourceModel struct {
	Attribute       types.String                                                                    `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                    `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                            `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagRulesConditionsClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                    `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagRulesConditionsClausesDataSourceModel struct {
	Attribute       types.String                                                                           `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                           `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                   `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagRulesConditionsClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                           `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagRulesConditionsClausesClausesDataSourceModel struct {
	Attribute       types.String                                                                                  `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                                  `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                          `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagRulesConditionsClausesClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                                  `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagRulesConditionsClausesClausesClausesDataSourceModel struct {
	Attribute       types.String                                                                                         `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                                         `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                                 `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagRulesConditionsClausesClausesClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                                         `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagRulesConditionsClausesClausesClausesClausesDataSourceModel struct {
	Attribute       types.String                                                                                                `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                                                                                                `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized                                                                                        `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.NestedObjectList[FlagshipFlagRulesConditionsClausesClausesClausesClausesClausesDataSourceModel] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                                                                                                `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagRulesConditionsClausesClausesClausesClausesClausesDataSourceModel struct {
	Attribute       types.String                   `tfsdk:"attribute" json:"attribute,computed"`
	Operator        types.String                   `tfsdk:"operator" json:"operator,computed"`
	Value           jsontypes.Normalized           `tfsdk:"value" json:"value,computed"`
	Clauses         customfield.List[types.String] `tfsdk:"clauses" json:"clauses,computed"`
	LogicalOperator types.String                   `tfsdk:"logical_operator" json:"logical_operator,computed"`
}

type FlagshipFlagRulesRolloutDataSourceModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,computed"`
	Attribute  types.String  `tfsdk:"attribute" json:"attribute,computed"`
}
