// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestMessageDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"request_id": schema.StringAttribute{
				Description: "UUID.",
				Required:    true,
			},
			"author": schema.StringAttribute{
				Description: "Author of message.",
				Computed:    true,
			},
			"content": schema.StringAttribute{
				Description: "Content of message.",
				Computed:    true,
			},
			"created": schema.StringAttribute{
				Description: "Defines the message creation time.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"id": schema.Int64Attribute{
				Description: "Message ID.",
				Computed:    true,
			},
			"is_follow_on_request": schema.BoolAttribute{
				Description: "Whether the message is a follow-on request.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Description: "Defines the message last updated time.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
		},
	}
}

func (d *CloudforceOneRequestMessageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CloudforceOneRequestMessageDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
