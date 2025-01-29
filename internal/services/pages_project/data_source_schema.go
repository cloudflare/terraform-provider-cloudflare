// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*PagesProjectDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"project_name": schema.StringAttribute{
				Description: "Name of the project.",
				Required:    true,
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
			"name": schema.StringAttribute{
				Description: "Name of the project.",
				Computed:    true,
			},
			"production_branch": schema.StringAttribute{
				Description: "Production branch of the project. Used to identify production deployments.",
				Computed:    true,
			},
			"subdomain": schema.StringAttribute{
				Description: "The Cloudflare subdomain associated with the project.",
				Computed:    true,
			},
			"domains": schema.ListAttribute{
				Description: "A list of associated custom domains for the project.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
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
			"canonical_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment to the repo.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Id of the deployment.",
						Computed:    true,
					},
					"aliases": schema.ListAttribute{
						Description: "A list of alias URLs pointing to this deployment.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
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
					"env_vars": schema.MapNestedAttribute{
						Description: "A dict of env variables to build this deploy.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectMapType[PagesProjectCanonicalDeploymentEnvVarsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"value": schema.StringAttribute{
									Description: "Environment variable value.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The type of environment variable.",
									Computed:    true,
								},
							},
						},
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
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"path_includes": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"pr_comments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"preview_branch_excludes": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"preview_branch_includes": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
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
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"preview": schema.SingleNestedAttribute{
						Description: "Configs for preview deploys.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapNestedAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewAIBindingsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"analytics_engine_datasets": schema.MapNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dataset": schema.StringAttribute{
											Description: "Name of the dataset.",
											Computed:    true,
										},
									},
								},
							},
							"browsers": schema.MapNestedAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewBrowsersDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{},
								},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Computed:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewD1DatabasesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "UUID of the D1 database.",
											Computed:    true,
										},
									},
								},
							},
							"durable_object_namespaces": schema.MapNestedAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the Durabble Object namespace.",
											Computed:    true,
										},
									},
								},
							},
							"env_vars": schema.MapNestedAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewEnvVarsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description: "Environment variable value.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: "The type of environment variable.",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
											},
										},
									},
								},
							},
							"hyperdrive_bindings": schema.MapNestedAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewHyperdriveBindingsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"kv_namespaces": schema.MapNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewKVNamespacesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the KV namespace.",
											Computed:    true,
										},
									},
								},
							},
							"mtls_certificates": schema.MapNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewMTLSCertificatesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
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
							"queue_producers": schema.MapNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewQueueProducersDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the Queue.",
											Computed:    true,
										},
									},
								},
							},
							"r2_buckets": schema.MapNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewR2BucketsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"jurisdiction": schema.StringAttribute{
											Description: "Jurisdiction of the R2 bucket.",
											Computed:    true,
										},
										"name": schema.StringAttribute{
											Description: "Name of the R2 bucket.",
											Computed:    true,
										},
									},
								},
							},
							"services": schema.MapNestedAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewServicesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"entrypoint": schema.StringAttribute{
											Description: "The entrypoint to bind to.",
											Computed:    true,
										},
										"environment": schema.StringAttribute{
											Description: "The Service environment.",
											Computed:    true,
										},
										"service": schema.StringAttribute{
											Description: "The Service name.",
											Computed:    true,
										},
									},
								},
							},
							"vectorize_bindings": schema.MapNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewVectorizeBindingsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"index_name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"production": schema.SingleNestedAttribute{
						Description: "Configs for production deploys.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionDataSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapNestedAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionAIBindingsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"analytics_engine_datasets": schema.MapNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dataset": schema.StringAttribute{
											Description: "Name of the dataset.",
											Computed:    true,
										},
									},
								},
							},
							"browsers": schema.MapNestedAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionBrowsersDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{},
								},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Computed:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionD1DatabasesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "UUID of the D1 database.",
											Computed:    true,
										},
									},
								},
							},
							"durable_object_namespaces": schema.MapNestedAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionDurableObjectNamespacesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the Durabble Object namespace.",
											Computed:    true,
										},
									},
								},
							},
							"env_vars": schema.MapNestedAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionEnvVarsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description: "Environment variable value.",
											Computed:    true,
										},
										"type": schema.StringAttribute{
											Description: "The type of environment variable.",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
											},
										},
									},
								},
							},
							"hyperdrive_bindings": schema.MapNestedAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionHyperdriveBindingsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
							"kv_namespaces": schema.MapNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionKVNamespacesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the KV namespace.",
											Computed:    true,
										},
									},
								},
							},
							"mtls_certificates": schema.MapNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionMTLSCertificatesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate_id": schema.StringAttribute{
											Computed: true,
										},
									},
								},
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
							"queue_producers": schema.MapNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionQueueProducersDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the Queue.",
											Computed:    true,
										},
									},
								},
							},
							"r2_buckets": schema.MapNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionR2BucketsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"jurisdiction": schema.StringAttribute{
											Description: "Jurisdiction of the R2 bucket.",
											Computed:    true,
										},
										"name": schema.StringAttribute{
											Description: "Name of the R2 bucket.",
											Computed:    true,
										},
									},
								},
							},
							"services": schema.MapNestedAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionServicesDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"entrypoint": schema.StringAttribute{
											Description: "The entrypoint to bind to.",
											Computed:    true,
										},
										"environment": schema.StringAttribute{
											Description: "The Service environment.",
											Computed:    true,
										},
										"service": schema.StringAttribute{
											Description: "The Service name.",
											Computed:    true,
										},
									},
								},
							},
							"vectorize_bindings": schema.MapNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionVectorizeBindingsDataSourceModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"index_name": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
				},
			},
			"latest_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment to the repo.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Description: "Id of the deployment.",
						Computed:    true,
					},
					"aliases": schema.ListAttribute{
						Description: "A list of alias URLs pointing to this deployment.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
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
					"env_vars": schema.MapNestedAttribute{
						Description: "A dict of env variables to build this deploy.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectMapType[PagesProjectLatestDeploymentEnvVarsDataSourceModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"value": schema.StringAttribute{
									Description: "Environment variable value.",
									Computed:    true,
								},
								"type": schema.StringAttribute{
									Description: "The type of environment variable.",
									Computed:    true,
								},
							},
						},
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
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"path_includes": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"pr_comments_enabled": schema.BoolAttribute{
										Computed: true,
									},
									"preview_branch_excludes": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"preview_branch_includes": schema.ListAttribute{
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
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
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"path_includes": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"pr_comments_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"preview_branch_excludes": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"preview_branch_includes": schema.ListAttribute{
								Computed:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
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
		},
	}
}

func (d *PagesProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PagesProjectDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
