package cloudflare

import (
	"context"
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareZoneRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"zone_id", "name"},
			},
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"zone_id", "name"},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"paused": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"plan": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_servers": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"vanity_name_servers": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceCloudflareZoneRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading Zones")
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	name := d.Get("name").(string)
	accountID := d.Get("account_id").(string)

	var zone cloudflare.Zone
	if name != "" && zoneID == "" {
		zoneFilter := cloudflare.WithZoneFilters(name, accountID, "")
		zonesResp, err := client.ListZonesContext(context.Background(), zoneFilter)

		if err != nil {
			return fmt.Errorf("error listing zones: %s", err)
		}

		if zonesResp.Total > 1 {
			return fmt.Errorf("more than one zone was returned; consider using `cloudflare_zones` data source with filtering to target the zone more specifically")
		}

		if zonesResp.Total == 0 {
			return fmt.Errorf("no zone found")
		}

		zone = zonesResp.Result[0]
	} else {
		var err error
		zone, err = client.ZoneDetails(context.Background(), zoneID.(string))
		if err != nil {
			return fmt.Errorf("error getting zone details: %s", err)
		}
	}

	d.SetId(zone.ID)
	d.Set("zone_id", zone.ID)
	d.Set("account_id", zone.Account.ID)
	d.Set("name", zone.Name)
	d.Set("status", zone.Status)
	d.Set("paused", zone.Paused)
	d.Set("plan", zone.Plan.Name)

	if err := d.Set("name_servers", zone.NameServers); err != nil {
		return fmt.Errorf("failed to set name_servers attribute: %s", err)
	}

	if err := d.Set("vanity_name_servers", zone.VanityNS); err != nil {
		return fmt.Errorf("failed to set vanity_name_servers attribute: %s", err)
	}

	return nil
}
