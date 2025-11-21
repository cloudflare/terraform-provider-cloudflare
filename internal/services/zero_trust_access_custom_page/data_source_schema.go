// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_access_custom_page

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustAccessCustomPageDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "UUID.",
				Computed:    true,
			},
			"custom_page_id": schema.StringAttribute{
				Description: "UUID.",
				Required:    true,
			},
			"account_id": schema.StringAttribute{
				Description: "Identifier.",
				Required:    true,
			},
			"custom_html": schema.StringAttribute{
				Description: "Custom page HTML.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Custom page name.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "Custom page type.\nAvailable values: \"identity_denied\", \"forbidden\".",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("identity_denied", "forbidden"),
				},
			},
			"uid": schema.StringAttribute{
				Description: "UUID.",
				Computed:    true,
			},
		},
	}
}

func (d *ZeroTrustAccessCustomPageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustAccessCustomPageDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
