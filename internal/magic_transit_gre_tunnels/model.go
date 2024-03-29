// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_gre_tunnels

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitGRETunnelsResultEnvelope struct {
	Result MagicTransitGRETunnelsModel `json:"result,computed"`
}

type MagicTransitGRETunnelsModel struct {
	AccountID        types.String                             `tfsdk:"account_id" path:"account_id"`
	TunnelIdentifier types.String                             `tfsdk:"tunnel_identifier" path:"tunnel_identifier"`
	GRETunnels       []*MagicTransitGRETunnelsGRETunnelsModel `tfsdk:"gre_tunnels" json:"gre_tunnels"`
}

type MagicTransitGRETunnelsGRETunnelsModel struct {
	CloudflareGREEndpoint types.String                                      `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint"`
	CustomerGREEndpoint   types.String                                      `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint"`
	InterfaceAddress      types.String                                      `tfsdk:"interface_address" json:"interface_address"`
	Name                  types.String                                      `tfsdk:"name" json:"name"`
	ID                    types.String                                      `tfsdk:"id" json:"id,computed"`
	CreatedOn             types.String                                      `tfsdk:"created_on" json:"created_on,computed"`
	Description           types.String                                      `tfsdk:"description" json:"description"`
	HealthCheck           *MagicTransitGRETunnelsGRETunnelsHealthCheckModel `tfsdk:"health_check" json:"health_check"`
	ModifiedOn            types.String                                      `tfsdk:"modified_on" json:"modified_on,computed"`
	Mtu                   types.Int64                                       `tfsdk:"mtu" json:"mtu"`
	TTL                   types.Int64                                       `tfsdk:"ttl" json:"ttl"`
}

type MagicTransitGRETunnelsGRETunnelsHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}
