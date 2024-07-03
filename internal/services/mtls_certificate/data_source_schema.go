// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &MTLSCertificateDataSource{}
var _ datasource.DataSourceWithValidateConfig = &MTLSCertificateDataSource{}

func (r MTLSCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"mtls_certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"ca": schema.BoolAttribute{
				Description: "Indicates whether the certificate is a CA or leaf certificate.",
				Optional:    true,
			},
			"certificates": schema.StringAttribute{
				Description: "The uploaded root CA certificate.",
				Optional:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate expires.",
				Optional:    true,
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Optional unique name for the certificate. Only used for human readability.",
				Optional:    true,
			},
			"serial_number": schema.StringAttribute{
				Description: "The certificate serial number.",
				Optional:    true,
			},
			"signature": schema.StringAttribute{
				Description: "The type of hash used for the certificate.",
				Optional:    true,
			},
			"uploaded_on": schema.StringAttribute{
				Description: "This is the time the certificate was uploaded.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *MTLSCertificateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *MTLSCertificateDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
