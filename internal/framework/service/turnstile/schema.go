package turnstile

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *TurnstileWidgetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: heredoc.Doc(`
		The [Turnstile Widget](https://developers.cloudflare.com/turnstile/) resource allows you to manage Cloudflare Turnstile Widgets.
`),

		Attributes: map[string]schema.Attribute{
			consts.IDSchemaKey: schema.StringAttribute{
				Computed:            true,
				Optional:            true,
				MarkdownDescription: consts.IDSchemaDescription + " This is the site key value.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			consts.AccountIDSchemaKey: schema.StringAttribute{
				MarkdownDescription: consts.AccountIDSchemaDescription,
				Required:            true,
			},
			"secret": schema.StringAttribute{
				MarkdownDescription: "Secret key for this widget.",
				Computed:            true,
				Sensitive:           true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Human readable widget name.",
				Required:            true,
			},
			"domains": schema.SetAttribute{
				MarkdownDescription: "Domains where the widget is deployed",
				Required:            true,
				ElementType:         types.StringType,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: fmt.Sprintf("Widget Mode. %s", utils.RenderAvailableDocumentationValuesStringSlice([]string{"non-interactive", "invisible", "managed"})),
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("non-interactive", "invisible", "managed"),
				},
			},
			"region": schema.StringAttribute{
				MarkdownDescription: "Region where this widget can be used.",
				Computed:            true,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("world"),
				},
			},
			"bot_fight_mode": schema.BoolAttribute{
				MarkdownDescription: "If bot_fight_mode is set to true, Cloudflare issues computationally expensive challenges in response to malicious bots (Enterprise only).",
				Computed:            true,
				Optional:            true,
			},
			"offlabel": schema.BoolAttribute{
				MarkdownDescription: "Do not show any Cloudflare branding on the widget (Enterprise only).",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
		},
	}
}
