package zero_trust_connectivity_settings

import "github.com/hashicorp/terraform-plugin-framework/types"

type ConnectivitySettingsModel struct {
	AccountID          types.String `tfsdk:"account_id"`
	IcmpProxyEnabled   types.Bool   `tfsdk:"icmp_proxy_enabled"`
	OfframpWARPEnabled types.Bool   `tfsdk:"offramp_warp_enabled"`
}
