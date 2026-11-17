// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*PageShieldSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Domain Page Shield",
				"Domain Page Shield Read",
				"Page Shield",
				"Page Shield Read",
				"Zone Settings Read",
				"Zone Settings Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "When true, indicates that Page Shield is enabled. Page Shield must be enabled for continuous script monitoring, malicious script detection, and code change detection to work.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The timestamp of when the setting was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"use_cloudflare_reporting_endpoint": schema.BoolAttribute{
				Description: "When true, CSP reports will be sent to the Cloudflare-managed reporting endpoint (https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report) so detected resources appear in the Page Shield dashboard.",
				Computed:    true,
			},
			"use_connection_url_path": schema.BoolAttribute{
				Description: "When true, the paths associated with connection URLs will also be analyzed.",
				Computed:    true,
			},
		},
	}
}

func (d *PageShieldSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PageShieldSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
