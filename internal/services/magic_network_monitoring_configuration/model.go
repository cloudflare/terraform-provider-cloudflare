// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration

import (
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicNetworkMonitoringConfigurationResultEnvelope struct {
	Result MagicNetworkMonitoringConfigurationModel `json:"result"`
}

type MagicNetworkMonitoringConfigurationModel struct {
	AccountID       types.String                                                                      `tfsdk:"account_id" path:"account_id,required"`
	Name            types.String                                                                      `tfsdk:"name" json:"name,required"`
	RouterIPs       *[]types.String                                                                   `tfsdk:"router_ips" json:"router_ips,optional"`
	DefaultSampling types.Float64                                                                     `tfsdk:"default_sampling" json:"default_sampling,computed_optional"`
	WARPDevices     customfield.NestedObjectList[MagicNetworkMonitoringConfigurationWARPDevicesModel] `tfsdk:"warp_devices" json:"warp_devices,computed_optional"`
}

func (m MagicNetworkMonitoringConfigurationModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicNetworkMonitoringConfigurationModel) MarshalJSONForUpdate(state MagicNetworkMonitoringConfigurationModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MagicNetworkMonitoringConfigurationWARPDevicesModel struct {
	ID       types.String `tfsdk:"id" json:"id,required"`
	Name     types.String `tfsdk:"name" json:"name,required"`
	RouterIP types.String `tfsdk:"router_ip" json:"router_ip,required"`
}
