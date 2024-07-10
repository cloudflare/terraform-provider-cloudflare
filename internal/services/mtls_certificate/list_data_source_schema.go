// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &MTLSCertificatesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &MTLSCertificatesDataSource{}

func (r MTLSCertificatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"ca": schema.BoolAttribute{
							Description: "Indicates whether the certificate is a CA or leaf certificate.",
							Computed:    true,
							Optional:    true,
						},
						"certificates": schema.StringAttribute{
							Description: "The uploaded root CA certificate.",
							Computed:    true,
							Optional:    true,
						},
						"expires_on": schema.StringAttribute{
							Description: "When the certificate expires.",
							Computed:    true,
						},
						"issuer": schema.StringAttribute{
							Description: "The certificate authority that issued the certificate.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Optional unique name for the certificate. Only used for human readability.",
							Computed:    true,
							Optional:    true,
						},
						"serial_number": schema.StringAttribute{
							Description: "The certificate serial number.",
							Computed:    true,
						},
						"signature": schema.StringAttribute{
							Description: "The type of hash used for the certificate.",
							Computed:    true,
						},
						"uploaded_on": schema.StringAttribute{
							Description: "This is the time the certificate was uploaded.",
							Computed:    true,
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *MTLSCertificatesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *MTLSCertificatesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
