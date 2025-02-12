// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/custom_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultDataSourceEnvelope struct {
	Result CustomSSLDataSourceModel `json:"result,computed"`
}

type CustomSSLDataSourceModel struct {
	ID                  types.String                                                      `tfsdk:"id" json:"-,computed"`
	CustomCertificateID types.String                                                      `tfsdk:"custom_certificate_id" path:"custom_certificate_id,optional"`
	ZoneID              types.String                                                      `tfsdk:"zone_id" path:"zone_id,computed"`
	BundleMethod        types.String                                                      `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn           timetypes.RFC3339                                                 `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer              types.String                                                      `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn          timetypes.RFC3339                                                 `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Policy              types.String                                                      `tfsdk:"policy" json:"policy,computed"`
	Priority            types.Float64                                                     `tfsdk:"priority" json:"priority,computed"`
	Signature           types.String                                                      `tfsdk:"signature" json:"signature,computed"`
	Status              types.String                                                      `tfsdk:"status" json:"status,computed"`
	UploadedOn          timetypes.RFC3339                                                 `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	Hosts               customfield.List[types.String]                                    `tfsdk:"hosts" json:"hosts,computed"`
	GeoRestrictions     customfield.NestedObject[CustomSSLGeoRestrictionsDataSourceModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed"`
	KeylessServer       jsontypes.Normalized                                              `tfsdk:"keyless_server" json:"keyless_server,computed"`
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
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
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

type CustomSSLFindOneByDataSourceModel struct {
	Match  types.String `tfsdk:"match" query:"match,computed_optional"`
	Status types.String `tfsdk:"status" query:"status,optional"`
}
