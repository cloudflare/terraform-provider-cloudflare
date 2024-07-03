// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &PagesProjectsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &PagesProjectsDataSource{}

func (r PagesProjectsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						"id": schema.StringAttribute{
							Description: "Id of the deployment.",
							Computed:    true,
						},
						"aliases": schema.ListAttribute{
							Description: "A list of alias URLs pointing to this deployment.",
							Computed:    true,
							ElementType: types.StringType,
						},
						"build_config": schema.StringAttribute{
							Computed: true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the deployment was created.",
							Computed:    true,
						},
						"env_vars": schema.StringAttribute{
							Description: "A dict of env variables to build this deploy.",
							Computed:    true,
						},
						"environment": schema.StringAttribute{
							Description: "Type of deploy.",
							Computed:    true,
						},
						"is_skipped": schema.BoolAttribute{
							Description: "If the deployment has been skipped.",
							Computed:    true,
						},
						"latest_stage": schema.StringAttribute{
							Computed: true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the deployment was last modified.",
							Computed:    true,
						},
						"project_id": schema.StringAttribute{
							Description: "Id of the project.",
							Computed:    true,
						},
						"project_name": schema.StringAttribute{
							Description: "Name of the project.",
							Computed:    true,
						},
						"short_id": schema.StringAttribute{
							Description: "Short Id (8 character) of the deployment.",
							Computed:    true,
						},
						"source": schema.StringAttribute{
							Computed: true,
						},
						"stages": schema.ListNestedAttribute{
							Description: "List of past stages.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ended_on": schema.StringAttribute{
										Description: "When the stage ended.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "The current build stage.",
										Computed:    true,
									},
									"started_on": schema.StringAttribute{
										Description: "When the stage started.",
										Computed:    true,
									},
									"status": schema.StringAttribute{
										Description: "State of the current stage.",
										Computed:    true,
									},
								},
							},
						},
						"url": schema.StringAttribute{
							Description: "The live URL to view this deployment.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *PagesProjectsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *PagesProjectsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
