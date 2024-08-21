// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/origin_tls_client_auth"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificatesResultListDataSourceEnvelope struct {
	Result *[]*AuthenticatedOriginPullsCertificatesResultDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificatesDataSourceModel struct {
	ZoneID   types.String                                                  `tfsdk:"zone_id" path:"zone_id"`
	MaxItems types.Int64                                                   `tfsdk:"max_items"`
	Result   *[]*AuthenticatedOriginPullsCertificatesResultDataSourceModel `tfsdk:"result"`
}

func (m *AuthenticatedOriginPullsCertificatesDataSourceModel) toListParams() (params origin_tls_client_auth.OriginTLSClientAuthListParams, diags diag.Diagnostics) {
	params = origin_tls_client_auth.OriginTLSClientAuthListParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

type AuthenticatedOriginPullsCertificatesResultDataSourceModel struct {
	ID          types.String      `tfsdk:"id" json:"id"`
	Certificate types.String      `tfsdk:"certificate" json:"certificate"`
	ExpiresOn   timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on,computed"`
	Issuer      types.String      `tfsdk:"issuer" json:"issuer,computed"`
	Signature   types.String      `tfsdk:"signature" json:"signature,computed"`
	Status      types.String      `tfsdk:"status" json:"status"`
	UploadedOn  timetypes.RFC3339 `tfsdk:"uploaded_on" json:"uploaded_on"`
}
