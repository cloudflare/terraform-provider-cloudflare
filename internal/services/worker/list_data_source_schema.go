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
