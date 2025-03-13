// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*ResourceGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
			},
			"resource_group_id": schema.StringAttribute{
				Description: "Resource Group identifier tag.",
				Required:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier of the group.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the resource group.",
				Computed:    true,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Attributes associated to the resource group.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[ResourceGroupMetaDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Computed: true,
					},
					"value": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"scope": schema.ListNestedAttribute{
				Description: "The scope associated to the resource group",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ResourceGroupScopeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Description: "This is a combination of pre-defined resource name and identifier (like Account ID etc.)",
							Computed:    true,
						},
						"objects": schema.ListNestedAttribute{
							Description: "A list of scope objects for additional context.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ResourceGroupScopeObjectsDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description: "This is a combination of pre-defined resource name and identifier (like Zone ID etc.)",
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *ResourceGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ResourceGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
