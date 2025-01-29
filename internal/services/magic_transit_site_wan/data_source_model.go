// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteWANResultDataSourceEnvelope struct {
	Result MagicTransitSiteWANDataSourceModel `json:"result,computed"`
}

type MagicTransitSiteWANDataSourceModel struct {
	AccountID        types.String                                                                 `tfsdk:"account_id" path:"account_id,required"`
	WANID            types.String                                                                 `tfsdk:"wan_id" path:"wan_id,required"`
	SiteID           types.String                                                                 `tfsdk:"site_id" path:"site_id,computed"`
	HealthCheckRate  types.String                                                                 `tfsdk:"health_check_rate" json:"health_check_rate,computed"`
	ID               types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Name             types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Physport         types.Int64                                                                  `tfsdk:"physport" json:"physport,computed"`
	Priority         types.Int64                                                                  `tfsdk:"priority" json:"priority,computed"`
	VlanTag          types.Int64                                                                  `tfsdk:"vlan_tag" json:"vlan_tag,computed"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteWANStaticAddressingDataSourceModel] `tfsdk:"static_addressing" json:"static_addressing,computed"`
}

func (m *MagicTransitSiteWANDataSourceModel) toReadParams(_ context.Context) (params magic_transit.SiteWANGetParams, diags diag.Diagnostics) {
	params = magic_transit.SiteWANGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteWANStaticAddressingDataSourceModel struct {
	Address          types.String `tfsdk:"address" json:"address,computed"`
	GatewayAddress   types.String `tfsdk:"gateway_address" json:"gateway_address,computed"`
	SecondaryAddress types.String `tfsdk:"secondary_address" json:"secondary_address,computed"`
}
