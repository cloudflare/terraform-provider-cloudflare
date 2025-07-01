// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZoneSubscriptionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Subscription identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Subscription identifier tag.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"frequency": schema.StringAttribute{
				Description: "How often the subscription is renewed automatically.\nAvailable values: \"weekly\", \"monthly\", \"quarterly\", \"yearly\".",
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
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "The ID of the rate plan.\nAvailable values: \"free\", \"lite\", \"pro\", \"pro_plus\", \"business\", \"enterprise\", \"partners_free\", \"partners_pro\", \"partners_business\", \"partners_enterprise\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"free",
								"lite",
								"pro",
								"pro_plus",
								"business",
								"enterprise",
								"partners_free",
								"partners_pro",
								"partners_business",
								"partners_enterprise",
							),
						},
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
			"currency": schema.StringAttribute{
				Description: "The monetary unit in which pricing information is displayed.",
				Computed:    true,
			},
			"current_period_end": schema.StringAttribute{
				Description: "The end of the current period and also when the next billing is due.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"current_period_start": schema.StringAttribute{
				Description: "When the current billing period started. May match initial_period_start if this is the first period.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"price": schema.Float64Attribute{
				Description: "The price of the subscription that will be billed, in US dollars.",
				Computed:    true,
			},
			"state": schema.StringAttribute{
				Description: "The state that the subscription is in.\nAvailable values: \"Trial\", \"Provisioned\", \"Paid\", \"AwaitingPayment\", \"Cancelled\", \"Failed\", \"Expired\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"Trial",
						"Provisioned",
						"Paid",
						"AwaitingPayment",
						"Cancelled",
						"Failed",
						"Expired",
					),
				},
			},
		},
	}
}

func (r *ZoneSubscriptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneSubscriptionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
