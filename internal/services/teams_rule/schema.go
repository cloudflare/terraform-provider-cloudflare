package teams_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ResourceSchema returns the Terraform schema for the deprecated cloudflare_teams_rule resource.
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
				Description: "The action to preform when the associated traffic, identity, and device posture expressions are either absent or evaluate to `true`.\nAvailable values: \"on\", \"off\", \"allow\", \"block\", \"scan\", \"noscan\", \"safesearch\", \"ytrestricted\", \"isolate\", \"noisolate\", \"override\", \"l4_override\", \"egress\", \"resolve\", \"quarantine\", \"redirect\".",
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
				Description: "The name of the rule.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "The description of the rule.",
				Optional:    true,
			},
			"precedence": schema.Int64Attribute{
				Description: "Precedence sets the order of your rules. Lower values indicate higher precedence. At each processing phase, applicable rules are evaluated in ascending order of this value. Refer to [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform) docs on how to manage precedence via Terraform.",
				Optional:    true,
			},
			"filters": schema.ListAttribute{
				Description: "The protocol or layer to evaluate the traffic, identity, and device posture expressions.",
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
			"schedule": schema.SingleNestedAttribute{
				Description: "The schedule for activating DNS policies. This does not apply to HTTP or network policies.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"fri": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Fridays, in increasing order from 00:00-24:00.  If this parameter is omitted, the rule will be deactivated on Fridays.",
						Optional:    true,
					},
					"mon": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Mondays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Mondays.",
						Optional:    true,
					},
					"sat": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Saturdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Saturdays.",
						Optional:    true,
					},
					"sun": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Sundays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Sundays.",
						Optional:    true,
					},
					"thu": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Thursdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Thursdays.",
						Optional:    true,
					},
					"time_zone": schema.StringAttribute{
						Description: "The timezone for the schedule.",
						Optional:    true,
					},
					"tue": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Tuesdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Tuesdays.",
						Optional:    true,
					},
					"wed": schema.StringAttribute{
						Description: "The time intervals when the rule will be active on Wednesdays, in increasing order from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on Wednesdays.",
						Optional:    true,
					},
				},
			},
			"device_posture": schema.StringAttribute{
				Description: "The device posture expression to be matched.",
				Optional:    true,
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the rule is enabled.",
				Optional:    true,
				Computed:    true,
			},
			"identity": schema.StringAttribute{
				Description: "The identity expression to be matched.",
				Optional:    true,
				Computed:    true,
			},
			"traffic": schema.StringAttribute{
				Description: "The traffic expression to be matched.",
				Optional:    true,
				Computed:    true,
			},
			"expiration": schema.SingleNestedAttribute{
				Description: "The expiration settings for the rule.",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"expires_at": schema.StringAttribute{
						Description: "The expiration date and time for the rule.",
						Optional:    true,
						Computed:    true,
					},
					"duration": schema.Int64Attribute{
						Description: "The duration in seconds for the rule to expire.",
						Optional:    true,
					},
					"expired": schema.BoolAttribute{
						Description: "Whether the rule has expired.",
						Computed:    true,
					},
				},
			},
			"rule_settings": schema.SingleNestedAttribute{
				Description: "The rule settings for the policy.",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"add_headers": schema.MapAttribute{
						Description: "Add custom headers to allowed requests in the form of key-value pairs.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"allow_child_bypass": schema.BoolAttribute{
						Description: "Allow child bypass.",
						Optional:    true,
					},
					"audit_ssh": schema.SingleNestedAttribute{
						Description: "Audit SSH settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"command_logging": schema.BoolAttribute{
								Description: "Enable command logging.",
								Optional:    true,
							},
						},
					},
					"biso_admin_controls": schema.SingleNestedAttribute{
						Description: "BISO admin controls.",
						Optional:    true,
						Computed:    true,
						Attributes: map[string]schema.Attribute{
							"copy": schema.StringAttribute{
								Description: "Copy setting.",
								Optional:    true,
							},
							"dcp": schema.BoolAttribute{
								Description: "DCP setting.",
								Computed:    true,
							},
							"dd": schema.BoolAttribute{
								Description: "DD setting.",
								Computed:    true,
							},
							"dk": schema.BoolAttribute{
								Description: "DK setting.",
								Computed:    true,
							},
							"download": schema.StringAttribute{
								Description: "Download setting.",
								Optional:    true,
							},
							"du": schema.BoolAttribute{
								Description: "DU setting.",
								Computed:    true,
							},
							"keyboard": schema.StringAttribute{
								Description: "Keyboard setting.",
								Optional:    true,
							},
							"paste": schema.StringAttribute{
								Description: "Paste setting.",
								Optional:    true,
							},
							"printing": schema.StringAttribute{
								Description: "Printing setting.",
								Optional:    true,
							},
							"upload": schema.StringAttribute{
								Description: "Upload setting.",
								Optional:    true,
							},
							"version": schema.StringAttribute{
								Description: "Version setting.",
								Computed:    true,
							},
						},
					},
					"block_page": schema.SingleNestedAttribute{
						Description: "Block page settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"target_uri": schema.StringAttribute{
								Description: "The target URI for the block page.",
								Required:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Include context in the block page.",
								Optional:    true,
							},
						},
					},
					"block_page_enabled": schema.BoolAttribute{
						Description: "Whether the block page is enabled.",
						Optional:    true,
					},
					"block_reason": schema.StringAttribute{
						Description: "The reason for blocking.",
						Optional:    true,
					},
					"bypass_parent_rule": schema.BoolAttribute{
						Description: "Whether to bypass parent rule.",
						Optional:    true,
					},
					"check_session": schema.SingleNestedAttribute{
						Description: "Check session settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"duration": schema.StringAttribute{
								Description: "The duration for the session check.",
								Optional:    true,
							},
							"enforce": schema.BoolAttribute{
								Description: "Whether to enforce the session check.",
								Optional:    true,
							},
						},
					},
					"dns_resolvers": schema.SingleNestedAttribute{
						Description: "DNS resolver settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.ListNestedAttribute{
								Description: "IPv4 DNS resolvers.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "The IP address of the DNS resolver.",
											Required:    true,
										},
										"port": schema.Int64Attribute{
											Description: "The port number for the DNS resolver.",
											Optional:    true,
										},
										"route_through_private_network": schema.BoolAttribute{
											Description: "Whether to route through private network.",
											Optional:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "The virtual network ID.",
											Optional:    true,
										},
									},
								},
							},
							"ipv6": schema.ListNestedAttribute{
								Description: "IPv6 DNS resolvers.",
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Description: "The IP address of the DNS resolver.",
											Required:    true,
										},
										"port": schema.Int64Attribute{
											Description: "The port number for the DNS resolver.",
											Optional:    true,
										},
										"route_through_private_network": schema.BoolAttribute{
											Description: "Whether to route through private network.",
											Optional:    true,
										},
										"vnet_id": schema.StringAttribute{
											Description: "The virtual network ID.",
											Optional:    true,
										},
									},
								},
							},
						},
					},
					"egress": schema.SingleNestedAttribute{
						Description: "Egress settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ipv4": schema.StringAttribute{
								Description: "IPv4 egress setting.",
								Optional:    true,
							},
							"ipv4_fallback": schema.StringAttribute{
								Description: "IPv4 fallback setting.",
								Optional:    true,
							},
							"ipv6": schema.StringAttribute{
								Description: "IPv6 egress setting.",
								Optional:    true,
							},
						},
					},
					"ignore_cname_category_matches": schema.BoolAttribute{
						Description: "Whether to ignore CNAME category matches.",
						Optional:    true,
					},
					"insecure_disable_dnssec_validation": schema.BoolAttribute{
						Description: "Whether to disable DNSSEC validation.",
						Optional:    true,
					},
					"ip_categories": schema.BoolAttribute{
						Description: "Whether to enable IP categories.",
						Optional:    true,
					},
					"ip_indicator_feeds": schema.BoolAttribute{
						Description: "Whether to enable IP indicator feeds.",
						Optional:    true,
					},
					"l4override": schema.SingleNestedAttribute{
						Description: "L4 override settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"ip": schema.StringAttribute{
								Description: "The IP address for L4 override.",
								Optional:    true,
							},
							"port": schema.Int64Attribute{
								Description: "The port number for L4 override.",
								Optional:    true,
							},
						},
					},
					"notification_settings": schema.SingleNestedAttribute{
						Description: "Notification settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether notifications are enabled.",
								Optional:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Whether to include context in notifications.",
								Optional:    true,
							},
							"msg": schema.StringAttribute{
								Description: "The notification message.",
								Optional:    true,
							},
							"support_url": schema.StringAttribute{
								Description: "The support URL for notifications.",
								Optional:    true,
							},
						},
					},
					"override_host": schema.StringAttribute{
						Description: "The override host.",
						Optional:    true,
					},
					"override_ips": schema.ListAttribute{
						Description: "The override IPs.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"payload_log": schema.SingleNestedAttribute{
						Description: "Payload log settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "Whether payload logging is enabled.",
								Optional:    true,
							},
						},
					},
					"quarantine": schema.SingleNestedAttribute{
						Description: "Quarantine settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"file_types": schema.ListAttribute{
								Description: "The file types to quarantine.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"redirect": schema.SingleNestedAttribute{
						Description: "Redirect settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"target_uri": schema.StringAttribute{
								Description: "The target URI for redirection.",
								Required:    true,
							},
							"include_context": schema.BoolAttribute{
								Description: "Whether to include context in redirection.",
								Optional:    true,
							},
							"preserve_path_and_query": schema.BoolAttribute{
								Description: "Whether to preserve path and query in redirection.",
								Optional:    true,
							},
						},
					},
					"resolve_dns_internally": schema.SingleNestedAttribute{
						Description: "Resolve DNS internally settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"fallback": schema.StringAttribute{
								Description: "The fallback DNS resolver.",
								Optional:    true,
							},
							"view_id": schema.StringAttribute{
								Description: "The view ID for DNS resolution.",
								Optional:    true,
							},
						},
					},
					"resolve_dns_through_cloudflare": schema.BoolAttribute{
						Description: "Whether to resolve DNS through Cloudflare.",
						Optional:    true,
					},
					"untrusted_cert": schema.SingleNestedAttribute{
						Description: "Untrusted certificate settings.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"action": schema.StringAttribute{
								Description: "The action for untrusted certificates.",
								Optional:    true,
							},
						},
					},
				},
			},
			"created_at": schema.StringAttribute{
				Description: "The creation timestamp.",
				Computed:    true,
			},
			"deleted_at": schema.StringAttribute{
				Description: "The deletion timestamp.",
				Computed:    true,
			},
			"updated_at": schema.StringAttribute{
				Description: "The update timestamp.",
				Computed:    true,
			},
			"version": schema.Int64Attribute{
				Description: "The version of the rule.",
				Computed:    true,
			},
			"warning_status": schema.StringAttribute{
				Description: "The warning status of the rule.",
				Computed:    true,
			},
		},
	}
}

// Duplicate Schema and ConfigValidators implementations were removed to avoid
// redeclaration errors. The implementations reside in `resource.go`.
