// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteACLsResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitSiteACLsResultDataSourceModel] `json:"result,computed"`
}

type MagicTransitSiteACLsDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	SiteID    types.String                                                            `tfsdk:"site_id" path:"site_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicTransitSiteACLsResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicTransitSiteACLsDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteACLListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteACLListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteACLsResultDataSourceModel struct {
	ID             types.String                                                      `tfsdk:"id" json:"id,computed"`
	Description    types.String                                                      `tfsdk:"description" json:"description,computed"`
	ForwardLocally types.Bool                                                        `tfsdk:"forward_locally" json:"forward_locally,computed"`
	LAN1           customfield.NestedObject[MagicTransitSiteACLsLAN1DataSourceModel] `tfsdk:"lan_1" json:"lan_1,computed"`
	LAN2           customfield.NestedObject[MagicTransitSiteACLsLAN2DataSourceModel] `tfsdk:"lan_2" json:"lan_2,computed"`
	Name           types.String                                                      `tfsdk:"name" json:"name,computed"`
	Protocols      customfield.List[jsontypes.Normalized]                            `tfsdk:"protocols" json:"protocols,computed"`
	Unidirectional types.Bool                                                        `tfsdk:"unidirectional" json:"unidirectional,computed"`
}

type MagicTransitSiteACLsLAN1DataSourceModel struct {
	LANID      types.String                   `tfsdk:"lan_id" json:"lan_id,computed"`
	LANName    types.String                   `tfsdk:"lan_name" json:"lan_name,computed"`
	PortRanges customfield.List[types.String] `tfsdk:"port_ranges" json:"port_ranges,computed"`
	Ports      customfield.List[types.Int64]  `tfsdk:"ports" json:"ports,computed"`
	Subnets    customfield.List[types.String] `tfsdk:"subnets" json:"subnets,computed"`
}

type MagicTransitSiteACLsLAN2DataSourceModel struct {
	LANID      types.String                   `tfsdk:"lan_id" json:"lan_id,computed"`
	LANName    types.String                   `tfsdk:"lan_name" json:"lan_name,computed"`
	PortRanges customfield.List[types.String] `tfsdk:"port_ranges" json:"port_ranges,computed"`
	Ports      customfield.List[types.Int64]  `tfsdk:"ports" json:"ports,computed"`
	Subnets    customfield.List[types.String] `tfsdk:"subnets" json:"subnets,computed"`
}
