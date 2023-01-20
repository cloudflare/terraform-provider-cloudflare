package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareWeb3Hostname() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareWeb3HostnameSchema(),
		CreateContext: resourceCloudflareWeb3HostnameCreate,
		ReadContext:   resourceCloudflareWeb3HostnameRead,
		UpdateContext: resourceCloudflareWeb3HostnameUpdate,
		DeleteContext: resourceCloudflareWeb3HostnameDelete,
		Description: heredoc.Doc(`
			Manages Web3 hostnames for IPFS and Ethereum gateways.
		`),
	}
}

func resourceCloudflareWeb3HostnameCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	hostname, err := client.CreateWeb3Hostname(ctx, cloudflare.Web3HostnameCreateParameters{
		ZoneID:      d.Get(consts.ZoneIDSchemaKey).(string),
		Name:        d.Get("name").(string),
		Target:      d.Get("target").(string),
		Description: d.Get("description").(string),
		DNSLink:     d.Get("dnslink").(string),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating web3hostname %q: %w", d.Get("name").(string), err))
	}

	d.SetId(hostname.ID)

	return resourceCloudflareWeb3HostnameRead(ctx, d, meta)
}

func resourceCloudflareWeb3HostnameRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	hostname, err := client.GetWeb3Hostname(ctx, cloudflare.Web3HostnameDetailsParameters{
		ZoneID:     d.Get(consts.ZoneIDSchemaKey).(string),
		Identifier: d.Id(),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading web3hostname %q: %w", d.Id(), err))
	}

	d.SetId(hostname.ID)

	return nil
}

func resourceCloudflareWeb3HostnameUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	hostname, err := client.UpdateWeb3Hostname(ctx, cloudflare.Web3HostnameUpdateParameters{
		ZoneID:      zoneID,
		Identifier:  d.Id(),
		Description: d.Get("description").(string),
		DNSLink:     d.Get("dnslink").(string),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating web3hostname %q: %w", d.Id(), err))
	}

	d.SetId(hostname.ID)

	return resourceCloudflareWeb3HostnameRead(ctx, d, meta)
}

func resourceCloudflareWeb3HostnameDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	_, err := client.DeleteWeb3Hostname(ctx, cloudflare.Web3HostnameDetailsParameters{
		ZoneID:     d.Get(consts.ZoneIDSchemaKey).(string),
		Identifier: d.Id(),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting web3hostname %q: %w", d.Id(), err))
	}

	return nil
}
