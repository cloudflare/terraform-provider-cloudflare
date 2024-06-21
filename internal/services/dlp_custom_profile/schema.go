// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_custom_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r DLPCustomProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"profile_id": schema.StringAttribute{
				Description: "The ID for this profile",
				Optional:    true,
			},
			"profiles": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allowed_match_count": schema.Float64Attribute{
							Description: "Related DLP policies will trigger when the match count exceeds the number set.",
							Computed:    true,
							Optional:    true,
							Default:     float64default.StaticFloat64(0),
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
			},
		},
	}
}
