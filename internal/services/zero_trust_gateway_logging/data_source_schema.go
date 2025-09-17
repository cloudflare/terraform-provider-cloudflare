// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_logging

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewayLoggingDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"redact_pii": schema.BoolAttribute{
				Description: "Indicate whether to redact personally identifiable information from activity logging (PII fields include source IP, user email, user ID, device ID, URL, referrer, and user agent).",
				Computed:    true,
			},
			"settings_by_rule_type": schema.SingleNestedAttribute{
				Description: "Configure logging settings for each rule type.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"dns": schema.SingleNestedAttribute{
						Description: "Configure logging settings for DNS firewall.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeDNSDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"log_all": schema.BoolAttribute{
								Description: "Specify whether to log all requests to this service.",
								Computed:    true,
							},
							"log_blocks": schema.BoolAttribute{
								Description: "Specify whether to log only blocking requests to this service.",
								Computed:    true,
							},
						},
					},
					"http": schema.SingleNestedAttribute{
						Description: "Configure logging settings for HTTP/HTTPS firewall.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeHTTPDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"log_all": schema.BoolAttribute{
								Description: "Specify whether to log all requests to this service.",
								Computed:    true,
							},
							"log_blocks": schema.BoolAttribute{
								Description: "Specify whether to log only blocking requests to this service.",
								Computed:    true,
							},
						},
					},
					"l4": schema.SingleNestedAttribute{
						Description: "Configure logging settings for Network firewall.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeL4DataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"log_all": schema.BoolAttribute{
								Description: "Specify whether to log all requests to this service.",
								Computed:    true,
							},
							"log_blocks": schema.BoolAttribute{
								Description: "Specify whether to log only blocking requests to this service.",
								Computed:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustGatewayLoggingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustGatewayLoggingDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
