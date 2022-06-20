package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareNotificationPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: "The account identifier to target for the resource.",
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
			ValidateFunc: validation.StringInSlice([]string{"billing_usage_alert", "health_check_status_notification", "g6_pool_toggle_alert", "real_origin_monitoring", "universal_ssl_event_type", "bgp_hijack_notification", "http_alert_origin_error", "workers_alert", "weekly_account_overview"}, false),
			Description:  fmt.Sprintf("The event type that will trigger the dispatch of a notification. See the developer documentation for descriptions of [available alert types](https://developers.cloudflare.com/fundamentals/notifications/notification-available/) %s", renderAvailableDocumentationValuesStringSlice([]string{"billing_usage_alert", "health_check_status_notification", "g6_pool_toggle_alert", "real_origin_monitoring", "universal_ssl_event_type", "bgp_hijack_notification", "http_alert_origin_error", "workers_alert", "weekly_account_overview"})),
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
					Optional:     true,
					RequiredWith: []string{"filters.0.enabled"},
					Description:  "Load balancer pool identifier.",
				},
				"slo": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional:    true,
					Description: "A numerical limit. Example: `99.9`",
				},
			},
		},
	}
}
