// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"fmt"
	"regexp"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
	APIKey                  types.String `tfsdk:"api_key" json:"api_key"`
	APIUserServiceKey       types.String `tfsdk:"api_user_service_key" json:"api_user_service_key"`
	Email                   types.String `tfsdk:"email" json:"email"`
	APIToken                types.String `tfsdk:"api_token" json:"api_token"`
	UserAgentOperatorSuffix types.String `tfsdk:"user_agent_operator_suffix" json:"user_agent_operator_suffix"`
	BaseURL                 types.String `tfsdk:"base_url" json:"base_url"`
}

func (p *CloudflareProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cloudflare"
	resp.Version = p.version
}

func (p CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			consts.EmailSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("A registered Cloudflare email address. Alternatively, can be configured using the `%s` environment variable. Required when using `api_key`. Conflicts with `api_token`.", consts.EmailEnvVarKey),
				Validators:          []validator.String{},
			},

			consts.APIKeySchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("The API key for operations. Alternatively, can be configured using the `%s` environment variable. API keys are [now considered legacy by Cloudflare](https://developers.cloudflare.com/fundamentals/api/get-started/keys/#limitations), API tokens should be used instead. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APIKeyEnvVarKey),
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`[0-9a-f]{37}`),
						"API key must be 37 characters long and only contain characters 0-9 and a-f (all lowercased)",
					),
					stringvalidator.AlsoRequires(path.Expressions{
						path.MatchRoot(consts.EmailSchemaKey),
					}...),
				},
			},

			consts.APITokenSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("The API Token for operations. Alternatively, can be configured using the `%s` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APITokenEnvVarKey),
				Validators: []validator.String{
					stringvalidator.RegexMatches(
						regexp.MustCompile(`[A-Za-z0-9-_]{40}`),
						"API tokens must be 40 characters long and only contain characters a-z, A-Z, 0-9, hyphens and underscores",
					),
				},
			},

			consts.APIUserServiceKeySchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("A special Cloudflare API key good for a restricted set of endpoints. Alternatively, can be configured using the `%s` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APIUserServiceKeyEnvVarKey),
			},

			consts.UserAgentOperatorSuffixSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("A value to append to the HTTP User Agent for all API calls. This value is not something most users need to modify however, if you are using a non-standard provider or operator configuration, this is recommended to assist in uniquely identifying your traffic. **Setting this value will remove the Terraform version from the HTTP User Agent string and may have unintended consequences**. Alternatively, can be configured using the `%s` environment variable.", consts.UserAgentOperatorSuffixEnvVarKey),
			},

			consts.BaseURLSchemaKey: schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: fmt.Sprintf("Value to override the default HTTP client base URL. Alternatively, can be configured using the `%s` environment variable.", consts.BaseURLSchemaKey),
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
	if !data.Email.IsNull() {
		opts = append(opts, option.WithAPIEmail(data.Email.ValueString()))
	}
	if !data.APIUserServiceKey.IsNull() {
		opts = append(opts, option.WithUserServiceKey(data.APIUserServiceKey.ValueString()))
	}

	pluginVersion := utils.FindGoModuleVersion("github.com/hashicorp/terraform-plugin-framework")
	framework := "terraform-plugin-framework"
	userAgentParams := utils.UserAgentBuilderParams{
		ProviderVersion: &p.version,
		PluginType:      &framework,
		PluginVersion:   pluginVersion,
	}

	if !data.UserAgentOperatorSuffix.IsNull() {
		operatorSuffix := data.UserAgentOperatorSuffix.String()
		userAgentParams.OperatorSuffix = &operatorSuffix
	} else {
		userAgentParams.TerraformVersion = &req.TerraformVersion
	}

	opts = append(opts, option.WithHeader("user-agent", userAgentParams.String()))
	opts = append(opts, option.WithHeader("x-stainless-package-version", p.version))
	opts = append(opts, option.WithHeader("x-stainless-runtime", framework))
	opts = append(opts, option.WithHeader("x-stainless-lang", "Terraform"))
	if pluginVersion != nil {
		opts = append(opts, option.WithHeader("x-stainless-runtime-version", *pluginVersion))
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
		load_balancer.NewLoadBalancerDataSource,
		load_balancer.NewLoadBalancersDataSource,
		load_balancer_monitor.NewLoadBalancerMonitorDataSource,
		load_balancer_monitor.NewLoadBalancerMonitorsDataSource,
		load_balancer_pool.NewLoadBalancerPoolDataSource,
		load_balancer_pool.NewLoadBalancerPoolsDataSource,
		certificate_pack.NewCertificatePackDataSource,
		certificate_pack.NewCertificatePacksDataSource,
		custom_ssl.NewCustomSSLDataSource,
		custom_ssl.NewCustomSSLsDataSource,
		custom_hostname.NewCustomHostnameDataSource,
		custom_hostname.NewCustomHostnamesDataSource,
		record.NewRecordDataSource,
		record.NewRecordsDataSource,
		email_routing_rule.NewEmailRoutingRuleDataSource,
		email_routing_rule.NewEmailRoutingRulesDataSource,
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
		authenticated_origin_pulls_certificate.NewAuthenticatedOriginPullsCertificateDataSource,
		authenticated_origin_pulls_certificate.NewAuthenticatedOriginPullsCertificatesDataSource,
		rate_limit.NewRateLimitDataSource,
		rate_limit.NewRateLimitsDataSource,
		waiting_room.NewWaitingRoomDataSource,
		waiting_room.NewWaitingRoomsDataSource,
		waiting_room_event.NewWaitingRoomEventDataSource,
		waiting_room_event.NewWaitingRoomEventsDataSource,
		waiting_room_rules.NewWaitingRoomRulesListDataSource,
		web3_hostname.NewWeb3HostnameDataSource,
		web3_hostname.NewWeb3HostnamesDataSource,
		worker_script.NewWorkerScriptDataSource,
		worker_script.NewWorkerScriptsDataSource,
		worker_domain.NewWorkerDomainDataSource,
		worker_domain.NewWorkerDomainsDataSource,
		workers_kv_namespace.NewWorkersKVNamespaceDataSource,
		workers_kv_namespace.NewWorkersKVNamespacesDataSource,
		queue.NewQueueDataSource,
		queue.NewQueuesDataSource,
		api_shield_operation.NewAPIShieldOperationsDataSource,
		api_shield_schema.NewAPIShieldSchemaDataSource,
		api_shield_schema.NewAPIShieldSchemasDataSource,
		ruleset.NewRulesetDataSource,
		ruleset.NewRulesetsDataSource,
		spectrum_application.NewSpectrumApplicationDataSource,
		spectrum_application.NewSpectrumApplicationsDataSource,
		regional_hostname.NewRegionalHostnameDataSource,
		regional_hostname.NewRegionalHostnamesDataSource,
		address_map.NewAddressMapDataSource,
		address_map.NewAddressMapsDataSource,
		byo_ip_prefix.NewByoIPPrefixDataSource,
		byo_ip_prefix.NewByoIPPrefixesDataSource,
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
		r2_bucket.NewR2BucketsDataSource,
		workers_for_platforms_namespace.NewWorkersForPlatformsNamespaceDataSource,
		workers_for_platforms_namespace.NewWorkersForPlatformsNamespacesDataSource,
		worker_secret.NewWorkerSecretsDataSource,
		device_dex_test.NewDeviceDEXTestDataSource,
		device_dex_test.NewDeviceDEXTestsDataSource,
		device_managed_networks.NewDeviceManagedNetworksDataSource,
		device_managed_networks.NewDeviceManagedNetworksListDataSource,
		device_settings_policy.NewDeviceSettingsPolicyDataSource,
		device_settings_policy.NewDeviceSettingsPoliciesDataSource,
		fallback_domain.NewFallbackDomainDataSource,
		fallback_domain.NewFallbackDomainsDataSource,
		device_posture_rule.NewDevicePostureRuleDataSource,
		device_posture_rule.NewDevicePostureRulesDataSource,
		device_posture_integration.NewDevicePostureIntegrationDataSource,
		device_posture_integration.NewDevicePostureIntegrationsDataSource,
		access_identity_provider.NewAccessIdentityProviderDataSource,
		access_identity_provider.NewAccessIdentityProvidersDataSource,
		access_application.NewAccessApplicationDataSource,
		access_application.NewAccessApplicationsDataSource,
		access_ca_certificate.NewAccessCACertificateDataSource,
		access_ca_certificate.NewAccessCACertificatesDataSource,
		access_policy.NewAccessPolicyDataSource,
		access_policy.NewAccessPoliciesDataSource,
		access_mutual_tls_certificate.NewAccessMutualTLSCertificateDataSource,
		access_mutual_tls_certificate.NewAccessMutualTLSCertificatesDataSource,
		access_group.NewAccessGroupDataSource,
		access_group.NewAccessGroupsDataSource,
		access_service_token.NewAccessServiceTokenDataSource,
		access_service_token.NewAccessServiceTokensDataSource,
		access_custom_page.NewAccessCustomPageDataSource,
		access_custom_page.NewAccessCustomPagesDataSource,
		access_tag.NewAccessTagDataSource,
		access_tag.NewAccessTagsDataSource,
		tunnel.NewTunnelDataSource,
		tunnel.NewTunnelsDataSource,
		teams_list.NewTeamsListDataSource,
		teams_list.NewTeamsListsDataSource,
		teams_location.NewTeamsLocationDataSource,
		teams_location.NewTeamsLocationsDataSource,
		teams_rule.NewTeamsRuleDataSource,
		teams_rule.NewTeamsRulesDataSource,
		tunnel_route.NewTunnelRoutesDataSource,
		tunnel_virtual_network.NewTunnelVirtualNetworksDataSource,
		turnstile_widget.NewTurnstileWidgetDataSource,
		turnstile_widget.NewTurnstileWidgetsDataSource,
		hyperdrive_config.NewHyperdriveConfigDataSource,
		hyperdrive_config.NewHyperdriveConfigsDataSource,
		web_analytics_site.NewWebAnalyticsSiteDataSource,
		web_analytics_site.NewWebAnalyticsSitesDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
