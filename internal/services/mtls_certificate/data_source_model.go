// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificateResultDataSourceEnvelope struct {
	Result MTLSCertificateDataSourceModel `json:"result,computed"`
}

type MTLSCertificateResultListDataSourceEnvelope struct {
	Result *[]*MTLSCertificateDataSourceModel `json:"result,computed"`
}

type MTLSCertificateDataSourceModel struct {
	AccountID         types.String                             `tfsdk:"account_id" path:"account_id"`
	MTLSCertificateID types.String                             `tfsdk:"mtls_certificate_id" path:"mtls_certificate_id"`
	ExpiresOn         timetypes.RFC3339                        `tfsdk:"expires_on" json:"expires_on,computed"`
	Issuer            types.String                             `tfsdk:"issuer" json:"issuer,computed"`
	SerialNumber      types.String                             `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature         types.String                             `tfsdk:"signature" json:"signature,computed"`
	CA                types.Bool                               `tfsdk:"ca" json:"ca"`
	Certificates      types.String                             `tfsdk:"certificates" json:"certificates"`
	ID                types.String                             `tfsdk:"id" json:"id"`
	Name              types.String                             `tfsdk:"name" json:"name"`
	UploadedOn        timetypes.RFC3339                        `tfsdk:"uploaded_on" json:"uploaded_on"`
	Filter            *MTLSCertificateFindOneByDataSourceModel `tfsdk:"filter"`
}

type MTLSCertificateFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
}
