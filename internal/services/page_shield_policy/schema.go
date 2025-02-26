// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*PageShieldPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"action": schema.StringAttribute{
				Description: "The action to take if the expression matches\navailable values: \"allow\", \"log\"",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "log"),
				},
			},
			"description": schema.StringAttribute{
				Description: "A description for the policy",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the policy is enabled",
				Required:    true,
			},
			"expression": schema.StringAttribute{
				Description: "The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax",
				Required:    true,
			},
			"value": schema.StringAttribute{
				Description: "The policy which will be applied",
				Required:    true,
			},
		},
	}
}

func (r *PageShieldPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PageShieldPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
