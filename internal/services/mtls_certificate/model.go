// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificateResultEnvelope struct {
	Result MTLSCertificateModel `json:"result,computed"`
}

type MTLSCertificateResultDataSourceEnvelope struct {
	Result MTLSCertificateDataSourceModel `json:"result,computed"`
}

type MTLSCertificatesResultDataSourceEnvelope struct {
	Result MTLSCertificatesDataSourceModel `json:"result,computed"`
}

type MTLSCertificateModel struct {
	ID           types.String `tfsdk:"id" json:"id"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	CA           types.Bool   `tfsdk:"ca" json:"ca"`
	Certificates types.String `tfsdk:"certificates" json:"certificates"`
	Name         types.String `tfsdk:"name" json:"name"`
	PrivateKey   types.String `tfsdk:"private_key" json:"private_key"`
	ExpiresOn    types.String `tfsdk:"expires_on" json:"expires_on,computed"`
	Issuer       types.String `tfsdk:"issuer" json:"issuer,computed"`
	SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature    types.String `tfsdk:"signature" json:"signature,computed"`
	UpdatedAt    types.String `tfsdk:"updated_at" json:"updated_at,computed"`
	UploadedOn   types.String `tfsdk:"uploaded_on" json:"uploaded_on,computed"`
}

type MTLSCertificateDataSourceModel struct {
}

type MTLSCertificatesDataSourceModel struct {
}
