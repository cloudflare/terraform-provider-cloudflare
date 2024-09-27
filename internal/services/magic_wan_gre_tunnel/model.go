// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANGRETunnelResultEnvelope struct {
	Result MagicWANGRETunnelModel `json:"result"`
}

type MagicWANGRETunnelModel struct {
	AccountID             types.String                                                      `tfsdk:"account_id" path:"account_id,required"`
	GRETunnelID           types.String                                                      `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id,optional"`
	CloudflareGREEndpoint types.String                                                      `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,optional"`
	CustomerGREEndpoint   types.String                                                      `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,optional"`
	Description           types.String                                                      `tfsdk:"description" json:"description,optional"`
	InterfaceAddress      types.String                                                      `tfsdk:"interface_address" json:"interface_address,optional"`
	Name                  types.String                                                      `tfsdk:"name" json:"name,optional"`
	Mtu                   types.Int64                                                       `tfsdk:"mtu" json:"mtu,computed_optional"`
	TTL                   types.Int64                                                       `tfsdk:"ttl" json:"ttl,computed_optional"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelHealthCheckModel]       `tfsdk:"health_check" json:"health_check,computed_optional"`
	Deleted               types.Bool                                                        `tfsdk:"deleted" json:"deleted,computed"`
	Modified              types.Bool                                                        `tfsdk:"modified" json:"modified,computed"`
	DeletedGRETunnel      customfield.NestedObject[MagicWANGRETunnelDeletedGRETunnelModel]  `tfsdk:"deleted_gre_tunnel" json:"deleted_gre_tunnel,computed"`
	GRETunnel             customfield.NestedObject[MagicWANGRETunnelGRETunnelModel]         `tfsdk:"gre_tunnel" json:"gre_tunnel,computed"`
	GRETunnels            customfield.NestedObjectList[MagicWANGRETunnelGRETunnelsModel]    `tfsdk:"gre_tunnels" json:"gre_tunnels,computed"`
	ModifiedGRETunnel     customfield.NestedObject[MagicWANGRETunnelModifiedGRETunnelModel] `tfsdk:"modified_gre_tunnel" json:"modified_gre_tunnel,computed"`
}

type MagicWANGRETunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed_optional"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed_optional"`
	Target    types.String `tfsdk:"target" json:"target,optional"`
	Type      types.String `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANGRETunnelDeletedGRETunnelModel struct {
	CloudflareGREEndpoint types.String                                                                `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                                `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                                `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                                `tfsdk:"name" json:"name,computed"`
	ID                    types.String                                                                `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                                `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelDeletedGRETunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn            timetypes.RFC3339                                                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                                 `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                                 `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelDeletedGRETunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed"`
	Target    types.String `tfsdk:"target" json:"target,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type MagicWANGRETunnelGRETunnelModel struct {
	CloudflareGREEndpoint types.String                                                         `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                         `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                         `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                         `tfsdk:"name" json:"name,computed"`
	ID                    types.String                                                         `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                         `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelGRETunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn            timetypes.RFC3339                                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                          `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                          `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelGRETunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed"`
	Target    types.String `tfsdk:"target" json:"target,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type MagicWANGRETunnelGRETunnelsModel struct {
	CloudflareGREEndpoint types.String                                                          `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                          `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                          `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                          `tfsdk:"name" json:"name,computed"`
	ID                    types.String                                                          `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                                                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                          `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelGRETunnelsHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn            timetypes.RFC3339                                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                           `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                           `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelGRETunnelsHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed"`
	Target    types.String `tfsdk:"target" json:"target,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}

type MagicWANGRETunnelModifiedGRETunnelModel struct {
	CloudflareGREEndpoint types.String                                                                 `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                                 `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                                 `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                                 `tfsdk:"name" json:"name,computed"`
	ID                    types.String                                                                 `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                                                            `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                                 `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelModifiedGRETunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn            timetypes.RFC3339                                                            `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                                  `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                                  `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelModifiedGRETunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed"`
	Target    types.String `tfsdk:"target" json:"target,computed"`
	Type      types.String `tfsdk:"type" json:"type,computed"`
}
