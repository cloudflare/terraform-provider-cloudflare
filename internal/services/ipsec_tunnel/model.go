// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ipsec_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type IPSECTunnelResultEnvelope struct {
	Result IPSECTunnelModel `json:"result,computed"`
}

type IPSECTunnelModel struct {
	AccountID          types.String                 `tfsdk:"account_id" path:"account_id"`
	IPSECTunnelID      types.String                 `tfsdk:"ipsec_tunnel_id" path:"ipsec_tunnel_id"`
	CloudflareEndpoint types.String                 `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                 `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                 `tfsdk:"name" json:"name"`
	CustomerEndpoint   types.String                 `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                 `tfsdk:"description" json:"description"`
	HealthCheck        *IPSECTunnelHealthCheckModel `tfsdk:"health_check" json:"health_check"`
	PSK                types.String                 `tfsdk:"psk" json:"psk"`
	ReplayProtection   types.Bool                   `tfsdk:"replay_protection" json:"replay_protection"`
}

type IPSECTunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}
