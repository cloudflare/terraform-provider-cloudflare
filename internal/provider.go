// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package internal

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/access_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_member"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_role"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_subscription"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/account_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/address_map"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_discovery_operation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_operation_schema_validation_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_schema"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_shield_schema_validation_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/api_token_permissions_groups"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_smart_routing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/argo_tiered_caching"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/authenticated_origin_pulls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/bot_management"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/botnet_feed_config_asn"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/byo_ip_prefix"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/calls_sfu_app"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/calls_turn_app"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/certificate_pack"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloud_connector_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloudforce_one_request"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloudforce_one_request_asset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloudforce_one_request_message"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/cloudforce_one_request_priority"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/content_scanning_expression"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_hostname_fallback_origin"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/d1_database"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dcv_delegation"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_record"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_settings_internal_view"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_incoming"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_outgoing"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_peer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_zone_transfers_tsig"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_address"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_catch_all"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_dns"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_routing_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_block_sender"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_impersonation_registry"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/email_security_trusted_domains"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/filter"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/firewall_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/healthcheck"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hostname_tls_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/hyperdrive_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/image"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/image_variant"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ip_ranges"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/keyless_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/leaked_credential_check"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/leaked_credential_check_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/list_item"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_monitor"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/load_balancer_pool"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpull_retention"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_dataset_field"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_dataset_job"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_job"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/logpush_ownership_challenge"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_network_monitoring_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_network_monitoring_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_connector"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site_acl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site_lan"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_transit_site_wan"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_gre_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_ipsec_tunnel"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/magic_wan_static_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/managed_transforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/mtls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/notification_policy_webhooks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/observatory_scheduled_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/origin_ca_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_connections"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_cookies"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/page_shield_scripts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/pages_project"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/permission_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/queue_consumer"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_cors"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_event_notification"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_lifecycle"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_lock"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_bucket_sippy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_custom_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/r2_managed_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/rate_limit"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/regional_tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/registrar_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/resource_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/ruleset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippet_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/snippets"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/spectrum_application"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_audio_track"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_caption_language"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_download"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_key"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_live_input"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_watermark"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream_webhook"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/tiered_cache"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/total_tls"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/turnstile_widget"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/url_normalization_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/user"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/user_agent_blocking_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_event"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_rules"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/waiting_room_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/web3_hostname"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/web_analytics_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/web_analytics_site"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_cron_trigger"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_custom_domain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_deployment"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_for_platforms_dispatch_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_kv_namespace"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_script_subdomain"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/workers_secret"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_application"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_custom_page"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_group"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_identity_provider"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_infrastructure_target"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_key_configuration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_mtls_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_mtls_hostname_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_service_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_short_lived_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_access_tag"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_custom_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_custom_profile_local_domain_fallback"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile_certificates"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_default_profile_local_domain_fallback"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_managed_networks"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_device_posture_rule"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dex_test"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_custom_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_dataset"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_entry"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dlp_predefined_profile"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_dns_location"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_app_types"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_categories"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_certificate"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_policy"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_proxy_endpoint"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_gateway_settings"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_list"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_organization"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_risk_behavior"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_risk_scoring_integration"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_config"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_route"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_token"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zero_trust_tunnel_cloudflared_virtual_network"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_reserve"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_cache_variants"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_dnssec"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_hold"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_lockdown"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_setting"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/zone_subscription"
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

func ProviderSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
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

func (p *CloudflareProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = ProviderSchema(ctx)
}

func (p *CloudflareProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	var data CloudflareProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	opts := []option.RequestOption{}

	if !data.BaseURL.IsNull() {
		opts = append(opts, option.WithBaseURL(data.BaseURL.ValueString()))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_TOKEN"); ok {
		opts = append(opts, option.WithAPIToken(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		opts = append(opts, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_EMAIL"); ok {
		opts = append(opts, option.WithAPIEmail(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_USER_SERVICE_KEY"); ok {
		opts = append(opts, option.WithUserServiceKey(o))
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

func (p *CloudflareProvider) ConfigValidators(_ context.Context) []provider.ConfigValidator {
	return []provider.ConfigValidator{}
}

func (p *CloudflareProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		account.NewResource,
		account_member.NewResource,
		account_subscription.NewResource,
		account_token.NewResource,
		origin_ca_certificate.NewResource,
		user.NewResource,
		api_token.NewResource,
		zone.NewResource,
		zone_setting.NewResource,
		zone_hold.NewResource,
		zone_subscription.NewResource,
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
		dns_firewall.NewResource,
		zone_dnssec.NewResource,
		dns_record.NewResource,
		dns_settings.NewResource,
		dns_settings_internal_view.NewResource,
		dns_zone_transfers_incoming.NewResource,
		dns_zone_transfers_outgoing.NewResource,
		dns_zone_transfers_acl.NewResource,
		dns_zone_transfers_peer.NewResource,
		dns_zone_transfers_tsig.NewResource,
		email_security_block_sender.NewResource,
		email_security_impersonation_registry.NewResource,
		email_security_trusted_domains.NewResource,
		email_routing_settings.NewResource,
		email_routing_dns.NewResource,
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
		waiting_room_settings.NewResource,
		web3_hostname.NewResource,
		workers_route.NewResource,
		workers_script.NewResource,
		workers_script_subdomain.NewResource,
		workers_cron_trigger.NewResource,
		workers_deployment.NewResource,
		workers_custom_domain.NewResource,
		workers_kv_namespace.NewResource,
		workers_kv.NewResource,
		queue.NewResource,
		queue_consumer.NewResource,
		api_shield.NewResource,
		api_shield_discovery_operation.NewResource,
		api_shield_operation.NewResource,
		api_shield_operation_schema_validation_settings.NewResource,
		api_shield_schema_validation_settings.NewResource,
		api_shield_schema.NewResource,
		managed_transforms.NewResource,
		page_shield_policy.NewResource,
		ruleset.NewResource,
		url_normalization_settings.NewResource,
		spectrum_application.NewResource,
		regional_hostname.NewResource,
		address_map.NewResource,
		byo_ip_prefix.NewResource,
		image.NewResource,
		image_variant.NewResource,
		magic_wan_gre_tunnel.NewResource,
		magic_wan_ipsec_tunnel.NewResource,
		magic_wan_static_route.NewResource,
		magic_transit_site.NewResource,
		magic_transit_site_acl.NewResource,
		magic_transit_site_lan.NewResource,
		magic_transit_site_wan.NewResource,
		magic_transit_connector.NewResource,
		magic_network_monitoring_configuration.NewResource,
		magic_network_monitoring_rule.NewResource,
		mtls_certificate.NewResource,
		pages_project.NewResource,
		pages_domain.NewResource,
		registrar_domain.NewResource,
		list.NewResource,
		list_item.NewResource,
		stream.NewResource,
		stream_audio_track.NewResource,
		stream_key.NewResource,
		stream_live_input.NewResource,
		stream_watermark.NewResource,
		stream_webhook.NewResource,
		stream_caption_language.NewResource,
		stream_download.NewResource,
		notification_policy_webhooks.NewResource,
		notification_policy.NewResource,
		d1_database.NewResource,
		r2_bucket.NewResource,
		r2_bucket_lifecycle.NewResource,
		r2_bucket_cors.NewResource,
		r2_custom_domain.NewResource,
		r2_managed_domain.NewResource,
		r2_bucket_event_notification.NewResource,
		r2_bucket_lock.NewResource,
		r2_bucket_sippy.NewResource,
		workers_for_platforms_dispatch_namespace.NewResource,
		workers_secret.NewResource,
		zero_trust_dex_test.NewResource,
		zero_trust_device_managed_networks.NewResource,
		zero_trust_device_default_profile.NewResource,
		zero_trust_device_default_profile_local_domain_fallback.NewResource,
		zero_trust_device_default_profile_certificates.NewResource,
		zero_trust_device_custom_profile.NewResource,
		zero_trust_device_custom_profile_local_domain_fallback.NewResource,
		zero_trust_device_posture_rule.NewResource,
		zero_trust_device_posture_integration.NewResource,
		zero_trust_access_identity_provider.NewResource,
		zero_trust_organization.NewResource,
		zero_trust_access_infrastructure_target.NewResource,
		zero_trust_access_application.NewResource,
		zero_trust_access_short_lived_certificate.NewResource,
		zero_trust_access_mtls_certificate.NewResource,
		zero_trust_access_mtls_hostname_settings.NewResource,
		zero_trust_access_group.NewResource,
		zero_trust_access_service_token.NewResource,
		zero_trust_access_key_configuration.NewResource,
		zero_trust_access_custom_page.NewResource,
		zero_trust_access_tag.NewResource,
		zero_trust_access_policy.NewResource,
		zero_trust_tunnel_cloudflared.NewResource,
		zero_trust_tunnel_cloudflared_config.NewResource,
		zero_trust_dlp_dataset.NewResource,
		zero_trust_dlp_custom_profile.NewResource,
		zero_trust_dlp_predefined_profile.NewResource,
		zero_trust_dlp_entry.NewResource,
		zero_trust_gateway_settings.NewResource,
		zero_trust_list.NewResource,
		zero_trust_dns_location.NewResource,
		zero_trust_gateway_proxy_endpoint.NewResource,
		zero_trust_gateway_policy.NewResource,
		zero_trust_gateway_certificate.NewResource,
		zero_trust_tunnel_cloudflared_route.NewResource,
		zero_trust_tunnel_cloudflared_virtual_network.NewResource,
		zero_trust_risk_behavior.NewResource,
		zero_trust_risk_scoring_integration.NewResource,
		turnstile_widget.NewResource,
		hyperdrive_config.NewResource,
		web_analytics_site.NewResource,
		web_analytics_rule.NewResource,
		bot_management.NewResource,
		observatory_scheduled_test.NewResource,
		hostname_tls_setting.NewResource,
		snippets.NewResource,
		snippet_rules.NewResource,
		calls_sfu_app.NewResource,
		calls_turn_app.NewResource,
		cloudforce_one_request.NewResource,
		cloudforce_one_request_message.NewResource,
		cloudforce_one_request_priority.NewResource,
		cloudforce_one_request_asset.NewResource,
		cloud_connector_rules.NewResource,
		leaked_credential_check.NewResource,
		leaked_credential_check_rule.NewResource,
		content_scanning_expression.NewResource,
	}
}

func (p *CloudflareProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		account.NewAccountDataSource,
		account.NewAccountsDataSource,
		account_member.NewAccountMemberDataSource,
		account_member.NewAccountMembersDataSource,
		account_role.NewAccountRoleDataSource,
		account_role.NewAccountRolesDataSource,
		account_subscription.NewAccountSubscriptionDataSource,
		account_token.NewAccountTokenDataSource,
		account_token.NewAccountTokensDataSource,
		api_token_permissions_groups.NewAPITokenPermissionsGroupsDataSource,
		api_token_permissions_groups.NewAPITokenPermissionsGroupsListDataSource,
		origin_ca_certificate.NewOriginCACertificateDataSource,
		origin_ca_certificate.NewOriginCACertificatesDataSource,
		ip_ranges.NewIPRangesDataSource,
		user.NewUserDataSource,
		api_token.NewAPITokenDataSource,
		api_token.NewAPITokensDataSource,
		zone.NewZoneDataSource,
		zone.NewZonesDataSource,
		zone_setting.NewZoneSettingDataSource,
		zone_hold.NewZoneHoldDataSource,
		zone_subscription.NewZoneSubscriptionDataSource,
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
		dns_firewall.NewDNSFirewallDataSource,
		dns_firewall.NewDNSFirewallsDataSource,
		zone_dnssec.NewZoneDNSSECDataSource,
		dns_record.NewDNSRecordDataSource,
		dns_record.NewDNSRecordsDataSource,
		dns_settings.NewDNSSettingsDataSource,
		dns_settings_internal_view.NewDNSSettingsInternalViewDataSource,
		dns_settings_internal_view.NewDNSSettingsInternalViewsDataSource,
		dns_zone_transfers_incoming.NewDNSZoneTransfersIncomingDataSource,
		dns_zone_transfers_outgoing.NewDNSZoneTransfersOutgoingDataSource,
		dns_zone_transfers_acl.NewDNSZoneTransfersACLDataSource,
		dns_zone_transfers_acl.NewDNSZoneTransfersACLsDataSource,
		dns_zone_transfers_peer.NewDNSZoneTransfersPeerDataSource,
		dns_zone_transfers_peer.NewDNSZoneTransfersPeersDataSource,
		dns_zone_transfers_tsig.NewDNSZoneTransfersTSIGDataSource,
		dns_zone_transfers_tsig.NewDNSZoneTransfersTSIGsDataSource,
		email_security_block_sender.NewEmailSecurityBlockSenderDataSource,
		email_security_block_sender.NewEmailSecurityBlockSendersDataSource,
		email_security_impersonation_registry.NewEmailSecurityImpersonationRegistryDataSource,
		email_security_impersonation_registry.NewEmailSecurityImpersonationRegistriesDataSource,
		email_security_trusted_domains.NewEmailSecurityTrustedDomainsDataSource,
		email_security_trusted_domains.NewEmailSecurityTrustedDomainsListDataSource,
		email_routing_settings.NewEmailRoutingSettingsDataSource,
		email_routing_dns.NewEmailRoutingDNSDataSource,
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
		logpush_dataset_field.NewLogpushDatasetFieldDataSource,
		logpush_dataset_job.NewLogpushDatasetJobDataSource,
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
		waiting_room_settings.NewWaitingRoomSettingsDataSource,
		web3_hostname.NewWeb3HostnameDataSource,
		web3_hostname.NewWeb3HostnamesDataSource,
		workers_route.NewWorkersRouteDataSource,
		workers_route.NewWorkersRoutesDataSource,
		workers_script.NewWorkersScriptDataSource,
		workers_script.NewWorkersScriptsDataSource,
		workers_script_subdomain.NewWorkersScriptSubdomainDataSource,
		workers_cron_trigger.NewWorkersCronTriggerDataSource,
		workers_deployment.NewWorkersDeploymentDataSource,
		workers_custom_domain.NewWorkersCustomDomainDataSource,
		workers_custom_domain.NewWorkersCustomDomainsDataSource,
		workers_kv_namespace.NewWorkersKVNamespaceDataSource,
		workers_kv_namespace.NewWorkersKVNamespacesDataSource,
		workers_kv.NewWorkersKVDataSource,
		queue.NewQueueDataSource,
		queue.NewQueuesDataSource,
		queue_consumer.NewQueueConsumerDataSource,
		api_shield.NewAPIShieldDataSource,
		api_shield_discovery_operation.NewAPIShieldDiscoveryOperationsDataSource,
		api_shield_operation.NewAPIShieldOperationDataSource,
		api_shield_operation.NewAPIShieldOperationsDataSource,
		api_shield_operation_schema_validation_settings.NewAPIShieldOperationSchemaValidationSettingsDataSource,
		api_shield_schema_validation_settings.NewAPIShieldSchemaValidationSettingsDataSource,
		api_shield_schema.NewAPIShieldSchemaDataSource,
		api_shield_schema.NewAPIShieldSchemasDataSource,
		managed_transforms.NewManagedTransformsDataSource,
		page_shield_policy.NewPageShieldPolicyDataSource,
		page_shield_policy.NewPageShieldPoliciesDataSource,
		page_shield_connections.NewPageShieldConnectionsDataSource,
		page_shield_connections.NewPageShieldConnectionsListDataSource,
		page_shield_scripts.NewPageShieldScriptsDataSource,
		page_shield_scripts.NewPageShieldScriptsListDataSource,
		page_shield_cookies.NewPageShieldCookiesDataSource,
		page_shield_cookies.NewPageShieldCookiesListDataSource,
		ruleset.NewRulesetDataSource,
		ruleset.NewRulesetsDataSource,
		url_normalization_settings.NewURLNormalizationSettingsDataSource,
		spectrum_application.NewSpectrumApplicationDataSource,
		spectrum_application.NewSpectrumApplicationsDataSource,
		regional_hostname.NewRegionalHostnameDataSource,
		regional_hostname.NewRegionalHostnamesDataSource,
		address_map.NewAddressMapDataSource,
		address_map.NewAddressMapsDataSource,
		byo_ip_prefix.NewByoIPPrefixDataSource,
		byo_ip_prefix.NewByoIPPrefixesDataSource,
		image.NewImageDataSource,
		image.NewImagesDataSource,
		image_variant.NewImageVariantDataSource,
		magic_wan_gre_tunnel.NewMagicWANGRETunnelDataSource,
		magic_wan_ipsec_tunnel.NewMagicWANIPSECTunnelDataSource,
		magic_wan_static_route.NewMagicWANStaticRouteDataSource,
		magic_transit_site.NewMagicTransitSiteDataSource,
		magic_transit_site.NewMagicTransitSitesDataSource,
		magic_transit_site_acl.NewMagicTransitSiteACLDataSource,
		magic_transit_site_acl.NewMagicTransitSiteACLsDataSource,
		magic_transit_site_lan.NewMagicTransitSiteLANDataSource,
		magic_transit_site_lan.NewMagicTransitSiteLANsDataSource,
		magic_transit_site_wan.NewMagicTransitSiteWANDataSource,
		magic_transit_site_wan.NewMagicTransitSiteWANsDataSource,
		magic_transit_connector.NewMagicTransitConnectorDataSource,
		magic_transit_connector.NewMagicTransitConnectorsDataSource,
		magic_network_monitoring_configuration.NewMagicNetworkMonitoringConfigurationDataSource,
		magic_network_monitoring_rule.NewMagicNetworkMonitoringRuleDataSource,
		magic_network_monitoring_rule.NewMagicNetworkMonitoringRulesDataSource,
		mtls_certificate.NewMTLSCertificateDataSource,
		mtls_certificate.NewMTLSCertificatesDataSource,
		pages_project.NewPagesProjectDataSource,
		pages_project.NewPagesProjectsDataSource,
		pages_domain.NewPagesDomainDataSource,
		pages_domain.NewPagesDomainsDataSource,
		registrar_domain.NewRegistrarDomainDataSource,
		registrar_domain.NewRegistrarDomainsDataSource,
		list.NewListDataSource,
		list.NewListsDataSource,
		list_item.NewListItemDataSource,
		list_item.NewListItemsDataSource,
		stream.NewStreamDataSource,
		stream.NewStreamsDataSource,
		stream_audio_track.NewStreamAudioTrackDataSource,
		stream_key.NewStreamKeyDataSource,
		stream_live_input.NewStreamLiveInputDataSource,
		stream_watermark.NewStreamWatermarkDataSource,
		stream_watermark.NewStreamWatermarksDataSource,
		stream_webhook.NewStreamWebhookDataSource,
		stream_caption_language.NewStreamCaptionLanguageDataSource,
		stream_download.NewStreamDownloadDataSource,
		notification_policy_webhooks.NewNotificationPolicyWebhooksDataSource,
		notification_policy_webhooks.NewNotificationPolicyWebhooksListDataSource,
		notification_policy.NewNotificationPolicyDataSource,
		notification_policy.NewNotificationPoliciesDataSource,
		d1_database.NewD1DatabaseDataSource,
		d1_database.NewD1DatabasesDataSource,
		r2_bucket.NewR2BucketDataSource,
		r2_bucket_lifecycle.NewR2BucketLifecycleDataSource,
		r2_bucket_cors.NewR2BucketCORSDataSource,
		r2_custom_domain.NewR2CustomDomainDataSource,
		r2_bucket_event_notification.NewR2BucketEventNotificationDataSource,
		r2_bucket_lock.NewR2BucketLockDataSource,
		r2_bucket_sippy.NewR2BucketSippyDataSource,
		workers_for_platforms_dispatch_namespace.NewWorkersForPlatformsDispatchNamespaceDataSource,
		workers_for_platforms_dispatch_namespace.NewWorkersForPlatformsDispatchNamespacesDataSource,
		workers_secret.NewWorkersSecretDataSource,
		workers_secret.NewWorkersSecretsDataSource,
		zero_trust_dex_test.NewZeroTrustDEXTestDataSource,
		zero_trust_dex_test.NewZeroTrustDEXTestsDataSource,
		zero_trust_device_managed_networks.NewZeroTrustDeviceManagedNetworksDataSource,
		zero_trust_device_managed_networks.NewZeroTrustDeviceManagedNetworksListDataSource,
		zero_trust_device_default_profile.NewZeroTrustDeviceDefaultProfileDataSource,
		zero_trust_device_default_profile_local_domain_fallback.NewZeroTrustDeviceDefaultProfileLocalDomainFallbackDataSource,
		zero_trust_device_default_profile_certificates.NewZeroTrustDeviceDefaultProfileCertificatesDataSource,
		zero_trust_device_custom_profile.NewZeroTrustDeviceCustomProfileDataSource,
		zero_trust_device_custom_profile.NewZeroTrustDeviceCustomProfilesDataSource,
		zero_trust_device_custom_profile_local_domain_fallback.NewZeroTrustDeviceCustomProfileLocalDomainFallbackDataSource,
		zero_trust_device_posture_rule.NewZeroTrustDevicePostureRuleDataSource,
		zero_trust_device_posture_rule.NewZeroTrustDevicePostureRulesDataSource,
		zero_trust_device_posture_integration.NewZeroTrustDevicePostureIntegrationDataSource,
		zero_trust_device_posture_integration.NewZeroTrustDevicePostureIntegrationsDataSource,
		zero_trust_access_identity_provider.NewZeroTrustAccessIdentityProviderDataSource,
		zero_trust_access_identity_provider.NewZeroTrustAccessIdentityProvidersDataSource,
		zero_trust_organization.NewZeroTrustOrganizationDataSource,
		zero_trust_access_infrastructure_target.NewZeroTrustAccessInfrastructureTargetDataSource,
		zero_trust_access_infrastructure_target.NewZeroTrustAccessInfrastructureTargetsDataSource,
		zero_trust_access_application.NewZeroTrustAccessApplicationDataSource,
		zero_trust_access_application.NewZeroTrustAccessApplicationsDataSource,
		zero_trust_access_short_lived_certificate.NewZeroTrustAccessShortLivedCertificateDataSource,
		zero_trust_access_short_lived_certificate.NewZeroTrustAccessShortLivedCertificatesDataSource,
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
		zero_trust_access_policy.NewZeroTrustAccessPolicyDataSource,
		zero_trust_access_policy.NewZeroTrustAccessPoliciesDataSource,
		zero_trust_tunnel_cloudflared.NewZeroTrustTunnelCloudflaredDataSource,
		zero_trust_tunnel_cloudflared.NewZeroTrustTunnelCloudflaredsDataSource,
		zero_trust_tunnel_cloudflared_config.NewZeroTrustTunnelCloudflaredConfigDataSource,
		zero_trust_tunnel_cloudflared_token.NewZeroTrustTunnelCloudflaredTokenDataSource,
		zero_trust_dlp_dataset.NewZeroTrustDLPDatasetDataSource,
		zero_trust_dlp_dataset.NewZeroTrustDLPDatasetsDataSource,
		zero_trust_dlp_custom_profile.NewZeroTrustDLPCustomProfileDataSource,
		zero_trust_dlp_predefined_profile.NewZeroTrustDLPPredefinedProfileDataSource,
		zero_trust_dlp_entry.NewZeroTrustDLPEntryDataSource,
		zero_trust_dlp_entry.NewZeroTrustDLPEntriesDataSource,
		zero_trust_gateway_categories.NewZeroTrustGatewayCategoriesListDataSource,
		zero_trust_gateway_app_types.NewZeroTrustGatewayAppTypesListDataSource,
		zero_trust_gateway_settings.NewZeroTrustGatewaySettingsDataSource,
		zero_trust_list.NewZeroTrustListDataSource,
		zero_trust_list.NewZeroTrustListsDataSource,
		zero_trust_dns_location.NewZeroTrustDNSLocationDataSource,
		zero_trust_dns_location.NewZeroTrustDNSLocationsDataSource,
		zero_trust_gateway_proxy_endpoint.NewZeroTrustGatewayProxyEndpointDataSource,
		zero_trust_gateway_policy.NewZeroTrustGatewayPolicyDataSource,
		zero_trust_gateway_policy.NewZeroTrustGatewayPoliciesDataSource,
		zero_trust_gateway_certificate.NewZeroTrustGatewayCertificateDataSource,
		zero_trust_gateway_certificate.NewZeroTrustGatewayCertificatesDataSource,
		zero_trust_tunnel_cloudflared_route.NewZeroTrustTunnelCloudflaredRouteDataSource,
		zero_trust_tunnel_cloudflared_route.NewZeroTrustTunnelCloudflaredRoutesDataSource,
		zero_trust_tunnel_cloudflared_virtual_network.NewZeroTrustTunnelCloudflaredVirtualNetworkDataSource,
		zero_trust_tunnel_cloudflared_virtual_network.NewZeroTrustTunnelCloudflaredVirtualNetworksDataSource,
		zero_trust_risk_behavior.NewZeroTrustRiskBehaviorDataSource,
		zero_trust_risk_scoring_integration.NewZeroTrustRiskScoringIntegrationDataSource,
		zero_trust_risk_scoring_integration.NewZeroTrustRiskScoringIntegrationsDataSource,
		turnstile_widget.NewTurnstileWidgetDataSource,
		turnstile_widget.NewTurnstileWidgetsDataSource,
		hyperdrive_config.NewHyperdriveConfigDataSource,
		hyperdrive_config.NewHyperdriveConfigsDataSource,
		web_analytics_site.NewWebAnalyticsSiteDataSource,
		web_analytics_site.NewWebAnalyticsSitesDataSource,
		bot_management.NewBotManagementDataSource,
		observatory_scheduled_test.NewObservatoryScheduledTestDataSource,
		dcv_delegation.NewDCVDelegationDataSource,
		hostname_tls_setting.NewHostnameTLSSettingDataSource,
		snippets.NewSnippetsDataSource,
		snippets.NewSnippetsListDataSource,
		snippet_rules.NewSnippetRulesListDataSource,
		calls_sfu_app.NewCallsSFUAppDataSource,
		calls_sfu_app.NewCallsSFUAppsDataSource,
		calls_turn_app.NewCallsTURNAppDataSource,
		calls_turn_app.NewCallsTURNAppsDataSource,
		cloudforce_one_request.NewCloudforceOneRequestDataSource,
		cloudforce_one_request.NewCloudforceOneRequestsDataSource,
		cloudforce_one_request_message.NewCloudforceOneRequestMessageDataSource,
		cloudforce_one_request_priority.NewCloudforceOneRequestPriorityDataSource,
		cloudforce_one_request_asset.NewCloudforceOneRequestAssetDataSource,
		permission_group.NewPermissionGroupDataSource,
		permission_group.NewPermissionGroupsDataSource,
		resource_group.NewResourceGroupDataSource,
		resource_group.NewResourceGroupsDataSource,
		cloud_connector_rules.NewCloudConnectorRulesListDataSource,
		botnet_feed_config_asn.NewBotnetFeedConfigASNDataSource,
		leaked_credential_check.NewLeakedCredentialCheckDataSource,
		leaked_credential_check_rule.NewLeakedCredentialCheckRulesDataSource,
		content_scanning_expression.NewContentScanningExpressionsDataSource,
	}
}

func NewProvider(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CloudflareProvider{
			version: version,
		}
	}
}
