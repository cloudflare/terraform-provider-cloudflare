// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultDataSourceEnvelope struct {
	Result CertificatePackDataSourceModel `json:"result,computed"`
}

type CertificatePackDataSourceModel struct {
	ID                   types.String                                                                  `tfsdk:"id" path:"certificate_pack_id,computed"`
	CertificatePackID    types.String                                                                  `tfsdk:"certificate_pack_id" path:"certificate_pack_id,optional"`
	ZoneID               types.String                                                                  `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateAuthority types.String                                                                  `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CloudflareBranding   types.Bool                                                                    `tfsdk:"cloudflare_branding" json:"cloudflare_branding,computed"`
	PrimaryCertificate   types.String                                                                  `tfsdk:"primary_certificate" json:"primary_certificate,computed"`
	Status               types.String                                                                  `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                                                  `tfsdk:"type" json:"type,computed"`
	ValidationMethod     types.String                                                                  `tfsdk:"validation_method" json:"validation_method,computed"`
	ValidityDays         types.Int64                                                                   `tfsdk:"validity_days" json:"validity_days,computed"`
	Hosts                customfield.List[types.String]                                                `tfsdk:"hosts" json:"hosts,computed"`
	Certificates         customfield.NestedObjectList[CertificatePackCertificatesDataSourceModel]      `tfsdk:"certificates" json:"certificates,computed"`
	ValidationErrors     customfield.NestedObjectList[CertificatePackValidationErrorsDataSourceModel]  `tfsdk:"validation_errors" json:"validation_errors,computed"`
	ValidationRecords    customfield.NestedObjectList[CertificatePackValidationRecordsDataSourceModel] `tfsdk:"validation_records" json:"validation_records,computed"`
	Filter               *CertificatePackFindOneByDataSourceModel                                      `tfsdk:"filter"`
}

func (m *CertificatePackDataSourceModel) toReadParams(_ context.Context) (params ssl.CertificatePackGetParams, diags diag.Diagnostics) {
	params = ssl.CertificatePackGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *CertificatePackDataSourceModel) toListParams(_ context.Context) (params ssl.CertificatePackListParams, diags diag.Diagnostics) {
	params = ssl.CertificatePackListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Filter.Status.IsNull() {
		params.Status = cloudflare.F(ssl.CertificatePackListParamsStatus(m.Filter.Status.ValueString()))
	}

	return
}

type CertificatePackCertificatesDataSourceModel struct {
	ID              types.String                                                                        `tfsdk:"id" json:"id,computed"`
	Hosts           customfield.List[types.String]                                                      `tfsdk:"hosts" json:"hosts,computed"`
	Status          types.String                                                                        `tfsdk:"status" json:"status,computed"`
	BundleMethod    types.String                                                                        `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       timetypes.RFC3339                                                                   `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	GeoRestrictions customfield.NestedObject[CertificatePackCertificatesGeoRestrictionsDataSourceModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed"`
	Issuer          types.String                                                                        `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                                                                   `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority        types.Float64                                                                       `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                                                        `tfsdk:"signature" json:"signature,computed"`
	UploadedOn      timetypes.RFC3339                                                                   `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	ZoneID          types.String                                                                        `tfsdk:"zone_id" json:"zone_id,computed"`
}

type CertificatePackCertificatesGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
}

type CertificatePackValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CertificatePackValidationRecordsDataSourceModel struct {
	Emails   customfield.List[types.String] `tfsdk:"emails" json:"emails,computed"`
	HTTPBody types.String                   `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String                   `tfsdk:"http_url" json:"http_url,computed"`
	TXTName  types.String                   `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue types.String                   `tfsdk:"txt_value" json:"txt_value,computed"`
}

type CertificatePackFindOneByDataSourceModel struct {
	Status types.String `tfsdk:"status" query:"status,optional"`
}
