// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteResultEnvelope struct {
	Result MagicTransitSiteModel `json:"result"`
}

type MagicTransitSiteModel struct {
	ID                   types.String                   `tfsdk:"id" json:"id,computed"`
	AccountID            types.String                   `tfsdk:"account_id" path:"account_id,required"`
	HaMode               types.Bool                     `tfsdk:"ha_mode" json:"ha_mode,optional"`
	Name                 types.String                   `tfsdk:"name" json:"name,required"`
	ConnectorID          types.String                   `tfsdk:"connector_id" json:"connector_id,optional"`
	Description          types.String                   `tfsdk:"description" json:"description,optional"`
	SecondaryConnectorID types.String                   `tfsdk:"secondary_connector_id" json:"secondary_connector_id,optional"`
	Location             *MagicTransitSiteLocationModel `tfsdk:"location" json:"location,optional"`
}

func (m MagicTransitSiteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m MagicTransitSiteModel) MarshalJSONForUpdate(state MagicTransitSiteModel) (data []byte, err error) {
	return apijson.MarshalForUpdate(m, state)
}

type MagicTransitSiteLocationModel struct {
	Lat types.String `tfsdk:"lat" json:"lat,optional"`
	Lon types.String `tfsdk:"lon" json:"lon,optional"`
}
