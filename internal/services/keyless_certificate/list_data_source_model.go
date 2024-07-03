// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KeylessCertificatesResultListDataSourceEnvelope struct {
	Result *[]*KeylessCertificatesItemsDataSourceModel `json:"result,computed"`
}

type KeylessCertificatesDataSourceModel struct {
	ZoneID   types.String                                `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                                 `tfsdk:"max_items"`
	Items    *[]*KeylessCertificatesItemsDataSourceModel `tfsdk:"items"`
}

type KeylessCertificatesItemsDataSourceModel struct {
	ID          types.String    `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String    `tfsdk:"created_on" json:"created_on,computed"`
	Enabled     types.Bool      `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String    `tfsdk:"host" json:"host,computed"`
	ModifiedOn  types.String    `tfsdk:"modified_on" json:"modified_on,computed"`
	Name        types.String    `tfsdk:"name" json:"name,computed"`
	Permissions *[]types.String `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64   `tfsdk:"port" json:"port,computed"`
	Status      types.String    `tfsdk:"status" json:"status,computed"`
}

type KeylessCertificatesItemsTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}
