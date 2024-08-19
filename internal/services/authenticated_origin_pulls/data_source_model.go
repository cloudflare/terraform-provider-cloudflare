// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package authenticated_origin_pulls

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
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
	CERTUpdatedAt  timetypes.RFC3339 `tfsdk:"cert_updated_at" json:"cert_updated_at"`
	CERTUploadedOn timetypes.RFC3339 `tfsdk:"cert_uploaded_on" json:"cert_uploaded_on"`
	Certificate    types.String      `tfsdk:"certificate" json:"certificate"`
	CreatedAt      timetypes.RFC3339 `tfsdk:"created_at" json:"created_at"`
	Enabled        types.Bool        `tfsdk:"enabled" json:"enabled"`
	ExpiresOn      timetypes.RFC3339 `tfsdk:"expires_on" json:"expires_on"`
	Issuer         types.String      `tfsdk:"issuer" json:"issuer"`
	SerialNumber   types.String      `tfsdk:"serial_number" json:"serial_number"`
	Signature      types.String      `tfsdk:"signature" json:"signature"`
	Status         types.String      `tfsdk:"status" json:"status"`
	UpdatedAt      timetypes.RFC3339 `tfsdk:"updated_at" json:"updated_at"`
}
