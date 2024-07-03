// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRuleResultDataSourceEnvelope struct {
	Result AccessRuleDataSourceModel `json:"result,computed"`
}

type AccessRuleResultListDataSourceEnvelope struct {
	Result *[]*AccessRuleDataSourceModel `json:"result,computed"`
}

type AccessRuleDataSourceModel struct {
	Identifier types.String                        `tfsdk:"identifier" path:"identifier"`
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id"`
	ZoneID     types.String                        `tfsdk:"zone_id" path:"zone_id"`
	FindOneBy  *AccessRuleFindOneByDataSourceModel `tfsdk:"find_one_by"`
}

type AccessRuleFindOneByDataSourceModel struct {
	AccountID     types.String                                     `tfsdk:"account_id" path:"account_id"`
	ZoneID        types.String                                     `tfsdk:"zone_id" path:"zone_id"`
	Direction     types.String                                     `tfsdk:"direction" query:"direction"`
	EgsPagination *AccessRuleFindOneByEgsPaginationDataSourceModel `tfsdk:"egs_pagination" query:"egs-pagination"`
	Filters       *AccessRuleFindOneByFiltersDataSourceModel       `tfsdk:"filters" query:"filters"`
	Order         types.String                                     `tfsdk:"order" query:"order"`
	Page          types.Float64                                    `tfsdk:"page" query:"page"`
	PerPage       types.Float64                                    `tfsdk:"per_page" query:"per_page"`
}

type AccessRuleFindOneByEgsPaginationDataSourceModel struct {
	Json *AccessRuleFindOneByEgsPaginationJsonDataSourceModel `tfsdk:"json" json:"json"`
}

type AccessRuleFindOneByEgsPaginationJsonDataSourceModel struct {
	Page    types.Float64 `tfsdk:"page" json:"page"`
	PerPage types.Float64 `tfsdk:"per_page" json:"per_page"`
}

type AccessRuleFindOneByFiltersDataSourceModel struct {
	ConfigurationTarget types.String `tfsdk:"configuration_target" json:"configuration.target"`
	ConfigurationValue  types.String `tfsdk:"configuration_value" json:"configuration.value"`
	Match               types.String `tfsdk:"match" json:"match"`
	Mode                types.String `tfsdk:"mode" json:"mode"`
	Notes               types.String `tfsdk:"notes" json:"notes"`
}
