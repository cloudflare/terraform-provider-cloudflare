// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_cf1_site

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MagicTransitCf1SiteResultEnvelope struct {
	Result *[]*MagicTransitCf1SiteBodyModel `json:"result"`
}

type MagicTransitCf1SiteModel struct {
	ID          types.String                      `tfsdk:"id" json:"id,computed"`
	AccountID   types.String                      `tfsdk:"account_id" path:"account_id,required"`
	Body        *[]*MagicTransitCf1SiteBodyModel  `tfsdk:"body" json:"body,required,no_refresh"`
	Description types.String                      `tfsdk:"description" json:"description,optional"`
	Name        types.String                      `tfsdk:"name" json:"name,optional"`
	Location    *MagicTransitCf1SiteLocationModel `tfsdk:"location" json:"location,optional"`
	CreatedOn   timetypes.RFC3339                 `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	ModifiedOn  timetypes.RFC3339                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

func (m MagicTransitCf1SiteModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m.Body)
}

func (m MagicTransitCf1SiteModel) MarshalJSONForUpdate(state MagicTransitCf1SiteModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m.Body, state.Body)
}

type MagicTransitCf1SiteBodyModel struct {
	Name        types.String                          `tfsdk:"name" json:"name,required"`
	ID          types.String                          `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Description types.String                          `tfsdk:"description" json:"description,optional"`
	Location    *MagicTransitCf1SiteBodyLocationModel `tfsdk:"location" json:"location,optional"`
	ModifiedOn  timetypes.RFC3339                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
}

type MagicTransitCf1SiteBodyLocationModel struct {
	Lat  types.Float64 `tfsdk:"lat" json:"lat,optional"`
	Long types.Float64 `tfsdk:"long" json:"long,optional"`
	Name types.String  `tfsdk:"name" json:"name,optional"`
}

type MagicTransitCf1SiteLocationModel struct {
	Lat  types.Float64 `tfsdk:"lat" json:"lat,optional"`
	Long types.Float64 `tfsdk:"long" json:"long,optional"`
	Name types.String  `tfsdk:"name" json:"name,optional"`
}
