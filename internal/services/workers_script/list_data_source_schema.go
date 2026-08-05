// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_script

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersScriptsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Optional:    true,
			},
			"tags": schema.StringAttribute{
				Description: "Filter scripts by tags. Format: comma-separated list of tag:allowed pairs where allowed is 'yes' or 'no'.",
				Optional:    true,
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
				CustomType:  customfield.NewNestedObjectListType[WorkersScriptsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The name used to identify the script.",
							Computed:    true,
						},
						"cache_options": schema.SingleNestedAttribute{
							Description: "Global CacheW configuration for the Worker. When caching is on,\nthe platform provisions a `cloudflare.app` zone for the Worker.\nA `type: worker` entry in the `exports` map can override this\nvalue for a single entrypoint.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersScriptsCacheOptionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether caching is enabled for this Worker.",
									Computed:    true,
								},
								"cross_version_cache": schema.BoolAttribute{
									Description: "Whether cached responses are shared across Worker version\nuploads. This is independent of `enabled`. It can stay true\nwhile caching is off, so the preference survives turning\ncaching off and back on.",
									Computed:    true,
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
						"created_on": schema.StringAttribute{
							Description: "When the script was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"etag": schema.StringAttribute{
							Description: "Hashed script content, can be used in a If-None-Match header when updating.",
							Computed:    true,
						},
						"exports": schema.MapNestedAttribute{
							Description: "Declarative exports for the Worker's most recent version,\nincluding Durable Object classes (with their `storage`\nbackend) and named Worker entrypoints. Tombstoned lifecycle\nentries are omitted, so only live exports (`created` and\n`expecting-transfer`) are returned.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectMapType[WorkersScriptsExportsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description: "The kind of export.\nAvailable values: \"worker\", \"durable-object\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("worker", "durable-object"),
										},
									},
									"cache": schema.SingleNestedAttribute{
										Description: "Cache override for this entrypoint. It applies only to\n`type: worker` entries and overrides the Worker's global\n`cache_options.enabled` for that entrypoint.",
										Computed:    true,
										CustomType:  customfield.NewNestedObjectType[WorkersScriptsExportsCacheDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description: "Whether caching is enabled for this entrypoint.",
												Computed:    true,
											},
										},
									},
									"renamed_to": schema.StringAttribute{
										Description: "Destination class name for a `state: renamed` tombstone. The\ntarget must appear as a live (`created`) entry in the same\n`exports` map. Write-only: never present in GET responses.",
										Computed:    true,
									},
									"state": schema.StringAttribute{
										Description: "Lifecycle state of the export entry. Defaults to `created`\n(a normal, live export) when omitted.\n\n`deleted`, `renamed`, and `transferred` are tombstones:\nwrite-only lifecycle operations that retire, rename, or hand\noff a provisioned Durable Object namespace. They are applied\nat upload and are filtered out of GET responses, so a read\nonly ever returns `created` or `expecting-transfer`.\n\n`expecting-transfer` is a live export whose data is being\nreceived from another script via the two-phase transfer flow;\nit carries `storage` and `transfer_from`.\nAvailable values: \"created\", \"deleted\", \"renamed\", \"transferred\", \"expecting-transfer\".",
										Computed:    true,
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
										Description: "Storage backend for a `type: durable-object` export. Required\nfor live Durable Object entries (`created` and\n`expecting-transfer`). `sqlite` selects SQLite-backed storage;\n`legacy-kv` selects the legacy key-value storage.\nAvailable values: \"sqlite\", \"legacy-kv\".",
										Computed:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("sqlite", "legacy-kv"),
										},
									},
									"transfer_from": schema.StringAttribute{
										Description: "Source script for a `state: expecting-transfer` entry. The\nnamespace on this script is materialised from the source\nscript's data via the pending-transfer flow. Present on reads\nfor `expecting-transfer` entries.",
										Computed:    true,
									},
									"transferred_to": schema.StringAttribute{
										Description: "Destination script for a `state: transferred` tombstone. Must\nreference a script in the same account; cross-dispatch-namespace\ntransfers are rejected. Write-only: never present in GET\nresponses.",
										Computed:    true,
									},
								},
							},
						},
						"handlers": schema.ListAttribute{
							Description: "The names of handlers exported as part of the default export.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"has_assets": schema.BoolAttribute{
							Description: "Whether a Worker contains assets.",
							Computed:    true,
						},
						"has_modules": schema.BoolAttribute{
							Description: "Whether a Worker contains modules.",
							Computed:    true,
						},
						"last_deployed_from": schema.StringAttribute{
							Description: "The client most recently used to deploy this Worker.",
							Computed:    true,
						},
						"logpush": schema.BoolAttribute{
							Description: "Whether Logpush is turned on for the Worker.",
							Computed:    true,
						},
						"migration_tag": schema.StringAttribute{
							Description: "The tag of the Durable Object migration that was most recently applied for this Worker.",
							Computed:    true,
						},
						"modified_on": schema.StringAttribute{
							Description: "When the script was last modified.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"named_handlers": schema.ListNestedAttribute{
							Description: "Named exports, such as Durable Object class implementations and named entrypoints.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[WorkersScriptsNamedHandlersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"handlers": schema.ListAttribute{
										Description: "The names of handlers exported as part of the named export.",
										Computed:    true,
										CustomType:  customfield.NewListType[types.String](ctx),
										ElementType: types.StringType,
									},
									"name": schema.StringAttribute{
										Description: "The name of the export.",
										Computed:    true,
									},
								},
							},
						},
						"observability": schema.SingleNestedAttribute{
							Description: "Observability settings for the Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersScriptsObservabilityDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether observability is enabled for the Worker.",
									Computed:    true,
								},
								"head_sampling_rate": schema.Float64Attribute{
									Description: "The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.",
									Computed:    true,
								},
								"logs": schema.SingleNestedAttribute{
									Description: "Log settings for the Worker.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[WorkersScriptsObservabilityLogsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description: "Whether logs are enabled for the Worker.",
											Computed:    true,
										},
										"invocation_logs": schema.BoolAttribute{
											Description: "Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.",
											Computed:    true,
										},
										"destinations": schema.ListAttribute{
											Description: "A list of destinations where logs will be exported to.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"head_sampling_rate": schema.Float64Attribute{
											Description: "The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.",
											Computed:    true,
										},
										"persist": schema.BoolAttribute{
											Description: "Whether log persistence is enabled for the Worker.",
											Computed:    true,
										},
									},
								},
								"traces": schema.SingleNestedAttribute{
									Description: "Trace settings for the Worker.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[WorkersScriptsObservabilityTracesDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"destinations": schema.ListAttribute{
											Description: "A list of destinations where traces will be exported to.",
											Computed:    true,
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
										"enabled": schema.BoolAttribute{
											Description: "Whether traces are enabled for the Worker.",
											Computed:    true,
										},
										"head_sampling_rate": schema.Float64Attribute{
											Description: "The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.",
											Computed:    true,
										},
										"persist": schema.BoolAttribute{
											Description: "Whether trace persistence is enabled for the Worker.",
											Computed:    true,
										},
										"propagation_policy": schema.StringAttribute{
											Description: "Controls how inbound trace context (traceparent/tracestate) headers on incoming requests are handled. \"authenticated\" (default) honors inbound trace context only when accompanied by a valid trace auth token. \"accept\" unconditionally accepts inbound trace context. Requires the trace propagation feature to be enabled.\nAvailable values: \"authenticated\", \"accept\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("authenticated", "accept"),
											},
										},
									},
								},
							},
						},
						"placement": schema.SingleNestedAttribute{
							Description: "Configuration for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Specify mode='smart' for Smart Placement, or one of region/hostname/host.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersScriptsPlacementDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"mode": schema.StringAttribute{
									Description: "Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"smart\", \"targeted\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive("smart", "targeted"),
									},
								},
								"last_analyzed_at": schema.StringAttribute{
									Description: "The last time the script was analyzed for [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"status": schema.StringAttribute{
									Description: "Status of [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement).\nAvailable values: \"SUCCESS\", \"UNSUPPORTED_APPLICATION\", \"INSUFFICIENT_INVOCATIONS\".",
									Computed:    true,
									Validators: []validator.String{
										stringvalidator.OneOfCaseInsensitive(
											"SUCCESS",
											"UNSUPPORTED_APPLICATION",
											"INSUFFICIENT_INVOCATIONS",
										),
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
									CustomType:  customfield.NewNestedObjectListType[WorkersScriptsPlacementTargetDataSourceModel](ctx),
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
						"placement_mode": schema.StringAttribute{
							Description:        `Available values: "smart", "targeted".`,
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive("smart", "targeted"),
							},
						},
						"placement_status": schema.StringAttribute{
							Description:        `Available values: "SUCCESS", "UNSUPPORTED_APPLICATION", "INSUFFICIENT_INVOCATIONS".`,
							Computed:           true,
							DeprecationMessage: "This attribute is deprecated.",
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"SUCCESS",
									"UNSUPPORTED_APPLICATION",
									"INSUFFICIENT_INVOCATIONS",
								),
							},
						},
						"routes": schema.ListNestedAttribute{
							Description: "Routes associated with the Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[WorkersScriptsRoutesDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"id": schema.StringAttribute{
										Description: "Identifier.",
										Computed:    true,
									},
									"pattern": schema.StringAttribute{
										Description: "Pattern to match incoming requests against. [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).",
										Computed:    true,
									},
									"script": schema.StringAttribute{
										Description: "Name of the script to run if the route matches.",
										Computed:    true,
									},
								},
							},
						},
						"tag": schema.StringAttribute{
							Description: "The immutable ID of the script.",
							Computed:    true,
						},
						"tags": schema.SetAttribute{
							Description: "Tags associated with the Worker.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"tail_consumers": schema.SetNestedAttribute{
							Description: "List of Workers that will consume logs from the attached Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectSetType[WorkersScriptsTailConsumersDataSourceModel](ctx),
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
						"usage_model": schema.StringAttribute{
							Description: "Usage model for the Worker invocations.\nAvailable values: \"standard\", \"bundled\", \"unbound\".",
							Computed:    true,
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

func (d *WorkersScriptsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersScriptsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
