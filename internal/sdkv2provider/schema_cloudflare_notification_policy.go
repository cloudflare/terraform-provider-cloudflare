package sdkv2provider

import (
	"fmt"
	"sort"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

var notificationAffectedComponents = map[string]string{
	"Access":                              "w4k8yvhfb3vp",
	"Always Online":                       "xm3cq0t85y10",
	"Analytics":                           "4c231tkdlpcl",
	"API":                                 "g4tb35rs9yw7",
	"API Shield":                          "g9yx473yjk9t",
	"Apps Marketplace":                    "g9dgngpcdt1x",
	"Argo Smart Routing":                  "z9w398bsjvnq",
	"Audit Logs":                          "2469qcw8rvjp",
	"Authoritative DNS":                   "dp8ppfycqxcs",
	"Billing":                             "ll1x88wwz4fq",
	"Bring Your Own IP (BYOIP)":           "4msl4k5wdcbv",
	"Browser Isolation":                   "q0dfbn0p6hyt",
	"Bot Management":                      "s0991jwsqllx",
	"Cache Reserve":                       "3q1jnbdbn845",
	"CDN/Cache":                           "5wnz34mhfhrk",
	"CDN Cache Purge":                     "fbvx0hxhhdj0",
	"Challenge Platform":                  "x0tkn0hzrtw7",
	"Cloud Access Security Broker (CASB)": "h2p0jj4ltvcq",
	"Community Site":                      "qgh1bfr4hxrl",
	"Data Loss Prevention (DLP)":          "rppy995xymxv",
	"Dashboard":                           "3sq3s4d20ywk",
	"Developer's Site":                    "rzcwwk4rgb0w",
	"Digital Experience Monitoring (DEX)": "nmp96vgn1hpl",
	"Distributed Web Gateway":             "5scwd3vmnsyj",
	"DNS Root Servers":                    "4l9qztbt6rbj",
	"DNS Updates":                         "7j656z7tqk7f",
	"Durable Objects":                     "bty1yz6dhh0v",
	"Email Routing":                       "gjb0yzvvrpf6",
	"Ethereum Gateway":                    "7yyjz9qdsjbx",
	"Firewall":                            "d1r0plwsl5qb",
	"Gateway":                             "gyx2yygg7lmd",
	"Geo-Key Manager":                     "4tw744y7kfmw",
	"Image Resizing":                      "dw7t39j5syzl",
	"Images":                              "3lbj8lp3d750",
	"Infrastructure":                      "8qwmwg7ytljv",
	"Load Balancing and Monitoring":       "8sn2w5kyxfnp",
	"Lists":                               "tfnrx45s2b48",
	"Logs":                                "k0mgxrls5y1b",
	"Magic Firewall":                      "m1cm5tqpkqtm",
	"Magic Transit":                       "bjlxcss20fsl",
	"Magic WAN":                           "qgnxz00j1f2v",
	"Magic WAN Connector":                 "q3t6mnpmpgt8",
	"Marketing Site":                      "6239kkkfzfnf",
	"Mirage":                              "j9jl2gb9zywx",
	"Network":                             "4n0gb0kh02gf",
	"Notifications":                       "5xvn0m7tthlf",
	"Observatory":                         "mfty6kskddpf",
	"Page Shield":                         "hk5dqm69klkp",
	"Pages":                               "vgxj684rcw7t",
	"R2":                                  "hb7g5sq2zz0h",
	"Radar":                               "0fw91jq1bzxx",
	"Randomness Beacon":                   "yd553hxj8dbj",
	"Recursive DNS":                       "8w536gxk7dvq",
	"Registrar":                           "kn2xkt469vyh",
	"Registration Data Access Protocol (RDAP)": "vsj8h17tq59r",
	"Security Center":                          "18qkc83zzmxb",
	"Snippets":                                 "570kfpd0dgg7",
	"Spectrum":                                 "6dd6ssg7plt0",
	"Speed Optimizations":                      "fcx388ss9k9x",
	"Stream":                                   "47xg28c02lnk",
	"SSL Certificate Provisioning":             "cghykwlwsmn5",
	"SSL for SaaS Provisioning":                "9p2qlpt19nqb",
	"Support Site":                             "jzcwkvrc4w4q",
	"Time Services":                            "ggcvp9h5v6rv",
	"Trace":                                    "f0jjgwcxtmk8",
	"Tunnel":                                   "y98zlwj1d7zh",
	"Turnstile":                                "m4jywscr0n0k",
	"Waiting Room":                             "9c7cbxnhk1dq",
	"WARP":                                     "k04jkcpzxn94",
	"Web Analytics":                            "qt59p3cr1grx",
	"Workers":                                  "57srcl8zcn7c",
	"Workers Preview":                          "wjvmzdf21d4l",
	"Workers KV":                               "tmh50tx2nprs",
	"Zaraz":                                    "qgt2kv10g1yn",
	"Zero Trust":                               "kf0ktv29xrfy",
	"Zero Trust Dashboard":                     "276xk3r83js7",
	"Zone Versioning":                          "4tv81hqpt2jt",
}

func affectedComponentKeys() []string {
	var keys []string

	for key := range notificationAffectedComponents {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

var notificationPolicyAlertTypes = []string{
	"advanced_http_alert_error",
	"access_custom_certificate_expiration_type",
	"advanced_ddos_attack_l4_alert",
	"advanced_ddos_attack_l7_alert",
	"bgp_hijack_notification",
	"billing_usage_alert",
	"block_notification_block_removed",
	"block_notification_new_block",
	"block_notification_review_rejected",
	"brand_protection_alert",
	"brand_protection_digest",
	"clickhouse_alert_fw_anomaly",
	"clickhouse_alert_fw_ent_anomaly",
	"custom_ssl_certificate_event_type",
	"dedicated_ssl_certificate_event_type",
	"dos_attack_l4",
	"dos_attack_l7",
	"expiring_service_token_alert",
	"failing_logpush_job_disabled_alert",
	"fbm_auto_advertisement",
	"fbm_dosd_attack",
	"fbm_volumetric_attack",
	"health_check_status_notification",
	"hostname_aop_custom_certificate_expiration_type",
	"http_alert_edge_error",
	"http_alert_origin_error",
	"incident_alert",
	"load_balancing_health_alert",
	"load_balancing_pool_enablement_alert",
	"logo_match_alert",
	"magic_tunnel_health_check_event",
	"maintenance_event_notification",
	"mtls_certificate_store_certificate_expiration_type",
	"pages_event_alert",
	"radar_notification",
	"real_origin_monitoring",
	"scriptmonitor_alert_new_code_change_detections",
	"scriptmonitor_alert_new_hosts",
	"scriptmonitor_alert_new_malicious_hosts",
	"scriptmonitor_alert_new_malicious_scripts",
	"scriptmonitor_alert_new_malicious_url",
	"scriptmonitor_alert_new_max_length_resource_url",
	"scriptmonitor_alert_new_resources",
	"secondary_dns_all_primaries_failing",
	"secondary_dns_primaries_failing",
	"secondary_dns_zone_successfully_updated",
	"secondary_dns_zone_validation_warning",
	"sentinel_alert",
	"stream_live_notifications",
	"traffic_anomalies_alert",
	"tunnel_health_event",
	"tunnel_update_event",
	"universal_ssl_event_type",
	"web_analytics_metrics_update",
	"weekly_account_overview",
	"workers_alert",
	"zone_aop_custom_certificate_expiration_type",
}

var notificationPolicyIncidentImpactLevels = []string{
	"INCIDENT_IMPACT_NONE",
	"INCIDENT_IMPACT_MINOR",
	"INCIDENT_IMPACT_MAJOR",
	"INCIDENT_IMPACT_CRITICAL",
}

func resourceCloudflareNotificationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the notification policy.",
		},
		"description": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Description of the notification policy.",
		},
		"enabled": {
			Type:        schema.TypeBool,
			Required:    true,
			Description: "The status of the notification policy.",
		},
		"alert_type": {
			Type:         schema.TypeString,
			Required:     true,
			ValidateFunc: validation.StringInSlice(notificationPolicyAlertTypes, false),
			Description:  fmt.Sprintf("The event type that will trigger the dispatch of a notification. See the developer documentation for descriptions of [available alert types](https://developers.cloudflare.com/fundamentals/notifications/notification-available/). %s", renderAvailableDocumentationValuesStringSlice(notificationPolicyAlertTypes)),
		},
		"filters": notificationPolicyFilterSchema(),
		"created": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "When the notification policy was created.",
		},
		"modified": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "When the notification policy was last modified.",
		},
		"email_integration": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        mechanismData,
			Description: "The email ID to which the notification should be dispatched.",
		},
		"webhooks_integration": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        mechanismData,
			Description: "The unique ID of a configured webhooks endpoint to which the notification should be dispatched.",
		},
		"pagerduty_integration": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        mechanismData,
			Description: "The unique ID of a configured pagerduty endpoint to which the notification should be dispatched.",
		},
	}
}

var mechanismData = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}

func notificationPolicyFilterSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Optional:    true,
		MaxItems:    1,
		Description: "An optional nested block of filters that applies to the selected `alert_type`. A key-value map that specifies the type of filter and the values to match against (refer to the alert type block for available fields).",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"actions": {
					Type:        schema.TypeSet,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Optional:    true,
					Description: "Targeted actions for alert.",
				},
				"airport_code": {
					Type:        schema.TypeSet,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Optional:    true,
					Description: "Filter on Points of Presence.",
				},
				"affected_components": {
					Type:        schema.TypeSet,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Optional:    true,
					Description: fmt.Sprintf("Affected components for alert. %s", renderAvailableDocumentationValuesStringSlice(affectedComponentKeys())),
				},
				"status": {
					Type:        schema.TypeSet,
					Elem:        &schema.Schema{Type: schema.TypeString},
					Optional:    true,
					Description: "Status to alert on.",
				},
				"health_check_id": {
					Type:         schema.TypeSet,
					Elem:         &schema.Schema{Type: schema.TypeString},
					Optional:     true,
					RequiredWith: []string{"filters.0.status"},
					Description:  "Identifier health check.",
				},
				"zones": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "A list of zone identifiers.",
				},
				"services": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
				"product": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: fmt.Sprintf("Product name. %s", renderAvailableDocumentationValuesStringSlice([]string{"worker_requests", "worker_durable_objects_requests", "worker_durable_objects_duration", "worker_durable_objects_data_transfer", "worker_durable_objects_stored_data", "worker_durable_objects_storage_deletes", "worker_durable_objects_storage_writes", "worker_durable_objects_storage_reads"})),
				},
				"limit": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "A numerical limit. Example: `100`",
				},
				"enabled": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "State of the pool to alert on.",
				},
				"pool_id": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Load balancer pool identifier.",
				},
				"slo": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "A numerical limit. Example: `99.9`.",
				},
				"where": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Filter for alert.",
				},
				"group_by": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Alert grouping.",
				},
				"alert_trigger_preferences": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Alert trigger preferences. Example: `slo`.",
				},
				"requests_per_second": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Requests per second threshold for dos alert.",
				},
				"target_zone_name": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Target domain to alert on.",
				},
				"target_hostname": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Target host to alert on for dos.",
				},
				"packets_per_second": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Packets per second threshold for dos alert.",
				},
				"protocol": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Protocol to alert on for dos.",
				},
				"project_id": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Identifier of pages project.",
				},
				"environment": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
					Description: fmt.Sprintf("Environment of pages. %s", renderAvailableDocumentationValuesStringSlice([]string{
						"ENVIRONMENT_PREVIEW",
						"ENVIRONMENT_PRODUCTION",
					})),
				},
				"event": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
					Description: fmt.Sprintf("Pages event to alert. %s", renderAvailableDocumentationValuesStringSlice([]string{
						"EVENT_DEPLOYMENT_STARTED",
						"EVENT_DEPLOYMENT_FAILED",
						"EVENT_DEPLOYMENT_SUCCESS",
					})),
				},
				"event_source": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Source configuration to alert on for pool or origin.",
				},
				"new_health": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Health status to alert on for pool or origin.",
				},
				"input_id": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Stream input id to alert on.",
				},
				"event_type": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Stream event type to alert on.",
				},
				"megabits_per_second": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Megabits per second threshold for dos alert.",
				},
				"incident_impact": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type:         schema.TypeString,
						ValidateFunc: validation.StringInSlice(notificationPolicyIncidentImpactLevels, false),
					},
					Optional:    true,
					Description: fmt.Sprintf("The incident impact level that will trigger the dispatch of a notification. %s", renderAvailableDocumentationValuesStringSlice(notificationPolicyIncidentImpactLevels)),
				},
				"new_status": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Tunnel health status to alert on.",
				},
				"selectors": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Selectors for alert. Valid options depend on the alert type.",
				},
				"tunnel_id": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "Tunnel IDs to alert on.",
				},
			},
		},
	}
}
