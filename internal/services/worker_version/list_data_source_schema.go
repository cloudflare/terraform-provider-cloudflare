// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkerVersionsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"worker_id": schema.StringAttribute{
				Description: "Identifier for the Worker, which can be ID or name.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkerVersionsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Version identifier.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the version was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"number": schema.Int64Attribute{
							Description: "The integer version number, starting from one.",
							Computed:    true,
						},
						"annotations": schema.SingleNestedAttribute{
							Description: "Metadata about the version.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkerVersionsAnnotationsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"workers_message": schema.StringAttribute{
									Description: "Human-readable message about the version.",
									Computed:    true,
								},
								"workers_tag": schema.StringAttribute{
									Description: "User-provided identifier for the version.",
									Computed:    true,
								},
								"workers_triggered_by": schema.StringAttribute{
									Description: "Operation that triggered the creation of the version.",
									Computed:    true,
								},
							},
						},
						"assets": schema.SingleNestedAttribute{
							Description: "Configuration for assets within a Worker.\n\n[`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers) and\n[`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/) files should be\nincluded as modules named `_headers` and `_redirects` with content type `text/plain`.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkerVersionsAssetsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"config": schema.SingleNestedAttribute{
									Description: "Configuration for assets within a Worker.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[WorkerVersionsAssetsConfigDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"html_handling": schema.StringAttribute{
											Description: "Determines the redirects and rewrites of requests for HTML content.\nAvailable values: \"auto-trailing-slash\", \"force-trailing-slash\", \"drop-trailing-slash\", \"none\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"auto-trailing-slash",
													"force-trailing-slash",
													"drop-trailing-slash",
													"none",
												),
											},
										},
										"not_found_handling": schema.StringAttribute{
											Description: "Determines the response when a request does not match a static asset, and there is no Worker script.\nAvailable values: \"none\", \"404-page\", \"single-page-application\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"none",
													"404-page",
													"single-page-application",
												),
											},
										},
										"run_worker_first": schema.ListAttribute{
											Description: "Contains a list path rules to control routing to either the Worker or assets. Glob (*) and negative (!) rules are supported. Rules must start with either '/' or '!/'. At least one non-negative rule must be provided, and negative rules have higher precedence than non-negative rules.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
									},
								},
								"jwt": schema.StringAttribute{
									Description: "Token provided upon successful upload of all files from a registered manifest.",
									Computed:    true,
									Sensitive:   true,
								},
							},
						},
						"bindings": schema.ListNestedAttribute{
							Description: "List of bindings attached to a Worker. You can find more about bindings on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[WorkerVersionsBindingsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "A JavaScript variable name for the binding.",
										Computed:    true,
									},
									"type": schema.StringAttribute{
										Description: "The kind of resource that the binding provides.\nAvailable values: \"ai\", \"analytics_engine\", \"assets\", \"browser\", \"d1\", \"data_blob\", \"dispatch_namespace\", \"durable_object_namespace\", \"hyperdrive\", \"inherit\", \"images\", \"json\", \"kv_namespace\", \"mtls_certificate\", \"plain_text\", \"pipelines\", \"queue\", \"r2_bucket\", \"secret_text\", \"send_email\", \"service\", \"text_blob\", \"vectorize\", \"version_metadata\", \"secrets_store_secret\", \"secret_key\", \"workflow\", \"wasm_module\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"ai",
												"analytics_engine",
												"assets",
												"browser",
												"d1",
												"data_blob",
												"dispatch_namespace",
												"durable_object_namespace",
												"hyperdrive",
												"inherit",
												"images",
												"json",
												"kv_namespace",
												"mtls_certificate",
												"plain_text",
												"pipelines",
												"queue",
												"r2_bucket",
												"secret_text",
												"send_email",
												"service",
												"text_blob",
												"vectorize",
												"version_metadata",
												"secrets_store_secret",
												"secret_key",
												"workflow",
												"wasm_module",
											),
										},
									},
									"dataset": schema.StringAttribute{
										Description: "The name of the dataset to bind to.",
										Computed:    true,
									},
									"id": schema.StringAttribute{
										Description: "Identifier of the D1 database to bind to.",
										Computed:    true,
									},
									"part": schema.StringAttribute{
										Description: "The name of the file containing the data content. Only accepted for `service worker syntax` Workers.",
										Computed:    true,
									},
									"namespace": schema.StringAttribute{
										Description: "The name of the dispatch namespace.",
										Computed:    true,
									},
									"outbound": schema.SingleNestedAttribute{
										Description: "Outbound worker.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[WorkerVersionsBindingsOutboundDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"params": schema.ListAttribute{
												Description: "Pass information from the Dispatch Worker to the Outbound Worker through the parameters.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"worker": schema.SingleNestedAttribute{
												Description: "Outbound worker.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectType[WorkerVersionsBindingsOutboundWorkerDataSourceModel](ctx),
												Attributes: map[string]schema.Attribute{
													"environment": schema.StringAttribute{
														Description: "Environment of the outbound worker.",
														Computed:    true,
													},
													"service": schema.StringAttribute{
														Description: "Name of the outbound worker.",
														Computed:    true,
													},
												},
											},
										},
									},
									"class_name": schema.StringAttribute{
										Description: "The exported class name of the Durable Object.",
										Computed:    true,
									},
									"environment": schema.StringAttribute{
										Description: "The environment of the script_name to bind to.",
										Computed:    true,
									},
									"namespace_id": schema.StringAttribute{
										Description: "Namespace identifier tag.",
										Computed:    true,
									},
									"script_name": schema.StringAttribute{
										Description: "The script where the Durable Object is defined, if it is external to this Worker.",
										Computed:    true,
									},
									"old_name": schema.StringAttribute{
										Description: "The old name of the inherited binding. If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.",
										Computed:    true,
									},
									"version_id": schema.StringAttribute{
										Description: `Identifier for the version to inherit the binding from, which can be the version ID or the literal "latest" to inherit from the latest version. Defaults to inheriting the binding from the latest version.`,
										Computed:    true,
									},
									"json": schema.StringAttribute{
										Description: "JSON data to use.",
										Computed:    true,
									},
									"certificate_id": schema.StringAttribute{
										Description: "Identifier of the certificate to bind to.",
										Computed:    true,
									},
									"text": schema.StringAttribute{
										Description: "The text value to use.",
										Computed:    true,
										Sensitive:   true,
									},
									"pipeline": schema.StringAttribute{
										Description: "Name of the Pipeline to bind to.",
										Computed:    true,
									},
									"queue_name": schema.StringAttribute{
										Description: "Name of the Queue to bind to.",
										Computed:    true,
									},
									"bucket_name": schema.StringAttribute{
										Description: "R2 bucket to bind to.",
										Computed:    true,
									},
									"jurisdiction": schema.StringAttribute{
										Description: "The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket.\nAvailable values: \"eu\", \"fedramp\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("eu", "fedramp"),
										},
									},
									"allowed_destination_addresses": schema.ListAttribute{
										Description: "List of allowed destination addresses.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"allowed_sender_addresses": schema.ListAttribute{
										Description: "List of allowed sender addresses.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"destination_address": schema.StringAttribute{
										Description: "Destination address for the email.",
										Computed:    true,
									},
									"service": schema.StringAttribute{
										Description: "Name of Worker to bind to.",
										Computed:    true,
									},
									"index_name": schema.StringAttribute{
										Description: "Name of the Vectorize index to bind to.",
										Computed:    true,
									},
									"secret_name": schema.StringAttribute{
										Description: "Name of the secret in the store.",
										Computed:    true,
									},
									"store_id": schema.StringAttribute{
										Description: "ID of the store containing the secret.",
										Computed:    true,
									},
									"algorithm": schema.StringAttribute{
										Description: "Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).",
										Computed:    true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"format": schema.StringAttribute{
										Description: "Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).\nAvailable values: \"raw\", \"pkcs8\", \"spki\", \"jwk\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"raw",
												"pkcs8",
												"spki",
												"jwk",
											),
										},
									},
									"usages": schema.SetAttribute{
										Description: "Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).",
										Computed:    true,
										CustomType:  customfield.NewSetType[types.String](ctx),
										ElementType: types.StringType,
									},
									"key_base64": schema.StringAttribute{
										Description: "Base64-encoded key data. Required if `format` is \"raw\", \"pkcs8\", or \"spki\".",
										Computed:    true,
										Sensitive:   true,
									},
									"key_jwk": schema.StringAttribute{
										Description: "Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is \"jwk\".",
										Computed:    true,
										Sensitive:   true,
										CustomType:  jsontypes.NormalizedType{},
									},
									"workflow_name": schema.StringAttribute{
										Description: "Name of the Workflow to bind to.",
										Computed:    true,
									},
								},
							},
						},
						"compatibility_date": schema.StringAttribute{
							Description: "Date indicating targeted support in the Workers runtime. Backwards incompatible fixes to the runtime following this date will not affect this Worker.",
							Computed:    true,
						},
						"compatibility_flags": schema.SetAttribute{
							Description: "Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"limits": schema.SingleNestedAttribute{
							Description: "Resource limits enforced at runtime.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkerVersionsLimitsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"cpu_ms": schema.Int64Attribute{
									Description: "CPU time limit in milliseconds.",
									Computed:    true,
								},
							},
						},
						"main_module": schema.StringAttribute{
							Description: "The name of the main module in the `modules` array (e.g. the name of the module that exports a `fetch` handler).",
							Computed:    true,
						},
						"migrations": schema.SingleNestedAttribute{
							Description: "Migrations for Durable Objects associated with the version. Migrations are applied when the version is deployed.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkerVersionsMigrationsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"deleted_classes": schema.ListAttribute{
									Description: "A list of classes to delete Durable Object namespaces from.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"new_classes": schema.ListAttribute{
									Description: "A list of classes to create Durable Object namespaces from.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"new_sqlite_classes": schema.ListAttribute{
									Description: "A list of classes to create Durable Object namespaces with SQLite from.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"new_tag": schema.StringAttribute{
									Description: "Tag to set as the latest migration tag.",
									Computed:    true,
								},
								"old_tag": schema.StringAttribute{
									Description: "Tag used to verify against the latest migration tag for this Worker. If they don't match, the upload is rejected.",
									Computed:    true,
								},
								"renamed_classes": schema.ListNestedAttribute{
									Description: "A list of classes with Durable Object namespaces that were renamed.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[WorkerVersionsMigrationsRenamedClassesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"from": schema.StringAttribute{
												Computed: true,
											},
											"to": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"transferred_classes": schema.ListNestedAttribute{
									Description: "A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[WorkerVersionsMigrationsTransferredClassesDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"from": schema.StringAttribute{
												Computed: true,
											},
											"from_script": schema.StringAttribute{
												Computed: true,
											},
											"to": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"steps": schema.ListNestedAttribute{
									Description: "Migrations to apply in order.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[WorkerVersionsMigrationsStepsDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"deleted_classes": schema.ListAttribute{
												Description: "A list of classes to delete Durable Object namespaces from.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"new_classes": schema.ListAttribute{
												Description: "A list of classes to create Durable Object namespaces from.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"new_sqlite_classes": schema.ListAttribute{
												Description: "A list of classes to create Durable Object namespaces with SQLite from.",
												Computed:    true,
												CustomType:  customfield.NewListType[types.String](ctx),
												ElementType: types.StringType,
											},
											"renamed_classes": schema.ListNestedAttribute{
												Description: "A list of classes with Durable Object namespaces that were renamed.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectListType[WorkerVersionsMigrationsStepsRenamedClassesDataSourceModel](ctx),
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"from": schema.StringAttribute{
															Computed: true,
														},
														"to": schema.StringAttribute{
															Computed: true,
														},
													},
												},
											},
											"transferred_classes": schema.ListNestedAttribute{
												Description: "A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.",
												Computed:    true,
												CustomType:  customfield.NewNestedObjectListType[WorkerVersionsMigrationsStepsTransferredClassesDataSourceModel](ctx),
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"from": schema.StringAttribute{
															Computed: true,
														},
														"from_script": schema.StringAttribute{
															Computed: true,
														},
														"to": schema.StringAttribute{
															Computed: true,
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"modules": schema.SetNestedAttribute{
							Description: "Code, sourcemaps, and other content used at runtime.\n\nThis includes [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers) and\n[`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/) files used to configure \n[Static Assets](https://developers.cloudflare.com/workers/static-assets/). `_headers` and `_redirects` files should be \nincluded as modules named `_headers` and `_redirects` with content type `text/plain`.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectSetType[WorkerVersionsModulesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"content_base64": schema.StringAttribute{
										Description: "The base64-encoded module content.",
										Computed:    true,
									},
									"content_type": schema.StringAttribute{
										Description: "The content type of the module.",
										Computed:    true,
									},
									"name": schema.StringAttribute{
										Description: "The name of the module.",
										Computed:    true,
									},
								},
							},
						},
						"placement": schema.SingleNestedAttribute{
							Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Specify mode='smart' for Smart Placement, or one of region/hostname/host.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkerVersionsPlacementDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"mode": schema.StringAttribute{
									Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"smart\", \"targeted\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("smart", "targeted"),
									},
								},
								"region": schema.StringAttribute{
									Description: "Cloud region for targeted placement in format 'provider:region'.",
									Computed:    true,
								},
								"hostname": schema.StringAttribute{
									Description: "HTTP hostname for targeted placement.",
									Computed:    true,
								},
								"host": schema.StringAttribute{
									Description: "TCP host and port for targeted placement.",
									Computed:    true,
								},
								"target": schema.ListNestedAttribute{
									Description: "Array of placement targets (currently limited to single target).",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[WorkerVersionsPlacementTargetDataSourceModel](ctx),
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"region": schema.StringAttribute{
												Description: "Cloud region in format 'provider:region'.",
												Computed:    true,
											},
											"hostname": schema.StringAttribute{
												Description: "HTTP hostname for targeted placement.",
												Computed:    true,
											},
											"host": schema.StringAttribute{
												Description: "TCP host:port for targeted placement.",
												Computed:    true,
											},
										},
									},
								},
							},
						},
						"source": schema.StringAttribute{
							Description: "The client used to create the version.",
							Computed:    true,
						},
						"startup_time_ms": schema.Int64Attribute{
							Description: "Time in milliseconds spent on [Worker startup](https://developers.cloudflare.com/workers/platform/limits/#worker-startup-time).",
							Computed:    true,
						},
						"usage_model": schema.StringAttribute{
							Description:        "Usage model for the version.\nAvailable values: \"standard\", \"bundled\", \"unbound\".",
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"standard",
									"bundled",
									"unbound",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (d *WorkerVersionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkerVersionsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
