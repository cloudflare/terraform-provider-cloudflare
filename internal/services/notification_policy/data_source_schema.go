// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package notification_policy

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
				Computed:    true,
			},
			"alert_interval": schema.StringAttribute{
				Description: "Optional specification of how often to re-alert from the same incident, not support on all alert types.",
				Computed:    true,
				Optional:    true,
			},
			"alert_type": schema.StringAttribute{
				Description: "Refers to which event will trigger a Notification dispatch. You can use the endpoint to get available alert types which then will give you a list of possible values.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("access_custom_certificate_expiration_type", "advanced_ddos_attack_l4_alert", "advanced_ddos_attack_l7_alert", "advanced_http_alert_error", "bgp_hijack_notification", "billing_usage_alert", "block_notification_block_removed", "block_notification_new_block", "block_notification_review_rejected", "brand_protection_alert", "brand_protection_digest", "clickhouse_alert_fw_anomaly", "clickhouse_alert_fw_ent_anomaly", "cloudforce_one_request_notification", "custom_analytics", "custom_ssl_certificate_event_type", "dedicated_ssl_certificate_event_type", "device_connectivity_anomaly_alert", "dos_attack_l4", "dos_attack_l7", "expiring_service_token_alert", "failing_logpush_job_disabled_alert", "fbm_auto_advertisement", "fbm_dosd_attack", "fbm_volumetric_attack", "health_check_status_notification", "hostname_aop_custom_certificate_expiration_type", "http_alert_edge_error", "http_alert_origin_error", "incident_alert", "load_balancing_health_alert", "load_balancing_pool_enablement_alert", "logo_match_alert", "magic_tunnel_health_check_event", "magic_wan_tunnel_health", "maintenance_event_notification", "mtls_certificate_store_certificate_expiration_type", "pages_event_alert", "radar_notification", "real_origin_monitoring", "scriptmonitor_alert_new_code_change_detections", "scriptmonitor_alert_new_hosts", "scriptmonitor_alert_new_malicious_hosts", "scriptmonitor_alert_new_malicious_scripts", "scriptmonitor_alert_new_malicious_url", "scriptmonitor_alert_new_max_length_resource_url", "scriptmonitor_alert_new_resources", "secondary_dns_all_primaries_failing", "secondary_dns_primaries_failing", "secondary_dns_warning", "secondary_dns_zone_successfully_updated", "secondary_dns_zone_validation_warning", "sentinel_alert", "stream_live_notifications", "synthetic_test_latency_alert", "synthetic_test_low_availability_alert", "traffic_anomalies_alert", "tunnel_health_event", "tunnel_update_event", "universal_ssl_event_type", "web_analytics_metrics_update", "zone_aop_custom_certificate_expiration_type"),
				},
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Description: "Optional description for the Notification policy.",
				Computed:    true,
				Optional:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether or not the Notification policy is enabled.",
				Computed:    true,
			},
			"filters": schema.SingleNestedAttribute{
				Description: "Optional filters that allow you to be alerted only on a subset of events for that alert type based on some criteria. This is only available for select alert types. See alert type documentation for more details.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"actions": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"affected_asns": schema.ListAttribute{
						Description: "Used for configuring radar_notification",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"affected_components": schema.ListAttribute{
						Description: "Used for configuring incident_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"affected_locations": schema.ListAttribute{
						Description: "Used for configuring radar_notification",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"airport_code": schema.ListAttribute{
						Description: "Used for configuring maintenance_event_notification",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"alert_trigger_preferences": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"alert_trigger_preferences_value": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"enabled": schema.ListAttribute{
						Description: "Used for configuring load_balancing_pool_enablement_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"environment": schema.ListAttribute{
						Description: "Used for configuring pages_event_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"event": schema.ListAttribute{
						Description: "Used for configuring pages_event_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"event_source": schema.ListAttribute{
						Description: "Used for configuring load_balancing_health_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"event_type": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"group_by": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"health_check_id": schema.ListAttribute{
						Description: "Used for configuring health_check_status_notification",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"incident_impact": schema.ListAttribute{
						Description: "Used for configuring incident_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"input_id": schema.ListAttribute{
						Description: "Used for configuring stream_live_notifications",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"limit": schema.ListAttribute{
						Description: "Used for configuring billing_usage_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"logo_tag": schema.ListAttribute{
						Description: "Used for configuring logo_match_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"megabits_per_second": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"new_health": schema.ListAttribute{
						Description: "Used for configuring load_balancing_health_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"new_status": schema.ListAttribute{
						Description: "Used for configuring tunnel_health_event",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"packets_per_second": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"pool_id": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"product": schema.ListAttribute{
						Description: "Used for configuring billing_usage_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"project_id": schema.ListAttribute{
						Description: "Used for configuring pages_event_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"protocol": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"query_tag": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"requests_per_second": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"selectors": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"services": schema.ListAttribute{
						Description: "Used for configuring clickhouse_alert_fw_ent_anomaly",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"slo": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"status": schema.ListAttribute{
						Description: "Used for configuring health_check_status_notification",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_hostname": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_ip": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l4_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"target_zone_name": schema.ListAttribute{
						Description: "Used for configuring advanced_ddos_attack_l7_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"traffic_exclusions": schema.ListAttribute{
						Description: "Used for configuring traffic_anomalies_alert",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"tunnel_id": schema.ListAttribute{
						Description: "Used for configuring tunnel_health_event",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"tunnel_name": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"where": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"zones": schema.ListAttribute{
						Description: "Usage depends on specific alert type",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
				},
			},
			"mechanisms": schema.MapAttribute{
				Description: "List of IDs that will be used when dispatching a notification. IDs for email type will be the email address.",
				Computed:    true,
				Optional:    true,
				ElementType: types.ListType{
					ElemType: jsontypes.NewNormalizedNull().Type(ctx),
				},
			},
			"modified": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Description: "Name of the policy.",
				Computed:    true,
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
