// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultEnvelope struct {
	Result CustomSSLModel `json:"result"`
}

type CustomSSLModel struct {
	ID              types.String                                          `tfsdk:"id" json:"id,computed"`
	ZoneID          types.String                                          `tfsdk:"zone_id" path:"zone_id"`
	Type            types.String                                          `tfsdk:"type" json:"type,computed_optional"`
	Certificate     types.String                                          `tfsdk:"certificate" json:"certificate"`
	PrivateKey      types.String                                          `tfsdk:"private_key" json:"private_key"`
	Policy          types.String                                          `tfsdk:"policy" json:"policy"`
	GeoRestrictions *CustomSSLGeoRestrictionsModel                        `tfsdk:"geo_restrictions" json:"geo_restrictions"`
	BundleMethod    types.String                                          `tfsdk:"bundle_method" json:"bundle_method,computed_optional"`
	ExpiresOn       timetypes.RFC3339                                     `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer          types.String                                          `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority        types.Float64                                         `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                          `tfsdk:"signature" json:"signature,computed"`
	Status          types.String                                          `tfsdk:"status" json:"status,computed"`
	UploadedOn      timetypes.RFC3339                                     `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	Hosts           types.List                                            `tfsdk:"hosts" json:"hosts,computed"`
	KeylessServer   customfield.NestedObject[CustomSSLKeylessServerModel] `tfsdk:"keyless_server" json:"keyless_server,computed"`
}

type CustomSSLGeoRestrictionsModel struct {
	Label types.String `tfsdk:"label" json:"label,computed_optional"`
}

type CustomSSLKeylessServerModel struct {
	ID          types.String                                                `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                           `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Enabled     types.Bool                                                  `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                                `tfsdk:"host" json:"host"`
	ModifiedOn  timetypes.RFC3339                                           `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                `tfsdk:"name" json:"name,computed"`
	Permissions types.List                                                  `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                               `tfsdk:"port" json:"port,computed_optional"`
	Status      types.String                                                `tfsdk:"status" json:"status,computed"`
	Tunnel      customfield.NestedObject[CustomSSLKeylessServerTunnelModel] `tfsdk:"tunnel" json:"tunnel,computed_optional"`
}

type CustomSSLKeylessServerTunnelModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id"`
}
