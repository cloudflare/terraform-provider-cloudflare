// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PagesProjectDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Name of the project.",
				Computed:    true,
				Optional:    true,
			},
			"environment": schema.StringAttribute{
				Description: "Type of deploy.",
				Optional:    true,
			},
			"is_skipped": schema.BoolAttribute{
				Description: "If the deployment has been skipped.",
				Optional:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the deployment was last modified.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "Name of the project.",
				Optional:    true,
			},
			"production_branch": schema.StringAttribute{
				Description: "Production branch of the project. Used to identify production deployments.",
				Optional:    true,
			},
			"project_id": schema.StringAttribute{
				Description: "Id of the project.",
				Optional:    true,
			},
			"short_id": schema.StringAttribute{
				Description: "Short Id (8 character) of the deployment.",
				Optional:    true,
			},
			"subdomain": schema.StringAttribute{
				Description: "The Cloudflare subdomain associated with the project.",
				Optional:    true,
			},
			"url": schema.StringAttribute{
				Description: "The live URL to view this deployment.",
				Optional:    true,
			},
			"aliases": schema.ListAttribute{
				Description: "A list of alias URLs pointing to this deployment.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"domains": schema.ListAttribute{
				Description: "A list of associated custom domains for the project.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"env_vars": schema.MapAttribute{
				Description: "A dict of env variables to build this deploy.",
				Optional:    true,
				ElementType: jsontypes.NormalizedType{},
			},
			"canonical_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment to the repo.",
				Optional:    true,
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentBuildConfigDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"build_caching": schema.BoolAttribute{
								Description: "Enable build caching for the project.",
								Computed:    true,
							},
							"build_command": schema.StringAttribute{
								Description: "Command used to build project.",
								Computed:    true,
							},
							"destination_dir": schema.StringAttribute{
								Description: "Output directory of the build.",
								Computed:    true,
							},
							"root_dir": schema.StringAttribute{
								Description: "Directory to run the command.",
								Computed:    true,
							},
							"web_analytics_tag": schema.StringAttribute{
								Description: "The classifying tag for analytics.",
								Computed:    true,
							},
							"web_analytics_token": schema.StringAttribute{
								Description: "The auth token for analytics.",
								Computed:    true,
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentDeploymentTriggerDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description: "Additional info about the trigger.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentDeploymentTriggerMetadataDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentLatestStageDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ended_on": schema.StringAttribute{
								Description: "When the stage ended.",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"name": schema.StringAttribute{
								Description: "The current build stage.",
								Computed:    true,
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
						CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceConfigDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"deployments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"owner": schema.StringAttribute{
										Computed: true,
									},
									"path_excludes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"path_includes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"pr_comments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"preview_branch_excludes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"preview_branch_includes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"preview_deployment_setting": schema.StringAttribute{
										Computed: true,
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
									},
									"production_deployments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"repo_name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"stages": schema.ListNestedAttribute{
						Description: "List of past stages.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[PagesProjectCanonicalDeploymentStagesDataSourceModel](ctx),
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
			"deployment_configs": schema.SingleNestedAttribute{
				Description: "Configs for deployments in a project.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"preview": schema.SingleNestedAttribute{
						Description: "Configs for preview deploys.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"analytics_engine_datasets": schema.MapAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"browsers": schema.MapAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Computed:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Computed:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"durable_object_namespaces": schema.MapAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"env_vars": schema.MapAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"hyperdrive_bindings": schema.MapAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"kv_namespaces": schema.MapAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"mtls_certificates": schema.MapAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewPlacementDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Computed:    true,
									},
								},
							},
							"queue_producers": schema.MapAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"r2_buckets": schema.MapAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"services": schema.MapAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"vectorize_bindings": schema.MapAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
						},
					},
					"production": schema.SingleNestedAttribute{
						Description: "Configs for production deploys.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"analytics_engine_datasets": schema.MapAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"browsers": schema.MapAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Computed:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Computed:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"durable_object_namespaces": schema.MapAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"env_vars": schema.MapAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"hyperdrive_bindings": schema.MapAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"kv_namespaces": schema.MapAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"mtls_certificates": schema.MapAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionPlacementDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Computed:    true,
									},
								},
							},
							"queue_producers": schema.MapAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"r2_buckets": schema.MapAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"services": schema.MapAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"vectorize_bindings": schema.MapAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								ElementType: jsontypes.NormalizedType{},
							},
						},
					},
				},
			},
			"deployment_trigger": schema.SingleNestedAttribute{
				Description: "Info about what caused the deployment.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"metadata": schema.SingleNestedAttribute{
						Description: "Additional info about the trigger.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentTriggerMetadataDataSourceModel](ctx),
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
			"latest_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment to the repo.",
				Optional:    true,
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentBuildConfigDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"build_caching": schema.BoolAttribute{
								Description: "Enable build caching for the project.",
								Computed:    true,
							},
							"build_command": schema.StringAttribute{
								Description: "Command used to build project.",
								Computed:    true,
							},
							"destination_dir": schema.StringAttribute{
								Description: "Output directory of the build.",
								Computed:    true,
							},
							"root_dir": schema.StringAttribute{
								Description: "Directory to run the command.",
								Computed:    true,
							},
							"web_analytics_tag": schema.StringAttribute{
								Description: "The classifying tag for analytics.",
								Computed:    true,
							},
							"web_analytics_token": schema.StringAttribute{
								Description: "The auth token for analytics.",
								Computed:    true,
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentDeploymentTriggerDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description: "Additional info about the trigger.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentDeploymentTriggerMetadataDataSourceModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentLatestStageDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ended_on": schema.StringAttribute{
								Description: "When the stage ended.",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"name": schema.StringAttribute{
								Description: "The current build stage.",
								Computed:    true,
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
						CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceConfigDataSourceModel](ctx),
								Attributes: map[string]schema.Attribute{
									"deployments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"owner": schema.StringAttribute{
										Computed: true,
									},
									"path_excludes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"path_includes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"pr_comments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"preview_branch_excludes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"preview_branch_includes": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"preview_deployment_setting": schema.StringAttribute{
										Computed: true,
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
									},
									"production_deployments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"repo_name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"type": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"stages": schema.ListNestedAttribute{
						Description: "List of past stages.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[PagesProjectLatestDeploymentStagesDataSourceModel](ctx),
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
			"latest_stage": schema.SingleNestedAttribute{
				Description: "The status of the deployment.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"ended_on": schema.StringAttribute{
						Description: "When the stage ended.",
						Computed:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"name": schema.StringAttribute{
						Description: "The current build stage.",
						Computed:    true,
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
			"stages": schema.ListNestedAttribute{
				Description: "List of past stages.",
				Optional:    true,
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
			"created_on": schema.StringAttribute{
				Description: "When the project was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "Id of the project.",
				Computed:    true,
			},
			"build_config": schema.SingleNestedAttribute{
				Description: "Configs for the project build process.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectBuildConfigDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"build_caching": schema.BoolAttribute{
						Description: "Enable build caching for the project.",
						Computed:    true,
					},
					"build_command": schema.StringAttribute{
						Description: "Command used to build project.",
						Computed:    true,
					},
					"destination_dir": schema.StringAttribute{
						Description: "Output directory of the build.",
						Computed:    true,
					},
					"root_dir": schema.StringAttribute{
						Description: "Directory to run the command.",
						Computed:    true,
					},
					"web_analytics_tag": schema.StringAttribute{
						Description: "The classifying tag for analytics.",
						Computed:    true,
					},
					"web_analytics_token": schema.StringAttribute{
						Description: "The auth token for analytics.",
						Computed:    true,
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesProjectSourceDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[PagesProjectSourceConfigDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"deployments_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"owner": schema.StringAttribute{
								Computed: true,
							},
							"path_excludes": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"path_includes": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"pr_comments_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"preview_branch_excludes": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"preview_branch_includes": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"preview_deployment_setting": schema.StringAttribute{
								Computed: true,
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
							},
							"production_deployments_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"repo_name": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"type": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *PagesProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PagesProjectDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("project_name")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("project_name")),
	}
}
