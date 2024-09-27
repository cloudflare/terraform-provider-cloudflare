// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*CallAppDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account identifier tag.",
				Optional:    true,
			},
			"app_id": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a item.",
				Optional:    true,
			},
			"created": schema.StringAttribute{
				Description: "The date and time the item was created.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"modified": schema.StringAttribute{
				Description: "The date and time the item was last modified.",
				Optional:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"uid": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a item.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A short description of Calls app, not shown to end users.",
				Computed:    true,
				Optional:    true,
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The account identifier tag.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *CallAppDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CallAppDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("app_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("app_id")),
	}
}
