// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_script

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r WorkerScriptResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The id of the script in the Workers system. Usually the script name.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"script_name": schema.StringAttribute{
				Description: "Name of the script, used in URLs and route configuration.",
				Required:    true,
			},
			"any_part_name": schema.StringAttribute{
				Description: "A module comprising a Worker script, often a javascript file. Multiple modules may be provided as separate named parts, but at least one module must be present and referenced in the metadata as `main_module` or `body_part` by part name. Source maps may also be included using the `application/source-map` content type.",
				Optional:    true,
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "JSON encoded metadata about the uploaded parts and Worker configuration.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"bindings": schema.ListAttribute{
						Description: "List of bindings available to the worker.",
						Optional:    true,
						ElementType: types.StringType,
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
							"deleted_classes": schema.StringAttribute{
								Description: "A list of classes to delete Durable Object namespaces from.",
								Optional:    true,
							},
							"new_classes": schema.StringAttribute{
								Description: "A list of classes to create Durable Object namespaces from.",
								Optional:    true,
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
					"placement": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"mode": schema.StringAttribute{
								Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Only `\"smart\"` is currently supported",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("smart"),
								},
							},
						},
					},
					"tags": schema.ListAttribute{
						Description: "List of strings to use as tags for this Worker",
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
						Description: "Usage model to apply to invocations.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("bundled", "unbound"),
						},
					},
					"version_tags": schema.StringAttribute{
						Description: "Key-value pairs to use as tags for this version of this Worker",
						Optional:    true,
					},
				},
			},
			"message": schema.StringAttribute{
				Description: "Rollback message to be associated with this deployment. Only parsed when query param `\"rollback_to\"` is present.",
				Optional:    true,
			},
		},
	}
}
