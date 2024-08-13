// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = &ZeroTrustDLPCustomProfileResource{}

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Description:   "Unique identifier for a DLP profile",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profiles": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allowed_match_count": schema.Float64Attribute{
							Description: "Related DLP policies will trigger when the match count exceeds the number set.",
							Computed:    true,
							Optional:    true,
							Validators: []validator.Float64{
								float64validator.Between(0, 1000),
							},
							Default: float64default.StaticFloat64(0),
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
							Description: "The description of the profile.",
							Optional:    true,
						},
						"entries": schema.ListNestedAttribute{
							Description: "The entries for this profile.",
							Optional:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Whether the entry is enabled or not.",
										Required:    true,
									},
									"name": schema.StringAttribute{
										Description: "The name of the entry.",
										Required:    true,
									},
									"pattern": schema.SingleNestedAttribute{
										Description: "A pattern that matches an entry",
										Required:    true,
										Attributes: map[string]schema.Attribute{
											"regex": schema.StringAttribute{
												Description: "The regex pattern.",
												Required:    true,
											},
											"validation": schema.StringAttribute{
												Description: "Validation algorithm for the pattern. This algorithm will get run on potential matches, and if it returns false, the entry will not be matched.",
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("luhn"),
												},
											},
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the profile.",
							Optional:    true,
						},
						"ocr_enabled": schema.BoolAttribute{
							Description: "If true, scan images via OCR to determine if any text present matches filters.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Description: "The description of the profile.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the profile.",
				Optional:    true,
			},
			"ocr_enabled": schema.BoolAttribute{
				Description: "If true, scan images via OCR to determine if any text present matches filters.",
				Optional:    true,
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
			"entries": schema.ListNestedAttribute{
				Description: "The custom entries for this profile. Array elements with IDs are modifying the existing entry with that ID. Elements without ID will create new entries. Any entry not in the list will be deleted.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for a DLP entry",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the entry is enabled or not.",
							Optional:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the entry.",
							Optional:    true,
						},
						"pattern": schema.SingleNestedAttribute{
							Description: "A pattern that matches an entry",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"regex": schema.StringAttribute{
									Description: "The regex pattern.",
									Required:    true,
								},
								"validation": schema.StringAttribute{
									Description: "Validation algorithm for the pattern. This algorithm will get run on potential matches, and if it returns false, the entry will not be matched.",
									Optional:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("luhn"),
									},
								},
							},
						},
						"profile_id": schema.StringAttribute{
							Description: "Unique identifier for a DLP profile",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
					},
				},
			},
			"shared_entries": schema.ListNestedAttribute{
				Description: "Entries from other profiles (e.g. pre-defined Cloudflare profiles, or your Microsoft Information Protection profiles).",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"enabled": schema.BoolAttribute{
							Description: "Whether the entry is enabled or not.",
							Optional:    true,
						},
						"entry_id": schema.StringAttribute{
							Description: "Unique identifier for a DLP entry",
							Computed:    true,
						},
					},
				},
			},
			"allowed_match_count": schema.Float64Attribute{
				Description: "Related DLP policies will trigger when the match count exceeds the number set.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1000),
				},
				Default: float64default.StaticFloat64(0),
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Unique identifier for a DLP profile",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the profile.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("custom"),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
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
