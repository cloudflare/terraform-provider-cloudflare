// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web_analytics_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r WebAnalyticsRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The Web Analytics rule identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"ruleset_id": schema.StringAttribute{
				Description: "The Web Analytics ruleset identifier.",
				Required:    true,
			},
			"host": schema.StringAttribute{
				Optional: true,
			},
			"inclusive": schema.BoolAttribute{
				Description: "Whether the rule includes or excludes traffic from being measured.",
				Optional:    true,
			},
			"is_paused": schema.BoolAttribute{
				Description: "Whether the rule is paused or not.",
				Optional:    true,
			},
			"paths": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}
