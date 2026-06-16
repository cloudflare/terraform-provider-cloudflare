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

type MagicTransitCf1SiteResultDataSourceEnvelope struct {
	Result MagicTransitCf1SiteDataSourceModel `json:"result,computed"`
}

type MagicTransitCf1SiteDataSourceModel struct {
	ID          types.String                                                         `tfsdk:"id" path:"cf1_site_id,computed"`
	Cf1SiteID   types.String                                                         `tfsdk:"cf1_site_id" path:"cf1_site_id,required"`
	AccountID   types.String                                                         `tfsdk:"account_id" path:"account_id,required"`
	CreatedOn   timetypes.RFC3339                                                    `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                                                         `tfsdk:"description" json:"description,computed"`
	ModifiedOn  timetypes.RFC3339                                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                         `tfsdk:"name" json:"name,computed"`
	Location    customfield.NestedObject[MagicTransitCf1SiteLocationDataSourceModel] `tfsdk:"location" json:"location,computed"`
}

func (m *MagicTransitCf1SiteDataSourceModel) toReadParams(_ context.Context) (params magic_transit.Cf1SiteGetParams, diags diag.Diagnostics) {
	params = magic_transit.Cf1SiteGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}

type MagicTransitCf1SiteLocationDataSourceModel struct {
	Lat  types.Float64 `tfsdk:"lat" json:"lat,computed"`
	Long types.Float64 `tfsdk:"long" json:"long,computed"`
	Name types.String  `tfsdk:"name" json:"name,computed"`
}
