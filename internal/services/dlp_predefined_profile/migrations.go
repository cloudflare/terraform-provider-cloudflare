// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dlp_predefined_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithUpgradeState = &DLPPredefinedProfileResource{}

func (r *DLPPredefinedProfileResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description:   "Unique identifier for a DLP profile",
						Computed:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"profile_id": schema.StringAttribute{
						Description:   "Unique identifier for a DLP profile",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
					},
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
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
						Description: "The entries for this profile.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "Unique identifier for a DLP entry",
									Computed:    true,
								},
								"enabled": schema.BoolAttribute{
									Description: "Whether the entry is enabled or not.",
									Optional:    true,
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
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state DLPPredefinedProfileModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
