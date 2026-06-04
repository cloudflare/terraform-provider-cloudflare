// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_deployment_groups

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceDeploymentGroupsDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"group_id": schema.StringAttribute{
				Required: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Description: "The RFC3339Nano timestamp when the deployment group was created.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "A user-friendly name for the deployment group.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The RFC3339Nano timestamp when the deployment group was last updated.",
				Computed:    true,
			},
			"policy_ids": schema.ListAttribute{
				Description: "Contains a list of policy IDs assigned to this deployment group.",
				Computed:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"version_config": schema.ListNestedAttribute{
				Description: "Contains version configurations for different target environments.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceDeploymentGroupsVersionConfigDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"target_environment": schema.StringAttribute{
							Description: "The target environment for the client version (e.g., windows, macos).",
							Computed:    true,
						},
						"version": schema.StringAttribute{
							Description: "The specific client version to deploy.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceDeploymentGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceDeploymentGroupsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
