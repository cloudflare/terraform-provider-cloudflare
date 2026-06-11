// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_warp_connector_config

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelWARPConnectorConfigResultEnvelope struct {
	Result ZeroTrustTunnelWARPConnectorConfigModel `json:"result"`
}

type ZeroTrustTunnelWARPConnectorConfigModel struct {
	ID                   types.String                                   `tfsdk:"id" json:"-,computed"`
	TunnelID             types.String                                   `tfsdk:"tunnel_id" path:"tunnel_id,required"`
	AccountID            types.String                                   `tfsdk:"account_id" path:"account_id,required"`
	HaMode               types.String                                   `tfsdk:"ha_mode" json:"ha_mode,required"`
	Config               *ZeroTrustTunnelWARPConnectorConfigConfigModel `tfsdk:"config" json:"config,optional"`
	ConfigurationVersion types.Int64                                    `tfsdk:"configuration_version" json:"configuration_version,computed"`
	CreatedAt            timetypes.RFC3339                              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	UpdatedAt            timetypes.RFC3339                              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}

func (m ZeroTrustTunnelWARPConnectorConfigModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m ZeroTrustTunnelWARPConnectorConfigModel) MarshalJSONForUpdate(state ZeroTrustTunnelWARPConnectorConfigModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type ZeroTrustTunnelWARPConnectorConfigConfigModel struct {
	FnrID        types.String                                                  `tfsdk:"fnr_id" json:"fnr_id,optional"`
	Vips         *[]*ZeroTrustTunnelWARPConnectorConfigConfigVipsModel         `tfsdk:"vips" json:"vips,optional"`
	VipsPrevious *[]*ZeroTrustTunnelWARPConnectorConfigConfigVipsPreviousModel `tfsdk:"vips_previous" json:"vips_previous,optional"`
}

type ZeroTrustTunnelWARPConnectorConfigConfigVipsModel struct {
	Address types.String `tfsdk:"address" json:"address,required"`
}

type ZeroTrustTunnelWARPConnectorConfigConfigVipsPreviousModel struct {
	Address types.String `tfsdk:"address" json:"address,required"`
}
