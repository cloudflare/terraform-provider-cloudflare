// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request_message

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
			"page": schema.Int64Attribute{
				Description: "Page number of results.",
				Required:    true,
			},
			"per_page": schema.Int64Attribute{
				Description: "Number of results per page.",
				Required:    true,
			},
			"after": schema.StringAttribute{
				Description: "Retrieve mes  ges created after this time.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"before": schema.StringAttribute{
				Description: "Retrieve messages created before this time.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"sort_by": schema.StringAttribute{
				Description: "Field to sort results by.",
				Optional:    true,
			},
			"sort_order": schema.StringAttribute{
				Description: "Sort order (asc or desc).\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
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
