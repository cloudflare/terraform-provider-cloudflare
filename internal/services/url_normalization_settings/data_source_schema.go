// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &URLNormalizationSettingsDataSource{}

func (d *URLNormalizationSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the URL normalization.",
				Optional:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of URL normalization performed by Cloudflare.",
				Optional:    true,
			},
		},
	}
}

func (d *URLNormalizationSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
