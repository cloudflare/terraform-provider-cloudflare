// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANIPSECTunnelResultEnvelope struct {
	Result MagicWANIPSECTunnelModel `json:"result"`
}

type MagicWANIPSECTunnelModel struct {
	AccountID           types.String                                                          `tfsdk:"account_id" path:"account_id"`
	IPSECTunnelID       types.String                                                          `tfsdk:"ipsec_tunnel_id" path:"ipsec_tunnel_id"`
	CloudflareEndpoint  types.String                                                          `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress    types.String                                                          `tfsdk:"interface_address" json:"interface_address"`
	Name                types.String                                                          `tfsdk:"name" json:"name"`
	CustomerEndpoint    types.String                                                          `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description         types.String                                                          `tfsdk:"description" json:"description"`
	PSK                 types.String                                                          `tfsdk:"psk" json:"psk"`
	HealthCheck         *MagicWANIPSECTunnelHealthCheckModel                                  `tfsdk:"health_check" json:"health_check"`
	ReplayProtection    types.Bool                                                            `tfsdk:"replay_protection" json:"replay_protection,computed_optional"`
	Deleted             types.Bool                                                            `tfsdk:"deleted" json:"deleted,computed"`
	Modified            types.Bool                                                            `tfsdk:"modified" json:"modified,computed"`
	DeletedIPSECTunnel  customfield.NestedObject[MagicWANIPSECTunnelDeletedIPSECTunnelModel]  `tfsdk:"deleted_ipsec_tunnel" json:"deleted_ipsec_tunnel,computed"`
	IPSECTunnel         customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelModel]         `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel,computed"`
	IPSECTunnels        customfield.NestedObjectList[MagicWANIPSECTunnelIPSECTunnelsModel]    `tfsdk:"ipsec_tunnels" json:"ipsec_tunnels,computed"`
	ModifiedIPSECTunnel customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelModel] `tfsdk:"modified_ipsec_tunnel" json:"modified_ipsec_tunnel,computed"`
}

type MagicWANIPSECTunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction,computed_optional"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate      types.String `tfsdk:"rate" json:"rate,computed_optional"`
	Target    types.String `tfsdk:"target" json:"target,computed_optional"`
	Type      types.String `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANIPSECTunnelDeletedIPSECTunnelModel struct {
	CloudflareEndpoint types.String                                                                          `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                                          `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                                          `tfsdk:"name" json:"name"`
	ID                 types.String                                                                          `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                                            `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed_optional"`
	CreatedOn          timetypes.RFC3339                                                                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                                          `tfsdk:"customer_endpoint" json:"customer_endpoint,computed_optional"`
	Description        types.String                                                                          `tfsdk:"description" json:"description,computed_optional"`
	ModifiedOn         timetypes.RFC3339                                                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelDeletedIPSECTunnelPSKMetadataModel]       `tfsdk:"psk_metadata" json:"psk_metadata,computed_optional"`
	ReplayProtection   types.Bool                                                                            `tfsdk:"replay_protection" json:"replay_protection,computed_optional"`
	TunnelHealthCheck  customfield.NestedObject[MagicWANIPSECTunnelDeletedIPSECTunnelTunnelHealthCheckModel] `tfsdk:"tunnel_health_check" json:"tunnel_health_check,computed_optional"`
}

type MagicWANIPSECTunnelDeletedIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}

type MagicWANIPSECTunnelDeletedIPSECTunnelTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate    types.String `tfsdk:"rate" json:"rate,computed_optional"`
	Target  types.String `tfsdk:"target" json:"target,computed_optional"`
	Type    types.String `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANIPSECTunnelIPSECTunnelModel struct {
	CloudflareEndpoint types.String                                                                   `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                                   `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                                   `tfsdk:"name" json:"name"`
	ID                 types.String                                                                   `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                                     `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed_optional"`
	CreatedOn          timetypes.RFC3339                                                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                                   `tfsdk:"customer_endpoint" json:"customer_endpoint,computed_optional"`
	Description        types.String                                                                   `tfsdk:"description" json:"description,computed_optional"`
	ModifiedOn         timetypes.RFC3339                                                              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelPSKMetadataModel]       `tfsdk:"psk_metadata" json:"psk_metadata,computed_optional"`
	ReplayProtection   types.Bool                                                                     `tfsdk:"replay_protection" json:"replay_protection,computed_optional"`
	TunnelHealthCheck  customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckModel] `tfsdk:"tunnel_health_check" json:"tunnel_health_check,computed_optional"`
}

type MagicWANIPSECTunnelIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}

type MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate    types.String `tfsdk:"rate" json:"rate,computed_optional"`
	Target  types.String `tfsdk:"target" json:"target,computed_optional"`
	Type    types.String `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANIPSECTunnelIPSECTunnelsModel struct {
	CloudflareEndpoint types.String                                                                    `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                                    `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                                    `tfsdk:"name" json:"name"`
	ID                 types.String                                                                    `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                                      `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed_optional"`
	CreatedOn          timetypes.RFC3339                                                               `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                                    `tfsdk:"customer_endpoint" json:"customer_endpoint,computed_optional"`
	Description        types.String                                                                    `tfsdk:"description" json:"description,computed_optional"`
	ModifiedOn         timetypes.RFC3339                                                               `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelsPSKMetadataModel]       `tfsdk:"psk_metadata" json:"psk_metadata,computed_optional"`
	ReplayProtection   types.Bool                                                                      `tfsdk:"replay_protection" json:"replay_protection,computed_optional"`
	TunnelHealthCheck  customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelsTunnelHealthCheckModel] `tfsdk:"tunnel_health_check" json:"tunnel_health_check,computed_optional"`
}

type MagicWANIPSECTunnelIPSECTunnelsPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}

type MagicWANIPSECTunnelIPSECTunnelsTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate    types.String `tfsdk:"rate" json:"rate,computed_optional"`
	Target  types.String `tfsdk:"target" json:"target,computed_optional"`
	Type    types.String `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelModel struct {
	CloudflareEndpoint types.String                                                                           `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                                           `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                                           `tfsdk:"name" json:"name"`
	ID                 types.String                                                                           `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                                             `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed_optional"`
	CreatedOn          timetypes.RFC3339                                                                      `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                                           `tfsdk:"customer_endpoint" json:"customer_endpoint,computed_optional"`
	Description        types.String                                                                           `tfsdk:"description" json:"description,computed_optional"`
	ModifiedOn         timetypes.RFC3339                                                                      `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelPSKMetadataModel]       `tfsdk:"psk_metadata" json:"psk_metadata,computed_optional"`
	ReplayProtection   types.Bool                                                                             `tfsdk:"replay_protection" json:"replay_protection,computed_optional"`
	TunnelHealthCheck  customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelTunnelHealthCheckModel] `tfsdk:"tunnel_health_check" json:"tunnel_health_check,computed_optional"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate    types.String `tfsdk:"rate" json:"rate,computed_optional"`
	Target  types.String `tfsdk:"target" json:"target,computed_optional"`
	Type    types.String `tfsdk:"type" json:"type,computed_optional"`
}
