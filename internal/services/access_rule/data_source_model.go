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
	AccountID  types.String                        `tfsdk:"account_id" path:"account_id"`
	Identifier types.String                        `tfsdk:"identifier" path:"identifier"`
	ZoneID     types.String                        `tfsdk:"zone_id" path:"zone_id"`
	Filter     *AccessRuleFindOneByDataSourceModel `tfsdk:"filter"`
}

type AccessRuleFindOneByDataSourceModel struct {
	Configuration *AccessRuleConfigurationDataSourceModel `tfsdk:"configuration" query:"configuration"`
	Direction     types.String                            `tfsdk:"direction" query:"direction"`
	Match         types.String                            `tfsdk:"match" query:"match"`
	Mode          types.String                            `tfsdk:"mode" query:"mode"`
	Notes         types.String                            `tfsdk:"notes" query:"notes"`
	Order         types.String                            `tfsdk:"order" query:"order"`
}

type AccessRuleConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}
