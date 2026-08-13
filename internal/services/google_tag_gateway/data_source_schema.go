// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package google_tag_gateway

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*GoogleTagGatewayDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
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
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Enables or disables Google Tag Gateway for this zone.",
				Computed:    true,
			},
			"endpoint": schema.StringAttribute{
				Description: "Specifies the endpoint path for proxying Google Tag Manager requests. Use an absolute path starting with '/', with no nested paths and alphanumeric characters only (e.g. /metrics).",
				Computed:    true,
			},
			"hide_original_ip": schema.BoolAttribute{
				Description: "Hides the original client IP address from Google when enabled.",
				Computed:    true,
			},
			"measurement_id": schema.StringAttribute{
				Description: "Specify the Google Tag Manager container or measurement ID (e.g. GTM-XXXXXXX or G-XXXXXXXXXX).",
				Computed:    true,
			},
			"set_up_tag": schema.BoolAttribute{
				Description: "Set up the associated Google Tag on the zone automatically when enabled.",
				Computed:    true,
			},
		},
	}
}

func (d *GoogleTagGatewayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *GoogleTagGatewayDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
