// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomSSLResultEnvelope struct {
	Result CustomSSLModel `json:"result,computed"`
}

type CustomSSLModel struct {
	ZoneID              types.String                   `tfsdk:"zone_id" path:"zone_id"`
	CustomCertificateID types.String                   `tfsdk:"custom_certificate_id" path:"custom_certificate_id"`
	Type                types.String                   `tfsdk:"type" json:"type"`
	Certificate         types.String                   `tfsdk:"certificate" json:"certificate"`
	PrivateKey          types.String                   `tfsdk:"private_key" json:"private_key"`
	Policy              types.String                   `tfsdk:"policy" json:"policy"`
	GeoRestrictions     *CustomSSLGeoRestrictionsModel `tfsdk:"geo_restrictions" json:"geo_restrictions"`
	BundleMethod        types.String                   `tfsdk:"bundle_method" json:"bundle_method"`
	ID                  types.String                   `tfsdk:"id" json:"id,computed"`
}

type CustomSSLGeoRestrictionsModel struct {
	Label types.String `tfsdk:"label" json:"label"`
}
