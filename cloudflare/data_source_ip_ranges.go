package cloudflare

import (
	"encoding/json"
	"sort"
	"strconv"

	cf "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceCloudflareIPRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareIPRangesRead,

		Schema: map[string]*schema.Schema{
			"cidr_blocks_v4": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cidr_blocks_v6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceCloudflareIPRangesRead(d *schema.ResourceData, meta interface{}) error {
	ips, err := cf.IPs()
	if err != nil {
		return err
	}

	serialized, err := json.Marshal(ips)
	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(hashcode.String(string(serialized))))

	sort.Strings(ips.IPv4CIDRs)
	d.Set("cidr_blocks_v4", ips.IPv4CIDRs)

	sort.Strings(ips.IPv6CIDRs)
	d.Set("cidr_blocks_v6", ips.IPv6CIDRs)

	return nil
}
