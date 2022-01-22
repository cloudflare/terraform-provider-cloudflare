package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareIPsecTunnel() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareIPsecTunnelSchema(),
		Create: resourceCloudflareIPsecTunnelCreate,
		Read:   resourceCloudflareIPsecTunnelRead,
		Update: resourceCloudflareIPsecTunnelUpdate,
		Delete: resourceCloudflareIPsecTunnelDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareIPsecTunnelImport,
		},
	}
}

func resourceCloudflareIPsecTunnelCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	newTunnel, err := client.CreateMagicTransitIPsecTunnels(context.Background(), []cloudflare.MagicTransitIPsecTunnel{
		IPsecTunnelFromResource(d),
	})

	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error creating IPSec tunnel %s", d.Get("name").(string)))
	}

	d.SetId(newTunnel[0].ID)

	return resourceCloudflareIPsecTunnelRead(d, meta)
}

func resourceCloudflareIPsecTunnelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/tunnelID\"", d.Id())
	}

	accountID, tunnelID := attributes[0], attributes[1]
	d.SetId(tunnelID)
	d.Set("account_id", accountID)
	client.AccountID = accountID

	resourceCloudflareIPsecTunnelRead(d, meta)

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareIPsecTunnelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	tunnel, err := client.GetMagicTransitIPsecTunnel(context.Background(), d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "IPsec tunnel not found") {
			log.Printf("[INFO] IPsec tunnel %s not found", d.Id())
			d.SetId("")
			return nil
		}
		return errors.Wrap(err, fmt.Sprintf("error reading IPsec tunnel ID %q", d.Id()))
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

func resourceCloudflareIPsecTunnelUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	_, err := client.UpdateMagicTransitIPsecTunnel(context.Background(), d.Id(), IPsecTunnelFromResource(d))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error updating IPsec tunnel %q", d.Id()))
	}

	return resourceCloudflareStaticRouteRead(d, meta)
}

func resourceCloudflareIPsecTunnelDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	client.AccountID = d.Get("account_id").(string)

	log.Printf("[INFO] Deleting IPsec tunnel:  %s", d.Id())

	_, err := client.DeleteMagicTransitIPsecTunnel(context.Background(), d.Id())
	if err != nil {
		return fmt.Errorf("error deleting IPsec tunnel: %s", err)
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
