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
						"action": schema.StringAttribute{
							Description: "Specify the action to perform when the associated traffic, identity, and device posture expressions either absent or evaluate to `true`.\nAvailable values: \"on\", \"off\", \"allow\", \"block\", \"scan\", \"noscan\", \"safesearch\", \"ytrestricted\", \"isolate\", \"noisolate\", \"override\", \"l4_override\", \"egress\", \"resolve\", \"quarantine\", \"redirect\".",
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
									"redirect",
								),
							},
						},
						"enabled": schema.BoolAttribute{
							Description: "Specify whether the rule is enabled.",
							Computed:    true,
						},
						"filters": schema.ListAttribute{
							Description: "Specify the protocol or layer to evaluate the traffic, identity, and device posture expressions.",
							Computed:    true,
							Validators: []validator.List{
								listvalidator.ValueStringsAre(
									stringvalidator.OneOfCaseInsensitive(
										"http",
										"dns",
										"l4",
										"egress",
										"dns_resolver",
									),
								),
							},
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
						"name": schema.StringAttribute{
							Description: "Specify the rule name.",
							Computed:    true,
						},
						"precedence": schema.Int64Attribute{
							Description: "Set the order of your rules. Lower values indicate higher precedence. At each processing phase, evaluate applicable rules in ascending order of this value. Refer to [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform) to manage precedence via Terraform.",
							Computed:    true,
						},
						"traffic": schema.StringAttribute{
							Description: "Specify the wirefilter expression used for traffic matching. The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.",
							Computed:    true,
						},
						"id": schema.StringAttribute{
							Description: "Identify the API resource with a UUID.",
							Computed:    true,
						},
						"created_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"deleted_at": schema.StringAttribute{
							Description: "Indicate the date of deletion, if any.",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"description": schema.StringAttribute{
							Description: "Specify the rule description.",
							Computed:    true,
						},
						"device_posture": schema.StringAttribute{
							Description: "Specify the wirefilter expression used for device posture check. The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.",
							Computed:    true,
						},
						"expiration": schema.SingleNestedAttribute{
							Description: "Defines the expiration time stamp and default duration of a DNS policy. Takes precedence over the policy's `schedule` configuration, if any. This  does not apply to HTTP or network policies.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesExpirationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"expires_at": schema.StringAttribute{
									Description: "Show the timestamp when the policy expires and stops applying.  The value must follow RFC 3339 and include a UTC offset.  The system accepts non-zero offsets but converts them to the equivalent UTC+00:00  value and returns timestamps with a trailing Z. Expiration policies ignore client  timezones and expire globally at the specified expires_at time.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"duration": schema.Int64Attribute{
									Description: "Defines the default duration a policy active in minutes. Must set in order to use the `reset_expiration` endpoint on this rule.",
									Computed:    true,
									Validators: []validator.Int64{
										int64validator.AtLeast(5),
									},
								},
								"expired": schema.BoolAttribute{
									Description: "Indicates whether the policy is expired.",
									Computed:    true,
								},
							},
						},
						"identity": schema.StringAttribute{
							Description: "Specify the wirefilter expression used for identity matching. The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.",
							Computed:    true,
						},
						"read_only": schema.BoolAttribute{
							Description: "Indicate that this rule is shared via the Orgs API and read only.",
							Computed:    true,
						},
						"rule_settings": schema.SingleNestedAttribute{
							Description: "Set settings related to this rule.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"add_headers": schema.MapAttribute{
									Description: "Add custom headers to allowed requests as key-value pairs. Use header names as keys that map to arrays of header values.",
									Computed:    true,
									CustomType:  customfield.NewMapType[customfield.List[types.String]](ctx),
									ElementType: types.ListType{
										ElemType: types.StringType,
									},
								},
								"allow_child_bypass": schema.BoolAttribute{
									Description: "Set to enable MSP children to bypass this rule. Only parent MSP accounts can set this. this rule.",
									Computed:    true,
								},
								"audit_ssh": schema.SingleNestedAttribute{
									Description: "Define the settings for the Audit SSH action.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsAuditSSHDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"command_logging": schema.BoolAttribute{
											Description: "Enable SSH command logging.",
											Computed:    true,
										},
									},
								},
								"biso_admin_controls": schema.SingleNestedAttribute{
									Description: "Configure browser isolation behavior.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsBISOAdminControlsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"copy": schema.StringAttribute{
											Description: "Configure copy behavior. If set to remote_only, users cannot copy isolated content from the remote browser to the local clipboard. If this field is absent, copying remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\", \"remote_only\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"enabled",
													"disabled",
													"remote_only",
												),
											},
										},
										"dcp": schema.BoolAttribute{
											Description: "Set to false to enable copy-pasting. Only applies when `version == \"v1\"`.",
											Computed:    true,
										},
										"dd": schema.BoolAttribute{
											Description: "Set to false to enable downloading. Only applies when `version == \"v1\"`.",
											Computed:    true,
										},
										"dk": schema.BoolAttribute{
											Description: "Set to false to enable keyboard usage. Only applies when `version == \"v1\"`.",
											Computed:    true,
										},
										"download": schema.StringAttribute{
											Description: "Configure download behavior. When set to remote_only, users can view downloads but cannot save them. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\", \"remote_only\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"enabled",
													"disabled",
													"remote_only",
												),
											},
										},
										"dp": schema.BoolAttribute{
											Description: "Set to false to enable printing. Only applies when `version == \"v1\"`.",
											Computed:    true,
										},
										"du": schema.BoolAttribute{
											Description: "Set to false to enable uploading. Only applies when `version == \"v1\"`.",
											Computed:    true,
										},
										"keyboard": schema.StringAttribute{
											Description: "Configure keyboard usage behavior. If this field is absent, keyboard usage remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
											},
										},
										"paste": schema.StringAttribute{
											Description: "Configure paste behavior. If set to remote_only, users cannot paste content from the local clipboard into isolated pages. If this field is absent, pasting remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\", \"remote_only\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive(
													"enabled",
													"disabled",
													"remote_only",
												),
											},
										},
										"printing": schema.StringAttribute{
											Description: "Configure print behavior. Default, Printing is enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
											},
										},
										"upload": schema.StringAttribute{
											Description: "Configure upload behavior. If this field is absent, uploading remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
											},
										},
										"version": schema.StringAttribute{
											Description: "Indicate which version of the browser isolation controls should apply.\nAvailable values: \"v1\", \"v2\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("v1", "v2"),
											},
										},
									},
								},
								"block_page": schema.SingleNestedAttribute{
									Description: "Configure custom block page settings. If missing or null, use the account settings.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsBlockPageDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"target_uri": schema.StringAttribute{
											Description: "Specify the URI to which the user is redirected.",
											Computed:    true,
										},
										"include_context": schema.BoolAttribute{
											Description: "Specify whether to pass the context information as query parameters.",
											Computed:    true,
										},
									},
								},
								"block_page_enabled": schema.BoolAttribute{
									Description: "Enable the custom block page.",
									Computed:    true,
								},
								"block_reason": schema.StringAttribute{
									Description: "Explain why the rule blocks the request. The custom block page shows this text (if enabled).",
									Computed:    true,
								},
								"bypass_parent_rule": schema.BoolAttribute{
									Description: "Set to enable MSP accounts to bypass their parent's rules. Only MSP child accounts can set this.",
									Computed:    true,
								},
								"check_session": schema.SingleNestedAttribute{
									Description: "Configure session check behavior.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsCheckSessionDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description: "Sets the required session freshness threshold. The API returns a normalized version of this value.",
											Computed:    true,
										},
										"enforce": schema.BoolAttribute{
											Description: "Enable session enforcement.",
											Computed:    true,
										},
									},
								},
								"dns_resolvers": schema.SingleNestedAttribute{
									Description: "Configure custom resolvers to route queries that match the resolver policy. Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally' settings. DNS queries get routed to the address closest to their origin. Only valid when a rule's action is set to 'resolve'.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsDNSResolversDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ipv4": schema.ListNestedAttribute{
											Computed:   true,
											CustomType: customfield.NewNestedObjectListType[ZeroTrustGatewayPoliciesRuleSettingsDNSResolversIPV4DataSourceModel](ctx),
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"ip": schema.StringAttribute{
														Description: "Specify the IPv4 address of the upstream resolver.",
														Computed:    true,
													},
													"port": schema.Int64Attribute{
														Description: "Specify a port number to use for the upstream resolver. Defaults to 53 if unspecified.",
														Computed:    true,
													},
													"route_through_private_network": schema.BoolAttribute{
														Description: "Indicate whether to connect to this resolver over a private network. Must set when vnet_id set.",
														Computed:    true,
													},
													"vnet_id": schema.StringAttribute{
														Description: "Specify an optional virtual network for this resolver. Uses default virtual network id if omitted.",
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
														Description: "Specify the IPv6 address of the upstream resolver.",
														Computed:    true,
													},
													"port": schema.Int64Attribute{
														Description: "Specify a port number to use for the upstream resolver. Defaults to 53 if unspecified.",
														Computed:    true,
													},
													"route_through_private_network": schema.BoolAttribute{
														Description: "Indicate whether to connect to this resolver over a private network. Must set when vnet_id set.",
														Computed:    true,
													},
													"vnet_id": schema.StringAttribute{
														Description: "Specify an optional virtual network for this resolver. Uses default virtual network id if omitted.",
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
											Description: "Specify the IPv4 address to use for egress.",
											Computed:    true,
										},
										"ipv4_fallback": schema.StringAttribute{
											Description: "Specify the fallback IPv4 address to use for egress when the primary IPv4 fails. Set '0.0.0.0' to indicate local egress via WARP IPs.",
											Computed:    true,
										},
										"ipv6": schema.StringAttribute{
											Description: "Specify the IPv6 range to use for egress.",
											Computed:    true,
										},
									},
								},
								"ignore_cname_category_matches": schema.BoolAttribute{
									Description: "Ignore category matches at CNAME domains in a response. When off, evaluate categories in this rule against all CNAME domain categories in the response.",
									Computed:    true,
								},
								"insecure_disable_dnssec_validation": schema.BoolAttribute{
									Description: "Specify whether to disable DNSSEC validation (for Allow actions) [INSECURE].",
									Computed:    true,
								},
								"ip_categories": schema.BoolAttribute{
									Description: "Enable IPs in DNS resolver category blocks. The system blocks only domain name categories unless you enable this setting.",
									Computed:    true,
								},
								"ip_indicator_feeds": schema.BoolAttribute{
									Description: "Indicates whether to include IPs in DNS resolver indicator feed blocks. Default, indicator feeds block only domain names.",
									Computed:    true,
								},
								"l4override": schema.SingleNestedAttribute{
									Description: "Send matching traffic to the supplied destination IP address. and port.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsL4overrideDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "Defines the IPv4 or IPv6 address.",
											Computed:    true,
										},
										"port": schema.Int64Attribute{
											Description: "Defines a port number to use for TCP/UDP overrides.",
											Computed:    true,
										},
									},
								},
								"notification_settings": schema.SingleNestedAttribute{
									Description: "Configure a notification to display on the user's device when this rule matched.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsNotificationSettingsDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"enabled": schema.BoolAttribute{
											Description: "Enable notification.",
											Computed:    true,
										},
										"include_context": schema.BoolAttribute{
											Description: "Indicates whether to pass the context information as query parameters.",
											Computed:    true,
										},
										"msg": schema.StringAttribute{
											Description: "Customize the message shown in the notification.",
											Computed:    true,
										},
										"support_url": schema.StringAttribute{
											Description: "Defines an optional URL to direct users to additional information. If unset, the notification opens a block page.",
											Computed:    true,
										},
									},
								},
								"override_host": schema.StringAttribute{
									Description: "Defines a hostname for override, for the matching DNS queries.",
									Computed:    true,
								},
								"override_ips": schema.ListAttribute{
									Description: "Defines a an IP or set of IPs for overriding matched DNS queries.",
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
											Description: "Enable DLP payload logging for this rule.",
											Computed:    true,
										},
									},
								},
								"quarantine": schema.SingleNestedAttribute{
									Description: "Configure settings that apply to quarantine rules.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsQuarantineDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"file_types": schema.ListAttribute{
											Description: "Specify the types of files to sandbox.",
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
								"redirect": schema.SingleNestedAttribute{
									Description: "Apply settings to redirect rules.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsRedirectDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"target_uri": schema.StringAttribute{
											Description: "Specify the URI to which the user is redirected.",
											Computed:    true,
										},
										"include_context": schema.BoolAttribute{
											Description: "Specify whether to pass the context information as query parameters.",
											Computed:    true,
										},
										"preserve_path_and_query": schema.BoolAttribute{
											Description: "Specify whether to append the path and query parameters from the original request to target_uri.",
											Computed:    true,
										},
									},
								},
								"resolve_dns_internally": schema.SingleNestedAttribute{
									Description: "Configure to forward the query to the internal DNS service, passing the specified 'view_id' as input. Not used when 'dns_resolvers' is specified or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is set to 'resolve'.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsResolveDNSInternallyDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"fallback": schema.StringAttribute{
											Description: "Specify the fallback behavior to apply when the internal DNS response code differs from 'NOERROR' or when the response data contains only CNAME records for 'A' or 'AAAA' queries.\nAvailable values: \"none\", \"public_dns\".",
											Computed:    true,
											Validators: []validator.String{
												stringvalidator.OneOfCaseInsensitive("none", "public_dns"),
											},
										},
										"view_id": schema.StringAttribute{
											Description: "Specify the internal DNS view identifier to pass to the internal DNS service.",
											Computed:    true,
										},
									},
								},
								"resolve_dns_through_cloudflare": schema.BoolAttribute{
									Description: "Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot set when 'dns_resolvers' specified or 'resolve_dns_internally' is set. Only valid when a rule's action set to 'resolve'.",
									Computed:    true,
								},
								"untrusted_cert": schema.SingleNestedAttribute{
									Description: "Configure behavior when an upstream certificate is invalid or an SSL error occurs.",
									Computed:    true,
									CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesRuleSettingsUntrustedCERTDataSourceModel](ctx),
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description: "Defines the action performed when an untrusted certificate seen. The default action an error with HTTP code 526.\nAvailable values: \"pass_through\", \"block\", \"error\".",
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
							Description: "Defines the schedule for activating DNS policies. (HTTP/Egress or L4 unsupported).",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPoliciesScheduleDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"fri": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Fridays, in the increasing order from 00:00-24:00.  If this parameter omitted, the rule is deactivated on Fridays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
								"mon": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Mondays, in the increasing order from 00:00-24:00(capped at maximum of 6 time splits). If this parameter omitted, the rule is deactivated on Mondays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
								"sat": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Saturdays, in the increasing order from 00:00-24:00.  If this parameter omitted, the rule is deactivated on Saturdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
								"sun": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Sundays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Sundays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
								"thu": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Thursdays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Thursdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
								"time_zone": schema.StringAttribute{
									Description: "Specify the time zone for rule evaluation. When a [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) is provided, Gateway always uses the current time for that time zone. When this parameter is omitted, Gateway uses the time zone determined from the user's IP address. Colo time zone is used when the user's IP address does not resolve to a location.",
									Computed:    true,
								},
								"tue": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Tuesdays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Tuesdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
								"wed": schema.StringAttribute{
									Description: "Specify the time intervals when the rule is active on Wednesdays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Wednesdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
									Computed:    true,
								},
							},
						},
						"sharable": schema.BoolAttribute{
							Description: "Indicate that this rule is sharable via the Orgs API.",
							Computed:    true,
						},
						"source_account": schema.StringAttribute{
							Description: "Provide the account tag of the account that created the rule.",
							Computed:    true,
						},
						"updated_at": schema.StringAttribute{
							Computed:   true,
							CustomType: timetypes.RFC3339Type{},
						},
						"version": schema.Int64Attribute{
							Description: "Indicate the version number of the rule(read-only).",
							Computed:    true,
						},
						"warning_status": schema.StringAttribute{
							Description: "Indicate a warning for a misconfigured rule, if any.",
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
