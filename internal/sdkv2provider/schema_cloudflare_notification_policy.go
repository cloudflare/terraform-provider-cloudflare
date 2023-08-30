package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

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
			Type:     schema.TypeString,
			Required: true,
			ValidateFunc: validation.StringInSlice([]string{
				"access_custom_certificate_expiration_type",
				"advanced_ddos_attack_l4_alert",
				"advanced_ddos_attack_l7_alert",
				"bgp_hijack_notification",
				"billing_usage_alert",
				"block_notification_block_removed",
				"block_notification_new_block",
				"block_notification_review_rejected",
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
				"load_balancing_health_alert",
				"load_balancing_pool_enablement_alert",
				"pages_event_alert",
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
				"tunnel_health_event",
				"tunnel_update_event",
				"universal_ssl_event_type",
				"web_analytics_metrics_update",
				"weekly_account_overview",
				"workers_alert",
				"zone_aop_custom_certificate_expiration_type",
			}, false),
			Description: fmt.Sprintf("The event type that will trigger the dispatch of a notification. See the developer documentation for descriptions of [available alert types](https://developers.cloudflare.com/fundamentals/notifications/notification-available/). %s", renderAvailableDocumentationValuesStringSlice([]string{
				"access_custom_certificate_expiration_type",
				"advanced_ddos_attack_l4_alert",
				"advanced_ddos_attack_l7_alert",
				"bgp_hijack_notification",
				"billing_usage_alert",
				"block_notification_block_removed",
				"block_notification_new_block",
				"block_notification_review_rejected",
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
				"load_balancing_health_alert",
				"load_balancing_pool_enablement_alert",
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
				"tunnel_health_event",
				"tunnel_update_event",
				"universal_ssl_event_type",
				"web_analytics_metrics_update",
				"weekly_account_overview",
				"workers_alert",
				"zone_aop_custom_certificate_expiration_type",
			})),
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
			Description: "The email id to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.",
		},
		"webhooks_integration": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        mechanismData,
			Description: "The unique id of a configured webhooks endpoint to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.",
		},
		"pagerduty_integration": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        mechanismData,
			Description: "The unique id of a configured pagerduty endpoint to which the notification should be dispatched. One of email, webhooks, or PagerDuty mechanisms is required.",
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
			},
		},
	}
}
