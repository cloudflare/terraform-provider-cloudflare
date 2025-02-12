// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing_settings

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type EmailRoutingSettingsResultEnvelope struct {
	Result EmailRoutingSettingsModel `json:"result"`
}

type EmailRoutingSettingsModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
}

func (m EmailRoutingSettingsModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m EmailRoutingSettingsModel) MarshalJSONForUpdate(state EmailRoutingSettingsModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}
