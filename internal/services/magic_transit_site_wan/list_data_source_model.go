// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_wan

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteWANsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitSiteWANsResultDataSourceModel] `json:"result,computed"`
}

type MagicTransitSiteWANsDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	SiteID    types.String                                                            `tfsdk:"site_id" path:"site_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicTransitSiteWANsResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicTransitSiteWANsDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteWANListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteWANListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteWANsResultDataSourceModel struct {
	ID               types.String                                                                  `tfsdk:"id" json:"id,computed"`
	Name             types.String                                                                  `tfsdk:"name" json:"name,computed"`
	Physport         types.Int64                                                                   `tfsdk:"physport" json:"physport,computed"`
	Priority         types.Int64                                                                   `tfsdk:"priority" json:"priority,computed"`
	SiteID           types.String                                                                  `tfsdk:"site_id" json:"site_id,computed"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteWANsStaticAddressingDataSourceModel] `tfsdk:"static_addressing" json:"static_addressing,computed"`
	VlanTag          types.Int64                                                                   `tfsdk:"vlan_tag" json:"vlan_tag,computed"`
}

type MagicTransitSiteWANsStaticAddressingDataSourceModel struct {
	Address          types.String `tfsdk:"address" json:"address,computed"`
	GatewayAddress   types.String `tfsdk:"gateway_address" json:"gateway_address,computed"`
	SecondaryAddress types.String `tfsdk:"secondary_address" json:"secondary_address,computed"`
}
