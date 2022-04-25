package cloudflare

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelRoute() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareTunnelRouteSchema(),
		Create: resourceCloudflareTunnelRouteCreate,
		Read:   resourceCloudflareTunnelRouteRead,
		Update: resourceCloudflareTunnelRouteUpdate,
		Delete: resourceCloudflareTunnelRouteDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareTunnelRouteImport,
		},
	}
}

func resourceCloudflareTunnelRouteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	network := d.Get("network").(string)

	tunnelRoutes, err := client.ListTunnelRoutes(context.Background(), cloudflare.TunnelRoutesListParams{
		AccountID:       accountID,
		IsDeleted:       cloudflare.BoolPtr(false),
		NetworkSubset:   network,
		NetworkSuperset: network,
	})

	if err != nil {
		return fmt.Errorf("failed to fetch Tunnel Route: %w", err)
	}

	if len(tunnelRoutes) < 1 {
		log.Printf("[INFO] Tunnel Route for network %s in account %s not found", network, accountID)
		d.SetId("")
		return nil
	}

	tunnelRoute := tunnelRoutes[0]

	d.Set("tunnel_id", tunnelRoute.TunnelID)
	d.Set("network", tunnelRoute.Network)
	if len(tunnelRoute.Comment) > 0 {
		d.Set("comment", tunnelRoute.Comment)
	}

	return nil
}

func resourceCloudflareTunnelRouteCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	resource := cloudflare.TunnelRoutesCreateParams{
		AccountID: d.Get("account_id").(string),
		TunnelID:  d.Get("tunnel_id").(string),
		Network:   d.Get("network").(string),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	newTunnelRoute, err := client.CreateTunnelRoute(context.Background(), resource)
	if err != nil {
		return fmt.Errorf("error creating Tunnel Route for Network %q: %w", d.Get("network").(string), err)
	}

	d.SetId(newTunnelRoute.Network)

	return resourceCloudflareTunnelRouteRead(d, meta)
}

func resourceCloudflareTunnelRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	resource := cloudflare.TunnelRoutesUpdateParams{
		AccountID: d.Get("account_id").(string),
		TunnelID:  d.Get("tunnel_id").(string),
		Network:   d.Get("network").(string),
		Comment:   "",
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	_, err := client.UpdateTunnelRoute(context.Background(), resource)
	if err != nil {
		return fmt.Errorf("error updating Tunnel Route for Network %q: %w", d.Get("network").(string), err)
	}

	return resourceCloudflareTunnelRouteRead(d, meta)
}

func resourceCloudflareTunnelRouteDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	err := client.DeleteTunnelRoute(context.Background(), cloudflare.TunnelRoutesDeleteParams{
		AccountID: d.Get("account_id").(string),
		Network:   d.Get("network").(string),
	})
	if err != nil {
		return fmt.Errorf("error deleting Tunnel Route for Network %q: %w", d.Get("network").(string), err)
	}

	return nil
}

func resourceCloudflareTunnelRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/network"`, d.Id())
	}

	accountID, network := attributes[0], attributes[1]

	d.SetId(network)
	d.Set("account_id", accountID)
	d.Set("network", network)

	err := resourceCloudflareTunnelRouteRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
