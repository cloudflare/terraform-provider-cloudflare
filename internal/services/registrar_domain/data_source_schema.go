// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_domain

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*RegistrarDomainDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"domain_name": schema.StringAttribute{
				Description: "Fully qualified domain name (FQDN) including the extension\n(e.g., `example.com`, `mybrand.app`). The domain name uniquely\nidentifies a registration — the same domain cannot be registered\ntwice, making it a natural idempotency key for registration requests.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
		},
	}
}

func (d *RegistrarDomainDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *RegistrarDomainDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
