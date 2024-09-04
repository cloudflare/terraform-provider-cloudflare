// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewayPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "The API resource UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"action": schema.StringAttribute{
				Description: "The action to preform when the associated traffic, identity, and device posture expressions are either absent or evaluate to `true`.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"on",
						"off",
						"allow",
						"block",
						"scan",
						"noscan",
						"safesearch",
						"ytrestricted",
						"isolate",
						"noisolate",
						"override",
						"l4_override",
						"egress",
						"audit_ssh",
						"resolve",
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "The name of the rule.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the rule.",
				Computed:    true,
				Optional:    true,
			},
			"device_posture": schema.StringAttribute{
				Description: "The wirefilter expression used for device posture check matching.",
				Computed:    true,
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "True if the rule is enabled.",
				Computed:    true,
				Optional:    true,
			},
			"identity": schema.StringAttribute{
				Description: "The wirefilter expression used for identity matching.",
				Computed:    true,
				Optional:    true,
			},
			"precedence": schema.Int64Attribute{
				Description: "Precedence sets the order of your rules. Lower values indicate higher precedence. At each processing phase, applicable rules are evaluated in ascending order of this value.",
				Computed:    true,
				Optional:    true,
			},
			"traffic": schema.StringAttribute{
				Description: "The wirefilter expression used for traffic matching.",
				Computed:    true,
				Optional:    true,
			},
			"filters": schema.ListAttribute{
				Description: "The protocol or layer to evaluate the traffic, identity, and device posture expressions.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"http",
							"dns",
							"l4",
							"egress",
						),
					),
				},
				CustomType:  customfield.NewListType[types.String](ctx),
				ElementType: types.StringType,
			},
			"rule_settings": schema.SingleNestedAttribute{
				Description: "Additional settings that modify the rule's action.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"add_headers": schema.MapAttribute{
						Description: "Add custom headers to allowed requests, in the form of key-value pairs. Keys are header names, pointing to an array with its header value(s).",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewMapType[types.String](ctx),
						ElementType: types.StringType,
					},
					"allow_child_bypass": schema.BoolAttribute{
						Description: "Set by parent MSP accounts to enable their children to bypass this rule.",
						Computed:    true,
						Optional:    true,
					},
					"audit_ssh": schema.SingleNestedAttribute{
						Description: "Settings for the Audit SSH action.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsAuditSSHModel](ctx),
						Attributes: map[string]schema.Attribute{
							"command_logging": schema.BoolAttribute{
								Description: "Enable to turn on SSH command logging.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"biso_admin_controls": schema.SingleNestedAttribute{
						Description: "Configure how browser isolation behaves.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsBISOAdminControlsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"dcp": schema.BoolAttribute{
								Description: "Set to false to enable copy-pasting.",
								Computed:    true,
								Optional:    true,
							},
							"dd": schema.BoolAttribute{
								Description: "Set to false to enable downloading.",
								Computed:    true,
								Optional:    true,
							},
							"dk": schema.BoolAttribute{
								Description: "Set to false to enable keyboard usage.",
								Computed:    true,
								Optional:    true,
							},
							"dp": schema.BoolAttribute{
								Description: "Set to false to enable printing.",
								Computed:    true,
								Optional:    true,
							},
							"du": schema.BoolAttribute{
								Description: "Set to false to enable uploading.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"block_page_enabled": schema.BoolAttribute{
						Description: "Enable the custom block page.",
						Computed:    true,
						Optional:    true,
					},
					"block_reason": schema.StringAttribute{
						Description: "The text describing why this block occurred, displayed on the custom block page (if enabled).",
						Computed:    true,
						Optional:    true,
					},
					"bypass_parent_rule": schema.BoolAttribute{
						Description: "Set by children MSP accounts to bypass their parent's rules.",
						Computed:    true,
						Optional:    true,
					},
					"check_session": schema.SingleNestedAttribute{
						Description: "Configure how session check behaves.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsCheckSessionModel](ctx),
						Attributes: map[string]schema.Attribute{
							"duration": schema.StringAttribute{
								Description: "Configure how fresh the session needs to be to be considered valid.",
								Computed:    true,
								Optional:    true,
							},
							"enforce": schema.BoolAttribute{
								Description: "Set to true to enable session enforcement.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"dns_resolvers": schema.SingleNestedAttribute{
						Description: "Add your own custom resolvers to route queries that match the resolver policy. Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will route to the address closest to their origin. Only valid when a rule's action is set to 'resolve'.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsDNSResolversModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.ListNestedAttribute{
								Computed:   true,
								Optional:   true,
								CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV4Model](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "IPv4 address of upstream resolver.",
											Required:    true,
										},
										"port": schema.Int64Attribute{
											Description: "A port number to use for upstream resolver. Defaults to 53 if unspecified.",
											Computed:    true,
											Optional:    true,
										},
										"route_through_private_network": schema.BoolAttribute{
											Description: "Whether to connect to this resolver over a private network. Must be set when vnet_id is set.",
											Computed:    true,
											Optional:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "Optionally specify a virtual network for this resolver. Uses default virtual network id if omitted.",
											Computed:    true,
											Optional:    true,
										},
									},
								},
							},
							"ipv6": schema.ListNestedAttribute{
								Computed:   true,
								Optional:   true,
								CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayPolicyRuleSettingsDNSResolversIPV6Model](ctx),
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "IPv6 address of upstream resolver.",
											Required:    true,
										},
										"port": schema.Int64Attribute{
											Description: "A port number to use for upstream resolver. Defaults to 53 if unspecified.",
											Computed:    true,
											Optional:    true,
										},
										"route_through_private_network": schema.BoolAttribute{
											Description: "Whether to connect to this resolver over a private network. Must be set when vnet_id is set.",
											Computed:    true,
											Optional:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "Optionally specify a virtual network for this resolver. Uses default virtual network id if omitted.",
											Computed:    true,
											Optional:    true,
										},
									},
								},
							},
						},
					},
					"egress": schema.SingleNestedAttribute{
						Description: "Configure how Gateway Proxy traffic egresses. You can enable this setting for rules with Egress actions and filters, or omit it to indicate local egress via WARP IPs.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsEgressModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.StringAttribute{
								Description: "The IPv4 address to be used for egress.",
								Computed:    true,
								Optional:    true,
							},
							"ipv4_fallback": schema.StringAttribute{
								Description: "The fallback IPv4 address to be used for egress in the event of an error egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via WARP IPs.",
								Computed:    true,
								Optional:    true,
							},
							"ipv6": schema.StringAttribute{
								Description: "The IPv6 range to be used for egress.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"ignore_cname_category_matches": schema.BoolAttribute{
						Description: "Set to true, to ignore the category matches at CNAME domains in a response. If unchecked, the categories in this rule will be checked against all the CNAME domain categories in a response.",
						Computed:    true,
						Optional:    true,
					},
					"insecure_disable_dnssec_validation": schema.BoolAttribute{
						Description: "INSECURE - disable DNSSEC validation (for Allow actions).",
						Computed:    true,
						Optional:    true,
					},
					"ip_categories": schema.BoolAttribute{
						Description: "Set to true to enable IPs in DNS resolver category blocks. By default categories only block based on domain names.",
						Computed:    true,
						Optional:    true,
					},
					"ip_indicator_feeds": schema.BoolAttribute{
						Description: "Set to true to include IPs in DNS resolver indicator feed blocks. By default indicator feeds only block based on domain names.",
						Computed:    true,
						Optional:    true,
					},
					"l4override": schema.SingleNestedAttribute{
						Description: "Send matching traffic to the supplied destination IP address and port.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsL4overrideModel](ctx),
						Attributes: map[string]schema.Attribute{
							"ip": schema.StringAttribute{
								Description: "IPv4 or IPv6 address.",
								Computed:    true,
								Optional:    true,
							},
							"port": schema.Int64Attribute{
								Description: "A port number to use for TCP/UDP overrides.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"notification_settings": schema.SingleNestedAttribute{
						Description: "Configure a notification to display on the user's device when this rule is matched.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsNotificationSettingsModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Set notification on",
								Computed:    true,
								Optional:    true,
							},
							"msg": schema.StringAttribute{
								Description: "Customize the message shown in the notification.",
								Computed:    true,
								Optional:    true,
							},
							"support_url": schema.StringAttribute{
								Description: "Optional URL to direct users to additional information. If not set, the notification will open a block page.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"override_host": schema.StringAttribute{
						Description: "Override matching DNS queries with a hostname.",
						Computed:    true,
						Optional:    true,
					},
					"override_ips": schema.ListAttribute{
						Description: "Override matching DNS queries with an IP or set of IPs.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"payload_log": schema.SingleNestedAttribute{
						Description: "Configure DLP payload logging.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsPayloadLogModel](ctx),
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Set to true to enable DLP payload logging for this rule.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"resolve_dns_through_cloudflare": schema.BoolAttribute{
						Description: "Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot be set when dns_resolvers are specified. Only valid when a rule's action is set to 'resolve'.",
						Computed:    true,
						Optional:    true,
					},
					"untrusted_cert": schema.SingleNestedAttribute{
						Description: "Configure behavior when an upstream cert is invalid or an SSL error occurs.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsUntrustedCERTModel](ctx),
						Attributes: map[string]schema.Attribute{
							"action": schema.StringAttribute{
								Description: "The action performed when an untrusted certificate is seen. The default action is an error with HTTP code 526.",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive(
										"pass_through",
										"block",
										"error",
									),
								},
							},
						},
					},
				},
			},
			"schedule": schema.SingleNestedAttribute{
				Description: "The schedule for activating DNS policies. This does not apply to HTTP or network policies.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyScheduleModel](ctx),
				Attributes: map[string]schema.Attribute{
					"fri": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Fridays, in increasing order from 00:00-24:00.  If this parameter is omitted, the rule will be deactivated on Fridays.",
						Computed:    true,
						Optional:    true,
					},
					"mon": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Mondays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Mondays.",
						Computed:    true,
						Optional:    true,
					},
					"sat": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Saturdays, in increasing order from 00:00-24:00.  If this parameter is omitted, the rule will be deactivated on Saturdays.",
						Computed:    true,
						Optional:    true,
					},
					"sun": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Sundays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Sundays.",
						Computed:    true,
						Optional:    true,
					},
					"thu": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Thursdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Thursdays.",
						Computed:    true,
						Optional:    true,
					},
					"time_zone": schema.StringAttribute{
						Description: "The time zone the rule will be evaluated against. If a [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) is provided, Gateway will always use the current time at that time zone. If this parameter is omitted, then Gateway will use the time zone inferred from the user's source IP to evaluate the rule. If Gateway cannot determine the time zone from the IP, we will fall back to the time zone of the user's connected data center.",
						Computed:    true,
						Optional:    true,
					},
					"tue": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Tuesdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Tuesdays.",
						Computed:    true,
						Optional:    true,
					},
					"wed": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Wednesdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Wednesdays.",
						Computed:    true,
						Optional:    true,
					},
				},
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"deleted_at": schema.StringAttribute{
				Description: "Date of deletion, if any.",
				Computed:    true,
				CustomType:  timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
		},
	}
}

func (r *ZeroTrustGatewayPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewayPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
