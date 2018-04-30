package cloudflare

import (
	"fmt"
	"log"
	"strings"

	"time"

	"regexp"

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
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringMatch(regexp.MustCompile("[-_a-zA-Z0-9]+"), "Only alphanumeric characters, hyphens and underscores are allowed."),
			},

			"origins": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem:     originsElem,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: true,
			},

			"minimum_origins": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
				ForceNew: true,
			},

			"check_regions": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
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

var originsElem = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},

		"address": {
			Type:     schema.TypeString,
			Required: true,
			Elem: &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validateStringIP,
			},
		},

		"weight": {
			Type:         schema.TypeFloat,
			Optional:     true,
			Default:      1.0,
			ValidateFunc: floatBetween(0.0, 1.0),
		},

		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
	},
}

func resourceCloudFlareLoadBalancerPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerPool := cloudflare.LoadBalancerPool{
		Name:           d.Get("name").(string),
		Origins:        expandLoadBalancerOrigins(d.Get("origins").(*schema.Set)),
		Enabled:        d.Get("enabled").(bool),
		MinimumOrigins: d.Get("minimum_origins").(int),
	}

	if checkRegions, ok := d.GetOk("check_regions"); ok {
		loadBalancerPool.CheckRegions = expandInterfaceToStringList(checkRegions.(*schema.Set).List())
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

	log.Printf("[DEBUG] Creating CloudFlare Load Balancer Pool from struct: %+v", loadBalancerPool)

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
			Weight:  o["weight"].(float64),
		}
		origins = append(origins, origin)
	}
	return
}

func resourceCloudFlareLoadBalancerPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerPool, err := client.LoadBalancerPoolDetails(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Load balancer pool %s no longer exists", d.Id())
			d.SetId("")
			return nil
		} else {
			return errors.Wrap(err,
				fmt.Sprintf("Error reading load balancer pool from API for resource %s ", d.Id()))
		}
	}
	log.Printf("[DEBUG] Read CloudFlare Load Balancer Pool from API as struct: %+v", loadBalancerPool)

	d.Set("name", loadBalancerPool.Name)
	d.Set("enabled", loadBalancerPool.Enabled)
	d.Set("minimum_origins", loadBalancerPool.MinimumOrigins)
	d.Set("description", loadBalancerPool.Description)
	d.Set("monitor", loadBalancerPool.Monitor)
	d.Set("notification_email", loadBalancerPool.NotificationEmail)
	d.Set("created_on", loadBalancerPool.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancerPool.ModifiedOn.Format(time.RFC3339Nano))

	if err := d.Set("origins", flattenLoadBalancerOrigins(loadBalancerPool.Origins)); err != nil {
		log.Printf("[WARN] Error setting origins on load balancer pool %q: %s", d.Id(), err)
	}

	if err := d.Set("check_regions", schema.NewSet(schema.HashString, flattenStringList(loadBalancerPool.CheckRegions))); err != nil {
		log.Printf("[WARN] Error setting check_regions on load balancer pool %q: %s", d.Id(), err)
	}

	return nil
}

func flattenLoadBalancerOrigins(origins []cloudflare.LoadBalancerOrigin) *schema.Set {
	flattened := make([]interface{}, 0)
	for _, o := range origins {
		cfg := map[string]interface{}{
			"name":    o.Name,
			"address": o.Address,
			"enabled": o.Enabled,
			"weight":  o.Weight,
		}
		flattened = append(flattened, cfg)
	}
	return schema.NewSet(schema.HashResource(originsElem), flattened)
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

// floatBetween returns a validate function that can be used in schema definitions.
func floatBetween(min, max float64) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(float64)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be float64", k))
			return
		}

		if v < min || v > max {
			es = append(es, fmt.Errorf("expected %s to be within %v and %v", k, min, max))
			return
		}

		return
	}
}
