// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_wan_gre_tunnel

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicWANGRETunnelResultDataSourceEnvelope struct {
	Result MagicWANGRETunnelDataSourceModel `json:"result,computed"`
}

type MagicWANGRETunnelDataSourceModel struct {
	AccountID   types.String                                                        `tfsdk:"account_id" path:"account_id,required"`
	GRETunnelID types.String                                                        `tfsdk:"gre_tunnel_id" path:"gre_tunnel_id,required"`
	GRETunnel   customfield.NestedObject[MagicWANGRETunnelGRETunnelDataSourceModel] `tfsdk:"gre_tunnel" json:"gre_tunnel,computed"`
}

func (m *MagicWANGRETunnelDataSourceModel) toReadParams(_ context.Context) (params magic_transit.GRETunnelGetParams, diags diag.Diagnostics) {
	params = magic_transit.GRETunnelGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicWANGRETunnelGRETunnelDataSourceModel struct {
	CloudflareGREEndpoint types.String                                                                   `tfsdk:"cloudflare_gre_endpoint" json:"cloudflare_gre_endpoint,computed"`
	CustomerGREEndpoint   types.String                                                                   `tfsdk:"customer_gre_endpoint" json:"customer_gre_endpoint,computed"`
	InterfaceAddress      types.String                                                                   `tfsdk:"interface_address" json:"interface_address,computed"`
	Name                  types.String                                                                   `tfsdk:"name" json:"name,computed"`
	ID                    types.String                                                                   `tfsdk:"id" json:"id,computed"`
	CreatedOn             timetypes.RFC3339                                                              `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description           types.String                                                                   `tfsdk:"description" json:"description,computed"`
	HealthCheck           customfield.NestedObject[MagicWANGRETunnelGRETunnelHealthCheckDataSourceModel] `tfsdk:"health_check" json:"health_check,computed"`
	ModifiedOn            timetypes.RFC3339                                                              `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Mtu                   types.Int64                                                                    `tfsdk:"mtu" json:"mtu,computed"`
	TTL                   types.Int64                                                                    `tfsdk:"ttl" json:"ttl,computed"`
}

type MagicWANGRETunnelGRETunnelHealthCheckDataSourceModel struct {
	Direction types.String                                                                         `tfsdk:"direction" json:"direction,computed"`
	Enabled   types.Bool                                                                           `tfsdk:"enabled" json:"enabled,computed"`
	Rate      jsontypes.Normalized                                                                 `tfsdk:"rate" json:"rate,computed"`
	Target    customfield.NestedObject[MagicWANGRETunnelGRETunnelHealthCheckTargetDataSourceModel] `tfsdk:"target" json:"target,computed"`
	Type      jsontypes.Normalized                                                                 `tfsdk:"type" json:"type,computed"`
}

type MagicWANGRETunnelGRETunnelHealthCheckTargetDataSourceModel struct {
	Effective types.String `tfsdk:"effective" json:"effective,computed"`
	Saved     types.String `tfsdk:"saved" json:"saved,computed"`
}
