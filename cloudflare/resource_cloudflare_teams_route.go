package cloudflare

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsRoute() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Create: resourceCloudflareTeamsRouteCreate,
		Read:   resourceCloudflareTeamsRouteRead,
		Update: resourceCloudflareTeamsRouteUpdate,
		Delete: resourceCloudflareTeamsRouteDelete,
		Importer: &schema.ResourceImporter{
			State: resourceCloudflareTeamsRouteImport,
		},
	}
}

func resourceCloudflareTeamsRouteRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	tunnelRoute, err := client.GetTunnelRouteForIP(context.Background(), cloudflare.TunnelRoutesForIPParams{
		AccountID: d.Get("account_id").(string),
		Network:   d.Get("network").(string),
	})
	if err != nil {
		// TODO: implement logic for missing route

		return fmt.Errorf("error reading Tunnel Route for Network %q: %w", d.Id(), err)
	}

	d.Set("tunnel_id", tunnelRoute.TunnelID)
	d.Set("network", tunnelRoute.Network)

	if len(tunnelRoute.Comment) > 0 {
		d.Set("comment", tunnelRoute.Comment)
	}

	return nil
}

func resourceCloudflareTeamsRouteCreate(d *schema.ResourceData, meta interface{}) error {
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

	return resourceCloudflareTeamsRouteRead(d, meta)
}

func resourceCloudflareTeamsRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	resource := cloudflare.TunnelRoutesUpdateParams{
		AccountID: d.Get("account_id").(string),
		TunnelID:  d.Get("tunnel_id").(string),
		Network:   d.Get("network").(string),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	_, err := client.UpdateTunnelRoute(context.Background(), resource)
	if err != nil {
		return fmt.Errorf("error updating Tunnel Route for Network %q: %w", d.Get("network").(string), err)
	}

	return resourceCloudflareTeamsRouteRead(d, meta)
}

func resourceCloudflareTeamsRouteDelete(d *schema.ResourceData, meta interface{}) error {
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

func resourceCloudflareTeamsRouteImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/tunnelID/network"`, d.Id())
	}

	accountID, tunnelID, network := attributes[0], attributes[1], attributes[2]

	d.SetId(network)
	d.Set("account_id", accountID)
	d.Set("tunnel_id", tunnelID)
	d.Set("network", network)

	err := resourceCloudflareTeamsRouteRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
