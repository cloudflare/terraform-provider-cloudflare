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

type MagicTransitSiteWANResultDataSourceEnvelope struct {
	Result MagicTransitSiteWANDataSourceModel `json:"result,computed"`
}

type MagicTransitSiteWANResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitSiteWANDataSourceModel] `json:"result,computed"`
}

type MagicTransitSiteWANDataSourceModel struct {
	AccountID        types.String                                                                 `tfsdk:"account_id" path:"account_id,optional"`
	WANID            types.String                                                                 `tfsdk:"wan_id" path:"wan_id,optional"`
	SiteID           types.String                                                                 `tfsdk:"site_id" path:"site_id,computed_optional"`
	HealthCheckRate  types.String                                                                 `tfsdk:"health_check_rate" json:"health_check_rate,computed"`
	ID               types.String                                                                 `tfsdk:"id" json:"id,computed"`
	Name             types.String                                                                 `tfsdk:"name" json:"name,computed"`
	Physport         types.Int64                                                                  `tfsdk:"physport" json:"physport,computed"`
	Priority         types.Int64                                                                  `tfsdk:"priority" json:"priority,computed"`
	VlanTag          types.Int64                                                                  `tfsdk:"vlan_tag" json:"vlan_tag,computed"`
	StaticAddressing customfield.NestedObject[MagicTransitSiteWANStaticAddressingDataSourceModel] `tfsdk:"static_addressing" json:"static_addressing,computed"`
	Filter           *MagicTransitSiteWANFindOneByDataSourceModel                                 `tfsdk:"filter"`
}

func (m *MagicTransitSiteWANDataSourceModel) toReadParams(_ context.Context) (params magic_transit.SiteWANGetParams, diags diag.Diagnostics) {
	params = magic_transit.SiteWANGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MagicTransitSiteWANDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteWANListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteWANListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteWANStaticAddressingDataSourceModel struct {
	Address          types.String `tfsdk:"address" json:"address,computed"`
	GatewayAddress   types.String `tfsdk:"gateway_address" json:"gateway_address,computed"`
	SecondaryAddress types.String `tfsdk:"secondary_address" json:"secondary_address,computed"`
}

type MagicTransitSiteWANFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	SiteID    types.String `tfsdk:"site_id" path:"site_id,required"`
}
