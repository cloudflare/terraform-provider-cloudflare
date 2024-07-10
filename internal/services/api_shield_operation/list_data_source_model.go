// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_shield_operation

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type APIShieldOperationsResultListDataSourceEnvelope struct {
	Result *[]*APIShieldOperationsItemsDataSourceModel `json:"result,computed"`
}

type APIShieldOperationsDataSourceModel struct {
	ZoneID    types.String                                `tfsdk:"zone_id" path:"zone_id"`
	Diff      types.Bool                                  `tfsdk:"diff" query:"diff"`
	Direction types.String                                `tfsdk:"direction" query:"direction"`
	Endpoint  types.String                                `tfsdk:"endpoint" query:"endpoint"`
	Host      *[]types.String                             `tfsdk:"host" query:"host"`
	Method    *[]types.String                             `tfsdk:"method" query:"method"`
	Order     types.String                                `tfsdk:"order" query:"order"`
	Origin    types.String                                `tfsdk:"origin" query:"origin"`
	Page      types.String                                `tfsdk:"page" query:"page"`
	PerPage   types.String                                `tfsdk:"per_page" query:"per_page"`
	State     types.String                                `tfsdk:"state" query:"state"`
	MaxItems  types.Int64                                 `tfsdk:"max_items"`
	Items     *[]*APIShieldOperationsItemsDataSourceModel `tfsdk:"items"`
}

type APIShieldOperationsItemsDataSourceModel struct {
	ID          types.String    `tfsdk:"id" json:"id,computed"`
	Endpoint    types.String    `tfsdk:"endpoint" json:"endpoint,computed"`
	Host        types.String    `tfsdk:"host" json:"host,computed"`
	LastUpdated types.String    `tfsdk:"last_updated" json:"last_updated,computed"`
	Method      types.String    `tfsdk:"method" json:"method,computed"`
	Origin      *[]types.String `tfsdk:"origin" json:"origin,computed"`
	State       types.String    `tfsdk:"state" json:"state,computed"`
}
