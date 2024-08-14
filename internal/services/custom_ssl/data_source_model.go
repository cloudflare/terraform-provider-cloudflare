// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultDataSourceEnvelope struct {
	Result CustomSSLDataSourceModel `json:"result,computed"`
}

type CustomSSLResultListDataSourceEnvelope struct {
	Result *[]*CustomSSLDataSourceModel `json:"result,computed"`
}

type CustomSSLDataSourceModel struct {
	CustomCertificateID types.String                             `tfsdk:"custom_certificate_id" path:"custom_certificate_id"`
	ZoneID              types.String                             `tfsdk:"zone_id" path:"zone_id"`
	BundleMethod        types.String                             `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn           timetypes.RFC3339                        `tfsdk:"expires_on" json:"expires_on,computed"`
	ID                  types.String                             `tfsdk:"id" json:"id,computed"`
	Issuer              types.String                             `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn          timetypes.RFC3339                        `tfsdk:"modified_on" json:"modified_on,computed"`
	Priority            types.Float64                            `tfsdk:"priority" json:"priority,computed"`
	Signature           types.String                             `tfsdk:"signature" json:"signature,computed"`
	Status              types.String                             `tfsdk:"status" json:"status,computed"`
	UploadedOn          timetypes.RFC3339                        `tfsdk:"uploaded_on" json:"uploaded_on,computed"`
	Hosts               *[]types.String                          `tfsdk:"hosts" json:"hosts,computed"`
	Policy              types.String                             `tfsdk:"policy" json:"policy"`
	GeoRestrictions     *CustomSSLGeoRestrictionsDataSourceModel `tfsdk:"geo_restrictions" json:"geo_restrictions"`
	KeylessServer       *CustomSSLKeylessServerDataSourceModel   `tfsdk:"keyless_server" json:"keyless_server"`
	Filter              *CustomSSLFindOneByDataSourceModel       `tfsdk:"filter"`
}

type CustomSSLGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label"`
}

type CustomSSLKeylessServerDataSourceModel struct {
	ID          types.String                                 `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                            `tfsdk:"created_on" json:"created_on,computed"`
	Enabled     types.Bool                                   `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                 `tfsdk:"host" json:"host,computed"`
	ModifiedOn  timetypes.RFC3339                            `tfsdk:"modified_on" json:"modified_on,computed"`
	Name        types.String                                 `tfsdk:"name" json:"name,computed"`
	Permissions *[]types.String                              `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                `tfsdk:"port" json:"port,computed"`
	Status      types.String                                 `tfsdk:"status" json:"status,computed"`
	Tunnel      *CustomSSLKeylessServerTunnelDataSourceModel `tfsdk:"tunnel" json:"tunnel"`
}

type CustomSSLKeylessServerTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type CustomSSLFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Match  types.String `tfsdk:"match" query:"match"`
	Status types.String `tfsdk:"status" query:"status"`
}
