package origin_ca_certificate

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudflareOriginCACertificateModel struct {
	ID          types.String `tfsdk:"id"`
	Certificate types.String `tfsdk:"certificate"`
	Hostnames   types.List   `tfsdk:"hostnames"`
	ExpiresOn   types.String `tfsdk:"expires_on"`
	RequestType types.String `tfsdk:"request_type"`
	RevokedAt   types.String `tfsdk:"revoked_at"`
}
