// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneHoldDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Access: Apps and Policies Read",
				"Access: Apps and Policies Revoke",
				"Access: Apps and Policies Write",
				"Access: Mutual TLS Certificates Write",
				"Access: Organizations, Identity Providers, and Groups Write",
				"Analytics Read",
				"Apps Write",
				"Cache Purge",
				"DNS Read",
				"DNS Write",
				"Firewall Services Read",
				"Firewall Services Write",
				"Load Balancers Read",
				"Load Balancers Write",
				"Logs Read",
				"Logs Write",
				"Page Rules Read",
				"Page Rules Write",
				"SSL and Certificates Read",
				"SSL and Certificates Write",
				"Stream Read",
				"Stream Write",
				"Trust and Safety Read",
				"Trust and Safety Write",
				"Workers Routes Read",
				"Workers Routes Write",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Zaraz Admin",
				"Zaraz Edit",
				"Zaraz Read",
				"Zero Trust: PII Read",
				"Zone Read",
				"Zone Settings Read",
				"Zone Settings Write",
				"Zone Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"hold": schema.BoolAttribute{
				Computed: true,
			},
			"hold_after": schema.StringAttribute{
				Computed: true,
			},
			"include_subdomains": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *ZoneHoldDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZoneHoldDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
