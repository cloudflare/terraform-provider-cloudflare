// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*ZoneSettingResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Setting name",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"setting_id": schema.StringAttribute{
				Description:   "Setting name",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"value": schema.DynamicAttribute{
				Description: "Current value of the zone setting.",
				Required:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "ssl-recommender enrollment setting.",
				Computed:    true,
				Optional:    true,
			},
			"editable": schema.BoolAttribute{
				Description: "Whether or not this setting can be modified for this zone (based on your Cloudflare plan level).",
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"modified_on": schema.StringAttribute{
				Description: "last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"time_remaining": schema.Float64Attribute{
				Description: "Value of the zone setting.\nNotes: The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is false.",
				Computed:    true,
			},
		},
	}
}

func (r *ZoneSettingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneSettingResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
