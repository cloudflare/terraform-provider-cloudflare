// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring_configuration

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicNetworkMonitoringConfigurationResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"name": schema.StringAttribute{
				Description: "The account name.",
				Required:    true,
			},
			"default_sampling": schema.Float64Attribute{
				Description: "Fallback sampling rate of flow messages being sent in packets per second. This should match the packet sampling rate configured on the router.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.Float64{
					float64validator.AtLeast(1),
				},
				Default: float64default.StaticFloat64(1),
			},
			"router_ips": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"warp_devices": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[MagicNetworkMonitoringConfigurationWARPDevicesModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Unique identifier for the warp device.",
							Required:    true,
						},
						"name": schema.StringAttribute{
							Description: "Name of the warp device.",
							Required:    true,
						},
						"router_ip": schema.StringAttribute{
							Description: "IPv4 CIDR of the router sourcing flow data associated with this warp device. Only /32 addresses are currently supported.",
							Required:    true,
						},
					},
				},
			},
		},
	}
}

func (r *MagicNetworkMonitoringConfigurationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicNetworkMonitoringConfigurationResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
