// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ManagedTransformsDataSource)(nil)

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
				Optional:    true,
			},
			"managed_request_headers": schema.ListNestedAttribute{
				Description: "The list of Managed Request Transforms.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ManagedTransformsManagedRequestHeadersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Computed:    true,
						},
					},
				},
			},
			"managed_response_headers": schema.ListNestedAttribute{
				Description: "The list of Managed Response Transforms.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ManagedTransformsManagedResponseHeadersDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The human-readable identifier of the Managed Transform.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the Managed Transform is enabled.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ManagedTransformsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ManagedTransformsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
