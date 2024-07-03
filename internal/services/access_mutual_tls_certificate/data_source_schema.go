// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_mutual_tls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessMutualTLSCertificateDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessMutualTLSCertificateDataSource{}

func (r AccessMutualTLSCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"certificate_id": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The ID of the application that will use this certificate.",
				Optional:    true,
			},
			"associated_hostnames": schema.StringAttribute{
				Description: "The hostnames of the applications that will use this certificate.",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Optional: true,
			},
			"expires_on": schema.StringAttribute{
				Optional: true,
			},
			"fingerprint": schema.StringAttribute{
				Description: "The MD5 fingerprint of the certificate.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the certificate.",
				Optional:    true,
			},
			"updated_at": schema.StringAttribute{
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
						Optional:    true,
					},
					"zone_id": schema.StringAttribute{
						Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *AccessMutualTLSCertificateDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessMutualTLSCertificateDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
