// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessMutualTLSCertificateResultEnvelope struct {
	Result AccessMutualTLSCertificateModel `json:"result,computed"`
}

type AccessMutualTLSCertificateModel struct {
	ID                  types.String      `tfsdk:"id" json:"id,computed"`
	AccountID           types.String      `tfsdk:"account_id" path:"account_id"`
	ZoneID              types.String      `tfsdk:"zone_id" path:"zone_id"`
	Certificate         types.String      `tfsdk:"certificate" json:"certificate"`
	Name                types.String      `tfsdk:"name" json:"name"`
	AssociatedHostnames *[]types.String   `tfsdk:"associated_hostnames" json:"associated_hostnames"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"created_at,computed"`
	ExpiresOn           timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed"`
	Fingerprint         types.String      `tfsdk:"fingerprint" json:"fingerprint,computed"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at,computed"`
}
