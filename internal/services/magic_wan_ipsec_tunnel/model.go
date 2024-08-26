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
	ReplayProtection    types.Bool                                                            `tfsdk:"replay_protection" json:"replay_protection"`
	Deleted             types.Bool                                                            `tfsdk:"deleted" json:"deleted,computed"`
	Modified            types.Bool                                                            `tfsdk:"modified" json:"modified,computed"`
	DeletedIPSECTunnel  customfield.NestedObject[MagicWANIPSECTunnelDeletedIPSECTunnelModel]  `tfsdk:"deleted_ipsec_tunnel" json:"deleted_ipsec_tunnel,computed"`
	IPSECTunnel         customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelModel]         `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel,computed"`
	IPSECTunnels        customfield.NestedObjectList[MagicWANIPSECTunnelIPSECTunnelsModel]    `tfsdk:"ipsec_tunnels" json:"ipsec_tunnels,computed"`
	ModifiedIPSECTunnel customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelModel] `tfsdk:"modified_ipsec_tunnel" json:"modified_ipsec_tunnel,computed"`
}

type MagicWANIPSECTunnelHealthCheckModel struct {
	Direction types.String `tfsdk:"direction" json:"direction"`
	Enabled   types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate      types.String `tfsdk:"rate" json:"rate"`
	Target    types.String `tfsdk:"target" json:"target"`
	Type      types.String `tfsdk:"type" json:"type"`
}

type MagicWANIPSECTunnelDeletedIPSECTunnelModel struct {
	CloudflareEndpoint types.String                                                 `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                 `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                 `tfsdk:"name" json:"name"`
	ID                 types.String                                                 `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                   `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          timetypes.RFC3339                                            `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                                 `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                                 `tfsdk:"description" json:"description"`
	ModifiedOn         timetypes.RFC3339                                            `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *MagicWANIPSECTunnelDeletedIPSECTunnelPSKMetadataModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                                   `tfsdk:"replay_protection" json:"replay_protection"`
	TunnelHealthCheck  *MagicWANIPSECTunnelDeletedIPSECTunnelTunnelHealthCheckModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type MagicWANIPSECTunnelDeletedIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type MagicWANIPSECTunnelDeletedIPSECTunnelTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate    types.String `tfsdk:"rate" json:"rate"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type"`
}

type MagicWANIPSECTunnelIPSECTunnelModel struct {
	CloudflareEndpoint types.String                                          `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                          `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                          `tfsdk:"name" json:"name"`
	ID                 types.String                                          `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                            `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          timetypes.RFC3339                                     `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                          `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                          `tfsdk:"description" json:"description"`
	ModifiedOn         timetypes.RFC3339                                     `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *MagicWANIPSECTunnelIPSECTunnelPSKMetadataModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                            `tfsdk:"replay_protection" json:"replay_protection"`
	TunnelHealthCheck  *MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type MagicWANIPSECTunnelIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate    types.String `tfsdk:"rate" json:"rate"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type"`
}

type MagicWANIPSECTunnelIPSECTunnelsModel struct {
	CloudflareEndpoint types.String                                           `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                           `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                           `tfsdk:"name" json:"name"`
	ID                 types.String                                           `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                             `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          timetypes.RFC3339                                      `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                           `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                           `tfsdk:"description" json:"description"`
	ModifiedOn         timetypes.RFC3339                                      `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *MagicWANIPSECTunnelIPSECTunnelsPSKMetadataModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                             `tfsdk:"replay_protection" json:"replay_protection"`
	TunnelHealthCheck  *MagicWANIPSECTunnelIPSECTunnelsTunnelHealthCheckModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type MagicWANIPSECTunnelIPSECTunnelsPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelsTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate    types.String `tfsdk:"rate" json:"rate"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelModel struct {
	CloudflareEndpoint types.String                                                  `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint"`
	InterfaceAddress   types.String                                                  `tfsdk:"interface_address" json:"interface_address"`
	Name               types.String                                                  `tfsdk:"name" json:"name"`
	ID                 types.String                                                  `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                    `tfsdk:"allow_null_cipher" json:"allow_null_cipher"`
	CreatedOn          timetypes.RFC3339                                             `tfsdk:"created_on" json:"created_on,computed"`
	CustomerEndpoint   types.String                                                  `tfsdk:"customer_endpoint" json:"customer_endpoint"`
	Description        types.String                                                  `tfsdk:"description" json:"description"`
	ModifiedOn         timetypes.RFC3339                                             `tfsdk:"modified_on" json:"modified_on,computed"`
	PSKMetadata        *MagicWANIPSECTunnelModifiedIPSECTunnelPSKMetadataModel       `tfsdk:"psk_metadata" json:"psk_metadata"`
	ReplayProtection   types.Bool                                                    `tfsdk:"replay_protection" json:"replay_protection"`
	TunnelHealthCheck  *MagicWANIPSECTunnelModifiedIPSECTunnelTunnelHealthCheckModel `tfsdk:"tunnel_health_check" json:"tunnel_health_check"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelTunnelHealthCheckModel struct {
	Enabled types.Bool   `tfsdk:"enabled" json:"enabled"`
	Rate    types.String `tfsdk:"rate" json:"rate"`
	Target  types.String `tfsdk:"target" json:"target"`
	Type    types.String `tfsdk:"type" json:"type"`
}
