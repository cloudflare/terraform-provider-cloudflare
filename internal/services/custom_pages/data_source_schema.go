// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_pages

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*CustomPagesDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Error Page Types\nAvailable values: \"1000_errors\", \"500_errors\", \"basic_challenge\", \"country_challenge\", \"ip_block\", \"managed_challenge\", \"ratelimit_block\", \"under_attack\", \"waf_block\", \"waf_challenge\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"1000_errors",
						"500_errors",
						"basic_challenge",
						"country_challenge",
						"ip_block",
						"managed_challenge",
						"ratelimit_block",
						"under_attack",
						"waf_block",
						"waf_challenge",
					),
				},
			},
			"identifier": schema.StringAttribute{
				Description: "Error Page Types\nAvailable values: \"1000_errors\", \"500_errors\", \"basic_challenge\", \"country_challenge\", \"ip_block\", \"managed_challenge\", \"ratelimit_block\", \"under_attack\", \"waf_block\", \"waf_challenge\".",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"1000_errors",
						"500_errors",
						"basic_challenge",
						"country_challenge",
						"ip_block",
						"managed_challenge",
						"ratelimit_block",
						"under_attack",
						"waf_block",
						"waf_challenge",
					),
				},
			},
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"preview_target": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Description: "The custom page state.\nAvailable values: \"default\", \"customized\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("default", "customized"),
				},
			},
			"url": schema.StringAttribute{
				Description: "The URL associated with the custom page.",
				Computed:    true,
			},
			"required_tokens": schema.ListAttribute{
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
		},
	}
}

func (d *CustomPagesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CustomPagesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.Conflicting(path.MatchRoot("account_id"), path.MatchRoot("zone_id")),
	}
}
