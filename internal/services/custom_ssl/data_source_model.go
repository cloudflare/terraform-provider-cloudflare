// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultDataSourceEnvelope struct {
	Result CustomSSLDataSourceModel `json:"result,computed"`
}

type CustomSSLResultListDataSourceEnvelope struct {
	Result *[]*CustomSSLDataSourceModel `json:"result,computed"`
}

type CustomSSLDataSourceModel struct {
	ZoneID              types.String                             `tfsdk:"zone_id" json:"zone_id"`
	CustomCertificateID types.String                             `tfsdk:"custom_certificate_id" path:"custom_certificate_id"`
	ID                  types.String                             `tfsdk:"id" json:"id"`
	BundleMethod        types.String                             `tfsdk:"bundle_method" json:"bundle_method"`
	ExpiresOn           types.String                             `tfsdk:"expires_on" json:"expires_on"`
	Hosts               types.String                             `tfsdk:"hosts" json:"hosts"`
	Issuer              types.String                             `tfsdk:"issuer" json:"issuer"`
	ModifiedOn          types.String                             `tfsdk:"modified_on" json:"modified_on"`
	Priority            types.Float64                            `tfsdk:"priority" json:"priority"`
	Signature           types.String                             `tfsdk:"signature" json:"signature"`
	Status              types.String                             `tfsdk:"status" json:"status"`
	UploadedOn          types.String                             `tfsdk:"uploaded_on" json:"uploaded_on"`
	GeoRestrictions     *CustomSSLGeoRestrictionsDataSourceModel `tfsdk:"geo_restrictions" json:"geo_restrictions"`
	KeylessServer       *CustomSSLKeylessServerDataSourceModel   `tfsdk:"keyless_server" json:"keyless_server"`
	Policy              types.String                             `tfsdk:"policy" json:"policy"`
	FindOneBy           *CustomSSLFindOneByDataSourceModel       `tfsdk:"find_one_by"`
}

type CustomSSLGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label"`
}

type CustomSSLKeylessServerDataSourceModel struct {
	ID          types.String                                 `tfsdk:"id" json:"id,computed"`
	CreatedOn   types.String                                 `tfsdk:"created_on" json:"created_on,computed"`
	Enabled     types.Bool                                   `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                 `tfsdk:"host" json:"host,computed"`
	ModifiedOn  types.String                                 `tfsdk:"modified_on" json:"modified_on,computed"`
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
	ZoneID  types.String  `tfsdk:"zone_id" path:"zone_id"`
	Match   types.String  `tfsdk:"match" query:"match"`
	Page    types.Float64 `tfsdk:"page" query:"page"`
	PerPage types.Float64 `tfsdk:"per_page" query:"per_page"`
	Status  types.String  `tfsdk:"status" query:"status"`
}
