package cloudflare

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudflareUniversalSSL() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareUniversalSSLCreate,
		Read:   resourceCloudflareUniversalSSLRead,
		Update: resourceCloudflareUniversalSSLUpdate,
		Delete: resourceCloudflareUniversalSSLDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareUniversalSSLImport,
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"settings": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: resourceCloudflareUniversalSSLSchema,
				},
			},
			"initial_settings": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: resourceCloudflareUniversalSSLSchema,
				},
			},
		},
	}
}

var resourceCloudflareUniversalSSLSchema = map[string]*schema.Schema{
	"status": {
		Type:         schema.TypeString,
		ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
		Required:     true,
	},
}

func resourceCloudflareUniversalSSLCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	d.SetId(zoneID)

	ussl, err := client.UniversalSSLSettingDetails(d.Id())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading universal SSL status for zone %q", d.Id()))
	}

	if err := buildAndSetUniversalSSLStatus(d, ussl, "initial_settings"); err != nil {
		log.Printf("[WARN] Error setting initial_settings for zone %q: %s", d.Id(), err)
	}

	return resourceCloudflareUniversalSSLUpdate(d, meta)
}

func resourceCloudflareUniversalSSLImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	zoneID := d.Id()
	d.Set("zone_id", zoneID)

	ussl, err := client.UniversalSSLSettingDetails(zoneID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error reading universal SSL status for zone %q", d.Id()))
	}

	if err := buildAndSetUniversalSSLStatus(d, ussl, "initial_settings"); err != nil {
		log.Printf("[WARN] Error setting initial_settings for zone %q: %s", d.Id(), err)
	}

	resourceCloudflareUniversalSSLRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareUniversalSSLRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	d.SetId(zoneID)

	ussl, err := client.UniversalSSLSettingDetails(d.Id())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error reading universal SSL status for zone %q", d.Id()))
	}

	return buildAndSetUniversalSSLStatus(d, ussl, "settings")
}

func resourceCloudflareUniversalSSLUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	d.SetId(zoneID)

	_, err := client.EditUniversalSSLSetting(d.Id(), cloudflare.UniversalSSLSetting{Enabled: boolFromString(d.Get("settings.0.status").(string))})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot update universal SSL status for zone %q", d.Id()))
	}

	return nil
}

func resourceCloudflareUniversalSSLDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	d.SetId(zoneID)

	_, err := client.EditUniversalSSLSetting(d.Id(), cloudflare.UniversalSSLSetting{Enabled: boolFromString(d.Get("initial_settings.0.status").(string))})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Cannot update universal SSL status for zone %q", d.Id()))
	}
	return nil
}

func buildAndSetUniversalSSLStatus(d *schema.ResourceData, ussl cloudflare.UniversalSSLSetting, key string) error {
	settings := map[string]interface{}{}
	settings["status"] = stringFromBool(ussl.Enabled)

	return d.Set(key, []map[string]interface{}{settings})
}

func boolFromString(status string) bool {
	if status == "on" {
		return true
	}
	return false
}

func stringFromBool(status bool) string {
	if status {
		return "on"
	}
	return "off"
}
