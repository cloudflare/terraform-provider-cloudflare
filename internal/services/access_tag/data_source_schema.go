// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_tag

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessTagDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessTagDataSource{}

func (r AccessTagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"tag_name": schema.StringAttribute{
				Description: "The name of the tag",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the tag",
				Optional:    true,
			},
			"app_count": schema.Int64Attribute{
				Description: "The number of applications that have this tag",
				Optional:    true,
			},
			"created_at": schema.StringAttribute{
				Optional: true,
			},
			"updated_at": schema.StringAttribute{
				Optional: true,
			},
			"find_one_by": schema.SingleNestedAttribute{
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

func (r *AccessTagDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessTagDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
