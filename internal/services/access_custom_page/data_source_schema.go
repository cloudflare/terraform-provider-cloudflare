// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_custom_page

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &AccessCustomPageDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessCustomPageDataSource{}

func (r AccessCustomPageDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"custom_page_id": schema.StringAttribute{
				Description: "UUID",
				Optional:    true,
			},
			"custom_html": schema.StringAttribute{
				Description: "Custom page HTML.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Custom page name.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "Custom page type.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("identity_denied", "forbidden"),
				},
			},
			"app_count": schema.Int64Attribute{
				Description: "Number of apps the custom page is assigned to.",
				Computed:    true,
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"uid": schema.StringAttribute{
				Description: "UUID",
				Computed:    true,
				Optional:    true,
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *AccessCustomPageDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessCustomPageDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
