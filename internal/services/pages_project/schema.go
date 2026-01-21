// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description:   "Name of the project.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"production_branch": schema.StringAttribute{
				Description: "Production branch of the project. Used to identify production deployments.",
				Required:    true,
			},
		"build_config": schema.SingleNestedAttribute{
			Description: "Configs for the project build process.",
			Optional:    true,
			Computed:    true,
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
						Optional:    true,
						Computed:    true,
						Sensitive:   true,
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Description: "Configs for the project source control.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Required: true,
						Attributes: map[string]schema.Attribute{
							"deployments_enabled": schema.BoolAttribute{
								Description:        "Whether to enable automatic deployments when pushing to the source repository.\nWhen disabled, no deployments (production or preview) will be triggered automatically.",
								Computed:           true,
								Optional:           true,
								DeprecationMessage: "Use `production_deployments_enabled` and `preview_deployment_setting` for more granular control.",
							},
							"owner": schema.StringAttribute{
								Description: "The owner of the repository.",
								Computed:    true,
								Optional:    true,
							},
							"owner_id": schema.StringAttribute{
								Description: "The owner ID of the repository.",
								Computed:    true,
								Optional:    true,
							},
							"path_excludes": schema.ListAttribute{
								Description: "A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"path_includes": schema.ListAttribute{
								Description: "A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"pr_comments_enabled": schema.BoolAttribute{
								Description: "Whether to enable PR comments.",
								Computed:    true,
								Optional:    true,
							},
							"preview_branch_excludes": schema.ListAttribute{
								Description: "A list of branches that should not trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"preview_branch_includes": schema.ListAttribute{
								Description: "A list of branches that should trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
							},
							"preview_deployment_setting": schema.StringAttribute{
								Description: "Controls whether commits to preview branches trigger a preview deployment.\nAvailable values: \"all\", \"none\", \"custom\".",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"all",
										"none",
										"custom",
									),
								},
							},
							"production_branch": schema.StringAttribute{
								Description: "The production branch of the repository.",
								Computed:    true,
								Optional:    true,
							},
							"production_deployments_enabled": schema.BoolAttribute{
								Description: "Whether to trigger a production deployment on commits to the production branch.",
								Computed:    true,
								Optional:    true,
							},
							"repo_id": schema.StringAttribute{
								Description: "The ID of the repository.",
								Computed:    true,
								Optional:    true,
							},
							"repo_name": schema.StringAttribute{
								Description: "The name of the repository.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"type": schema.StringAttribute{
						Description: "The source control management provider.\nAvailable values: \"github\", \"gitlab\".",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("github", "gitlab"),
						},
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
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"always_use_latest_compatibility_date": schema.BoolAttribute{
								Description:   "Whether to always use the latest compatibility date for Pages Functions.",
								Computed:      true,
								Optional:      true,
								Default:       booldefault.StaticBool(false),
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"analytics_engine_datasets": schema.MapNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dataset": schema.StringAttribute{
											Description: "Name of the dataset.",
											Required:    true,
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
							"build_image_major_version": schema.Int64Attribute{
								Description:   "The major version of the build image to use for Pages Functions.",
								Computed:      true,
								Optional:      true,
								Default:       int64default.StaticInt64(3),
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"compatibility_date": schema.StringAttribute{
								Description:   "Compatibility date used for Pages Functions.",
								Computed:      true,
								Optional:      true,
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "UUID of the D1 database.",
											Required:    true,
										},
									},
								},
							},
							"durable_object_namespaces": schema.MapNestedAttribute{
								Description: "Durable Object namespaces used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the Durable Object namespace.",
											Required:    true,
										},
									},
								},
							},
							"env_vars": schema.MapNestedAttribute{
								Description: "Environment variables used for builds and Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description: `Available values: "plain_text", "secret_text".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
											},
										},
										"value": schema.StringAttribute{
											Description: "Environment variable value.",
											Required:    true,
											Sensitive:   true,
										},
									},
								},
							},
							"fail_open": schema.BoolAttribute{
								Description:   "Whether to fail open when the deployment config cannot be applied.",
								Computed:      true,
								Optional:      true,
								Default:       booldefault.StaticBool(true),
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"hyperdrive_bindings": schema.MapNestedAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"kv_namespaces": schema.MapNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the KV namespace.",
											Required:    true,
										},
									},
								},
							},
							"limits": schema.SingleNestedAttribute{
								Description: "Limits for Pages Functions.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"cpu_ms": schema.Int64Attribute{
										Description: "CPU time limit in milliseconds.",
										Required:    true,
									},
								},
							},
							"mtls_certificates": schema.MapNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.MapNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the Queue.",
											Required:    true,
										},
									},
								},
							},
							"r2_buckets": schema.MapNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the R2 bucket.",
											Required:    true,
										},
										"jurisdiction": schema.StringAttribute{
											Description: "Jurisdiction of the R2 bucket.",
											Optional:    true,
										},
									},
								},
							},
							"services": schema.MapNestedAttribute{
								Description: "Services used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"service": schema.StringAttribute{
											Description: "The Service name.",
											Required:    true,
										},
										"entrypoint": schema.StringAttribute{
											Description: "The entrypoint to bind to.",
											Optional:    true,
										},
										"environment": schema.StringAttribute{
											Description: "The Service environment.",
											Optional:    true,
										},
									},
								},
							},
							"usage_model": schema.StringAttribute{
								Description:        "The usage model for Pages Functions.\nAvailable values: \"standard\", \"bundled\", \"unbound\".",
								Computed:           true,
								Optional:           true,
								DeprecationMessage: "All new projects now use the Standard usage model.",
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"standard",
										"bundled",
										"unbound",
									),
								},
								Default: stringdefault.StaticString("standard"),
							},
							"vectorize_bindings": schema.MapNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"index_name": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"wrangler_config_hash": schema.StringAttribute{
								Description: "Hash of the Wrangler configuration used for the deployment.",
								Optional:    true,
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
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"always_use_latest_compatibility_date": schema.BoolAttribute{
								Description:   "Whether to always use the latest compatibility date for Pages Functions.",
								Computed:      true,
								Optional:      true,
								Default:       booldefault.StaticBool(false),
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"analytics_engine_datasets": schema.MapNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"dataset": schema.StringAttribute{
											Description: "Name of the dataset.",
											Required:    true,
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
							"build_image_major_version": schema.Int64Attribute{
								Description:   "The major version of the build image to use for Pages Functions.",
								Computed:      true,
								Optional:      true,
								Default:       int64default.StaticInt64(3),
								PlanModifiers: []planmodifier.Int64{int64planmodifier.UseStateForUnknown()},
							},
							"compatibility_date": schema.StringAttribute{
								Description:   "Compatibility date used for Pages Functions.",
								Computed:      true,
								Optional:      true,
								PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
							},
							"compatibility_flags": schema.ListAttribute{
								Description: "Compatibility flags used for Pages Functions.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"d1_databases": schema.MapNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Description: "UUID of the D1 database.",
											Required:    true,
										},
									},
								},
							},
							"durable_object_namespaces": schema.MapNestedAttribute{
								Description: "Durable Object namespaces used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the Durable Object namespace.",
											Required:    true,
										},
									},
								},
							},
							"env_vars": schema.MapNestedAttribute{
								Description: "Environment variables used for builds and Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description: `Available values: "plain_text", "secret_text".`,
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
											},
										},
										"value": schema.StringAttribute{
											Description: "Environment variable value.",
											Required:    true,
											Sensitive:   true,
										},
									},
								},
							},
							"fail_open": schema.BoolAttribute{
								Description:   "Whether to fail open when the deployment config cannot be applied.",
								Computed:      true,
								Optional:      true,
								Default:       booldefault.StaticBool(true),
								PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
							},
							"hyperdrive_bindings": schema.MapNestedAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"kv_namespaces": schema.MapNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"namespace_id": schema.StringAttribute{
											Description: "ID of the KV namespace.",
											Required:    true,
										},
									},
								},
							},
							"limits": schema.SingleNestedAttribute{
								Description: "Limits for Pages Functions.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"cpu_ms": schema.Int64Attribute{
										Description: "CPU time limit in milliseconds.",
										Required:    true,
									},
								},
							},
							"mtls_certificates": schema.MapNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate_id": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.MapNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the Queue.",
											Required:    true,
										},
									},
								},
							},
							"r2_buckets": schema.MapNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description: "Name of the R2 bucket.",
											Required:    true,
										},
										"jurisdiction": schema.StringAttribute{
											Description: "Jurisdiction of the R2 bucket.",
											Optional:    true,
										},
									},
								},
							},
							"services": schema.MapNestedAttribute{
								Description: "Services used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"service": schema.StringAttribute{
											Description: "The Service name.",
											Required:    true,
										},
										"entrypoint": schema.StringAttribute{
											Description: "The entrypoint to bind to.",
											Optional:    true,
										},
										"environment": schema.StringAttribute{
											Description: "The Service environment.",
											Optional:    true,
										},
									},
								},
							},
							"usage_model": schema.StringAttribute{
								Description:        "The usage model for Pages Functions.\nAvailable values: \"standard\", \"bundled\", \"unbound\".",
								Computed:           true,
								Optional:           true,
								DeprecationMessage: "All new projects now use the Standard usage model.",
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"standard",
										"bundled",
										"unbound",
									),
								},
								Default: stringdefault.StaticString("standard"),
							},
							"vectorize_bindings": schema.MapNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"index_name": schema.StringAttribute{
											Required: true,
										},
									},
								},
							},
							"wrangler_config_hash": schema.StringAttribute{
								Description: "Hash of the Wrangler configuration used for the deployment.",
								Optional:    true,
							},
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description:   "When the project was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"framework": schema.StringAttribute{
				Description:   "Framework the project is using.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"framework_version": schema.StringAttribute{
				Description:   "Version of the framework the project is using.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"preview_script_name": schema.StringAttribute{
				Description:   "Name of the preview script.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"production_script_name": schema.StringAttribute{
				Description:   "Name of the production script.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"subdomain": schema.StringAttribute{
				Description:   "The Cloudflare subdomain associated with the project.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"uses_functions": schema.BoolAttribute{
				Description:   "Whether the project uses functions.",
				Computed:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.UseStateForUnknown()},
			},
			"domains": schema.ListAttribute{
				Description:   "A list of associated custom domains for the project.",
				Computed:      true,
				CustomType:    customfield.NewListType[types.String](ctx),
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
			},
			"canonical_deployment": schema.SingleNestedAttribute{
				Description: "Most recent production deployment of the project.",
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
							"web_analytics_tag": schema.StringAttribute{
								Description: "The classifying tag for analytics.",
								Computed:    true,
							},
							"web_analytics_token": schema.StringAttribute{
								Description: "The auth token for analytics.",
								Computed:    true,
								Sensitive:   true,
							},
							"build_caching": schema.BoolAttribute{
								Description: "Enable build caching for the project.",
								Computed:    true,
							},
							"build_command": schema.StringAttribute{
								Description: "Command used to build project.",
								Computed:    true,
							},
							"destination_dir": schema.StringAttribute{
								Description: "Assets output directory of the build.",
								Computed:    true,
							},
							"root_dir": schema.StringAttribute{
								Description: "Directory to run the command.",
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
									"commit_dirty": schema.BoolAttribute{
										Description: "Whether the deployment trigger commit was dirty.",
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
								Description: "What caused the deployment.\nAvailable values: \"github:push\", \"ad_hoc\", \"deploy_hook\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"github:push",
										"ad_hoc",
										"deploy_hook",
									),
								},
							},
						},
					},
					"env_vars": schema.MapNestedAttribute{
						Description: "Environment variables used for builds and Pages Functions.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectMapType[PagesProjectCanonicalDeploymentEnvVarsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: `Available values: "plain_text", "secret_text".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
									},
								},
								"value": schema.StringAttribute{
									Description: "Environment variable value.",
									Computed:    true,
									Sensitive:   true,
								},
							},
						},
					},
					"environment": schema.StringAttribute{
						Description: "Type of deploy.\nAvailable values: \"preview\", \"production\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("preview", "production"),
						},
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
								Description: "The current build stage.\nAvailable values: \"queued\", \"initialize\", \"clone_repo\", \"build\", \"deploy\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"queued",
										"initialize",
										"clone_repo",
										"build",
										"deploy",
									),
								},
							},
							"started_on": schema.StringAttribute{
								Description: "When the stage started.",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"status": schema.StringAttribute{
								Description: "State of the current stage.\nAvailable values: \"success\", \"idle\", \"active\", \"failure\", \"canceled\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"success",
										"idle",
										"active",
										"failure",
										"canceled",
									),
								},
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
						Description: "Configs for the project source control.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectCanonicalDeploymentSourceConfigModel](ctx),
								Attributes: map[string]schema.Attribute{
									"deployments_enabled": schema.BoolAttribute{
										Description:        "Whether to enable automatic deployments when pushing to the source repository.\nWhen disabled, no deployments (production or preview) will be triggered automatically.",
										Computed:           true,
										DeprecationMessage: "Use `production_deployments_enabled` and `preview_deployment_setting` for more granular control.",
									},
									"owner": schema.StringAttribute{
										Description: "The owner of the repository.",
										Computed:    true,
									},
									"owner_id": schema.StringAttribute{
										Description: "The owner ID of the repository.",
										Computed:    true,
									},
									"path_excludes": schema.ListAttribute{
										Description: "A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"path_includes": schema.ListAttribute{
										Description: "A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"pr_comments_enabled": schema.BoolAttribute{
										Description: "Whether to enable PR comments.",
										Computed:    true,
									},
									"preview_branch_excludes": schema.ListAttribute{
										Description: "A list of branches that should not trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"preview_branch_includes": schema.ListAttribute{
										Description: "A list of branches that should trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"preview_deployment_setting": schema.StringAttribute{
										Description: "Controls whether commits to preview branches trigger a preview deployment.\nAvailable values: \"all\", \"none\", \"custom\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"all",
												"none",
												"custom",
											),
										},
									},
									"production_branch": schema.StringAttribute{
										Description: "The production branch of the repository.",
										Computed:    true,
									},
									"production_deployments_enabled": schema.BoolAttribute{
										Description: "Whether to trigger a production deployment on commits to the production branch.",
										Computed:    true,
									},
									"repo_id": schema.StringAttribute{
										Description: "The ID of the repository.",
										Computed:    true,
									},
									"repo_name": schema.StringAttribute{
										Description: "The name of the repository.",
										Computed:    true,
									},
								},
							},
							"type": schema.StringAttribute{
								Description: "The source control management provider.\nAvailable values: \"github\", \"gitlab\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("github", "gitlab"),
								},
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
									Description: "The current build stage.\nAvailable values: \"queued\", \"initialize\", \"clone_repo\", \"build\", \"deploy\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"queued",
											"initialize",
											"clone_repo",
											"build",
											"deploy",
										),
									},
								},
								"started_on": schema.StringAttribute{
									Description: "When the stage started.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"status": schema.StringAttribute{
									Description: "State of the current stage.\nAvailable values: \"success\", \"idle\", \"active\", \"failure\", \"canceled\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"success",
											"idle",
											"active",
											"failure",
											"canceled",
										),
									},
								},
							},
						},
					},
					"url": schema.StringAttribute{
						Description: "The live URL to view this deployment.",
						Computed:    true,
					},
					"uses_functions": schema.BoolAttribute{
						Description: "Whether the deployment uses functions.",
						Computed:    true,
					},
				},
			},
			"latest_deployment": schema.SingleNestedAttribute{
				Description: "Most recent deployment of the project.",
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
							"web_analytics_tag": schema.StringAttribute{
								Description: "The classifying tag for analytics.",
								Computed:    true,
							},
							"web_analytics_token": schema.StringAttribute{
								Description: "The auth token for analytics.",
								Computed:    true,
								Sensitive:   true,
							},
							"build_caching": schema.BoolAttribute{
								Description: "Enable build caching for the project.",
								Computed:    true,
							},
							"build_command": schema.StringAttribute{
								Description: "Command used to build project.",
								Computed:    true,
							},
							"destination_dir": schema.StringAttribute{
								Description: "Assets output directory of the build.",
								Computed:    true,
							},
							"root_dir": schema.StringAttribute{
								Description: "Directory to run the command.",
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
									"commit_dirty": schema.BoolAttribute{
										Description: "Whether the deployment trigger commit was dirty.",
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
								Description: "What caused the deployment.\nAvailable values: \"github:push\", \"ad_hoc\", \"deploy_hook\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"github:push",
										"ad_hoc",
										"deploy_hook",
									),
								},
							},
						},
					},
					"env_vars": schema.MapNestedAttribute{
						Description: "Environment variables used for builds and Pages Functions.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectMapType[PagesProjectLatestDeploymentEnvVarsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: `Available values: "plain_text", "secret_text".`,
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
									},
								},
								"value": schema.StringAttribute{
									Description: "Environment variable value.",
									Computed:    true,
									Sensitive:   true,
								},
							},
						},
					},
					"environment": schema.StringAttribute{
						Description: "Type of deploy.\nAvailable values: \"preview\", \"production\".",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("preview", "production"),
						},
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
								Description: "The current build stage.\nAvailable values: \"queued\", \"initialize\", \"clone_repo\", \"build\", \"deploy\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"queued",
										"initialize",
										"clone_repo",
										"build",
										"deploy",
									),
								},
							},
							"started_on": schema.StringAttribute{
								Description: "When the stage started.",
								Computed:    true,
								CustomType:  timetypes.RFC3339Type{},
							},
							"status": schema.StringAttribute{
								Description: "State of the current stage.\nAvailable values: \"success\", \"idle\", \"active\", \"failure\", \"canceled\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"success",
										"idle",
										"active",
										"failure",
										"canceled",
									),
								},
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
						Description: "Configs for the project source control.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceModel](ctx),
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Computed:   true,
								CustomType: customfield.NewNestedObjectType[PagesProjectLatestDeploymentSourceConfigModel](ctx),
								Attributes: map[string]schema.Attribute{
									"deployments_enabled": schema.BoolAttribute{
										Description:        "Whether to enable automatic deployments when pushing to the source repository.\nWhen disabled, no deployments (production or preview) will be triggered automatically.",
										Computed:           true,
										DeprecationMessage: "Use `production_deployments_enabled` and `preview_deployment_setting` for more granular control.",
									},
									"owner": schema.StringAttribute{
										Description: "The owner of the repository.",
										Computed:    true,
									},
									"owner_id": schema.StringAttribute{
										Description: "The owner ID of the repository.",
										Computed:    true,
									},
									"path_excludes": schema.ListAttribute{
										Description: "A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"path_includes": schema.ListAttribute{
										Description: "A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"pr_comments_enabled": schema.BoolAttribute{
										Description: "Whether to enable PR comments.",
										Computed:    true,
									},
									"preview_branch_excludes": schema.ListAttribute{
										Description: "A list of branches that should not trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"preview_branch_includes": schema.ListAttribute{
										Description: "A list of branches that should trigger a preview deployment. Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"preview_deployment_setting": schema.StringAttribute{
										Description: "Controls whether commits to preview branches trigger a preview deployment.\nAvailable values: \"all\", \"none\", \"custom\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"all",
												"none",
												"custom",
											),
										},
									},
									"production_branch": schema.StringAttribute{
										Description: "The production branch of the repository.",
										Computed:    true,
									},
									"production_deployments_enabled": schema.BoolAttribute{
										Description: "Whether to trigger a production deployment on commits to the production branch.",
										Computed:    true,
									},
									"repo_id": schema.StringAttribute{
										Description: "The ID of the repository.",
										Computed:    true,
									},
									"repo_name": schema.StringAttribute{
										Description: "The name of the repository.",
										Computed:    true,
									},
								},
							},
							"type": schema.StringAttribute{
								Description: "The source control management provider.\nAvailable values: \"github\", \"gitlab\".",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("github", "gitlab"),
								},
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
									Description: "The current build stage.\nAvailable values: \"queued\", \"initialize\", \"clone_repo\", \"build\", \"deploy\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"queued",
											"initialize",
											"clone_repo",
											"build",
											"deploy",
										),
									},
								},
								"started_on": schema.StringAttribute{
									Description: "When the stage started.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"status": schema.StringAttribute{
									Description: "State of the current stage.\nAvailable values: \"success\", \"idle\", \"active\", \"failure\", \"canceled\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"success",
											"idle",
											"active",
											"failure",
											"canceled",
										),
									},
								},
							},
						},
					},
					"url": schema.StringAttribute{
						Description: "The live URL to view this deployment.",
						Computed:    true,
					},
					"uses_functions": schema.BoolAttribute{
						Description: "Whether the deployment uses functions.",
						Computed:    true,
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
