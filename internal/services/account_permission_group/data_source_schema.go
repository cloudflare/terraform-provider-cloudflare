// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account_permission_group

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountPermissionGroupDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Account Firewall Access Rules Read",
				"Account Firewall Access Rules Write",
				"Account Settings Read",
				"Account Settings Write",
				"Billing Read",
				"Billing Write",
				"DDoS Botnet Feed Read",
				"DDoS Botnet Feed Write",
				"DDoS Protection Read",
				"DDoS Protection Write",
				"DNS Firewall Read",
				"DNS Firewall Write",
				"DNS View Read",
				"DNS View Write",
				"Load Balancers Account Read",
				"Load Balancers Account Write",
				"Load Balancing: Monitors and Pools Read",
				"Load Balancing: Monitors and Pools Write",
				"SCIM Provisioning",
				"Trust and Safety Read",
				"Trust and Safety Write",
				"Workers KV Storage Read",
				"Workers KV Storage Write",
				"Workers R2 Storage Read",
				"Workers R2 Storage Write",
				"Workers Scripts Read",
				"Workers Scripts Write",
				"Workers Tail Read",
				"Zero Trust: PII Read",
			},
		}.String(),
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
				Description: "Identifier of the permission group.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the permission group.",
				Computed:    true,
			},
			"meta": schema.SingleNestedAttribute{
				Description: "Attributes associated to the permission group.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectType[AccountPermissionGroupMetaDataSourceModel](ctx),
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

func (d *AccountPermissionGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *AccountPermissionGroupDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
