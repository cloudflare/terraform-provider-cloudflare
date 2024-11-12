// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_role

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountRoleDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Account identifier tag.",
				Optional:    true,
			},
			"role_id": schema.StringAttribute{
				Description: "Role identifier tag.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "Description of role's permissions.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Role identifier tag.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Role name.",
				Computed:    true,
			},
			"permissions": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[AccountRolePermissionsDataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"analytics": schema.SingleNestedAttribute{
						Computed:   true,
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsAnalyticsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsBillingDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsCachePurgeDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsDNSDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsDNSRecordsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsLBDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsLogsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsOrganizationDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsSSLDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsWAFDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsZoneSettingsDataSourceModel](ctx),
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
						CustomType: customfield.NewNestedObjectType[AccountRolePermissionsZonesDataSourceModel](ctx),
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
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "Account identifier tag.",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *AccountRoleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccountRoleDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("role_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("role_id")),
	}
}
