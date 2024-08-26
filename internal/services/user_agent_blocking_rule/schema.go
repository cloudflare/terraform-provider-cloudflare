// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user_agent_blocking_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*UserAgentBlockingRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Description:   "The unique identifier of the User Agent Blocking rule.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"block",
						"challenge",
						"whitelist",
						"js_challenge",
						"managed_challenge",
					),
				},
			},
			"configuration": schema.SingleNestedAttribute{
				Description: "The rule configuration.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"target": schema.StringAttribute{
						Description: "The configuration target. You must set the target to `ip` when specifying an IP address in the rule.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"ip",
								"ip6",
								"ip_range",
								"asn",
								"country",
							),
						},
					},
					"value": schema.StringAttribute{
						Description: "The IP address to match. This address will be compared to the IP address of incoming requests.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *UserAgentBlockingRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *UserAgentBlockingRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
