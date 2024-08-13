// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package observatory_scheduled_test

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = &ObservatoryScheduledTestResource{}

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "A URL.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"url": schema.StringAttribute{
				Description:   "A URL.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"frequency": schema.StringAttribute{
				Description: "The frequency of the test.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
				},
			},
			"item_count": schema.Float64Attribute{
				Description: "Number of items affected.",
				Computed:    true,
			},
			"region": schema.StringAttribute{
				Description: "A test region.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"asia-east1",
						"asia-northeast1",
						"asia-northeast2",
						"asia-south1",
						"asia-southeast1",
						"australia-southeast1",
						"europe-north1",
						"europe-southwest1",
						"europe-west1",
						"europe-west2",
						"europe-west3",
						"europe-west4",
						"europe-west8",
						"europe-west9",
						"me-west1",
						"southamerica-east1",
						"us-central1",
						"us-east1",
						"us-east4",
						"us-south1",
						"us-west1",
					),
				},
			},
			"schedule": schema.SingleNestedAttribute{
				Description: "The test schedule.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ObservatoryScheduledTestScheduleModel](ctx),
				Attributes: map[string]schema.Attribute{
					"frequency": schema.StringAttribute{
						Description: "The frequency of the test.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
						},
					},
					"region": schema.StringAttribute{
						Description: "A test region.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"asia-east1",
								"asia-northeast1",
								"asia-northeast2",
								"asia-south1",
								"asia-southeast1",
								"australia-southeast1",
								"europe-north1",
								"europe-southwest1",
								"europe-west1",
								"europe-west2",
								"europe-west3",
								"europe-west4",
								"europe-west8",
								"europe-west9",
								"me-west1",
								"southamerica-east1",
								"us-central1",
								"us-east1",
								"us-east4",
								"us-south1",
								"us-west1",
							),
						},
					},
					"url": schema.StringAttribute{
						Description: "A URL.",
						Optional:    true,
					},
				},
			},
			"test": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ObservatoryScheduledTestTestModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "UUID",
						Optional:    true,
					},
					"date": schema.StringAttribute{
						Optional:   true,
						CustomType: timetypes.RFC3339Type{},
					},
					"desktop_report": schema.SingleNestedAttribute{
						Description: "The Lighthouse report.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"cls": schema.Float64Attribute{
								Description: "Cumulative Layout Shift.",
								Optional:    true,
							},
							"device_type": schema.StringAttribute{
								Description: "The type of device.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("DESKTOP", "MOBILE"),
								},
							},
							"error": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"code": schema.StringAttribute{
										Description: "The error code of the Lighthouse result.",
										Optional:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"NOT_REACHABLE",
												"DNS_FAILURE",
												"NOT_HTML",
												"LIGHTHOUSE_TIMEOUT",
												"UNKNOWN",
											),
										},
									},
									"detail": schema.StringAttribute{
										Description: "Detailed error message.",
										Optional:    true,
									},
									"final_displayed_url": schema.StringAttribute{
										Description: "The final URL displayed to the user.",
										Optional:    true,
									},
								},
							},
							"fcp": schema.Float64Attribute{
								Description: "First Contentful Paint.",
								Optional:    true,
							},
							"json_report_url": schema.StringAttribute{
								Description: "The URL to the full Lighthouse JSON report.",
								Optional:    true,
							},
							"lcp": schema.Float64Attribute{
								Description: "Largest Contentful Paint.",
								Optional:    true,
							},
							"performance_score": schema.Float64Attribute{
								Description: "The Lighthouse performance score.",
								Optional:    true,
							},
							"si": schema.Float64Attribute{
								Description: "Speed Index.",
								Optional:    true,
							},
							"state": schema.StringAttribute{
								Description: "The state of the Lighthouse report.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"RUNNING",
										"COMPLETE",
										"FAILED",
									),
								},
							},
							"tbt": schema.Float64Attribute{
								Description: "Total Blocking Time.",
								Optional:    true,
							},
							"ttfb": schema.Float64Attribute{
								Description: "Time To First Byte.",
								Optional:    true,
							},
							"tti": schema.Float64Attribute{
								Description: "Time To Interactive.",
								Optional:    true,
							},
						},
					},
					"mobile_report": schema.SingleNestedAttribute{
						Description: "The Lighthouse report.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"cls": schema.Float64Attribute{
								Description: "Cumulative Layout Shift.",
								Optional:    true,
							},
							"device_type": schema.StringAttribute{
								Description: "The type of device.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("DESKTOP", "MOBILE"),
								},
							},
							"error": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"code": schema.StringAttribute{
										Description: "The error code of the Lighthouse result.",
										Optional:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"NOT_REACHABLE",
												"DNS_FAILURE",
												"NOT_HTML",
												"LIGHTHOUSE_TIMEOUT",
												"UNKNOWN",
											),
										},
									},
									"detail": schema.StringAttribute{
										Description: "Detailed error message.",
										Optional:    true,
									},
									"final_displayed_url": schema.StringAttribute{
										Description: "The final URL displayed to the user.",
										Optional:    true,
									},
								},
							},
							"fcp": schema.Float64Attribute{
								Description: "First Contentful Paint.",
								Optional:    true,
							},
							"json_report_url": schema.StringAttribute{
								Description: "The URL to the full Lighthouse JSON report.",
								Optional:    true,
							},
							"lcp": schema.Float64Attribute{
								Description: "Largest Contentful Paint.",
								Optional:    true,
							},
							"performance_score": schema.Float64Attribute{
								Description: "The Lighthouse performance score.",
								Optional:    true,
							},
							"si": schema.Float64Attribute{
								Description: "Speed Index.",
								Optional:    true,
							},
							"state": schema.StringAttribute{
								Description: "The state of the Lighthouse report.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"RUNNING",
										"COMPLETE",
										"FAILED",
									),
								},
							},
							"tbt": schema.Float64Attribute{
								Description: "Total Blocking Time.",
								Optional:    true,
							},
							"ttfb": schema.Float64Attribute{
								Description: "Time To First Byte.",
								Optional:    true,
							},
							"tti": schema.Float64Attribute{
								Description: "Time To Interactive.",
								Optional:    true,
							},
						},
					},
					"region": schema.SingleNestedAttribute{
						Description: "A test region with a label.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"label": schema.StringAttribute{
								Optional: true,
							},
							"value": schema.StringAttribute{
								Description: "A test region.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"asia-east1",
										"asia-northeast1",
										"asia-northeast2",
										"asia-south1",
										"asia-southeast1",
										"australia-southeast1",
										"europe-north1",
										"europe-southwest1",
										"europe-west1",
										"europe-west2",
										"europe-west3",
										"europe-west4",
										"europe-west8",
										"europe-west9",
										"me-west1",
										"southamerica-east1",
										"us-central1",
										"us-east1",
										"us-east4",
										"us-south1",
										"us-west1",
									),
								},
							},
						},
					},
					"schedule_frequency": schema.StringAttribute{
						Description: "The frequency of the test.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("DAILY", "WEEKLY"),
						},
					},
					"url": schema.StringAttribute{
						Description: "A URL.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *ObservatoryScheduledTestResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ObservatoryScheduledTestResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
