// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gre_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GRETunnelResultEnvelope struct {
	Result GRETunnelModel `json:"result,computed"`
}

type GRETunnelModel struct {
	AccountID             types.String                 `tfsdk:"account_id" path:"account_id"`
	GRETunnelID           types.String                 `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id"`
	CloudflareGREEndpoint types.String                 `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint"`
	CustomerGREEndpoint   types.String                 `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint"`
	InterfaceAddress      types.String                 `tfsdk:"interface_address" json:"interface_address"`
	Name                  types.String                 `tfsdk:"name" json:"name"`
	Description           types.String                 `tfsdk:"description" json:"description"`
	HealthCheck           *GRETunnelHealthCheckModel   `tfsdk:"health_check" json:"health_check"`
	Mtu                   types.Int64                  `tfsdk:"mtu" json:"mtu"`
	TTL                   types.Int64                  `tfsdk:"ttl" json:"ttl"`
	GRETunnels            *[]*GRETunnelGRETunnelsModel `tfsdk:"gre_tunnels" json:"gre_tunnels,computed"`
	Modified              types.Bool                   `tfsdk:"modified" json:"modified,computed"`
	ModifiedGRETunnel     types.String                 `tfsdk:"modified_gre_tunnel" json:"modified_gre_tunnel,computed"`
	Deleted               types.Bool                   `tfsdk:"deleted" json:"deleted,computed"`
	DeletedGRETunnel      types.String                 `tfsdk:"deleted_gre_tunnel" json:"deleted_gre_tunnel,computed"`
	GRETunnel             types.String                 `tfsdk:"gre_tunnel" json:"gre_tunnel,computed"`
}

type GRETunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}

type GRETunnelGRETunnelsModel struct {
	CloudflareGREEndpoint types.String                         `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint"`
	CustomerGREEndpoint   types.String                         `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint"`
	InterfaceAddress      types.String                         `tfsdk:"interface_address" json:"interface_address"`
	Name                  types.String                         `tfsdk:"name" json:"name"`
	ID                    types.String                         `tfsdk:"id" json:"id,computed"`
	CreatedOn             types.String                         `tfsdk:"created_on" json:"created_on,computed"`
	Description           types.String                         `tfsdk:"description" json:"description"`
	HealthCheck           *GRETunnelGRETunnelsHealthCheckModel `tfsdk:"health_check" json:"health_check"`
	ModifiedOn            types.String                         `tfsdk:"modified_on" json:"modified_on,computed"`
	Mtu                   types.Int64                          `tfsdk:"mtu" json:"mtu"`
	TTL                   types.Int64                          `tfsdk:"ttl" json:"ttl"`
}

type GRETunnelGRETunnelsHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}
