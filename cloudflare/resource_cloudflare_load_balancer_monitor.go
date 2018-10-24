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

func resourceCloudflareLoadBalancerMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceCloudflareLoadBalancerPoolMonitorCreate,
		Read:   resourceCloudflareLoadBalancerPoolMonitorRead,
		Update: resourceCloudflareLoadBalancerPoolMonitorUpdate,
		Delete: resourceCloudflareLoadBalancerPoolMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		SchemaVersion: 0,
		Schema: map[string]*schema.Schema{
			"expected_body": {
				Type:     schema.TypeString,
				Required: true,
			},

			"expected_codes": {
				Type:     schema.TypeString,
				Required: true,
			},

			"method": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GET",
			},

			"timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      5,
				ValidateFunc: validation.IntBetween(1, 10),
			},

			"path": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/",
			},

			"interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      60,
				ValidateFunc: validation.IntBetween(60, 3600),
			},
			// interval has to be larger than (retries+1) * probe_timeout:

			"retries": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      2,
				ValidateFunc: validation.IntBetween(1, 5),
			},

			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "http",
				ValidateFunc: validation.StringInSlice([]string{"http", "https"}, false),
			},

			"header": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": {
							Type:     schema.TypeString,
							Required: true,
						},

						"values": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
				Set: HashByMapKey("header"),
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
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

func resourceCloudflareLoadBalancerPoolMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerMonitor := cloudflare.LoadBalancerMonitor{
		ExpectedBody:  d.Get("expected_body").(string),
		ExpectedCodes: d.Get("expected_codes").(string),
		Method:        d.Get("method").(string),
		Timeout:       d.Get("timeout").(int),
		Type:          d.Get("type").(string),
		Path:          d.Get("path").(string),
		Interval:      d.Get("interval").(int),
		Retries:       d.Get("retries").(int),
	}

	if header, ok := d.GetOk("header"); ok {
		loadBalancerMonitor.Header = expandLoadBalancerMonitorHeader(header)
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerMonitor.Description = description.(string)
	}

	log.Printf("[DEBUG] Creating Cloudflare Load Balancer Monitor from struct: %+v", loadBalancerMonitor)

	r, err := client.CreateLoadBalancerMonitor(loadBalancerMonitor)
	if err != nil {
		return errors.Wrap(err, "error creating load balancer monitor")
	}

	if r.ID == "" {
		return fmt.Errorf("cailed to find id in create response; resource was empty")
	}

	d.SetId(r.ID)

	log.Printf("[INFO] New Cloudflare Load Balancer Monitor created with  ID: %s", d.Id())

	return resourceCloudflareLoadBalancerPoolMonitorRead(d, meta)
}

func resourceCloudflareLoadBalancerPoolMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerMonitor := cloudflare.LoadBalancerMonitor{
		ID:            d.Id(),
		ExpectedBody:  d.Get("expected_body").(string),
		ExpectedCodes: d.Get("expected_codes").(string),
		Method:        d.Get("method").(string),
		Timeout:       d.Get("timeout").(int),
		Type:          d.Get("type").(string),
		Path:          d.Get("path").(string),
		Interval:      d.Get("interval").(int),
		Retries:       d.Get("retries").(int),
	}

	if header, ok := d.GetOk("header"); ok {
		loadBalancerMonitor.Header = expandLoadBalancerMonitorHeader(header)
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerMonitor.Description = description.(string)
	}

	log.Printf("[DEBUG] Update Cloudflare Load Balancer Monitor from struct: %+v", loadBalancerMonitor)

	_, err := client.ModifyLoadBalancerMonitor(loadBalancerMonitor)
	if err != nil {
		return errors.Wrap(err, "error modifying load balancer monitor")
	}

	log.Printf("[INFO] Cloudflare Load Balancer Monitor %q was modified", d.Id())

	return resourceCloudflareLoadBalancerPoolMonitorRead(d, meta)
}

func expandLoadBalancerMonitorHeader(cfgSet interface{}) map[string][]string {
	header := make(map[string][]string)
	cfgList := cfgSet.(*schema.Set).List()
	for _, item := range cfgList {
		cfg := item.(map[string]interface{})
		header[cfg["header"].(string)] = expandInterfaceToStringList(cfg["values"].(*schema.Set).List())
	}
	return header
}

func resourceCloudflareLoadBalancerPoolMonitorRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	loadBalancerMonitor, err := client.LoadBalancerMonitorDetails(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Load balancer monitor %s no longer exists", d.Id())
			d.SetId("")
			return nil
		} else {
			return errors.Wrap(err,
				fmt.Sprintf("Error reading load balancer monitor from API for resource %s ", d.Id()))
		}
	}
	log.Printf("[DEBUG] Read Cloudflare Load Balancer Monitor from API as struct: %+v", loadBalancerMonitor)

	d.Set("expected_body", loadBalancerMonitor.ExpectedBody)
	d.Set("expected_codes", loadBalancerMonitor.ExpectedCodes)
	d.Set("method", loadBalancerMonitor.Method)
	d.Set("timeout", loadBalancerMonitor.Timeout)
	d.Set("path", loadBalancerMonitor.Path)
	d.Set("interval", loadBalancerMonitor.Interval)
	d.Set("retries", loadBalancerMonitor.Retries)
	d.Set("type", loadBalancerMonitor.Type)
	d.Set("description", loadBalancerMonitor.Description)
	d.Set("created_on", loadBalancerMonitor.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancerMonitor.ModifiedOn.Format(time.RFC3339Nano))

	if err := d.Set("header", flattenLoadBalancerMonitorHeader(loadBalancerMonitor.Header)); err != nil {
		log.Printf("[WARN] Error setting header for load balancer monitor %q: %s", d.Id(), err)
	}

	return nil
}

func flattenLoadBalancerMonitorHeader(header map[string][]string) *schema.Set {
	flattened := make([]interface{}, 0)
	for k, v := range header {
		cfg := map[string]interface{}{
			"header": k,
			"values": schema.NewSet(schema.HashString, flattenStringList(v)),
		}
		flattened = append(flattened, cfg)
	}
	return schema.NewSet(HashByMapKey("header"), flattened)
}

func resourceCloudflareLoadBalancerPoolMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting Cloudflare Load Balancer Monitor: %s ", d.Id())

	err := client.DeleteLoadBalancerMonitor(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Load balancer monitor %s no longer exists", d.Id())
			return nil
		} else {
			return errors.Wrap(err, "error deleting cloudflare load balancer monitor")
		}
	}

	return nil
}
