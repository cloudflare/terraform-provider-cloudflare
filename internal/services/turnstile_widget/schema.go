// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile_widget

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*TurnstileWidgetResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Settings Read",
				"Account Settings Write",
				"Turnstile Sites Read",
				"Turnstile Sites Write",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Widget item identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"sitekey": schema.StringAttribute{
				Description:   "Widget item identifier tag.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"direction": schema.StringAttribute{
				Description: "Direction to order widgets.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"filter": schema.StringAttribute{
				Description: "Filter widgets by field using case-insensitive substring matching.\nFormat: `field:value`\n\nSupported fields:\n- `name` - Filter by widget name (e.g., `filter=name:login-form`)\n- `sitekey` - Filter by sitekey (e.g., `filter=sitekey:0x4AAA`)\n\nReturns 400 Bad Request if the field is unsupported or format is invalid.\nAn empty filter value returns all results.",
				Optional:    true,
			},
			"order": schema.StringAttribute{
				Description: "Field to order widgets by.\nAvailable values: \"id\", \"sitekey\", \"name\", \"created_on\", \"modified_on\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"id",
						"sitekey",
						"name",
						"created_on",
						"modified_on",
					),
				},
			},
			"page": schema.Float64Attribute{
				Description: "Page number of paginated results.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
				Default: float64default.StaticFloat64(1),
			},
			"per_page": schema.Float64Attribute{
				Description: "Number of items per page.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.Between(5, 1000),
				},
				Default: float64default.StaticFloat64(25),
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
				Computed:    true,
				Optional:    true,
			},
			"clearance_level": schema.StringAttribute{
				Description: "If Turnstile is embedded on a Cloudflare site and the widget should grant challenge clearance,\nthis setting can determine the clearance level to be set\nAvailable values: \"no_clearance\", \"jschallenge\", \"managed\", \"interactive\".",
				Computed:    true,
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
				Computed:    true,
				Optional:    true,
			},
			"offlabel": schema.BoolAttribute{
				Description: "Do not show any Cloudflare branding on the widget (ENT only).",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"region": schema.StringAttribute{
				Description: "Region where this widget can be used. This cannot be changed after creation.\nAvailable values: \"world\", \"china\".",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("world", "china"),
				},
				Default: stringdefault.StaticString("world"),
			},
			"created_on": schema.StringAttribute{
				Description:   "When the widget was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"deployed_via": schema.StringAttribute{
				Description: "Origin that created this widget, recorded at creation time and\nimmutable afterward. Server-derived from the create request; not\nclient-settable. Omitted from the response for widgets created\nbefore this field existed.\nAvailable values: \"wrangler\", \"dashboard\", \"spin\", \"api\", \"unknown\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"wrangler",
						"dashboard",
						"spin",
						"api",
						"unknown",
					),
				},
			},
			"last_modified_via": schema.StringAttribute{
				Description: "Origin of the most recent mutation (create, update, delete, or\nsecret rotation). Server-derived; not client-settable. Omitted for\nwidgets last mutated before this field existed.\nAvailable values: \"wrangler\", \"dashboard\", \"spin\", \"api\", \"unknown\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"wrangler",
						"dashboard",
						"spin",
						"api",
						"unknown",
					),
				},
			},
			"modified_on": schema.StringAttribute{
				Description: "When the widget was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"secret": schema.StringAttribute{
				Description: "Secret key for this widget.",
				Computed:    true,
				Sensitive:   true,
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
