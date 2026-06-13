// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship_flag

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FlagshipFlagResultEnvelope struct {
	Result FlagshipFlagModel `json:"result"`
}

type FlagshipFlagModel struct {
	AccountID        types.String               `tfsdk:"account_id" path:"account_id,required"`
	AppID            types.String               `tfsdk:"app_id" path:"app_id,required"`
	FlagKey          types.String               `tfsdk:"flag_key" path:"flag_key,optional"`
	DefaultVariation types.String               `tfsdk:"default_variation" json:"default_variation,required"`
	Enabled          types.Bool                 `tfsdk:"enabled" json:"enabled,required"`
	Key              types.String               `tfsdk:"key" json:"key,required"`
	Variations       *map[string]types.String   `tfsdk:"variations" json:"variations,required"`
	Rules            *[]*FlagshipFlagRulesModel `tfsdk:"rules" json:"rules,required"`
	Description      types.String               `tfsdk:"description" json:"description,optional"`
	Type             types.String               `tfsdk:"type" json:"type,optional"`
	UpdatedAt        types.String               `tfsdk:"updated_at" json:"updated_at,computed"`
	UpdatedBy        types.String               `tfsdk:"updated_by" json:"updated_by,computed"`
}

func (m FlagshipFlagModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m FlagshipFlagModel) MarshalJSONForUpdate(state FlagshipFlagModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type FlagshipFlagRulesModel struct {
	Conditions     *[]*FlagshipFlagRulesConditionsModel `tfsdk:"conditions" json:"conditions,required"`
	Priority       types.Int64                          `tfsdk:"priority" json:"priority,required"`
	ServeVariation types.String                         `tfsdk:"serve_variation" json:"serve_variation,required"`
	Rollout        *FlagshipFlagRulesRolloutModel       `tfsdk:"rollout" json:"rollout,optional"`
}

type FlagshipFlagRulesConditionsModel struct {
	Attribute       types.String                                `tfsdk:"attribute" json:"attribute,optional"`
	Operator        types.String                                `tfsdk:"operator" json:"operator,optional"`
	Value           jsontypes.Normalized                        `tfsdk:"value" json:"value,optional"`
	Clauses         *[]*FlagshipFlagRulesConditionsClausesModel `tfsdk:"clauses" json:"clauses,optional"`
	LogicalOperator types.String                                `tfsdk:"logical_operator" json:"logical_operator,optional"`
}

type FlagshipFlagRulesConditionsClausesModel struct {
	Attribute       types.String                                       `tfsdk:"attribute" json:"attribute,optional"`
	Operator        types.String                                       `tfsdk:"operator" json:"operator,optional"`
	Value           jsontypes.Normalized                               `tfsdk:"value" json:"value,optional"`
	Clauses         *[]*FlagshipFlagRulesConditionsClausesClausesModel `tfsdk:"clauses" json:"clauses,optional"`
	LogicalOperator types.String                                       `tfsdk:"logical_operator" json:"logical_operator,optional"`
}

type FlagshipFlagRulesConditionsClausesClausesModel struct {
	Attribute       types.String                                              `tfsdk:"attribute" json:"attribute,optional"`
	Operator        types.String                                              `tfsdk:"operator" json:"operator,optional"`
	Value           jsontypes.Normalized                                      `tfsdk:"value" json:"value,optional"`
	Clauses         *[]*FlagshipFlagRulesConditionsClausesClausesClausesModel `tfsdk:"clauses" json:"clauses,optional"`
	LogicalOperator types.String                                              `tfsdk:"logical_operator" json:"logical_operator,optional"`
}

type FlagshipFlagRulesConditionsClausesClausesClausesModel struct {
	Attribute       types.String                                                     `tfsdk:"attribute" json:"attribute,optional"`
	Operator        types.String                                                     `tfsdk:"operator" json:"operator,optional"`
	Value           jsontypes.Normalized                                             `tfsdk:"value" json:"value,optional"`
	Clauses         *[]*FlagshipFlagRulesConditionsClausesClausesClausesClausesModel `tfsdk:"clauses" json:"clauses,optional"`
	LogicalOperator types.String                                                     `tfsdk:"logical_operator" json:"logical_operator,optional"`
}

type FlagshipFlagRulesConditionsClausesClausesClausesClausesModel struct {
	Attribute       types.String                                                            `tfsdk:"attribute" json:"attribute,optional"`
	Operator        types.String                                                            `tfsdk:"operator" json:"operator,optional"`
	Value           jsontypes.Normalized                                                    `tfsdk:"value" json:"value,optional"`
	Clauses         *[]*FlagshipFlagRulesConditionsClausesClausesClausesClausesClausesModel `tfsdk:"clauses" json:"clauses,optional"`
	LogicalOperator types.String                                                            `tfsdk:"logical_operator" json:"logical_operator,optional"`
}

type FlagshipFlagRulesConditionsClausesClausesClausesClausesClausesModel struct {
	Attribute       types.String         `tfsdk:"attribute" json:"attribute,optional"`
	Operator        types.String         `tfsdk:"operator" json:"operator,optional"`
	Value           jsontypes.Normalized `tfsdk:"value" json:"value,optional"`
	Clauses         *[]types.String      `tfsdk:"clauses" json:"clauses,optional"`
	LogicalOperator types.String         `tfsdk:"logical_operator" json:"logical_operator,optional"`
}

type FlagshipFlagRulesRolloutModel struct {
	Percentage types.Float64 `tfsdk:"percentage" json:"percentage,required"`
	Attribute  types.String  `tfsdk:"attribute" json:"attribute,optional"`
}
