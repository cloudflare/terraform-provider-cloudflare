package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// V4Model represents the v4 schema structure for authenticated_origin_pulls_certificate
// This includes ALL fields from v4, including the 'type' field used for filtering
type V4Model struct {
	ID           types.String `tfsdk:"id"`
	ZoneID       types.String `tfsdk:"zone_id"`
	Certificate  types.String `tfsdk:"certificate"`
	PrivateKey   types.String `tfsdk:"private_key"`
	Type         types.String `tfsdk:"type"`
	Issuer       types.String `tfsdk:"issuer"`
	Signature    types.String `tfsdk:"signature"`
	SerialNumber types.String `tfsdk:"serial_number"`
	ExpiresOn    types.String `tfsdk:"expires_on"`
	Status       types.String `tfsdk:"status"`
	UploadedOn   types.String `tfsdk:"uploaded_on"`
}

// V5Model represents the v5 schema structure for per-zone certificates
// Note: Does NOT include 'type' field (removed in v5)
type V5Model struct {
	ZoneID        types.String `tfsdk:"zone_id"`
	CertificateID types.String `tfsdk:"certificate_id"`
	Certificate   types.String `tfsdk:"certificate"`
	PrivateKey    types.String `tfsdk:"private_key"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	ExpiresOn     types.String `tfsdk:"expires_on"`
	ID            types.String `tfsdk:"id"`
	Issuer        types.String `tfsdk:"issuer"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	Signature     types.String `tfsdk:"signature"`
	Status        types.String `tfsdk:"status"`
	UploadedOn    types.String `tfsdk:"uploaded_on"`
}
