// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLsResultListDataSourceEnvelope struct {
	Result *[]*CustomSSLsItemsDataSourceModel `json:"result,computed"`
}

type CustomSSLsDataSourceModel struct {
	ZoneID   types.String                       `tfsdk:"zone_id" path:"zone_id"`
	Match    types.String                       `tfsdk:"match" query:"match"`
	Page     types.Float64                      `tfsdk:"page" query:"page"`
	PerPage  types.Float64                      `tfsdk:"per_page" query:"per_page"`
	Status   types.String                       `tfsdk:"status" query:"status"`
	MaxItems types.Int64                        `tfsdk:"max_items"`
	Items    *[]*CustomSSLsItemsDataSourceModel `tfsdk:"items"`
}

type CustomSSLsItemsDataSourceModel struct {
	ID              types.String                                   `tfsdk:"id" json:"id,computed"`
	BundleMethod    types.String                                   `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       types.String                                   `tfsdk:"expires_on" json:"expires_on,computed"`
	Hosts           *[]types.String                                `tfsdk:"hosts" json:"hosts,computed"`
	Issuer          types.String                                   `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      types.String                                   `tfsdk:"modified_on" json:"modified_on,computed"`
	Priority        types.Float64                                  `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                   `tfsdk:"signature" json:"signature,computed"`
	Status          types.String                                   `tfsdk:"status" json:"status,computed"`
	UploadedOn      types.String                                   `tfsdk:"uploaded_on" json:"uploaded_on,computed"`
	ZoneID          types.String                                   `tfsdk:"zone_id" json:"zone_id,computed"`
	GeoRestrictions *CustomSSLsItemsGeoRestrictionsDataSourceModel `tfsdk:"geo_restrictions" json:"geo_restrictions"`
	KeylessServer   *CustomSSLsItemsKeylessServerDataSourceModel   `tfsdk:"keyless_server" json:"keyless_server"`
	Policy          types.String                                   `tfsdk:"policy" json:"policy"`
}

type CustomSSLsItemsGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label"`
}

type CustomSSLsItemsKeylessServerDataSourceModel struct {
	ID          types.String                                       `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String                                       `tfsdk:"created_on" json:"created_on,computed"`
	Enabled     types.Bool                                         `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                       `tfsdk:"host" json:"host,computed"`
	ModifiedOn  types.String                                       `tfsdk:"modified_on" json:"modified_on,computed"`
	Name        types.String                                       `tfsdk:"name" json:"name,computed"`
	Permissions *[]types.String                                    `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                      `tfsdk:"port" json:"port,computed"`
	Status      types.String                                       `tfsdk:"status" json:"status,computed"`
	Tunnel      *CustomSSLsItemsKeylessServerTunnelDataSourceModel `tfsdk:"tunnel" json:"tunnel"`
}

type CustomSSLsItemsKeylessServerTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}
