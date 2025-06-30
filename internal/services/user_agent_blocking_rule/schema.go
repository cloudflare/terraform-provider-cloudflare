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
			"id": schema.StringAttribute{
				Description:   "The unique identifier of the User Agent Blocking rule.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Defines an identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"mode": schema.StringAttribute{
				Description: "The action to apply to a matched request.\nAvailable values: \"block\", \"challenge\", \"whitelist\", \"js_challenge\", \"managed_challenge\".",
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
				Required: true,
				Attributes: map[string]schema.Attribute{
					"target": schema.StringAttribute{
						Description: "The configuration target. You must set the target to `ua` when specifying a user agent in the rule.\nAvailable values: \"ua\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("ua"),
						},
					},
					"value": schema.StringAttribute{
						Description: "the user agent to exactly match",
						Optional:    true,
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the rule. This value is sanitized and any tags will be removed.",
				Optional:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the rule is currently paused.",
				Optional:    true,
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
