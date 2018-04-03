package cloudflare

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_EMAIL", nil),
				Description: "A registered CloudFlare email address.",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_TOKEN", nil),
				Description: "The token key for API operations.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cloudflare_ip_ranges": dataSourceCloudflareIPRanges(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cloudflare_load_balancer_monitor":  resourceCloudFlareLoadBalancerMonitor(),
			"cloudflare_page_rule":              resourceCloudFlarePageRule(),
			"cloudflare_record":                 resourceCloudFlareRecord(),
			"cloudflare_rate_limit":             resourceCloudFlareRateLimit(),
			"cloudflare_zone_settings_override": resourceCloudFlareZoneSettingsOverride(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Email: d.Get("email").(string),
		Token: d.Get("token").(string),
	}

	return config.Client()
}
