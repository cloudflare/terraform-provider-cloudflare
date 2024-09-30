// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_subscription

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*AccountSubscriptionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"subscription_identifier": schema.StringAttribute{
				Description:   "Subscription identifier tag.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"frequency": schema.StringAttribute{
				Description: "How often the subscription is renewed automatically.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"weekly",
						"monthly",
						"quarterly",
						"yearly",
					),
				},
			},
			"rate_plan": schema.SingleNestedAttribute{
				Description: "The rate plan applied to the subscription.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[AccountSubscriptionRatePlanModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The ID of the rate plan.",
						Optional:    true,
					},
					"currency": schema.StringAttribute{
						Description: "The currency applied to the rate plan subscription.",
						Optional:    true,
					},
					"externally_managed": schema.BoolAttribute{
						Description: "Whether this rate plan is managed externally from Cloudflare.",
						Optional:    true,
					},
					"is_contract": schema.BoolAttribute{
						Description: "Whether a rate plan is enterprise-based (or newly adopted term contract).",
						Optional:    true,
					},
					"public_name": schema.StringAttribute{
						Description: "The full name of the rate plan.",
						Optional:    true,
					},
					"scope": schema.StringAttribute{
						Description: "The scope that this rate plan applies to.",
						Optional:    true,
					},
					"sets": schema.ListAttribute{
						Description: "The list of sets this rate plan applies to.",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"subscription_id": schema.StringAttribute{
				Description: "Subscription identifier tag.",
				Computed:    true,
			},
		},
	}
}

func (r *AccountSubscriptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *AccountSubscriptionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
