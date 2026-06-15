// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_warp_connector_config

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelWARPConnectorConfigResultDataSourceEnvelope struct {
	Result ZeroTrustTunnelWARPConnectorConfigDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelWARPConnectorConfigDataSourceModel struct {
	AccountID            types.String                                                                      `tfsdk:"account_id" path:"account_id,required"`
	TunnelID             types.String                                                                      `tfsdk:"tunnel_id" path:"tunnel_id,required"`
	ConfigurationVersion types.Int64                                                                       `tfsdk:"configuration_version" json:"configuration_version,computed"`
	CreatedAt            timetypes.RFC3339                                                                 `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	HaMode               types.String                                                                      `tfsdk:"ha_mode" json:"ha_mode,computed"`
	UpdatedAt            timetypes.RFC3339                                                                 `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
	Config               customfield.NestedObject[ZeroTrustTunnelWARPConnectorConfigConfigDataSourceModel] `tfsdk:"config" json:"config,computed"`
}

func (m *ZeroTrustTunnelWARPConnectorConfigDataSourceModel) toReadParams(_ context.Context) (params zero_trust.TunnelWARPConnectorConfigurationGetParams, diags diag.Diagnostics) {
	params = zero_trust.TunnelWARPConnectorConfigurationGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type ZeroTrustTunnelWARPConnectorConfigConfigDataSourceModel struct {
	FnrID        types.String                                                                                      `tfsdk:"fnr_id" json:"fnr_id,computed"`
	Vips         customfield.NestedObjectList[ZeroTrustTunnelWARPConnectorConfigConfigVipsDataSourceModel]         `tfsdk:"vips" json:"vips,computed"`
	VipsPrevious customfield.NestedObjectList[ZeroTrustTunnelWARPConnectorConfigConfigVipsPreviousDataSourceModel] `tfsdk:"vips_previous" json:"vips_previous,computed"`
}

type ZeroTrustTunnelWARPConnectorConfigConfigVipsDataSourceModel struct {
	Address types.String `tfsdk:"address" json:"address,computed"`
}

type ZeroTrustTunnelWARPConnectorConfigConfigVipsPreviousDataSourceModel struct {
	Address types.String `tfsdk:"address" json:"address,computed"`
}
