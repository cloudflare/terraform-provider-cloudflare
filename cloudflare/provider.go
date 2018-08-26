package cloudflare

import (
	"log"
	"os"

	"fmt"

	"github.com/cloudflare/cloudflare-go"
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
				Description: "A registered Cloudflare email address.",
			},

			"token": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_TOKEN", nil),
				Description: "The token key for API operations.",
			},

			"rps": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_RPS", 4),
				Description: "RPS limit to apply when making calls to the API",
			},

			"retries": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_RETRIES", 3),
				Description: "Maximum number of retries to perform when an API request fails",
			},

			"min_backoff": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_MIN_BACKOFF", 1),
				Description: "Minimum backoff period in seconds after failed API calls",
			},

			"max_backoff": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_MAX_BACKOFF", 30),
				Description: "Maximum backoff period in seconds after failed API calls",
			},

			"api_client_logging": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_API_CLIENT_LOGGING", false),
				Description: "Whether to print logs from the API client (using the default log library logger)",
			},

			"use_org_from_zone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_ORG_ZONE", nil),
				Description: "If specified zone is owned by an organization, configure API client to always use that organization",
			},

			"org_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDFLARE_ORG_ID", nil),
				Description: "Configure API client to always use that organization. If set this will override 'user_owner_from_zone'",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cloudflare_ip_ranges": dataSourceCloudflareIPRanges(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cloudflare_access_rule":            resourceCloudflareAccessRule(),
			"cloudflare_load_balancer_monitor":  resourceCloudflareLoadBalancerMonitor(),
			"cloudflare_load_balancer_pool":     resourceCloudflareLoadBalancerPool(),
			"cloudflare_load_balancer":          resourceCloudflareLoadBalancer(),
			"cloudflare_page_rule":              resourceCloudflarePageRule(),
			"cloudflare_rate_limit":             resourceCloudflareRateLimit(),
			"cloudflare_record":                 resourceCloudflareRecord(),
			"cloudflare_waf_rule":               resourceCloudflareWAFRule(),
			"cloudflare_zone_settings_override": resourceCloudflareZoneSettingsOverride(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	limitOpt := cloudflare.UsingRateLimit(float64(d.Get("rps").(int)))
	retryOpt := cloudflare.UsingRetryPolicy(d.Get("retries").(int), d.Get("min_backoff").(int), d.Get("max_backoff").(int))
	options := []cloudflare.Option{limitOpt, retryOpt}
	if d.Get("api_client_logging").(bool) {
		options = append(options, cloudflare.UsingLogger(log.New(os.Stderr, "", log.LstdFlags)))
	}

	config := Config{
		Email:   d.Get("email").(string),
		Token:   d.Get("token").(string),
		Options: options,
	}

	client, err := config.Client()
	if err != nil {
		return nil, err
	}

	if orgId, ok := d.GetOk("org_id"); ok {
		log.Printf("[INFO] Using specified organization id %s in Cloudflare provider", orgId.(string))
		options = append(options, cloudflare.UsingOrganization(orgId.(string)))
	} else if zoneName, ok := d.GetOk("use_org_from_zone"); ok {
		zoneId, err := client.ZoneIDByName(zoneName.(string))
		if err != nil {
			return nil, fmt.Errorf("error finding zone %q: %s", zoneName.(string), err)
		}

		zone, err := client.ZoneDetails(zoneId)
		if err != nil {
			return nil, err
		}
		log.Printf("[DEBUG] Looked up zone to match organization details to: %#v", zone)

		orgs, _, err := client.ListOrganizations()
		if err != nil {
			return nil, fmt.Errorf("error listing organizations: %s", err.Error())
		}
		log.Printf("[DEBUG] Found organizations for current user: %#v", orgs)

		orgIds := make([]string, len(orgs))
		for _, org := range orgs {
			orgIds = append(orgIds, org.ID)
		}

		if contains(orgIds, zone.Owner.ID) {
			log.Printf("[INFO] Using organization %#v in Cloudflare provider", zone.Owner)
			options = append(options, cloudflare.UsingOrganization(zone.Owner.ID))
		} else {
			log.Printf("[INFO] Zone ownership specified but organization owner not found. Falling back to using user API for Cloudflare provider")
		}
	} else {
		return client, err
	}

	config = Config{
		Email:   d.Get("email").(string),
		Token:   d.Get("token").(string),
		Options: options,
	}

	client, err = config.Client()
	if err != nil {
		return nil, err
	}

	return client, err
}
