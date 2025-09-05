// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPPredefinedProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The id of the profile (uuid).",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"context_awareness": schema.SingleNestedAttribute{
				Description:        "Scan the context of predefined entries to only return matches surrounded by keywords.",
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
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
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
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
			"ai_context_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"allowed_match_count": schema.Int64Attribute{
				Computed: true,
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(0, 1000),
				},
				Default: int64default.StaticInt64(0),
			},
			"confidence_threshold": schema.StringAttribute{
				Computed: true,
				Optional: true,
				Default:  stringdefault.StaticString("low"),
			},
			"ocr_enabled": schema.BoolAttribute{
				Computed: true,
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
			"created_at": schema.StringAttribute{
				Description: "When the profile was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Description: "The description of the profile.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the profile.",
				Computed:    true,
			},
			"open_access": schema.BoolAttribute{
				Description: "Whether this profile can be accessed by anyone.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: `Available values: "custom", "predefined", "integration".`,
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"custom",
						"predefined",
						"integration",
					),
				},
			},
			"updated_at": schema.StringAttribute{
				Description: "When the profile was lasted updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
