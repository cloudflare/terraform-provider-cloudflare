// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustDeviceCustomProfileResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Device ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"policy_id": schema.StringAttribute{
				Description:   "Device ID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"lan_allow_minutes": schema.Float64Attribute{
				Description:   "The amount of time in minutes a user is allowed access to their LAN. A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"lan_allow_subnet_size": schema.Float64Attribute{
				Description:   "The size of the subnet for the local access network. Note that this field is omitted from the response if null or unset.",
				Optional:      true,
				PlanModifiers: []planmodifier.Float64{float64planmodifier.RequiresReplace()},
			},
			"match": schema.StringAttribute{
				Description: "The wirefilter expression to match devices.",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the device settings profile.",
				Required:    true,
			},
			"precedence": schema.Float64Attribute{
				Description: "The precedence of the policy. Lower values indicate higher precedence. Policies will be evaluated in ascending order of this field.",
				Required:    true,
			},
			"allow_mode_switch": schema.BoolAttribute{
				Description: "Whether to allow the user to switch WARP between modes.",
				Optional:    true,
			},
			"allow_updates": schema.BoolAttribute{
				Description: "Whether to receive update notifications when a new version of the client is available.",
				Optional:    true,
			},
			"allowed_to_leave": schema.BoolAttribute{
				Description: "Whether to allow devices to leave the organization.",
				Optional:    true,
			},
			"auto_connect": schema.Float64Attribute{
				Description: "The amount of time in seconds to reconnect after having been disabled.",
				Optional:    true,
			},
			"captive_portal": schema.Float64Attribute{
				Description: "Turn on the captive portal after the specified amount of time.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the policy.",
				Optional:    true,
			},
			"disable_auto_fallback": schema.BoolAttribute{
				Description: "If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the policy will be applied to matching devices.",
				Optional:    true,
			},
			"exclude_office_ips": schema.BoolAttribute{
				Description: "Whether to add Microsoft IPs to Split Tunnel exclusions.",
				Optional:    true,
			},
			"register_interface_ip_with_dns": schema.BoolAttribute{
				Description: "Determines if the operating system will register WARP's local interface IP with your on-premises DNS server.",
				Optional:    true,
			},
			"support_url": schema.StringAttribute{
				Description: "The URL to launch when the Send Feedback button is clicked.",
				Optional:    true,
			},
			"switch_locked": schema.BoolAttribute{
				Description: "Whether to allow the user to turn off the WARP switch and disconnect the client.",
				Optional:    true,
			},
			"tunnel_protocol": schema.StringAttribute{
				Description: "Determines which tunnel protocol to use.",
				Optional:    true,
			},
			"exclude": schema.ListNestedAttribute{
				Description: "List of routes excluded in the WARP client's tunnel. Both 'exclude' and 'include' cannot be set in the same request.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfileExcludeModel](ctx),
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
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfileIncludeModel](ctx),
				Validators: []validator.List{
					listvalidator.ConflictsWith(path.MatchRoot("exclude")),
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
			"service_mode_v2": schema.SingleNestedAttribute{
				Computed:   true,
				Optional:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDeviceCustomProfileServiceModeV2Model](ctx),
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
			"default": schema.BoolAttribute{
				Description: "Whether the policy is the default policy for an account.",
				Computed:    true,
			},
			"gateway_unique_id": schema.StringAttribute{
				Computed: true,
			},
			"fallback_domains": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfileFallbackDomainsModel](ctx),
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
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
			"target_tests": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfileTargetTestsModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The id of the DEX test targeting this policy",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the DEX test targeting this policy",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (r *ZeroTrustDeviceCustomProfileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustDeviceCustomProfileResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
