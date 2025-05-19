// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID.",
				Computed:    true,
			},
			"request_id": schema.StringAttribute{
				Description: "UUID.",
				Optional:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"completed": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"content": schema.StringAttribute{
				Description: "Request content.",
				Computed:    true,
			},
			"created": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"message_tokens": schema.Int64Attribute{
				Description: "Tokens for the request messages.",
				Computed:    true,
			},
			"priority": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"readable_id": schema.StringAttribute{
				Description: "Readable Request ID.",
				Computed:    true,
			},
			"request": schema.StringAttribute{
				Description: "Requested information from request.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "Request Status.\nAvailable values: \"open\", \"accepted\", \"reported\", \"approved\", \"completed\", \"declined\".",
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
			"summary": schema.StringAttribute{
				Description: "Brief description of the request.",
				Computed:    true,
			},
			"tlp": schema.StringAttribute{
				Description: "The CISA defined Traffic Light Protocol (TLP).\nAvailable values: \"clear\", \"amber\", \"amber-strict\", \"green\", \"red\".",
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
			"tokens": schema.Int64Attribute{
				Description: "Tokens for the request.",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"page": schema.Int64Attribute{
						Description: "Page number of results.",
						Required:    true,
					},
					"per_page": schema.Int64Attribute{
						Description: "Number of results per page.",
						Required:    true,
					},
					"completed_after": schema.StringAttribute{
						Description: "Retrieve requests completed after this time.",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"completed_before": schema.StringAttribute{
						Description: "Retrieve requests completed before this time.",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"created_after": schema.StringAttribute{
						Description: "Retrieve requests created after this time.",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"created_before": schema.StringAttribute{
						Description: "Retrieve requests created before this time.",
						Optional:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"request_type": schema.StringAttribute{
						Description: "Requested information from request.",
						Optional:    true,
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
					"status": schema.StringAttribute{
						Description: "Request Status.\nAvailable values: \"open\", \"accepted\", \"reported\", \"approved\", \"completed\", \"declined\".",
						Optional:    true,
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
				},
			},
		},
	}
}

func (d *CloudforceOneRequestDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CloudforceOneRequestDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("request_id"), path.MatchRoot("filter")),
	}
}
