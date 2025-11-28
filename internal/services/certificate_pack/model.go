// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultEnvelope struct {
	Result CertificatePackModel `json:"result"`
}

type CertificatePackModel struct {
	ID                   types.String                                                        `tfsdk:"id" json:"id,computed"`
	ZoneID               types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateAuthority types.String                                                        `tfsdk:"certificate_authority" json:"certificate_authority,required"`
	Type                 types.String                                                        `tfsdk:"type" json:"type,required"`
	ValidationMethod     types.String                                                        `tfsdk:"validation_method" json:"validation_method,required"`
	ValidityDays         types.Int64                                                         `tfsdk:"validity_days" json:"validity_days,required"`
	Hosts                *[]types.String                                                     `tfsdk:"hosts" json:"hosts,required"`
	CloudflareBranding   types.Bool                                                          `tfsdk:"cloudflare_branding" json:"cloudflare_branding,optional"`
	PrimaryCertificate   types.String                                                        `tfsdk:"primary_certificate" json:"primary_certificate,computed"`
	Status               types.String                                                        `tfsdk:"status" json:"status,computed"`
	Certificates         customfield.NestedObjectList[CertificatePackCertificatesModel]      `tfsdk:"certificates" json:"certificates,computed"`
	ValidationErrors     customfield.NestedObjectList[CertificatePackValidationErrorsModel]  `tfsdk:"validation_errors" json:"validation_errors,computed"`
	ValidationRecords    customfield.NestedObjectList[CertificatePackValidationRecordsModel] `tfsdk:"validation_records" json:"validation_records,computed"`
}

func (m CertificatePackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CertificatePackModel) MarshalJSONForUpdate(state CertificatePackModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
}

type CertificatePackCertificatesModel struct {
	ID              types.String                                                              `tfsdk:"id" json:"id,computed"`
	Hosts           customfield.List[types.String]                                            `tfsdk:"hosts" json:"hosts,computed"`
	Status          types.String                                                              `tfsdk:"status" json:"status,computed"`
	BundleMethod    types.String                                                              `tfsdk:"bundle_method" json:"bundle_method,computed"`
	ExpiresOn       timetypes.RFC3339                                                         `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	GeoRestrictions customfield.NestedObject[CertificatePackCertificatesGeoRestrictionsModel] `tfsdk:"geo_restrictions" json:"geo_restrictions,computed"`
	Issuer          types.String                                                              `tfsdk:"issuer" json:"issuer,computed"`
	ModifiedOn      timetypes.RFC3339                                                         `tfsdk:"modified_on" json:"modified_on,computed" format:"date-time"`
	Priority        types.Float64                                                             `tfsdk:"priority" json:"priority,computed"`
	Signature       types.String                                                              `tfsdk:"signature" json:"signature,computed"`
	UploadedOn      timetypes.RFC3339                                                         `tfsdk:"uploaded_on" json:"uploaded_on,computed" format:"date-time"`
	ZoneID          types.String                                                              `tfsdk:"zone_id" json:"zone_id,computed"`
}

type CertificatePackCertificatesGeoRestrictionsModel struct {
	Label types.String `tfsdk:"label" json:"label,computed"`
}

type CertificatePackValidationErrorsModel struct {
	Message types.String `tfsdk:"message" json:"message,computed"`
}

type CertificatePackValidationRecordsModel struct {
	Emails   customfield.List[types.String] `tfsdk:"emails" json:"emails,computed"`
	HTTPBody types.String                   `tfsdk:"http_body" json:"http_body,computed"`
	HTTPURL  types.String                   `tfsdk:"http_url" json:"http_url,computed"`
	TXTName  types.String                   `tfsdk:"txt_name" json:"txt_name,computed"`
	TXTValue types.String                   `tfsdk:"txt_value" json:"txt_value,computed"`
}
