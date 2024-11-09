// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package permission_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*PermissionGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"permission_group_id": schema.StringAttribute{
				Description: "Permission Group identifier tag.",
				Optional:    true,
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
					"id": schema.StringAttribute{
						Description: "ID of the permission group to be fetched.",
						Optional:    true,
					},
					"label": schema.StringAttribute{
						Description: "Label of the permission group to be fetched.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the permission group to be fetched.",
						Optional:    true,
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
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("permission_group_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("permission_group_id")),
	}
}
