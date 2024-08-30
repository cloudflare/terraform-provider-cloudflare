// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
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
			"deployment_configs": schema.SingleNestedAttribute{
				Description: "Configs for deployments in a project.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"preview": schema.SingleNestedAttribute{
						Description: "Configs for preview deploys.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"analytics_engine_datasets": schema.MapAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"browsers": schema.MapAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Computed:    true,
								Optional:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"durable_object_namespaces": schema.MapAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"env_vars": schema.MapAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"hyperdrive_bindings": schema.MapAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"kv_namespaces": schema.MapAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"mtls_certificates": schema.MapAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsPreviewPlacementModel](ctx),
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.MapAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"r2_buckets": schema.MapAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"services": schema.MapAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"vectorize_bindings": schema.MapAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
						},
					},
					"production": schema.SingleNestedAttribute{
						Description: "Configs for production deploys.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.MapAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"analytics_engine_datasets": schema.MapAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"browsers": schema.MapAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"compatibility_date": schema.StringAttribute{
								Description: "Compatibility date used for Pages Functions.",
								Computed:    true,
								Optional:    true,
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"durable_object_namespaces": schema.MapAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"env_vars": schema.MapAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"hyperdrive_bindings": schema.MapAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"kv_namespaces": schema.MapAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"mtls_certificates": schema.MapAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[PagesProjectDeploymentConfigsProductionPlacementModel](ctx),
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.MapAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"r2_buckets": schema.MapAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"services": schema.MapAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
							},
							"vectorize_bindings": schema.MapAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								ElementType: jsontypes.NormalizedType{},
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
						ElementType: types.StringType,
					},
					"build_config": schema.SingleNestedAttribute{
						Description: "Configs for the project build process.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentBuildConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentDeploymentTriggerModel](ctx),
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description: "Additional info about the trigger.",
								Computed:    true,
								Optional:    true,
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
						CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceConfigModel](ctx),
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
						ElementType: types.StringType,
					},
					"build_config": schema.SingleNestedAttribute{
						Description: "Configs for the project build process.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentBuildConfigModel](ctx),
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
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentDeploymentTriggerModel](ctx),
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description: "Additional info about the trigger.",
								Computed:    true,
								Optional:    true,
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
						CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								Optional:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceConfigModel](ctx),
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
			"source": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[PagesProjectSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[PagesProjectSourceConfigModel](ctx),
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
		},
	}
}

func (r *PagesProjectResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *PagesProjectResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
