package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareIPsecTunnel() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareIPsecTunnelSchema(),
		CreateContext: resourceCloudflareIPsecTunnelCreate,
		ReadContext:   resourceCloudflareIPsecTunnelRead,
		UpdateContext: resourceCloudflareIPsecTunnelUpdate,
		DeleteContext: resourceCloudflareIPsecTunnelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareIPsecTunnelImport,
		},
	}
}

func resourceCloudflareIPsecTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get("account_id").(string)
	client := meta.(*cloudflare.API)

	newTunnel, err := client.CreateMagicTransitIPsecTunnels(ctx, accountID, []cloudflare.MagicTransitIPsecTunnel{
		IPsecTunnelFromResource(d),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating IPSec tunnel %s: %w", d.Get("name").(string), err))
	}

	d.SetId(newTunnel[0].ID)

	return resourceCloudflareIPsecTunnelRead(ctx, d, meta)
}

func resourceCloudflareIPsecTunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid id (\"%s\") specified, should be in format \"accountID/tunnelID\"", d.Id()))
	}

	accountID, tunnelID := attributes[0], attributes[1]
	d.SetId(tunnelID)
	d.Set("account_id", accountID)

	readDiags := resourceCloudflareIPsecTunnelRead(ctx, d, meta)
	if readDiags != nil {
		return nil, errors.New("failed to read IPSec Tunnel state")
	}

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareIPsecTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get("account_id").(string)
	client := meta.(*cloudflare.API)

	tunnel, err := client.GetMagicTransitIPsecTunnel(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "IPsec tunnel not found") {
			log.Printf("[INFO] IPsec tunnel %s not found", d.Id())
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading IPsec tunnel ID %q: %w", d.Id(), err))
	}

	d.Set("name", tunnel.Name)
	d.Set("customer_endpoint", tunnel.CustomerEndpoint)
	d.Set("cloudflare_endpoint", tunnel.CloudflareEndpoint)
	d.Set("interface_address", tunnel.InterfaceAddress)

	if len(tunnel.Description) > 0 {
		d.Set("description", tunnel.Description)
	}

	return nil
}

func resourceCloudflareIPsecTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get("account_id").(string)
	client := meta.(*cloudflare.API)

	_, err := client.UpdateMagicTransitIPsecTunnel(ctx, accountID, d.Id(), IPsecTunnelFromResource(d))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error updating IPsec tunnel %q", d.Id())))
	}

	return resourceCloudflareIPsecTunnelRead(ctx, d, meta)
}

func resourceCloudflareIPsecTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get("account_id").(string)
	client := meta.(*cloudflare.API)

	log.Printf("[INFO] Deleting IPsec tunnel:  %s", d.Id())

	_, err := client.DeleteMagicTransitIPsecTunnel(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting IPsec tunnel: %w", err))
	}

	return nil
}

func IPsecTunnelFromResource(d *schema.ResourceData) cloudflare.MagicTransitIPsecTunnel {
	tunnel := cloudflare.MagicTransitIPsecTunnel{
		Name:               d.Get("name").(string),
		CustomerEndpoint:   d.Get("customer_endpoint").(string),
		CloudflareEndpoint: d.Get("cloudflare_endpoint").(string),
		InterfaceAddress:   d.Get("interface_address").(string),
	}

	description, descriptionOk := d.GetOk("description")
	if descriptionOk {
		tunnel.Description = description.(string)
	}

	return tunnel
}
