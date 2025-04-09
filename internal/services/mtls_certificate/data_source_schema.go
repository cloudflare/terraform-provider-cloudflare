// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*MTLSCertificateDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"mtls_certificate_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ca": schema.BoolAttribute{
				Description: "Indicates whether the certificate is a CA or leaf certificate.",
				Computed:    true,
			},
			"certificates": schema.StringAttribute{
				Description: "The uploaded root CA certificate.",
				Computed:    true,
			},
			"expires_on": schema.StringAttribute{
				Description: "When the certificate expires.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"issuer": schema.StringAttribute{
				Description: "The certificate authority that issued the certificate.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Optional unique name for the certificate. Only used for human readability.",
				Computed:    true,
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
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (d *MTLSCertificateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MTLSCertificateDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
