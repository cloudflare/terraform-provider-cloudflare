// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessRuleResultEnvelope struct {
	Result AccessRuleModel `json:"result"`
}

type AccessRuleModel struct {
	AccountID     types.String                  `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID        types.String                  `tfsdk:"zone_id" path:"zone_id,optional"`
	Mode          types.String                  `tfsdk:"mode" json:"mode,required"`
	Configuration *AccessRuleConfigurationModel `tfsdk:"configuration" json:"configuration,required"`
	Notes         types.String                  `tfsdk:"notes" json:"notes,optional"`
}

func (m AccessRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m AccessRuleModel) MarshalJSONForUpdate(state AccessRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type AccessRuleConfigurationModel struct {
	Target types.String `tfsdk:"target" json:"target,optional"`
	Value  types.String `tfsdk:"value" json:"value,optional"`
}
