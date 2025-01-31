// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitSiteLANResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"site_id": schema.StringAttribute{
				Description:   "Identifier",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"lan_id": schema.StringAttribute{
				Description:   "Identifier",
				Optional:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ha_link": schema.BoolAttribute{
				Description:   "mark true to use this LAN for HA probing. only works for site with HA turned on. only one LAN can be set as the ha_link.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"physport": schema.Int64Attribute{
				Required: true,
			},
			"vlan_tag": schema.Int64Attribute{
				Description: "VLAN port number.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"nat": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANNatModel](ctx),
				Attributes: map[string]schema.Attribute{
					"static_prefix": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Optional:    true,
					},
				},
			},
			"routed_subnets": schema.ListNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectListType[MagicTransitSiteLANRoutedSubnetsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"next_hop": schema.StringAttribute{
							Description: "A valid IPv4 address.",
							Required:    true,
						},
						"prefix": schema.StringAttribute{
							Description: "A valid CIDR notation representing an IP range.",
							Required:    true,
						},
						"nat": schema.SingleNestedAttribute{
							Computed:   true,
							Optional:   true,
							CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANRoutedSubnetsNatModel](ctx),
							Attributes: map[string]schema.Attribute{
								"static_prefix": schema.StringAttribute{
									Description: "A valid CIDR notation representing an IP range.",
									Optional:    true,
								},
							},
						},
					},
				},
			},
			"static_addressing": schema.SingleNestedAttribute{
				Description: "If the site is not configured in high availability mode, this configuration is optional (if omitted, use DHCP). However, if in high availability mode, static_address is required along with secondary and virtual address.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[MagicTransitSiteLANStaticAddressingModel](ctx),
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Required:    true,
					},
					"dhcp_relay": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANStaticAddressingDHCPRelayModel](ctx),
						Attributes: map[string]schema.Attribute{
							"server_addresses": schema.ListAttribute{
								Description: "List of DHCP server IPs.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"dhcp_server": schema.SingleNestedAttribute{
						Computed:   true,
						Optional:   true,
						CustomType: customfield.NewNestedObjectType[MagicTransitSiteLANStaticAddressingDHCPServerModel](ctx),
						Attributes: map[string]schema.Attribute{
							"dhcp_pool_end": schema.StringAttribute{
								Description: "A valid IPv4 address.",
								Optional:    true,
							},
							"dhcp_pool_start": schema.StringAttribute{
								Description: "A valid IPv4 address.",
								Optional:    true,
							},
							"dns_server": schema.StringAttribute{
								Description: "A valid IPv4 address.",
								Optional:    true,
							},
							"dns_servers": schema.ListAttribute{
								Optional:    true,
								ElementType: types.StringType,
							},
							"reservations": schema.MapAttribute{
								Description: "Mapping of MAC addresses to IP addresses",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"secondary_address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Optional:    true,
					},
					"virtual_address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Optional:    true,
					},
				},
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
		},
	}
}

func (r *MagicTransitSiteLANResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitSiteLANResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
