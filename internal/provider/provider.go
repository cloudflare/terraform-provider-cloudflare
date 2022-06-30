package provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/meta"
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

		if s.ConflictsWith != nil && len(s.ConflictsWith) > 0 {
			conflicts := make([]string, len(s.ConflictsWith))
			for i, c := range s.ConflictsWith {
				conflicts[i] = fmt.Sprintf("`%s`", c)
			}
			desc += fmt.Sprintf(" Conflicts with %s.", strings.Join(conflicts, ", "))
		}

		return strings.TrimSpace(desc)
	}
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"email": {
					Type:          schema.TypeString,
					Optional:      true,
					DefaultFunc:   schema.EnvDefaultFunc("CLOUDFLARE_EMAIL", nil),
					Description:   "A registered Cloudflare email address. Alternatively, can be configured using the `CLOUDFLARE_EMAIL` environment variable.",
					ConflictsWith: []string{"api_token"},
					RequiredWith:  []string{"api_key"},
				},

				"api_key": {
					Type:         schema.TypeString,
					Optional:     true,
					DefaultFunc:  schema.EnvDefaultFunc("CLOUDFLARE_API_KEY", nil),
					Description:  "The API key for operations. Alternatively, can be configured using the `CLOUDFLARE_API_KEY` environment variable. API keys are [now considered legacy by Cloudflare](https://developers.cloudflare.com/api/keys/#limitations), API tokens should be used instead.",
					ExactlyOneOf: []string{"api_key", "api_token"},
					ValidateFunc: validation.StringMatch(regexp.MustCompile("[0-9a-f]{37}"), "API key must only contain characters 0-9 and a-f (all lowercased)"),
				},

				"api_token": {
					Type:         schema.TypeString,
					Optional:     true,
					DefaultFunc:  schema.EnvDefaultFunc("CLOUDFLARE_API_TOKEN", nil),
					Description:  "The API Token for operations. Alternatively, can be configured using the `CLOUDFLARE_API_TOKEN` environment variable.",
					ValidateFunc: validation.StringMatch(regexp.MustCompile("[A-Za-z0-9-_]{40}"), "API tokens must only contain characters a-z, A-Z, 0-9, hyphens and underscores"),
				},

				"api_user_service_key": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_USER_SERVICE_KEY", nil),
					Description: "A special Cloudflare API key good for a restricted set of endpoints. Alternatively, can be configured using the `CLOUDFLARE_API_USER_SERVICE_KEY` environment variable.",
				},

				"rps": {
					Type:        schema.TypeInt,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_RPS", 4),
					Description: "RPS limit to apply when making calls to the API. Alternatively, can be configured using the `CLOUDFLARE_RPS` environment variable.",
				},

				"retries": {
					Type:        schema.TypeInt,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_RETRIES", 3),
					Description: "Maximum number of retries to perform when an API request fails. Alternatively, can be configured using the `CLOUDFLARE_RETRIES` environment variable.",
				},

				"min_backoff": {
					Type:        schema.TypeInt,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_MIN_BACKOFF", 1),
					Description: "Minimum backoff period in seconds after failed API calls. Alternatively, can be configured using the `CLOUDFLARE_MIN_BACKOFF` environment variable.",
				},

				"max_backoff": {
					Type:        schema.TypeInt,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_MAX_BACKOFF", 30),
					Description: "Maximum backoff period in seconds after failed API calls. Alternatively, can be configured using the `CLOUDFLARE_MAX_BACKOFF` environment variable.",
				},

				"api_client_logging": {
					Type:        schema.TypeBool,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_CLIENT_LOGGING", false),
					Description: "Whether to print logs from the API client (using the default log library logger). Alternatively, can be configured using the `CLOUDFLARE_API_CLIENT_LOGGING` environment variable.",
				},

				"account_id": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_ACCOUNT_ID", nil),
					Description: "Configure API client to always use a specific account. Alternatively, can be configured using the `CLOUDFLARE_ACCOUNT_ID` environment variable.",
					Deprecated:  "Use resource specific `account_id` attributes instead.",
				},

				"api_hostname": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_HOSTNAME", "api.cloudflare.com"),
					Description: "Configure the hostname used by the API client. Alternatively, can be configured using the `CLOUDFLARE_API_HOSTNAME` environment variable.",
				},

				"api_base_path": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_BASE_PATH", "/client/v4"),
					Description: "Configure the base path used by the API client. Alternatively, can be configured using the `CLOUDFLARE_API_BASE_PATH` environment variable.",
				},
			},

			DataSourcesMap: map[string]*schema.Resource{
				"cloudflare_access_identity_provider":    dataSourceCloudflareAccessIdentityProvider(),
				"cloudflare_account_roles":               dataSourceCloudflareAccountRoles(),
				"cloudflare_api_token_permission_groups": dataSourceCloudflareApiTokenPermissionGroups(),
				"cloudflare_devices":                     dataSourceCloudflareDevices(),
				"cloudflare_ip_ranges":                   dataSourceCloudflareIPRanges(),
				"cloudflare_origin_ca_root_certificate":  dataSourceCloudflareOriginCARootCertificate(),
				"cloudflare_waf_groups":                  dataSourceCloudflareWAFGroups(),
				"cloudflare_waf_packages":                dataSourceCloudflareWAFPackages(),
				"cloudflare_waf_rules":                   dataSourceCloudflareWAFRules(),
				"cloudflare_zone_dnssec":                 dataSourceCloudflareZoneDNSSEC(),
				"cloudflare_zone":                        dataSourceCloudflareZone(),
				"cloudflare_zones":                       dataSourceCloudflareZones(),
			},

			ResourcesMap: map[string]*schema.Resource{
				"cloudflare_access_application":                     resourceCloudflareAccessApplication(),
				"cloudflare_access_ca_certificate":                  resourceCloudflareAccessCACertificate(),
				"cloudflare_access_group":                           resourceCloudflareAccessGroup(),
				"cloudflare_access_identity_provider":               resourceCloudflareAccessIdentityProvider(),
				"cloudflare_access_keys_configuration":              resourceCloudflareAccessKeysConfiguration(),
				"cloudflare_access_mutual_tls_certificate":          resourceCloudflareAccessMutualTLSCertificate(),
				"cloudflare_access_policy":                          resourceCloudflareAccessPolicy(),
				"cloudflare_access_rule":                            resourceCloudflareAccessRule(),
				"cloudflare_access_service_token":                   resourceCloudflareAccessServiceToken(),
				"cloudflare_access_bookmark":                        resourceCloudflareAccessBookmark(),
				"cloudflare_account_member":                         resourceCloudflareAccountMember(),
				"cloudflare_api_token":                              resourceCloudflareApiToken(),
				"cloudflare_argo_tunnel":                            resourceCloudflareArgoTunnel(),
				"cloudflare_argo":                                   resourceCloudflareArgo(),
				"cloudflare_authenticated_origin_pulls_certificate": resourceCloudflareAuthenticatedOriginPullsCertificate(),
				"cloudflare_authenticated_origin_pulls":             resourceCloudflareAuthenticatedOriginPulls(),
				"cloudflare_byo_ip_prefix":                          resourceCloudflareBYOIPPrefix(),
				"cloudflare_certificate_pack":                       resourceCloudflareCertificatePack(),
				"cloudflare_custom_hostname_fallback_origin":        resourceCloudflareCustomHostnameFallbackOrigin(),
				"cloudflare_custom_hostname":                        resourceCloudflareCustomHostname(),
				"cloudflare_custom_pages":                           resourceCloudflareCustomPages(),
				"cloudflare_custom_ssl":                             resourceCloudflareCustomSsl(),
				"cloudflare_device_posture_rule":                    resourceCloudflareDevicePostureRule(),
				"cloudflare_device_policy_certificates":             resourceCloudflareDevicePolicyCertificates(),
				"cloudflare_device_posture_integration":             resourceCloudflareDevicePostureIntegration(),
				"cloudflare_fallback_domain":                        resourceCloudflareFallbackDomain(),
				"cloudflare_filter":                                 resourceCloudflareFilter(),
				"cloudflare_firewall_rule":                          resourceCloudflareFirewallRule(),
				"cloudflare_gre_tunnel":                             resourceCloudflareGRETunnel(),
				"cloudflare_healthcheck":                            resourceCloudflareHealthcheck(),
				"cloudflare_ip_list":                                resourceCloudflareIPList(),
				"cloudflare_ipsec_tunnel":                           resourceCloudflareIPsecTunnel(),
				"cloudflare_list":                                   resourceCloudflareList(),
				"cloudflare_load_balancer_monitor":                  resourceCloudflareLoadBalancerMonitor(),
				"cloudflare_load_balancer_pool":                     resourceCloudflareLoadBalancerPool(),
				"cloudflare_load_balancer":                          resourceCloudflareLoadBalancer(),
				"cloudflare_logpull_retention":                      resourceCloudflareLogpullRetention(),
				"cloudflare_logpush_job":                            resourceCloudflareLogpushJob(),
				"cloudflare_logpush_ownership_challenge":            resourceCloudflareLogpushOwnershipChallenge(),
				"cloudflare_magic_firewall_ruleset":                 resourceCloudflareMagicFirewallRuleset(),
				"cloudflare_managed_headers":                        resourceCloudflareManagedHeaders(),
				"cloudflare_notification_policy_webhooks":           resourceCloudflareNotificationPolicyWebhooks(),
				"cloudflare_notification_policy":                    resourceCloudflareNotificationPolicy(),
				"cloudflare_origin_ca_certificate":                  resourceCloudflareOriginCACertificate(),
				"cloudflare_page_rule":                              resourceCloudflarePageRule(),
				"cloudflare_rate_limit":                             resourceCloudflareRateLimit(),
				"cloudflare_record":                                 resourceCloudflareRecord(),
				"cloudflare_ruleset":                                resourceCloudflareRuleset(),
				"cloudflare_spectrum_application":                   resourceCloudflareSpectrumApplication(),
				"cloudflare_split_tunnel":                           resourceCloudflareSplitTunnel(),
				"cloudflare_static_route":                           resourceCloudflareStaticRoute(),
				"cloudflare_teams_account":                          resourceCloudflareTeamsAccount(),
				"cloudflare_teams_list":                             resourceCloudflareTeamsList(),
				"cloudflare_teams_location":                         resourceCloudflareTeamsLocation(),
				"cloudflare_teams_rule":                             resourceCloudflareTeamsRule(),
				"cloudflare_teams_proxy_endpoint":                   resourceCloudflareTeamsProxyEndpoint(),
				"cloudflare_tunnel_route":                           resourceCloudflareTunnelRoute(),
				"cloudflare_tunnel_virtual_network":                 resourceCloudflareTunnelVirtualNetwork(),
				"cloudflare_waf_group":                              resourceCloudflareWAFGroup(),
				"cloudflare_waf_override":                           resourceCloudflareWAFOverride(),
				"cloudflare_waf_package":                            resourceCloudflareWAFPackage(),
				"cloudflare_waf_rule":                               resourceCloudflareWAFRule(),
				"cloudflare_waiting_room":                           resourceCloudflareWaitingRoom(),
				"cloudflare_waiting_room_event":                     resourceCloudflareWaitingRoomEvent(),
				"cloudflare_worker_cron_trigger":                    resourceCloudflareWorkerCronTrigger(),
				"cloudflare_worker_route":                           resourceCloudflareWorkerRoute(),
				"cloudflare_worker_script":                          resourceCloudflareWorkerScript(),
				"cloudflare_workers_kv_namespace":                   resourceCloudflareWorkersKVNamespace(),
				"cloudflare_workers_kv":                             resourceCloudflareWorkerKV(),
				"cloudflare_zone_cache_variants":                    resourceCloudflareZoneCacheVariants(),
				"cloudflare_zone_dnssec":                            resourceCloudflareZoneDNSSEC(),
				"cloudflare_zone_lockdown":                          resourceCloudflareZoneLockdown(),
				"cloudflare_zone_settings_override":                 resourceCloudflareZoneSettingsOverride(),
				"cloudflare_zone":                                   resourceCloudflareZone(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {

	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		var diags diag.Diagnostics

		baseURL := cloudflare.BaseURL(
			"https://" + d.Get("api_hostname").(string) + d.Get("api_base_path").(string),
		)
		limitOpt := cloudflare.UsingRateLimit(float64(d.Get("rps").(int)))
		retryOpt := cloudflare.UsingRetryPolicy(d.Get("retries").(int), d.Get("min_backoff").(int), d.Get("max_backoff").(int))
		options := []cloudflare.Option{limitOpt, retryOpt, baseURL}

		if d.Get("api_client_logging").(bool) {
			options = append(options, cloudflare.UsingLogger(log.New(os.Stderr, "", log.LstdFlags)))
		}

		c := cleanhttp.DefaultClient()
		c.Transport = logging.NewTransport("Cloudflare", c.Transport)
		options = append(options, cloudflare.HTTPClient(c))

		ua := fmt.Sprintf("terraform/%s terraform-plugin-sdk/%s terraform-provider-cloudflare/%s", p.TerraformVersion, meta.SDKVersionString(), version)
		options = append(options, cloudflare.UserAgent(ua))

		config := Config{Options: options}

		if v, ok := d.GetOk("api_token"); ok {
			config.APIToken = v.(string)
		} else if v, ok := d.GetOk("api_key"); ok {
			config.APIKey = v.(string)
			if v, ok = d.GetOk("email"); ok {
				config.Email = v.(string)
			} else {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "email is not set correctly",
				})

				return nil, diags
			}
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "credentials are not set correctly",
			})
			return nil, diags
		}

		if v, ok := d.GetOk("api_user_service_key"); ok {
			config.APIUserServiceKey = v.(string)
		}

		client, err := config.Client()
		if err != nil {
			return nil, diag.FromErr(err)
		}

		if accountID, ok := d.GetOk("account_id"); ok {
			tflog.Info(ctx, fmt.Sprintf("using specified account id %s in Cloudflare provider", accountID.(string)))
			options = append(options, cloudflare.UsingAccount(accountID.(string)))
		} else {
			return client, diag.FromErr(err)
		}

		config.Options = options

		client, err = config.Client()
		if err != nil {
			return nil, diag.FromErr(err)
		}

		return client, nil
	}
}
