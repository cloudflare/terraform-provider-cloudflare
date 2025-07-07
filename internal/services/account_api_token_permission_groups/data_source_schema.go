// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_api_token_permission_groups

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountAPITokenPermissionGroupsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "Filter by the name of the permission group.\nThe value must be URL-encoded.",
				Optional:    true,
			},
			"scope": schema.StringAttribute{
				Description: "Filter by the scope of the permission group.\nThe value must be URL-encoded.",
				Optional:    true,
			},
		},
	}
}

func (d *AccountAPITokenPermissionGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccountAPITokenPermissionGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
