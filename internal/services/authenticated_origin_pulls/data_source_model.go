// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/origin_tls_client_auth"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AuthenticatedOriginPullsResultDataSourceEnvelope struct {
	Result AuthenticatedOriginPullsDataSourceModel `json:"result,computed"`
}

type AuthenticatedOriginPullsDataSourceModel struct {
	ZoneID         types.String      `tfsdk:"zone_id" path:"zone_id"`
	Hostname       types.String      `tfsdk:"hostname" path:"hostname,computed"`
	CERTID         types.String      `tfsdk:"cert_id" json:"cert_id"`
	CERTStatus     types.String      `tfsdk:"cert_status" json:"cert_status"`
	CERTUpdatedAt  timetypes.RFC3339 `tfsdk:"cert_updated_at" json:"cert_updated_at" format:"date-time"`
	CERTUploadedOn timetypes.RFC3339 `tfsdk:"cert_uploaded_on" json:"cert_uploaded_on" format:"date-time"`
	Certificate    types.String      `tfsdk:"certificate" json:"certificate"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at" format:"date-time"`
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled"`
	ExpiresOn      timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on" format:"date-time"`
	Issuer         types.String      `tfsdk:"issuer" json:"issuer"`
	SerialNumber   types.String      `tfsdk:"serial_number" json:"serial_number"`
	Signature      types.String      `tfsdk:"signature" json:"signature"`
	Status         types.String      `tfsdk:"status" json:"status"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at" format:"date-time"`
}

func (m *AuthenticatedOriginPullsDataSourceModel) toReadParams(_ context.Context) (params origin_tls_client_auth.HostnameGetParams, diags diag.Diagnostics) {
	params = origin_tls_client_auth.HostnameGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}
