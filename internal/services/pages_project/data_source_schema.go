// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = &PagesProjectDataSource{}
var _ datasource.DataSourceWithValidateConfig = &PagesProjectDataSource{}

func (r PagesProjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
			"id": schema.StringAttribute{
				Description: "Id of the project.",
				Computed:    true,
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
			"canonical_deployment": schema.SingleNestedAttribute{
				Optional: true,
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
									Optional:    true,
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
			"created_on": schema.StringAttribute{
				Description: "When the project was created.",
				Computed:    true,
			},
			"deployment_configs": schema.SingleNestedAttribute{
				Description: "Configs for deployments in a project.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"preview": schema.SingleNestedAttribute{
						Description: "Configs for preview deploys.",
						Computed:    true,
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.SingleNestedAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"ai_binding": schema.SingleNestedAttribute{
										Description: "AI binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"project_id": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
							"analytics_engine_datasets": schema.SingleNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"analytics_engine_binding": schema.SingleNestedAttribute{
										Description: "Analytics Engine binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"dataset": schema.StringAttribute{
												Description: "Name of the dataset.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"browsers": schema.SingleNestedAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"browser": schema.StringAttribute{
										Description: "Browser binding.",
										Computed:    true,
										Optional:    true,
									},
								},
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
							"d1_databases": schema.SingleNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"d1_binding": schema.SingleNestedAttribute{
										Description: "D1 binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "UUID of the D1 database.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"durable_object_namespaces": schema.SingleNestedAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"do_binding": schema.SingleNestedAttribute{
										Description: "Durabble Object binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"namespace_id": schema.StringAttribute{
												Description: "ID of the Durabble Object namespace.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"env_vars": schema.SingleNestedAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"environment_variable": schema.SingleNestedAttribute{
										Description: "Environment variable.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												Description: "The type of environment variable (plain text or secret)",
												Computed:    true,
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
												},
											},
											"value": schema.StringAttribute{
												Description: "Environment variable value.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"hyperdrive_bindings": schema.SingleNestedAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"hyperdrive": schema.SingleNestedAttribute{
										Description: "Hyperdrive binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
							"kv_namespaces": schema.SingleNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"kv_binding": schema.SingleNestedAttribute{
										Description: "KV binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"namespace_id": schema.StringAttribute{
												Description: "ID of the KV namespace.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"mtls_certificates": schema.SingleNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"mtls": schema.SingleNestedAttribute{
										Description: "mTLS binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"certificate_id": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.SingleNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"queue_producer_binding": schema.SingleNestedAttribute{
										Description: "Queue Producer binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "Name of the Queue.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"r2_buckets": schema.SingleNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"r2_binding": schema.SingleNestedAttribute{
										Description: "R2 binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"jurisdiction": schema.StringAttribute{
												Description: "Jurisdiction of the R2 bucket.",
												Computed:    true,
												Optional:    true,
											},
											"name": schema.StringAttribute{
												Description: "Name of the R2 bucket.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"services": schema.SingleNestedAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"service_binding": schema.SingleNestedAttribute{
										Description: "Service binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"entrypoint": schema.StringAttribute{
												Description: "The entrypoint to bind to.",
												Computed:    true,
												Optional:    true,
											},
											"environment": schema.StringAttribute{
												Description: "The Service environment.",
												Computed:    true,
												Optional:    true,
											},
											"service": schema.StringAttribute{
												Description: "The Service name.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"vectorize_bindings": schema.SingleNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"vectorize": schema.SingleNestedAttribute{
										Description: "Vectorize binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"index_name": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
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
						Attributes: map[string]schema.Attribute{
							"ai_bindings": schema.SingleNestedAttribute{
								Description: "Constellation bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"ai_binding": schema.SingleNestedAttribute{
										Description: "AI binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"project_id": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
							"analytics_engine_datasets": schema.SingleNestedAttribute{
								Description: "Analytics Engine bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"analytics_engine_binding": schema.SingleNestedAttribute{
										Description: "Analytics Engine binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"dataset": schema.StringAttribute{
												Description: "Name of the dataset.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"browsers": schema.SingleNestedAttribute{
								Description: "Browser bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"browser": schema.StringAttribute{
										Description: "Browser binding.",
										Computed:    true,
										Optional:    true,
									},
								},
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
							"d1_databases": schema.SingleNestedAttribute{
								Description: "D1 databases used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"d1_binding": schema.SingleNestedAttribute{
										Description: "D1 binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description: "UUID of the D1 database.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"durable_object_namespaces": schema.SingleNestedAttribute{
								Description: "Durabble Object namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"do_binding": schema.SingleNestedAttribute{
										Description: "Durabble Object binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"namespace_id": schema.StringAttribute{
												Description: "ID of the Durabble Object namespace.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"env_vars": schema.SingleNestedAttribute{
								Description: "Environment variables for build configs.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"environment_variable": schema.SingleNestedAttribute{
										Description: "Environment variable.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												Description: "The type of environment variable (plain text or secret)",
												Computed:    true,
												Optional:    true,
												Validators: []validator.String{
													stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
												},
											},
											"value": schema.StringAttribute{
												Description: "Environment variable value.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"hyperdrive_bindings": schema.SingleNestedAttribute{
								Description: "Hyperdrive bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"hyperdrive": schema.SingleNestedAttribute{
										Description: "Hyperdrive binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
							"kv_namespaces": schema.SingleNestedAttribute{
								Description: "KV namespaces used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"kv_binding": schema.SingleNestedAttribute{
										Description: "KV binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"namespace_id": schema.StringAttribute{
												Description: "ID of the KV namespace.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"mtls_certificates": schema.SingleNestedAttribute{
								Description: "mTLS bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"mtls": schema.SingleNestedAttribute{
										Description: "mTLS binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"certificate_id": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
							"placement": schema.SingleNestedAttribute{
								Description: "Placement setting used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Description: "Placement mode.",
										Computed:    true,
										Optional:    true,
									},
								},
							},
							"queue_producers": schema.SingleNestedAttribute{
								Description: "Queue Producer bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"queue_producer_binding": schema.SingleNestedAttribute{
										Description: "Queue Producer binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "Name of the Queue.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"r2_buckets": schema.SingleNestedAttribute{
								Description: "R2 buckets used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"r2_binding": schema.SingleNestedAttribute{
										Description: "R2 binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"jurisdiction": schema.StringAttribute{
												Description: "Jurisdiction of the R2 bucket.",
												Computed:    true,
												Optional:    true,
											},
											"name": schema.StringAttribute{
												Description: "Name of the R2 bucket.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"services": schema.SingleNestedAttribute{
								Description: "Services used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"service_binding": schema.SingleNestedAttribute{
										Description: "Service binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"entrypoint": schema.StringAttribute{
												Description: "The entrypoint to bind to.",
												Computed:    true,
												Optional:    true,
											},
											"environment": schema.StringAttribute{
												Description: "The Service environment.",
												Computed:    true,
												Optional:    true,
											},
											"service": schema.StringAttribute{
												Description: "The Service name.",
												Computed:    true,
												Optional:    true,
											},
										},
									},
								},
							},
							"vectorize_bindings": schema.SingleNestedAttribute{
								Description: "Vectorize bindings used for Pages Functions.",
								Computed:    true,
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"vectorize": schema.SingleNestedAttribute{
										Description: "Vectorize binding.",
										Computed:    true,
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"index_name": schema.StringAttribute{
												Computed: true,
												Optional: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"domains": schema.ListAttribute{
				Description: "A list of associated custom domains for the project.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"latest_deployment": schema.SingleNestedAttribute{
				Optional: true,
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
									Optional:    true,
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
			"name": schema.StringAttribute{
				Description: "Name of the project.",
				Optional:    true,
			},
			"production_branch": schema.StringAttribute{
				Description: "Production branch of the project. Used to identify production deployments.",
				Optional:    true,
			},
			"source": schema.StringAttribute{
				Computed: true,
			},
			"subdomain": schema.StringAttribute{
				Description: "The Cloudflare subdomain associated with the project.",
				Optional:    true,
			},
			"aliases": schema.ListAttribute{
				Description: "A list of alias URLs pointing to this deployment.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"deployment_trigger": schema.SingleNestedAttribute{
				Description: "Info about what caused the deployment.",
				Optional:    true,
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
			"env_vars": schema.StringAttribute{
				Description: "A dict of env variables to build this deploy.",
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
			"latest_stage": schema.StringAttribute{
				Optional: true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the deployment was last modified.",
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
			"stages": schema.ListNestedAttribute{
				Description: "List of past stages.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ended_on": schema.StringAttribute{
							Description: "When the stage ended.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The current build stage.",
							Computed:    true,
							Optional:    true,
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
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
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

func (r *PagesProjectDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *PagesProjectDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
