// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewayCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identify the certificate with a UUID.",
				Computed:    true,
			},
			"certificate_id": schema.StringAttribute{
				Description: "Identify the certificate with a UUID.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"binding_status": schema.StringAttribute{
				Description: "Indicate the read-only deployment status of the certificate on Cloudflare's edge. Gateway TLS interception can use certificates in the 'available' (previously called 'active') state.\nAvailable values: \"pending_deployment\", \"available\", \"pending_deletion\", \"inactive\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"pending_deployment",
						"available",
						"pending_deletion",
						"inactive",
					),
				},
			},
			"certificate": schema.StringAttribute{
				Description: "Provide the CA certificate (read-only).",
				Computed:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"expires_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"fingerprint": schema.StringAttribute{
				Description: "Provide the SHA256 fingerprint of the certificate (read-only).",
				Computed:    true,
			},
			"in_use": schema.BoolAttribute{
				Description: "Indicate whether Gateway TLS interception uses this certificate (read-only). You cannot set this value directly. To configure interception, use the Gateway configuration setting named `certificate` (read-only).",
				Computed:    true,
			},
			"issuer_org": schema.StringAttribute{
				Description: "Indicate the organization that issued the certificate (read-only).",
				Computed:    true,
			},
			"issuer_raw": schema.StringAttribute{
				Description: "Provide the entire issuer field of the certificate (read-only).",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "Indicate the read-only certificate type, BYO-PKI (custom) or Gateway-managed.\nAvailable values: \"custom\", \"gateway_managed\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("custom", "gateway_managed"),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"uploaded_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (d *ZeroTrustGatewayCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustGatewayCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
