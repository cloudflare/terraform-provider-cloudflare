// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_cloud_region

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*OriginCloudRegionDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"origin_ip": schema.StringAttribute{
				Required: true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "Time this mapping was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"region": schema.StringAttribute{
				Description: "Cloud vendor region identifier.",
				Computed:    true,
			},
			"vendor": schema.StringAttribute{
				Description: "Cloud vendor hosting the origin.\nAvailable values: \"aws\", \"azure\", \"gcp\", \"oci\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"aws",
						"azure",
						"gcp",
						"oci",
					),
				},
			},
		},
	}
}

func (d *OriginCloudRegionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *OriginCloudRegionDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
