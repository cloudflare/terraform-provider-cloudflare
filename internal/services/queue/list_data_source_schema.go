// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*QueuesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "A Resource identifier.",
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
				CustomType:  customfield.NewNestedObjectListType[QueuesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"consumers": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[QueuesConsumersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"consumer_id": schema.StringAttribute{
										Description: "A Resource identifier.",
										Computed:    true,
									},
									"created_on": schema.StringAttribute{
										Computed: true,
									},
									"queue_id": schema.StringAttribute{
										Description: "A Resource identifier.",
										Computed:    true,
									},
									"script": schema.StringAttribute{
										Description: "Name of a Worker",
										Computed:    true,
									},
									"script_name": schema.StringAttribute{
										Description: "Name of a Worker",
										Computed:    true,
									},
									"settings": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[QueuesConsumersSettingsDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"batch_size": schema.Float64Attribute{
												Description: "The maximum number of messages to include in a batch.",
												Computed:    true,
											},
											"max_concurrency": schema.Float64Attribute{
												Description: "Maximum number of concurrent consumers that may consume from this Queue. Set to `null` to automatically opt in to the platform's maximum (recommended).",
												Computed:    true,
											},
											"max_retries": schema.Float64Attribute{
												Description: "The maximum number of retries",
												Computed:    true,
											},
											"max_wait_time_ms": schema.Float64Attribute{
												Description: "The number of milliseconds to wait for a batch to fill up before attempting to deliver it",
												Computed:    true,
											},
											"retry_delay": schema.Float64Attribute{
												Description: "The number of seconds to delay before making the message available for another attempt.",
												Computed:    true,
											},
											"visibility_timeout_ms": schema.Float64Attribute{
												Description: "The number of milliseconds that a message is exclusively leased. After the timeout, the message becomes available for another attempt.",
												Computed:    true,
											},
										},
									},
									"type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("worker", "http_pull"),
										},
									},
								},
							},
						},
						"consumers_total_count": schema.Float64Attribute{
							Computed: true,
						},
						"created_on": schema.StringAttribute{
							Computed: true,
						},
						"modified_on": schema.StringAttribute{
							Computed: true,
						},
						"producers": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[QueuesProducersDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"script": schema.StringAttribute{
										Computed: true,
									},
									"type": schema.StringAttribute{
										Computed: true,
										Validators: []validator.String{
											stringvalidator.OneOfCaseInsensitive("worker", "r2_bucket"),
										},
									},
									"bucket_name": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"producers_total_count": schema.Float64Attribute{
							Computed: true,
						},
						"queue_id": schema.StringAttribute{
							Computed: true,
						},
						"queue_name": schema.StringAttribute{
							Computed: true,
						},
						"settings": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[QueuesSettingsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"delivery_delay": schema.Float64Attribute{
									Description: "Number of seconds to delay delivery of all messages to consumers.",
									Computed:    true,
								},
								"message_retention_period": schema.Float64Attribute{
									Description: "Number of seconds after which an unconsumed message will be delayed.",
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

func (d *QueuesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *QueuesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
