// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

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
			"message": schema.StringAttribute{
				Description: "Rollback message to be associated with this deployment. Only parsed when query param `\"rollback_to\"` is present.",
				Optional:    true,
			},
			"any_part_name": schema.ListAttribute{
				Description: "A module comprising a Worker script, often a javascript file. Multiple modules may be provided as separate named parts, but at least one module must be present and referenced in the metadata as `main_module` or `body_part` by part name. Source maps may also be included using the `application/source-map` content type.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"metadata": schema.SingleNestedAttribute{
				Description: "JSON encoded metadata about the uploaded parts and Worker configuration.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkersScriptMetadataModel](ctx),
				Attributes: map[string]schema.Attribute{
					"bindings": schema.ListNestedAttribute{
						Description: "List of bindings available to the worker.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataBindingsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "Name of the binding variable.",
									Optional:    true,
								},
								"type": schema.StringAttribute{
									Description: "Type of binding. You can find more about bindings on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.",
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
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkersScriptMetadataMigrationsModel](ctx),
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
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataMigrationsRenamedClassesModel](ctx),
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
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataMigrationsTransferredClassesModel](ctx),
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
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataMigrationsStepsModel](ctx),
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
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataMigrationsStepsRenamedClassesModel](ctx),
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
											Computed:    true,
											Optional:    true,
											CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataMigrationsStepsTransferredClassesModel](ctx),
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
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[WorkersScriptMetadataPlacementModel](ctx),
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
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkersScriptMetadataTailConsumersModel](ctx),
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
					"version_tags": schema.MapAttribute{
						Description: "Key-value pairs to use as tags for this version of this Worker",
						Optional:    true,
						ElementType: types.StringType,
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
				Description: "Specifies the placement mode for the Worker (e.g. 'smart').",
				Computed:    true,
			},
			"startup_time_ms": schema.Int64Attribute{
				Computed: true,
			},
			"usage_model": schema.StringAttribute{
				Description: "Specifies the usage model for the Worker (e.g. 'bundled' or 'unbound').",
				Computed:    true,
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
