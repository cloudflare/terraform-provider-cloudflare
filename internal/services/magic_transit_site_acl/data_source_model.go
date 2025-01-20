// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteACLResultDataSourceEnvelope struct {
	Result MagicTransitSiteACLDataSourceModel `json:"result,computed"`
}

type MagicTransitSiteACLResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitSiteACLDataSourceModel] `json:"result,computed"`
}

type MagicTransitSiteACLDataSourceModel struct {
	AccountID      types.String                                                     `tfsdk:"account_id" path:"account_id,optional"`
	ACLID          types.String                                                     `tfsdk:"acl_id" path:"acl_id,optional"`
	SiteID         types.String                                                     `tfsdk:"site_id" path:"site_id,optional"`
	Description    types.String                                                     `tfsdk:"description" json:"description,computed"`
	ForwardLocally types.Bool                                                       `tfsdk:"forward_locally" json:"forward_locally,computed"`
	ID             types.String                                                     `tfsdk:"id" json:"id,computed"`
	Name           types.String                                                     `tfsdk:"name" json:"name,computed"`
	Protocols      customfield.List[types.String]                                   `tfsdk:"protocols" json:"protocols,computed"`
	LAN1           customfield.NestedObject[MagicTransitSiteACLLAN1DataSourceModel] `tfsdk:"lan_1" json:"lan_1,computed"`
	LAN2           customfield.NestedObject[MagicTransitSiteACLLAN2DataSourceModel] `tfsdk:"lan_2" json:"lan_2,computed"`
	Filter         *MagicTransitSiteACLFindOneByDataSourceModel                     `tfsdk:"filter"`
}

func (m *MagicTransitSiteACLDataSourceModel) toReadParams(_ context.Context) (params magic_transit.SiteACLGetParams, diags diag.Diagnostics) {
	params = magic_transit.SiteACLGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MagicTransitSiteACLDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteACLListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteACLListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	return
}

type MagicTransitSiteACLLAN1DataSourceModel struct {
	LANID   types.String                   `tfsdk:"lan_id" json:"lan_id,computed"`
	LANName types.String                   `tfsdk:"lan_name" json:"lan_name,computed"`
	Ports   customfield.List[types.Int64]  `tfsdk:"ports" json:"ports,computed"`
	Subnets customfield.List[types.String] `tfsdk:"subnets" json:"subnets,computed"`
}

type MagicTransitSiteACLLAN2DataSourceModel struct {
	LANID   types.String                   `tfsdk:"lan_id" json:"lan_id,computed"`
	LANName types.String                   `tfsdk:"lan_name" json:"lan_name,computed"`
	Ports   customfield.List[types.Int64]  `tfsdk:"ports" json:"ports,computed"`
	Subnets customfield.List[types.String] `tfsdk:"subnets" json:"subnets,computed"`
}

type MagicTransitSiteACLFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
	SiteID    types.String `tfsdk:"site_id" path:"site_id,required"`
}
