// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPCustomProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profiles": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"entries": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Required: true,
									},
									"name": schema.StringAttribute{
										Required: true,
									},
									"pattern": schema.SingleNestedAttribute{
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"regex": schema.StringAttribute{
												Required: true,
											},
											"validation": schema.StringAttribute{
												Optional: true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("luhn"),
												},
											},
										},
									},
									"words": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Required: true,
						},
						"allowed_match_count": schema.Int64Attribute{
							Description: "Related DLP policies will trigger when the match count exceeds the number set.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.Int64{
								int64validator.Between(0, 1000),
							},
							Default: int64default.StaticInt64(0),
						},
						"confidence_threshold": schema.StringAttribute{
							Optional: true,
						},
						"context_awareness": schema.SingleNestedAttribute{
							Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "If true, scan the context of predefined entries to only return matches surrounded by keywords.",
									Required:    true,
								},
								"skip": schema.SingleNestedAttribute{
									Description: "Content types to exclude from context analysis and return all matches.",
									Required:    true,
									Attributes: map[string]schema.Attribute{
										"files": schema.BoolAttribute{
											Description: "If the content type is a file, skip context analysis and return all matches.",
											Required:    true,
										},
									},
								},
							},
						},
						"description": schema.StringAttribute{
							Description: "The description of the profile",
							Optional:    true,
						},
						"ocr_enabled": schema.BoolAttribute{
							Optional: true,
						},
						"shared_entries": schema.ListNestedAttribute{
							Description: "Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your Microsoft Information Protection profiles).",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Required: true,
									},
									"entry_id": schema.StringAttribute{
										Required: true,
									},
									"entry_type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"custom",
												"predefined",
												"integration",
												"exact_data",
											),
										},
									},
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"allowed_match_count": schema.Int64Attribute{
				Optional: true,
			},
			"confidence_threshold": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the profile",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"ocr_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"context_awareness": schema.SingleNestedAttribute{
				Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDLPCustomProfileContextAwarenessModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "If true, scan the context of predefined entries to only return matches surrounded by keywords.",
						Required:    true,
					},
					"skip": schema.SingleNestedAttribute{
						Description: "Content types to exclude from context analysis and return all matches.",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"files": schema.BoolAttribute{
								Description: "If the content type is a file, skip context analysis and return all matches.",
								Required:    true,
							},
						},
					},
				},
			},
			"entries": schema.ListNestedAttribute{
				Description: "Custom entries from this profile",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDLPCustomProfileEntriesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Required: true,
						},
						"entry_id": schema.StringAttribute{
							Optional: true,
						},
						"name": schema.StringAttribute{
							Required: true,
						},
						"pattern": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"regex": schema.StringAttribute{
									Required: true,
								},
								"validation": schema.StringAttribute{
									Optional: true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("luhn"),
									},
								},
							},
						},
					},
				},
			},
			"shared_entries": schema.ListNestedAttribute{
				Description: "Other entries, e.g. predefined or integration.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDLPCustomProfileSharedEntriesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Required: true,
						},
						"entry_id": schema.StringAttribute{
							Required: true,
						},
						"entry_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"predefined",
									"integration",
									"exact_data",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustDLPCustomProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDLPCustomProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
