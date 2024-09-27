// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicNetworkMonitoringConfigurationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Description: "The account name.",
				Optional:    true,
			},
			"router_ips": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"warp_devices": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the warp device.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the warp device.",
							Computed:    true,
						},
						"router_ip": schema.StringAttribute{
							Description: "IPv4 CIDR of the router sourcing flow data associated with this warp device. Only /32 addresses are currently supported.",
							Computed:    true,
						},
					},
				},
			},
			"default_sampling": schema.Float64Attribute{
				Description: "Fallback sampling rate of flow messages being sent in packets per second. This should match the packet sampling rate configured on the router.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
			},
		},
	}
}

func (d *MagicNetworkMonitoringConfigurationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *MagicNetworkMonitoringConfigurationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
