// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_role

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/v5/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountRolesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Required:    true,
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
				CustomType:  customfield.NewNestedObjectListType[AccountRolesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Role identifier tag.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Description of role's permissions.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Role name.",
							Computed:    true,
						},
						"permissions": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"analytics": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsAnalyticsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"billing": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsBillingDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"cache_purge": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsCachePurgeDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"dns": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsDNSDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"dns_records": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsDNSRecordsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"lb": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsLBDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"logs": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsLogsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"organization": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsOrganizationDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"ssl": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsSSLDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"waf": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsWAFDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"zone_settings": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsZoneSettingsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
									},
								},
								"zones": schema.SingleNestedAttribute{
									Computed:   true,
									CustomType: customfield.NewNestedObjectType[AccountRolesPermissionsZonesDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"read": schema.BoolAttribute{
											Computed: true,
										},
										"write": schema.BoolAttribute{
											Computed: true,
										},
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

func (d *AccountRolesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountRolesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
