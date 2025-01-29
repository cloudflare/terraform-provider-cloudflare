// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*PagesProjectResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Name of the project.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Name of the project.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"production_branch": schema.StringAttribute{
				Description: "Production branch of the project. Used to identify production deployments.",
				Optional:    true,
			},
			"build_config": schema.SingleNestedAttribute{
				Description: "Configs for the project build process.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectBuildConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"build_caching": schema.BoolAttribute{
						Description: "Enable build caching for the project.",
						Optional:    true,
					},
					"build_command": schema.StringAttribute{
						Description: "Command used to build project.",
						Optional:    true,
					},
					"destination_dir": schema.StringAttribute{
						Description: "Output directory of the build.",
						Optional:    true,
					},
					"root_dir": schema.StringAttribute{
						Description: "Directory to run the command.",
						Optional:    true,
					},
					"web_analytics_tag": schema.StringAttribute{
						Description: "The classifying tag for analytics.",
						Optional:    true,
					},
					"web_analytics_token": schema.StringAttribute{
						Description: "The auth token for analytics.",
						Optional:    true,
					},
				},
			},
			"deployment_configs": schema.SingleNestedAttribute{
				Description: "Configs for deployments in a project.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"preview": schema.SingleNestedAttribute{
						Description: "Configs for preview deploys.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapNestedAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewAIBindingsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"analytics_engine_datasets": schema.MapNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewAnalyticsEngineDatasetsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dataset": schema.StringAttribute{
											Description: "Name of the dataset.",
											Optional:    true,
										},
									},
								},
							},
							"browsers": schema.MapNestedAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{},
								},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Optional:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewD1DatabasesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "UUID of the D1 database.",
											Optional:    true,
										},
									},
								},
							},
							"durable_object_namespaces": schema.MapNestedAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewDurableObjectNamespacesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the Durabble Object namespace.",
											Optional:    true,
										},
									},
								},
							},
							"env_vars": schema.MapNestedAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewEnvVarsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description: "Environment variable value.",
											Required:    true,
										},
										"type": schema.StringAttribute{
											Description: "The type of environment variable.",
											Optional:    true,
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
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewHyperdriveBindingsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"kv_namespaces": schema.MapNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewKVNamespacesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the KV namespace.",
											Optional:    true,
										},
									},
								},
							},
							"mtls_certificates": schema.MapNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewMTLSCertificatesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate_id": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewPlacementModel](ctx),
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.MapNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewQueueProducersModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the Queue.",
											Optional:    true,
										},
									},
								},
							},
							"r2_buckets": schema.MapNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewR2BucketsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"jurisdiction": schema.StringAttribute{
											Description: "Jurisdiction of the R2 bucket.",
											Optional:    true,
										},
										"name": schema.StringAttribute{
											Description: "Name of the R2 bucket.",
											Optional:    true,
										},
									},
								},
							},
							"services": schema.MapNestedAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewServicesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"entrypoint": schema.StringAttribute{
											Description: "The entrypoint to bind to.",
											Optional:    true,
										},
										"environment": schema.StringAttribute{
											Description: "The Service environment.",
											Optional:    true,
										},
										"service": schema.StringAttribute{
											Description: "The Service name.",
											Optional:    true,
										},
									},
								},
							},
							"vectorize_bindings": schema.MapNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsPreviewVectorizeBindingsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"index_name": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
					"production": schema.SingleNestedAttribute{
						Description: "Configs for production deploys.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapNestedAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionAIBindingsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"analytics_engine_datasets": schema.MapNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionAnalyticsEngineDatasetsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dataset": schema.StringAttribute{
											Description: "Name of the dataset.",
											Optional:    true,
										},
									},
								},
							},
							"browsers": schema.MapNestedAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{},
								},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Optional:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionD1DatabasesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "UUID of the D1 database.",
											Optional:    true,
										},
									},
								},
							},
							"durable_object_namespaces": schema.MapNestedAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionDurableObjectNamespacesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the Durabble Object namespace.",
											Optional:    true,
										},
									},
								},
							},
							"env_vars": schema.MapNestedAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionEnvVarsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"value": schema.StringAttribute{
											Description: "Environment variable value.",
											Required:    true,
										},
										"type": schema.StringAttribute{
											Description: "The type of environment variable.",
											Optional:    true,
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
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionHyperdriveBindingsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"kv_namespaces": schema.MapNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionKVNamespacesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the KV namespace.",
											Optional:    true,
										},
									},
								},
							},
							"mtls_certificates": schema.MapNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionMTLSCertificatesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate_id": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionPlacementModel](ctx),
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.MapNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionQueueProducersModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the Queue.",
											Optional:    true,
										},
									},
								},
							},
							"r2_buckets": schema.MapNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionR2BucketsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"jurisdiction": schema.StringAttribute{
											Description: "Jurisdiction of the R2 bucket.",
											Optional:    true,
										},
										"name": schema.StringAttribute{
											Description: "Name of the R2 bucket.",
											Optional:    true,
										},
									},
								},
							},
							"services": schema.MapNestedAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionServicesModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"entrypoint": schema.StringAttribute{
											Description: "The entrypoint to bind to.",
											Optional:    true,
										},
										"environment": schema.StringAttribute{
											Description: "The Service environment.",
											Optional:    true,
										},
										"service": schema.StringAttribute{
											Description: "The Service name.",
											Optional:    true,
										},
									},
								},
							},
							"vectorize_bindings": schema.MapNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectMapType[PagesProjectDeploymentConfigsProductionVectorizeBindingsModel](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"index_name": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the project was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
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
			"canonical_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment to the repo.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentBuildConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentDeploymentTriggerModel](ctx),
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description: "Additional info about the trigger.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentDeploymentTriggerMetadataModel](ctx),
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
						CustomType:  customfield.NewNestedObjectMapType[PagesProjectCanonicalDeploymentEnvVarsModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentLatestStageModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[PagesProjectCanonicalDeploymentStagesModel](ctx),
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
			"latest_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment to the repo.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentBuildConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentDeploymentTriggerModel](ctx),
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description: "Additional info about the trigger.",
								Computed:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentDeploymentTriggerMetadataModel](ctx),
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
						CustomType:  customfield.NewNestedObjectMapType[PagesProjectLatestDeploymentEnvVarsModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentLatestStageModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectListType[PagesProjectLatestDeploymentStagesModel](ctx),
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
				CustomType: customfield.NewNestedObjectType[PagesProjectSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[PagesProjectSourceConfigModel](ctx),
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

func (r *PagesProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PagesProjectResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
