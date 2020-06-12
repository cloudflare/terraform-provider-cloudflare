package cloudflare

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceCloudflareZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareZonesRead,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"paused": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"zones": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareZonesRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading Zones")
	client := meta.(*cloudflare.API)
	filter, err := expandFilter(d.Get("filter"))
	if err != nil {
		return err
	}

	zoneFilter := cloudflare.WithZoneFilters(
		fmt.Sprintf("contains:%s", filter.name),
		"",
		filter.status,
	)

	zones, err := client.ListZonesContext(context.TODO(), zoneFilter)
	if err != nil {
		return fmt.Errorf("error listing Zone: %s", err)
	}

	zoneDetails := make([]interface{}, 0)
	for _, v := range zones.Result {
		if filter.paused != v.Paused {
			continue
		}

		zoneDetails = append(zoneDetails, map[string]interface{}{
			"id":   v.ID,
			"name": v.Name,
		})
	}

	err = d.Set("zones", zoneDetails)
	if err != nil {
		return fmt.Errorf("Error setting zones: %s", err)
	}

	d.SetId(time.Now().UTC().String())
	return nil
}

func expandFilter(d interface{}) (*searchFilter, error) {
	cfg := d.([]interface{})
	filter := &searchFilter{}

	m := cfg[0].(map[string]interface{})
	name, ok := m["name"]
	if ok {
		filter.name = name.(string)
	}

	paused, ok := m["paused"]
	if ok {
		filter.paused = paused.(bool)
	}

	status, ok := m["status"]
	if ok {
		filter.status = status.(string)
	}

	return filter, nil
}

type searchFilter struct {
	name   string
	status string
	paused bool
}
