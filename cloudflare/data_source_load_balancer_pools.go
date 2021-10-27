package cloudflare

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareLoadBalancerPools() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareLoadBalancerPoolsRead,

		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"pools": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"origins": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     originsElem,
						},

						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"minimum_origins": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"latitude": {
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"longitude": {
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"check_regions": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"monitor": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"notification_email": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"load_shedding": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     loadShedElem,
						},

						"created_on": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"modified_on": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCloudflareLoadBalancerPoolsRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading Load Balancer Pools")
	client := meta.(*cloudflare.API)

	filter, err := expandFilterLoadBalancerPools(d.Get("filter"))
	if err != nil {
		return err
	}

	poolsObj, err := client.ListLoadBalancerPools(context.Background())
	if err != nil {
		return fmt.Errorf("error listing load balancer pools: %w", err)
	}

	poolIds := make([]string, 0)
	pools := make([]map[string]interface{}, 0, len(poolsObj))
	for _, p := range poolsObj {
		if filter.Name != nil && !filter.Name.Match([]byte(p.Name)) {
			continue
		}

		pools = append(pools, map[string]interface{}{
			"id":                 p.ID,
			"name":               p.Name,
			"origins":            flattenLoadBalancerOrigins(d, p.Origins),
			"enabled":            p.Enabled,
			"minimum_origins":    p.MinimumOrigins,
			"latitude":           p.Latitude,
			"longitude":          p.Longitude,
			"check_regions":      p.CheckRegions,
			"description":        p.Description,
			"monitor":            p.Monitor,
			"notification_email": p.NotificationEmail,
			"load_shedding":      flattenLoadBalancerLoadShedding(p.LoadShedding),
			"created_on":         p.CreatedOn.Format(time.RFC3339Nano),
			"modified_on":        p.ModifiedOn.Format(time.RFC3339Nano),
		})
		poolIds = append(poolIds, p.ID)
	}

	if err := d.Set("pools", pools); err != nil {
		return fmt.Errorf("error setting load balancer pools: %w", err)
	}

	d.SetId(stringListChecksum(poolIds))
	return nil
}

type searchLoadBalancerPools struct {
	Name *regexp.Regexp
}

func expandFilterLoadBalancerPools(d interface{}) (*searchLoadBalancerPools, error) {
	cfg := d.([]interface{})
	filter := &searchLoadBalancerPools{}
	if len(cfg) == 0 || cfg[0] == nil {
		return filter, nil
	}

	m := cfg[0].(map[string]interface{})
	name, ok := m["name"]
	if ok {
		match, err := regexp.Compile(name.(string))
		if err != nil {
			return nil, err
		}
		filter.Name = match
	}

	return filter, nil
}
