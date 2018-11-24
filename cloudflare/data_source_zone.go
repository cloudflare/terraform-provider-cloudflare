package cloudflare

import (
	"fmt"
	"log"
	"regexp"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func dataSourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareZoneRead,

		Schema: map[string]*schema.Schema{
			"zone": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("^([a-zA-Z0-9][\\-a-zA-Z0-9]*\\.)+[\\-a-zA-Z0-9]{2,20}$"), ""),
			},
			"paused": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"vanity_name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"meta": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"wildcard_proxiable": {
							Type: schema.TypeBool,
						},
						"phishing_detected": {
							Type: schema.TypeBool,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceCloudflareZoneRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading Zone")
	client := meta.(*cloudflare.API)
	name := d.Get("zone").(string)
	zones, err := client.ListZones(name)
	if err != nil {
		return fmt.Errorf("error listing Zone: %s", err)
	}

	if len(zones) > 1 {
		return fmt.Errorf("multiple zones for name %q found", name)
	}

	zone := zones[0]
	d.SetId(zone.ID)
	d.Set("zone", zone.Name)
	d.Set("paused", zone.Paused)
	d.Set("vanity_name_servers", zone.VanityNS)
	d.Set("status", zone.Status)
	d.Set("type", zone.Type)
	d.Set("name_servers", zone.NameServers)
	d.Set("meta", flattenMeta(d, zone.Meta))

	return nil
}
