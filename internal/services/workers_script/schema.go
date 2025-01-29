// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WorkersScriptResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Name of the script, used in URLs and route configuration.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"script_name": schema.StringAttribute{
				Description:   "Name of the script, used in URLs and route configuration.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "JSON encoded metadata about the uploaded parts and Worker configuration.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"assets": schema.SingleNestedAttribute{
						Description: "Configuration for assets within a Worker",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Description: "Configuration for assets within a Worker.",
								Optional:    true,
								Attributes: map[string]schema.Attribute{
									"html_handling": schema.StringAttribute{
										Description: "Determines the redirects and rewrites of requests for HTML content.",
										Optional:    true,
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
										Description: "Determines the response when a request does not match a static asset, and there is no Worker script.",
										Optional:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive(
												"none",
												"404-page",
												"single-page-application",
											),
										},
									},
									"run_worker_first": schema.BoolAttribute{
										Description: "When true, requests will always invoke the Worker script. Otherwise, attempt to serve an asset matching the request, falling back to the Worker script.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
									"serve_directly": schema.BoolAttribute{
										Description: "When true and the incoming request matches an asset, that will be served instead of invoking the Worker script. When false, requests will always invoke the Worker script.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(true),
									},
								},
							},
							"jwt": schema.StringAttribute{
								Description: "Token provided upon successful upload of all files from a registered manifest.",
								Optional:    true,
							},
						},
					},
					"bindings": schema.ListNestedAttribute{
						Description: "List of bindings attached to a Worker. You can find more about bindings on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "A JavaScript variable name for the binding.",
									Required:    true,
								},
								"type": schema.StringAttribute{
									Description: "The kind of resource that the binding provides.",
									Required:    true,
								},
								"dataset": schema.StringAttribute{
									Description: "The dataset name to bind to.",
									Optional:    true,
								},
								"id": schema.StringAttribute{
									Description: "Identifier of the D1 database to bind to.",
									Optional:    true,
								},
								"namespace": schema.StringAttribute{
									Description: "Namespace to bind to.",
									Optional:    true,
								},
								"outbound": schema.SingleNestedAttribute{
									Description: "Outbound worker.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"params": schema.ListAttribute{
											Description: "Pass information from the Dispatch Worker to the Outbound Worker through the parameters.",
											Optional:    true,
											ElementType: types.StringType,
										},
										"worker": schema.SingleNestedAttribute{
											Description: "Outbound worker.",
											Optional:    true,
											Attributes: map[string]schema.Attribute{
												"environment": schema.StringAttribute{
													Description: "Environment of the outbound worker.",
													Optional:    true,
												},
												"service": schema.StringAttribute{
													Description: "Name of the outbound worker.",
													Optional:    true,
												},
											},
										},
									},
								},
								"class_name": schema.StringAttribute{
									Description: "The exported class name of the Durable Object.",
									Optional:    true,
								},
								"environment": schema.StringAttribute{
									Description: "The environment of the script_name to bind to.",
									Optional:    true,
								},
								"namespace_id": schema.StringAttribute{
									Description: "Namespace identifier tag.",
									Optional:    true,
								},
								"script_name": schema.StringAttribute{
									Description: "The script where the Durable Object is defined, if it is external to this Worker.",
									Optional:    true,
								},
								"json": schema.StringAttribute{
									Description: "JSON data to use.",
									Optional:    true,
								},
								"certificate_id": schema.StringAttribute{
									Description: "Identifier of the certificate to bind to.",
									Optional:    true,
								},
								"text": schema.StringAttribute{
									Description: "The text value to use.",
									Optional:    true,
								},
								"queue_name": schema.StringAttribute{
									Description: "Name of the Queue to bind to.",
									Optional:    true,
								},
								"bucket_name": schema.StringAttribute{
									Description: "R2 bucket to bind to.",
									Optional:    true,
								},
								"service": schema.StringAttribute{
									Description: "Name of Worker to bind to.",
									Optional:    true,
								},
								"index_name": schema.StringAttribute{
									Description: "Name of the Vectorize index to bind to.",
									Optional:    true,
								},
							},
						},
					},
					"body_part": schema.StringAttribute{
						Description: "Name of the part in the multipart request that contains the script (e.g. the file adding a listener to the `fetch` event). Indicates a `service worker syntax` Worker.",
						Optional:    true,
					},
					"compatibility_date": schema.StringAttribute{
						Description: "Date indicating targeted support in the Workers runtime. Backwards incompatible fixes to the runtime following this date will not affect this Worker.",
						Optional:    true,
					},
					"compatibility_flags": schema.ListAttribute{
						Description: "Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"keep_assets": schema.BoolAttribute{
						Description: "Retain assets which exist for a previously uploaded Worker version; used in lieu of providing a completion token.",
						Optional:    true,
					},
					"keep_bindings": schema.ListAttribute{
						Description: "List of binding types to keep from previous_upload.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"logpush": schema.BoolAttribute{
						Description: "Whether Logpush is turned on for the Worker.",
						Optional:    true,
					},
					"main_module": schema.StringAttribute{
						Description: "Name of the part in the multipart request that contains the main module (e.g. the file exporting a `fetch` handler). Indicates a `module syntax` Worker.",
						Optional:    true,
					},
					"migrations": schema.SingleNestedAttribute{
						Description: "Migrations to apply for Durable Objects associated with this Worker.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"deleted_classes": schema.ListAttribute{
								Description: "A list of classes to delete Durable Object namespaces from.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"new_classes": schema.ListAttribute{
								Description: "A list of classes to create Durable Object namespaces from.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"new_sqlite_classes": schema.ListAttribute{
								Description: "A list of classes to create Durable Object namespaces with SQLite from.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"new_tag": schema.StringAttribute{
								Description: "Tag to set as the latest migration tag.",
								Optional:    true,
							},
							"old_tag": schema.StringAttribute{
								Description: "Tag used to verify against the latest migration tag for this Worker. If they don't match, the upload is rejected.",
								Optional:    true,
							},
							"renamed_classes": schema.ListNestedAttribute{
								Description: "A list of classes with Durable Object namespaces that were renamed.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from": schema.StringAttribute{
											Optional: true,
										},
										"to": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"transferred_classes": schema.ListNestedAttribute{
								Description: "A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"from": schema.StringAttribute{
											Optional: true,
										},
										"from_script": schema.StringAttribute{
											Optional: true,
										},
										"to": schema.StringAttribute{
											Optional: true,
										},
									},
								},
							},
							"steps": schema.ListNestedAttribute{
								Description: "Migrations to apply in order.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"deleted_classes": schema.ListAttribute{
											Description: "A list of classes to delete Durable Object namespaces from.",
											Optional:    true,
											ElementType: types.StringType,
										},
										"new_classes": schema.ListAttribute{
											Description: "A list of classes to create Durable Object namespaces from.",
											Optional:    true,
											ElementType: types.StringType,
										},
										"new_sqlite_classes": schema.ListAttribute{
											Description: "A list of classes to create Durable Object namespaces with SQLite from.",
											Optional:    true,
											ElementType: types.StringType,
										},
										"renamed_classes": schema.ListNestedAttribute{
											Description: "A list of classes with Durable Object namespaces that were renamed.",
											Optional:    true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"from": schema.StringAttribute{
														Optional: true,
													},
													"to": schema.StringAttribute{
														Optional: true,
													},
												},
											},
										},
										"transferred_classes": schema.ListNestedAttribute{
											Description: "A list of transfers for Durable Object namespaces from a different Worker and class to a class defined in this Worker.",
											Optional:    true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"from": schema.StringAttribute{
														Optional: true,
													},
													"from_script": schema.StringAttribute{
														Optional: true,
													},
													"to": schema.StringAttribute{
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
					"observability": schema.SingleNestedAttribute{
						Description: "Observability settings for the Worker.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether observability is enabled for the Worker.",
								Required:    true,
							},
							"head_sampling_rate": schema.Float64Attribute{
								Description: "The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.",
								Optional:    true,
							},
						},
					},
					"placement": schema.SingleNestedAttribute{
						Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("smart"),
								},
							},
							"status": schema.StringAttribute{
								Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
								Computed:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"SUCCESS",
										"UNSUPPORTED_APPLICATION",
										"INSUFFICIENT_INVOCATIONS",
									),
								},
							},
						},
					},
					"tags": schema.ListAttribute{
						Description: "List of strings to use as tags for this Worker.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"tail_consumers": schema.ListNestedAttribute{
						Description: "List of Workers that will consume logs from the attached Worker.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"service": schema.StringAttribute{
									Description: "Name of Worker that is to be the consumer.",
									Required:    true,
								},
								"environment": schema.StringAttribute{
									Description: "Optional environment if the Worker utilizes one.",
									Optional:    true,
								},
								"namespace": schema.StringAttribute{
									Description: "Optional dispatch namespace the script belongs to.",
									Optional:    true,
								},
							},
						},
					},
					"usage_model": schema.StringAttribute{
						Description: "Usage model for the Worker invocations.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("standard"),
						},
					},
				},
			},
			"created_on": schema.StringAttribute{
				Description: "When the script was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"etag": schema.StringAttribute{
				Description: "Hashed script content, can be used in a If-None-Match header when updating.",
				Computed:    true,
			},
			"has_assets": schema.BoolAttribute{
				Description: "Whether a Worker contains assets.",
				Computed:    true,
			},
			"has_modules": schema.BoolAttribute{
				Description: "Whether a Worker contains modules.",
				Computed:    true,
			},
			"logpush": schema.BoolAttribute{
				Description: "Whether Logpush is turned on for the Worker.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Description: "When the script was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"placement_mode": schema.StringAttribute{
				Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("smart"),
				},
			},
			"placement_status": schema.StringAttribute{
				Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"SUCCESS",
						"UNSUPPORTED_APPLICATION",
						"INSUFFICIENT_INVOCATIONS",
					),
				},
			},
			"startup_time_ms": schema.Int64Attribute{
				Computed: true,
			},
			"usage_model": schema.StringAttribute{
				Description: "Usage model for the Worker invocations.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("standard"),
				},
			},
			"placement": schema.SingleNestedAttribute{
				Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WorkersScriptPlacementModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("smart"),
						},
					},
					"status": schema.StringAttribute{
						Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive(
								"SUCCESS",
								"UNSUPPORTED_APPLICATION",
								"INSUFFICIENT_INVOCATIONS",
							),
						},
					},
				},
			},
			"tail_consumers": schema.ListNestedAttribute{
				Description: "List of Workers that will consume logs from the attached Worker.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkersScriptTailConsumersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"service": schema.StringAttribute{
							Description: "Name of Worker that is to be the consumer.",
							Computed:    true,
						},
						"environment": schema.StringAttribute{
							Description: "Optional environment if the Worker utilizes one.",
							Computed:    true,
						},
						"namespace": schema.StringAttribute{
							Description: "Optional dispatch namespace the script belongs to.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *WorkersScriptResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkersScriptResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
