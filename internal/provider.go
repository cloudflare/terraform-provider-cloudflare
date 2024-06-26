// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_application"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_ca_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_custom_page"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_identity_provider"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_keys_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_mutual_tls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_mutual_tls_hostname_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_organization"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_service_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_tag"
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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/device_dex_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/device_managed_networks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/device_posture_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/device_posture_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/device_settings_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dlp_custom_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dlp_predefined_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_address"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_catch_all"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/fallback_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/filter"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/firewall_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/gre_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/healthcheck"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hostname_tls_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ipsec_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/keyless_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpull_retention"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_job"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_ownership_challenge"
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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/record"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/spectrum_application"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/static_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/teams_account"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/teams_list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/teams_location"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/teams_proxy_endpoint"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/teams_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/total_tls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tunnel_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tunnel_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tunnel_virtual_network"
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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_cron_trigger"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_script"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/worker_secret"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_reserve"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_variants"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dnssec"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_hold"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_lockdown"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &CloudflareProvider{}

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

func (p CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
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
		record.NewResource,
		zone_dnssec.NewResource,
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
		worker_script.NewResource,
		worker_cron_trigger.NewResource,
		worker_domain.NewResource,
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
		gre_tunnel.NewResource,
		ipsec_tunnel.NewResource,
		static_route.NewResource,
		mtls_certificate.NewResource,
		pages_project.NewResource,
		pages_domain.NewResource,
		list.NewResource,
		list_item.NewResource,
		notification_policy_webhooks.NewResource,
		notification_policy.NewResource,
		d1_database.NewResource,
		r2_bucket.NewResource,
		workers_for_platforms_namespace.NewResource,
		worker_secret.NewResource,
		device_dex_test.NewResource,
		device_managed_networks.NewResource,
		device_settings_policy.NewResource,
		fallback_domain.NewResource,
		device_posture_rule.NewResource,
		device_posture_integration.NewResource,
		access_identity_provider.NewResource,
		access_organization.NewResource,
		access_application.NewResource,
		access_ca_certificate.NewResource,
		access_policy.NewResource,
		access_mutual_tls_certificate.NewResource,
		access_mutual_tls_hostname_settings.NewResource,
		access_group.NewResource,
		access_service_token.NewResource,
		access_keys_configuration.NewResource,
		access_custom_page.NewResource,
		access_tag.NewResource,
		tunnel.NewResource,
		tunnel_config.NewResource,
		dlp_custom_profile.NewResource,
		dlp_predefined_profile.NewResource,
		teams_account.NewResource,
		teams_list.NewResource,
		teams_location.NewResource,
		teams_proxy_endpoint.NewResource,
		teams_rule.NewResource,
		tunnel_route.NewResource,
		tunnel_virtual_network.NewResource,
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
	return []func() datasource.DataSource{}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
