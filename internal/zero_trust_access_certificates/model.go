// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_certificates

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessCertificatesResultEnvelope struct {
	Result ZeroTrustAccessCertificatesModel `json:"result,computed"`
}

type ZeroTrustAccessCertificatesModel struct {
	AccountID           types.String   `tfsdk:"account_id" path:"account_id"`
	ZoneID              types.String   `tfsdk:"zone_id" path:"zone_id"`
	UUID                types.String   `tfsdk:"uuid" path:"uuid"`
	Certificate         types.String   `tfsdk:"certificate" json:"certificate"`
	Name                types.String   `tfsdk:"name" json:"name"`
	AssociatedHostnames []types.String `tfsdk:"associated_hostnames" json:"associated_hostnames"`
	ID                  types.String   `tfsdk:"id" json:"id"`
	CreatedAt           types.String   `tfsdk:"created_at" json:"created_at,computed"`
	ExpiresOn           types.String   `tfsdk:"expires_on" json:"expires_on,computed"`
	Fingerprint         types.String   `tfsdk:"fingerprint" json:"fingerprint"`
	UpdatedAt           types.String   `tfsdk:"updated_at" json:"updated_at,computed"`
}
