package sdkv2provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomPages() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareCustomPagesSchema(),
		CreateContext: resourceCloudflareCustomPagesUpdate,
		ReadContext:   resourceCloudflareCustomPagesRead,
		UpdateContext: resourceCloudflareCustomPagesUpdate,
		DeleteContext: resourceCloudflareCustomPagesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareCustomPagesImport,
		},
		Description: "Provides a resource which manages Cloudflare custom error pages.",
	}
}

func resourceCloudflareCustomPagesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	pageType := d.Get("type").(string)

	if accountID == "" && zoneID == "" {
		return diag.FromErr(fmt.Errorf("either `account_id` or `zone_id` must be set"))
	}

	var (
		pageOptions cloudflare.CustomPageOptions
		identifier  string
	)

	if accountID != "" {
		pageOptions = cloudflare.CustomPageOptions{AccountID: accountID}
		identifier = accountID
	} else {
		pageOptions = cloudflare.CustomPageOptions{ZoneID: zoneID}
		identifier = zoneID
	}

	page, err := client.CustomPage(ctx, &pageOptions, pageType)
	if err != nil {
		return diag.FromErr(err)
	}

	// If the `page.State` comes back as "default", it's safe to assume we
	// don't need to keep the ID managed anymore as it will be relying on
	// Cloudflare's default pages.
	if page.State == "default" {
		log.Printf("[INFO] removing custom page configuration for '%s' as it is marked as being in the default state", pageType)
		d.SetId("")
		return nil
	}

	checksum := stringChecksum(fmt.Sprintf("%s/%s", identifier, page.ID))
	d.SetId(checksum)

	d.Set("state", page.State)
	d.Set("url", page.URL)
	d.Set("type", page.ID)

	return nil
}

func resourceCloudflareCustomPagesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var pageOptions cloudflare.CustomPageOptions
	if accountID != "" {
		pageOptions = cloudflare.CustomPageOptions{AccountID: accountID}
	} else {
		pageOptions = cloudflare.CustomPageOptions{ZoneID: zoneID}
	}

	pageType := d.Get("type").(string)
	customPageParameters := cloudflare.CustomPageParameters{
		URL:   d.Get("url").(string),
		State: "customized",
	}
	_, err := client.UpdateCustomPage(ctx, &pageOptions, pageType, customPageParameters)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to update '%s' custom page", pageType)))
	}

	return resourceCloudflareCustomPagesRead(ctx, d, meta)
}

func resourceCloudflareCustomPagesDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	var pageOptions cloudflare.CustomPageOptions
	if accountID != "" {
		pageOptions = cloudflare.CustomPageOptions{AccountID: accountID}
	} else {
		pageOptions = cloudflare.CustomPageOptions{ZoneID: zoneID}
	}

	pageType := d.Get("type").(string)
	customPageParameters := cloudflare.CustomPageParameters{
		URL:   nil,
		State: "default",
	}
	_, err := client.UpdateCustomPage(ctx, &pageOptions, pageType, customPageParameters)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to update '%s' custom page", pageType)))
	}

	return resourceCloudflareCustomPagesRead(ctx, d, meta)
}

func resourceCloudflareCustomPagesImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)
	if len(attributes) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"requestType/ID/pageType\"", d.Id())
	}
	requestType, identifier, pageType := attributes[0], attributes[1], attributes[2]

	d.Set("type", pageType)

	if requestType == "account" {
		d.Set(consts.AccountIDSchemaKey, identifier)
	} else {
		d.Set(consts.ZoneIDSchemaKey, identifier)
	}

	checksum := stringChecksum(fmt.Sprintf("%s/%s", identifier, pageType))
	d.SetId(checksum)

	resourceCloudflareCustomPagesRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
