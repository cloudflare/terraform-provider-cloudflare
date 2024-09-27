// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_network_monitoring"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicNetworkMonitoringConfigurationResultDataSourceEnvelope struct {
	Result MagicNetworkMonitoringConfigurationDataSourceModel `json:"result,computed"`
}

type MagicNetworkMonitoringConfigurationDataSourceModel struct {
	AccountID       types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	Name            types.String                                                      `tfsdk:"name" json:"name,optional"`
	RouterIPs       *[]types.String                                                   `tfsdk:"router_ips" json:"router_ips,optional"`
	WARPDevices     *[]*MagicNetworkMonitoringConfigurationWARPDevicesDataSourceModel `tfsdk:"warp_devices" json:"warp_devices,optional"`
	DefaultSampling types.Float64                                                     `tfsdk:"default_sampling" json:"default_sampling,computed_optional"`
}

func (m *MagicNetworkMonitoringConfigurationDataSourceModel) toReadParams(_ context.Context) (params magic_network_monitoring.ConfigGetParams, diags diag.Diagnostics) {
	params = magic_network_monitoring.ConfigGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicNetworkMonitoringConfigurationWARPDevicesDataSourceModel struct {
	ID       types.String `tfsdk:"id" json:"id,computed"`
	Name     types.String `tfsdk:"name" json:"name,computed"`
	RouterIP types.String `tfsdk:"router_ip" json:"router_ip,computed"`
}
