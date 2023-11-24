package origin_ca_certificate

import "github.com/hashicorp/terraform-plugin-framework/types"

type CloudflareOriginCACertificateModel struct {
	ID                types.String `tfsdk:"id"`
	Certificate       types.String `tfsdk:"certificate"`
	CSR               types.String `tfsdk:"csr"`
	Hostnames         types.List   `tfsdk:"hostnames"`
	ExpiresOn         types.String `tfsdk:"expires_on"`
	RequestType       types.String `tfsdk:"request_type"`
	RequestedValidity types.Int64  `tfsdk:"requested_validity"`
	RevokedAt         types.String `tfsdk:"revoked_at"`
}
