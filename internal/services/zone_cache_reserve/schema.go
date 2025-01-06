// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zone_cache_reserve

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
)

var _ resource.ResourceWithConfigValidators = (*ZoneCacheReserveResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"value": schema.StringAttribute{
				Description: "Value of the Cache Reserve zone setting.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("on", "off"),
				},
				Default: stringdefault.StaticString("off"),
			},
			"editable": schema.BoolAttribute{
				Description: "Whether the setting is editable",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "Last time this setting was modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"zone_setting_id": schema.StringAttribute{
				Description: "ID of the zone setting.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("cache_reserve"),
				},
			},
		},
	}
}

func (r *ZoneCacheReserveResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZoneCacheReserveResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
