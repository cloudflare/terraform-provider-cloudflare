// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_member"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/address_map"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation_schema_validation_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_schema"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_schema_validation_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_smart_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_tiered_caching"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/bot_management"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/certificate_pack"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname_fallback_origin"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/d1_database"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_dnssec"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_address"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_catch_all"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/filter"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/firewall_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/healthcheck"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hostname_tls_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/keyless_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpull_retention"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_job"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_ownership_challenge"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_gre_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_ipsec_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_static_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/managed_headers"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/mtls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy_webhooks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/observatory_scheduled_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_ca_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/rate_limit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/spectrum_application"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/total_tls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/turnstile_widget"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/url_normalization_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/user_agent_blocking_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_event"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/web3_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/web_analytics_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/web_analytics_site"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_cron_trigger"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_custom_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_dispatch_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_secret"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_application"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_custom_page"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_identity_provider"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_key_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_mtls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_mtls_hostname_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_service_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_short_lived_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_tag"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_managed_networks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_profiles"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dex_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_custom_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_predefined_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dns_location"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_proxy_endpoint"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_local_domain_fallback"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_virtual_network"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_reserve"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_variants"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_hold"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_lockdown"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.ProviderWithConfigValidators = (*CloudflareProvider)(nil)

// CloudflareProvider defines the provider implementation.
type CloudflareProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CloudflareProviderModel describes the provider data model.
type CloudflareProviderModel struct {
	BaseURL        types.String `tfsdk:"base_url" json:"base_url"`
	APIToken       types.String `tfsdk:"api_token" json:"api_token"`
	APIKey         types.String `tfsdk:"api_key" json:"api_key"`
	APIEmail       types.String `tfsdk:"api_email" json:"api_email"`
	UserServiceKey types.String `tfsdk:"user_service_key" json:"user_service_key"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Description: "Set the base url that the provider connects to. This can be used for testing in other environments.",
				Optional:    true,
			},
			"api_token": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional: true,
			},
			"api_email": schema.StringAttribute{
				Optional: true,
			},
			"user_service_key": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (p *CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	// TODO(terraform): apiKey := os.Getenv("API_KEY")

	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if !data.APIToken.IsNull() {
		opts = append(opts, option.WithAPIToken(data.APIToken.ValueString()))
	}
	if !data.APIKey.IsNull() {
		opts = append(opts, option.WithAPIKey(data.APIKey.ValueString()))
	}
	if !data.APIEmail.IsNull() {
		opts = append(opts, option.WithAPIEmail(data.APIEmail.ValueString()))
	}
	if !data.UserServiceKey.IsNull() {
		opts = append(opts, option.WithUserServiceKey(data.UserServiceKey.ValueString()))
	}

	client := cloudflare.NewClient(
		opts...,
	)

	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *CloudflareProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		account.NewResource,
		account_member.NewResource,
		origin_ca_certificate.NewResource,
		api_token.NewResource,
		zone.NewResource,
		zone_setting.NewResource,
		zone_hold.NewResource,
		load_balancer.NewResource,
		load_balancer_monitor.NewResource,
		load_balancer_pool.NewResource,
		zone_cache_reserve.NewResource,
		tiered_cache.NewResource,
		zone_cache_variants.NewResource,
		regional_tiered_cache.NewResource,
		certificate_pack.NewResource,
		total_tls.NewResource,
		argo_smart_routing.NewResource,
		argo_tiered_caching.NewResource,
		custom_ssl.NewResource,
		custom_hostname.NewResource,
		custom_hostname_fallback_origin.NewResource,
		dns_record.NewResource,
		dns_zone_dnssec.NewResource,
		email_routing_settings.NewResource,
		email_routing_rule.NewResource,
		email_routing_catch_all.NewResource,
		email_routing_address.NewResource,
		filter.NewResource,
		zone_lockdown.NewResource,
		firewall_rule.NewResource,
		access_rule.NewResource,
		user_agent_blocking_rule.NewResource,
		healthcheck.NewResource,
		keyless_certificate.NewResource,
		logpush_job.NewResource,
		logpush_ownership_challenge.NewResource,
		logpull_retention.NewResource,
		authenticated_origin_pulls_certificate.NewResource,
		authenticated_origin_pulls.NewResource,
		page_rule.NewResource,
		rate_limit.NewResource,
		waiting_room.NewResource,
		waiting_room_event.NewResource,
		waiting_room_rules.NewResource,
		waiting_room_setting.NewResource,
		web3_hostname.NewResource,
		workers_script.NewResource,
		workers_cron_trigger.NewResource,
		workers_custom_domain.NewResource,
		workers_kv_namespace.NewResource,
		workers_kv.NewResource,
		queue.NewResource,
		api_shield.NewResource,
		api_shield_operation.NewResource,
		api_shield_operation_schema_validation_settings.NewResource,
		api_shield_schema_validation_settings.NewResource,
		api_shield_schema.NewResource,
		managed_headers.NewResource,
		ruleset.NewResource,
		url_normalization_settings.NewResource,
		spectrum_application.NewResource,
		regional_hostname.NewResource,
		address_map.NewResource,
		byo_ip_prefix.NewResource,
		magic_wan_gre_tunnel.NewResource,
		magic_wan_ipsec_tunnel.NewResource,
		magic_wan_static_route.NewResource,
		mtls_certificate.NewResource,
		pages_project.NewResource,
		pages_domain.NewResource,
		list.NewResource,
		list_item.NewResource,
		notification_policy_webhooks.NewResource,
		notification_policy.NewResource,
		d1_database.NewResource,
		r2_bucket.NewResource,
		workers_for_platforms_dispatch_namespace.NewResource,
		workers_secret.NewResource,
		zero_trust_dex_test.NewResource,
		zero_trust_device_managed_networks.NewResource,
		zero_trust_device_profiles.NewResource,
		zero_trust_device_certificates.NewResource,
		zero_trust_local_domain_fallback.NewResource,
		zero_trust_device_posture_rule.NewResource,
		zero_trust_device_posture_integration.NewResource,
		zero_trust_access_identity_provider.NewResource,
		zero_trust_organization.NewResource,
		zero_trust_access_application.NewResource,
		zero_trust_access_short_lived_certificate.NewResource,
		zero_trust_access_policy.NewResource,
		zero_trust_access_mtls_certificate.NewResource,
		zero_trust_access_mtls_hostname_settings.NewResource,
		zero_trust_access_group.NewResource,
		zero_trust_access_service_token.NewResource,
		zero_trust_access_key_configuration.NewResource,
		zero_trust_access_custom_page.NewResource,
		zero_trust_access_tag.NewResource,
		zero_trust_tunnel_cloudflared.NewResource,
		zero_trust_tunnel_cloudflared_config.NewResource,
		zero_trust_dlp_custom_profile.NewResource,
		zero_trust_dlp_predefined_profile.NewResource,
		zero_trust_gateway_settings.NewResource,
		zero_trust_list.NewResource,
		zero_trust_dns_location.NewResource,
		zero_trust_gateway_proxy_endpoint.NewResource,
		zero_trust_gateway_policy.NewResource,
		zero_trust_tunnel_cloudflared_route.NewResource,
		zero_trust_tunnel_cloudflared_virtual_network.NewResource,
		turnstile_widget.NewResource,
		hyperdrive_config.NewResource,
		web_analytics_site.NewResource,
		web_analytics_rule.NewResource,
		bot_management.NewResource,
		observatory_scheduled_test.NewResource,
		hostname_tls_setting.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		account.NewAccountDataSource,
		account.NewAccountsDataSource,
		account_member.NewAccountMemberDataSource,
		account_member.NewAccountMembersDataSource,
		origin_ca_certificate.NewOriginCACertificateDataSource,
		origin_ca_certificate.NewOriginCACertificatesDataSource,
		api_token.NewAPITokenDataSource,
		api_token.NewAPITokensDataSource,
		zone.NewZoneDataSource,
		zone.NewZonesDataSource,
		zone_setting.NewZoneSettingDataSource,
		zone_hold.NewZoneHoldDataSource,
		load_balancer.NewLoadBalancerDataSource,
		load_balancer.NewLoadBalancersDataSource,
		load_balancer_monitor.NewLoadBalancerMonitorDataSource,
		load_balancer_monitor.NewLoadBalancerMonitorsDataSource,
		load_balancer_pool.NewLoadBalancerPoolDataSource,
		load_balancer_pool.NewLoadBalancerPoolsDataSource,
		zone_cache_reserve.NewZoneCacheReserveDataSource,
		tiered_cache.NewTieredCacheDataSource,
		zone_cache_variants.NewZoneCacheVariantsDataSource,
		regional_tiered_cache.NewRegionalTieredCacheDataSource,
		certificate_pack.NewCertificatePackDataSource,
		certificate_pack.NewCertificatePacksDataSource,
		total_tls.NewTotalTLSDataSource,
		argo_smart_routing.NewArgoSmartRoutingDataSource,
		argo_tiered_caching.NewArgoTieredCachingDataSource,
		custom_ssl.NewCustomSSLDataSource,
		custom_ssl.NewCustomSSLsDataSource,
		custom_hostname.NewCustomHostnameDataSource,
		custom_hostname.NewCustomHostnamesDataSource,
		custom_hostname_fallback_origin.NewCustomHostnameFallbackOriginDataSource,
		dns_record.NewDNSRecordDataSource,
		dns_record.NewDNSRecordsDataSource,
		dns_zone_dnssec.NewDNSZoneDNSSECDataSource,
		email_routing_settings.NewEmailRoutingSettingsDataSource,
		email_routing_rule.NewEmailRoutingRuleDataSource,
		email_routing_rule.NewEmailRoutingRulesDataSource,
		email_routing_catch_all.NewEmailRoutingCatchAllDataSource,
		email_routing_address.NewEmailRoutingAddressDataSource,
		email_routing_address.NewEmailRoutingAddressesDataSource,
		filter.NewFilterDataSource,
		filter.NewFiltersDataSource,
		zone_lockdown.NewZoneLockdownDataSource,
		zone_lockdown.NewZoneLockdownsDataSource,
		firewall_rule.NewFirewallRuleDataSource,
		firewall_rule.NewFirewallRulesDataSource,
		access_rule.NewAccessRuleDataSource,
		access_rule.NewAccessRulesDataSource,
		user_agent_blocking_rule.NewUserAgentBlockingRuleDataSource,
		user_agent_blocking_rule.NewUserAgentBlockingRulesDataSource,
		healthcheck.NewHealthcheckDataSource,
		healthcheck.NewHealthchecksDataSource,
		keyless_certificate.NewKeylessCertificateDataSource,
		keyless_certificate.NewKeylessCertificatesDataSource,
		logpush_job.NewLogpushJobDataSource,
		logpush_job.NewLogpushJobsDataSource,
		logpull_retention.NewLogpullRetentionDataSource,
		authenticated_origin_pulls_certificate.NewAuthenticatedOriginPullsCertificateDataSource,
		authenticated_origin_pulls_certificate.NewAuthenticatedOriginPullsCertificatesDataSource,
		authenticated_origin_pulls.NewAuthenticatedOriginPullsDataSource,
		page_rule.NewPageRuleDataSource,
		rate_limit.NewRateLimitDataSource,
		rate_limit.NewRateLimitsDataSource,
		waiting_room.NewWaitingRoomDataSource,
		waiting_room.NewWaitingRoomsDataSource,
		waiting_room_event.NewWaitingRoomEventDataSource,
		waiting_room_event.NewWaitingRoomEventsDataSource,
		waiting_room_rules.NewWaitingRoomRulesDataSource,
		waiting_room_setting.NewWaitingRoomSettingDataSource,
		web3_hostname.NewWeb3HostnameDataSource,
		web3_hostname.NewWeb3HostnamesDataSource,
		workers_script.NewWorkersScriptDataSource,
		workers_script.NewWorkersScriptsDataSource,
		workers_cron_trigger.NewWorkersCronTriggerDataSource,
		workers_custom_domain.NewWorkersCustomDomainDataSource,
		workers_custom_domain.NewWorkersCustomDomainsDataSource,
		workers_kv_namespace.NewWorkersKVNamespaceDataSource,
		workers_kv_namespace.NewWorkersKVNamespacesDataSource,
		workers_kv.NewWorkersKVDataSource,
		queue.NewQueueDataSource,
		queue.NewQueuesDataSource,
		api_shield.NewAPIShieldDataSource,
		api_shield_operation.NewAPIShieldOperationDataSource,
		api_shield_operation.NewAPIShieldOperationsDataSource,
		api_shield_operation_schema_validation_settings.NewAPIShieldOperationSchemaValidationSettingsDataSource,
		api_shield_schema_validation_settings.NewAPIShieldSchemaValidationSettingsDataSource,
		api_shield_schema.NewAPIShieldSchemaDataSource,
		api_shield_schema.NewAPIShieldSchemasDataSource,
		managed_headers.NewManagedHeadersDataSource,
		ruleset.NewRulesetDataSource,
		ruleset.NewRulesetsDataSource,
		url_normalization_settings.NewURLNormalizationSettingsDataSource,
		spectrum_application.NewSpectrumApplicationDataSource,
		regional_hostname.NewRegionalHostnameDataSource,
		regional_hostname.NewRegionalHostnamesDataSource,
		address_map.NewAddressMapDataSource,
		address_map.NewAddressMapsDataSource,
		byo_ip_prefix.NewByoIPPrefixDataSource,
		byo_ip_prefix.NewByoIPPrefixesDataSource,
		magic_wan_gre_tunnel.NewMagicWANGRETunnelDataSource,
		magic_wan_ipsec_tunnel.NewMagicWANIPSECTunnelDataSource,
		magic_wan_static_route.NewMagicWANStaticRouteDataSource,
		mtls_certificate.NewMTLSCertificateDataSource,
		mtls_certificate.NewMTLSCertificatesDataSource,
		pages_project.NewPagesProjectDataSource,
		pages_project.NewPagesProjectsDataSource,
		pages_domain.NewPagesDomainDataSource,
		pages_domain.NewPagesDomainsDataSource,
		list.NewListDataSource,
		list.NewListsDataSource,
		list_item.NewListItemDataSource,
		list_item.NewListItemsDataSource,
		notification_policy_webhooks.NewNotificationPolicyWebhooksDataSource,
		notification_policy_webhooks.NewNotificationPolicyWebhooksListDataSource,
		notification_policy.NewNotificationPolicyDataSource,
		notification_policy.NewNotificationPoliciesDataSource,
		d1_database.NewD1DatabaseDataSource,
		d1_database.NewD1DatabasesDataSource,
		r2_bucket.NewR2BucketDataSource,
		workers_for_platforms_dispatch_namespace.NewWorkersForPlatformsDispatchNamespaceDataSource,
		workers_for_platforms_dispatch_namespace.NewWorkersForPlatformsDispatchNamespacesDataSource,
		workers_secret.NewWorkersSecretDataSource,
		workers_secret.NewWorkersSecretsDataSource,
		zero_trust_dex_test.NewZeroTrustDEXTestDataSource,
		zero_trust_dex_test.NewZeroTrustDEXTestsDataSource,
		zero_trust_device_managed_networks.NewZeroTrustDeviceManagedNetworksDataSource,
		zero_trust_device_managed_networks.NewZeroTrustDeviceManagedNetworksListDataSource,
		zero_trust_device_profiles.NewZeroTrustDeviceProfilesDataSource,
		zero_trust_device_profiles.NewZeroTrustDeviceProfilesListDataSource,
		zero_trust_device_certificates.NewZeroTrustDeviceCertificatesDataSource,
		zero_trust_local_domain_fallback.NewZeroTrustLocalDomainFallbackDataSource,
		zero_trust_local_domain_fallback.NewZeroTrustLocalDomainFallbacksDataSource,
		zero_trust_device_posture_rule.NewZeroTrustDevicePostureRuleDataSource,
		zero_trust_device_posture_rule.NewZeroTrustDevicePostureRulesDataSource,
		zero_trust_device_posture_integration.NewZeroTrustDevicePostureIntegrationDataSource,
		zero_trust_device_posture_integration.NewZeroTrustDevicePostureIntegrationsDataSource,
		zero_trust_access_identity_provider.NewZeroTrustAccessIdentityProviderDataSource,
		zero_trust_access_identity_provider.NewZeroTrustAccessIdentityProvidersDataSource,
		zero_trust_organization.NewZeroTrustOrganizationDataSource,
		zero_trust_access_application.NewZeroTrustAccessApplicationDataSource,
		zero_trust_access_application.NewZeroTrustAccessApplicationsDataSource,
		zero_trust_access_short_lived_certificate.NewZeroTrustAccessShortLivedCertificateDataSource,
		zero_trust_access_short_lived_certificate.NewZeroTrustAccessShortLivedCertificatesDataSource,
		zero_trust_access_policy.NewZeroTrustAccessPolicyDataSource,
		zero_trust_access_policy.NewZeroTrustAccessPoliciesDataSource,
		zero_trust_access_mtls_certificate.NewZeroTrustAccessMTLSCertificateDataSource,
		zero_trust_access_mtls_certificate.NewZeroTrustAccessMTLSCertificatesDataSource,
		zero_trust_access_mtls_hostname_settings.NewZeroTrustAccessMTLSHostnameSettingsDataSource,
		zero_trust_access_group.NewZeroTrustAccessGroupDataSource,
		zero_trust_access_group.NewZeroTrustAccessGroupsDataSource,
		zero_trust_access_service_token.NewZeroTrustAccessServiceTokenDataSource,
		zero_trust_access_service_token.NewZeroTrustAccessServiceTokensDataSource,
		zero_trust_access_key_configuration.NewZeroTrustAccessKeyConfigurationDataSource,
		zero_trust_access_custom_page.NewZeroTrustAccessCustomPageDataSource,
		zero_trust_access_custom_page.NewZeroTrustAccessCustomPagesDataSource,
		zero_trust_access_tag.NewZeroTrustAccessTagDataSource,
		zero_trust_access_tag.NewZeroTrustAccessTagsDataSource,
		zero_trust_tunnel_cloudflared.NewZeroTrustTunnelCloudflaredDataSource,
		zero_trust_tunnel_cloudflared.NewZeroTrustTunnelCloudflaredsDataSource,
		zero_trust_tunnel_cloudflared_config.NewZeroTrustTunnelCloudflaredConfigDataSource,
		zero_trust_dlp_custom_profile.NewZeroTrustDLPCustomProfileDataSource,
		zero_trust_dlp_predefined_profile.NewZeroTrustDLPPredefinedProfileDataSource,
		zero_trust_gateway_settings.NewZeroTrustGatewaySettingsDataSource,
		zero_trust_list.NewZeroTrustListDataSource,
		zero_trust_list.NewZeroTrustListsDataSource,
		zero_trust_dns_location.NewZeroTrustDNSLocationDataSource,
		zero_trust_dns_location.NewZeroTrustDNSLocationsDataSource,
		zero_trust_gateway_proxy_endpoint.NewZeroTrustGatewayProxyEndpointDataSource,
		zero_trust_gateway_policy.NewZeroTrustGatewayPolicyDataSource,
		zero_trust_gateway_policy.NewZeroTrustGatewayPoliciesDataSource,
		zero_trust_tunnel_cloudflared_route.NewZeroTrustTunnelCloudflaredRouteDataSource,
		zero_trust_tunnel_cloudflared_route.NewZeroTrustTunnelCloudflaredRoutesDataSource,
		zero_trust_tunnel_cloudflared_virtual_network.NewZeroTrustTunnelCloudflaredVirtualNetworkDataSource,
		zero_trust_tunnel_cloudflared_virtual_network.NewZeroTrustTunnelCloudflaredVirtualNetworksDataSource,
		turnstile_widget.NewTurnstileWidgetDataSource,
		turnstile_widget.NewTurnstileWidgetsDataSource,
		hyperdrive_config.NewHyperdriveConfigDataSource,
		hyperdrive_config.NewHyperdriveConfigsDataSource,
		web_analytics_site.NewWebAnalyticsSiteDataSource,
		web_analytics_site.NewWebAnalyticsSitesDataSource,
		bot_management.NewBotManagementDataSource,
		observatory_scheduled_test.NewObservatoryScheduledTestDataSource,
		hostname_tls_setting.NewHostnameTLSSettingDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
