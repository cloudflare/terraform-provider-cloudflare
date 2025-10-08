// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package regional_hostname

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*RegionalHostnameDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are supported for one level, e.g `*.example.com`",
				Computed:    true,
			},
			"hostname": schema.StringAttribute{
				Description: "DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are supported for one level, e.g `*.example.com`",
				Computed:    true,
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the regional hostname was created",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"region_key": schema.StringAttribute{
				Description: "Identifying key for the region",
				Computed:    true,
			},
			"routing": schema.StringAttribute{
				Description: "Configure which routing method to use for the regional hostname",
				Computed:    true,
			},
		},
	}
}

func (d *RegionalHostnameDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *RegionalHostnameDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
