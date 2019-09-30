package cloudflare

import (
	"fmt"
	"log"
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/httpclient"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-cloudflare/version"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_EMAIL", nil),
				Description: "A registered Cloudflare email address.",
			},

			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_KEY", nil),
				Description: "The API key for operations.",
			},

			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_TOKEN", nil),
				Description: "The API Token for operations.",
			},

			"rps": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_RPS", 4),
				Description: "RPS limit to apply when making calls to the API",
			},

			"retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_RETRIES", 3),
				Description: "Maximum number of retries to perform when an API request fails",
			},

			"min_backoff": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_MIN_BACKOFF", 1),
				Description: "Minimum backoff period in seconds after failed API calls",
			},

			"max_backoff": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_MAX_BACKOFF", 30),
				Description: "Maximum backoff period in seconds after failed API calls",
			},

			"api_client_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_CLIENT_LOGGING", false),
				Description: "Whether to print logs from the API client (using the default log library logger)",
			},

			"account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_ACCOUNT_ID", nil),
				Description: "Configure API client to always use that account.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cloudflare_ip_ranges": dataSourceCloudflareIPRanges(),
			"cloudflare_zones":     dataSourceCloudflareZones(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cloudflare_access_application":     resourceCloudflareAccessApplication(),
			"cloudflare_access_policy":          resourceCloudflareAccessPolicy(),
			"cloudflare_access_rule":            resourceCloudflareAccessRule(),
			"cloudflare_account_member":         resourceCloudflareAccountMember(),
			"cloudflare_argo":                   resourceCloudflareArgo(),
			"cloudflare_custom_pages":           resourceCloudflareCustomPages(),
			"cloudflare_custom_ssl":             resourceCloudflareCustomSsl(),
			"cloudflare_filter":                 resourceCloudflareFilter(),
			"cloudflare_firewall_rule":          resourceCloudflareFirewallRule(),
			"cloudflare_load_balancer_monitor":  resourceCloudflareLoadBalancerMonitor(),
			"cloudflare_load_balancer_pool":     resourceCloudflareLoadBalancerPool(),
			"cloudflare_load_balancer":          resourceCloudflareLoadBalancer(),
			"cloudflare_logpush_job":            resourceCloudflareLogpushJob(),
			"cloudflare_page_rule":              resourceCloudflarePageRule(),
			"cloudflare_rate_limit":             resourceCloudflareRateLimit(),
			"cloudflare_record":                 resourceCloudflareRecord(),
			"cloudflare_spectrum_application":   resourceCloudflareSpectrumApplication(),
			"cloudflare_waf_rule":               resourceCloudflareWAFRule(),
			"cloudflare_worker_route":           resourceCloudflareWorkerRoute(),
			"cloudflare_worker_script":          resourceCloudflareWorkerScript(),
			"cloudflare_zone_lockdown":          resourceCloudflareZoneLockdown(),
			"cloudflare_zone_settings_override": resourceCloudflareZoneSettingsOverride(),
			"cloudflare_zone":                   resourceCloudflareZone(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}

	return provider
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	limitOpt := cloudflare.UsingRateLimit(float64(d.Get("rps").(int)))
	retryOpt := cloudflare.UsingRetryPolicy(d.Get("retries").(int), d.Get("min_backoff").(int), d.Get("max_backoff").(int))
	options := []cloudflare.Option{limitOpt, retryOpt}

	if d.Get("api_client_logging").(bool) {
		options = append(options, cloudflare.UsingLogger(log.New(os.Stderr, "", log.LstdFlags)))
	}

	c := cleanhttp.DefaultClient()
	c.Transport = logging.NewTransport("Cloudflare", c.Transport)
	options = append(options, cloudflare.HTTPClient(c))

	tfUserAgent := httpclient.TerraformUserAgent(terraformVersion)
	providerUserAgent := fmt.Sprintf("terraform-provider-cloudflare/%s", version.ProviderVersion)
	ua := fmt.Sprintf("%s %s", tfUserAgent, providerUserAgent)
	options = append(options, cloudflare.UserAgent(ua))

	config := Config{Options: options}

	if v, ok := d.GetOk("api_token"); ok {
		config.APIToken = v.(string)
	} else if v, ok := d.GetOk("api_key"); ok {
		config.APIKey = v.(string)
		if v, ok = d.GetOk("email"); ok {
			config.Email = v.(string)
		} else {
			return nil, fmt.Errorf("email is not set correctly")
		}
	} else {
		return nil, fmt.Errorf("credentials are not set correctly")
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	if accountID, ok := d.GetOk("account_id"); ok {
		log.Printf("[INFO] Using specified account id %s in Cloudflare provider", accountID.(string))
		options = append(options, cloudflare.UsingAccount(accountID.(string)))
	} else {
		return client, err
	}

	config.Options = options

	client, err = config.Client()
	if err != nil {
		return nil, err
	}

	return client, err
}
