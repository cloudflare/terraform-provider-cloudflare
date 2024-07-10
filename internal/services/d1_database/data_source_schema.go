// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &D1DatabaseDataSource{}
var _ datasource.DataSourceWithValidateConfig = &D1DatabaseDataSource{}

func (r D1DatabaseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"database_id": schema.StringAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Description: "Specifies the timestamp the resource was created as an ISO8601 string.",
				Computed:    true,
			},
			"file_size": schema.Float64Attribute{
				Description: "The D1 database's size, in bytes.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"num_tables": schema.Float64Attribute{
				Optional: true,
			},
			"uuid": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"version": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
					"name": schema.StringAttribute{
						Description: "a database name to search for.",
						Optional:    true,
					},
					"page": schema.Float64Attribute{
						Description: "Page number of paginated results.",
						Computed:    true,
						Optional:    true,
					},
					"per_page": schema.Float64Attribute{
						Description: "Number of items per page.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
		},
	}
}

func (r *D1DatabaseDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *D1DatabaseDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
