// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_subscription

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZoneSubscriptionsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZoneSubscriptionsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Subscription identifier tag.",
							Computed:    true,
						},
						"app": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZoneSubscriptionsAppDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"install_id": schema.StringAttribute{
									Description: "app install id.",
									Computed:    true,
								},
							},
						},
						"component_values": schema.ListNestedAttribute{
							Description: "The list of add-ons subscribed to.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZoneSubscriptionsComponentValuesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"default": schema.Float64Attribute{
										Description: "The default amount assigned.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "The name of the component value.",
										Computed:    true,
									},
									"price": schema.Float64Attribute{
										Description: "The unit price for the component value.",
										Computed:    true,
									},
									"value": schema.Float64Attribute{
										Description: "The amount of the component value assigned.",
										Computed:    true,
									},
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
						"frequency": schema.StringAttribute{
							Description: "How often the subscription is renewed automatically.",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"weekly",
									"monthly",
									"quarterly",
									"yearly",
								),
							},
						},
						"price": schema.Float64Attribute{
							Description: "The price of the subscription that will be billed, in US dollars.",
							Computed:    true,
						},
						"rate_plan": schema.SingleNestedAttribute{
							Description: "The rate plan applied to the subscription.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZoneSubscriptionsRatePlanDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "The ID of the rate plan.",
									Computed:    true,
								},
								"currency": schema.StringAttribute{
									Description: "The currency applied to the rate plan subscription.",
									Computed:    true,
								},
								"externally_managed": schema.BoolAttribute{
									Description: "Whether this rate plan is managed externally from Cloudflare.",
									Computed:    true,
								},
								"is_contract": schema.BoolAttribute{
									Description: "Whether a rate plan is enterprise-based (or newly adopted term contract).",
									Computed:    true,
								},
								"public_name": schema.StringAttribute{
									Description: "The full name of the rate plan.",
									Computed:    true,
								},
								"scope": schema.StringAttribute{
									Description: "The scope that this rate plan applies to.",
									Computed:    true,
								},
								"sets": schema.ListAttribute{
									Description: "The list of sets this rate plan applies to.",
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
						"state": schema.StringAttribute{
							Description: "The state that the subscription is in.",
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
						"zone": schema.SingleNestedAttribute{
							Description: "A simple zone object. May have null properties if not a zone subscription.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZoneSubscriptionsZoneDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Identifier",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "The domain name",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ZoneSubscriptionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZoneSubscriptionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
