package cloudflare

import (
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
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

		if filter.name != nil {
			if !filter.name.Match([]byte(v.Name)) {
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
	name, ok := m["name"]
	if ok {
		match, err := regexp.Compile(name.(string))
		if err != nil {
			return nil, err
		}

		filter.name = match
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
	name   *regexp.Regexp
	status string
	paused bool
}
