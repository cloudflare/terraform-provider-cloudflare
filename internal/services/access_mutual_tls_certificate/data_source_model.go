// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessMutualTLSCertificateResultDataSourceEnvelope struct {
	Result AccessMutualTLSCertificateDataSourceModel `json:"result,computed"`
}

type AccessMutualTLSCertificateResultListDataSourceEnvelope struct {
	Result *[]*AccessMutualTLSCertificateDataSourceModel `json:"result,computed"`
}

type AccessMutualTLSCertificateDataSourceModel struct {
	CertificateID       types.String `tfsdk:"certificate_id" path:"certificate_id"`
	AccountID           types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID              types.String `tfsdk:"zone_id" path:"zone_id"`
	ID                  types.String `tfsdk:"id" json:"id"`
	AssociatedHostnames types.String `tfsdk:"associated_hostnames" json:"associated_hostnames"`
	CreatedAt           types.String `tfsdk:"created_at" json:"created_at"`
	ExpiresOn           types.String `tfsdk:"expires_on" json:"expires_on"`
	Fingerprint         types.String `tfsdk:"fingerprint" json:"fingerprint"`
	Name                types.String `tfsdk:"name" json:"name"`
	UpdatedAt           types.String `tfsdk:"updated_at" json:"updated_at"`
}
