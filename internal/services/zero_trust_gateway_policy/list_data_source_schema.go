// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustGatewayPoliciesDataSource)(nil)

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
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustGatewayPoliciesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The API resource UUID.",
							Computed:    true,
						},
						"action": schema.StringAttribute{
							Description: "The action to preform when the associated traffic, identity, and device posture expressions are either absent or evaluate to `true`.",
							Computed:    true,
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
									"resolve",
									"quarantine",
								),
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
						"description": schema.StringAttribute{
							Description: "The description of the rule.",
							Computed:    true,
						},
						"device_posture": schema.StringAttribute{
							Description: "The wirefilter expression used for device posture check matching.",
							Computed:    true,
						},
						"enabled": schema.BoolAttribute{
							Description: "True if the rule is enabled.",
							Computed:    true,
						},
						"expiration": schema.SingleNestedAttribute{
							Description: "The expiration time stamp and default duration of a DNS policy. Takes\nprecedence over the policy's `schedule` configuration, if any.\n\nThis does not apply to HTTP or network policies.\n",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesExpirationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"expires_at": schema.StringAttribute{
									Description: "The time stamp at which the policy will expire and cease to be\napplied.\n\nMust adhere to RFC 3339 and include a UTC offset. Non-zero\noffsets are accepted but will be converted to the equivalent\nvalue with offset zero (UTC+00:00) and will be returned as time\nstamps with offset zero denoted by a trailing 'Z'.\n\nPolicies with an expiration do not consider the timezone of\nclients they are applied to, and expire \"globally\" at the point\ngiven by their `expires_at` value.\n",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"duration": schema.Int64Attribute{
									Description: "The default duration a policy will be active in minutes. Must be set in order to use the `reset_expiration` endpoint on this rule.",
									Computed:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(5),
									},
								},
								"expired": schema.BoolAttribute{
									Description: "Whether the policy has expired.",
									Computed:    true,
								},
							},
						},
						"filters": schema.ListAttribute{
							Description: "The protocol or layer to evaluate the traffic, identity, and device posture expressions.",
							Computed:    true,
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
						"identity": schema.StringAttribute{
							Description: "The wirefilter expression used for identity matching.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "The name of the rule.",
							Computed:    true,
						},
						"precedence": schema.Int64Attribute{
							Description: "Precedence sets the order of your rules. Lower values indicate higher precedence. At each processing phase, applicable rules are evaluated in ascending order of this value.",
							Computed:    true,
						},
						"rule_settings": schema.SingleNestedAttribute{
							Description: "Additional settings that modify the rule's action.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"add_headers": schema.MapAttribute{
									Description: "Add custom headers to allowed requests, in the form of key-value pairs. Keys are header names, pointing to an array with its header value(s).",
									Computed:    true,
									CustomType:  customfield.NewMapType[types.String](ctx),
									ElementType: types.StringType,
								},
								"allow_child_bypass": schema.BoolAttribute{
									Description: "Set by parent MSP accounts to enable their children to bypass this rule.",
									Computed:    true,
								},
								"audit_ssh": schema.SingleNestedAttribute{
									Description: "Settings for the Audit SSH action.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsAuditSSHDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"command_logging": schema.BoolAttribute{
											Description: "Enable to turn on SSH command logging.",
											Computed:    true,
										},
									},
								},
								"biso_admin_controls": schema.SingleNestedAttribute{
									Description: "Configure how browser isolation behaves.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsBISOAdminControlsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"dcp": schema.BoolAttribute{
											Description: "Set to false to enable copy-pasting.",
											Computed:    true,
										},
										"dd": schema.BoolAttribute{
											Description: "Set to false to enable downloading.",
											Computed:    true,
										},
										"dk": schema.BoolAttribute{
											Description: "Set to false to enable keyboard usage.",
											Computed:    true,
										},
										"dp": schema.BoolAttribute{
											Description: "Set to false to enable printing.",
											Computed:    true,
										},
										"du": schema.BoolAttribute{
											Description: "Set to false to enable uploading.",
											Computed:    true,
										},
									},
								},
								"block_page_enabled": schema.BoolAttribute{
									Description: "Enable the custom block page.",
									Computed:    true,
								},
								"block_reason": schema.StringAttribute{
									Description: "The text describing why this block occurred, displayed on the custom block page (if enabled).",
									Computed:    true,
								},
								"bypass_parent_rule": schema.BoolAttribute{
									Description: "Set by children MSP accounts to bypass their parent's rules.",
									Computed:    true,
								},
								"check_session": schema.SingleNestedAttribute{
									Description: "Configure how session check behaves.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsCheckSessionDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description: "Configure how fresh the session needs to be to be considered valid.",
											Computed:    true,
										},
										"enforce": schema.BoolAttribute{
											Description: "Set to true to enable session enforcement.",
											Computed:    true,
										},
									},
								},
								"dns_resolvers": schema.SingleNestedAttribute{
									Description: "Add your own custom resolvers to route queries that match the resolver policy. Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally' are set. DNS queries will route to the address closest to their origin. Only valid when a rule's action is set to 'resolve'.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsDNSResolversDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ipv4": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV4DataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"ip": schema.StringAttribute{
														Description: "IPv4 address of upstream resolver.",
														Computed:    true,
													},
													"port": schema.Int64Attribute{
														Description: "A port number to use for upstream resolver. Defaults to 53 if unspecified.",
														Computed:    true,
													},
													"route_through_private_network": schema.BoolAttribute{
														Description: "Whether to connect to this resolver over a private network. Must be set when vnet_id is set.",
														Computed:    true,
													},
													"vnet_id": schema.StringAttribute{
														Description: "Optionally specify a virtual network for this resolver. Uses default virtual network id if omitted.",
														Computed:    true,
													},
												},
											},
										},
										"ipv6": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV6DataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"ip": schema.StringAttribute{
														Description: "IPv6 address of upstream resolver.",
														Computed:    true,
													},
													"port": schema.Int64Attribute{
														Description: "A port number to use for upstream resolver. Defaults to 53 if unspecified.",
														Computed:    true,
													},
													"route_through_private_network": schema.BoolAttribute{
														Description: "Whether to connect to this resolver over a private network. Must be set when vnet_id is set.",
														Computed:    true,
													},
													"vnet_id": schema.StringAttribute{
														Description: "Optionally specify a virtual network for this resolver. Uses default virtual network id if omitted.",
														Computed:    true,
													},
												},
											},
										},
									},
								},
								"egress": schema.SingleNestedAttribute{
									Description: "Configure how Gateway Proxy traffic egresses. You can enable this setting for rules with Egress actions and filters, or omit it to indicate local egress via WARP IPs.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsEgressDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ipv4": schema.StringAttribute{
											Description: "The IPv4 address to be used for egress.",
											Computed:    true,
										},
										"ipv4_fallback": schema.StringAttribute{
											Description: "The fallback IPv4 address to be used for egress in the event of an error egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via WARP IPs.",
											Computed:    true,
										},
										"ipv6": schema.StringAttribute{
											Description: "The IPv6 range to be used for egress.",
											Computed:    true,
										},
									},
								},
								"ignore_cname_category_matches": schema.BoolAttribute{
									Description: "Set to true, to ignore the category matches at CNAME domains in a response. If unchecked, the categories in this rule will be checked against all the CNAME domain categories in a response.",
									Computed:    true,
								},
								"insecure_disable_dnssec_validation": schema.BoolAttribute{
									Description: "INSECURE - disable DNSSEC validation (for Allow actions).",
									Computed:    true,
								},
								"ip_categories": schema.BoolAttribute{
									Description: "Set to true to enable IPs in DNS resolver category blocks. By default categories only block based on domain names.",
									Computed:    true,
								},
								"ip_indicator_feeds": schema.BoolAttribute{
									Description: "Set to true to include IPs in DNS resolver indicator feed blocks. By default indicator feeds only block based on domain names.",
									Computed:    true,
								},
								"l4override": schema.SingleNestedAttribute{
									Description: "Send matching traffic to the supplied destination IP address and port.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsL4overrideDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "IPv4 or IPv6 address.",
											Computed:    true,
										},
										"port": schema.Int64Attribute{
											Description: "A port number to use for TCP/UDP overrides.",
											Computed:    true,
										},
									},
								},
								"notification_settings": schema.SingleNestedAttribute{
									Description: "Configure a notification to display on the user's device when this rule is matched.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsNotificationSettingsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description: "Set notification on",
											Computed:    true,
										},
										"msg": schema.StringAttribute{
											Description: "Customize the message shown in the notification.",
											Computed:    true,
										},
										"support_url": schema.StringAttribute{
											Description: "Optional URL to direct users to additional information. If not set, the notification will open a block page.",
											Computed:    true,
										},
									},
								},
								"override_host": schema.StringAttribute{
									Description: "Override matching DNS queries with a hostname.",
									Computed:    true,
								},
								"override_ips": schema.ListAttribute{
									Description: "Override matching DNS queries with an IP or set of IPs.",
									Computed:    true,
									CustomType:  customfield.NewListType[types.String](ctx),
									ElementType: types.StringType,
								},
								"payload_log": schema.SingleNestedAttribute{
									Description: "Configure DLP payload logging.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsPayloadLogDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description: "Set to true to enable DLP payload logging for this rule.",
											Computed:    true,
										},
									},
								},
								"quarantine": schema.SingleNestedAttribute{
									Description: "Settings that apply to quarantine rules",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsQuarantineDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"file_types": schema.ListAttribute{
											Description: "Types of files to sandbox.",
											Computed:    true,
											Validators: []validator.List{
												listvalidator.ValueStringsAre(
													stringvalidator.OneOfCaseInsensitive(
														"exe",
														"pdf",
														"doc",
														"docm",
														"docx",
														"rtf",
														"ppt",
														"pptx",
														"xls",
														"xlsm",
														"xlsx",
														"zip",
														"rar",
													),
												),
											},
											CustomType:  customfield.NewListType[types.String](ctx),
											ElementType: types.StringType,
										},
									},
								},
								"resolve_dns_internally": schema.SingleNestedAttribute{
									Description: "Configure to forward the query to the internal DNS service, passing the specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is set to 'resolve'.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsResolveDNSInternallyDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"fallback": schema.StringAttribute{
											Description: "The fallback behavior to apply when the internal DNS response code is different from 'NOERROR' or when the response data only contains CNAME records for 'A' or 'AAAA' queries.",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("none", "public_dns"),
											},
										},
										"view_id": schema.StringAttribute{
											Description: "The internal DNS view identifier that's passed to the internal DNS service.",
											Computed:    true,
										},
									},
								},
								"resolve_dns_through_cloudflare": schema.BoolAttribute{
									Description: "Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot be set when 'dns_resolvers' are specified or 'resolve_dns_internally' is set. Only valid when a rule's action is set to 'resolve'.",
									Computed:    true,
								},
								"untrusted_cert": schema.SingleNestedAttribute{
									Description: "Configure behavior when an upstream cert is invalid or an SSL error occurs.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsUntrustedCERTDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description: "The action performed when an untrusted certificate is seen. The default action is an error with HTTP code 526.",
											Computed:    true,
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
							CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesScheduleDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"fri": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Fridays, in increasing order from 00:00-24:00.  If this parameter is omitted, the rule will be deactivated on Fridays.",
									Computed:    true,
								},
								"mon": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Mondays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Mondays.",
									Computed:    true,
								},
								"sat": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Saturdays, in increasing order from 00:00-24:00.  If this parameter is omitted, the rule will be deactivated on Saturdays.",
									Computed:    true,
								},
								"sun": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Sundays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Sundays.",
									Computed:    true,
								},
								"thu": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Thursdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Thursdays.",
									Computed:    true,
								},
								"time_zone": schema.StringAttribute{
									Description: "The time zone the rule will be evaluated against. If a [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) is provided, Gateway will always use the current time at that time zone. If this parameter is omitted, then Gateway will use the time zone inferred from the user's source IP to evaluate the rule. If Gateway cannot determine the time zone from the IP, we will fall back to the time zone of the user's connected data center.",
									Computed:    true,
								},
								"tue": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Tuesdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Tuesdays.",
									Computed:    true,
								},
								"wed": schema.StringAttribute{
									Description: "The time intervals when the rule will be active on Wednesdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Wednesdays.",
									Computed:    true,
								},
							},
						},
						"traffic": schema.StringAttribute{
							Description: "The wirefilter expression used for traffic matching.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"version": schema.Int64Attribute{
							Description: "version number of the rule",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *ZeroTrustGatewayPoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *ZeroTrustGatewayPoliciesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
