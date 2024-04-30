// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteResultEnvelope struct {
	Result MagicTransitSiteModel `json:"result,computed"`
}

type MagicTransitSiteModel struct {
	ID                   types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID            types.String                   `tfsdk:"account_id" path:"account_id"`
	Name                 types.String                   `tfsdk:"name" json:"name"`
	ConnectorID          types.String                   `tfsdk:"connector_id" json:"connector_id"`
	Description          types.String                   `tfsdk:"description" json:"description"`
	HaMode               types.Bool                     `tfsdk:"ha_mode" json:"ha_mode"`
	Location             *MagicTransitSiteLocationModel `tfsdk:"location" json:"location"`
	SecondaryConnectorID types.String                   `tfsdk:"secondary_connector_id" json:"secondary_connector_id"`
}

type MagicTransitSiteLocationModel struct {
	Lat types.String `tfsdk:"lat" json:"lat"`
	Lon types.String `tfsdk:"lon" json:"lon"`
}
