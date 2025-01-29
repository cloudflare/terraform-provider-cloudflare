// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
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
				CustomType:  customfield.NewNestedObjectListType[CloudforceOneRequestsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID",
							Computed:    true,
						},
						"created": schema.StringAttribute{
							Description: "Request creation time",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"priority": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"routine",
									"high",
									"urgent",
								),
							},
						},
						"request": schema.StringAttribute{
							Description: "Requested information from request",
							Computed:    true,
						},
						"summary": schema.StringAttribute{
							Description: "Brief description of the request",
							Computed:    true,
						},
						"tlp": schema.StringAttribute{
							Description: "The CISA defined Traffic Light Protocol (TLP)",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"clear",
									"amber",
									"amber-strict",
									"green",
									"red",
								),
							},
						},
						"updated": schema.StringAttribute{
							Description: "Request last updated time",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"completed": schema.StringAttribute{
							Description: "Request completion time",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"message_tokens": schema.Int64Attribute{
							Description: "Tokens for the request messages",
							Computed:    true,
						},
						"readable_id": schema.StringAttribute{
							Description: "Readable Request ID",
							Computed:    true,
						},
						"status": schema.StringAttribute{
							Description: "Request Status",
							Computed:    true,
							Validators: []validator.String{
								stringvalidator.OneOfCaseInsensitive(
									"open",
									"accepted",
									"reported",
									"approved",
									"completed",
									"declined",
								),
							},
						},
						"tokens": schema.Int64Attribute{
							Description: "Tokens for the request",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *CloudforceOneRequestsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *CloudforceOneRequestsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
