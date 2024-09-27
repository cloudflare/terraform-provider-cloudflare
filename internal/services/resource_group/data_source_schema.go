// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ResourceGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"resource_group_id": schema.StringAttribute{
				Description: "Resource Group identifier tag.",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier of the group.",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the resource group.",
				Optional:    true,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Attributes associated to the resource group.",
				Optional:    true,
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
				Optional:    true,
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
					"id": schema.StringAttribute{
						Description: "ID of the resource group to be fetched.",
						Optional:    true,
					},
					"name": schema.StringAttribute{
						Description: "Name of the resource group to be fetched.",
						Optional:    true,
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
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("resource_group_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("resource_group_id")),
	}
}
