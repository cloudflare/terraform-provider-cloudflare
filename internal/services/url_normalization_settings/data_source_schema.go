// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*URLNormalizationSettingsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Rulesets Read",
				"Account Rulesets Write",
				"Account WAF Read",
				"Account WAF Write",
				"Bot Management Read",
				"Bot Management Write",
				"Cache Settings Read",
				"Cache Settings Write",
				"Config Settings Read",
				"Config Settings Write",
				"Custom Errors Read",
				"Custom Errors Write",
				"Dynamic URL Redirects Read",
				"Dynamic URL Redirects Write",
				"HTTP DDoS Managed Ruleset Read",
				"HTTP DDoS Managed Ruleset Write",
				"L4 DDoS Managed Ruleset Read",
				"L4 DDoS Managed Ruleset Write",
				"Logs Read",
				"Logs Write",
				"Magic Firewall Read",
				"Magic Firewall Write",
				"Managed headers Read",
				"Managed headers Write",
				"Mass URL Redirects Read",
				"Mass URL Redirects Write",
				"Origin Read",
				"Origin Write",
				"Response Compression Read",
				"Response Compression Write",
				"Sanitize Read",
				"Sanitize Write",
				"Select Configuration Read",
				"Select Configuration Write",
				"Transform Rules Read",
				"Transform Rules Write",
				"Zone Transform Rules Read",
				"Zone Transform Rules Write",
				"Zone WAF Read",
				"Zone WAF Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Computed:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The unique ID of the zone.",
				Required:    true,
			},
			"scope": schema.StringAttribute{
				Description: "The scope of the URL normalization.\nAvailable values: \"incoming\", \"both\", \"none\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"incoming",
						"both",
						"none",
					),
				},
			},
			"type": schema.StringAttribute{
				Description: "The type of URL normalization performed by Cloudflare.\nAvailable values: \"cloudflare\", \"rfc3986\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("cloudflare", "rfc3986"),
				},
			},
		},
	}
}

func (d *URLNormalizationSettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *URLNormalizationSettingsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
