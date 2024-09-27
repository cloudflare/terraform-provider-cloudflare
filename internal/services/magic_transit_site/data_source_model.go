// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v3"
	"github.com/cloudflare/cloudflare-go/v3/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSiteResultDataSourceEnvelope struct {
	Result MagicTransitSiteDataSourceModel `json:"result,computed"`
}

type MagicTransitSiteResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitSiteDataSourceModel] `json:"result,computed"`
}

type MagicTransitSiteDataSourceModel struct {
	AccountID            types.String                                                      `tfsdk:"account_id" path:"account_id,optional"`
	SiteID               types.String                                                      `tfsdk:"site_id" path:"site_id,optional"`
	ConnectorID          types.String                                                      `tfsdk:"connector_id" json:"connector_id,computed"`
	Description          types.String                                                      `tfsdk:"description" json:"description,computed"`
	HaMode               types.Bool                                                        `tfsdk:"ha_mode" json:"ha_mode,computed"`
	ID                   types.String                                                      `tfsdk:"id" json:"id,computed"`
	Name                 types.String                                                      `tfsdk:"name" json:"name,computed"`
	SecondaryConnectorID types.String                                                      `tfsdk:"secondary_connector_id" json:"secondary_connector_id,computed"`
	Location             customfield.NestedObject[MagicTransitSiteLocationDataSourceModel] `tfsdk:"location" json:"location,computed"`
	Filter               *MagicTransitSiteFindOneByDataSourceModel                         `tfsdk:"filter"`
}

func (m *MagicTransitSiteDataSourceModel) toReadParams(_ context.Context) (params magic_transit.SiteGetParams, diags diag.Diagnostics) {
	params = magic_transit.SiteGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

func (m *MagicTransitSiteDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteListParams, diags diag.Diagnostics) {
	params = magic_transit.SiteListParams{
		AccountID: cloudflare.F(m.Filter.AccountID.ValueString()),
	}

	if !m.Filter.ConnectorIdentifier.IsNull() {
		params.ConnectorIdentifier = cloudflare.F(m.Filter.ConnectorIdentifier.ValueString())
	}

	return
}

type MagicTransitSiteLocationDataSourceModel struct {
	Lat types.String `tfsdk:"lat" json:"lat,computed"`
	Lon types.String `tfsdk:"lon" json:"lon,computed"`
}

type MagicTransitSiteFindOneByDataSourceModel struct {
	AccountID           types.String `tfsdk:"account_id" path:"account_id,required"`
	ConnectorIdentifier types.String `tfsdk:"connector_identifier" query:"connector_identifier,optional"`
}
