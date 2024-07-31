// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRulesResultListDataSourceEnvelope struct {
	Result *[]*AccessRulesResultDataSourceModel `json:"result,computed"`
}

type AccessRulesDataSourceModel struct {
	AccountID     types.String                             `tfsdk:"account_id" path:"account_id"`
	ZoneID        types.String                             `tfsdk:"zone_id" path:"zone_id"`
	Direction     types.String                             `tfsdk:"direction" query:"direction"`
	EgsPagination *AccessRulesEgsPaginationDataSourceModel `tfsdk:"egs_pagination" query:"egs-pagination"`
	Filters       *AccessRulesFiltersDataSourceModel       `tfsdk:"filters" query:"filters"`
	Order         types.String                             `tfsdk:"order" query:"order"`
	Page          types.Float64                            `tfsdk:"page" query:"page"`
	PerPage       types.Float64                            `tfsdk:"per_page" query:"per_page"`
	MaxItems      types.Int64                              `tfsdk:"max_items"`
	Result        *[]*AccessRulesResultDataSourceModel     `tfsdk:"result"`
}

type AccessRulesEgsPaginationDataSourceModel struct {
	Json *AccessRulesEgsPaginationJsonDataSourceModel `tfsdk:"json" json:"json"`
}

type AccessRulesEgsPaginationJsonDataSourceModel struct {
	Page    types.Float64 `tfsdk:"page" json:"page"`
	PerPage types.Float64 `tfsdk:"per_page" json:"per_page"`
}

type AccessRulesFiltersDataSourceModel struct {
	ConfigurationTarget types.String `tfsdk:"configuration_target" json:"configuration.target"`
	ConfigurationValue  types.String `tfsdk:"configuration_value" json:"configuration.value"`
	Match               types.String `tfsdk:"match" json:"match"`
	Mode                types.String `tfsdk:"mode" json:"mode"`
	Notes               types.String `tfsdk:"notes" json:"notes"`
}

type AccessRulesResultDataSourceModel struct {
}
