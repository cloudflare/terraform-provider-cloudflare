// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r NotificationPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:   "UUID",
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"account_id": schema.StringAttribute{
				Description: "The account id",
				Required:    true,
			},
			"alert_type": schema.StringAttribute{
				Description: "Refers to which event will trigger a Notification dispatch. You can use the endpoint to get available alert types which then will give you a list of possible values.",
				Required:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("access_custom_certificate_expiration_type", "advanced_ddos_attack_l4_alert", "advanced_ddos_attack_l7_alert", "advanced_http_alert_error", "bgp_hijack_notification", "billing_usage_alert", "block_notification_block_removed", "block_notification_new_block", "block_notification_review_rejected", "brand_protection_alert", "brand_protection_digest", "clickhouse_alert_fw_anomaly", "clickhouse_alert_fw_ent_anomaly", "custom_ssl_certificate_event_type", "dedicated_ssl_certificate_event_type", "dos_attack_l4", "dos_attack_l7", "expiring_service_token_alert", "failing_logpush_job_disabled_alert", "fbm_auto_advertisement", "fbm_dosd_attack", "fbm_volumetric_attack", "health_check_status_notification", "hostname_aop_custom_certificate_expiration_type", "http_alert_edge_error", "http_alert_origin_error", "incident_alert", "load_balancing_health_alert", "load_balancing_pool_enablement_alert", "logo_match_alert", "magic_tunnel_health_check_event", "maintenance_event_notification", "mtls_certificate_store_certificate_expiration_type", "pages_event_alert", "radar_notification", "real_origin_monitoring", "scriptmonitor_alert_new_code_change_detections", "scriptmonitor_alert_new_hosts", "scriptmonitor_alert_new_malicious_hosts", "scriptmonitor_alert_new_malicious_scripts", "scriptmonitor_alert_new_malicious_url", "scriptmonitor_alert_new_max_length_resource_url", "scriptmonitor_alert_new_resources", "secondary_dns_all_primaries_failing", "secondary_dns_primaries_failing", "secondary_dns_zone_successfully_updated", "secondary_dns_zone_validation_warning", "sentinel_alert", "stream_live_notifications", "traffic_anomalies_alert", "tunnel_health_event", "tunnel_update_event", "universal_ssl_event_type", "web_analytics_metrics_update", "zone_aop_custom_certificate_expiration_type"),
				},
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether or not the Notification policy is enabled.",
				Computed:    true,
				Optional:    true,
				Default:     booldefault.StaticBool(true),
			},
			"mechanisms": schema.MapAttribute{
				Description: "List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.",
				Required:    true,
				ElementType: types.ListType{
					ElemType: types.StringType,
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the policy.",
				Required:    true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description for the Notification policy.",
				Optional:    true,
			},
			"filters": schema.SingleNestedAttribute{
				Description: "Optional filters that allow you to be alerted only on a subset of events for that alert type based on some criteria. This is only available for select alert types. See alert type documentation for more details.",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"actions": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"affected_asns": schema.ListAttribute{
						Description: "Used for configuring radar_notification",
						Optional:    true,
						ElementType: types.StringType,
					},
					"affected_components": schema.ListAttribute{
						Description: "Used for configuring incident_alert. A list of identifiers for each component to monitor.",
						Optional:    true,
						ElementType: types.StringType,
					},
					"affected_locations": schema.ListAttribute{
						Description: "Used for configuring radar_notification",
						Optional:    true,
						ElementType: types.StringType,
					},
					"airport_code": schema.ListAttribute{
						Description: "Used for configuring maintenance_event_notification",
						Optional:    true,
						ElementType: types.StringType,
					},
					"alert_trigger_preferences": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"alert_trigger_preferences_value": schema.ListAttribute{
						Description: "Used for configuring magic_tunnel_health_check_event",
						Optional:    true,
						ElementType: types.StringType,
					},
					"enabled": schema.ListAttribute{
						Description: "Used for configuring load_balancing_pool_enablement_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"environment": schema.ListAttribute{
						Description: "Used for configuring pages_event_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"event": schema.ListAttribute{
						Description: "Used for configuring pages_event_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"event_source": schema.ListAttribute{
						Description: "Used for configuring load_balancing_health_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"event_type": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"group_by": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"health_check_id": schema.ListAttribute{
						Description: "Used for configuring health_check_status_notification",
						Optional:    true,
						ElementType: types.StringType,
					},
					"incident_impact": schema.ListAttribute{
						Description: "Used for configuring incident_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"input_id": schema.ListAttribute{
						Description: "Used for configuring stream_live_notifications",
						Optional:    true,
						ElementType: types.StringType,
					},
					"limit": schema.ListAttribute{
						Description: "Used for configuring billing_usage_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"logo_tag": schema.ListAttribute{
						Description: "Used for configuring logo_match_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"megabits_per_second": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"new_health": schema.ListAttribute{
						Description: "Used for configuring load_balancing_health_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"new_status": schema.ListAttribute{
						Description: "Used for configuring tunnel_health_event",
						Optional:    true,
						ElementType: types.StringType,
					},
					"packets_per_second": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"pool_id": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"product": schema.ListAttribute{
						Description: "Used for configuring billing_usage_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"project_id": schema.ListAttribute{
						Description: "Used for configuring pages_event_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"protocol": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"query_tag": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"requests_per_second": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"selectors": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"services": schema.ListAttribute{
						Description: "Used for configuring clickhouse_alert_fw_ent_anomaly",
						Optional:    true,
						ElementType: types.StringType,
					},
					"slo": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"status": schema.ListAttribute{
						Description: "Used for configuring health_check_status_notification",
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_hostname": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_ip": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_zone_name": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"traffic_exclusions": schema.ListAttribute{
						Description: "Used for configuring traffic_anomalies_alert",
						Optional:    true,
						ElementType: types.StringType,
					},
					"tunnel_id": schema.ListAttribute{
						Description: "Used for configuring tunnel_health_event",
						Optional:    true,
						ElementType: types.StringType,
					},
					"tunnel_name": schema.ListAttribute{
						Description: "Used for configuring magic_tunnel_health_check_event",
						Optional:    true,
						ElementType: types.StringType,
					},
					"where": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
					"zones": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}
