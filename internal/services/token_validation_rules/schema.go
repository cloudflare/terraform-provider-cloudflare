// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation_rules

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*TokenValidationRulesResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"action": schema.StringAttribute{
				Description: "Action to take on requests that match operations included in `selector` and fail `expression`.\nAvailable values: \"log\", \"block\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("log", "block"),
				},
			},
			"description": schema.StringAttribute{
				Description: "A human-readable description that gives more details than `title`.",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Toggle rule on or off.",
				Required:    true,
			},
			"expression": schema.StringAttribute{
				Description: "Rule expression. Requests that fail to match this expression will be subject to `action`.\n\nFor details on expressions, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).",
				Required:    true,
			},
			"title": schema.StringAttribute{
				Description: "A human-readable name for the rule.",
				Required:    true,
			},
			"selector": schema.SingleNestedAttribute{
				Description: "Select operations covered by this rule.\n\nFor details on selectors, see the [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"exclude": schema.ListNestedAttribute{
						Description: "Ignore operations that were otherwise included by `include`.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"operation_ids": schema.ListAttribute{
									Description: "Excluded operation IDs.",
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
					"include": schema.ListNestedAttribute{
						Description: "Select all matching operations.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"host": schema.ListAttribute{
									Description: "Included hostnames.",
									Optional:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
			"position": schema.SingleNestedAttribute{
				Description: "Update rule order among zone rules.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"index": schema.Int64Attribute{
						Description: "Move rule to this position",
						Optional:    true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
					"before": schema.StringAttribute{
						Description: "Move rule to before rule with this ID.",
						Optional:    true,
					},
					"after": schema.StringAttribute{
						Description: "Move rule to after rule with this ID.",
						Optional:    true,
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"last_updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *TokenValidationRulesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TokenValidationRulesResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
