// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDEXTestResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The unique identifier for the test.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"test_id": schema.StringAttribute{
				Description:   "The unique identifier for the test.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "Determines whether or not the test is active.",
				Required:    true,
			},
			"interval": schema.StringAttribute{
				Description: "How often the test will run.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the DEX test. Must be unique.",
				Required:    true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "The configuration object which contains the details for the WARP client to conduct the test.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description: "The desired endpoint to test.",
						Optional:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of test.",
						Optional:    true,
					},
					"method": schema.StringAttribute{
						Description: "The HTTP request method type.",
						Optional:    true,
					},
				},
			},
			"description": schema.StringAttribute{
				Description: "Additional details about the test.",
				Optional:    true,
			},
			"targeted": schema.BoolAttribute{
				Optional: true,
			},
			"target_policies": schema.ListNestedAttribute{
				Description: "DEX rules targeted by this test",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the DEX rule",
							Optional:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the DEX rule is the account default",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the DEX rule",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustDEXTestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDEXTestResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
