// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r TurnstileWidgetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
				Description: "Identifier",
				Required:    true,
			},
			"domains": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"mode": schema.StringAttribute{
				Description: "Widget Mode",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("non-interactive", "invisible", "managed"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Human readable widget name. Not unique. Cloudflare suggests that you\nset this to a meaningful string to make it easier to identify your\nwidget, and where it is used.\n",
				Required:    true,
			},
			"bot_fight_mode": schema.BoolAttribute{
				Description: "If bot_fight_mode is set to `true`, Cloudflare issues computationally\nexpensive challenges in response to malicious bots (ENT only).\n",
				Optional:    true,
			},
			"clearance_level": schema.StringAttribute{
				Description: "If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,\nthis setting can determine the clearance level to be set\n",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("no_clearance", "jschallenge", "managed", "interactive"),
				},
			},
			"offlabel": schema.BoolAttribute{
				Description: "Do not show any Cloudflare branding on the widget (ENT only).\n",
				Optional:    true,
			},
			"region": schema.StringAttribute{
				Description: "Region where this widget can be used.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("world"),
				},
				Default: stringdefault.StaticString("world"),
			},
		},
	}
}
