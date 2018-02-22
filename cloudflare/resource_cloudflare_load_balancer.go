package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/pkg/errors"
)

func resourceCloudFlareLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlareLoadBalancerCreate,
		Read:   resourceCloudFlareLoadBalancerRead,
		Update: resourceCloudFlareLoadBalancerUpdate,
		Delete: resourceCloudFlareLoadBalancerDelete,
		Exists: resourceCloudFlareLoadBalancerExists,
		Importer: &schema.ResourceImporter{
			State: resourceCloudFlareLoadBalancerImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"load_balancer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": { // "dns_name" ??
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false, // allow reusing same load balancer for different dns names
				// todo: validate dns name
			},

			"fallback_pool_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 32),
			},

			"default_pool_ids": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringLenBetween(1, 32),
				},
			},

			"proxied": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"ttl": {
				Type:     schema.TypeInt,
				Optional: true,
				// TODO is this computed / does it have a default?
			},

			// see https://github.com/hashicorp/terraform/issues/6215
			// nb enterprise only
			"pop_pools": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pop": {
							Type:     schema.TypeString,
							Required: true,
							// let the api handle validating pops
						},

						"pool_ids": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringLenBetween(1, 32),
							},
						},
					},
				},
				Set: HashByMapKey("pop"),
			},

			"region_pools": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:     schema.TypeString,
							Required: true,
							// let the api handle validating regions
						},

						"pool_ids": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type:         schema.TypeString,
								ValidateFunc: validation.StringLenBetween(1, 32),
							},
						},
					},
				},
				Set: HashByMapKey("region"),
			},

			"created_on": {
				Type:     schema.TypeBool,
				Computed: true,
			},

			"modified_on": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func resourceCloudFlareLoadBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newLoadBalancer := cloudflare.LoadBalancer{
		Name:         d.Get("name").(string),
		FallbackPool: d.Get("fallback_pool_id").(string),
		DefaultPools: expandInterfaceToStringList(d.Get("default_pool_ids")),
		Proxied:      d.Get("proxied").(bool),
	}

	if description, ok := d.GetOk("description"); ok {
		newLoadBalancer.Description = description.(string)
	}

	// TODO in default case, this is sending ttl=0, is this ok??
	if ttl, ok := d.GetOk("ttl"); ok {
		newLoadBalancer.TTL = ttl.(int)
	}

	if regionPools, ok := d.GetOk("region_pools"); ok {
		newLoadBalancer.RegionPools = expandGeoPools(regionPools, "region")
	}

	if popPools, ok := d.GetOk("pop_pools"); ok {
		newLoadBalancer.PopPools = expandGeoPools(popPools, "pop")
	}

	zoneName := d.Get("zone").(string)
	zoneId, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return fmt.Errorf("error finding zone %q: %s", zoneName, err)
	}

	log.Printf("[INFO] Creating CloudFlare Load Balancer from struct: %+v", newLoadBalancer)

	r, err := client.CreateLoadBalancer(zoneId, newLoadBalancer)
	if err != nil {
		return errors.Wrap(err, "error creating load balancer for zone")
	}

	if r.ID == "" {
		return fmt.Errorf("cailed to find id in Create response; resource was empty")
	}

	// terraform id is *not* the same as the resource id, is is the combination with the zoneId
	// this makes it easier to import and also matches the keys needed for cloudflare-go operations
	d.SetId(zoneName + "_" + r.ID)
	d.Set("load_balancer_id", r.ID)

	log.Printf("[INFO] CloudFlare Load Balancer ID: %s", d.Id())

	return resourceCloudFlareLoadBalancerRead(d, meta)
}

func resourceCloudFlareLoadBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	// since api only supports replace, update looks a lot like create...
	client := meta.(*cloudflare.API)

	newLoadBalancer := cloudflare.LoadBalancer{
		Name:         d.Get("name").(string),
		FallbackPool: d.Get("fallback_pool_id").(string),
		DefaultPools: expandInterfaceToStringList(d.Get("default_pool_ids")),
		Proxied:      d.Get("proxied").(bool),
	}

	if description, ok := d.GetOk("description"); ok {
		newLoadBalancer.Description = description.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		newLoadBalancer.TTL = ttl.(int)
	}

	if regionPools, ok := d.GetOk("region_pools"); ok {
		newLoadBalancer.RegionPools = expandGeoPools(regionPools, "region")
	}

	if popPools, ok := d.GetOk("pop_pools"); ok {
		newLoadBalancer.PopPools = expandGeoPools(popPools, "pop")
	}

	zoneName := d.Get("zone").(string)
	zoneId, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return fmt.Errorf("error finding zone %q: %s", zoneName, err)
	}

	log.Printf("[INFO] Creating CloudFlare Load Balancer from struct: %+v", newLoadBalancer)

	_, err = client.ModifyLoadBalancer(zoneId, newLoadBalancer)
	if err != nil {
		return errors.Wrap(err, "error creating load balancer for zone")
	}

	return resourceCloudFlareLoadBalancerRead(d, meta)
}

func expandGeoPools(pool interface{}, geoType string) map[string][]string {
	cfg := pool.(*schema.Set).List()
	expanded := make(map[string][]string)
	for _, v := range cfg {
		locationConfig := v.(map[string]interface{})
		// assume for now that lists are of type interface{} by default
		expanded[locationConfig[geoType].(string)] = expandInterfaceToStringList(locationConfig["pool_ids"])
	}
	return expanded
}

func resourceCloudFlareLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)

	loadBalancer, err := client.LoadBalancerDetails(zoneId, loadBalancerId)
	if err != nil {
		return errors.Wrap(err,
			fmt.Sprintf("Error reading load balancer resource from API for resource %s in zone %s", zoneId, loadBalancerId))
	}
	log.Printf("[INFO] Read CloudFlare Load Balancer from API as struct: %+v", loadBalancer)

	// api generally tries to populate everything, so just assume all data is present
	// start by setting required values
	d.Set("name", loadBalancer.Name)
	d.Set("fallback_pool_id", loadBalancer.FallbackPool)
	d.Set("default_pool_ids", loadBalancer.DefaultPools) // ok to pass []string to []interface{}

	d.Set("proxied", loadBalancer.Proxied)
	d.Set("description", loadBalancer.Description)
	d.Set("ttl", loadBalancer.TTL) // ok to pass []string to []interface{}

	d.Set("pop_pools", flattenGeoPools(loadBalancer.PopPools, "pop"))
	d.Set("region_pools", flattenGeoPools(loadBalancer.RegionPools, "region"))

	d.Set("created_on", loadBalancer.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancer.ModifiedOn.Format(time.RFC3339Nano))

	return nil
}

func flattenGeoPools(pools map[string][]string, geoType string) *schema.Set {
	flattened := make([]interface{}, 0)
	for k, v := range pools {
		geoConf := map[string]interface{}{
			geoType:    k,
			"pool_ids": flattenStringList(v),
		}
		flattened = append(flattened, geoConf)
		// assume for now that lists are of type interface{} by default
	}
	return schema.NewSet(HashByMapKey(geoType), flattened)
}

func resourceCloudFlareLoadBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)

	log.Printf("[INFO] Deleting CloudFlare Load Balancer: %s in zone: %s", loadBalancerId, zoneId)

	err := client.DeleteLoadBalancer(zoneId, loadBalancerId)
	if err != nil {
		return fmt.Errorf("error deleting CloudFlare Load Balancer: %s", err)
	}

	return nil
}

func resourceCloudFlareLoadBalancerExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	loadBalancerId := d.Get("load_balancer_id").(string)

	_, err := client.LoadBalancerDetails(zoneId, loadBalancerId)
	if err != nil {
		log.Printf("[INFO] Error found when checking if  load balancer exists: %s", err.Error())
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Found status 404 looking for resource %s in zone %s", loadBalancerId, zoneId)
			return false, nil
		} else {
			return false, errors.Wrap(err,
				fmt.Sprintf("Error reading load balancer resource from API for resource %s in zone %s", loadBalancerId, zoneId))
		}
	}

	return true, nil
}

func resourceCloudFlareLoadBalancerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "_", 2)
	var zoneName string
	var loadBalancerId string
	if len(idAttr) == 2 {
		zoneName = idAttr[0]
		loadBalancerId = idAttr[1]
		d.Set("zone", zoneName)
		d.Set("load_balancer_id", loadBalancerId)
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneName_loadBalancerId\"", d.Id())
	}
	zoneId, err := client.ZoneIDByName(zoneName)
	d.Set("zone_id", zoneId)
	if err != nil {
		return nil, fmt.Errorf("error finding zoneName %q: %s", zoneName, err)
	}
	return []*schema.ResourceData{d}, nil
}

func HashByMapKey(key string) func(v interface{}) int {
	return func(v interface{}) int {
		m := v.(map[string]interface{})
		return schema.HashString(m[key])
	}
}
