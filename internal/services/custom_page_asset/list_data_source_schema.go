// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_page_asset

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomPageAssetsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[CustomPageAssetsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique name of the custom asset. Can only contain letters (A-Z, a-z), numbers (0-9), and underscores (_).",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A short description of the custom asset.",
							Computed:    true,
						},
						"last_updated": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"name": schema.StringAttribute{
							Description: "The unique name of the custom asset. Can only contain letters (A-Z, a-z), numbers (0-9), and underscores (_).",
							Computed:    true,
						},
						"size_bytes": schema.Int64Attribute{
							Description: "The size of the asset content in bytes.",
							Computed:    true,
						},
						"url": schema.StringAttribute{
							Description: "The URL where the asset content is fetched from.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *CustomPageAssetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CustomPageAssetsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
