// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WorkerVersionResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Version identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"worker_id": schema.StringAttribute{
				Description:   "Identifier for the Worker, which can be ID or name.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"compatibility_date": schema.StringAttribute{
				Description:   "Date indicating targeted support in the Workers runtime. Backwards incompatible fixes to the runtime following this date will not affect this Worker.",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"main_module": schema.StringAttribute{
				Description:   "The name of the main module in the `modules` array (e.g. the name of the module that exports a `fetch` handler).",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"migrations": schema.SingleNestedAttribute{
				Description: "Migrations for Durable Objects associated with the version. Migrations are applied when the version is deployed.",
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
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"modules": schema.SetNestedAttribute{
				Description: "Code, sourcemaps, and other content used at runtime.\n\nThis includes [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers) and\n[`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/) files used to configure \n[Static Assets](https://developers.cloudflare.com/workers/static-assets/). `_headers` and `_redirects` files should be \nincluded as modules named `_headers` and `_redirects` with content type `text/plain`.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"content_base64": schema.StringAttribute{
							Description: "The base64-encoded module content.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("content_file")),
								stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("content_file")),
							},
						},
						"content_file": schema.StringAttribute{
							Description: "The file path of the module content.",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.ConflictsWith(path.MatchRelative().AtParent().AtName("content_base64")),
								stringvalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("content_base64")),
							},
						},
						"content_type": schema.StringAttribute{
							Description: "The content type of the module.",
							Required:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the module.",
							Required:    true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
						},
						"content_sha256": schema.StringAttribute{
							Description: "The SHA-256 hash of the module content.",
							Computed:    true,
							PlanModifiers: []planmodifier.String{
								ComputeSHA256HashOfContent(),
								stringplanmodifier.RequiresReplace(),
							},
						},
					},
				},
			},
			"placement": schema.SingleNestedAttribute{
				Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Specify mode='smart' for Smart Placement, or one of region/hostname/host.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"smart\", \"targeted\".",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("smart", "targeted"),
						},
					},
					"region": schema.StringAttribute{
						Description: "Cloud region for targeted placement in format 'provider:region'.",
						Optional:    true,
					},
					"hostname": schema.StringAttribute{
						Description: "HTTP hostname for targeted placement.",
						Optional:    true,
					},
					"host": schema.StringAttribute{
						Description: "TCP host and port for targeted placement.",
						Optional:    true,
					},
					"target": schema.ListNestedAttribute{
						Description: "Array of placement targets (currently limited to single target).",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"region": schema.StringAttribute{
									Description: "Cloud region in format 'provider:region'.",
									Optional:    true,
								},
								"hostname": schema.StringAttribute{
									Description: "HTTP hostname for targeted placement.",
									Optional:    true,
								},
								"host": schema.StringAttribute{
									Description: "TCP host:port for targeted placement.",
									Optional:    true,
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplace()},
			},
			"usage_model": schema.StringAttribute{
				Description:        "Usage model for the version.\nAvailable values: \"standard\", \"bundled\", \"unbound\".",
				Computed:           true,
				Optional:           true,
				DeprecationMessage: "This attribute is deprecated.",
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"standard",
						"bundled",
						"unbound",
					),
				},
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplaceIfConfigured()},
				Default:       stringdefault.StaticString("standard"),
			},
			"compatibility_flags": schema.SetAttribute{
				Description:   "Flags that enable or disable certain features in the Workers runtime. Used to enable upcoming features or opt in or out of specific changes not included in a `compatibility_date`.",
				Computed:      true,
				Optional:      true,
				CustomType:    customfield.NewSetType[types.String](ctx),
				ElementType:   types.StringType,
				PlanModifiers: []planmodifier.Set{setplanmodifier.RequiresReplaceIfConfigured()},
			},
			"annotations": schema.SingleNestedAttribute{
				Description: "Metadata about the version.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerVersionAnnotationsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"workers_message": schema.StringAttribute{
						Description: "Human-readable message about the version.",
						Optional:    true,
					},
					"workers_tag": schema.StringAttribute{
						Description: "User-provided identifier for the version.",
						Optional:    true,
					},
					"workers_triggered_by": schema.StringAttribute{
						Description: "Operation that triggered the creation of the version.",
						Computed:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{RequiresReplaceIfConfiguredIgnoringComputedDiff("workers_triggered_by")},
			},
			"assets": schema.SingleNestedAttribute{
				Description: "Configuration for assets within a Worker.\n\n[`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers) and\n[`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/) files should be\nincluded as modules named `_headers` and `_redirects` with content type `text/plain`.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Description: "Configuration for assets within a Worker.",
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerVersionAssetsConfigModel](ctx),
						Attributes: map[string]schema.Attribute{
							"html_handling": schema.StringAttribute{
								Description: "Determines the redirects and rewrites of requests for HTML content.\nAvailable values: \"auto-trailing-slash\", \"force-trailing-slash\", \"drop-trailing-slash\", \"none\".",
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
								Description: "Determines the response when a request does not match a static asset, and there is no Worker script.\nAvailable values: \"none\", \"404-page\", \"single-page-application\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"none",
										"404-page",
										"single-page-application",
									),
								},
							},
							"run_worker_first": schema.DynamicAttribute{
								Description:   "When a boolean true, requests will always invoke the Worker script. Otherwise, attempt to serve an asset matching the request, falling back to the Worker script. When a list of strings, contains path rules to control routing to either the Worker or assets. Glob (*) and negative (!) rules are supported. Rules must start with either '/' or '!/'. At least one non-negative rule must be provided, and negative rules have higher precedence than non-negative rules.",
								Optional:      true,
								Validators:    []validator.Dynamic{runWorkerFirstValidator{}},
								CustomType:    customfield.NormalizedDynamicType{},
								PlanModifiers: []planmodifier.Dynamic{customfield.NormalizeDynamicPlanModifier()},
							},
						},
					},
					"jwt": schema.StringAttribute{
						Description: "Token provided upon successful upload of all files from a registered manifest.",
						Optional:    true,
						Sensitive:   true,
					},
					"directory": schema.StringAttribute{
						Description: "Path to the directory containing asset files to upload.",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.ConflictsWith(path.MatchRoot("assets").AtName("jwt")),
						},
					},
					"asset_manifest_sha256": schema.StringAttribute{
						Description: "The SHA-256 hash of the asset manifest of files to upload.",
						Computed:    true,
						PlanModifiers: []planmodifier.String{
							ComputeSHA256HashOfAssetManifest(),
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
			},
			"bindings": schema.ListNestedAttribute{
				Description: "List of bindings attached to a Worker. You can find more about bindings on our docs: https://developers.cloudflare.com/workers/configuration/multipart-upload-metadata/#bindings.",
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[WorkerVersionBindingsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "A JavaScript variable name for the binding.",
							Required:    true,
						},
						"type": schema.StringAttribute{
							Description: "The kind of resource that the binding provides.\nAvailable values: \"ai\", \"analytics_engine\", \"assets\", \"browser\", \"d1\", \"data_blob\", \"dispatch_namespace\", \"durable_object_namespace\", \"hyperdrive\", \"inherit\", \"images\", \"json\", \"kv_namespace\", \"mtls_certificate\", \"plain_text\", \"pipelines\", \"queue\", \"r2_bucket\", \"secret_text\", \"send_email\", \"service\", \"text_blob\", \"vectorize\", \"version_metadata\", \"secrets_store_secret\", \"secret_key\", \"workflow\", \"wasm_module\".",
							Required:    true,
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
							Optional:    true,
						},
						"id": schema.StringAttribute{
							Description: "Identifier of the D1 database to bind to.",
							Optional:    true,
						},
						"part": schema.StringAttribute{
							Description: "The name of the file containing the data content. Only accepted for `service worker syntax` Workers.",
							Optional:    true,
						},
						"namespace": schema.StringAttribute{
							Description: "The name of the dispatch namespace.",
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
							Computed:    true,
							Optional:    true,
							PlanModifiers: []planmodifier.String{
								UnknownOnlyIf("type", "durable_object_namespace"),
							},
						},
						"environment": schema.StringAttribute{
							Description: "The environment of the script_name to bind to.",
							Optional:    true,
						},
						"namespace_id": schema.StringAttribute{
							Description: "Namespace identifier tag.",
							Computed:    true,
							Optional:    true,
							PlanModifiers: []planmodifier.String{
								UnknownOnlyIf("type", "durable_object_namespace"),
							},
						},
						"script_name": schema.StringAttribute{
							Description: "The script where the Durable Object is defined, if it is external to this Worker.",
							Computed:    true,
							Optional:    true,
							PlanModifiers: []planmodifier.String{
								UnknownOnlyIf("type", "durable_object_namespace"),
							},
						},
						"old_name": schema.StringAttribute{
							Description: "The old name of the inherited binding. If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.",
							Optional:    true,
						},
						"version_id": schema.StringAttribute{
							Description: `Identifier for the version to inherit the binding from, which can be the version ID or the literal "latest" to inherit from the latest version. Defaults to inheriting the binding from the latest version.`,
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
							Sensitive:   true,
						},
						"pipeline": schema.StringAttribute{
							Description: "Name of the Pipeline to bind to.",
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
						"jurisdiction": schema.StringAttribute{
							Description: "The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket.\nAvailable values: \"eu\", \"fedramp\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("eu", "fedramp"),
							},
						},
						"allowed_destination_addresses": schema.ListAttribute{
							Description: "List of allowed destination addresses.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"allowed_sender_addresses": schema.ListAttribute{
							Description: "List of allowed sender addresses.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"destination_address": schema.StringAttribute{
							Description: "Destination address for the email.",
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
						"secret_name": schema.StringAttribute{
							Description: "Name of the secret in the store.",
							Optional:    true,
						},
						"store_id": schema.StringAttribute{
							Description: "ID of the store containing the secret.",
							Optional:    true,
						},
						"algorithm": schema.StringAttribute{
							Description: "Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).",
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"format": schema.StringAttribute{
							Description: "Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).\nAvailable values: \"raw\", \"pkcs8\", \"spki\", \"jwk\".",
							Optional:    true,
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
							Optional:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"key_base64": schema.StringAttribute{
							Description: "Base64-encoded key data. Required if `format` is \"raw\", \"pkcs8\", or \"spki\".",
							Optional:    true,
							Sensitive:   true,
						},
						"key_jwk": schema.StringAttribute{
							Description: "Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is \"jwk\".",
							Optional:    true,
							Sensitive:   true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"workflow_name": schema.StringAttribute{
							Description: "Name of the Workflow to bind to.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{RequiresReplaceIfConfiguredIgnoringSensitiveTextDiff()},
			},
			"limits": schema.SingleNestedAttribute{
				Description: "Resource limits enforced at runtime.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerVersionLimitsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cpu_ms": schema.Int64Attribute{
						Description: "CPU time limit in milliseconds.",
						Required:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
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
			"source": schema.StringAttribute{
				Description: "The client used to create the version.",
				Computed:    true,
			},
			"main_script_base64": schema.StringAttribute{
				Description: "The base64-encoded main script content. This is only returned for service worker syntax workers (not ES modules). Used when importing existing workers that use the older service worker syntax.",
				Computed:    true,
			},
			"startup_time_ms": schema.Int64Attribute{
				Description: "Time in milliseconds spent on [Worker startup](https://developers.cloudflare.com/workers/platform/limits/#worker-startup-time).",
				Computed:    true,
			},
		},
	}
}

func (r *WorkerVersionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkerVersionResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
