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

func (r PageShieldPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID of the policy",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"action": schema.StringAttribute{
				Description: "The action to take if the expression matches",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "log"),
				},
			},
			"description": schema.StringAttribute{
				Description: "A description for the policy",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the policy is enabled",
				Optional:    true,
			},
			"expression": schema.StringAttribute{
				Description: "The expression which must match for the policy to be applied, using the Cloudflare Firewall rule expression syntax",
				Optional:    true,
			},
			"value": schema.StringAttribute{
				Description: "The policy which will be applied",
				Optional:    true,
			},
		},
	}
}
