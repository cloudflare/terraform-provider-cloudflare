// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLsResultListDataSourceEnvelope struct {
	Result *[]*CustomSSLsResultDataSourceModel `json:"result,computed"`
}

type CustomSSLsDataSourceModel struct {
	ZoneID   types.String                        `tfsdk:"zone_id" path:"zone_id"`
	Status   types.String                        `tfsdk:"status" query:"status"`
	Match    types.String                        `tfsdk:"match" query:"match"`
	MaxItems types.Int64                         `tfsdk:"max_items"`
	Result   *[]*CustomSSLsResultDataSourceModel `tfsdk:"result"`
}

type CustomSSLsResultDataSourceModel struct {
	ID              types.String                              `tfsdk:"id" json:"id,computed"`
	BundleMethod    types.String                              `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       timetypes.RFC3339                         `tfsdk:"expires_on" json:"expires_on,computed"`
	Hosts           *[]types.String                           `tfsdk:"hosts" json:"hosts,computed"`
	Issuer          types.String                              `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                         `tfsdk:"modified_on" json:"modified_on,computed"`
	Priority        types.Float64                             `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                              `tfsdk:"signature" json:"signature,computed"`
	Status          types.String                              `tfsdk:"status" json:"status,computed"`
	UploadedOn      timetypes.RFC3339                         `tfsdk:"uploaded_on" json:"uploaded_on,computed"`
	ZoneID          types.String                              `tfsdk:"zone_id" json:"zone_id,computed"`
	GeoRestrictions *CustomSSLsGeoRestrictionsDataSourceModel `tfsdk:"geo_restrictions" json:"geo_restrictions"`
	KeylessServer   *CustomSSLsKeylessServerDataSourceModel   `tfsdk:"keyless_server" json:"keyless_server"`
	Policy          types.String                              `tfsdk:"policy" json:"policy"`
}

type CustomSSLsGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label"`
}

type CustomSSLsKeylessServerDataSourceModel struct {
	ID          types.String                                  `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                             `tfsdk:"created_on" json:"created_on,computed"`
	Enabled     types.Bool                                    `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                  `tfsdk:"host" json:"host,computed"`
	ModifiedOn  timetypes.RFC3339                             `tfsdk:"modified_on" json:"modified_on,computed"`
	Name        types.String                                  `tfsdk:"name" json:"name,computed"`
	Permissions *[]jsontypes.Normalized                       `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                 `tfsdk:"port" json:"port,computed"`
	Status      types.String                                  `tfsdk:"status" json:"status,computed"`
	Tunnel      *CustomSSLsKeylessServerTunnelDataSourceModel `tfsdk:"tunnel" json:"tunnel"`
}

type CustomSSLsKeylessServerTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}
