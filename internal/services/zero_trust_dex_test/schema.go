// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_test

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
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
				Description:   "The name of the DEX test. Must be unique.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "The name of the DEX test. Must be unique.",
				Computed:      true,
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "Additional details about the test.",
				Computed:    true,
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Determines whether or not the test is active.",
				Computed:    true,
				Optional:    true,
			},
			"interval": schema.StringAttribute{
				Description: "How often the test will run.",
				Computed:    true,
				Optional:    true,
			},
			"targeted": schema.BoolAttribute{
				Computed: true,
				Optional: true,
			},
			"data": schema.SingleNestedAttribute{
				Description: "The configuration object which contains the details for the WARP client to conduct the test.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDEXTestDataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						Description: "The desired endpoint to test.",
						Computed:    true,
						Optional:    true,
					},
					"kind": schema.StringAttribute{
						Description: "The type of test.",
						Computed:    true,
						Optional:    true,
					},
					"method": schema.StringAttribute{
						Description: "The HTTP request method type.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"target_policies": schema.ListNestedAttribute{
				Description: "Device settings profiles targeted by this test",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDEXTestTargetPoliciesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the device settings profile",
							Computed:    true,
							Optional:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the profile is the account default",
							Computed:    true,
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device settings profile",
							Computed:    true,
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
