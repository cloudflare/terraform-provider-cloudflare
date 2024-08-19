// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSCertificateResultDataSourceEnvelope struct {
	Result ZeroTrustAccessMTLSCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessMTLSCertificateResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustAccessMTLSCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessMTLSCertificateDataSourceModel struct {
	AccountID           types.String                                            `tfsdk:"account_id" path:"account_id"`
	CertificateID       types.String                                            `tfsdk:"certificate_id" path:"certificate_id"`
	ZoneID              types.String                                            `tfsdk:"zone_id" path:"zone_id"`
	CreatedAt           timetypes.RFC3339                                       `tfsdk:"created_at" json:"created_at,computed"`
	UpdatedAt           timetypes.RFC3339                                       `tfsdk:"updated_at" json:"updated_at,computed"`
	ExpiresOn           timetypes.RFC3339                                       `tfsdk:"expires_on" json:"expires_on"`
	Fingerprint         types.String                                            `tfsdk:"fingerprint" json:"fingerprint"`
	ID                  types.String                                            `tfsdk:"id" json:"id"`
	Name                types.String                                            `tfsdk:"name" json:"name"`
	AssociatedHostnames *[]types.String                                         `tfsdk:"associated_hostnames" json:"associated_hostnames"`
	Filter              *ZeroTrustAccessMTLSCertificateFindOneByDataSourceModel `tfsdk:"filter"`
}

type ZeroTrustAccessMTLSCertificateFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
