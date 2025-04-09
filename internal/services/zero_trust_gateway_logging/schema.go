// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_logging

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewayLoggingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"redact_pii": schema.BoolAttribute{
				Description: "Redact personally identifiable information from activity logging (PII fields are: source IP, user email, user ID, device ID, URL, referrer, user agent).",
				Optional:    true,
			},
			"settings_by_rule_type": schema.SingleNestedAttribute{
				Description: "Logging settings by rule type.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeModel](ctx),
				Attributes: map[string]schema.Attribute{
					"dns": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeDNSModel](ctx),
						Attributes: map[string]schema.Attribute{
							"log_all": schema.BoolAttribute{
								Description: "Log all requests to this service.",
								Optional:    true,
							},
							"log_blocks": schema.BoolAttribute{
								Description: "Log only blocking requests to this service.",
								Optional:    true,
							},
						},
					},
					"http": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeHTTPModel](ctx),
						Attributes: map[string]schema.Attribute{
							"log_all": schema.BoolAttribute{
								Description: "Log all requests to this service.",
								Optional:    true,
							},
							"log_blocks": schema.BoolAttribute{
								Description: "Log only blocking requests to this service.",
								Optional:    true,
							},
						},
					},
					"l4": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[ZeroTrustGatewayLoggingSettingsByRuleTypeL4Model](ctx),
						Attributes: map[string]schema.Attribute{
							"log_all": schema.BoolAttribute{
								Description: "Log all requests to this service.",
								Optional:    true,
							},
							"log_blocks": schema.BoolAttribute{
								Description: "Log only blocking requests to this service.",
								Optional:    true,
							},
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustGatewayLoggingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewayLoggingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
