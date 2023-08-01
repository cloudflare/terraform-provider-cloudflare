package sdkv2provider

import (
	"context"
	"errors"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareZoneCacheReserve() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareZoneCacheReserveRead,

		Schema: map[string]*schema.Schema{
			consts.ZoneIDSchemaKey: {
				Type:        schema.TypeString,
				Required:    true,
				Description: consts.ZoneIDSchemaDescription,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The status of Cache Reserve support.",
			},
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare data source to look up Cache Reserve
			status for a given zone.
		`),
	}
}

func dataSourceCloudflareZoneCacheReserveRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	tflog.Info(ctx, "reading Cache Reserve", map[string]interface{}{
		"zone_id": zoneID,
	})

	params := cloudflare.GetCacheReserveParams{}
	output, err := client.GetCacheReserve(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			return diag.Errorf("unable to find zone: %s", zoneID)
		}
		return diag.Errorf("unable to read Cache Reserve for zone %q: %s", zoneID, err)
	}

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("enabled", output.Value == cacheReserveEnabled)

	d.SetId(stringChecksum(fmt.Sprintf("%s/cache-reserve", zoneID)))

	return nil
}
