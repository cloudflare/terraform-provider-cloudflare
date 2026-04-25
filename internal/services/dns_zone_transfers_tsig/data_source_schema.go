// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_zone_transfers_tsig

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*DNSZoneTransfersTSIGDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Settings Read",
				"Account Settings Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"tsig_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"algo": schema.StringAttribute{
				Description: "TSIG algorithm.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "TSIG key name.",
				Computed:    true,
			},
			"secret": schema.StringAttribute{
				Description: "TSIG secret.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (d *DNSZoneTransfersTSIGDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *DNSZoneTransfersTSIGDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
