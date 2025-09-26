// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_gateway_policy

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.ResourceWithConfigValidators = (*ZeroTrustGatewayPolicyResource)(nil)

func ResourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "Identify the API resource with a UUID.",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"action": schema.StringAttribute{
				Description: "Specify the action to perform when the associated traffic, identity, and device posture expressions either absent or evaluate to `true`.\nAvailable values: \"on\", \"off\", \"allow\", \"block\", \"scan\", \"noscan\", \"safesearch\", \"ytrestricted\", \"isolate\", \"noisolate\", \"override\", \"l4_override\", \"egress\", \"resolve\", \"quarantine\", \"redirect\".",
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
						"resolve",
						"quarantine",
						"redirect",
					),
				},
			},
			"name": schema.StringAttribute{
				Description: "Specify the rule name.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Specify the rule description.",
				Optional:    true,
			},
			"filters": schema.ListAttribute{
				Description: "Specify the protocol or layer to evaluate the traffic, identity, and device posture expressions.",
				Optional:    true,
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
				ElementType: types.StringType,
			},
			"device_posture": schema.StringAttribute{
				Description: "Specify the wirefilter expression used for device posture check. The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"enabled": schema.BoolAttribute{
				Description: "Specify whether the rule is enabled.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(false),
			},
			"identity": schema.StringAttribute{
				Description: "Specify the wirefilter expression used for identity matching. The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"precedence": schema.Int64Attribute{
				Description: "Set the order of your rules. Lower values indicate higher precedence. At each processing phase, evaluate applicable rules in ascending order of this value. Refer to [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform) to manage precedence via Terraform.",
				Computed:    true,
				Optional:    true,
			},
			"traffic": schema.StringAttribute{
				Description: "Specify the wirefilter expression used for traffic matching. The API automatically formats and sanitizes expressions before storing them. To prevent Terraform state drift, use the formatted expression returned in the API response.",
				Computed:    true,
				Optional:    true,
				Default:     stringdefault.StaticString(""),
			},
			"expiration": schema.SingleNestedAttribute{
				Description: "Defines the expiration time stamp and default duration of a DNS policy. Takes precedence over the policy's `schedule` configuration, if any. This  does not apply to HTTP or network policies. Settable only for `dns` rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyExpirationModel](ctx),
				Attributes: map[string]schema.Attribute{
					"expires_at": schema.StringAttribute{
						Description: "Show the timestamp when the policy expires and stops applying.  The value must follow RFC 3339 and include a UTC offset.  The system accepts non-zero offsets but converts them to the equivalent UTC+00:00  value and returns timestamps with a trailing Z. Expiration policies ignore client  timezones and expire globally at the specified expires_at time.",
						Required:    true,
						CustomType:  timetypes.RFC3339Type{},
					},
					"duration": schema.Int64Attribute{
						Description: "Defines the default duration a policy active in minutes. Must set in order to use the `reset_expiration` endpoint on this rule.",
						Optional:    true,
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
			"rule_settings": schema.SingleNestedAttribute{
				Description: "Set settings related to this rule. Each setting is only valid for specific rule types and can only be used with the appropriate selectors. If Terraform drift is observed in these setting values, verify that the setting is supported for the given rule type and that the API response reflects the requested value. If the API response returns sanitized or modified values that differ from the request, use the API-provided values in Terraform to ensure consistency.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyRuleSettingsModel](ctx),
				Attributes: map[string]schema.Attribute{
					"add_headers": schema.MapAttribute{
						Description: "Add custom headers to allowed requests as key-value pairs. Use header names as keys that map to arrays of header values. Settable only for `http` rules with the action set to `allow`.",
						Optional:    true,
						ElementType: types.ListType{
							ElemType: types.StringType,
						},
					},
					"allow_child_bypass": schema.BoolAttribute{
						Description: "Set to enable MSP children to bypass this rule. Only parent MSP accounts can set this. this rule. Settable for all types of rules.",
						Computed:    true,
						Optional:    true,
					},
					"audit_ssh": schema.SingleNestedAttribute{
						Description: "Define the settings for the Audit SSH action. Settable only for `l4` rules with `audit_ssh` action.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"command_logging": schema.BoolAttribute{
								Description: "Enable SSH command logging.",
								Optional:    true,
							},
						},
					},
					"biso_admin_controls": schema.SingleNestedAttribute{
						Description: "Configure browser isolation behavior. Settable only for `http` rules with the action set to `isolate`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"copy": schema.StringAttribute{
								Description: "Configure copy behavior. If set to remote_only, users cannot copy isolated content from the remote browser to the local clipboard. If this field is absent, copying remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\", \"remote_only\".",
								Optional:    true,
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
								Optional:    true,
							},
							"dd": schema.BoolAttribute{
								Description: "Set to false to enable downloading. Only applies when `version == \"v1\"`.",
								Computed:    true,
								Optional:    true,
							},
							"dk": schema.BoolAttribute{
								Description: "Set to false to enable keyboard usage. Only applies when `version == \"v1\"`.",
								Computed:    true,
								Optional:    true,
							},
							"download": schema.StringAttribute{
								Description: "Configure download behavior. When set to remote_only, users can view downloads but cannot save them. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\", \"remote_only\".",
								Optional:    true,
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
								Optional:    true,
							},
							"du": schema.BoolAttribute{
								Description: "Set to false to enable uploading. Only applies when `version == \"v1\"`.",
								Computed:    true,
								Optional:    true,
							},
							"keyboard": schema.StringAttribute{
								Description: "Configure keyboard usage behavior. If this field is absent, keyboard usage remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
								},
							},
							"paste": schema.StringAttribute{
								Description: "Configure paste behavior. If set to remote_only, users cannot paste content from the local clipboard into isolated pages. If this field is absent, pasting remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\", \"remote_only\".",
								Optional:    true,
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
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
								},
							},
							"upload": schema.StringAttribute{
								Description: "Configure upload behavior. If this field is absent, uploading remains enabled. Applies only when version == \"v2\".\nAvailable values: \"enabled\", \"disabled\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("enabled", "disabled"),
								},
							},
							"version": schema.StringAttribute{
								Description: "Indicate which version of the browser isolation controls should apply.\nAvailable values: \"v1\", \"v2\".",
								Computed:    true,
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("v1", "v2"),
								},
								Default: stringdefault.StaticString("v1"),
							},
						},
					},
					"block_page": schema.SingleNestedAttribute{
						Description: "Configure custom block page settings. If missing or null, use the account settings. Settable only for `http` rules with the action set to `block`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"target_uri": schema.StringAttribute{
								Description: "Specify the URI to which the user is redirected.",
								Required:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Specify whether to pass the context information as query parameters.",
								Optional:    true,
							},
						},
					},
					"block_page_enabled": schema.BoolAttribute{
						Description: "Enable the custom block page. Settable only for `dns` rules with action `block`.",
						Computed:    true,
						Optional:    true,
					},
					"block_reason": schema.StringAttribute{
						Description: "Explain why the rule blocks the request. The custom block page shows this text (if enabled). Settable only for `dns`, `l4`, and `http` rules when the action set to `block`.",
						Computed:    true,
						Optional:    true,
					},
					"bypass_parent_rule": schema.BoolAttribute{
						Description: "Set to enable MSP accounts to bypass their parent's rules. Only MSP child accounts can set this. Settable for all types of rules.",
						Optional:    true,
					},
					"check_session": schema.SingleNestedAttribute{
						Description: "Configure session check behavior. Settable only for `l4` and `http` rules with the action set to `allow`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"duration": schema.StringAttribute{
								Description: "Sets the required session freshness threshold. The API returns a normalized version of this value.",
								Optional:    true,
							},
							"enforce": schema.BoolAttribute{
								Description: "Enable session enforcement.",
								Optional:    true,
							},
						},
					},
					"dns_resolvers": schema.SingleNestedAttribute{
						Description: "Configure custom resolvers to route queries that match the resolver policy. Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally' settings. DNS queries get routed to the address closest to their origin. Only valid when a rule's action set to 'resolve'. Settable only for `dns_resolver` rules.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.ListNestedAttribute{
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "Specify the IPv4 address of the upstream resolver.",
											Required:    true,
										},
										"port": schema.Int64Attribute{
											Description: "Specify a port number to use for the upstream resolver. Defaults to 53 if unspecified.",
											Optional:    true,
										},
										"route_through_private_network": schema.BoolAttribute{
											Description: "Indicate whether to connect to this resolver over a private network. Must set when vnet_id set.",
											Optional:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "Specify an optional virtual network for this resolver. Uses default virtual network id if omitted.",
											Optional:    true,
										},
									},
								},
							},
							"ipv6": schema.ListNestedAttribute{
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "Specify the IPv6 address of the upstream resolver.",
											Required:    true,
										},
										"port": schema.Int64Attribute{
											Description: "Specify a port number to use for the upstream resolver. Defaults to 53 if unspecified.",
											Optional:    true,
										},
										"route_through_private_network": schema.BoolAttribute{
											Description: "Indicate whether to connect to this resolver over a private network. Must set when vnet_id set.",
											Optional:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "Specify an optional virtual network for this resolver. Uses default virtual network id if omitted.",
											Optional:    true,
										},
									},
								},
							},
						},
					},
					"egress": schema.SingleNestedAttribute{
						Description: "Configure how Gateway Proxy traffic egresses. You can enable this setting for rules with Egress actions and filters, or omit it to indicate local egress via WARP IPs. Settable only for `egress` rules.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.StringAttribute{
								Description: "Specify the IPv4 address to use for egress.",
								Optional:    true,
							},
							"ipv4_fallback": schema.StringAttribute{
								Description: "Specify the fallback IPv4 address to use for egress when the primary IPv4 fails. Set '0.0.0.0' to indicate local egress via WARP IPs.",
								Optional:    true,
							},
							"ipv6": schema.StringAttribute{
								Description: "Specify the IPv6 range to use for egress.",
								Optional:    true,
							},
						},
					},
					"ignore_cname_category_matches": schema.BoolAttribute{
						Description: "Ignore category matches at CNAME domains in a response. When off, evaluate categories in this rule against all CNAME domain categories in the response. Settable only for `dns` and `dns_resolver` rules.",
						Computed:    true,
						Optional:    true,
					},
					"insecure_disable_dnssec_validation": schema.BoolAttribute{
						Description: "Specify whether to disable DNSSEC validation (for Allow actions) [INSECURE]. Settable only for `dns` rules.",
						Computed:    true,
						Optional:    true,
					},
					"ip_categories": schema.BoolAttribute{
						Description: "Enable IPs in DNS resolver category blocks. The system blocks only domain name categories unless you enable this setting. Settable only for `dns` and `dns_resolver` rules.",
						Computed:    true,
						Optional:    true,
					},
					"ip_indicator_feeds": schema.BoolAttribute{
						Description: "Indicates whether to include IPs in DNS resolver indicator feed blocks. Default, indicator feeds block only domain names. Settable only for `dns` and `dns_resolver` rules.",
						Computed:    true,
						Optional:    true,
					},
					"l4override": schema.SingleNestedAttribute{
						Description: "Send matching traffic to the supplied destination IP address and port. Settable only for `l4` rules with the action set to `l4_override`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ip": schema.StringAttribute{
								Description: "Defines the IPv4 or IPv6 address.",
								Optional:    true,
							},
							"port": schema.Int64Attribute{
								Description: "Defines a port number to use for TCP/UDP overrides.",
								Optional:    true,
							},
						},
					},
					"notification_settings": schema.SingleNestedAttribute{
						Description: "Configure a notification to display on the user's device when this rule matched. Settable for all types of rules with the action set to `block`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable notification.",
								Optional:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Indicates whether to pass the context information as query parameters.",
								Optional:    true,
							},
							"msg": schema.StringAttribute{
								Description: "Customize the message shown in the notification.",
								Optional:    true,
							},
							"support_url": schema.StringAttribute{
								Description: "Defines an optional URL to direct users to additional information. If unset, the notification opens a block page.",
								Optional:    true,
							},
						},
					},
					"override_host": schema.StringAttribute{
						Description: "Defines a hostname for override, for the matching DNS queries. Settable only for `dns` rules with the action set to `override`.",
						Computed:    true,
						Optional:    true,
					},
					"override_ips": schema.ListAttribute{
						Description: "Defines a an IP or set of IPs for overriding matched DNS queries. Settable only for `dns` rules with the action set to `override`.",
						Computed:    true,
						Optional:    true,
						CustomType:  customfield.NewListType[types.String](ctx),
						ElementType: types.StringType,
					},
					"payload_log": schema.SingleNestedAttribute{
						Description: "Configure DLP payload logging. Settable only for `http` rules.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Enable DLP payload logging for this rule.",
								Optional:    true,
							},
						},
					},
					"quarantine": schema.SingleNestedAttribute{
						Description: "Configure settings that apply to quarantine rules. Settable only for `http` rules.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"file_types": schema.ListAttribute{
								Description: "Specify the types of files to sandbox.",
								Optional:    true,
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
								ElementType: types.StringType,
							},
						},
					},
					"redirect": schema.SingleNestedAttribute{
						Description: "Apply settings to redirect rules. Settable only for `http` rules with the action set to `redirect`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"target_uri": schema.StringAttribute{
								Description: "Specify the URI to which the user is redirected.",
								Required:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Specify whether to pass the context information as query parameters.",
								Optional:    true,
							},
							"preserve_path_and_query": schema.BoolAttribute{
								Description: "Specify whether to append the path and query parameters from the original request to target_uri.",
								Optional:    true,
							},
						},
					},
					"resolve_dns_internally": schema.SingleNestedAttribute{
						Description: "Configure to forward the query to the internal DNS service, passing the specified 'view_id' as input. Not used when 'dns_resolvers' is specified or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action set to 'resolve'. Settable only for `dns_resolver` rules.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"fallback": schema.StringAttribute{
								Description: "Specify the fallback behavior to apply when the internal DNS response code differs from 'NOERROR' or when the response data contains only CNAME records for 'A' or 'AAAA' queries.\nAvailable values: \"none\", \"public_dns\".",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("none", "public_dns"),
								},
							},
							"view_id": schema.StringAttribute{
								Description: "Specify the internal DNS view identifier to pass to the internal DNS service.",
								Optional:    true,
							},
						},
					},
					"resolve_dns_through_cloudflare": schema.BoolAttribute{
						Description: "Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot set when 'dns_resolvers' specified or 'resolve_dns_internally' is set. Only valid when a rule's action set to 'resolve'. Settable only for `dns_resolver` rules.",
						Computed:    true,
						Optional:    true,
					},
					"untrusted_cert": schema.SingleNestedAttribute{
						Description: "Configure behavior when an upstream certificate is invalid or an SSL error occurs. Settable only for `http` rules with the action set to `allow`.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"action": schema.StringAttribute{
								Description: "Defines the action performed when an untrusted certificate seen. The default action an error with HTTP code 526.\nAvailable values: \"pass_through\", \"block\", \"error\".",
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
				Description: "Defines the schedule for activating DNS policies. Settable only for `dns` and `dns_resolver` rules.",
				Computed:    true,
				Optional:    true,
				CustomType:  customfield.NewNestedObjectType[ZeroTrustGatewayPolicyScheduleModel](ctx),
				Attributes: map[string]schema.Attribute{
					"fri": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Fridays, in the increasing order from 00:00-24:00.  If this parameter omitted, the rule is deactivated on Fridays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
					"mon": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Mondays, in the increasing order from 00:00-24:00(capped at maximum of 6 time splits). If this parameter omitted, the rule is deactivated on Mondays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
					"sat": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Saturdays, in the increasing order from 00:00-24:00.  If this parameter omitted, the rule is deactivated on Saturdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
					"sun": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Sundays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Sundays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
					"thu": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Thursdays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Thursdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
					"time_zone": schema.StringAttribute{
						Description: "Specify the time zone for rule evaluation. When a [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) is provided, Gateway always uses the current time for that time zone. When this parameter is omitted, Gateway uses the time zone determined from the user's IP address. Colo time zone is used when the user's IP address does not resolve to a location.",
						Optional:    true,
					},
					"tue": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Tuesdays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Tuesdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
					"wed": schema.StringAttribute{
						Description: "Specify the time intervals when the rule is active on Wednesdays, in the increasing order from 00:00-24:00. If this parameter omitted, the rule is deactivated on Wednesdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.",
						Optional:    true,
					},
				},
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
			"read_only": schema.BoolAttribute{
				Description: "Indicate that this rule is shared via the Orgs API and read only.",
				Computed:    true,
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
	}
}

func (r *ZeroTrustGatewayPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = ResourceSchema(ctx)
}

func (r *ZeroTrustGatewayPolicyResource) ConfigValidators(_ context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{}
}
