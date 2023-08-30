package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWebAnalyticsSite() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareWebAnalyticsSiteSchema(),

		CreateContext: resourceCloudflareWebAnalyticsSiteCreate,
		ReadContext:   resourceCloudflareWebAnalyticsSiteRead,
		//UpdateContext: resourceCloudflareWebAnalyticsSiteUpdate,
		DeleteContext: resourceCloudflareWebAnalyticsSiteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareWebAnalyticsSiteImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
		},
		Description: "Provides a Cloudflare Web Analytics Site resource.",
	}
}

func buildCreateWebAnalyticsSiteParams(d *schema.ResourceData) cloudflare.CreateWebAnalyticsSiteParams {
	return cloudflare.CreateWebAnalyticsSiteParams{
		Host:        d.Get("host").(string),
		ZoneTag:     d.Get("zone_tag").(string),
		AutoInstall: cloudflare.BoolPtr(d.Get("auto_install").(bool)),
	}
}

func resourceCloudflareWebAnalyticsSiteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := buildCreateWebAnalyticsSiteParams(d)
	site, err := client.CreateWebAnalyticsSite(ctx, cloudflare.AccountIdentifier(accountID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating web analytics site %q: %w", d.Id(), err))
	}

	d.SetId(site.SiteTag)

	return resourceCloudflareWebAnalyticsSiteRead(ctx, d, meta)
}

func resourceCloudflareWebAnalyticsSiteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := cloudflare.UpdateWebAnalyticsSiteParams{
		SiteTag:     d.Id(),
		Host:        d.Get("host").(string),
		ZoneTag:     d.Get("zone_tag").(string),
		AutoInstall: cloudflare.BoolPtr(d.Get("auto_install").(bool)),
	}
	site, err := client.UpdateWebAnalyticsSite(ctx, cloudflare.AccountIdentifier(accountID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating web analytics site %q: %w", d.Id(), err))
	}

	d.SetId(site.SiteTag)

	return resourceCloudflareWebAnalyticsSiteRead(ctx, d, meta)
}

func resourceCloudflareWebAnalyticsSiteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	siteTag := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	site, err := client.GetWebAnalyticsSite(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.GetWebAnalyticsSiteParams{
		SiteTag: siteTag,
	})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing web analytic site from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error getting web analytics site %q: %w", d.Id(), err))
	}
	d.SetId(site.SiteTag)
	d.Set("site_tag", site.SiteTag)
	d.Set("auto_install", site.AutoInstall)
	d.Set("site_token", site.SiteToken)
	d.Set("snippet", site.Snippet)
	d.Set("ruleset_id", site.Ruleset.ID)
	return nil
}

func resourceCloudflareWebAnalyticsSiteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	siteTag := d.Id()
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	params := cloudflare.DeleteWebAnalyticsSiteParams{
		SiteTag: siteTag,
	}
	_, err := client.DeleteWebAnalyticsSite(ctx, cloudflare.AccountIdentifier(accountID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting web analytics site %q: %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareWebAnalyticsSiteImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), "/", 2)
	var accountID string
	var siteTag string
	if len(idAttr) == 2 {
		accountID = idAttr[0]
		siteTag = idAttr[1]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/siteTag\" for import", d.Id())
	}

	site, err := client.GetWebAnalyticsSite(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.GetWebAnalyticsSiteParams{
		SiteTag: siteTag,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch web analytics site: %s", siteTag)
	}

	d.SetId(site.SiteTag)
	d.Set(consts.AccountIDSchemaKey, accountID)

	resourceCloudflareWebAnalyticsSiteRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
