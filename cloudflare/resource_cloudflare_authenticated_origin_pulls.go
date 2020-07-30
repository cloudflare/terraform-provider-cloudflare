package cloudflare

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareAuthenticatedOriginPulls() *schema.Resource {
	return &schema.Resource{
		// AOP is a toggleable feature, editing is the same as creating.
		Create: resourceCloudflareAuthenticatedOriginPullsCreate,
		Read:   resourceCloudflareAuthenticatedOriginPullsRead,
		Update: resourceCloudflareAuthenticatedOriginPullsCreate,
		Delete: resourceCloudflareAuthenticatedOriginPullsDelete,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"authenticated_origin_pulls_certificate": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}
}

func resourceCloudflareAuthenticatedOriginPullsCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	hostname := d.Get("hostname").(string)
	aopCert := d.Get("authenticated_origin_pulls_certificate").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)
	var checksum string
	switch isEnabled, ok := d.GetOk("enabled"); ok {
	case hostname != "" && aopCert != "":
		// Per Hostname AOP
		conf := []cloudflare.PerHostnameAuthenticatedOriginPullsConfig{{
			CertID:   aopCert,
			Hostname: hostname,
			Enabled:  isEnabled.(bool),
		}}
		_, err := client.EditPerHostnameAuthenticatedOriginPullsConfig(zoneID, conf)
		if err != nil {
			return fmt.Errorf("error creating Per-Hostname Authenticated Origin Pulls resource on zone %q: %s", zoneID, err)
		}
		checksum = stringChecksum(fmt.Sprintf("PerHostnameAOP/%s/%s/%s", zoneID, hostname, aopCert))
		
	case aopCert != "":
		// Per Zone AOP
		_, err := client.SetPerZoneAuthenticatedOriginPullsStatus(zoneID, isEnabled.(bool))
		if err != nil {
			return fmt.Errorf("error creating Per-Zone Authenticated Origin Pulls resource on zone %q: %s", zoneID, err)
		}
		checksum = stringChecksum(fmt.Sprintf("PerZoneAOP/%s/%s", zoneID, aopCert))

	default:
		// Global AOP
		_, err := client.SetAuthenticatedOriginPullsStatus(zoneID, isEnabled.(bool))
		if err != nil {
			return fmt.Errorf("error creating Global Authenticated Origin Pulls resource on zone %q: %s", zoneID, err)
		}
		checksum = stringChecksum(fmt.Sprintf("GlobalAOP/%s/", zoneID))
	}

  d.SetId(resourceChecksum)
	return resourceCloudflareAuthenticatedOriginPullsRead(d, meta)
}

func resourceCloudflareAuthenticatedOriginPullsRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	hostname := d.Get("hostname").(string)
	aopCert := d.Get("authenticated_origin_pulls_certificate").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	if hostname != "" && aopCert != "" {
		// Per Hostname AOP
		res, err := client.GetPerHostnameAuthenticatedOriginPullsConfig(zoneID, hostname)
		if err != nil {
			return errors.Wrap(err, "failed to get Per-Hostname Authenticated Origin Pulls setting")
		}
		d.Set("enabled", res.Enabled)
	} else if aopCert != "" {
		// Per Zone AOP
		res, err := client.GetPerZoneAuthenticatedOriginPullsStatus(zoneID)
		if err != nil {
			return errors.Wrap(err, "failed to get Per-Zone Authenticated Origin Pulls setting")
		}
		d.Set("enabled", res.Enabled)
	} else {
		// Global AOP
		res, err := client.GetAuthenticatedOriginPullsStatus(zoneID)
		if err != nil {
			return errors.Wrap(err, "failed to get Global Authenticated Origin Pulls setting")
		}
		if res.Value == "on" {
			d.Set("enabled", true)
		} else {
			d.Set("enabled", false)
		}
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	hostname := d.Get("hostname").(string)
	aopCert := d.Get("authenticated_origin_pulls_certificate").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	if hostname != "" && aopCert != "" {
		// Per Hostname AOP
		conf := []cloudflare.PerHostnameAuthenticatedOriginPullsConfig{{
			CertID:   aopCert,
			Hostname: hostname,
			Enabled:  false,
		}}
		_, err := client.EditPerHostnameAuthenticatedOriginPullsConfig(zoneID, conf)
		if err != nil {
			return fmt.Errorf("Error disabling Per-Hostname Authenticated Origin Pulls resource on zone %q: %s", zoneID, err)
		}
	} else if aopCert != "" {
		// Per Zone AOP
		_, err := client.SetPerZoneAuthenticatedOriginPullsStatus(zoneID, false)
		if err != nil {
			return fmt.Errorf("Error disabling Per-Zone Authenticated Origin Pulls resource on zone %q: %s", zoneID, err)
		}
	} else {
		// Global AOP
		_, err := client.SetAuthenticatedOriginPullsStatus(zoneID, false)
		if err != nil {
			return fmt.Errorf("Error disabling Global Authenticated Origin Pulls resource on zone %q: %s", zoneID, err)
		}
	}
	return nil
}
