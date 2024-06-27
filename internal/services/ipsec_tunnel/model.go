// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ipsec_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IPSECTunnelResultEnvelope struct {
	Result IPSECTunnelModel `json:"result,computed"`
}

type IPSECTunnelModel struct {
	AccountID           types.String                     `tfsdk:"account_id" path:"account_id"`
	IPSECTunnelID       types.String                     `tfsdk:"ipsec_tunnel_id" path:"ipsec_tunnel_id"`
	CloudflareEndpoint  types.String                     `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress    types.String                     `tfsdk:"interface_address" json:"interface_address"`
	Name                types.String                     `tfsdk:"name" json:"name"`
	CustomerEndpoint    types.String                     `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description         types.String                     `tfsdk:"description" json:"description"`
	HealthCheck         *IPSECTunnelHealthCheckModel     `tfsdk:"health_check" json:"health_check"`
	PSK                 types.String                     `tfsdk:"psk" json:"psk"`
	ReplayProtection    types.Bool                       `tfsdk:"replay_protection" json:"replay_protection"`
	IPSECTunnels        *[]*IPSECTunnelIPSECTunnelsModel `tfsdk:"ipsec_tunnels" json:"ipsec_tunnels,computed"`
	Modified            types.Bool                       `tfsdk:"modified" json:"modified,computed"`
	ModifiedIPSECTunnel types.String                     `tfsdk:"modified_ipsec_tunnel" json:"modified_ipsec_tunnel,computed"`
	IPSECTunnel         types.String                     `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel,computed"`
	Deleted             types.Bool                       `tfsdk:"deleted" json:"deleted,computed"`
	DeletedIPSECTunnel  types.String                     `tfsdk:"deleted_ipsec_tunnel" json:"deleted_ipsec_tunnel,computed"`
}

type IPSECTunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}

type IPSECTunnelIPSECTunnelsModel struct {
	CloudflareEndpoint types.String                                   `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                   `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                   `tfsdk:"name" json:"name"`
	ID                 types.String                                   `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                     `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          types.String                                   `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                   `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                   `tfsdk:"description" json:"description"`
	ModifiedOn         types.String                                   `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *IPSECTunnelIPSECTunnelsPSKMetadataModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                     `tfsdk:"replay_protection" json:"replay_protection"`
	TunnelHealthCheck  *IPSECTunnelIPSECTunnelsTunnelHealthCheckModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type IPSECTunnelIPSECTunnelsPSKMetadataModel struct {
	LastGeneratedOn types.String `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type IPSECTunnelIPSECTunnelsTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate    types.String `tfsdk:"rate" json:"rate"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type"`
}
