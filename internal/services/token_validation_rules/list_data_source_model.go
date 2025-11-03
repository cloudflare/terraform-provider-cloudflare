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

type TokenValidationRulesListResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[TokenValidationRulesListResultDataSourceModel] `json:"result,computed"`
}

type TokenValidationRulesListDataSourceModel struct {
	ZoneID             types.String                                                                `tfsdk:"zone_id" path:"zone_id,required"`
	Action             types.String                                                                `tfsdk:"action" query:"action,optional"`
	Enabled            types.Bool                                                                  `tfsdk:"enabled" query:"enabled,optional"`
	Host               types.String                                                                `tfsdk:"host" query:"host,optional"`
	Hostname           types.String                                                                `tfsdk:"hostname" query:"hostname,optional"`
	ID                 types.String                                                                `tfsdk:"id" query:"id,optional"`
	RuleID             types.String                                                                `tfsdk:"rule_id" query:"rule_id,optional"`
	TokenConfiguration *[]types.String                                                             `tfsdk:"token_configuration" query:"token_configuration,optional"`
	MaxItems           types.Int64                                                                 `tfsdk:"max_items"`
	Result             customfield.NestedObjectList[TokenValidationRulesListResultDataSourceModel] `tfsdk:"result"`
}

func (m *TokenValidationRulesListDataSourceModel) toListParams(_ context.Context) (params token_validation.RuleListParams, diags diag.Diagnostics) {
	mTokenConfiguration := []string{}
	if m.TokenConfiguration != nil {
		for _, item := range *m.TokenConfiguration {
			mTokenConfiguration = append(mTokenConfiguration, item.ValueString())
		}
	}

	params = token_validation.RuleListParams{
		ZoneID:             cloudflare.F(m.ZoneID.ValueString()),
		TokenConfiguration: cloudflare.F(mTokenConfiguration),
	}

	if !m.ID.IsNull() {
		params.ID = cloudflare.F(m.ID.ValueString())
	}
	if !m.Action.IsNull() {
		params.Action = cloudflare.F(token_validation.RuleListParamsAction(m.Action.ValueString()))
	}
	if !m.Enabled.IsNull() {
		params.Enabled = cloudflare.F(m.Enabled.ValueBool())
	}
	if !m.Host.IsNull() {
		params.Host = cloudflare.F(m.Host.ValueString())
	}
	if !m.Hostname.IsNull() {
		params.Hostname = cloudflare.F(m.Hostname.ValueString())
	}
	if !m.RuleID.IsNull() {
		params.RuleID = cloudflare.F(m.RuleID.ValueString())
	}

	return
}

type TokenValidationRulesListResultDataSourceModel struct {
	Action      types.String                                                              `tfsdk:"action" json:"action,computed"`
	Description types.String                                                              `tfsdk:"description" json:"description,computed"`
	Enabled     types.Bool                                                                `tfsdk:"enabled" json:"enabled,computed"`
	Expression  types.String                                                              `tfsdk:"expression" json:"expression,computed"`
	Selector    customfield.NestedObject[TokenValidationRulesListSelectorDataSourceModel] `tfsdk:"selector" json:"selector,computed"`
	Title       types.String                                                              `tfsdk:"title" json:"title,computed"`
	ID          types.String                                                              `tfsdk:"id" json:"id,computed"`
	CreatedAt   timetypes.RFC3339                                                         `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	LastUpdated timetypes.RFC3339                                                         `tfsdk:"last_updated" json:"last_updated,computed" format:"date-time"`
	ModifiedBy  types.String                                                              `tfsdk:"modified_by" json:"modified_by,computed"`
}

type TokenValidationRulesListSelectorDataSourceModel struct {
	Exclude customfield.NestedObjectList[TokenValidationRulesListSelectorExcludeDataSourceModel] `tfsdk:"exclude" json:"exclude,computed"`
	Include customfield.NestedObjectList[TokenValidationRulesListSelectorIncludeDataSourceModel] `tfsdk:"include" json:"include,computed"`
}

type TokenValidationRulesListSelectorExcludeDataSourceModel struct {
	OperationIDs customfield.List[types.String] `tfsdk:"operation_ids" json:"operation_ids,computed"`
}

type TokenValidationRulesListSelectorIncludeDataSourceModel struct {
	Host customfield.List[types.String] `tfsdk:"host" json:"host,computed"`
}
