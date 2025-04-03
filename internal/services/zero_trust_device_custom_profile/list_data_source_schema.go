// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_custom_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceCustomProfilesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfilesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_mode_switch": schema.BoolAttribute{
							Description: "Whether to allow the user to switch WARP between modes.",
							Computed:    true,
						},
						"allow_updates": schema.BoolAttribute{
							Description: "Whether to receive update notifications when a new version of the client is available.",
							Computed:    true,
						},
						"allowed_to_leave": schema.BoolAttribute{
							Description: "Whether to allow devices to leave the organization.",
							Computed:    true,
						},
						"auto_connect": schema.Float64Attribute{
							Description: "The amount of time in seconds to reconnect after having been disabled.",
							Computed:    true,
						},
						"captive_portal": schema.Float64Attribute{
							Description: "Turn on the captive portal after the specified amount of time.",
							Computed:    true,
						},
						"default": schema.BoolAttribute{
							Description: "Whether the policy is the default policy for an account.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the policy.",
							Computed:    true,
						},
						"disable_auto_fallback": schema.BoolAttribute{
							Description: "If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "Whether the policy will be applied to matching devices.",
							Computed:    true,
						},
						"exclude": schema.ListNestedAttribute{
							Description: "List of routes excluded in the WARP client's tunnel.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfilesExcludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description: "The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "A description of the Split Tunnel item, displayed in the client UI.",
										Computed:    true,
									},
									"host": schema.StringAttribute{
										Description: "The domain name to exclude from the tunnel. If `host` is present, `address` must not be present.",
										Computed:    true,
									},
								},
							},
						},
						"exclude_office_ips": schema.BoolAttribute{
							Description: "Whether to add Microsoft IPs to Split Tunnel exclusions.",
							Computed:    true,
						},
						"fallback_domains": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfilesFallbackDomainsDataSourceModel](ctx),
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
						"gateway_unique_id": schema.StringAttribute{
							Computed: true,
						},
						"include": schema.ListNestedAttribute{
							Description: "List of routes included in the WARP client's tunnel.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfilesIncludeDataSourceModel](ctx),
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description: "The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.",
										Computed:    true,
									},
									"description": schema.StringAttribute{
										Description: "A description of the Split Tunnel item, displayed in the client UI.",
										Computed:    true,
									},
									"host": schema.StringAttribute{
										Description: "The domain name to exclude from the tunnel. If `host` is present, `address` must not be present.",
										Computed:    true,
									},
								},
							},
						},
						"lan_allow_minutes": schema.Float64Attribute{
							Description: "The amount of time in minutes a user is allowed access to their LAN. A value of 0 will allow LAN access until the next WARP reconnection, such as a reboot or a laptop waking from sleep. Note that this field is omitted from the response if null or unset.",
							Computed:    true,
						},
						"lan_allow_subnet_size": schema.Float64Attribute{
							Description: "The size of the subnet for the local access network. Note that this field is omitted from the response if null or unset.",
							Computed:    true,
						},
						"match": schema.StringAttribute{
							Description: "The wirefilter expression to match devices.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the device settings profile.",
							Computed:    true,
						},
						"policy_id": schema.StringAttribute{
							Description: "Device ID.",
							Computed:    true,
						},
						"precedence": schema.Float64Attribute{
							Description: "The precedence of the policy. Lower values indicate higher precedence. Policies will be evaluated in ascending order of this field.",
							Computed:    true,
						},
						"register_interface_ip_with_dns": schema.BoolAttribute{
							Description: "Determines if the operating system will register WARP's local interface IP with your on-premises DNS server.",
							Computed:    true,
						},
						"service_mode_v2": schema.SingleNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectType[ZeroTrustDeviceCustomProfilesServiceModeV2DataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"mode": schema.StringAttribute{
									Description: "The mode to run the WARP client under.",
									Computed:    true,
								},
								"port": schema.Float64Attribute{
									Description: "The port number when used with proxy mode.",
									Computed:    true,
								},
							},
						},
						"support_url": schema.StringAttribute{
							Description: "The URL to launch when the Send Feedback button is clicked.",
							Computed:    true,
						},
						"switch_locked": schema.BoolAttribute{
							Description: "Whether to allow the user to turn off the WARP switch and disconnect the client.",
							Computed:    true,
						},
						"target_tests": schema.ListNestedAttribute{
							Computed:   true,
							CustomType: customfield.NewNestedObjectListType[ZeroTrustDeviceCustomProfilesTargetTestsDataSourceModel](ctx),
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
						"tunnel_protocol": schema.StringAttribute{
							Description: "Determines which tunnel protocol to use.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceCustomProfilesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceCustomProfilesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
