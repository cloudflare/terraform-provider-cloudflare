package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// V4Model represents the v4 schema structure for cloudflare_authenticated_origin_pulls_certificate
// with type="per-hostname". This is the source schema for MoveState operations.
type V4Model struct {
	ID           types.String      `tfsdk:"id"`
	ZoneID       types.String      `tfsdk:"zone_id"`
	Certificate  types.String      `tfsdk:"certificate"`
	PrivateKey   types.String      `tfsdk:"private_key"`
	Type         types.String      `tfsdk:"type"` // This field is removed in v5
	Issuer       types.String      `tfsdk:"issuer"`
	Signature    types.String      `tfsdk:"signature"`
	SerialNumber types.String      `tfsdk:"serial_number"`
	ExpiresOn    timetypes.RFC3339 `tfsdk:"expires_on"`
	Status       types.String      `tfsdk:"status"`
	UploadedOn   timetypes.RFC3339 `tfsdk:"uploaded_on"`
}

// V5Model represents the v5 schema structure for authenticated_origin_pulls_hostname_certificate
// Note: Does NOT include 'type' field (removed in v5)
type V5Model struct {
	ID           types.String      `tfsdk:"id"`
	ZoneID       types.String      `tfsdk:"zone_id"`
	Certificate  types.String      `tfsdk:"certificate"`
	PrivateKey   types.String      `tfsdk:"private_key"`
	Issuer       types.String      `tfsdk:"issuer"`
	Signature    types.String      `tfsdk:"signature"`
	SerialNumber types.String      `tfsdk:"serial_number"`
	ExpiresOn    timetypes.RFC3339 `tfsdk:"expires_on"`
	Status       types.String      `tfsdk:"status"`
	UploadedOn   timetypes.RFC3339 `tfsdk:"uploaded_on"`
}
