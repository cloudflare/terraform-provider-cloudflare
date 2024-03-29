// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_ipsec_tunnels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitIPSECTunnelsResultEnvelope struct {
	Result MagicTransitIPSECTunnelsModel `json:"result,computed"`
}

type MagicTransitIPSECTunnelsModel struct {
	AccountID          types.String                                 `tfsdk:"account_id" path:"account_id"`
	TunnelIdentifier   types.String                                 `tfsdk:"tunnel_identifier" path:"tunnel_identifier"`
	CloudflareEndpoint types.String                                 `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                 `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                 `tfsdk:"name" json:"name"`
	CustomerEndpoint   types.String                                 `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                 `tfsdk:"description" json:"description"`
	HealthCheck        *MagicTransitIPSECTunnelsHealthCheckModel    `tfsdk:"health_check" json:"health_check"`
	PSK                types.String                                 `tfsdk:"psk" json:"psk"`
	ReplayProtection   types.Bool                                   `tfsdk:"replay_protection" json:"replay_protection"`
	IPSECTunnels       []*MagicTransitIPSECTunnelsIPSECTunnelsModel `tfsdk:"ipsec_tunnels" json:"ipsec_tunnels"`
}

type MagicTransitIPSECTunnelsHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}

type MagicTransitIPSECTunnelsIPSECTunnelsModel struct {
	CloudflareEndpoint types.String                                                `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                `tfsdk:"name" json:"name"`
	ID                 types.String                                                `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                  `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          types.String                                                `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                                `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                                `tfsdk:"description" json:"description"`
	ModifiedOn         types.String                                                `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *MagicTransitIPSECTunnelsIPSECTunnelsPSKMetadataModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                                  `tfsdk:"replay_protection" json:"replay_protection"`
	TunnelHealthCheck  *MagicTransitIPSECTunnelsIPSECTunnelsTunnelHealthCheckModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type MagicTransitIPSECTunnelsIPSECTunnelsPSKMetadataModel struct {
	LastGeneratedOn types.String `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type MagicTransitIPSECTunnelsIPSECTunnelsTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate    types.String `tfsdk:"rate" json:"rate"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type"`
}
