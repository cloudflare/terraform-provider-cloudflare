package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareLoadBalancerMonitor() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareLoadBalancerMonitorSchema(),
		CreateContext: resourceCloudflareLoadBalancerPoolMonitorCreate,
		ReadContext:   resourceCloudflareLoadBalancerPoolMonitorRead,
		UpdateContext: resourceCloudflareLoadBalancerPoolMonitorUpdate,
		DeleteContext: resourceCloudflareLoadBalancerPoolMonitorDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareLoadBalancerPoolMonitorImport,
		},
		Description: heredoc.Doc(`
			If Cloudflare's Load Balancing to load-balance across multiple
			origin servers or data centers, you configure one of these Monitors
			to actively check the availability of those servers over HTTP(S) or
			TCP.
		`),
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

	tflog.Debug(ctx, fmt.Sprintf("Creating Cloudflare Load Balancer Monitor from struct: %+v", loadBalancerMonitor))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	r, err := client.CreateLoadBalancerMonitor(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.CreateLoadBalancerMonitorParams{LoadBalancerMonitor: loadBalancerMonitor})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error creating load balancer monitor"))
	}

	if r.ID == "" {
		return diag.FromErr(fmt.Errorf("failed to find id in create response; resource was empty"))
	}

	d.SetId(r.ID)

	tflog.Info(ctx, fmt.Sprintf("New Cloudflare Load Balancer Monitor created with  ID: %s", d.Id()))

	return resourceCloudflareLoadBalancerPoolMonitorRead(ctx, d, meta)
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

	tflog.Debug(ctx, fmt.Sprintf("Update Cloudflare Load Balancer Monitor from struct: %+v", loadBalancerMonitor))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	_, err := client.UpdateLoadBalancerMonitor(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.UpdateLoadBalancerMonitorParams{LoadBalancerMonitor: loadBalancerMonitor})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "error modifying load balancer monitor"))
	}

	tflog.Info(ctx, fmt.Sprintf("Cloudflare Load Balancer Monitor %q was modified", d.Id()))

	return resourceCloudflareLoadBalancerPoolMonitorRead(ctx, d, meta)
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

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	loadBalancerMonitor, err := client.GetLoadBalancerMonitor(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Load balancer monitor %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		} else {
			return diag.FromErr(errors.Wrap(err,
				fmt.Sprintf("Error reading load balancer monitor from API for resource %s ", d.Id())))
		}
	}
	tflog.Debug(ctx, fmt.Sprintf("Read Cloudflare Load Balancer Monitor from API as struct: %+v", loadBalancerMonitor))

	if loadBalancerMonitor.Type == "http" || loadBalancerMonitor.Type == "https" {
		d.Set("allow_insecure", loadBalancerMonitor.AllowInsecure)
		d.Set("expected_body", loadBalancerMonitor.ExpectedBody)
		d.Set("expected_codes", loadBalancerMonitor.ExpectedCodes)
		d.Set("follow_redirects", loadBalancerMonitor.FollowRedirects)
		d.Set("path", loadBalancerMonitor.Path)
		d.Set("probe_zone", loadBalancerMonitor.ProbeZone)

		if err := d.Set("header", flattenLoadBalancerMonitorHeader(loadBalancerMonitor.Header)); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Error setting header for load balancer monitor %q: %s", d.Id(), err))
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

	tflog.Info(ctx, fmt.Sprintf("Deleting Cloudflare Load Balancer Monitor: %s ", d.Id()))

	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	err := client.DeleteLoadBalancerMonitor(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Load balancer monitor %s no longer exists", d.Id()))
			return nil
		} else {
			return diag.FromErr(errors.Wrap(err, "error deleting cloudflare load balancer monitor"))
		}
	}

	return nil
}

func resourceCloudflareLoadBalancerPoolMonitorImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var lbMonitorID string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		lbMonitorID = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/loadBalancerMonitorID\"", d.Id())
	}

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.SetId(lbMonitorID)

	resourceCloudflareLoadBalancerPoolMonitorRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
