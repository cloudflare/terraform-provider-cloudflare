// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package call_app_turn_key

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*CallAppTURNKeyDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account identifier tag.",
				Optional:    true,
			},
			"key_id": schema.StringAttribute{
				Description: "A Cloudflare-generated unique identifier for a item.",
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

func (d *CallAppTURNKeyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *CallAppTURNKeyDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("key_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("key_id")),
	}
}
