// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_mtls_certificate

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustAccessMTLSCertificateResultDataSourceEnvelope struct {
	Result ZeroTrustAccessMTLSCertificateDataSourceModel `json:"result,computed"`
}

type ZeroTrustAccessMTLSCertificateResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[ZeroTrustAccessMTLSCertificateDataSourceModel] `json:"result,computed"`
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

func (m *ZeroTrustAccessMTLSCertificateDataSourceModel) toReadParams() (params zero_trust.AccessCertificateGetParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCertificateGetParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

func (m *ZeroTrustAccessMTLSCertificateDataSourceModel) toListParams() (params zero_trust.AccessCertificateListParams, diags diag.Diagnostics) {
	params = zero_trust.AccessCertificateListParams{}

	if !m.Filter.AccountID.IsNull() {
		params.AccountID = cloudflare.F(m.Filter.AccountID.ValueString())
	} else {
		params.ZoneID = cloudflare.F(m.Filter.ZoneID.ValueString())
	}

	return
}

type ZeroTrustAccessMTLSCertificateFindOneByDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ZoneID    types.String `tfsdk:"zone_id" path:"zone_id"`
}
