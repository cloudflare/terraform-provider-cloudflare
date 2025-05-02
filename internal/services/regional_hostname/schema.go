// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.ResourceWithConfigValidators = (*RegionalHostnameResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are supported for one level, e.g `*.example.com`",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"hostname": schema.StringAttribute{
				Description:   "DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are supported for one level, e.g `*.example.com`",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"zone_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"routing": schema.StringAttribute{
				Description:   "Configure which routing method to use for the regional hostname",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"region_key": schema.StringAttribute{
				Description: "Identifying key for the region",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the regional hostname was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *RegionalHostnameResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *RegionalHostnameResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
