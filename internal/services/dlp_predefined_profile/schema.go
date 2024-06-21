// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_predefined_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

func (r DLPPredefinedProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The ID for this profile",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"profile_id": schema.StringAttribute{
				Description: "The ID for this profile",
				Required:    true,
			},
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
			},
			"ocr_enabled": schema.BoolAttribute{
				Description: "If true, scan images via OCR to determine if any text present matches filters.",
				Optional:    true,
			},
		},
	}
}
