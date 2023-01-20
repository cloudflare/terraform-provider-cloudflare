package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudflareRecord() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc(`
			Use this data source to lookup a single [DNS Record](https://api.cloudflare.com/#dns-records-for-a-zone-properties).
		`),
		ReadContext: dataSourceCloudflareRecordRead,
		Schema: map[string]*schema.Schema{
			consts.ZoneIDSchemaKey: {
				Description: "The zone identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Hostname to filter DNS record results on.",
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "A",
				ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS"}, false),
				Description:  "DNS record type to filter record results on.",
			},
			"priority": {
				Type:             schema.TypeInt,
				Optional:         true,
				DiffSuppressFunc: suppressPriority,
				Description:      "DNS priority to filter record results on.",
			},
			"value": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Value of the found DNS record.",
			},
			"proxied": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Proxied status of the found DNS record.",
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "TTL of the found DNS record.",
			},
			"proxiable": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Proxiable status of the found DNS record.",
			},
			"locked": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Locked status of the found DNS record.",
			},
			"zone_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zone name of the found DNS record.",
			},
		},
	}
}

func dataSourceCloudflareRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	searchRecord := cloudflare.ListDNSRecordsParams{
		Name: d.Get("hostname").(string),
		Type: d.Get("type").(string),
	}

	records, _, err := client.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), searchRecord)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error listing DNS records: %w", err))
	}

	if len(records) == 0 {
		return diag.Errorf("didn't get any DNS records for hostname: %s", searchRecord.Name)
	}

	if len(records) != 1 && !contains([]string{"MX", "URI"}, searchRecord.Type) {
		return diag.Errorf("only wanted 1 DNS record. Got %d records", len(records))
	} else {
		var p uint16
		if priority, ok := d.GetOkExists("priority"); ok {
			p = uint16(priority.(int))
		}
		for _, record := range records {
			if cloudflare.Uint16(record.Priority) == p {
				records = []cloudflare.DNSRecord{record}
				break
			}
		}
		if len(records) != 1 {
			return diag.Errorf("unable to find single record for %s type %s", searchRecord.Name, searchRecord.Type)
		}
	}

	record := records[0]
	d.SetId(record.ID)
	d.Set("type", record.Type)
	d.Set("value", record.Content)
	d.Set("proxied", record.Proxied)
	d.Set("ttl", record.TTL)
	d.Set("proxiable", record.Proxiable)
	d.Set("locked", record.Locked)
	d.Set("zone_name", record.ZoneName)

	if record.Priority != nil {
		d.Set("priority", int(cloudflare.Uint16(record.Priority)))
	}

	return nil
}
