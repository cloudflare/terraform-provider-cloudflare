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
	Mode          types.String                             `tfsdk:"mode" query:"mode"`
	Notes         types.String                             `tfsdk:"notes" query:"notes"`
	Order         types.String                             `tfsdk:"order" query:"order"`
	Configuration *AccessRulesConfigurationDataSourceModel `tfsdk:"configuration" query:"configuration"`
	Match         types.String                             `tfsdk:"match" query:"match"`
	MaxItems      types.Int64                              `tfsdk:"max_items"`
	Result        *[]*AccessRulesResultDataSourceModel     `tfsdk:"result"`
}

type AccessRulesConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}

type AccessRulesResultDataSourceModel struct {
}
