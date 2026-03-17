// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package worker

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*WorkersDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
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
				CustomType:  customfield.NewNestedObjectListType[WorkersResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Immutable ID of the Worker.",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "When the Worker was created.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"logpush": schema.BoolAttribute{
							Description: "Whether logpush is enabled for the Worker.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the Worker.",
							Computed:    true,
						},
						"observability": schema.SingleNestedAttribute{
							Description: "Observability settings for the Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersObservabilityDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether observability is enabled for the Worker.",
									Computed:    true,
								},
								"head_sampling_rate": schema.Float64Attribute{
									Description: "The sampling rate for observability. From 0 to 1 (1 = 100%, 0.1 = 10%).",
									Computed:    true,
								},
								"logs": schema.SingleNestedAttribute{
									Description: "Log settings for the Worker.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[WorkersObservabilityLogsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description: "Whether logs are enabled for the Worker.",
											Computed:    true,
										},
										"head_sampling_rate": schema.Float64Attribute{
											Description: "The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%).",
											Computed:    true,
										},
										"invocation_logs": schema.BoolAttribute{
											Description: "Whether [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs) are enabled for the Worker.",
											Computed:    true,
										},
									},
								},
							},
						},
						"references": schema.SingleNestedAttribute{
							Description: "Other resources that reference the Worker and depend on it existing.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersReferencesDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"dispatch_namespace_outbounds": schema.ListNestedAttribute{
									Description: "Other Workers that reference the Worker as an outbound for a dispatch namespace.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectListType[WorkersReferencesDispatchNamespaceOutboundsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[WorkersReferencesDomainsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[WorkersReferencesDurableObjectsDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[WorkersReferencesQueuesDataSourceModel](ctx),
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
									CustomType:  customfield.NewNestedObjectListType[WorkersReferencesWorkersDataSourceModel](ctx),
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
						"subdomain": schema.SingleNestedAttribute{
							Description: "Subdomain settings for the Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[WorkersSubdomainDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description: "Whether the *.workers.dev subdomain is enabled for the Worker.",
									Computed:    true,
								},
								"previews_enabled": schema.BoolAttribute{
									Description: "Whether [preview URLs](https://developers.cloudflare.com/workers/configuration/previews/) are enabled for the Worker.",
									Computed:    true,
								},
							},
						},
						"tags": schema.SetAttribute{
							Description: "Tags associated with the Worker.",
							Computed:    true,
							CustomType:  customfield.NewSetType[types.String](ctx),
							ElementType: types.StringType,
						},
						"tail_consumers": schema.SetNestedAttribute{
							Description: "Other Workers that should consume logs from the Worker.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectSetType[WorkersTailConsumersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description: "Name of the consumer Worker.",
										Computed:    true,
									},
								},
							},
						},
						"updated_on": schema.StringAttribute{
							Description: "When the Worker was most recently updated.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
					},
				},
			},
		},
	}
}

func (d *WorkersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *WorkersDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
