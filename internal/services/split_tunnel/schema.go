package split_tunnel

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ResourceSchema returns the Terraform schema for the deprecated cloudflare_split_tunnel resource.
func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown(), stringplanmodifier.RequiresReplace()},
			},
			"lan_allow_minutes": schema.Float64Attribute{
				Description: "The amount of time in minutes a user is allowed access to their LAN. A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.",
				Optional:    true,
			},
			"lan_allow_subnet_size": schema.Float64Attribute{
				Description: "The size of the subnet for the local access network. Note that this field is omitted from the response if null or unset.",
				Optional:    true,
			},
			"exclude": schema.ListNestedAttribute{
				Description: "List of routes excluded in the WARP client's tunnel. Both 'exclude' and 'include' cannot be set in the same request.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ConflictsWith(path.MatchRoot("include")),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.",
							Optional:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the Split Tunnel item, displayed in the client UI.",
							Optional:    true,
						},
						"host": schema.StringAttribute{
							Description: "The domain name to exclude from the tunnel. If `host` is present, `address` must not be present.",
							Optional:    true,
						},
					},
				},
			},
			"include": schema.ListNestedAttribute{
				Description: "List of routes included in the WARP client's tunnel. Both 'exclude' and 'include' cannot be set in the same request.",
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ConflictsWith(path.MatchRoot("exclude")),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The address in CIDR format to include in the tunnel. If `address` is present, `host` must not be present.",
							Optional:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the Split Tunnel item, displayed in the client UI.",
							Optional:    true,
						},
						"host": schema.StringAttribute{
							Description: "The domain name to include in the tunnel. If `host` is present, `address` must not be present.",
							Optional:    true,
						},
					},
				},
			},
			"service_mode_v2": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "The mode to run the WARP client under.",
						Optional:    true,
					},
					"port": schema.Float64Attribute{
						Description: "The port number when used with proxy mode.",
						Optional:    true,
					},
				},
			},
			"allow_mode_switch": schema.BoolAttribute{
				Description: "Whether to allow mode switch for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"allow_updates": schema.BoolAttribute{
				Description: "Whether to allow updates for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"allowed_to_leave": schema.BoolAttribute{
				Description: "Whether the profile is allowed to leave the organization.",
				Optional:    true,
				Computed:    true,
			},
			"auto_connect": schema.Float64Attribute{
				Description: "The amount of time in minutes to reconnect after a disconnection.",
				Optional:    true,
				Computed:    true,
			},
			"captive_portal": schema.Float64Attribute{
				Description: "The captive portal setting for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"disable_auto_fallback": schema.BoolAttribute{
				Description: "Whether to disable auto fallback for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"exclude_office_ips": schema.BoolAttribute{
				Description: "Whether to exclude office IPs from this profile.",
				Optional:    true,
				Computed:    true,
			},
			"register_interface_ip_with_dns": schema.BoolAttribute{
				Description: "Whether to register interface IP with DNS.",
				Optional:    true,
				Computed:    true,
			},
			"sccm_vpn_boundary_support": schema.BoolAttribute{
				Description: "Whether to enable SCCM VPN boundary support.",
				Optional:    true,
				Computed:    true,
			},
			"support_url": schema.StringAttribute{
				Description: "The support URL for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"switch_locked": schema.BoolAttribute{
				Description: "Whether the switch is locked for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"tunnel_protocol": schema.StringAttribute{
				Description: "The tunnel protocol for this profile.",
				Optional:    true,
				Computed:    true,
			},
			"default": schema.BoolAttribute{
				Description: "Whether this is the default profile.",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether this profile is enabled.",
				Computed:    true,
			},
			"gateway_unique_id": schema.StringAttribute{
				Description: "The unique ID of the gateway for this profile.",
				Computed:    true,
			},
			"fallback_domains": schema.ListNestedAttribute{
				Description: "The fallback domains for this profile.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"suffix": schema.StringAttribute{
							Description: "The domain suffix to match when resolving locally.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the fallback domain, displayed in the client UI.",
							Computed:    true,
						},
						"dns_server": schema.ListAttribute{
							Description: "A list of IP addresses to handle domain resolution.",
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

// Duplicate Schema and ConfigValidators implementations were removed to avoid
// redeclaration errors. The implementations reside in `resource.go`.
