// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package bot_management

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ resource.ResourceWithConfigValidators = (*BotManagementResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"ai_bots_protection": schema.StringAttribute{
				Description: "Enable rule to block AI Scrapers and Crawlers.\nAvailable values: \"block\", \"disabled\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("block", "disabled"),
				},
			},
			"auto_update_model": schema.BoolAttribute{
				Description: "Automatically update to the newest bot detection models created by Cloudflare as they are released. [Learn more.](https://developers.cloudflare.com/bots/reference/machine-learning-models#model-versions-and-release-notes)",
				Optional:    true,
			},
			"enable_js": schema.BoolAttribute{
				Description: "Use lightweight, invisible JavaScript detections to improve Bot Management. [Learn more about JavaScript Detections](https://developers.cloudflare.com/bots/reference/javascript-detections/).",
				Optional:    true,
			},
			"fight_mode": schema.BoolAttribute{
				Description: "Whether to enable Bot Fight Mode.",
				Optional:    true,
			},
			"optimize_wordpress": schema.BoolAttribute{
				Description: "Whether to optimize Super Bot Fight Mode protections for Wordpress.",
				Optional:    true,
			},
			"sbfm_definitely_automated": schema.StringAttribute{
				Description: "Super Bot Fight Mode (SBFM) action to take on definitely automated requests.\nAvailable values: \"allow\", \"block\", \"managed_challenge\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"allow",
						"block",
						"managed_challenge",
					),
				},
			},
			"sbfm_likely_automated": schema.StringAttribute{
				Description: "Super Bot Fight Mode (SBFM) action to take on likely automated requests.\nAvailable values: \"allow\", \"block\", \"managed_challenge\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"allow",
						"block",
						"managed_challenge",
					),
				},
			},
			"sbfm_static_resource_protection": schema.BoolAttribute{
				Description: "Super Bot Fight Mode (SBFM) to enable static resource protection.\nEnable if static resources on your application need bot protection.\nNote: Static resource protection can also result in legitimate traffic being blocked.",
				Optional:    true,
			},
			"sbfm_verified_bots": schema.StringAttribute{
				Description: "Super Bot Fight Mode (SBFM) action to take on verified bots requests.\nAvailable values: \"allow\", \"block\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("allow", "block"),
				},
			},
			"suppress_session_score": schema.BoolAttribute{
				Description: "Whether to disable tracking the highest bot score for a session in the Bot Management cookie.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"using_latest_model": schema.BoolAttribute{
				Description: "A read-only field that indicates whether the zone currently is running the latest ML model.",
				Computed:    true,
			},
			"stale_zone_configuration": schema.SingleNestedAttribute{
				Description: "A read-only field that shows which unauthorized settings are currently active on the zone. These settings typically result from upgrades or downgrades.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[BotManagementStaleZoneConfigurationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"optimize_wordpress": schema.BoolAttribute{
						Description: "Indicates that the zone's wordpress optimization for SBFM is turned on.",
						Computed:    true,
					},
					"sbfm_definitely_automated": schema.StringAttribute{
						Description: "Indicates that the zone's definitely automated requests are being blocked or challenged.",
						Computed:    true,
					},
					"sbfm_likely_automated": schema.StringAttribute{
						Description: "Indicates that the zone's likely automated requests are being blocked or challenged.",
						Computed:    true,
					},
					"sbfm_static_resource_protection": schema.StringAttribute{
						Description: "Indicates that the zone's static resource protection is turned on.",
						Computed:    true,
					},
					"sbfm_verified_bots": schema.StringAttribute{
						Description: "Indicates that the zone's verified bot requests are being blocked.",
						Computed:    true,
					},
					"suppress_session_score": schema.BoolAttribute{
						Description: "Indicates that the zone's session score tracking is disabled.",
						Computed:    true,
					},
					"fight_mode": schema.BoolAttribute{
						Description: "Indicates that the zone's Bot Fight Mode is turned on.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *BotManagementResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *BotManagementResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
