// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificate_associations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*MTLSCertificateAssociationsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"mtls_certificate_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"service": schema.StringAttribute{
				Description: "The service using the certificate.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Certificate deployment status for the given service.",
				Computed:    true,
			},
		},
	}
}

func (d *MTLSCertificateAssociationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MTLSCertificateAssociationsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
