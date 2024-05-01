// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteACLResultEnvelope struct {
	Result MagicTransitSiteACLModel `json:"result,computed"`
}

type MagicTransitSiteACLModel struct {
	ID             types.String                  `tfsdk:"id" json:"id,computed"`
	AccountID      types.String                  `tfsdk:"account_id" path:"account_id"`
	SiteID         types.String                  `tfsdk:"site_id" path:"site_id"`
	LAN1           *MagicTransitSiteACLLAN1Model `tfsdk:"lan_1" json:"lan_1"`
	LAN2           *MagicTransitSiteACLLAN2Model `tfsdk:"lan_2" json:"lan_2"`
	Name           types.String                  `tfsdk:"name" json:"name"`
	Description    types.String                  `tfsdk:"description" json:"description"`
	ForwardLocally types.Bool                    `tfsdk:"forward_locally" json:"forward_locally"`
	Protocols      *[]types.String               `tfsdk:"protocols" json:"protocols"`
}

type MagicTransitSiteACLLAN1Model struct {
	LANID   types.String    `tfsdk:"lan_id" json:"lan_id"`
	LANName types.String    `tfsdk:"lan_name" json:"lan_name"`
	Ports   *[]types.Int64  `tfsdk:"ports" json:"ports"`
	Subnets *[]types.String `tfsdk:"subnets" json:"subnets"`
}

type MagicTransitSiteACLLAN2Model struct {
	LANID   types.String    `tfsdk:"lan_id" json:"lan_id"`
	LANName types.String    `tfsdk:"lan_name" json:"lan_name"`
	Ports   *[]types.Int64  `tfsdk:"ports" json:"ports"`
	Subnets *[]types.String `tfsdk:"subnets" json:"subnets"`
}
