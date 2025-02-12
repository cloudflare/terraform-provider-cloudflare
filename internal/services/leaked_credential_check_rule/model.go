// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_check_rule

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LeakedCredentialCheckRuleResultEnvelope struct {
	Result LeakedCredentialCheckRuleModel `json:"result"`
}

type LeakedCredentialCheckRuleModel struct {
	ID       jsontypes.Normalized `tfsdk:"id" json:"id,required"`
	ZoneID   types.String         `tfsdk:"zone_id" path:"zone_id,required"`
	Password types.String         `tfsdk:"password" json:"password,optional"`
	Username types.String         `tfsdk:"username" json:"username,optional"`
}

func (m LeakedCredentialCheckRuleModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m LeakedCredentialCheckRuleModel) MarshalJSONForUpdate(state LeakedCredentialCheckRuleModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
