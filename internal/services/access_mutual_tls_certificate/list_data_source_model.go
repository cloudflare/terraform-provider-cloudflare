// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AccessMutualTLSCertificatesResultListDataSourceEnvelope struct {
	Result *[]*AccessMutualTLSCertificatesItemsDataSourceModel `json:"result,computed"`
}

type AccessMutualTLSCertificatesDataSourceModel struct {
	AccountID types.String                                        `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String                                        `tfsdk:"zone_id" path:"zone_id"`
	MaxItems  types.Int64                                         `tfsdk:"max_items"`
	Items     *[]*AccessMutualTLSCertificatesItemsDataSourceModel `tfsdk:"items"`
}

type AccessMutualTLSCertificatesItemsDataSourceModel struct {
	ID                  types.String      `tfsdk:"id" json:"id"`
	AssociatedHostnames *[]types.String   `tfsdk:"associated_hostnames" json:"associated_hostnames"`
	CreatedAt           timetypes.RFC3339 `tfsdk:"created_at" json:"created_at"`
	ExpiresOn           timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on"`
	Fingerprint         types.String      `tfsdk:"fingerprint" json:"fingerprint"`
	Name                types.String      `tfsdk:"name" json:"name"`
	UpdatedAt           timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at"`
}
