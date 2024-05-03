// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificateResultEnvelope struct {
	Result MTLSCertificateModel `json:"result,computed"`
}

type MTLSCertificateModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	AccountID    types.String `tfsdk:"account_id" path:"account_id"`
	CA           types.Bool   `tfsdk:"ca" json:"ca"`
	Certificates types.String `tfsdk:"certificates" json:"certificates"`
	Name         types.String `tfsdk:"name" json:"name"`
	PrivateKey   types.String `tfsdk:"private_key" json:"private_key"`
}
