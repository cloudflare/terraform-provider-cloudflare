package sdkv2provider

import (
	"bytes"
	"context"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const (
	MAXIMUM_NUMBER_OF_ENTITIES_REACHED_SUMMARY = "You've attempted to add a new %[1]s to the `terraform-plugin-sdkv2` which is no longer considered suitable for use."
	MAXIMUM_NUMBER_OF_ENTITIES_REACHED_DETAIL  = "Due the number of known internal issues with `terraform-plugin-sdkv2` (most notably handling of zero values), we are no longer recommending using it and instead, advise using `terraform-plugin-framework` exclusively. If you must use terraform-plugin-sdkv2 for this new %[1]s you should first discuss it with a maintainer to fully understand the impact and potential ramifications. Only then should you bump %[2]s to include your %[1]s."
	MAXIMUM_ALLOWED_SDKV2_RESOURCES            = 108
	MAXIMUM_ALLOWED_SDKV2_DATASOURCES          = 19
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown

	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		desc = strings.TrimSpace(desc)

		if !bytes.HasSuffix([]byte(s.Description), []byte(".")) && s.Description != "" {
			desc += "."
		}

		if s.Default != nil {
			if s.Default == "" {
				desc += " Defaults to `\"\"`."
			} else {
				desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
			}
		}

		if s.RequiredWith != nil && len(s.RequiredWith) > 0 && !contains(s.RequiredWith, consts.APIKeySchemaKey) {
			requiredWith := make([]string, len(s.RequiredWith))
			for i, c := range s.RequiredWith {
				requiredWith[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Required when using %s.", strings.Join(requiredWith, ", "))
		}

		if s.ConflictsWith != nil && len(s.ConflictsWith) > 0 && !contains(s.ConflictsWith, consts.APITokenSchemaKey) {
			conflicts := make([]string, len(s.ConflictsWith))
			for i, c := range s.ConflictsWith {
				conflicts[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Conflicts with %s.", strings.Join(conflicts, ", "))
		}

		if s.ExactlyOneOf != nil && len(s.ExactlyOneOf) > 0 && (!contains(s.ExactlyOneOf, consts.APIKeySchemaKey) || !contains(s.ExactlyOneOf, consts.APITokenSchemaKey) || !contains(s.ExactlyOneOf, consts.APIUserServiceKeySchemaKey)) {
			exactlyOneOfs := make([]string, len(s.ExactlyOneOf))
			for i, c := range s.ExactlyOneOf {
				exactlyOneOfs[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Must provide only one of %s.", strings.Join(exactlyOneOfs, ", "))
		}

		if s.AtLeastOneOf != nil && len(s.AtLeastOneOf) > 0 {
			atLeastOneOfs := make([]string, len(s.AtLeastOneOf))
			for i, c := range s.AtLeastOneOf {
				atLeastOneOfs[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Must provide at least one of %s.", strings.Join(atLeastOneOfs, ", "))
		}

		if s.ForceNew {
			desc += " **Modifying this attribute will force creation of a new resource.**"
		}

		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				consts.EmailSchemaKey: {
					Type:          schema.TypeString,
					Optional:      true,
					Description:   fmt.Sprintf("A registered Cloudflare email address. Alternatively, can be configured using the `%s` environment variable. Required when using `api_key`. Conflicts with `api_token`.", consts.EmailEnvVarKey),
					ConflictsWith: []string{consts.APITokenSchemaKey},
					RequiredWith:  []string{consts.APIKeySchemaKey},
				},

				consts.APIKeySchemaKey: {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  fmt.Sprintf("The API key for operations. Alternatively, can be configured using the `%s` environment variable. API keys are [now considered legacy by Cloudflare](https://developers.cloudflare.com/fundamentals/api/get-started/keys/#limitations), API tokens should be used instead. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APIKeyEnvVarKey),
					ValidateFunc: validation.StringMatch(regexp.MustCompile("[0-9a-f]{37}"), "API key must be 37 characters long and only contain characters 0-9 and a-f (all lowercased)"),
				},

				consts.APITokenSchemaKey: {
					Type:         schema.TypeString,
					Optional:     true,
					Description:  fmt.Sprintf("The API Token for operations. Alternatively, can be configured using the `%s` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APITokenEnvVarKey),
					ValidateFunc: validation.StringMatch(regexp.MustCompile("[A-Za-z0-9-_]{40}"), "API tokens must be 40 characters long and only contain characters a-z, A-Z, 0-9, hyphens and underscores"),
				},

				consts.APIUserServiceKeySchemaKey: {
					Type:        schema.TypeString,
					Optional:    true,
					Description: fmt.Sprintf("A special Cloudflare API key good for a restricted set of endpoints. Alternatively, can be configured using the `%s` environment variable. Must provide only one of `api_key`, `api_token`, `api_user_service_key`.", consts.APIUserServiceKeyEnvVarKey),
				},

				consts.RPSSchemaKey: {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: fmt.Sprintf("RPS limit to apply when making calls to the API. Alternatively, can be configured using the `%s` environment variable.", consts.RPSEnvVarKey),
				},

				consts.RetriesSchemaKey: {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: fmt.Sprintf("Maximum number of retries to perform when an API request fails. Alternatively, can be configured using the `%s` environment variable.", consts.RetriesEnvVarKey),
				},

				consts.MinimumBackoffSchemaKey: {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: fmt.Sprintf("Minimum backoff period in seconds after failed API calls. Alternatively, can be configured using the `%s` environment variable.", consts.MinimumBackoffEnvVar),
				},

				consts.MaximumBackoffSchemaKey: {
					Type:        schema.TypeInt,
					Optional:    true,
					Description: fmt.Sprintf("Maximum backoff period in seconds after failed API calls. Alternatively, can be configured using the `%s` environment variable.", consts.MaximumBackoffEnvVarKey),
				},

				consts.APIClientLoggingSchemaKey: {
					Type:        schema.TypeBool,
					Optional:    true,
					Description: fmt.Sprintf("Whether to print logs from the API client (using the default log library logger). Alternatively, can be configured using the `%s` environment variable.", consts.APIClientLoggingEnvVarKey),
				},

				consts.APIHostnameSchemaKey: {
					Type:        schema.TypeString,
					Optional:    true,
					Description: fmt.Sprintf("Configure the hostname used by the API client. Alternatively, can be configured using the `%s` environment variable.", consts.APIHostnameEnvVarKey),
				},

				consts.APIBasePathSchemaKey: {
					Type:        schema.TypeString,
					Optional:    true,
					Description: fmt.Sprintf("Configure the base path used by the API client. Alternatively, can be configured using the `%s` environment variable.", consts.APIBasePathEnvVarKey),
				},

				consts.UserAgentOperatorSuffixSchemaKey: {
					Type:        schema.TypeString,
					Optional:    true,
					Description: fmt.Sprintf("A value to append to the HTTP User Agent for all API calls. This value is not something most users need to modify however, if you are using a non-standard provider or operator configuration, this is recommended to assist in uniquely identifying your traffic. **Setting this value will remove the Terraform version from the HTTP User Agent string and may have unintended consequences**. Alternatively, can be configured using the `%s` environment variable.", consts.UserAgentOperatorSuffixEnvVarKey),
				},
			},

			DataSourcesMap: map[string]*schema.Resource{
				"cloudflare_access_application":         dataSourceCloudflareAccessApplication(),
				"cloudflare_access_identity_provider":   dataSourceCloudflareAccessIdentityProvider(),
				"cloudflare_account_roles":              dataSourceCloudflareAccountRoles(),
				"cloudflare_accounts":                   dataSourceCloudflareAccounts(),
				"cloudflare_devices":                    dataSourceCloudflareDevices(),
				"cloudflare_device_posture_rules":       dataSourceCloudflareDevicePostureRules(),
				"cloudflare_ip_ranges":                  dataSourceCloudflareIPRanges(),
				"cloudflare_list":                       dataSourceCloudflareList(),
				"cloudflare_lists":                      dataSourceCloudflareLists(),
				"cloudflare_tunnel_virtual_network":     dataSourceCloudflareTunnelVirtualNetwork(),
				"cloudflare_load_balancer_pools":        dataSourceCloudflareLoadBalancerPools(),
				"cloudflare_origin_ca_root_certificate": dataSourceCloudflareOriginCARootCertificate(),
				"cloudflare_record":                     dataSourceCloudflareRecord(),
				"cloudflare_rulesets":                   dataSourceCloudflareRulesets(),
				"cloudflare_zone_cache_reserve":         dataSourceCloudflareZoneCacheReserve(),
				"cloudflare_tunnel":                     dataSourceCloudflareTunnel(),
				"cloudflare_zone_dnssec":                dataSourceCloudflareZoneDNSSEC(),
				"cloudflare_zone":                       dataSourceCloudflareZone(),
				"cloudflare_zones":                      dataSourceCloudflareZones(),
			},

			ResourcesMap: map[string]*schema.Resource{
				"cloudflare_access_application":                              resourceCloudflareAccessApplication(),
				"cloudflare_access_ca_certificate":                           resourceCloudflareAccessCACertificate(),
				"cloudflare_access_group":                                    resourceCloudflareAccessGroup(),
				"cloudflare_access_identity_provider":                        resourceCloudflareAccessIdentityProvider(),
				"cloudflare_access_custom_page":                              resourceCloudflareAccessCustomPage(),
				"cloudflare_access_keys_configuration":                       resourceCloudflareAccessKeysConfiguration(),
				"cloudflare_access_mutual_tls_certificate":                   resourceCloudflareAccessMutualTLSCertificate(),
				"cloudflare_access_organization":                             resourceCloudflareAccessOrganization(),
				"cloudflare_access_policy":                                   resourceCloudflareAccessPolicy(),
				"cloudflare_access_rule":                                     resourceCloudflareAccessRule(),
				"cloudflare_access_service_token":                            resourceCloudflareAccessServiceToken(),
				"cloudflare_access_tag":                                      resourceCloudflareAccessTag(),
				"cloudflare_account_member":                                  resourceCloudflareAccountMember(),
				"cloudflare_account":                                         resourceCloudflareAccount(),
				"cloudflare_address_map":                                     resourceCloudflareAddressMap(),
				"cloudflare_api_shield":                                      resourceCloudflareAPIShield(),
				"cloudflare_api_shield_operation":                            resourceCloudflareAPIShieldOperation(),
				"cloudflare_api_shield_operation_schema_validation_settings": resourceCloudflareAPIShieldOperationSchemaValidationSettings(),
				"cloudflare_api_shield_schema":                               resourceCloudflareAPIShieldSchemas(),
				"cloudflare_api_shield_schema_validation_settings":           resourceCloudflareAPIShieldSchemaValidationSettings(),
				"cloudflare_api_token":                                       resourceCloudflareApiToken(),
				"cloudflare_argo":                                            resourceCloudflareArgo(),
				"cloudflare_authenticated_origin_pulls_certificate":          resourceCloudflareAuthenticatedOriginPullsCertificate(),
				"cloudflare_authenticated_origin_pulls":                      resourceCloudflareAuthenticatedOriginPulls(),
				"cloudflare_byo_ip_prefix":                                   resourceCloudflareBYOIPPrefix(),
				"cloudflare_certificate_pack":                                resourceCloudflareCertificatePack(),
				"cloudflare_custom_hostname_fallback_origin":                 resourceCloudflareCustomHostnameFallbackOrigin(),
				"cloudflare_custom_hostname":                                 resourceCloudflareCustomHostname(),
				"cloudflare_custom_pages":                                    resourceCloudflareCustomPages(),
				"cloudflare_custom_ssl":                                      resourceCloudflareCustomSsl(),
				"cloudflare_device_dex_test":                                 resourceCloudflareDeviceDexTest(),
				"cloudflare_device_managed_networks":                         resourceCloudflareDeviceManagedNetworks(),
				"cloudflare_device_policy_certificates":                      resourceCloudflareDevicePolicyCertificates(),
				"cloudflare_device_posture_integration":                      resourceCloudflareDevicePostureIntegration(),
				"cloudflare_device_posture_rule":                             resourceCloudflareDevicePostureRule(),
				"cloudflare_device_settings_policy":                          resourceCloudflareDeviceSettingsPolicy(),
				"cloudflare_dlp_profile":                                     resourceCloudflareDLPProfile(),
				"cloudflare_email_routing_catch_all":                         resourceCloudflareEmailRoutingCatchAll(),
				"cloudflare_email_routing_settings":                          resourceCloudflareEmailRoutingSettings(),
				"cloudflare_fallback_domain":                                 resourceCloudflareFallbackDomain(),
				"cloudflare_filter":                                          resourceCloudflareFilter(),
				"cloudflare_firewall_rule":                                   resourceCloudflareFirewallRule(),
				"cloudflare_gre_tunnel":                                      resourceCloudflareGRETunnel(),
				"cloudflare_healthcheck":                                     resourceCloudflareHealthcheck(),
				"cloudflare_hostname_tls_setting":                            resourceCloudflareHostnameTLSSetting(),
				"cloudflare_hostname_tls_setting_ciphers":                    resourceCloudflareHostnameTLSSettingCiphers(),
				"cloudflare_ipsec_tunnel":                                    resourceCloudflareIPsecTunnel(),
				"cloudflare_keyless_certificate":                             resourceCloudflareKeylessCertificate(),
				"cloudflare_list":                                            resourceCloudflareList(),
				"cloudflare_load_balancer_monitor":                           resourceCloudflareLoadBalancerMonitor(),
				"cloudflare_load_balancer_pool":                              resourceCloudflareLoadBalancerPool(),
				"cloudflare_load_balancer":                                   resourceCloudflareLoadBalancer(),
				"cloudflare_logpull_retention":                               resourceCloudflareLogpullRetention(),
				"cloudflare_logpush_job":                                     resourceCloudflareLogpushJob(),
				"cloudflare_logpush_ownership_challenge":                     resourceCloudflareLogpushOwnershipChallenge(),
				"cloudflare_magic_firewall_ruleset":                          resourceCloudflareMagicFirewallRuleset(),
				"cloudflare_managed_headers":                                 resourceCloudflareManagedHeaders(),
				"cloudflare_mtls_certificate":                                resourceCloudflareMTLSCertificate(),
				"cloudflare_notification_policy_webhooks":                    resourceCloudflareNotificationPolicyWebhook(),
				"cloudflare_notification_policy":                             resourceCloudflareNotificationPolicy(),
				"cloudflare_observatory_scheduled_test":                      resourceCloudflareObservatoryScheduledTest(),
				"cloudflare_origin_ca_certificate":                           resourceCloudflareOriginCACertificate(),
				"cloudflare_page_rule":                                       resourceCloudflarePageRule(),
				"cloudflare_pages_domain":                                    resourceCloudflarePagesDomain(),
				"cloudflare_pages_project":                                   resourceCloudflarePagesProject(),
				"cloudflare_queue":                                           resourceCloudflareQueue(),
				"cloudflare_rate_limit":                                      resourceCloudflareRateLimit(),
				"cloudflare_record":                                          resourceCloudflareRecord(),
				"cloudflare_regional_hostname":                               resourceCloudflareRegionalHostname(),
				"cloudflare_regional_tiered_cache":                           resourceCloudflareRegionalTieredCache(),
				"cloudflare_spectrum_application":                            resourceCloudflareSpectrumApplication(),
				"cloudflare_split_tunnel":                                    resourceCloudflareSplitTunnel(),
				"cloudflare_static_route":                                    resourceCloudflareStaticRoute(),
				"cloudflare_bot_management":                                  resourceCloudflareBotManagement(),
				"cloudflare_teams_account":                                   resourceCloudflareTeamsAccount(),
				"cloudflare_teams_list":                                      resourceCloudflareTeamsList(),
				"cloudflare_teams_location":                                  resourceCloudflareTeamsLocation(),
				"cloudflare_teams_proxy_endpoint":                            resourceCloudflareTeamsProxyEndpoint(),
				"cloudflare_teams_rule":                                      resourceCloudflareTeamsRule(),
				"cloudflare_tiered_cache":                                    resourceCloudflareTieredCache(),
				"cloudflare_total_tls":                                       resourceCloudflareTotalTLS(),
				"cloudflare_tunnel_config":                                   resourceCloudflareTunnelConfig(),
				"cloudflare_tunnel_route":                                    resourceCloudflareTunnelRoute(),
				"cloudflare_tunnel_virtual_network":                          resourceCloudflareTunnelVirtualNetwork(),
				"cloudflare_tunnel":                                          resourceCloudflareTunnel(),
				"cloudflare_url_normalization_settings":                      resourceCloudflareURLNormalizationSettings(),
				"cloudflare_user_agent_blocking_rule":                        resourceCloudflareUserAgentBlockingRules(),
				"cloudflare_waiting_room_event":                              resourceCloudflareWaitingRoomEvent(),
				"cloudflare_waiting_room_rules":                              resourceCloudflareWaitingRoomRules(),
				"cloudflare_waiting_room_settings":                           resourceCloudflareWaitingRoomSettings(),
				"cloudflare_waiting_room":                                    resourceCloudflareWaitingRoom(),
				"cloudflare_web3_hostname":                                   resourceCloudflareWeb3Hostname(),
				"cloudflare_web_analytics_rule":                              resourceCloudflareWebAnalyticsRule(),
				"cloudflare_web_analytics_site":                              resourceCloudflareWebAnalyticsSite(),
				"cloudflare_worker_cron_trigger":                             resourceCloudflareWorkerCronTrigger(),
				"cloudflare_worker_domain":                                   resourceCloudflareWorkerDomain(),
				"cloudflare_worker_route":                                    resourceCloudflareWorkerRoute(),
				"cloudflare_worker_script":                                   resourceCloudflareWorkerScript(),
				"cloudflare_worker_secret":                                   resourceCloudflareWorkerSecret(),
				"cloudflare_workers_kv_namespace":                            resourceCloudflareWorkersKVNamespace(),
				"cloudflare_workers_kv":                                      resourceCloudflareWorkerKV(),
				"cloudflare_zone_cache_reserve":                              resourceCloudflareZoneCacheReserve(),
				"cloudflare_zone_cache_variants":                             resourceCloudflareZoneCacheVariants(),
				"cloudflare_zone_dnssec":                                     resourceCloudflareZoneDNSSEC(),
				"cloudflare_zone_lockdown":                                   resourceCloudflareZoneLockdown(),
				"cloudflare_zone_settings_override":                          resourceCloudflareZoneSettingsOverride(),
				"cloudflare_zone_hold":                                       resourceCloudflareZoneHold(),
				"cloudflare_zone":                                            resourceCloudflareZone(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var (
			diags diag.Diagnostics

			email             string
			apiKey            string
			apiToken          string
			apiUserServiceKey string
			rps               int64
			retries           int64
			minBackOff        int64
			maxBackOff        int64
			baseHostname      string
			basePath          string
		)

		if len(p.ResourcesMap) > MAXIMUM_ALLOWED_SDKV2_RESOURCES {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf(MAXIMUM_NUMBER_OF_ENTITIES_REACHED_SUMMARY, "resource"),
				Detail:   fmt.Sprintf(MAXIMUM_NUMBER_OF_ENTITIES_REACHED_DETAIL, "resource", "MAXIMUM_ALLOWED_SDKV2_RESOURCES"),
			})

			return nil, diags
		}

		if len(p.DataSourcesMap) > MAXIMUM_ALLOWED_SDKV2_DATASOURCES {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf(MAXIMUM_NUMBER_OF_ENTITIES_REACHED_SUMMARY, "datasource"),
				Detail:   fmt.Sprintf(MAXIMUM_NUMBER_OF_ENTITIES_REACHED_DETAIL, "datasource", "MAXIMUM_ALLOWED_SDKV2_DATASOURCES"),
			})

			return nil, diags
		}

		if d.Get(consts.APIHostnameSchemaKey).(string) != "" {
			baseHostname = d.Get(consts.APIHostnameSchemaKey).(string)
		} else {
			baseHostname = utils.GetDefaultFromEnv(consts.APIHostnameEnvVarKey, consts.APIHostnameDefault)
		}

		if d.Get(consts.APIBasePathSchemaKey).(string) != "" {
			basePath = d.Get(consts.APIBasePathSchemaKey).(string)
		} else {
			basePath = utils.GetDefaultFromEnv(consts.APIBasePathEnvVarKey, consts.APIBasePathDefault)
		}
		baseURL := cloudflare.BaseURL(fmt.Sprintf("https://%s%s", baseHostname, basePath))

		if _, ok := d.GetOk(consts.RPSSchemaKey); ok {
			rps = int64(d.Get(consts.RPSSchemaKey).(int))
		} else {
			i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.RPSEnvVarKey, consts.RPSDefault), 10, 64)
			rps = i
		}
		limitOpt := cloudflare.UsingRateLimit(float64(rps))

		if _, ok := d.GetOk(consts.RetriesSchemaKey); ok {
			retries = int64(d.Get(consts.RetriesSchemaKey).(int))
		} else {
			i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.RetriesEnvVarKey, consts.RetriesDefault), 10, 64)
			retries = i
		}

		if _, ok := d.GetOk(consts.MinimumBackoffSchemaKey); ok {
			minBackOff = int64(d.Get(consts.MinimumBackoffSchemaKey).(int))
		} else {
			i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.MinimumBackoffEnvVar, consts.MinimumBackoffDefault), 10, 64)
			minBackOff = i
		}

		if _, ok := d.GetOk(consts.MaximumBackoffSchemaKey); ok {
			maxBackOff = int64(d.Get(consts.MaximumBackoffSchemaKey).(int))
		} else {
			i, _ := strconv.ParseInt(utils.GetDefaultFromEnv(consts.MaximumBackoffEnvVarKey, consts.MaximumBackoffDefault), 10, 64)
			maxBackOff = i
		}

		if retries >= math.MaxInt32 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("retries value of %d is too large, try a smaller value.", retries),
			})

			return nil, diags
		}

		if minBackOff >= math.MaxInt32 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("min_backoff value of %d is too large, try a smaller value.", minBackOff),
			})

			return nil, diags
		}

		if maxBackOff >= math.MaxInt32 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("max_backoff value of %d is too large, try a smaller value.", maxBackOff),
			})

			return nil, diags
		}

		retryOpt := cloudflare.UsingRetryPolicy(int(retries), int(minBackOff), int(maxBackOff))
		options := []cloudflare.Option{limitOpt, retryOpt, baseURL}

		options = append(options, cloudflare.Debug(logging.IsDebugOrHigher()))

		pluginVersion := utils.FindGoModuleVersion("github.com/hashicorp/terraform-plugin-sdk/v2")
		userAgentParams := utils.UserAgentBuilderParams{
			ProviderVersion: cloudflare.StringPtr(version),
			PluginType:      cloudflare.StringPtr("terraform-plugin-sdk"),
			PluginVersion:   pluginVersion,
		}
		if v, ok := d.GetOk(consts.UserAgentOperatorSuffixSchemaKey); ok {
			userAgentParams.OperatorSuffix = cloudflare.StringPtr(v.(string))
		} else {
			userAgentParams.TerraformVersion = cloudflare.StringPtr(p.TerraformVersion)
		}
		options = append(options, cloudflare.UserAgent(userAgentParams.String()))

		config := Config{Options: options}

		if v, ok := d.GetOk(consts.APITokenSchemaKey); ok {
			apiToken = v.(string)
		} else {
			apiToken = utils.GetDefaultFromEnv(consts.APITokenEnvVarKey, "")
		}

		if apiToken != "" {
			config.APIToken = apiToken
		}

		if v, ok := d.GetOk(consts.APIKeySchemaKey); ok {
			apiKey = v.(string)
		} else {
			apiKey = utils.GetDefaultFromEnv(consts.APIKeyEnvVarKey, "")
		}

		if apiKey != "" {
			config.APIKey = apiKey

			if v, ok := d.GetOk(consts.EmailSchemaKey); ok {
				email = v.(string)
			} else {
				email = utils.GetDefaultFromEnv(consts.EmailEnvVarKey, "")
			}

			if email == "" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("%q is not set correctly", consts.EmailSchemaKey),
				})

				return nil, diags
			}

			if email != "" {
				config.Email = email
			}
		}

		if v, ok := d.GetOk(consts.APIUserServiceKeySchemaKey); ok {
			apiUserServiceKey = v.(string)
		} else {
			apiUserServiceKey = utils.GetDefaultFromEnv(consts.APIUserServiceKeyEnvVarKey, "")
		}

		if apiUserServiceKey != "" {
			config.APIUserServiceKey = apiUserServiceKey
		}

		if apiKey == "" && apiToken == "" && apiUserServiceKey == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("must provide exactly one of %q, %q or %q.", consts.APIKeySchemaKey, consts.APITokenSchemaKey, consts.APIUserServiceKeySchemaKey),
			})
			return nil, diags
		}

		config.Options = options
		client, err := config.Client(ctx)
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return client, nil
	}
}
