package cloudflare

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareZoneRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"name"},
			},
			"account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"zone_id"},
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

	zoneID, zoneIDExists := d.GetOk("zone_id")
	zoneID = zoneID.(string)

	name, nameExists := d.GetOk("name")
	name = name.(string)

	accountID := d.Get("account_id")
	accountID = accountID.(string)

	if nameExists && zoneIDExists {
		return fmt.Errorf("zone_id and name arguments can't be used together")
	}

	if !nameExists && !zoneIDExists {
		return fmt.Errorf("either zone_id or name must be set")
	}

	var zone cloudflare.Zone
	if nameExists && !zoneIDExists { // if only the name was provided
		zoneFilter := cloudflare.WithZoneFilters(name.(string), accountID.(string), "")
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
	d.Set("name_servers", zone.NameServers)
	d.Set("vanity_name_servers", zone.VanityNS)
	return nil
}
