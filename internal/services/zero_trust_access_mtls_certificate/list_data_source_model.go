// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_certificate

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSCertificatesResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessMTLSCertificatesResultDataSourceModel] `json:"result,computed"`
}

type ZeroTrustAccessMTLSCertificatesDataSourceModel struct {
	AccountID types.String                                                                       `tfsdk:"account_id" path:"account_id,optional"`
	ZoneID    types.String                                                                       `tfsdk:"zone_id" path:"zone_id,optional"`
	MaxItems  types.Int64                                                                        `tfsdk:"max_items"`
	Result    customfield.NestedObjectList[ZeroTrustAccessMTLSCertificatesResultDataSourceModel] `tfsdk:"result"`
}

func (m *ZeroTrustAccessMTLSCertificatesDataSourceModel) toListParams(_ context.Context) (params zero_trust.AccessCertificateListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCertificateListParams{}

	if !m.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessMTLSCertificatesResultDataSourceModel struct {
	ID                  types.String                   `tfsdk:"id" json:"id,computed"`
	AssociatedHostnames customfield.List[types.String] `tfsdk:"associated_hostnames" json:"associated_hostnames,computed"`
	CreatedAt           timetypes.RFC3339              `tfsdk:"created_at" json:"created_at,computed" format:"date-time"`
	ExpiresOn           timetypes.RFC3339              `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Fingerprint         types.String                   `tfsdk:"fingerprint" json:"fingerprint,computed"`
	Name                types.String                   `tfsdk:"name" json:"name,computed"`
	UpdatedAt           timetypes.RFC3339              `tfsdk:"updated_at" json:"updated_at,computed" format:"date-time"`
}
