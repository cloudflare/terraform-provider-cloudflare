// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
  "context"

  "github.com/cloudflare/cloudflare-go/v4"
  "github.com/cloudflare/cloudflare-go/v4/magic_transit"
  "github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
  "github.com/hashicorp/terraform-plugin-framework/diag"
  "github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitSitesResultListDataSourceEnvelope struct {
Result customfield.NestedObjectList[MagicTransitSitesResultDataSourceModel] `json:"result,computed"`
}

type MagicTransitSitesDataSourceModel struct {
AccountID types.String `tfsdk:"account_id" path:"account_id,required"`
Connectorid types.String `tfsdk:"connectorid" query:"connectorid,optional"`
MaxItems types.Int64 `tfsdk:"max_items"`
Result customfield.NestedObjectList[MagicTransitSitesResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicTransitSitesDataSourceModel) toListParams(_ context.Context) (params magic_transit.SiteListParams, diags diag.Diagnostics) {
  params = magic_transit.SiteListParams{
    AccountID: cloudflare.F(m.AccountID.ValueString()),
  }

  if !m.Connectorid.IsNull() {
    params.Connectorid = cloudflare.F(m.Connectorid.ValueString())
  }

  return
}

type MagicTransitSitesResultDataSourceModel struct {
ID types.String `tfsdk:"id" json:"id,computed"`
ConnectorID types.String `tfsdk:"connector_id" json:"connector_id,computed"`
Description types.String `tfsdk:"description" json:"description,computed"`
HaMode types.Bool `tfsdk:"ha_mode" json:"ha_mode,computed"`
Location customfield.NestedObject[MagicTransitSitesLocationDataSourceModel] `tfsdk:"location" json:"location,computed"`
Name types.String `tfsdk:"name" json:"name,computed"`
SecondaryConnectorID types.String `tfsdk:"secondary_connector_id" json:"secondary_connector_id,computed"`
}

type MagicTransitSitesLocationDataSourceModel struct {
Lat types.String `tfsdk:"lat" json:"lat,computed"`
Lon types.String `tfsdk:"lon" json:"lon,computed"`
}
