// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomHostnamesResultEnvelope struct {
	Result CustomHostnamesModel `json:"result,computed"`
}

type CustomHostnamesModel struct {
	ID             types.String                        `tfsdk:"id" json:"id,computed"`
	ZoneID         types.String                        `tfsdk:"zone_id" path:"zone_id"`
	Hostname       types.String                        `tfsdk:"hostname" json:"hostname"`
	SSL            *CustomHostnamesSSLModel            `tfsdk:"ssl" json:"ssl"`
	CustomMetadata *CustomHostnamesCustomMetadataModel `tfsdk:"custom_metadata" json:"custom_metadata"`
}

type CustomHostnamesSSLModel struct {
	BundleMethod         types.String                     `tfsdk:"bundle_method" json:"bundle_method"`
	CertificateAuthority types.String                     `tfsdk:"certificate_authority" json:"certificate_authority"`
	CustomCertificate    types.String                     `tfsdk:"custom_certificate" json:"custom_certificate"`
	CustomKey            types.String                     `tfsdk:"custom_key" json:"custom_key"`
	Method               types.String                     `tfsdk:"method" json:"method"`
	Settings             *CustomHostnamesSSLSettingsModel `tfsdk:"settings" json:"settings"`
	Type                 types.String                     `tfsdk:"type" json:"type"`
	Wildcard             types.Bool                       `tfsdk:"wildcard" json:"wildcard"`
}

type CustomHostnamesSSLSettingsModel struct {
	Ciphers       *[]types.String `tfsdk:"ciphers" json:"ciphers"`
	EarlyHints    types.String    `tfsdk:"early_hints" json:"early_hints"`
	HTTP2         types.String    `tfsdk:"http2" json:"http2"`
	MinTLSVersion types.String    `tfsdk:"min_tls_version" json:"min_tls_version"`
	TLS1_3        types.String    `tfsdk:"tls_1_3" json:"tls_1_3"`
}

type CustomHostnamesCustomMetadataModel struct {
	Key types.String `tfsdk:"key" json:"key"`
}
