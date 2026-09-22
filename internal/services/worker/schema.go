// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*WorkerResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
			},
		}.String(),
		Version: 500,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Immutable ID of the Worker.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseNonNullStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description:   "Identifier.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "Name of the Worker.",
				Required:    true,
			},
			"logpush": schema.BoolAttribute{
				Description: "Whether logpush is enabled for the Worker.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"tags": schema.SetAttribute{
				Description: "Tags associated with the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewSetType[types.String](ctx),
				ElementType: types.StringType,
				Default:     setdefault.StaticValue(customfield.NewSetMust[types.String](ctx, nil).SetValue),
			},
			"observability": schema.SingleNestedAttribute{
				Description: "Observability settings for the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerObservabilityModel](ctx),
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether observability is enabled for the Worker.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"head_sampling_rate": schema.Float64Attribute{
						Description: "The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).",
						Computed:    true,
						Optional:    true,
						Default:     float64default.StaticFloat64(1),
						PlanModifiers: []planmodifier.Float64{
							NormalizeFloat64(),
						},
					},
					"issues": schema.SingleNestedAttribute{
						Description: "Real-time Issues settings for the Worker.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerObservabilityIssuesModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether real-time Issues are enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
						},
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
					},
					"logs": schema.SingleNestedAttribute{
						Description: "Log settings for the Worker.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerObservabilityLogsModel](ctx),
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{
							"destinations": schema.ListAttribute{
								Description: "A list of destinations where logs will be exported to.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
								PlanModifiers: []planmodifier.List{
									listplanmodifier.UseStateForUnknown(),
								},
							},
							"enabled": schema.BoolAttribute{
								Description: "Whether logs are enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"head_sampling_rate": schema.Float64Attribute{
								Description: "The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).",
								Computed:    true,
								Optional:    true,
								Default:     float64default.StaticFloat64(1),
								PlanModifiers: []planmodifier.Float64{
									NormalizeFloat64(),
								},
							},
							"invocation_logs": schema.BoolAttribute{
								Description: "Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(true),
							},
							"persist": schema.BoolAttribute{
								Description: "Whether log persistence is enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(true),
							},
						},
					},

					"traces": schema.SingleNestedAttribute{
						Description: "Trace settings for the Worker.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerObservabilityTracesModel](ctx),
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{
							"destinations": schema.ListAttribute{
								Description: "A list of destinations where traces will be exported to.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewListType[types.String](ctx),
								ElementType: types.StringType,
								PlanModifiers: []planmodifier.List{
									listplanmodifier.UseStateForUnknown(),
								},
							},
							"enabled": schema.BoolAttribute{
								Description: "Whether traces are enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"head_sampling_rate": schema.Float64Attribute{
								Description: "The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%).",
								Computed:    true,
								Optional:    true,
								Default:     float64default.StaticFloat64(1),
							},
							"persist": schema.BoolAttribute{
								Description: "Whether trace persistence is enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(true),
							},
							"propagation_policy": schema.StringAttribute{
								Description: "Controls how inbound trace context (traceparent/tracestate) headers on incoming requests are handled. \"authenticated\" (default) honors inbound trace context only when accompanied by a valid trace auth token. \"accept\" unconditionally accepts inbound trace context. Requires the trace propagation feature to be enabled.\nAvailable values: \"authenticated\", \"accept\".",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("authenticated", "accept"),
								},
								PlanModifiers: []planmodifier.String{
									PropagationPolicyDefault(),
								},
							},
						},
						Default: objectdefault.StaticValue(customfield.NewObjectMust(ctx, &WorkerObservabilityTracesModel{
							Enabled:          types.BoolValue(false),
							HeadSamplingRate: types.Float64Value(1),
							Persist:          types.BoolValue(true),
							Destinations:     customfield.NewListMust[types.String](ctx, nil),
						}).ObjectValue),
					},
				},
				Default: objectdefault.StaticValue(customfield.NewObjectMust(ctx, &WorkerObservabilityModel{
					Enabled:          types.BoolValue(false),
					HeadSamplingRate: types.Float64Value(1),
					Logs: customfield.NewObjectMust(ctx, &WorkerObservabilityLogsModel{
						Enabled:          types.BoolValue(false),
						HeadSamplingRate: types.Float64Value(1),
						InvocationLogs:   types.BoolValue(true),
						Persist:          types.BoolValue(true),
						Destinations:     customfield.NewListMust[types.String](ctx, nil),
					}),
					Traces: customfield.NewObjectMust(ctx, &WorkerObservabilityTracesModel{
						Enabled:          types.BoolValue(false),
						HeadSamplingRate: types.Float64Value(1),
						Persist:          types.BoolValue(true),
						Destinations:     customfield.NewListMust[types.String](ctx, nil),
					}),
				}).ObjectValue),
			},
			"previews_base_config": schema.SingleNestedAttribute{
				Description: "Template configuration used when creating new Previews for this Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigModel](ctx),
				Attributes: map[string]schema.Attribute{
					"cache_options": schema.SingleNestedAttribute{
						Description: "Cache options used when creating new Previews.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigCacheOptionsModel](ctx),
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
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
					},
					"env": schema.MapNestedAttribute{
						Description: "Bindings used when creating new Previews, keyed by binding name.",
						Optional:    true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description: "The kind of resource that the binding provides.",
									Required:    true,
								},
							},
						},
					},
					"limits": schema.SingleNestedAttribute{
						Description: "Resource limits enforced at runtime for newly created Previews.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigLimitsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"cpu_ms": schema.Int64Attribute{
								Description: "The amount of CPU time this Worker can use in milliseconds.",
								Optional:    true,
							},
							"subrequests": schema.Int64Attribute{
								Description: "The number of subrequests this Worker can make per request.",
								Optional:    true,
							},
						},
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
					},
					"logpush": schema.BoolAttribute{
						Description: "Whether logpush is enabled when creating new Previews.",
						Optional:    true,
					},
					"observability": schema.SingleNestedAttribute{
						Description: "Observability settings used when creating new Previews.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigObservabilityModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether observability is enabled for the Worker.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"head_sampling_rate": schema.Float64Attribute{
								Description: "The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).",
								Computed:    true,
								Optional:    true,
								Default:     float64default.StaticFloat64(1),
							},
							"issues": schema.SingleNestedAttribute{
								Description: "Real-time Issues settings for the Worker.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigObservabilityIssuesModel](ctx),
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description: "Whether real-time Issues are enabled for the Worker.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
								},
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
							},
							"logs": schema.SingleNestedAttribute{
								Description: "Log settings for the Worker.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigObservabilityLogsModel](ctx),
								Attributes: map[string]schema.Attribute{
									"destinations": schema.ListAttribute{
										Description:   "A list of destinations where logs will be exported to.",
										Computed:      true,
										Optional:      true,
										CustomType:    customfield.NewListType[types.String](ctx),
										ElementType:   types.StringType,
										PlanModifiers: []planmodifier.List{listplanmodifier.UseNonNullStateForUnknown()},
									},
									"enabled": schema.BoolAttribute{
										Description: "Whether logs are enabled for the Worker.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
									"head_sampling_rate": schema.Float64Attribute{
										Description: "The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).",
										Computed:    true,
										Optional:    true,
										Default:     float64default.StaticFloat64(1),
									},
									"invocation_logs": schema.BoolAttribute{
										Description: "Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(true),
									},
									"persist": schema.BoolAttribute{
										Description: "Whether log persistence is enabled for the Worker.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(true),
									},
								},
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
							},
							"redact_query_string": schema.BoolAttribute{
								Description: "Whether query strings are removed from request URLs in logs and traces.",
								Computed:    true,
								Optional:    true,
								Default:     booldefault.StaticBool(false),
							},
							"traces": schema.SingleNestedAttribute{
								Description: "Trace settings for the Worker.",
								Computed:    true,
								Optional:    true,
								CustomType:  customfield.NewNestedObjectType[WorkerPreviewsBaseConfigObservabilityTracesModel](ctx),
								Attributes: map[string]schema.Attribute{
									"destinations": schema.ListAttribute{
										Description:   "A list of destinations where traces will be exported to.",
										Computed:      true,
										Optional:      true,
										CustomType:    customfield.NewListType[types.String](ctx),
										ElementType:   types.StringType,
										PlanModifiers: []planmodifier.List{listplanmodifier.UseNonNullStateForUnknown()},
									},
									"enabled": schema.BoolAttribute{
										Description: "Whether traces are enabled for the Worker.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(false),
									},
									"head_sampling_rate": schema.Float64Attribute{
										Description: "The sampling rate for traces. From 0 to 1 (1 = 100%, 0.1 = 10%).",
										Computed:    true,
										Optional:    true,
										Default:     float64default.StaticFloat64(1),
									},
									"persist": schema.BoolAttribute{
										Description: "Whether trace persistence is enabled for the Worker.",
										Computed:    true,
										Optional:    true,
										Default:     booldefault.StaticBool(true),
									},
									"propagation_policy": schema.StringAttribute{
										Description: "Controls how inbound trace context (traceparent/tracestate) headers on incoming requests are handled. \"authenticated\" honors inbound trace context only when accompanied by a valid trace auth token. \"accept\" unconditionally accepts inbound trace context. Requires the trace propagation feature to be enabled. Returns null when the trace propagation feature is not enabled for the account.\nAvailable values: \"authenticated\", \"accept\".",
										Optional:    true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("authenticated", "accept"),
										},
									},
								},
								PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
							},
						},
						PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
					},
					"placement": schema.SingleNestedAttribute{
						Description: "Placement configuration used when creating new Previews.",
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
					},
					"tail_consumers": schema.SetNestedAttribute{
						Description: "Other Workers that should consume logs from newly created Previews.",
						Optional:    true,
						CustomType:  customfield.NewNestedObjectSetType[WorkerPreviewsBaseConfigTailConsumersModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description: "Name of the consumer Worker.",
									Required:    true,
								},
							},
						},
					},
				},
				PlanModifiers: []planmodifier.Object{objectplanmodifier.UseNonNullStateForUnknown()},
			},
			"subdomain": schema.SingleNestedAttribute{
				Description: "Subdomain settings for the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerSubdomainModel](ctx),
				Attributes: map[string]schema.Attribute{
					"enabled": schema.BoolAttribute{
						Description: "Whether the *.workers.dev subdomain is enabled for the Worker.",
						Computed:    true,
						Optional:    true,
						Default:     booldefault.StaticBool(false),
					},
					"preview_url_suffix": schema.StringAttribute{
						Description: "Prepend a version or preview prefix to this host suffix to form the *.workers.dev [preview URL](https://developers.cloudflare.com/workers/configuration/previews/) the Worker would serve on once previews are enabled, e.g. `https://<prefix>-my-worker.my-subdomain.workers.dev`. Present whenever the account owns a workers.dev subdomain, regardless of whether `previews_enabled` is true, so presence does not imply preview URLs are currently live. Absent only when the account owns no workers.dev subdomain.",
						Computed:    true,
					},
					"previews_enabled": schema.BoolAttribute{
						Description: "Whether [preview URLs](https://developers.cloudflare.com/workers/configuration/previews/) are enabled for the Worker.",
						Computed:    true,
						Optional:    true,
						PlanModifiers: []planmodifier.Bool{
							DefaultSubdomainPreviewsEnabledToEnabled(),
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"url": schema.StringAttribute{
						Description: "The address the Worker would serve on once its *.workers.dev subdomain is enabled. Present whenever the account owns a workers.dev subdomain, regardless of whether `enabled` is true, so presence does not imply the Worker is currently live at this URL. Absent only when the account owns no workers.dev subdomain.",
						Computed:    true,
					},
				},
				Default: objectdefault.StaticValue(customfield.NewObjectMust(ctx, &WorkerSubdomainModel{
					Enabled:         types.BoolValue(false),
					PreviewsEnabled: types.BoolValue(false),
				}).ObjectValue),
			},
			"tail_consumers": schema.SetNestedAttribute{
				Description: "Other Workers that should consume logs from the Worker.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectSetType[WorkerTailConsumersModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Description: "Name of the consumer Worker.",
							Required:    true,
						},
					},
				},
				Default: setdefault.StaticValue(customfield.NewSetMust[customfield.NestedObject[WorkerTailConsumersModel]](ctx, nil).SetValue),
			},
			"created_on": schema.StringAttribute{
				Description: "When the Worker was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{
					// created_on is set on Worker creation and never changes
					// after that.
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"deployed_on": schema.StringAttribute{
				Description: "When the Worker's most recent deployment was created. `null` if the Worker has never been deployed.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_on": schema.StringAttribute{
				Description: "When the Worker was most recently updated.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"references": schema.SingleNestedAttribute{
				Description: "Other resources that reference the Worker and depend on it existing.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[WorkerReferencesModel](ctx),
				Attributes: map[string]schema.Attribute{
					"dispatch_namespace_outbounds": schema.ListNestedAttribute{
						Description: "Other Workers that reference the Worker as an outbound for a dispatch namespace.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerReferencesDispatchNamespaceOutboundsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"namespace_id": schema.StringAttribute{
									Description: "ID of the dispatch namespace.",
									Computed:    true,
								},
								"namespace_name": schema.StringAttribute{
									Description: "Name of the dispatch namespace.",
									Computed:    true,
								},
								"worker_id": schema.StringAttribute{
									Description: "ID of the Worker using the dispatch namespace.",
									Computed:    true,
								},
								"worker_name": schema.StringAttribute{
									Description: "Name of the Worker using the dispatch namespace.",
									Computed:    true,
								},
							},
						},
					},
					"domains": schema.ListNestedAttribute{
						Description: "Custom domains connected to the Worker.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerReferencesDomainsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "ID of the custom domain.",
									Computed:    true,
								},
								"certificate_id": schema.StringAttribute{
									Description: "ID of the TLS certificate issued for the custom domain.",
									Computed:    true,
								},
								"hostname": schema.StringAttribute{
									Description: "Full hostname of the custom domain, including the zone name.",
									Computed:    true,
								},
								"zone_id": schema.StringAttribute{
									Description: "ID of the zone.",
									Computed:    true,
								},
								"zone_name": schema.StringAttribute{
									Description: "Name of the zone.",
									Computed:    true,
								},
							},
						},
					},
					"durable_objects": schema.ListNestedAttribute{
						Description: "Other Workers that reference Durable Object classes implemented by the Worker.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerReferencesDurableObjectsModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"namespace_id": schema.StringAttribute{
									Description: "ID of the Durable Object namespace being used.",
									Computed:    true,
								},
								"namespace_name": schema.StringAttribute{
									Description: "Name of the Durable Object namespace being used.",
									Computed:    true,
								},
								"worker_id": schema.StringAttribute{
									Description: "ID of the Worker using the Durable Object implementation.",
									Computed:    true,
								},
								"worker_name": schema.StringAttribute{
									Description: "Name of the Worker using the Durable Object implementation.",
									Computed:    true,
								},
							},
						},
					},
					"queues": schema.ListNestedAttribute{
						Description: "Queues that send messages to the Worker.",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerReferencesQueuesModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"queue_consumer_id": schema.StringAttribute{
									Description: "ID of the queue consumer configuration.",
									Computed:    true,
								},
								"queue_id": schema.StringAttribute{
									Description: "ID of the queue.",
									Computed:    true,
								},
								"queue_name": schema.StringAttribute{
									Description: "Name of the queue.",
									Computed:    true,
								},
							},
						},
					},
					"workers": schema.ListNestedAttribute{
						Description: "Other Workers that reference the Worker using [service bindings](https://developers.cloudflare.com/workers/runtime-apis/bindings/service-bindings/).",
						Computed:    true,
						CustomType:  customfield.NewNestedObjectListType[WorkerReferencesWorkersModel](ctx),
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description: "ID of the referencing Worker.",
									Computed:    true,
								},
								"name": schema.StringAttribute{
									Description: "Name of the referencing Worker.",
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

func (r *WorkerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *WorkerResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
