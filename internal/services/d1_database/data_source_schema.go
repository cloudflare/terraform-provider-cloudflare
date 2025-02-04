// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1_database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*D1DatabaseDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"database_id": schema.StringAttribute{
				Optional: true,
			},
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"created_at": schema.StringAttribute{
				Description: "Specifies the timestamp the resource was created as an ISO8601 string.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"file_size": schema.Float64Attribute{
				Description: "The D1 database's size, in bytes.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"num_tables": schema.Float64Attribute{
				Computed: true,
			},
			"uuid": schema.StringAttribute{
				Computed: true,
			},
			"version": schema.StringAttribute{
				Computed: true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description: "a database name to search for.",
						Optional:    true,
					},
				},
			},
		},
	}
}

func (d *D1DatabaseDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *D1DatabaseDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("database_id"), path.MatchRoot("filter")),
	}
}
