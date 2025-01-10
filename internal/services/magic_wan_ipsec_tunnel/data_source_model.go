// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_ipsec_tunnel

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANIPSECTunnelResultDataSourceEnvelope struct {
	Result MagicWANIPSECTunnelDataSourceModel `json:"result,computed"`
}

type MagicWANIPSECTunnelDataSourceModel struct {
	AccountID     types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	IPSECTunnelID types.String                                                            `tfsdk:"ipsec_tunnel_id" path:"ipsec_tunnel_id,required"`
	IPSECTunnel   customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelDataSourceModel] `tfsdk:"ipsec_tunnel" json:"ipsec_tunnel,computed"`
}

func (m *MagicWANIPSECTunnelDataSourceModel) toReadParams(_ context.Context) (params magic_transit.IPSECTunnelGetParams, diags diag.Diagnostics) {
	params = magic_transit.IPSECTunnelGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicWANIPSECTunnelIPSECTunnelDataSourceModel struct {
	CloudflareEndpoint types.String                                                                       `tfsdk:"cloudflare_endpoint" json:"cloudflare_endpoint,computed"`
	InterfaceAddress   types.String                                                                       `tfsdk:"interface_address" json:"interface_address,computed"`
	Name               types.String                                                                       `tfsdk:"name" json:"name,computed"`
	ID                 types.String                                                                       `tfsdk:"id" json:"id,computed"`
	AllowNullCipher    types.Bool                                                                         `tfsdk:"allow_null_cipher" json:"allow_null_cipher,computed"`
	CreatedOn          timetypes.RFC3339                                                                  `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	CustomerEndpoint   types.String                                                                       `tfsdk:"customer_endpoint" json:"customer_endpoint,computed"`
	Description        types.String                                                                       `tfsdk:"description" json:"description,computed"`
	HealthCheck        customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelHealthCheckDataSourceModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn         timetypes.RFC3339                                                                  `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	PSKMetadata        customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelPSKMetadataDataSourceModel] `tfsdk:"psk_metadata" json:"psk_metadata,computed"`
	ReplayProtection   types.Bool                                                                         `tfsdk:"replay_protection" json:"replay_protection,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelHealthCheckDataSourceModel struct {
	Direction types.String                                                                             `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool                                                                               `tfsdk:"enabled" json:"enabled,computed"`
	Rate      types.String                                                                             `tfsdk:"rate" json:"rate,computed"`
	Target    customfield.NestedObject[MagicWANIPSECTunnelIPSECTunnelHealthCheckTargetDataSourceModel] `tfsdk:"target" json:"target,computed"`
	Type      types.String                                                                             `tfsdk:"type" json:"type,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelHealthCheckTargetDataSourceModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,computed"`
}

type MagicWANIPSECTunnelIPSECTunnelPSKMetadataDataSourceModel struct {
	LastGeneratedOn timetypes.RFC3339 `tfsdk:"last_generated_on" json:"last_generated_on,computed" format:"date-time"`
}
