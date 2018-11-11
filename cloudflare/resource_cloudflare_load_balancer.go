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

func resourceCloudflareLoadBalancer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareLoadBalancerCreate,
		Read:   resourceCloudflareLoadBalancerRead,
		Update: resourceCloudflareLoadBalancerUpdate,
		Delete: resourceCloudflareLoadBalancerDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareLoadBalancerImport,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"zone": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
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
				Type:          schema.TypeBool,
				Optional:      true,
				Default:       false,
				ConflictsWith: []string{"ttl"},
			},

			"ttl": {
				Type:          schema.TypeInt,
				Optional:      true,
				Computed:      true,
				ConflictsWith: []string{"proxied"}, // this is set to zero regardless of config when proxied=true
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
			},

			"steering_policy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"off", "geo", "dynamic_latency", ""}, false),
				Default:      "",
			},

			// nb enterprise only
			"pop_pools": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     popPoolElem,
			},

			"region_pools": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem:     regionPoolElem,
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
	}
}

var popPoolElem = &schema.Resource{
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
}

var regionPoolElem = &schema.Resource{
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
}

var localPoolElems = map[string]*schema.Resource{
	"pop":    popPoolElem,
	"region": regionPoolElem,
}

func resourceCloudflareLoadBalancerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	newLoadBalancer := cloudflare.LoadBalancer{
		Name:           d.Get("name").(string),
		FallbackPool:   d.Get("fallback_pool_id").(string),
		DefaultPools:   expandInterfaceToStringList(d.Get("default_pool_ids")),
		Proxied:        d.Get("proxied").(bool),
		TTL:            d.Get("ttl").(int),
		SteeringPolicy: d.Get("steering_policy").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		newLoadBalancer.Description = description.(string)
	}

	if ttl, ok := d.GetOk("ttl"); ok {
		newLoadBalancer.TTL = ttl.(int)
	}

	if regionPools, ok := d.GetOk("region_pools"); ok {
		expandedRegionPools, err := expandGeoPools(regionPools, "region")
		if err != nil {
			return err
		}
		newLoadBalancer.RegionPools = expandedRegionPools
	}

	if popPools, ok := d.GetOk("pop_pools"); ok {
		expandedPopPools, err := expandGeoPools(popPools, "pop")
		if err != nil {
			return err
		}
		newLoadBalancer.PopPools = expandedPopPools
	}

	zoneName := d.Get("zone").(string)
	zoneId, err := client.ZoneIDByName(zoneName)
	if err != nil {
		return fmt.Errorf("error finding zone %q: %s", zoneName, err)
	}
	d.Set("zone_id", zoneId)

	log.Printf("[INFO] Creating Cloudflare Load Balancer from struct: %+v", newLoadBalancer)

	r, err := client.CreateLoadBalancer(zoneId, newLoadBalancer)
	if err != nil {
		return errors.Wrap(err, "error creating load balancer for zone")
	}

	if r.ID == "" {
		return fmt.Errorf("failed to find id in Create response; resource was empty")
	}

	d.SetId(r.ID)

	log.Printf("[INFO] Cloudflare Load Balancer ID: %s", d.Id())

	return resourceCloudflareLoadBalancerRead(d, meta)
}

func resourceCloudflareLoadBalancerUpdate(d *schema.ResourceData, meta interface{}) error {
	// since api only supports replace, update looks a lot like create...
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)

	loadBalancer := cloudflare.LoadBalancer{
		ID:             d.Id(),
		Name:           d.Get("name").(string),
		FallbackPool:   d.Get("fallback_pool_id").(string),
		DefaultPools:   expandInterfaceToStringList(d.Get("default_pool_ids")),
		Proxied:        d.Get("proxied").(bool),
		TTL:            d.Get("ttl").(int),
		SteeringPolicy: d.Get("steering_policy").(string),
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancer.Description = description.(string)
	}

	if regionPools, ok := d.GetOk("region_pools"); ok {
		expandedRegionPools, err := expandGeoPools(regionPools, "region")
		if err != nil {
			return err
		}
		loadBalancer.RegionPools = expandedRegionPools
	}

	if popPools, ok := d.GetOk("pop_pools"); ok {
		expandedPopPools, err := expandGeoPools(popPools, "pop")
		if err != nil {
			return err
		}
		loadBalancer.PopPools = expandedPopPools
	}

	log.Printf("[INFO] Updating Cloudflare Load Balancer from struct: %+v", loadBalancer)

	_, err := client.ModifyLoadBalancer(zoneId, loadBalancer)
	if err != nil {
		return errors.Wrap(err, "error creating load balancer for zone")
	}

	return resourceCloudflareLoadBalancerRead(d, meta)
}

func expandGeoPools(pool interface{}, geoType string) (map[string][]string, error) {
	cfg := pool.(*schema.Set).List()
	expanded := make(map[string][]string)
	for _, v := range cfg {
		locationConfig := v.(map[string]interface{})
		// lists are of type interface{} by default
		location := locationConfig[geoType].(string)
		if _, present := expanded[location]; !present {
			expanded[location] = expandInterfaceToStringList(locationConfig["pool_ids"])
		} else {
			return nil, fmt.Errorf("duplicate entry specified for %s pool in location %q. each location must only be specified once", geoType, location)
		}
	}
	return expanded, nil
}

func resourceCloudflareLoadBalancerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	loadBalancerId := d.Id()

	loadBalancer, err := client.LoadBalancerDetails(zoneId, loadBalancerId)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Load balancer %s in zone %s not found", loadBalancerId, zoneId)
			d.SetId("")
			return nil
		}
		return errors.Wrap(err,
			fmt.Sprintf("Error reading load balancer resource from API for resource %s in zone %s", zoneId, loadBalancerId))
	}

	d.Set("name", loadBalancer.Name)
	d.Set("fallback_pool_id", loadBalancer.FallbackPool)
	d.Set("proxied", loadBalancer.Proxied)
	d.Set("description", loadBalancer.Description)
	d.Set("ttl", loadBalancer.TTL)
	d.Set("steering_policy", loadBalancer.SteeringPolicy)
	d.Set("created_on", loadBalancer.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancer.ModifiedOn.Format(time.RFC3339Nano))

	if err := d.Set("default_pool_ids", loadBalancer.DefaultPools); err != nil {
		log.Printf("[WARN] Error setting default_pool_ids on load balancer %q: %s", d.Id(), err)
	}

	if err := d.Set("pop_pools", flattenGeoPools(loadBalancer.PopPools, "pop")); err != nil {
		log.Printf("[WARN] Error setting pop_pools on load balancer %q: %s", d.Id(), err)
	}

	if err := d.Set("region_pools", flattenGeoPools(loadBalancer.RegionPools, "region")); err != nil {
		log.Printf("[WARN] Error setting region_pools on load balancer %q: %s", d.Id(), err)
	}

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
	}
	return schema.NewSet(schema.HashResource(localPoolElems[geoType]), flattened)
}

func resourceCloudflareLoadBalancerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneId := d.Get("zone_id").(string)
	loadBalancerId := d.Id()

	log.Printf("[INFO] Deleting Cloudflare Load Balancer: %s in zone: %s", loadBalancerId, zoneId)

	err := client.DeleteLoadBalancer(zoneId, loadBalancerId)
	if err != nil {
		return fmt.Errorf("error deleting Cloudflare Load Balancer: %s", err)
	}

	return nil
}

func resourceCloudflareLoadBalancerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)

	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var zoneName string
	var loadBalancerId string
	if len(idAttr) == 2 {
		zoneName = idAttr[0]
		loadBalancerId = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneName/loadBalancerId\"", d.Id())
	}
	zoneId, err := client.ZoneIDByName(zoneName)

	if err != nil {
		return nil, fmt.Errorf("error finding zoneName %q: %s", zoneName, err)
	}

	d.Set("zone", zoneName)
	d.Set("zone_id", zoneId)
	d.SetId(loadBalancerId)
	return []*schema.ResourceData{d}, nil
}
