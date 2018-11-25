package cloudflare

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"regexp"
	"time"
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
						"zone": {
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
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
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

	zones, err := client.ListZones()
	if err != nil {
		return fmt.Errorf("error listing Zone: %s", err)
	}

	var zoneNames []string
	for _, v := range zones {

		if filter.zone != nil {
			if !filter.zone.Match([]byte(v.Name)) {
				continue
			}
		}

		if filter.paused != v.Paused {
			continue
		}

		if filter.status != "" && filter.status != v.Status {
			continue
		}

		zoneNames = append(zoneNames, v.Name)
	}

	err = d.Set("zones", zoneNames)
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
	zone, ok := m["zone"]
	if ok {
		match, err := regexp.Compile(zone.(string))
		if err != nil {
			return nil, err
		}

		filter.zone = match
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
	zone   *regexp.Regexp
	status string
	paused bool
}
