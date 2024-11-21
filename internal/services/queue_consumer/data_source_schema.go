// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package queue_consumer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*QueueConsumerDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "A Resource identifier.",
				Required:    true,
			},
			"queue_id": schema.StringAttribute{
				Description: "A Resource identifier.",
				Required:    true,
			},
		},
	}
}

func (d *QueueConsumerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *QueueConsumerDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
