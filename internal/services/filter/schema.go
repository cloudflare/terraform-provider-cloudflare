// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package filter

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*FilterResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		DeprecationMessage: "The Filters API is deprecated in favour of using the Ruleset Engine. See https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api for full details.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique identifier of the filter.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Defines an identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"body": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the filter.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "An informative summary of the filter.",
							Optional:    true,
						},
						"expression": schema.StringAttribute{
							Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
							Optional:    true,
						},
						"paused": schema.BoolAttribute{
							Description: "When true, indicates that the filter is currently paused.",
							Optional:    true,
						},
						"ref": schema.StringAttribute{
							Description: "A short reference tag. Allows you to select related filters.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "An informative summary of the filter.",
				Optional:    true,
			},
			"expression": schema.StringAttribute{
				Description: "The filter expression. For more information, refer to [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).",
				Optional:    true,
			},
			"paused": schema.BoolAttribute{
				Description: "When true, indicates that the filter is currently paused.",
				Optional:    true,
			},
			"ref": schema.StringAttribute{
				Description: "A short reference tag. Allows you to select related filters.",
				Optional:    true,
			},
		},
	}
}

func (r *FilterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *FilterResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
