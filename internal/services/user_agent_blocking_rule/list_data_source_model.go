// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type UserAgentBlockingRulesResultListDataSourceEnvelope struct {
	Result *[]*UserAgentBlockingRulesItemsDataSourceModel `json:"result,computed"`
}

type UserAgentBlockingRulesDataSourceModel struct {
	ZoneIdentifier    types.String                                   `tfsdk:"zone_identifier" path:"zone_identifier"`
	Description       types.String                                   `tfsdk:"description" query:"description"`
	DescriptionSearch types.String                                   `tfsdk:"description_search" query:"description_search"`
	Page              types.Float64                                  `tfsdk:"page" query:"page"`
	PerPage           types.Float64                                  `tfsdk:"per_page" query:"per_page"`
	UASearch          types.String                                   `tfsdk:"ua_search" query:"ua_search"`
	MaxItems          types.Int64                                    `tfsdk:"max_items"`
	Items             *[]*UserAgentBlockingRulesItemsDataSourceModel `tfsdk:"items"`
}

type UserAgentBlockingRulesItemsDataSourceModel struct {
	ID            types.String                                             `tfsdk:"id" json:"id,computed"`
	Configuration *UserAgentBlockingRulesItemsConfigurationDataSourceModel `tfsdk:"configuration" json:"configuration"`
	Description   types.String                                             `tfsdk:"description" json:"description"`
	Mode          types.String                                             `tfsdk:"mode" json:"mode"`
	Paused        types.Bool                                               `tfsdk:"paused" json:"paused"`
}

type UserAgentBlockingRulesItemsConfigurationDataSourceModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}
