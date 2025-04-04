// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersForPlatformsDispatchNamespaceDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Computed:    true,
			},
			"dispatch_namespace": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"created_by": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description: "When the script was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified_by": schema.StringAttribute{
				Description: "Identifier.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the script was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"namespace_id": schema.StringAttribute{
				Description: "API Resource UUID tag.",
				Computed:    true,
			},
			"namespace_name": schema.StringAttribute{
				Description: "Name of the Workers for Platforms dispatch namespace.",
				Computed:    true,
			},
			"script_count": schema.Int64Attribute{
				Description: "The current number of scripts in this Dispatch Namespace",
				Computed:    true,
			},
		},
	}
}

func (d *WorkersForPlatformsDispatchNamespaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *WorkersForPlatformsDispatchNamespaceDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
