// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one_request

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customvalidator"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var _ datasource.DataSourceWithConfigValidators = (*CloudforceOneRequestDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"request_identifier": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"content": schema.StringAttribute{
				Description: "Request content",
				Optional:    true,
			},
			"completed": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"created": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"id": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
			},
			"message_tokens": schema.Int64Attribute{
				Description: "Tokens for the request messages",
				Computed:    true,
			},
			"readable_id": schema.StringAttribute{
				Description: "Readable Request ID",
				Computed:    true,
			},
			"request": schema.StringAttribute{
				Description: "Requested information from request",
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
			"tokens": schema.Int64Attribute{
				Description: "Tokens for the request",
				Computed:    true,
			},
			"updated": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"priority": schema.DynamicAttribute{
				Computed: true,
				Validators: []validator.Dynamic{
					customvalidator.AllowedSubtypes(basetypes.StringType{}, basetypes.StringType{}),
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_identifier": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
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
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_identifier"), path.MatchRoot("request_identifier")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_identifier")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("request_identifier")),
	}
}
