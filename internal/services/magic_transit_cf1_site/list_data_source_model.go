// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_cf1_site

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v7"
	"github.com/cloudflare/cloudflare-go/v7/magic_transit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitCf1SitesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[MagicTransitCf1SitesResultDataSourceModel] `json:"result,computed"`
}

type MagicTransitCf1SitesDataSourceModel struct {
	AccountID types.String                                                            `tfsdk:"account_id" path:"account_id,required"`
	MaxItems  types.Int64                                                             `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[MagicTransitCf1SitesResultDataSourceModel] `tfsdk:"result"`
}

func (m *MagicTransitCf1SitesDataSourceModel) toListParams(_ context.Context) (params magic_transit.Cf1SiteListParams, diags diag.Diagnostics) {
	params = magic_transit.Cf1SiteListParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitCf1SitesResultDataSourceModel struct {
	Name        types.String                                                          `tfsdk:"name" json:"name,computed"`
	ID          types.String                                                          `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                          `tfsdk:"description" json:"description,computed"`
	Location    customfield.NestedObject[MagicTransitCf1SitesLocationDataSourceModel] `tfsdk:"location" json:"location,computed"`
	ModifiedOn  timetypes.RFC3339                                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type MagicTransitCf1SitesLocationDataSourceModel struct {
	Lat  types.Float64 `tfsdk:"lat" json:"lat,computed"`
	Long types.Float64 `tfsdk:"long" json:"long,computed"`
	Name types.String  `tfsdk:"name" json:"name,computed"`
}
