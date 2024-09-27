// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/custom_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultDataSourceEnvelope struct {
	Result CustomSSLDataSourceModel `json:"result,computed"`
}

type CustomSSLResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CustomSSLDataSourceModel] `json:"result,computed"`
}

type CustomSSLDataSourceModel struct {
	CustomCertificateID types.String                                                      `tfsdk:"custom_certificate_id" path:"custom_certificate_id,optional"`
	ZoneID              types.String                                                      `tfsdk:"zone_id" path:"zone_id,computed_optional"`
	BundleMethod        types.String                                                      `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn           timetypes.RFC3339                                                 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	ID                  types.String                                                      `tfsdk:"id" json:"id,computed"`
	Issuer              types.String                                                      `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn          timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Policy              types.String                                                      `tfsdk:"policy" json:"policy,computed"`
	Priority            types.Float64                                                     `tfsdk:"priority" json:"priority,computed"`
	Signature           types.String                                                      `tfsdk:"signature" json:"signature,computed"`
	Status              types.String                                                      `tfsdk:"status" json:"status,computed"`
	UploadedOn          timetypes.RFC3339                                                 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	Hosts               customfield.List[types.String]                                    `tfsdk:"hosts" json:"hosts,computed"`
	GeoRestrictions     customfield.NestedObject[CustomSSLGeoRestrictionsDataSourceModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed"`
	KeylessServer       customfield.NestedObject[CustomSSLKeylessServerDataSourceModel]   `tfsdk:"keyless_server" json:"keyless_server,computed"`
	Filter              *CustomSSLFindOneByDataSourceModel                                `tfsdk:"filter"`
}

func (m *CustomSSLDataSourceModel) toReadParams(_ context.Context) (params custom_certificates.CustomCertificateGetParams, diags diag.Diagnostics) {
	params = custom_certificates.CustomCertificateGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *CustomSSLDataSourceModel) toListParams(_ context.Context) (params custom_certificates.CustomCertificateListParams, diags diag.Diagnostics) {
	params = custom_certificates.CustomCertificateListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	if !m.Filter.Match.IsNull() {
		params.Match = cloudflare.F(custom_certificates.CustomCertificateListParamsMatch(m.Filter.Match.ValueString()))
	}
	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(custom_certificates.CustomCertificateListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type CustomSSLGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
}

type CustomSSLKeylessServerDataSourceModel struct {
	ID          types.String                                                          `tfsdk:"id" json:"id,computed"`
	CreatedOn   timetypes.RFC3339                                                     `tfsdk:"created_on" json:"created_on,computed" format:"date-time"`
	Enabled     types.Bool                                                            `tfsdk:"enabled" json:"enabled,computed"`
	Host        types.String                                                          `tfsdk:"host" json:"host,computed"`
	ModifiedOn  timetypes.RFC3339                                                     `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Name        types.String                                                          `tfsdk:"name" json:"name,computed"`
	Permissions customfield.List[types.String]                                        `tfsdk:"permissions" json:"permissions,computed"`
	Port        types.Float64                                                         `tfsdk:"port" json:"port,computed"`
	Status      types.String                                                          `tfsdk:"status" json:"status,computed"`
	Tunnel      customfield.NestedObject[CustomSSLKeylessServerTunnelDataSourceModel] `tfsdk:"tunnel" json:"tunnel,computed"`
}

type CustomSSLKeylessServerTunnelDataSourceModel struct {
	PrivateIP types.String `tfsdk:"private_ip" json:"private_ip,computed"`
	VnetID    types.String `tfsdk:"vnet_id" json:"vnet_id,computed"`
}

type CustomSSLFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id,required"`
	Match  types.String `tfsdk:"match" query:"match,computed_optional"`
	Status types.String `tfsdk:"status" query:"status,optional"`
}
