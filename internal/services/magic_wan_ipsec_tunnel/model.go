// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANIPSECTunnelResultEnvelope struct {
	Result MagicWANIPSECTunnelModel `json:"result"`
}

type MagicWANIPSECTunnelModel struct {
	ID                  types.String                                                          `tfsdk:"id" json:"id,computed"`
	AccountID           types.String                                                          `tfsdk:"account_id" path:"account_id,required"`
	CloudflareEndpoint  types.String                                                          `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint,required,no_refresh"`
	InterfaceAddress    types.String                                                          `tfsdk:"interface_address" json:"interface_address,required,no_refresh"`
	Name                types.String                                                          `tfsdk:"name" json:"name,required,no_refresh"`
	CustomerEndpoint    types.String                                                          `tfsdk:"customer_endpoint" json:"customer_endpoint,optional,no_refresh"`
	Description         types.String                                                          `tfsdk:"description" json:"description,optional,no_refresh"`
	PSK                 types.String                                                          `tfsdk:"psk" json:"psk,optional,no_refresh"`
	ReplayProtection    types.Bool                                                            `tfsdk:"replay_protection" json:"replay_protection,computed_optional,no_refresh"`
	HealthCheck         customfield.NestedObject[MagicWANIPSECTunnelHealthCheckModel]         `tfsdk:"health_check" json:"health_check,computed_optional,no_refresh"`
	AllowNullCipher     types.Bool                                                            `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed,no_refresh"`
	CreatedOn           timetypes.RFC3339                                                     `tfsdk:"created_on" json:"created_on,computed,no_refresh" format:"date-time"`
	Modified            types.Bool                                                            `tfsdk:"modified" json:"modified,computed,no_refresh"`
	ModifiedOn          timetypes.RFC3339                                                     `tfsdk:"modified_on" json:"modified_on,computed,no_refresh" format:"date-time"`
	IPSECTunnel         customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelModel]         `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel,computed"`
	ModifiedIPSECTunnel customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelModel] `tfsdk:"modified_ipsec_tunnel" json:"modified_ipsec_tunnel,computed,no_refresh"`
	PSKMetadata         customfield.NestedObject[MagicWANIPSECTunnelPSKMetadataModel]         `tfsdk:"psk_metadata" json:"psk_metadata,computed,no_refresh"`
}

func (m MagicWANIPSECTunnelModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicWANIPSECTunnelModel) MarshalJSONForUpdate(state MagicWANIPSECTunnelModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MagicWANIPSECTunnelHealthCheckModel struct {
	Direction types.String                                                        `tfsdk:"direction" json:"direction,computed_optional"`
	Enabled   types.Bool                                                          `tfsdk:"enabled" json:"enabled,computed_optional"`
	Rate      types.String                                                        `tfsdk:"rate" json:"rate,computed_optional"`
	Target    customfield.NestedObject[MagicWANIPSECTunnelHealthCheckTargetModel] `tfsdk:"target" json:"target,computed_optional"`
	Type      types.String                                                        `tfsdk:"type" json:"type,computed_optional"`
}

type MagicWANIPSECTunnelHealthCheckTargetModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,optional"`
}

type MagicWANIPSECTunnelIPSECTunnelModel struct {
	ID                 types.String                                                             `tfsdk:"id" json:"id,computed"`
	CloudflareEndpoint types.String                                                             `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint,computed"`
	InterfaceAddress   types.String                                                             `tfsdk:"interface_address" json:"interface_address,computed"`
	Name               types.String                                                             `tfsdk:"name" json:"name,computed"`
	AllowNullCipher    types.Bool                                                               `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed"`
	CreatedOn          timetypes.RFC3339                                                        `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                             `tfsdk:"customer_endpoint" json:"customer_endpoint,computed"`
	Description        types.String                                                             `tfsdk:"description" json:"description,computed"`
	HealthCheck        customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn         timetypes.RFC3339                                                        `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelPSKMetadataModel] `tfsdk:"psk_metadata" json:"psk_metadata,computed"`
	ReplayProtection   types.Bool                                                               `tfsdk:"replay_protection" json:"replay_protection,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelHealthCheckModel struct {
	Direction types.String                                                                   `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool                                                                     `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String                                                                   `tfsdk:"rate" json:"rate,computed"`
	Target    customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelHealthCheckTargetModel] `tfsdk:"target" json:"target,computed"`
	Type      types.String                                                                   `tfsdk:"type" json:"type,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelHealthCheckTargetModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelModel struct {
	ID                 types.String                                                                     `tfsdk:"id" json:"id,computed"`
	CloudflareEndpoint types.String                                                                     `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint,computed"`
	InterfaceAddress   types.String                                                                     `tfsdk:"interface_address" json:"interface_address,computed"`
	Name               types.String                                                                     `tfsdk:"name" json:"name,computed"`
	AllowNullCipher    types.Bool                                                                       `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed"`
	CreatedOn          timetypes.RFC3339                                                                `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                                     `tfsdk:"customer_endpoint" json:"customer_endpoint,computed"`
	Description        types.String                                                                     `tfsdk:"description" json:"description,computed"`
	HealthCheck        customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelHealthCheckModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn         timetypes.RFC3339                                                                `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelPSKMetadataModel] `tfsdk:"psk_metadata" json:"psk_metadata,computed"`
	ReplayProtection   types.Bool                                                                       `tfsdk:"replay_protection" json:"replay_protection,computed"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelHealthCheckModel struct {
	Direction types.String                                                                           `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool                                                                             `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String                                                                           `tfsdk:"rate" json:"rate,computed"`
	Target    customfield.NestedObject[MagicWANIPSECTunnelModifiedIPSECTunnelHealthCheckTargetModel] `tfsdk:"target" json:"target,computed"`
	Type      types.String                                                                           `tfsdk:"type" json:"type,computed"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelHealthCheckTargetModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,computed"`
}

type MagicWANIPSECTunnelModifiedIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}

type MagicWANIPSECTunnelPSKMetadataModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}
