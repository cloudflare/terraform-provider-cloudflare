package sdkv2provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareHealthcheck() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCloudflareHealthcheckCreate,
		ReadContext:   resourceCloudflareHealthcheckRead,
		UpdateContext: resourceCloudflareHealthcheckUpdate,
		DeleteContext: resourceCloudflareHealthcheckDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareHealthcheckImport,
		},

		Schema: resourceCloudflareHealthcheckSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
		},
		Description: heredoc.Doc(`
			Standalone Health Checks provide a way to monitor origin servers
			without needing a Cloudflare Load Balancer.
		`),
	}
}

func resourceCloudflareHealthcheckRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	healthcheck, err := client.Healthcheck(ctx, zoneID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "object does not exist") {
			tflog.Info(ctx, fmt.Sprintf("Healthcheck %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error reading healthcheck information for %q", d.Id())))
	}

	switch healthcheck.Type {
	case "TCP":
		d.Set("method", healthcheck.TCPConfig.Method)
		d.Set("port", int(healthcheck.TCPConfig.Port))
	case "HTTP", "HTTPS":
		d.Set("method", healthcheck.HTTPConfig.Method)
		d.Set("port", int(healthcheck.HTTPConfig.Port))
		d.Set("path", healthcheck.HTTPConfig.Path)
		d.Set("expected_codes", healthcheck.HTTPConfig.ExpectedCodes)
		d.Set("expected_body", healthcheck.HTTPConfig.ExpectedBody)
		d.Set("follow_redirects", healthcheck.HTTPConfig.FollowRedirects)
		d.Set("allow_insecure", healthcheck.HTTPConfig.AllowInsecure)

		if err := d.Set("header", flattenHealthcheckHeader(healthcheck.HTTPConfig.Header)); err != nil {
			tflog.Warn(ctx, fmt.Sprintf("Error setting header for standalone healthcheck %q: %s", d.Id(), err))
		}
	}

	d.Set("name", healthcheck.Name)
	d.Set("description", healthcheck.Description)
	d.Set("suspended", healthcheck.Suspended)
	d.Set("address", healthcheck.Address)
	d.Set("consecutive_fails", healthcheck.ConsecutiveFails)
	d.Set("consecutive_successes", healthcheck.ConsecutiveSuccesses)
	d.Set("retries", healthcheck.Retries)
	d.Set("timeout", healthcheck.Timeout)
	d.Set("interval", healthcheck.Interval)
	d.Set("type", healthcheck.Type)
	d.Set("created_on", healthcheck.CreatedOn.Format(time.RFC3339Nano))
	d.Set("modified_on", healthcheck.ModifiedOn.Format(time.RFC3339Nano))
	d.Set("check_regions", healthcheck.CheckRegions)

	return nil
}

func resourceCloudflareHealthcheckCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	healthcheck, err := healthcheckSetStruct(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating healthcheck struct")))
	}

	retry := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		hc, err := client.CreateHealthcheck(ctx, zoneID, healthcheck)
		if err != nil {
			if strings.Contains(err.Error(), "no such host") {
				return resource.RetryableError(fmt.Errorf("hostname resolution failed"))
			}

			return resource.NonRetryableError(errors.Wrap(err, fmt.Sprintf("error creating standalone healthcheck")))
		}

		d.SetId(hc.ID)

		resourceCloudflareHealthcheckRead(ctx, d, meta)
		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareHealthcheckUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	healthcheck, err := healthcheckSetStruct(d)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating healthcheck struct")))
	}

	_, err = client.UpdateHealthcheck(ctx, zoneID, d.Id(), healthcheck)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error creating healthcheck")))
	}

	return resourceCloudflareHealthcheckRead(ctx, d, meta)
}

func resourceCloudflareHealthcheckDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteHealthcheck(ctx, zoneID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error deleting standalone healthcheck")))
	}

	return nil
}

func resourceCloudflareHealthcheckImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/HealthcheckId\"", d.Id())
	}

	zoneID, HealthcheckID := attributes[0], attributes[1]

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(HealthcheckID)

	resourceCloudflareHealthcheckRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func healthcheckSetStruct(d *schema.ResourceData) (cloudflare.Healthcheck, error) {
	healthcheck := cloudflare.Healthcheck{
		Name:                 d.Get("name").(string),
		Address:              d.Get("address").(string),
		Type:                 d.Get("type").(string),
		Retries:              d.Get("retries").(int),
		Timeout:              d.Get("timeout").(int),
		Interval:             d.Get("interval").(int),
		ConsecutiveFails:     d.Get("consecutive_fails").(int),
		ConsecutiveSuccesses: d.Get("consecutive_successes").(int),
	}

	if description, ok := d.GetOk("description"); ok {
		healthcheck.Description = description.(string)
	}

	if suspended, ok := d.GetOk("suspended"); ok {
		healthcheck.Suspended = suspended.(bool)
	}

	if region, ok := d.GetOk("check_regions"); ok {
		healthcheck.CheckRegions = expandInterfaceToStringList(region)
	}

	switch healthcheck.Type {
	case "TCP":
		tcpConfig := new(cloudflare.HealthcheckTCPConfig)

		if method, ok := d.GetOk("method"); ok {
			if method != "connection_established" {
				return cloudflare.Healthcheck{}, errors.New(fmt.Sprintf("cannot use %s as method for TCP healthchecks", method))
			}
			tcpConfig.Method = method.(string)
		} else {
			tcpConfig.Method = "connection_established"
		}

		if port, ok := d.GetOk("port"); ok {
			tcpConfig.Port = uint16(port.(int))
		}

		healthcheck.TCPConfig = tcpConfig
	case "HTTP", "HTTPS":
		httpConfig := new(cloudflare.HealthcheckHTTPConfig)

		if method, ok := d.GetOk("method"); ok {
			if method != "GET" && method != "HEAD" {
				return cloudflare.Healthcheck{}, errors.New(fmt.Sprintf("cannot use %s as method for HTTP/HTTPS healthchecks", method))
			}
			httpConfig.Method = method.(string)
		} else {
			httpConfig.Method = "GET"
		}

		if port, ok := d.GetOk("port"); ok {
			httpConfig.Port = uint16(port.(int))
		}

		if path, ok := d.GetOk("path"); ok {
			httpConfig.Path = path.(string)
		}

		if expectedCode, ok := d.GetOk("expected_codes"); ok {
			httpConfig.ExpectedCodes = expandInterfaceToStringList(expectedCode)
		}

		if expectedBody, ok := d.GetOk("expected_body"); ok {
			httpConfig.ExpectedBody = expectedBody.(string)
		}

		if followRedirects, ok := d.GetOk("follow_redirects"); ok {
			httpConfig.FollowRedirects = followRedirects.(bool)
		}

		if allowInsecure, ok := d.GetOk("allow_insecure"); ok {
			httpConfig.AllowInsecure = allowInsecure.(bool)
		}

		if header, ok := d.GetOk("header"); ok {
			httpConfig.Header = expandHealthcheckHeader(header)
		}

		healthcheck.HTTPConfig = httpConfig
	}

	return healthcheck, nil
}

func flattenHealthcheckHeader(header map[string][]string) *schema.Set {
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

func expandHealthcheckHeader(cfgSet interface{}) map[string][]string {
	header := make(map[string][]string)
	cfgList := cfgSet.(*schema.Set).List()
	for _, item := range cfgList {
		cfg := item.(map[string]interface{})
		header[cfg["header"].(string)] = expandInterfaceToStringList(cfg["values"].(*schema.Set).List())
	}
	return header
}
