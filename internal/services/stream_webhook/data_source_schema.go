// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*StreamWebhookDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account identifier tag.",
				Required:    true,
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the webhook was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"notification_url": schema.StringAttribute{
				Description: "The URL where webhooks will be sent.",
				Computed:    true,
			},
			"secret": schema.StringAttribute{
				Description: "The secret used to verify webhook signatures.",
				Computed:    true,
				Sensitive:   true,
			},
		},
	}
}

func (d *StreamWebhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *StreamWebhookDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
