package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTeamsRoute() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareTeamsRouteSchema(),
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
		// FIXME(2022-04-21): Until the API returns a valid v4 compatible envelope, we need to
		// check if the error message is related to problems unmarshalling the response _or_
		// an expected not found error.
		var notFoundError *cloudflare.NotFoundError
		if strings.Contains(err.Error(), "error unmarshalling the JSON response error body") || errors.As(err, &notFoundError) {
			d.SetId("")
			return nil
		}

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
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/network"`, d.Id())
	}

	accountID, network := attributes[0], attributes[1]

	d.SetId(network)
	d.Set("account_id", accountID)
	d.Set("network", network)

	err := resourceCloudflareTeamsRouteRead(d, meta)
	if err != nil {
		return nil, err
	}

	return []*schema.ResourceData{d}, nil
}
