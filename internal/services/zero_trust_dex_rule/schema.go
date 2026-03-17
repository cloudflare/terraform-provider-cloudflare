// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dex_rule

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDEXRuleResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "API Resource UUID tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"match": schema.StringAttribute{
				Description: "The wirefilter expression to match.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the Rule.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
			"targeted_tests": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDEXRuleTargetedTestsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"data": schema.SingleNestedAttribute{
							Description: "The configuration object which contains the details for the WARP client to conduct the test.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustDEXRuleTargetedTestsDataModel](ctx),
							Attributes: map[string]schema.Attribute{
								"host": schema.StringAttribute{
									Description: "The desired endpoint to test.",
									Computed:    true,
								},
								"kind": schema.StringAttribute{
									Description: "The type of test.\nAvailable values: \"http\", \"traceroute\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("http", "traceroute"),
									},
								},
								"method": schema.StringAttribute{
									Description: "The HTTP request method type.\nAvailable values: \"GET\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("GET"),
									},
								},
							},
						},
						"enabled": schema.BoolAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"test_id": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustDEXRuleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDEXRuleResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
