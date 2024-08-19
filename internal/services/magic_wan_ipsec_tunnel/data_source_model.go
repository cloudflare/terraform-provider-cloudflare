// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANIPSECTunnelResultDataSourceEnvelope struct {
	Result MagicWANIPSECTunnelDataSourceModel `json:"result,computed"`
}

type MagicWANIPSECTunnelDataSourceModel struct {
	AccountID     types.String                                   `tfsdk:"account_id" path:"account_id"`
	IPSECTunnelID types.String                                   `tfsdk:"ipsec_tunnel_id" path:"ipsec_tunnel_id"`
	IPSECTunnel   *MagicWANIPSECTunnelIPSECTunnelDataSourceModel `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel"`
}

type MagicWANIPSECTunnelIPSECTunnelDataSourceModel struct {
	CloudflareEndpoint types.String                                                    `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint,computed"`
	InterfaceAddress   types.String                                                    `tfsdk:"interface_address" json:"interface_address,computed"`
	Name               types.String                                                    `tfsdk:"name" json:"name,computed"`
	ID                 types.String                                                    `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                      `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          timetypes.RFC3339                                               `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                                    `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                                    `tfsdk:"description" json:"description"`
	ModifiedOn         timetypes.RFC3339                                               `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *MagicWANIPSECTunnelIPSECTunnelPSKMetadataDataSourceModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                                      `tfsdk:"replay_protection" json:"replay_protection,computed"`
	TunnelHealthCheck  *MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckDataSourceModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type MagicWANIPSECTunnelIPSECTunnelPSKMetadataDataSourceModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckDataSourceModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Rate    types.String `tfsdk:"rate" json:"rate,computed"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type,computed"`
}
