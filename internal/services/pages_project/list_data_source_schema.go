// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &PagesProjectsDataSource{}

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
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
						"build_config": schema.SingleNestedAttribute{
							Description: "Configs for the project build process.",
							Computed:    true,
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"build_caching": schema.BoolAttribute{
									Description: "Enable build caching for the project.",
									Computed:    true,
									Optional:    true,
								},
								"build_command": schema.StringAttribute{
									Description: "Command used to build project.",
									Computed:    true,
									Optional:    true,
								},
								"destination_dir": schema.StringAttribute{
									Description: "Output directory of the build.",
									Computed:    true,
									Optional:    true,
								},
								"root_dir": schema.StringAttribute{
									Description: "Directory to run the command.",
									Computed:    true,
									Optional:    true,
								},
								"web_analytics_tag": schema.StringAttribute{
									Description: "The classifying tag for analytics.",
									Computed:    true,
									Optional:    true,
								},
								"web_analytics_token": schema.StringAttribute{
									Description: "The auth token for analytics.",
									Computed:    true,
									Optional:    true,
								},
							},
						},
						"created_on": schema.StringAttribute{
							Description: "When the deployment was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"deployment_trigger": schema.SingleNestedAttribute{
							Description: "Info about what caused the deployment.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[PagesProjectsDeploymentTriggerDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"metadata": schema.SingleNestedAttribute{
									Description: "Additional info about the trigger.",
									Computed:    true,
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"branch": schema.StringAttribute{
											Description: "Where the trigger happened.",
											Computed:    true,
										},
										"commit_hash": schema.StringAttribute{
											Description: "Hash of the deployment trigger commit.",
											Computed:    true,
										},
										"commit_message": schema.StringAttribute{
											Description: "Message of the deployment trigger commit.",
											Computed:    true,
										},
									},
								},
								"type": schema.StringAttribute{
									Description: "What caused the deployment.",
									Computed:    true,
								},
							},
						},
						"env_vars": schema.MapAttribute{
							Description: "A dict of env variables to build this deploy.",
							Computed:    true,
							ElementType: jsontypes.NormalizedType{},
						},
						"environment": schema.StringAttribute{
							Description: "Type of deploy.",
							Computed:    true,
						},
						"is_skipped": schema.BoolAttribute{
							Description: "If the deployment has been skipped.",
							Computed:    true,
						},
						"latest_stage": schema.SingleNestedAttribute{
							Description: "The status of the deployment.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[PagesProjectsLatestStageDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"ended_on": schema.StringAttribute{
									Description: "When the stage ended.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"name": schema.StringAttribute{
									Description: "The current build stage.",
									Computed:    true,
									Optional:    true,
								},
								"started_on": schema.StringAttribute{
									Description: "When the stage started.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"status": schema.StringAttribute{
									Description: "State of the current stage.",
									Computed:    true,
								},
							},
						},
						"modified_on": schema.StringAttribute{
							Description: "When the deployment was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
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
						"source": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[PagesProjectsSourceDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"config": schema.SingleNestedAttribute{
									Computed: true,
									Optional: true,
									Attributes: map[string]schema.Attribute{
										"deployments_enabled": schema.BoolAttribute{
											Computed: true,
											Optional: true,
										},
										"owner": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
										"path_excludes": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
										"path_includes": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
										"pr_comments_enabled": schema.BoolAttribute{
											Computed: true,
											Optional: true,
										},
										"preview_branch_excludes": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
										"preview_branch_includes": schema.ListAttribute{
											Computed:    true,
											Optional:    true,
											ElementType: types.StringType,
										},
										"preview_deployment_setting": schema.StringAttribute{
											Computed: true,
											Optional: true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"all",
													"none",
													"custom",
												),
											},
										},
										"production_branch": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
										"production_deployments_enabled": schema.BoolAttribute{
											Computed: true,
											Optional: true,
										},
										"repo_name": schema.StringAttribute{
											Computed: true,
											Optional: true,
										},
									},
								},
								"type": schema.StringAttribute{
									Computed: true,
									Optional: true,
								},
							},
						},
						"stages": schema.ListNestedAttribute{
							Description: "List of past stages.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ended_on": schema.StringAttribute{
										Description: "When the stage ended.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
									},
									"name": schema.StringAttribute{
										Description: "The current build stage.",
										Computed:    true,
										Optional:    true,
									},
									"started_on": schema.StringAttribute{
										Description: "When the stage started.",
										Computed:    true,
										CustomType:  timetypes.RFC3339Type{},
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

func (d *PagesProjectsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *PagesProjectsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
