// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker_version

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
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
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Individual Workers Admin",
				"Individual Workers Content Read-Only",
				"Individual Workers Editor",
				"Workers Admin",
				"Workers Content Read-Only",
				"Workers Editor",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Version identifier.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown(), stringplanmodifier.RequiresReplace()},
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
			"deploy": schema.BoolAttribute{
				Description:   "If true, a deployment will be created that sends 100% of traffic to the new version.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"include": schema.StringAttribute{
				Description: "Optional properties to include in the response. `modules` includes the `modules` property, which contains code and may add several megabytes to the response size. `modules,sourcemaps` also includes the version's source maps in `modules`.\nAvailable values: \"modules\", \"modules,sourcemaps\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("modules", "modules,sourcemaps"),
				},
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
			"containers": schema.SetNestedAttribute{
				Description: "List of containers attached to a Worker. Containers can only be attached to Durable Object classes of this Worker script.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"class_name": schema.StringAttribute{
							Description: "Select which Durable Object class should get this container attached.",
							Required:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.Set{setplanmodifier.RequiresReplace()},
			},
			"exports": schema.MapNestedAttribute{
				Description: "Declarative exports for the version, including Durable Object\nclasses (with their `storage` backend) and named Worker\nentrypoints. On reads, tombstoned lifecycle entries are\nomitted, so only live exports (`created` and\n`expecting-transfer`) are returned. `exports` and `migrations`\nare mutually exclusive on upload.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Description: "For `type: worker`: Marks this entry as a Worker entrypoint export.\n\nFor `type: durable-object`: Marks this entry as a Durable Object export.\nAvailable values: \"worker\", \"durable-object\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("worker", "durable-object"),
							},
						},
						"cache": schema.SingleNestedAttribute{
							Description: "For `type: worker`: Cache override for this entrypoint. Overrides the Worker's\nglobal `cache_options.enabled` for this entrypoint only.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether caching is enabled for this entrypoint.",
									Required:    true,
								},
							},
						},
						"state": schema.StringAttribute{
							Description: "For `state: created`: Live export. May be omitted; defaults to `created`.\n\nFor `state: deleted`: Tombstone that deletes the namespace.\n\nFor `state: renamed`: Tombstone that renames the namespace's class.\n\nFor `state: transferred`: Tombstone that transfers the namespace to another script.\n\nFor `state: expecting-transfer`: Target side of a two-phase transfer.\nAvailable values: \"created\", \"deleted\", \"renamed\", \"transferred\", \"expecting-transfer\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"created",
									"deleted",
									"renamed",
									"transferred",
									"expecting-transfer",
								),
							},
						},
						"storage": schema.StringAttribute{
							Description: "For `type: durable-object` and `state: created`: Durable Object storage backend. `sqlite` is the recommended (and\nonly) backend for new namespaces. `legacy-kv` is accepted only for\na class whose namespace already exists as KV-backed; the `exports`\nflow never provisions a new `legacy-kv` namespace.\n\nFor `type: durable-object` and `state: expecting-transfer`: Durable Object storage backend. `sqlite` is the recommended (and\nonly) backend for new namespaces. `legacy-kv` is accepted only for\na class whose namespace already exists as KV-backed; the `exports`\nflow never provisions a new `legacy-kv` namespace.\nAvailable values: \"sqlite\", \"legacy-kv\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("sqlite", "legacy-kv"),
							},
						},
						"container": schema.StringAttribute{
							Description: "For `type: durable-object` and `state: created`: Name of the container (declared in the upload's\n`metadata.containers`) that backs this Durable Object. When\nset, the namespace is container-enabled. Valid only on live\nentries.\n\nFor `state: expecting-transfer`: Name of the container (declared in the upload's\n`metadata.containers`) that backs this Durable Object once the\ntransfer settles. Valid only on live entries.",
							Optional:    true,
						},
						"renamed_to": schema.StringAttribute{
							Description: "For `state: renamed`: The destination class name. Must differ from the source class\n(the map key) and must be declared as a live (`created`) entry\nin the same `exports` map. Write-only: never present in GET\nresponses.",
							Optional:    true,
						},
						"transferred_to": schema.StringAttribute{
							Description: "For `state: transferred`: The destination script name. Must be in the same account and\nthe same dispatch-namespace context (or both non-dispatch).\nCross-dispatch-namespace transfers are rejected. Write-only:\nnever present in GET responses.",
							Optional:    true,
						},
						"transfer_from": schema.StringAttribute{
							Description: "For `state: expecting-transfer`: The source script name to receive the namespace from. Must be\nin the same account and dispatch-namespace context. Present on\nreads for `expecting-transfer` entries.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.Map{mapplanmodifier.RequiresReplace()},
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
				Description: "Code, sourcemaps, and other content used at runtime.\n\nSource maps use content type `application/source-map`. Responses include them only when requested with\n`include=modules,sourcemaps`. Workers stores the source map fields used to remap stack traces, so a returned\nsource map is equivalent to the uploaded file but not byte-identical.\n\nThis includes [`_headers`](https://developers.cloudflare.com/workers/static-assets/headers/#custom-headers) and\n[`_redirects`](https://developers.cloudflare.com/workers/static-assets/redirects/) files used to configure\n[Static Assets](https://developers.cloudflare.com/workers/static-assets/). `_headers` and `_redirects` files should be\nincluded as modules named `_headers` and `_redirects` with content type `text/plain`.",
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
			"package_dependencies": schema.ListNestedAttribute{
				Description: "The list of npm packages that were installed and used when this Worker\nversion was built.",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"installed_version": schema.StringAttribute{
							Description: "The exact version that was resolved and installed by the package manager.",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "The npm package name.",
							Required:    true,
						},
						"package_json_version": schema.StringAttribute{
							Description: "The version constraint as written in package.json.",
							Required:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{listplanmodifier.RequiresReplace()},
			},
			"placement": schema.SingleNestedAttribute{
				Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Specify mode='smart' for Smart Placement, or one of region/hostname/host.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\n\nTargeted placement mode.\nAvailable values: \"smart\", \"targeted\".",
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
				Description:   "Enables or disables specific Workers runtime features, such as upcoming features or changes outside the `compatibility_date`.\n\nResponses list only the flags this version sets explicitly and exclude behavior that `compatibility_date` enables. Responses omit the field when the list is empty, so a missing field means the version sets zero flags.",
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
						Description: "Human-readable message about the version. Truncated to 1000 bytes if longer.",
						Optional:    true,
					},
					"workers_tag": schema.StringAttribute{
						Description: "User-provided identifier for the version. Maximum 100 bytes.",
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
				CustomType:  customfield.NewNestedObjectType[WorkerVersionAssetsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Description: "Configuration for assets within a Worker.",
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerVersionAssetsConfigModel](ctx),
						Attributes: map[string]schema.Attribute{
							"base_path": schema.StringAttribute{
								Description: "The public URL path prefix under which assets are served. A null request value resets it to `/`; responses represent the root as `/`. All versions in a gradual deployment must use the same canonical value. To change it, first deploy the version containing the change at 100%.",
								Computed:    true,
								Optional:    true,
								Default:     stringdefault.StaticString("/"),
							},
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
							Description: "For `type: ai` or `type: ai_search` or `type: ai_search_namespace` or `type: messaging` or `type: analytics_engine` or `type: artifacts` or `type: assets` or `type: browser` or `type: d1` or `type: data_blob` or `type: dispatch_namespace` or `type: durable_object_namespace` or `type: hyperdrive` or `type: images` or `type: json` or `type: kv_namespace` or `type: media` or `type: mtls_certificate` or `type: plain_text` or `type: pipelines` or `type: k2` or `type: queue` or `type: ratelimit` or `type: r2_bucket` or `type: secret_text` or `type: send_email` or `type: service` or `type: text_blob` or `type: vectorize` or `type: version_metadata` or `type: secrets_store_secret` or `type: flagship` or `type: secret_key` or `type: workflow` or `type: wasm_module` or `type: vpc_service` or `type: vpc_network` or `type: worker_loader`: A JavaScript variable name for the binding.\n\nFor `type: inherit`: The name of the inherited binding.",
							Required:    true,
						},
						"type": schema.StringAttribute{
							Description: "The kind of resource that the binding provides.\nAvailable values: \"ai\", \"ai_search\", \"ai_search_namespace\", \"messaging\", \"analytics_engine\", \"artifacts\", \"assets\", \"browser\", \"d1\", \"data_blob\", \"dispatch_namespace\", \"durable_object_namespace\", \"hyperdrive\", \"inherit\", \"images\", \"json\", \"kv_namespace\", \"media\", \"mtls_certificate\", \"plain_text\", \"pipelines\", \"k2\", \"queue\", \"ratelimit\", \"r2_bucket\", \"secret_text\", \"send_email\", \"service\", \"text_blob\", \"vectorize\", \"version_metadata\", \"secrets_store_secret\", \"flagship\", \"secret_key\", \"workflow\", \"wasm_module\", \"vpc_service\", \"vpc_network\", \"worker_loader\".",
							Required:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"ai",
									"ai_search",
									"ai_search_namespace",
									"messaging",
									"analytics_engine",
									"artifacts",
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
									"media",
									"mtls_certificate",
									"plain_text",
									"pipelines",
									"k2",
									"queue",
									"ratelimit",
									"r2_bucket",
									"secret_text",
									"send_email",
									"service",
									"text_blob",
									"vectorize",
									"version_metadata",
									"secrets_store_secret",
									"flagship",
									"secret_key",
									"workflow",
									"wasm_module",
									"vpc_service",
									"vpc_network",
									"worker_loader",
								),
							},
						},
						"instance_name": schema.StringAttribute{
							Description: "For `type: ai_search`: The user-chosen instance name. Must exist at deploy time. The worker can search, chat, update, and manage items/jobs on this instance.",
							Optional:    true,
						},
						"namespace": schema.StringAttribute{
							Description: "For `type: ai_search`: The namespace the instance belongs to. Defaults to \"default\" if omitted. Customers who don't use namespaces can simply omit this field.\n\nFor `type: ai_search_namespace`: The user-chosen namespace name. Must exist before deploy -- Wrangler handles auto-creation on deploy failure (R2 bucket pattern). The \"default\" namespace is auto-created by config-api for new accounts. Grants full access (CRUD + search + chat) to all instances within the namespace.\n\nFor `type: messaging`: The Messaging namespace to bind to.\n\nFor `type: artifacts`: The Artifacts namespace exposed to the Worker in the Worker's account. Must be 2-63 characters, start with an ASCII alphanumeric character, contain only ASCII alphanumeric characters, dots, underscores, and hyphens, and must not end with a hyphen. The namespace does not need to be created before binding it.\n\nFor `type: dispatch_namespace`: The name of the dispatch namespace.",
							Optional:    true,
						},
						"dataset": schema.StringAttribute{
							Description: "For `type: analytics_engine`: The name of the dataset to bind to.",
							Optional:    true,
						},
						"database_id": schema.StringAttribute{
							Description: "For `type: d1`: Identifier of the D1 database to bind to.",
							Computed:    true,
							Optional:    true,
						},
						"id": schema.StringAttribute{
							Description: "For `type: d1`: Identifier of the D1 database to bind to.\n\nFor `type: hyperdrive`: Identifier of the Hyperdrive connection to bind to.",
							Optional:    true,
						},
						"part": schema.StringAttribute{
							Description: "For `type: data_blob`: The name of the file containing the data content. Only accepted for `service worker syntax` Workers.\n\nFor `type: text_blob`: The name of the file containing the text content. Only accepted for `service worker syntax` Workers.\n\nFor `type: wasm_module`: The name of the file containing the WebAssembly module content. Only accepted for `service worker syntax` Workers.",
							Optional:    true,
						},
						"outbound": schema.SingleNestedAttribute{
							Description: "For `type: dispatch_namespace`: Outbound worker.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"params": schema.ListNestedAttribute{
									Description: "Pass information from the Dispatch Worker to the Outbound Worker through the parameters.",
									Optional:    true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description: "Name of the parameter.",
												Required:    true,
											},
										},
									},
								},
								"worker": schema.SingleNestedAttribute{
									Description: "Outbound worker.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"entrypoint": schema.StringAttribute{
											Description: "Entrypoint to invoke on the outbound worker.",
											Optional:    true,
										},
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
							Description: "For `type: durable_object_namespace`: The exported class name of the Durable Object.\n\nFor `type: workflow`: Class name of the Workflow. Should only be provided if the Workflow belongs to this script.",
							Computed:    true,
							Optional:    true,
							PlanModifiers: []planmodifier.String{
								UnknownOnlyIf("type", "durable_object_namespace"),
							},
						},
						"dispatch_namespace": schema.StringAttribute{
							Description: "For `type: durable_object_namespace`: The dispatch namespace the Durable Object script belongs to.",
							Optional:    true,
						},
						"environment": schema.StringAttribute{
							Description:   "For `type: durable_object_namespace`: The environment of the script_name to bind to.\n\nFor `type: service`: Optional environment if the Worker utilizes one.",
							Computed:      true,
							Optional:      true,
							PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
						},
						"namespace_id": schema.StringAttribute{
							Description: "For `type: durable_object_namespace` or `type: kv_namespace`: Namespace identifier tag.\n\nFor `type: ratelimit`: Identifier of the rate limit namespace to bind to.",
							Computed:    true,
							Optional:    true,
							PlanModifiers: []planmodifier.String{
								UnknownOnlyIf("type", "durable_object_namespace"),
							},
						},
						"script_name": schema.StringAttribute{
							Description: "For `type: durable_object_namespace`: The script where the Durable Object is defined, if it is external to this Worker.\n\nFor `type: workflow`: Script name that contains the Workflow. If not provided, defaults to this script name.",
							Computed:    true,
							Optional:    true,
							PlanModifiers: []planmodifier.String{
								UnknownOnlyIf("type", "durable_object_namespace"),
							},
						},
						"old_name": schema.StringAttribute{
							Description: "For `type: inherit`: The old name of the inherited binding. If set, the binding will be renamed from `old_name` to `name` in the new version. If not set, the binding will keep the same name between versions.",
							Optional:    true,
						},
						"version_id": schema.StringAttribute{
							Description: "For `type: inherit`: Identifier for the version to inherit the binding from. This can be a version ID, or the literal \"latest\" to inherit from the most recently uploaded version, which may not be the deployed version. Defaults to \"latest\". Only the Workers API version endpoints under `/accounts/{account_id}/workers/workers/{worker_id}/versions` accept a version ID. Other endpoints, including script upload and script version upload, accept only \"latest\" and reject a version ID with error 10057.",
							Optional:    true,
						},
						"json": schema.StringAttribute{
							Description: "For `type: json`: JSON data to use.",
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"certificate_id": schema.StringAttribute{
							Description: "For `type: mtls_certificate`: Identifier of the certificate to bind to.",
							Optional:    true,
						},
						"text": schema.StringAttribute{
							Description: "For `type: plain_text`: The text value to use.\n\nFor `type: secret_text`: The secret value to use.",
							Optional:    true,
							Sensitive:   true,
						},
						"pipeline": schema.StringAttribute{
							Description: "For `type: pipelines`: Name of the Pipeline to bind to.",
							Optional:    true,
						},
						"stream": schema.StringAttribute{
							Description: "For `type: k2`: ID of a K2 stream owned by the account deploying the Worker.",
							Optional:    true,
						},
						"queue_name": schema.StringAttribute{
							Description: "For `type: queue`: Name of the Queue to bind to.",
							Optional:    true,
						},
						"simple": schema.SingleNestedAttribute{
							Description: "For `type: ratelimit`: The rate limit configuration.",
							Optional:    true,
							Attributes: map[string]schema.Attribute{
								"limit": schema.Float64Attribute{
									Description: "The limit (requests per period).",
									Required:    true,
								},
								"period": schema.Int64Attribute{
									Description: "The period in seconds.",
									Required:    true,
								},
								"mitigation_timeout": schema.Int64Attribute{
									Description: "Duration in seconds to apply the mitigation action after the rate limit is exceeded. Valid values are 0 (disabled), 10, or multiples of 60 up to 86400. Must be greater than or equal to the period when non-zero.",
									Optional:    true,
								},
							},
						},
						"bucket_name": schema.StringAttribute{
							Description: "For `type: r2_bucket`: R2 bucket to bind to.",
							Optional:    true,
						},
						"jurisdiction": schema.StringAttribute{
							Description: "For `type: r2_bucket`: The [jurisdiction](https://developers.cloudflare.com/r2/reference/data-location/#jurisdictional-restrictions) of the R2 bucket.\nAvailable values: \"eu\", \"fedramp\", \"fedramp-high\", \"us\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"eu",
									"fedramp",
									"fedramp-high",
									"us",
								),
							},
						},
						"allowed_destination_addresses": schema.ListAttribute{
							Description: "For `type: send_email`: List of allowed destination addresses.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"allowed_sender_addresses": schema.ListAttribute{
							Description: "For `type: send_email`: List of allowed sender addresses.",
							Optional:    true,
							ElementType: types.StringType,
						},
						"destination_address": schema.StringAttribute{
							Description: "For `type: send_email`: Destination address for the email.",
							Optional:    true,
						},
						"service": schema.StringAttribute{
							Description: "For `type: service`: Name of Worker to bind to.",
							Optional:    true,
						},
						"entrypoint": schema.StringAttribute{
							Description: "For `type: service`: Entrypoint to invoke on the target Worker.",
							Optional:    true,
						},
						"index_name": schema.StringAttribute{
							Description: "For `type: vectorize`: Name of the Vectorize index to bind to.",
							Optional:    true,
						},
						"secret_name": schema.StringAttribute{
							Description: "For `type: secrets_store_secret`: Name of the secret in the store.",
							Optional:    true,
						},
						"store_id": schema.StringAttribute{
							Description: "For `type: secrets_store_secret`: ID of the store containing the secret.",
							Optional:    true,
						},
						"app_id": schema.StringAttribute{
							Description: "For `type: flagship`: ID of the Flagship app to bind to for feature flag evaluation.",
							Optional:    true,
						},
						"algorithm": schema.StringAttribute{
							Description: "For `type: secret_key`: Algorithm-specific key parameters. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#algorithm).",
							Optional:    true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"format": schema.StringAttribute{
							Description: "For `type: secret_key`: Data format of the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#format).\nAvailable values: \"raw\", \"pkcs8\", \"spki\", \"jwk\".",
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
							Description: "For `type: secret_key`: Allowed operations with the key. [Learn more](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#keyUsages).",
							Optional:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"key_base64": schema.StringAttribute{
							Description: "For `type: secret_key`: Base64-encoded key data. Required if `format` is \"raw\", \"pkcs8\", or \"spki\".",
							Optional:    true,
							Sensitive:   true,
						},
						"key_jwk": schema.StringAttribute{
							Description: "For `type: secret_key`: Key data in [JSON Web Key](https://developer.mozilla.org/en-US/docs/Web/API/SubtleCrypto/importKey#json_web_key) format. Required if `format` is \"jwk\".",
							Optional:    true,
							Sensitive:   true,
							CustomType:  jsontypes.NormalizedType{},
						},
						"workflow_name": schema.StringAttribute{
							Description: "For `type: workflow`: Name of the Workflow to bind to.",
							Optional:    true,
						},
						"service_id": schema.StringAttribute{
							Description: "For `type: vpc_service`: Identifier of the VPC service to bind to.",
							Optional:    true,
						},
						"identity": schema.StringAttribute{
							Description: "For `type: vpc_network`: Enables Gateway identity for the binding. Requires network_id to be \"cf1:network\" and cannot be combined with tunnel_id or an explicit \"cf1:vnet:<UUID>\" network_id.\nAvailable values: \"runtime-email-alpha\".",
							Optional:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("runtime-email-alpha"),
							},
						},
						"network_id": schema.StringAttribute{
							Description: "For `type: vpc_network`: Identifier of the network to bind to. Use \"cf1:network\" to select the account's default virtual network at runtime, or \"cf1:vnet:<UUID>\" to select an explicit virtual network in the Worker's account. Mutually exclusive with tunnel_id.",
							Optional:    true,
						},
						"tunnel_id": schema.StringAttribute{
							Description: "For `type: vpc_network`: UUID of the Cloudflare Tunnel to bind to. Mutually exclusive with network_id.",
							Optional:    true,
						},
					},
				},
				PlanModifiers: []planmodifier.List{RequiresReplaceIfConfiguredIgnoringSensitiveTextDiff()},
			},
			"cache_options": schema.SingleNestedAttribute{
				Description: "Global CacheW configuration for the Worker. When caching is on,\nthe platform provisions a `cloudflare.app` zone for the Worker.\nA `type: worker` entry in the `exports` map can override this\nvalue for a single entrypoint.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerVersionCacheOptionsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether caching is enabled for this Worker.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"cross_version_cache": schema.BoolAttribute{
						Description: "Whether cached responses are shared across Worker version\nuploads. This is independent of `enabled`. It can stay true\nwhile caching is off, so the preference survives turning\ncaching off and back on.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
			},
			"limits": schema.SingleNestedAttribute{
				Description: "Resource limits enforced at runtime.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerVersionLimitsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cpu_ms": schema.Int64Attribute{
						Description: "CPU time limit in milliseconds.",
						Optional:    true,
					},
					"subrequests": schema.Int64Attribute{
						Description: "Subrequest limit per request.",
						Optional:    true,
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.RequiresReplaceIfConfigured()},
			},
			"author_email": schema.StringAttribute{
				Description: "Email of the user who created the version.",
				Computed:    true,
			},
			"author_id": schema.StringAttribute{
				Description: "Identifier of the user who created the version.",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Description:   "When the version was created.",
				Computed:      true,
				CustomType:    timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"migration_tag": schema.StringAttribute{
				Description: "Durable Object migration tag. Set when the version is deployed. Omitted if the version has not been deployed or the Worker does not use Durable Objects.",
				Computed:    true,
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
			"urls": schema.ListAttribute{
				Description: "All routable URLs that always point to this version. Does not include alias URLs, since aliases can be updated to point to a different version.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"exports_reconciliation": schema.SingleNestedAttribute{
				Description: "Summary of the declarative exports reconciliation that ran on this upload. Populated only when the uploaded metadata included an `exports` block. Durable Object entries drive reconciliation; `type: worker` entries do not contribute to this summary.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerVersionExportsReconciliationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"created": schema.ListAttribute{
						Description: "Class names for which a new namespace was provisioned.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"deleted": schema.ListAttribute{
						Description: "Class names whose namespace was deleted by a `deleted` tombstone.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"info": schema.ListNestedAttribute{
						Description: "Non-blocking info entries (stale tombstones, tombstone applied with class still in code). See `exports_reconciliation_info`.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerVersionExportsReconciliationInfoModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"class": schema.StringAttribute{
									Description: "The class name the info entry is about.",
									Computed:    true,
								},
								"message": schema.StringAttribute{
									Description: "Human-readable explanation.",
									Computed:    true,
								},
								"scenario": schema.StringAttribute{
									Description: "Stable, machine-readable tag identifying which reconciliation scenario produced an error, warning, or info entry. Clients may branch on this value instead of parsing `message`.",
									Computed:    true,
								},
								"namespace_id": schema.StringAttribute{
									Description: "The provisioned namespace the entry relates to, when applicable.",
									Computed:    true,
								},
								"referencing_scripts": schema.ListAttribute{
									Description: "Other Workers in the account that still bind to the affected class. Advisory: while non-empty the tombstone is not yet safe to remove — redeploy these Workers with bindings re-pointed first.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
							},
						},
					},
					"removable_entries": schema.ListAttribute{
						Description: "Source class names whose tombstone entry is now stale and safe to delete from `exports` (no remaining referencing scripts).",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"renamed": schema.ListNestedAttribute{
						Description: "Applied `renamed` tombstones.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerVersionExportsReconciliationRenamedModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from": schema.StringAttribute{
									Description: "The original (source) class name.",
									Computed:    true,
								},
								"to": schema.StringAttribute{
									Description: "The new class name (`renamed_to`).",
									Computed:    true,
								},
							},
						},
					},
					"transfer_pending": schema.ListNestedAttribute{
						Description: "Phase-1 transfer hints recorded on the target side.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerVersionExportsReconciliationTransferPendingModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"class": schema.StringAttribute{
									Description: "The target-side class name awaiting transfer.",
									Computed:    true,
								},
								"from": schema.StringAttribute{
									Description: "The source script the namespace will be transferred from.",
									Computed:    true,
								},
							},
						},
					},
					"transferred": schema.ListNestedAttribute{
						Description: "Committed `transferred` tombstones (phase-2).",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerVersionExportsReconciliationTransferredModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"class": schema.StringAttribute{
									Description: "The source class name that was transferred.",
									Computed:    true,
								},
								"phase": schema.StringAttribute{
									Description: "The transfer phase. Currently always `committed`.",
									Computed:    true,
								},
								"to": schema.StringAttribute{
									Description: "The destination script that now owns the namespace.",
									Computed:    true,
								},
							},
						},
					},
					"updated": schema.ListAttribute{
						Description: "Class names whose provisioned namespace was mutated in place.",
						Computed:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"warnings": schema.ListNestedAttribute{
						Description: "Non-blocking warnings. See `exports_reconciliation_warning`.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerVersionExportsReconciliationWarningsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"class": schema.StringAttribute{
									Description: "The class name the warning is about.",
									Computed:    true,
								},
								"message": schema.StringAttribute{
									Description: "Human-readable explanation of the warning.",
									Computed:    true,
								},
								"scenario": schema.StringAttribute{
									Description: "Stable, machine-readable tag identifying which reconciliation scenario produced an error, warning, or info entry. Clients may branch on this value instead of parsing `message`.",
									Computed:    true,
								},
								"namespace_id": schema.StringAttribute{
									Description: "The provisioned namespace the warning relates to, when applicable.",
									Computed:    true,
								},
							},
						},
					},
				},
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
