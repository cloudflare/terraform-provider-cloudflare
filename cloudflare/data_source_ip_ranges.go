package cloudflare

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	urlIPV4s = "https://www.cloudflare.com/ips-v4"
	urlIPV6s = "https://www.cloudflare.com/ips-v6"
)

func dataSourceCloudflareIPRanges() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareIPRangesRead,

		Schema: map[string]*schema.Schema{
			"cidr_blocks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ipv4_cidr_blocks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ipv6_cidr_blocks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceCloudflareIPRangesRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] Reading IPv4 ranges")
	ipv4s, err := dataSourceCloudflareIPRangesGet(urlIPV4s)
	if err != nil {
		return fmt.Errorf("Error listing IPV4 ranges: %s", err)
	}
	sort.Strings(ipv4s)

	log.Printf("[DEBUG] Reading IPv6 ranges")
	ipv6s, err := dataSourceCloudflareIPRangesGet(urlIPV6s)
	if err != nil {
		return fmt.Errorf("Error listing IPV6 ranges: %s", err)
	}
	sort.Strings(ipv6s)

	all := append([]string{}, ipv4s...)
	all = append(all, ipv6s...)

	d.SetId(strconv.Itoa(hashcode.String(strings.Join(all, "|"))))

	if err := d.Set("cidr_blocks", all); err != nil {
		return fmt.Errorf("Error setting all cidr blocks: %s", err)
	}

	if err := d.Set("ipv4_cidr_blocks", ipv4s); err != nil {
		return fmt.Errorf("Error setting ipv4 cidr blocks: %s", err)
	}

	if err := d.Set("ipv6_cidr_blocks", ipv4s); err != nil {
		return fmt.Errorf("Error setting ipv6 cidr blocks: %s", err)
	}

	return nil
}

// dataSourceCloudflareIPRangesGet performs an HTTP GET on the given URL and
// parses each line as an IP address.
func dataSourceCloudflareIPRangesGet(url string) ([]string, error) {
	conn := cleanhttp.DefaultClient()

	res, err := conn.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	list := strings.Split(strings.TrimSpace(string(body)), "\n")
	return list, nil
}
