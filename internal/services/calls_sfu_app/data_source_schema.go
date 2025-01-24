// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls_sfu_app

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*CallsSFUAppDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account identifier tag.",
				Required:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a item.",
				Required:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the item was created.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the item was last modified.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"name": schema.StringAttribute{
				Description: "A short description of Calls app, not shown to end users.",
				Computed:    true,
			},
			"uid": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a item.",
				Computed:    true,
			},
		},
	}
}

func (d *CallsSFUAppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CallsSFUAppDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
