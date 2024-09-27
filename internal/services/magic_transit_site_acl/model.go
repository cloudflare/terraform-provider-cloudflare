// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_acl

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteACLResultEnvelope struct {
	Result MagicTransitSiteACLModel `json:"result"`
}

type MagicTransitSiteACLModel struct {
	ID             types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID      types.String                   `tfsdk:"account_id" path:"account_id,required"`
	SiteID         types.String                   `tfsdk:"site_id" path:"site_id,required"`
	Name           types.String                   `tfsdk:"name" json:"name,required"`
	LAN1           *MagicTransitSiteACLLAN1Model  `tfsdk:"lan_1" json:"lan_1,required"`
	LAN2           *MagicTransitSiteACLLAN2Model  `tfsdk:"lan_2" json:"lan_2,required"`
	Description    types.String                   `tfsdk:"description" json:"description,computed_optional"`
	ForwardLocally types.Bool                     `tfsdk:"forward_locally" json:"forward_locally,computed_optional"`
	Protocols      customfield.List[types.String] `tfsdk:"protocols" json:"protocols,computed_optional"`
}

type MagicTransitSiteACLLAN1Model struct {
	LANID   types.String                   `tfsdk:"lan_id" json:"lan_id,required"`
	LANName types.String                   `tfsdk:"lan_name" json:"lan_name,computed_optional"`
	Ports   customfield.List[types.Int64]  `tfsdk:"ports" json:"ports,computed_optional"`
	Subnets customfield.List[types.String] `tfsdk:"subnets" json:"subnets,computed_optional"`
}

type MagicTransitSiteACLLAN2Model struct {
	LANID   types.String                   `tfsdk:"lan_id" json:"lan_id,required"`
	LANName types.String                   `tfsdk:"lan_name" json:"lan_name,computed_optional"`
	Ports   customfield.List[types.Int64]  `tfsdk:"ports" json:"ports,computed_optional"`
	Subnets customfield.List[types.String] `tfsdk:"subnets" json:"subnets,computed_optional"`
}
