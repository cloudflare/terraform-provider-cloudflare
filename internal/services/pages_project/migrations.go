// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages_project

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r PagesProjectResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"project_name": schema.StringAttribute{
						Description:   "Name of the project.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"build_config": schema.SingleNestedAttribute{
						Description: "Configs for the project build process.",
						Optional:    true,
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
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
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
								ElementType: jsontypes.NewNormalizedNull().Type(ctx),
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
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
					},
					"deployment_configs": schema.SingleNestedAttribute{
						Description: "Configs for deployments in a project.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"preview": schema.SingleNestedAttribute{
								Description: "Configs for preview deploys.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"ai_bindings": schema.SingleNestedAttribute{
										Description: "Constellation bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"ai_binding": schema.SingleNestedAttribute{
												Description: "AI binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"project_id": schema.StringAttribute{
														Optional: true,
													},
												},
											},
										},
									},
									"analytics_engine_datasets": schema.SingleNestedAttribute{
										Description: "Analytics Engine bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"analytics_engine_binding": schema.SingleNestedAttribute{
												Description: "Analytics Engine binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"dataset": schema.StringAttribute{
														Description: "Name of the dataset.",
														Optional:    true,
													},
												},
											},
										},
									},
									"browsers": schema.SingleNestedAttribute{
										Description: "Browser bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"browser": schema.StringAttribute{
												Description: "Browser binding.",
												Optional:    true,
											},
										},
									},
									"compatibility_date": schema.StringAttribute{
										Description: "Compatibility date used for Pages Functions.",
										Optional:    true,
									},
									"compatibility_flags": schema.ListAttribute{
										Description: "Compatibility flags used for Pages Functions.",
										Optional:    true,
										ElementType: jsontypes.NewNormalizedNull().Type(ctx),
									},
									"d1_databases": schema.SingleNestedAttribute{
										Description: "D1 databases used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"d1_binding": schema.SingleNestedAttribute{
												Description: "D1 binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "UUID of the D1 database.",
														Optional:    true,
													},
												},
											},
										},
									},
									"durable_object_namespaces": schema.SingleNestedAttribute{
										Description: "Durabble Object namespaces used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"do_binding": schema.SingleNestedAttribute{
												Description: "Durabble Object binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"namespace_id": schema.StringAttribute{
														Description: "ID of the Durabble Object namespace.",
														Optional:    true,
													},
												},
											},
										},
									},
									"env_vars": schema.SingleNestedAttribute{
										Description: "Environment variables for build configs.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"environment_variable": schema.SingleNestedAttribute{
												Description: "Environment variable.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description: "The type of environment variable (plain text or secret)",
														Optional:    true,
														Validators: []validator.String{
															stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
														},
													},
													"value": schema.StringAttribute{
														Description: "Environment variable value.",
														Optional:    true,
													},
												},
											},
										},
									},
									"hyperdrive_bindings": schema.SingleNestedAttribute{
										Description: "Hyperdrive bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"hyperdrive": schema.SingleNestedAttribute{
												Description: "Hyperdrive binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Optional: true,
													},
												},
											},
										},
									},
									"kv_namespaces": schema.SingleNestedAttribute{
										Description: "KV namespaces used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"kv_binding": schema.SingleNestedAttribute{
												Description: "KV binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"namespace_id": schema.StringAttribute{
														Description: "ID of the KV namespace.",
														Optional:    true,
													},
												},
											},
										},
									},
									"mtls_certificates": schema.SingleNestedAttribute{
										Description: "mTLS bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"mtls": schema.SingleNestedAttribute{
												Description: "mTLS binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"certificate_id": schema.StringAttribute{
														Optional: true,
													},
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
									"queue_producers": schema.SingleNestedAttribute{
										Description: "Queue Producer bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"queue_producer_binding": schema.SingleNestedAttribute{
												Description: "Queue Producer binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description: "Name of the Queue.",
														Optional:    true,
													},
												},
											},
										},
									},
									"r2_buckets": schema.SingleNestedAttribute{
										Description: "R2 buckets used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"r2_binding": schema.SingleNestedAttribute{
												Description: "R2 binding.",
												Optional:    true,
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
									},
									"services": schema.SingleNestedAttribute{
										Description: "Services used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"service_binding": schema.SingleNestedAttribute{
												Description: "Service binding.",
												Optional:    true,
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
									},
									"vectorize_bindings": schema.SingleNestedAttribute{
										Description: "Vectorize bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"vectorize": schema.SingleNestedAttribute{
												Description: "Vectorize binding.",
												Optional:    true,
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
							"production": schema.SingleNestedAttribute{
								Description: "Configs for production deploys.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"ai_bindings": schema.SingleNestedAttribute{
										Description: "Constellation bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"ai_binding": schema.SingleNestedAttribute{
												Description: "AI binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"project_id": schema.StringAttribute{
														Optional: true,
													},
												},
											},
										},
									},
									"analytics_engine_datasets": schema.SingleNestedAttribute{
										Description: "Analytics Engine bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"analytics_engine_binding": schema.SingleNestedAttribute{
												Description: "Analytics Engine binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"dataset": schema.StringAttribute{
														Description: "Name of the dataset.",
														Optional:    true,
													},
												},
											},
										},
									},
									"browsers": schema.SingleNestedAttribute{
										Description: "Browser bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"browser": schema.StringAttribute{
												Description: "Browser binding.",
												Optional:    true,
											},
										},
									},
									"compatibility_date": schema.StringAttribute{
										Description: "Compatibility date used for Pages Functions.",
										Optional:    true,
									},
									"compatibility_flags": schema.ListAttribute{
										Description: "Compatibility flags used for Pages Functions.",
										Optional:    true,
										ElementType: jsontypes.NewNormalizedNull().Type(ctx),
									},
									"d1_databases": schema.SingleNestedAttribute{
										Description: "D1 databases used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"d1_binding": schema.SingleNestedAttribute{
												Description: "D1 binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Description: "UUID of the D1 database.",
														Optional:    true,
													},
												},
											},
										},
									},
									"durable_object_namespaces": schema.SingleNestedAttribute{
										Description: "Durabble Object namespaces used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"do_binding": schema.SingleNestedAttribute{
												Description: "Durabble Object binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"namespace_id": schema.StringAttribute{
														Description: "ID of the Durabble Object namespace.",
														Optional:    true,
													},
												},
											},
										},
									},
									"env_vars": schema.SingleNestedAttribute{
										Description: "Environment variables for build configs.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"environment_variable": schema.SingleNestedAttribute{
												Description: "Environment variable.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description: "The type of environment variable (plain text or secret)",
														Optional:    true,
														Validators: []validator.String{
															stringvalidator.OneOfCaseInsensitive("plain_text", "secret_text"),
														},
													},
													"value": schema.StringAttribute{
														Description: "Environment variable value.",
														Optional:    true,
													},
												},
											},
										},
									},
									"hyperdrive_bindings": schema.SingleNestedAttribute{
										Description: "Hyperdrive bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"hyperdrive": schema.SingleNestedAttribute{
												Description: "Hyperdrive binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"id": schema.StringAttribute{
														Optional: true,
													},
												},
											},
										},
									},
									"kv_namespaces": schema.SingleNestedAttribute{
										Description: "KV namespaces used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"kv_binding": schema.SingleNestedAttribute{
												Description: "KV binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"namespace_id": schema.StringAttribute{
														Description: "ID of the KV namespace.",
														Optional:    true,
													},
												},
											},
										},
									},
									"mtls_certificates": schema.SingleNestedAttribute{
										Description: "mTLS bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"mtls": schema.SingleNestedAttribute{
												Description: "mTLS binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"certificate_id": schema.StringAttribute{
														Optional: true,
													},
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
									"queue_producers": schema.SingleNestedAttribute{
										Description: "Queue Producer bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"queue_producer_binding": schema.SingleNestedAttribute{
												Description: "Queue Producer binding.",
												Optional:    true,
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description: "Name of the Queue.",
														Optional:    true,
													},
												},
											},
										},
									},
									"r2_buckets": schema.SingleNestedAttribute{
										Description: "R2 buckets used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"r2_binding": schema.SingleNestedAttribute{
												Description: "R2 binding.",
												Optional:    true,
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
									},
									"services": schema.SingleNestedAttribute{
										Description: "Services used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"service_binding": schema.SingleNestedAttribute{
												Description: "Service binding.",
												Optional:    true,
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
									},
									"vectorize_bindings": schema.SingleNestedAttribute{
										Description: "Vectorize bindings used for Pages Functions.",
										Optional:    true,
										Attributes: map[string]schema.Attribute{
											"vectorize": schema.SingleNestedAttribute{
												Description: "Vectorize binding.",
												Optional:    true,
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
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
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
								ElementType: jsontypes.NewNormalizedNull().Type(ctx),
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
						PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
					},
					"name": schema.StringAttribute{
						Description:   "Name of the project.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"production_branch": schema.StringAttribute{
						Description:   "Production branch of the project. Used to identify production deployments.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"id": schema.StringAttribute{
						Description: "Id of the project.",
						Computed:    true,
					},
					"created_on": schema.StringAttribute{
						Description: "When the project was created.",
						Computed:    true,
					},
					"domains": schema.ListAttribute{
						Description: "A list of associated custom domains for the project.",
						Computed:    true,
						ElementType: jsontypes.NewNormalizedNull().Type(ctx),
					},
					"source": schema.StringAttribute{
						Computed: true,
					},
					"subdomain": schema.StringAttribute{
						Description: "The Cloudflare subdomain associated with the project.",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state PagesProjectModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}
