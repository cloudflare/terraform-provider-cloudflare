// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANGRETunnelResultEnvelope struct {
	Result MagicWANGRETunnelModel `json:"result,computed"`
}

type MagicWANGRETunnelModel struct {
	AccountID             types.String                         `tfsdk:"account_id" path:"account_id"`
	GRETunnelID           types.String                         `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id"`
	CloudflareGREEndpoint types.String                         `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint"`
	CustomerGREEndpoint   types.String                         `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint"`
	Description           types.String                         `tfsdk:"description" json:"description"`
	InterfaceAddress      types.String                         `tfsdk:"interface_address" json:"interface_address"`
	Name                  types.String                         `tfsdk:"name" json:"name"`
	HealthCheck           *MagicWANGRETunnelHealthCheckModel   `tfsdk:"health_check" json:"health_check"`
	Mtu                   types.Int64                          `tfsdk:"mtu" json:"mtu"`
	TTL                   types.Int64                          `tfsdk:"ttl" json:"ttl"`
	Deleted               types.Bool                           `tfsdk:"deleted" json:"deleted,computed"`
	Modified              types.Bool                           `tfsdk:"modified" json:"modified,computed"`
	GRETunnels            *[]*MagicWANGRETunnelGRETunnelsModel `tfsdk:"gre_tunnels" json:"gre_tunnels,computed"`
	DeletedGRETunnel      jsontypes.Normalized                 `tfsdk:"deleted_gre_tunnel" json:"deleted_gre_tunnel,computed"`
	GRETunnel             jsontypes.Normalized                 `tfsdk:"gre_tunnel" json:"gre_tunnel,computed"`
	ModifiedGRETunnel     jsontypes.Normalized                 `tfsdk:"modified_gre_tunnel" json:"modified_gre_tunnel,computed"`
}

type MagicWANGRETunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}

type MagicWANGRETunnelGRETunnelsModel struct {
	CloudflareGREEndpoint types.String                                 `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint"`
	CustomerGREEndpoint   types.String                                 `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint"`
	InterfaceAddress      types.String                                 `tfsdk:"interface_address" json:"interface_address"`
	Name                  types.String                                 `tfsdk:"name" json:"name"`
	ID                    types.String                                 `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                            `tfsdk:"created_on" json:"created_on,computed"`
	Description           types.String                                 `tfsdk:"description" json:"description"`
	HealthCheck           *MagicWANGRETunnelGRETunnelsHealthCheckModel `tfsdk:"health_check" json:"health_check"`
	ModifiedOn            timetypes.RFC3339                            `tfsdk:"modified_on" json:"modified_on,computed"`
	Mtu                   types.Int64                                  `tfsdk:"mtu" json:"mtu"`
	TTL                   types.Int64                                  `tfsdk:"ttl" json:"ttl"`
}

type MagicWANGRETunnelGRETunnelsHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}
