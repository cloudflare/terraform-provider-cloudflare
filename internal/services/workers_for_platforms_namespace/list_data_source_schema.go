// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_namespace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &WorkersForPlatformsNamespacesDataSource{}
var _ datasource.DataSourceWithValidateConfig = &WorkersForPlatformsNamespacesDataSource{}

func (r WorkersForPlatformsNamespacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_by": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the script was created.",
							Computed:    true,
						},
						"modified_by": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the script was last modified.",
							Computed:    true,
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
				},
			},
		},
	}
}

func (r *WorkersForPlatformsNamespacesDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *WorkersForPlatformsNamespacesDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
