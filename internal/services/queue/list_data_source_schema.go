// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*QueuesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
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
									"created_on": schema.StringAttribute{
										Computed: true,
									},
									"environment": schema.StringAttribute{
										Computed: true,
									},
									"queue_name": schema.StringAttribute{
										Computed: true,
									},
									"service": schema.StringAttribute{
										Computed: true,
									},
									"settings": schema.SingleNestedAttribute{
										Computed:   true,
										CustomType: customfield.NewNestedObjectType[QueuesConsumersSettingsDataSourceModel](ctx),
										Attributes: map[string]schema.Attribute{
											"batch_size": schema.Float64Attribute{
												Description: "The maximum number of messages to include in a batch.",
												Computed:    true,
											},
											"max_retries": schema.Float64Attribute{
												Description: "The maximum number of retries",
												Computed:    true,
											},
											"max_wait_time_ms": schema.Float64Attribute{
												Computed: true,
											},
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
									"environment": schema.StringAttribute{
										Computed: true,
									},
									"service": schema.StringAttribute{
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
