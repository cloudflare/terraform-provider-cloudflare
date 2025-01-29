// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*PermissionGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"permission_group_id": schema.StringAttribute{
				Description: "Permission Group identifier tag.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier of the group.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the group.",
				Computed:    true,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Attributes associated to the permission group.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[PermissionGroupMetaDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Computed: true,
					},
					"value": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *PermissionGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *PermissionGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
