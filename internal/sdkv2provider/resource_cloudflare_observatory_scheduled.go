package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
)

func resourceCloudflareObservatoryScheduledTest() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareObservatoryScheduledTestSchema(),

		CreateContext: resourceCloudflareObservatoryScheduledTestCreate,
		ReadContext:   resourceCloudflareObservatoryScheduledTestRead,
		DeleteContext: resourceCloudflareObservatoryScheduledTestDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareObservatoryScheduledTestImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Second),
		},
		Description: "Provides a Cloudflare Observatory Scheduled Test resource.",
	}
}

func resourceCloudflareObservatoryScheduledTestCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	params := cloudflare.CreateObservatoryScheduledPageTestParams{
		URL:       d.Get("url").(string),
		Region:    d.Get("region").(string),
		Frequency: d.Get("frequency").(string),
	}
	test, err := client.CreateObservatoryScheduledPageTest(ctx, cloudflare.ZoneIdentifier(zoneID), params)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating observatory scheduled test %q: %w", params.URL, err))
	}

	d.SetId(stringChecksum(fmt.Sprintf("%s:%s", test.Schedule.URL, test.Schedule.Region)))

	return resourceCloudflareObservatoryScheduledTestRead(ctx, d, meta)
}

func resourceCloudflareObservatoryScheduledTestRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	schedule, err := client.GetObservatoryScheduledPageTest(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.GetObservatoryScheduledPageTestParams{
		URL:    d.Get("url").(string),
		Region: d.Get("region").(string),
	})
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, fmt.Sprintf("Removing observatory scheduled test from state because it's not found in API"))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error getting observatory scheduled test %q: %w", d.Id(), err))
	}
	d.SetId(stringChecksum(fmt.Sprintf("%s:%s", schedule.URL, schedule.Region)))
	d.Set("url", schedule.URL)
	d.Set("region", schedule.Region)
	d.Set("frequency", schedule.Frequency)
	return nil
}

func resourceCloudflareObservatoryScheduledTestDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	_, err := client.DeleteObservatoryScheduledPageTest(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.DeleteObservatoryScheduledPageTestParams{
		URL:    d.Get("url").(string),
		Region: d.Get("region").(string),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting observatory scheduled test %q: %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareObservatoryScheduledTestImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	idAttr := strings.SplitN(d.Id(), ":", 3)
	var zoneID string
	var url string
	var region string
	if len(idAttr) == 3 {
		zoneID = idAttr[0]
		url = idAttr[1]
		region = idAttr[2]
	} else {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID:url:region\" for import", d.Id())
	}

	schedule, err := client.GetObservatoryScheduledPageTest(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.GetObservatoryScheduledPageTestParams{
		URL:    url,
		Region: region,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch web analytics site: %s", url)
	}

	d.SetId(stringChecksum(fmt.Sprintf("%s:%s", schedule.URL, schedule.Region)))
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("url", schedule.URL)
	d.Set("region", schedule.Region)
	d.Set("frequency", schedule.Frequency)

	return []*schema.ResourceData{d}, nil
}
