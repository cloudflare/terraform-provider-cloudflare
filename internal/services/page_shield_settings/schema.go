// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield_settings

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*PageShieldSettingsResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 500,
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Domain Page Shield",
				"Domain Page Shield Read",
				"Page Shield",
				"Page Shield Read",
				"Zone Settings Read",
				"Zone Settings Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"enabled": schema.BoolAttribute{
				Description: "When true, indicates that Page Shield is enabled. Page Shield must be enabled for continuous script monitoring, malicious script detection, and code change detection to work.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"updated_at": schema.StringAttribute{
				Description: "The timestamp of when the setting was last updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"use_cloudflare_reporting_endpoint": schema.BoolAttribute{
				Description: "When true, CSP reports will be sent to the Cloudflare-managed reporting endpoint (https://csp-reporting.cloudflare.com/cdn-cgi/script_monitor/report) so detected resources appear in the Page Shield dashboard.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"use_connection_url_path": schema.BoolAttribute{
				Description: "When true, the paths associated with connection URLs will also be analyzed.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *PageShieldSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PageShieldSettingsResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
