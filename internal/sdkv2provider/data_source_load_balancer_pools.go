package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareLoadBalancerPools() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceCloudflareLoadBalancerPoolsRead,
		Description: "A datasource to find Load Balancer Pools.",
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the datasource lookups.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"filter": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "One or more values used to look up Load Balancer pools. If more than one value is given all values must match in order to be included.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A regular expression matching the name of the Load Balancer pool to lookup.",
						},
					},
				},
			},
			"pools": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "A list of Load Balancer Pools details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID for this load balancer pool.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Short name (tag) for the pool.",
						},

						"origins": {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        originsElem,
							Description: "The list of origins within this pool.",
						},

						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this pool is enabled. Disabled pools will not receive traffic and are excluded from health checks.",
						},

						"minimum_origins": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum number of origins that must be healthy for this pool to serve traffic.",
						},

						"latitude": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Latitude this pool is physically located at; used for proximity steering.",
						},

						"longitude": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Longitude this pool is physically located at; used for proximity steering.",
						},

						"check_regions": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "List of regions (specified by region code) from which to run health checks. Empty means every Cloudflare data center (the default), but requires an Enterprise plan. Region codes can be found [here](https://support.cloudflare.com/hc/en-us/articles/115000540888-Load-Balancing-Geographic-Regions).",
						},

						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Brief description of the Load Balancer Pool intention.",
						},

						"monitor": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the Monitor to use for health checking origins within this pool.",
						},

						"notification_email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Email address to send health status notifications to. Multiple emails are set as a comma delimited list.",
						},

						"load_shedding": {
							Type:        schema.TypeSet,
							Computed:    true,
							Elem:        loadShedElem,
							Description: "Setting for controlling load shedding for this pool.",
						},

						"created_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The RFC3339 timestamp of when the load balancer was created.",
						},

						"modified_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The RFC3339 timestamp of when the load balancer was last modified.",
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

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	poolsObj, err := client.ListLoadBalancerPools(context.Background(), cloudflare.AccountIdentifier(accountID), cloudflare.ListLoadBalancerPoolParams{})
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
