// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_permission_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountPermissionGroupsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AccountPermissionGroupsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier of the permission group.",
							Computed:    true,
						},
						"meta": schema.SingleNestedAttribute{
							Description: "Attributes associated to the permission group.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[AccountPermissionGroupsMetaDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Computed: true,
								},
								"value": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "Name of the permission group.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *AccountPermissionGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountPermissionGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
