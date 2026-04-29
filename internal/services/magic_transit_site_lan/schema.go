// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site_lan

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/schemata"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*MagicTransitSiteLANResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		MarkdownDescription: schemata.Description{
			Scopes: []string{
				"Magic Transit Read",
				"Magic Transit Write",
				"Magic WAN Read",
				"Magic WAN Write",
			},
		}.String(),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identifier",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
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
			"ha_link": schema.BoolAttribute{
				Description:   "mark true to use this LAN for HA probing. only works for site with HA turned on. only one LAN can be set as the ha_link.",
				Optional:      true,
				PlanModifiers: []planmodifier.Bool{boolplanmodifier.RequiresReplace()},
			},
			"bond_id": schema.Int64Attribute{
				Optional: true,
			},
			"is_breakout": schema.BoolAttribute{
				Description: "mark true to use this LAN for source-based breakout traffic",
				Optional:    true,
			},
			"is_prioritized": schema.BoolAttribute{
				Description: "mark true to use this LAN for source-based prioritized traffic",
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Optional: true,
			},
			"physport": schema.Int64Attribute{
				Optional: true,
			},
			"vlan_tag": schema.Int64Attribute{
				Description: "VLAN ID. Use zero for untagged.",
				Optional:    true,
			},
			"nat": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"static_prefix": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Optional:    true,
					},
				},
			},
			"routed_subnets": schema.ListNestedAttribute{
				Optional: true,
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
							Optional: true,
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
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"address": schema.StringAttribute{
						Description: "A valid CIDR notation representing an IP range.",
						Required:    true,
					},
					"dhcp_relay": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"server_addresses": schema.ListAttribute{
								Description: "List of DHCP server IPs.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"dhcp_server": schema.SingleNestedAttribute{
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"dhcp_options": schema.ListNestedAttribute{
								Description: "Optional list of custom DHCP options to include in DHCP responses. Only valid when DHCP server is enabled.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"code": schema.Int64Attribute{
											Description: "DHCP option number (1-254). Options 0 and 255 are reserved by RFC 2132. Options 3, 6, and 51 are not allowed because they conflict with connector-managed configuration.",
											Required:    true,
											Validators: []validator.Int64{
												int64validator.Between(1, 254),
											},
										},
										"type": schema.StringAttribute{
											Description: "The type of the option value. text: a string (max 255 bytes). hex: colon-separated hex bytes (e.g. \"01:04:aa:bb:cc\", max 255 bytes). ip: an IPv4 address (e.g. \"10.20.30.40\"). byte: an unsigned integer 0-255 (1 byte). short: an unsigned integer 0-65535 (2 bytes). integer: an unsigned integer 0-4294967295 (4 bytes).\nAvailable values: \"text\", \"hex\", \"ip\", \"byte\", \"short\", \"integer\".",
											Required:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"text",
													"hex",
													"ip",
													"byte",
													"short",
													"integer",
												),
											},
										},
										"value": schema.StringAttribute{
											Description: "The option value, interpreted according to the type field.",
											Required:    true,
										},
									},
								},
							},
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
		},
	}
}

func (r *MagicTransitSiteLANResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *MagicTransitSiteLANResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
