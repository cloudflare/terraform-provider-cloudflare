// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = &NotificationPolicyDataSource{}
var _ datasource.DataSourceWithValidateConfig = &NotificationPolicyDataSource{}

func (r NotificationPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "The account id",
				Optional:    true,
			},
			"policy_id": schema.StringAttribute{
				Description: "The unique identifier of a notification policy",
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Description: "The unique identifier of a notification policy",
				Optional:    true,
			},
			"alert_type": schema.StringAttribute{
				Description: "Refers to which event will trigger a Notification dispatch. You can use the endpoint to get available alert types which then will give you a list of possible values.",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("access_custom_certificate_expiration_type", "advanced_ddos_attack_l4_alert", "advanced_ddos_attack_l7_alert", "advanced_http_alert_error", "bgp_hijack_notification", "billing_usage_alert", "block_notification_block_removed", "block_notification_new_block", "block_notification_review_rejected", "brand_protection_alert", "brand_protection_digest", "clickhouse_alert_fw_anomaly", "clickhouse_alert_fw_ent_anomaly", "custom_ssl_certificate_event_type", "dedicated_ssl_certificate_event_type", "dos_attack_l4", "dos_attack_l7", "expiring_service_token_alert", "failing_logpush_job_disabled_alert", "fbm_auto_advertisement", "fbm_dosd_attack", "fbm_volumetric_attack", "health_check_status_notification", "hostname_aop_custom_certificate_expiration_type", "http_alert_edge_error", "http_alert_origin_error", "incident_alert", "load_balancing_health_alert", "load_balancing_pool_enablement_alert", "logo_match_alert", "magic_tunnel_health_check_event", "maintenance_event_notification", "mtls_certificate_store_certificate_expiration_type", "pages_event_alert", "radar_notification", "real_origin_monitoring", "scriptmonitor_alert_new_code_change_detections", "scriptmonitor_alert_new_hosts", "scriptmonitor_alert_new_malicious_hosts", "scriptmonitor_alert_new_malicious_scripts", "scriptmonitor_alert_new_malicious_url", "scriptmonitor_alert_new_max_length_resource_url", "scriptmonitor_alert_new_resources", "secondary_dns_all_primaries_failing", "secondary_dns_primaries_failing", "secondary_dns_zone_successfully_updated", "secondary_dns_zone_validation_warning", "sentinel_alert", "stream_live_notifications", "traffic_anomalies_alert", "tunnel_health_event", "tunnel_update_event", "universal_ssl_event_type", "web_analytics_metrics_update", "zone_aop_custom_certificate_expiration_type"),
				},
			},
			"created": schema.StringAttribute{
				Optional: true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description for the Notification policy.",
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether or not the Notification policy is enabled.",
				Computed:    true,
				Optional:    true,
			},
			"filters": schema.SingleNestedAttribute{
				Description: "Optional filters that allow you to be alerted only on a subset of events for that alert type based on some criteria. This is only available for select alert types. See alert type documentation for more details.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"actions": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"affected_asns": schema.StringAttribute{
						Description: "Used for configuring radar_notification",
						Optional:    true,
					},
					"affected_components": schema.StringAttribute{
						Description: "Used for configuring incident_alert. A list of identifiers for each component to monitor.",
						Optional:    true,
					},
					"affected_locations": schema.StringAttribute{
						Description: "Used for configuring radar_notification",
						Optional:    true,
					},
					"airport_code": schema.StringAttribute{
						Description: "Used for configuring maintenance_event_notification",
						Optional:    true,
					},
					"alert_trigger_preferences": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"alert_trigger_preferences_value": schema.StringAttribute{
						Description: "Used for configuring magic_tunnel_health_check_event",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("99.0", "98.0", "97.0"),
						},
					},
					"enabled": schema.StringAttribute{
						Description: "Used for configuring load_balancing_pool_enablement_alert",
						Optional:    true,
					},
					"environment": schema.StringAttribute{
						Description: "Used for configuring pages_event_alert",
						Optional:    true,
					},
					"event": schema.StringAttribute{
						Description: "Used for configuring pages_event_alert",
						Optional:    true,
					},
					"event_source": schema.StringAttribute{
						Description: "Used for configuring load_balancing_health_alert",
						Optional:    true,
					},
					"event_type": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"group_by": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"health_check_id": schema.StringAttribute{
						Description: "Used for configuring health_check_status_notification",
						Optional:    true,
					},
					"incident_impact": schema.StringAttribute{
						Description: "Used for configuring incident_alert",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("INCIDENT_IMPACT_NONE", "INCIDENT_IMPACT_MINOR", "INCIDENT_IMPACT_MAJOR", "INCIDENT_IMPACT_CRITICAL"),
						},
					},
					"input_id": schema.StringAttribute{
						Description: "Used for configuring stream_live_notifications",
						Optional:    true,
					},
					"limit": schema.StringAttribute{
						Description: "Used for configuring billing_usage_alert",
						Optional:    true,
					},
					"logo_tag": schema.StringAttribute{
						Description: "Used for configuring logo_match_alert",
						Optional:    true,
					},
					"megabits_per_second": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
					},
					"new_health": schema.StringAttribute{
						Description: "Used for configuring load_balancing_health_alert",
						Optional:    true,
					},
					"new_status": schema.StringAttribute{
						Description: "Used for configuring tunnel_health_event",
						Optional:    true,
					},
					"packets_per_second": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
					},
					"pool_id": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"product": schema.StringAttribute{
						Description: "Used for configuring billing_usage_alert",
						Optional:    true,
					},
					"project_id": schema.StringAttribute{
						Description: "Used for configuring pages_event_alert",
						Optional:    true,
					},
					"protocol": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
					},
					"query_tag": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"requests_per_second": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Optional:    true,
					},
					"selectors": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"services": schema.StringAttribute{
						Description: "Used for configuring clickhouse_alert_fw_ent_anomaly",
						Optional:    true,
					},
					"slo": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"status": schema.StringAttribute{
						Description: "Used for configuring health_check_status_notification",
						Optional:    true,
					},
					"target_hostname": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Optional:    true,
					},
					"target_ip": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
					},
					"target_zone_name": schema.StringAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Optional:    true,
					},
					"traffic_exclusions": schema.StringAttribute{
						Description: "Used for configuring traffic_anomalies_alert",
						Optional:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("security_events"),
						},
					},
					"tunnel_id": schema.StringAttribute{
						Description: "Used for configuring tunnel_health_event",
						Optional:    true,
					},
					"tunnel_name": schema.StringAttribute{
						Description: "Used for configuring magic_tunnel_health_check_event",
						Optional:    true,
					},
					"where": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
					"zones": schema.StringAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
					},
				},
			},
			"mechanisms": schema.StringAttribute{
				Description: "List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.",
				Optional:    true,
			},
			"modified": schema.StringAttribute{
				Optional: true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the policy.",
				Optional:    true,
			},
			"find_one_by": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description: "The account id",
						Required:    true,
					},
				},
			},
		},
	}
}

func (r *NotificationPolicyDataSource) ConfigValidators(ctx context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}

func (r *NotificationPolicyDataSource) ValidateConfig(ctx context.Context, req datasource.ValidateConfigRequest, resp *datasource.ValidateConfigResponse) {
}
