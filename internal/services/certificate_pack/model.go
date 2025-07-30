// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package certificate_pack

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CertificatePackResultEnvelope struct {
	Result CertificatePackModel `json:"result"`
}

type CertificatePackModel struct {
	ID                   types.String                                                        `tfsdk:"id" json:"id,computed"`
	ZoneID               types.String                                                        `tfsdk:"zone_id" path:"zone_id,required"`
	CertificateAuthority types.String                                                        `tfsdk:"certificate_authority" json:"certificate_authority,required,no_refresh"`
	Type                 types.String                                                        `tfsdk:"type" json:"type,required,no_refresh"`
	ValidationMethod     types.String                                                        `tfsdk:"validation_method" json:"validation_method,required,no_refresh"`
	ValidityDays         types.Int64                                                         `tfsdk:"validity_days" json:"validity_days,required,no_refresh"`
	Hosts                *[]types.String                                                     `tfsdk:"hosts" json:"hosts,required,no_refresh"`
	CloudflareBranding   types.Bool                                                          `tfsdk:"cloudflare_branding" json:"cloudflare_branding,optional,no_refresh"`
	Status               types.String                                                        `tfsdk:"status" json:"status,computed,no_refresh"`
	ValidationErrors     customfield.NestedObjectList[CertificatePackValidationErrorsModel]  `tfsdk:"validation_errors" json:"validation_errors,computed,no_refresh"`
	ValidationRecords    customfield.NestedObjectList[CertificatePackValidationRecordsModel] `tfsdk:"validation_records" json:"validation_records,computed,no_refresh"`
}

func (m CertificatePackModel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(m)
}

func (m CertificatePackModel) MarshalJSONForUpdate(state CertificatePackModel) (data []byte, err error) {
	return apijson.MarshalForPatch(m, state)
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
