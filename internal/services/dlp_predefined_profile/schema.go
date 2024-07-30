// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_predefined_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r DLPPredefinedProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID for this profile",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Description:   "The ID for this profile",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"allowed_match_count": schema.Float64Attribute{
				Description: "Related DLP policies will trigger when the match count exceeds the number set.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(0, 1000),
				},
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
				Default:       float64default.StaticFloat64(0),
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
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"entries": schema.ListNestedAttribute{
				Description: "The entries for this profile.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The ID for this entry",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the entry is enabled or not.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"ocr_enabled": schema.BoolAttribute{
				Description:   "If true, scan images via OCR to determine if any text present matches filters.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The name of the profile.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The type of the profile.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("predefined"),
				},
			},
		},
	}
}
