// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls_certificate

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/origin_tls_client_auth"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsCertificateResultDataSourceEnvelope struct {
	Result AuthenticatedOriginPullsCertificateDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificateResultListDataSourceEnvelope struct {
	Result customfield.NestedObjectList[AuthenticatedOriginPullsCertificateDataSourceModel] `json:"result,computed"`
}

type AuthenticatedOriginPullsCertificateDataSourceModel struct {
	CertificateID types.String                                                 `tfsdk:"certificate_id" path:"certificate_id"`
	ZoneID        types.String                                                 `tfsdk:"zone_id" path:"zone_id"`
	ExpiresOn     timetypes.RFC3339                                            `tfsdk:"expires_on" json:"expires_on,computed" format:"date-time"`
	Issuer        types.String                                                 `tfsdk:"issuer" json:"issuer,computed"`
	Signature     types.String                                                 `tfsdk:"signature" json:"signature,computed"`
	Certificate   types.String                                                 `tfsdk:"certificate" json:"certificate,computed_optional"`
	ID            types.String                                                 `tfsdk:"id" json:"id,computed_optional"`
	Status        types.String                                                 `tfsdk:"status" json:"status,computed_optional"`
	UploadedOn    timetypes.RFC3339                                            `tfsdk:"uploaded_on" json:"uploaded_on,computed_optional" format:"date-time"`
	Filter        *AuthenticatedOriginPullsCertificateFindOneByDataSourceModel `tfsdk:"filter"`
}

func (m *AuthenticatedOriginPullsCertificateDataSourceModel) toReadParams() (params origin_tls_client_auth.OriginTLSClientAuthGetParams, diags diag.Diagnostics) {
	params = origin_tls_client_auth.OriginTLSClientAuthGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}

func (m *AuthenticatedOriginPullsCertificateDataSourceModel) toListParams() (params origin_tls_client_auth.OriginTLSClientAuthListParams, diags diag.Diagnostics) {
	params = origin_tls_client_auth.OriginTLSClientAuthListParams{
		ZoneID: cloudflare.F(m.Filter.ZoneID.ValueString()),
	}

	return
}

type AuthenticatedOriginPullsCertificateFindOneByDataSourceModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
}
