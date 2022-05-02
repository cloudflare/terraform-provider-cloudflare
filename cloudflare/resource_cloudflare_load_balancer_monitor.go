package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareLoadBalancerMonitor() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareLoadBalancerMonitorSchema(),
		CreateContext: resourceCloudflareLoadBalancerPoolMonitorCreate,
		ReadContext: resourceCloudflareLoadBalancerPoolMonitorRead,
		UpdateContext: resourceCloudflareLoadBalancerPoolMonitorUpdate,
		DeleteContext: resourceCloudflareLoadBalancerPoolMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCloudflareLoadBalancerPoolMonitorCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	loadBalancerMonitor := cloudflare.LoadBalancerMonitor{
		Timeout:  d.Get("timeout").(int),
		Type:     d.Get("type").(string),
		Interval: d.Get("interval").(int),
		Retries:  d.Get("retries").(int),
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerMonitor.Description = description.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		loadBalancerMonitor.Port = uint16(port.(int))
	}

	switch loadBalancerMonitor.Type {
	case "tcp":
		if method, ok := d.GetOk("method"); ok {
			loadBalancerMonitor.Method = method.(string)
		} else {
			loadBalancerMonitor.Method = "connection_established"
		}
	case "http", "https":
		if allowInsecure, ok := d.GetOk("allow_insecure"); ok {
			loadBalancerMonitor.AllowInsecure = allowInsecure.(bool)
		}

		expectedBody := d.Get("expected_body")
		loadBalancerMonitor.ExpectedBody = expectedBody.(string)

		if expectedCodes, ok := d.GetOk("expected_codes"); ok {
			loadBalancerMonitor.ExpectedCodes = expectedCodes.(string)
		} else {
			return diag.FromErr(fmt.Errorf("expected_codes must be set"))
		}

		if followRedirects, ok := d.GetOk("follow_redirects"); ok {
			loadBalancerMonitor.FollowRedirects = followRedirects.(bool)
		}

		if method, ok := d.GetOk("method"); ok {
			loadBalancerMonitor.Method = method.(string)
		} else {
			loadBalancerMonitor.Method = "GET"
		}

		if header, ok := d.GetOk("header"); ok {
			loadBalancerMonitor.Header = expandLoadBalancerMonitorHeader(header)
		}

		if path, ok := d.GetOk("path"); ok {
			loadBalancerMonitor.Path = path.(string)
		} else {
			loadBalancerMonitor.Path = "/"
		}

		if probeZone, ok := d.GetOk("probe_zone"); ok {
			loadBalancerMonitor.ProbeZone = probeZone.(string)
		} else {
			loadBalancerMonitor.ProbeZone = ""
		}
	}

	log.Printf("[DEBUG] Creating Cloudflare Load Balancer Monitor from struct: %+v", loadBalancerMonitor)

	r, err := client.CreateLoadBalancerMonitor(context.Background(), loadBalancerMonitor)
	if err != nil {
		return err.Wrap(err, "error creating load balancer monitor")
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in create response; resource was empty"))
	}

	d.SetId(r.ID)

	log.Printf("[INFO] New Cloudflare Load Balancer Monitor created with  ID: %s", d.Id())

	return resourceCloudflareLoadBalancerPoolMonitorRead(d, meta)
}

func resourceCloudflareLoadBalancerPoolMonitorUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	loadBalancerMonitor := cloudflare.LoadBalancerMonitor{
		ID:       d.Id(),
		Timeout:  d.Get("timeout").(int),
		Type:     d.Get("type").(string),
		Interval: d.Get("interval").(int),
		Retries:  d.Get("retries").(int),
	}

	if description, ok := d.GetOk("description"); ok {
		loadBalancerMonitor.Description = description.(string)
	}

	if port, ok := d.GetOk("port"); ok {
		loadBalancerMonitor.Port = uint16(port.(int))
	}

	switch loadBalancerMonitor.Type {
	case "tcp":
		if method, ok := d.GetOk("method"); ok {
			loadBalancerMonitor.Method = method.(string)
		} else {
			loadBalancerMonitor.Method = "connection_established"
		}
	case "http", "https":
		if allowInsecure, ok := d.GetOk("allow_insecure"); ok {
			loadBalancerMonitor.AllowInsecure = allowInsecure.(bool)
		}

		if expectedBody, ok := d.GetOk("expected_body"); ok {
			loadBalancerMonitor.ExpectedBody = expectedBody.(string)
		} else {
			loadBalancerMonitor.ExpectedBody = ""
		}

		if expectedCodes, ok := d.GetOk("expected_codes"); ok {
			loadBalancerMonitor.ExpectedCodes = expectedCodes.(string)
		} else {
			return diag.FromErr(fmt.Errorf("expected_codes must be set"))
		}

		if header, ok := d.GetOk("header"); ok {
			loadBalancerMonitor.Header = expandLoadBalancerMonitorHeader(header)
		}

		if followRedirects, ok := d.GetOk("follow_redirects"); ok {
			loadBalancerMonitor.FollowRedirects = followRedirects.(bool)
		}

		if method, ok := d.GetOk("method"); ok {
			loadBalancerMonitor.Method = method.(string)
		} else {
			loadBalancerMonitor.Method = "GET"
		}

		if path, ok := d.GetOk("path"); ok {
			loadBalancerMonitor.Path = path.(string)
		} else {
			loadBalancerMonitor.Path = "/"
		}

		if probeZone, ok := d.GetOk("probe_zone"); ok {
			loadBalancerMonitor.ProbeZone = probeZone.(string)
		} else {
			loadBalancerMonitor.ProbeZone = ""
		}
	}

	log.Printf("[DEBUG] Update Cloudflare Load Balancer Monitor from struct: %+v", loadBalancerMonitor)

	_, err := client.ModifyLoadBalancerMonitor(context.Background(), loadBalancerMonitor)
	if err != nil {
		return err.Wrap(err, "error modifying load balancer monitor")
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

func resourceCloudflareLoadBalancerPoolMonitorRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	loadBalancerMonitor, err := client.LoadBalancerMonitorDetails(context.Background(), d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Load balancer monitor %s no longer exists", d.Id())
			d.SetId("")
			return nil
		} else {
			return err.Wrap(err,
				fmt.Sprintf("Error reading load balancer monitor from API for resource %s ", d.Id()))
		}
	}
	log.Printf("[DEBUG] Read Cloudflare Load Balancer Monitor from API as struct: %+v", loadBalancerMonitor)

	if loadBalancerMonitor.Type == "http" || loadBalancerMonitor.Type == "https" {
		d.Set("allow_insecure", loadBalancerMonitor.AllowInsecure)
		d.Set("expected_body", loadBalancerMonitor.ExpectedBody)
		d.Set("expected_codes", loadBalancerMonitor.ExpectedCodes)
		d.Set("follow_redirects", loadBalancerMonitor.FollowRedirects)
		d.Set("path", loadBalancerMonitor.Path)
		d.Set("probe_zone", loadBalancerMonitor.ProbeZone)

		if err := d.Set("header", flattenLoadBalancerMonitorHeader(loadBalancerMonitor.Header)); err != nil {
			log.Printf("[WARN] Error setting header for load balancer monitor %q: %s", d.Id(), err)
		}
	}

	d.Set("description", loadBalancerMonitor.Description)
	d.Set("interval", loadBalancerMonitor.Interval)
	d.Set("method", loadBalancerMonitor.Method)
	d.Set("port", int(loadBalancerMonitor.Port))
	d.Set("retries", loadBalancerMonitor.Retries)
	d.Set("timeout", loadBalancerMonitor.Timeout)
	d.Set("type", loadBalancerMonitor.Type)
	d.Set("created_on", loadBalancerMonitor.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", loadBalancerMonitor.ModifiedOn.Format(time.RFC3339Nano))

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

func resourceCloudflareLoadBalancerPoolMonitorDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting Cloudflare Load Balancer Monitor: %s ", d.Id())

	err := client.DeleteLoadBalancerMonitor(context.Background(), d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "HTTP status 404") {
			log.Printf("[INFO] Load balancer monitor %s no longer exists", d.Id())
			return nil
		} else {
			return err.Wrap(err, "error deleting cloudflare load balancer monitor")
		}
	}

	return nil
}
