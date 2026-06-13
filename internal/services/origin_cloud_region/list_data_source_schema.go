// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_cloud_region

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*OriginCloudRegionsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[OriginCloudRegionsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The origin IP address (IPv4 or IPv6). Normalized to canonical form (RFC 5952 for IPv6).",
							Computed:    true,
						},
						"origin_ip": schema.StringAttribute{
							Description: "The origin IP address (IPv4 or IPv6). Normalized to canonical form (RFC 5952 for IPv6).",
							Computed:    true,
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
						"modified_on": schema.StringAttribute{
							Description: "Time this mapping was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *OriginCloudRegionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *OriginCloudRegionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
