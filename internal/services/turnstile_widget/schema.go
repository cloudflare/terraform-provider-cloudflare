// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*TurnstileWidgetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Widget item identifier tag.",
				Computed:    true,
			},
			"sitekey": schema.StringAttribute{
				Description:   "Widget item identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"region": schema.StringAttribute{
				Description: "Region where this widget can be used.\nAvailable values: \"world\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("world"),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Default:       stringdefault.StaticString("world"),
			},
			"mode": schema.StringAttribute{
				Description: "Widget Mode\nAvailable values: \"non-interactive\", \"invisible\", \"managed\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"non-interactive",
						"invisible",
						"managed",
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "Human readable widget name. Not unique. Cloudflare suggests that you\nset this to a meaningful string to make it easier to identify your\nwidget, and where it is used.",
				Required:    true,
			},
			"domains": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"bot_fight_mode": schema.BoolAttribute{
				Description: "If bot_fight_mode is set to `true`, Cloudflare issues computationally\nexpensive challenges in response to malicious bots (ENT only).",
				Optional:    true,
			},
			"clearance_level": schema.StringAttribute{
				Description: "If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,\nthis setting can determine the clearance level to be set\nAvailable values: \"no_clearance\", \"jschallenge\", \"managed\", \"interactive\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"no_clearance",
						"jschallenge",
						"managed",
						"interactive",
					),
				},
			},
			"ephemeral_id": schema.BoolAttribute{
				Description: "Return the Ephemeral ID in /siteverify (ENT only).",
				Optional:    true,
			},
			"offlabel": schema.BoolAttribute{
				Description: "Do not show any Cloudflare branding on the widget (ENT only).",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the widget was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the widget was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"secret": schema.StringAttribute{
				Description: "Secret key for this widget.",
				Computed:    true,
			},
		},
	}
}

func (r *TurnstileWidgetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *TurnstileWidgetResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
