// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &QueuesDataSource{}

func (d *QueuesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"consumers": schema.ListNestedAttribute{
							Computed: true,
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
										Computed: true,
										Optional: true,
										Attributes: map[string]schema.Attribute{
											"batch_size": schema.Float64Attribute{
												Description: "The maximum number of messages to include in a batch.",
												Computed:    true,
												Optional:    true,
											},
											"max_retries": schema.Float64Attribute{
												Description: "The maximum number of retries",
												Computed:    true,
												Optional:    true,
											},
											"max_wait_time_ms": schema.Float64Attribute{
												Computed: true,
												Optional: true,
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
						"producers": schema.ListAttribute{
							Computed:    true,
							ElementType: jsontypes.NewNormalizedNull().Type(ctx),
						},
						"producers_total_count": schema.Float64Attribute{
							Computed: true,
						},
						"queue_id": schema.StringAttribute{
							Computed: true,
						},
						"queue_name": schema.StringAttribute{
							Computed: true,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func (d *QueuesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
