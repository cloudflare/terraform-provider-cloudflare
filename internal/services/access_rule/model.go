// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRuleResultEnvelope struct {
	Result AccessRuleModel `json:"result"`
}

type AccessRuleModel struct {
	AccountID     types.String                  `tfsdk:"account_id" path:"account_id"`
	Identifier    types.String                  `tfsdk:"identifier" path:"identifier"`
	ZoneID        types.String                  `tfsdk:"zone_id" path:"zone_id"`
	Mode          types.String                  `tfsdk:"mode" json:"mode"`
	Configuration *AccessRuleConfigurationModel `tfsdk:"configuration" json:"configuration"`
	Notes         types.String                  `tfsdk:"notes" json:"notes"`
	ID            types.String                  `tfsdk:"id" json:"id,computed"`
}

type AccessRuleConfigurationModel struct {
	Target types.String `tfsdk:"target" json:"target"`
	Value  types.String `tfsdk:"value" json:"value"`
}
