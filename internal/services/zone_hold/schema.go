// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_hold

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZoneHoldResource)(nil)

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
			"hold_after": schema.StringAttribute{
				Description: "If `hold_after` is provided and future-dated, the hold will be temporarily disabled,\nthen automatically re-enabled by the system at the time specified\nin this RFC3339-formatted timestamp. A past-dated `hold_after` value will have\nno effect on an existing, enabled hold. Providing an empty string will set its value\nto the current time.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"include_subdomains": schema.BoolAttribute{
				Description: "If `true`, the zone hold will extend to block any subdomain of the given zone, as well\nas SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with the hostname\n'example.com' and include_subdomains=true will block 'example.com',\n'staging.example.com', 'api.staging.example.com', etc.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func (r *ZoneHoldResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneHoldResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
