// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type MTLSCertificatesResultListDataSourceEnvelope struct {
	Result *[]*MTLSCertificatesItemsDataSourceModel `json:"result,computed"`
}

type MTLSCertificatesDataSourceModel struct {
	AccountID types.String                             `tfsdk:"account_id" path:"account_id"`
	MaxItems  types.Int64                              `tfsdk:"max_items"`
	Items     *[]*MTLSCertificatesItemsDataSourceModel `tfsdk:"items"`
}

type MTLSCertificatesItemsDataSourceModel struct {
	ID           types.String `tfsdk:"id" json:"id,computed"`
	CA           types.Bool   `tfsdk:"ca" json:"ca"`
	Certificates types.String `tfsdk:"certificates" json:"certificates"`
	ExpiresOn    types.String `tfsdk:"expires_on" json:"expires_on,computed"`
	Issuer       types.String `tfsdk:"issuer" json:"issuer,computed"`
	Name         types.String `tfsdk:"name" json:"name"`
	SerialNumber types.String `tfsdk:"serial_number" json:"serial_number,computed"`
	Signature    types.String `tfsdk:"signature" json:"signature,computed"`
	UploadedOn   types.String `tfsdk:"uploaded_on" json:"uploaded_on"`
}
