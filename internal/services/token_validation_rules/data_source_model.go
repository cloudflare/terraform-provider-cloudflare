// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_rules

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/token_validation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TokenValidationRulesResultDataSourceEnvelope struct {
	Result TokenValidationRulesDataSourceModel `json:"result,computed"`
}

type TokenValidationRulesDataSourceModel struct {
	ID          types.String                                                          `tfsdk:"id" path:"rule_id,computed"`
	RuleID      types.String                                                          `tfsdk:"rule_id" path:"rule_id,optional"`
	ZoneID      types.String                                                          `tfsdk:"zone_id" path:"zone_id,required"`
	Action      types.String                                                          `tfsdk:"action" json:"action,computed"`
	CreatedAt   timetypes.RFC3339                                                     `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	Description types.String                                                          `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String                                                          `tfsdk:"expression" json:"expression,computed"`
	LastUpdated timetypes.RFC3339                                                     `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	Title       types.String                                                          `tfsdk:"title" json:"title,computed"`
	Selector    customfield.NestedObject[TokenValidationRulesSelectorDataSourceModel] `tfsdk:"selector" json:"selector,computed"`
	Filter      *TokenValidationRulesFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *TokenValidationRulesDataSourceModel) toReadParams(_ context.Context) (params token_validation.RuleGetParams, diags diag.Diagnostics) {
	params = token_validation.RuleGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *TokenValidationRulesDataSourceModel) toListParams(_ context.Context) (params token_validation.RuleListParams, diags diag.Diagnostics) {
	mFilterTokenConfiguration := []string{}
	if m.Filter.TokenConfiguration != nil {
		for _, item := range *m.Filter.TokenConfiguration {
			mFilterTokenConfiguration = append(mFilterTokenConfiguration, item.ValueString())
		}
	}

	params = token_validation.RuleListParams{
		ZoneID:             cloudflare.F(m.ZoneID.ValueString()),
		TokenConfiguration: cloudflare.F(mFilterTokenConfiguration),
	}

	if !m.Filter.ID.IsNull() {
		params.ID = cloudflare.F(m.Filter.ID.ValueString())
	}
	if !m.Filter.Action.IsNull() {
		params.Action = cloudflare.F(token_validation.RuleListParamsAction(m.Filter.Action.ValueString()))
	}
	if !m.Filter.Enabled.IsNull() {
		params.Enabled = cloudflare.F(m.Filter.Enabled.ValueBool())
	}
	if !m.Filter.Host.IsNull() {
		params.Host = cloudflare.F(m.Filter.Host.ValueString())
	}
	if !m.Filter.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Filter.Hostname.ValueString())
	}

	return
}

type TokenValidationRulesSelectorDataSourceModel struct {
	Exclude customfield.NestedObjectList[TokenValidationRulesSelectorExcludeDataSourceModel] `tfsdk:"exclude" json:"exclude,computed"`
	Include customfield.NestedObjectList[TokenValidationRulesSelectorIncludeDataSourceModel] `tfsdk:"include" json:"include,computed"`
}

type TokenValidationRulesSelectorExcludeDataSourceModel struct {
	OperationIDs customfield.List[types.String] `tfsdk:"operation_ids" json:"operation_ids,computed"`
}

type TokenValidationRulesSelectorIncludeDataSourceModel struct {
	Host customfield.List[types.String] `tfsdk:"host" json:"host,computed"`
}

type TokenValidationRulesFindOneByDataSourceModel struct {
	ID                 types.String    `tfsdk:"id" query:"id,optional"`
	Action             types.String    `tfsdk:"action" query:"action,optional"`
	Enabled            types.Bool      `tfsdk:"enabled" query:"enabled,optional"`
	Host               types.String    `tfsdk:"host" query:"host,optional"`
	Hostname           types.String    `tfsdk:"hostname" query:"hostname,optional"`
	TokenConfiguration *[]types.String `tfsdk:"token_configuration" query:"token_configuration,optional"`
}
