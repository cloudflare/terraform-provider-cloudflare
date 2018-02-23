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

func resourceCloudFlareLoadBalancerPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudFlareLoadBalancerPoolCreate,
		Read:   resourceCloudFlareLoadBalancerPoolRead,
		Delete: resourceCloudFlareLoadBalancerPoolDelete,
		Exists: resourceCloudFlareLoadBalancerPoolExists,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// TODO check that order of origins is not significant
			"origins": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"address": {
							Type:     schema.TypeString,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								// TODO: validate IP address
							},
						},

						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
				Set: HashByMapKey("address"),
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},

			"minimum_origins": { // TODO not currently used as not in the API client
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: true,
			},

			"check_regions": { //TODO: set??
				Type:     schema.TypeList,
				Optional: true,
				Computed: true, // TODO this is messy for people on a restricted plan
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ForceNew: true,
			},

			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				ForceNew:     true,
			},

			"monitor": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(0, 32),
				ForceNew:     true,
			},

			"notification_email": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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

func resourceCloudFlareLoadBalancerPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerPool := cloudflare.LoadBalancerPool{
		Name:         d.Get("name").(string),
		Origins:      expandLoadBalancerOrigins(d.Get("origins").(*schema.Set)),
		Enabled:      d.Get("enabled").(bool),
		CheckRegions: expandInterfaceToStringList(d.Get("check_regions")),
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerPool.Description = description.(string)
	}

	if monitor, ok := d.GetOk("monitor"); ok {
		loadBalancerPool.Monitor = monitor.(string)
	}

	if notificationEmail, ok := d.GetOk("notification_email"); ok {
		loadBalancerPool.NotificationEmail = notificationEmail.(string)
	}

	log.Printf("[INFO] Creating CloudFlare Load Balancer Pool from struct: %+v", loadBalancerPool)

	r, err := client.CreateLoadBalancerPool(loadBalancerPool)
	if err != nil {
		return errors.Wrap(err, "error creating load balancer pool")
	}

	if r.ID == "" {
		return fmt.Errorf("cailed to find id in create response; resource was empty")
	}

	d.SetId(r.ID)

	log.Printf("[INFO] New CloudFlare Load Balancer Pool created with  ID: %s", d.Id())

	return resourceCloudFlareLoadBalancerPoolRead(d, meta)
}

func expandLoadBalancerOrigins(originSet *schema.Set) (origins []cloudflare.LoadBalancerOrigin) {
	for _, iface := range originSet.List() {
		o := iface.(map[string]interface{})
		origin := cloudflare.LoadBalancerOrigin{
			Name:    o["name"].(string),
			Address: o["address"].(string),
			Enabled: o["enabled"].(bool),
		}
		origins = append(origins, origin)
	}
	return
}

func resourceCloudFlareLoadBalancerPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerPool, err := client.LoadBalancerPoolDetails(d.Id())
	if err != nil {
		return errors.Wrap(err,
			fmt.Sprintf("Error reading load balancer resource from API for resource %s ", d.Id()))
	}
	log.Printf("[INFO] Read CloudFlare Load Balancer Pool from API as struct: %+v", loadBalancerPool)

	// api generally tries to populate everything, so just assume all data is present
	// start by setting required values
	d.Set("name", loadBalancerPool.Name)
	d.Set("origins", flattenLoadBalancerOrigins(loadBalancerPool.Origins))
	d.Set("enabled", loadBalancerPool.Enabled)

	d.Set("minimum_origins", 1) // TODO
	d.Set("check_regions", flattenStringList(loadBalancerPool.CheckRegions))

	// ok to set empty optional/ noncomputed values?
	d.Set("description", loadBalancerPool.Description)
	d.Set("monitor", loadBalancerPool.Monitor)
	d.Set("notification_email", loadBalancerPool.NotificationEmail)

	d.Set("created_on", loadBalancerPool.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancerPool.ModifiedOn.Format(time.RFC3339Nano))

	return nil
}

func flattenLoadBalancerOrigins(origins []cloudflare.LoadBalancerOrigin) *schema.Set {
	flattened := make([]interface{}, 0)
	for _, o := range origins {
		cfg := map[string]interface{}{
			"name":    o.Name,
			"address": o.Address,
			"enabled": o.Enabled,
		}
		flattened = append(flattened, cfg)
	}
	return schema.NewSet(HashByMapKey("address"), flattened)
}

func resourceCloudFlareLoadBalancerPoolDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting CloudFlare Load Balancer Pool: %s ", d.Id())

	err := client.DeleteLoadBalancerPool(d.Id())
	if err != nil {
		return errors.Wrap(err, "error deleting CloudFlare Load Balancer Pool")
	}

	return nil
}

func resourceCloudFlareLoadBalancerPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*cloudflare.API)

	_, err := client.LoadBalancerPoolDetails(d.Id())
	if err != nil {
		log.Printf("[INFO] Error found when checking if load balancer pool exists: %s", err.Error())
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Found status 404 looking for resource %s", d.Id())
			return false, nil
		} else {
			return false, errors.Wrap(err,
				fmt.Sprintf("Error reading load balancer resource from API for resource %s", d.Id()))
		}
	}

	return true, nil
}
