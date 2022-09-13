package provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareRecord() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareRecordRead,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Description: "The zone identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"record_id": {
				Description: "The record identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxied": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"proxiable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"locked": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"zone_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudflareRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	recordID := d.Get("record_id").(string)
	record, err := client.DNSRecord(ctx, zoneID, recordID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding record %q: %w", recordID, err))
	}
	d.SetId(record.ID)
	d.Set("hostname", record.Name)
	d.Set("type", record.Type)
	d.Set("value", record.Content)
	d.Set("proxied", record.Proxied)
	d.Set("ttl", record.TTL)
	d.Set("proxiable", record.Proxiable)
	d.Set("locked", record.Locked)
	d.Set("zone_name", record.ZoneName)
	return nil
}
