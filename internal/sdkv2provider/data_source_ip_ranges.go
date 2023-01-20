package sdkv2provider

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareIPRanges() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareIPRangesRead,
		Schema: map[string]*schema.Schema{
			"cidr_blocks": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The lexically ordered list of all non-China CIDR blocks.",
			},
			"ipv4_cidr_blocks": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The lexically ordered list of only the IPv4 CIDR blocks.",
			},
			"ipv6_cidr_blocks": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The lexically ordered list of only the IPv6 CIDR blocks.",
			},
			"china_ipv4_cidr_blocks": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The lexically ordered list of only the IPv4 China CIDR blocks.",
			},
			"china_ipv6_cidr_blocks": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The lexically ordered list of only the IPv6 China CIDR blocks.",
			},
		},
		Description: "Use this data source to get the [IP ranges](https://www.cloudflare.com/ips/) of Cloudflare network.",
	}
}

func dataSourceCloudflareIPRangesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ranges, err := cloudflare.IPs()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Cloudflare IP ranges: %w", err))
	}

	IPv4s := ranges.IPv4CIDRs
	IPv6s := ranges.IPv6CIDRs
	chinaIPv4s := ranges.ChinaIPv4CIDRs
	chinaIPv6s := ranges.ChinaIPv6CIDRs

	sort.Strings(IPv4s)
	sort.Strings(IPv6s)
	sort.Strings(chinaIPv4s)
	sort.Strings(chinaIPv6s)

	all := append([]string{}, IPv4s...)
	all = append(all, IPv6s...)
	sort.Strings(all)

	d.SetId(strconv.Itoa(hashCodeString(strings.Join(all, "|"))))

	if err := d.Set("cidr_blocks", all); err != nil {
		return diag.FromErr(fmt.Errorf("error setting all cidr blocks: %w", err))
	}

	if err := d.Set("ipv4_cidr_blocks", IPv4s); err != nil {
		return diag.FromErr(fmt.Errorf("error setting ipv4 cidr blocks: %w", err))
	}

	if err := d.Set("ipv6_cidr_blocks", IPv6s); err != nil {
		return diag.FromErr(fmt.Errorf("error setting ipv6 cidr blocks: %w", err))
	}

	if err := d.Set("china_ipv4_cidr_blocks", chinaIPv4s); err != nil {
		return diag.FromErr(fmt.Errorf("error setting china ipv4 cidr blocks: %w", err))
	}

	if err := d.Set("china_ipv6_cidr_blocks", chinaIPv6s); err != nil {
		return diag.FromErr(fmt.Errorf("error setting china ipv6 cidr blocks: %w", err))
	}

	return nil
}
