// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_group

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = &AccessGroupsDataSource{}
var _ datasource.DataSourceWithValidateConfig = &AccessGroupsDataSource{}

func (r AccessGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
			},
			"items": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "UUID",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed: true,
						},
						"exclude": schema.ListNestedAttribute{
							Description: "Rules evaluated with a NOT logical operator. To match a policy, a user cannot meet any of the Exclude rules.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
									},
								},
							},
						},
						"include": schema.ListNestedAttribute{
							Description: "Rules evaluated with an OR logical operator. A user needs to meet only one of the Include rules.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
									},
								},
							},
						},
						"is_default": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the Access group.",
							Computed:    true,
						},
						"require": schema.ListNestedAttribute{
							Description: "Rules evaluated with an AND logical operator. To match a policy, a user must meet all of the Require rules.",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"everyone": schema.StringAttribute{
										Description: "An empty object which matches on all users.",
										Computed:    true,
									},
									"certificate": schema.StringAttribute{
										Computed: true,
									},
									"any_valid_service_token": schema.StringAttribute{
										Description: "An empty object which matches on all service tokens.",
										Computed:    true,
									},
								},
							},
						},
						"updated_at": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (r *AccessGroupsDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *AccessGroupsDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
