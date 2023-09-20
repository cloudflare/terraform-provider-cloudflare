package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareAuthenticatedOriginPulls() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareAuthenticatedOriginPullsSchema(),
		CreateContext: resourceCloudflareAuthenticatedOriginPullsCreate,
		ReadContext:   resourceCloudflareAuthenticatedOriginPullsRead,
		UpdateContext: resourceCloudflareAuthenticatedOriginPullsCreate,
		DeleteContext: resourceCloudflareAuthenticatedOriginPullsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAuthenticatedOriginPullsImport,
		},
		Description: heredoc.Doc(fmt.Sprintf(`
			Provides a Cloudflare Authenticated Origin Pulls resource. A %s
			resource is required to use Per-Zone or Per-Hostname Authenticated
			Origin Pulls.`, "`cloudflare_authenticated_origin_pulls`")),
	}
}

func resourceCloudflareAuthenticatedOriginPullsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hostname := d.Get("hostname").(string)
	aopCert := d.Get("authenticated_origin_pulls_certificate").(string)

	var checksum string
	isEnabled := false
	if enabledVal, ok := d.GetOk("enabled"); ok {
		// if enabled is not the zero val, use that
		isEnabled = enabledVal.(bool)
	}
	switch {
	case hostname != "" && aopCert != "":
		// Per Hostname AOP
		conf := []cloudflare.PerHostnameAuthenticatedOriginPullsConfig{{
			CertID:   aopCert,
			Hostname: hostname,
			Enabled:  isEnabled,
		}}
		_, err := client.EditPerHostnameAuthenticatedOriginPullsConfig(ctx, zoneID, conf)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating Per-Hostname Authenticated Origin Pulls resource on zone %q for hostname %s: %w", zoneID, hostname, err))
		}
		checksum = stringChecksum(fmt.Sprintf("PerHostnameAOP/%s/%s/%s", zoneID, hostname, aopCert))

	case aopCert != "":
		// Per Zone AOP
		_, err := client.SetPerZoneAuthenticatedOriginPullsStatus(ctx, zoneID, isEnabled)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating Per-Zone Authenticated Origin Pulls resource on zone %q: %w", zoneID, err))
		}
		checksum = stringChecksum(fmt.Sprintf("PerZoneAOP/%s/%s", zoneID, aopCert))

	default:
		// Global AOP
		_, err := client.SetAuthenticatedOriginPullsStatus(ctx, zoneID, isEnabled)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error creating Global Authenticated Origin Pulls resource on zone %q: %w", zoneID, err))
		}
		checksum = stringChecksum(fmt.Sprintf("GlobalAOP/%s/", zoneID))
	}

	d.SetId(checksum)
	return resourceCloudflareAuthenticatedOriginPullsRead(ctx, d, meta)
}

func resourceCloudflareAuthenticatedOriginPullsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hostname := d.Get("hostname").(string)
	aopCert := d.Get("authenticated_origin_pulls_certificate").(string)

	if hostname != "" && aopCert != "" {
		// Per Hostname AOP
		res, err := client.GetPerHostnameAuthenticatedOriginPullsConfig(ctx, zoneID, hostname)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to get Per-Hostname Authenticated Origin Pulls setting"))
		}
		d.Set("enabled", res.Enabled)
	} else if aopCert != "" {
		// Per Zone AOP
		res, err := client.GetPerZoneAuthenticatedOriginPullsStatus(ctx, zoneID)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to get Per-Zone Authenticated Origin Pulls setting"))
		}
		d.Set("enabled", res.Enabled)
	} else {
		// Global AOP
		res, err := client.GetAuthenticatedOriginPullsStatus(ctx, zoneID)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, "failed to get Global Authenticated Origin Pulls setting"))
		}
		if res.Value == "on" {
			d.Set("enabled", true)
		} else {
			d.Set("enabled", false)
		}
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	hostname := d.Get("hostname").(string)
	aopCert := d.Get("authenticated_origin_pulls_certificate").(string)

	if hostname != "" && aopCert != "" {
		// Per Hostname AOP
		conf := []cloudflare.PerHostnameAuthenticatedOriginPullsConfig{{
			CertID:   aopCert,
			Hostname: hostname,
			Enabled:  false,
		}}
		_, err := client.EditPerHostnameAuthenticatedOriginPullsConfig(ctx, zoneID, conf)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error disabling Per-Hostname Authenticated Origin Pulls resource on zone %q: %w", zoneID, err))
		}
	} else if aopCert != "" {
		// Per Zone AOP
		_, err := client.SetPerZoneAuthenticatedOriginPullsStatus(ctx, zoneID, false)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error disabling Per-Zone Authenticated Origin Pulls resource on zone %q: %w", zoneID, err))
		}
	} else {
		// Global AOP
		_, err := client.SetAuthenticatedOriginPullsStatus(ctx, zoneID, false)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error disabling Global Authenticated Origin Pulls resource on zone %q: %w", zoneID, err))
		}
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 3)

	var zoneID string
	var certID string
	var hostname string
	var checksum string

	if len(idAttr) == 1 {
		zoneID = idAttr[0]
		checksum = stringChecksum(fmt.Sprintf("GlobalAOP/%s/", zoneID))
	} else if len(idAttr) == 2 {
		zoneID, certID = idAttr[0], idAttr[1]
		d.Set("authenticated_origin_pulls_certificate", certID)
		checksum = stringChecksum(fmt.Sprintf("PerZoneAOP/%s/%s", zoneID, certID))
	} else if len(idAttr) == 3 {
		zoneID, certID, hostname = idAttr[0], idAttr[1], idAttr[2]
		d.Set("hostname", hostname)
		d.Set("authenticated_origin_pulls_certificate", certID)
		checksum = stringChecksum(fmt.Sprintf("PerHostnameAOP/%s/%s/%s", zoneID, hostname, certID))
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be maximum 3 id specified for import", d.Id())
	}

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(checksum)
	resourceCloudflareAuthenticatedOriginPullsRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
