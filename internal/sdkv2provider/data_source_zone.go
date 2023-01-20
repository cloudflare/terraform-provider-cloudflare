package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareZone() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareZoneRead,

		Schema: map[string]*schema.Schema{
			consts.ZoneIDSchemaKey: {
				Description:  "The zone identifier to target for the resource.",
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{consts.ZoneIDSchemaKey, "name"},
			},
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{consts.ZoneIDSchemaKey, "name"},
				Description:  "The name of the zone.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of the zone.",
			},
			"paused": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the zone is paused on Cloudflare.",
			},
			"plan": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the plan associated with the zone.",
			},
			"name_servers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Cloudflare assigned name servers. This is only populated for zones that use Cloudflare DNS.",
			},
			"vanity_name_servers": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "List of Vanity Nameservers (if set).",
			},
		},
		Description: heredoc.Doc(fmt.Sprintf(`
			Use this data source to look up [zone](https://api.cloudflare.com/#zone-properties)
			info. This is the singular alternative to %s.
		`, "`cloudflare_zones`")),
	}
}

func dataSourceCloudflareZoneRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("Reading Zones"))
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	name := d.Get("name").(string)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	var zone cloudflare.Zone
	if name != "" && zoneID == "" {
		zoneFilter := cloudflare.WithZoneFilters(name, accountID, "")
		zonesResp, err := client.ListZonesContext(ctx, zoneFilter)

		if err != nil {
			return diag.FromErr(fmt.Errorf("error listing zones: %w", err))
		}

		if zonesResp.Total > 1 {
			return diag.FromErr(fmt.Errorf("more than one zone was returned; consider adding the `account_id` to the existing resource or use the `cloudflare_zones` data source with filtering to target the zone more specifically"))
		}

		if zonesResp.Total == 0 {
			return diag.FromErr(fmt.Errorf("no zone found"))
		}

		zone = zonesResp.Result[0]
	} else {
		var err error
		zone, err = client.ZoneDetails(ctx, zoneID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error getting zone details: %w", err))
		}
	}

	d.SetId(zone.ID)
	d.Set(consts.ZoneIDSchemaKey, zone.ID)
	d.Set(consts.AccountIDSchemaKey, zone.Account.ID)
	d.Set("name", zone.Name)
	d.Set("status", zone.Status)
	d.Set("paused", zone.Paused)
	d.Set("plan", zone.Plan.Name)

	if err := d.Set("name_servers", zone.NameServers); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set name_servers attribute: %w", err))
	}

	if err := d.Set("vanity_name_servers", zone.VanityNS); err != nil {
		return diag.FromErr(fmt.Errorf("failed to set vanity_name_servers attribute: %w", err))
	}

	return nil
}
