// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = &FilterResource{}

func (r *FilterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_identifier": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"id": schema.StringAttribute{
				Description:   "The unique identifier of the filter.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"expression": schema.StringAttribute{
				Description:   "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the filter.",
				Computed:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the filter is currently paused.",
				Computed:    true,
			},
			"ref": schema.StringAttribute{
				Description: "A short reference tag. Allows you to select related filters.",
				Computed:    true,
			},
		},
	}
}

func (r *FilterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
