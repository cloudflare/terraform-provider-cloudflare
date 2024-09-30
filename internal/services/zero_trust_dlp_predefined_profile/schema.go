// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPPredefinedProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"entries": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Required: true,
						},
						"enabled": schema.BoolAttribute{
							Required: true,
						},
					},
				},
			},
			"allowed_match_count": schema.Int64Attribute{
				Optional: true,
			},
			"ocr_enabled": schema.BoolAttribute{
				Optional: true,
			},
			"context_awareness": schema.SingleNestedAttribute{
				Description: "Scan the context of predefined entries to only return matches surrounded by keywords.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustDLPPredefinedProfileContextAwarenessModel](ctx),
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
		},
	}
}

func (r *ZeroTrustDLPPredefinedProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDLPPredefinedProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
