// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_predefined_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDLPPredefinedProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"profile_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"enabled_entries": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
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
			"entries": schema.ListNestedAttribute{
				Computed:           true,
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
				CustomType:         customfield.NewNestedObjectListType[ZeroTrustDLPPredefinedProfileEntriesModel](ctx),
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
			"name": schema.StringAttribute{
				Description: "The name of the predefined profile.",
				Computed:    true,
			},
			"open_access": schema.BoolAttribute{
				Description: "Whether this profile can be accessed by anyone.",
				Computed:    true,
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
