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

type CertificatePacksResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[CertificatePacksResultDataSourceModel] `json:"result,computed"`
}

type CertificatePacksDataSourceModel struct {
	ZoneID   types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	Status   types.String                                                        `tfsdk:"status" query:"status,optional"`
	MaxItems types.Int64                                                         `tfsdk:"max_items"`
	Result   customfield.NestedObjectList[CertificatePacksResultDataSourceModel] `tfsdk:"result"`
}

func (m *CertificatePacksDataSourceModel) toListParams(_ context.Context) (params ssl.CertificatePackListParams, diags diag.Diagnostics) {
	params = ssl.CertificatePackListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	if !m.Status.IsNull() {
		params.Status = cloudflare.F(ssl.CertificatePackListParamsStatus(m.Status.ValueString()))
	}

	return
}

type CertificatePacksResultDataSourceModel struct {
	ID                   types.String                                                                   `tfsdk:"id" json:"id,computed"`
	Certificates         customfield.NestedObjectList[CertificatePacksCertificatesDataSourceModel]      `tfsdk:"certificates" json:"certificates,computed"`
	Hosts                customfield.Set[types.String]                                                  `tfsdk:"hosts" json:"hosts,computed"`
	Status               types.String                                                                   `tfsdk:"status" json:"status,computed"`
	Type                 types.String                                                                   `tfsdk:"type" json:"type,computed"`
	CertificateAuthority types.String                                                                   `tfsdk:"certificate_authority" json:"certificate_authority,computed"`
	CloudflareBranding   types.Bool                                                                     `tfsdk:"cloudflare_branding" json:"cloudflare_branding,computed"`
	PrimaryCertificate   types.String                                                                   `tfsdk:"primary_certificate" json:"primary_certificate,computed"`
	ValidationErrors     customfield.NestedObjectList[CertificatePacksValidationErrorsDataSourceModel]  `tfsdk:"validation_errors" json:"validation_errors,computed"`
	ValidationMethod     types.String                                                                   `tfsdk:"validation_method" json:"validation_method,computed"`
	ValidationRecords    customfield.NestedObjectList[CertificatePacksValidationRecordsDataSourceModel] `tfsdk:"validation_records" json:"validation_records,computed"`
	ValidityDays         types.Int64                                                                    `tfsdk:"validity_days" json:"validity_days,computed"`
}

type CertificatePacksCertificatesDataSourceModel struct {
	ID              types.String                                                                         `tfsdk:"id" json:"id,computed"`
	Hosts           customfield.List[types.String]                                                       `tfsdk:"hosts" json:"hosts,computed"`
	Status          types.String                                                                         `tfsdk:"status" json:"status,computed"`
	BundleMethod    types.String                                                                         `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       timetypes.RFC3339                                                                    `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	GeoRestrictions customfield.NestedObject[CertificatePacksCertificatesGeoRestrictionsDataSourceModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed"`
	Issuer          types.String                                                                         `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                                                                    `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority        types.Float64                                                                        `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                                                         `tfsdk:"signature" json:"signature,computed"`
	UploadedOn      timetypes.RFC3339                                                                    `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	ZoneID          types.String                                                                         `tfsdk:"zone_id" json:"zone_id,computed"`
}

type CertificatePacksCertificatesGeoRestrictionsDataSourceModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
}

type CertificatePacksValidationErrorsDataSourceModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CertificatePacksValidationRecordsDataSourceModel struct {
	Emails   customfield.List[types.String] `tfsdk:"emails" json:"emails,computed"`
	HTTPBody types.String                   `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String                   `tfsdk:"http_url" json:"http_url,computed"`
	TXTName  types.String                   `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue types.String                   `tfsdk:"txt_value" json:"txt_value,computed"`
}
