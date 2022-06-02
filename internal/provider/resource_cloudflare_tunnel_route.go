package provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTunnelRoute() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTunnelRouteSchema(),
		CreateContext: resourceCloudflareTunnelRouteCreate,
		ReadContext:   resourceCloudflareTunnelRouteRead,
		UpdateContext: resourceCloudflareTunnelRouteUpdate,
		DeleteContext: resourceCloudflareTunnelRouteDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTunnelRouteImport,
		},
	}
}

func resourceCloudflareTunnelRouteRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	network := d.Get("network").(string)

	resource := cloudflare.TunnelRoutesListParams{
		AccountID:       accountID,
		IsDeleted:       cloudflare.BoolPtr(false),
		NetworkSubset:   network,
		NetworkSuperset: network,
	}

	if virtualNetworkID, ok := d.Get("virtual_network_id").(string); ok {
		resource.VirtualNetworkID = virtualNetworkID
	}

	tunnelRoutes, err := client.ListTunnelRoutes(ctx, resource)

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Tunnel Route: %w", err))
	}

	if len(tunnelRoutes) < 1 {
		tflog.Info(ctx, fmt.Sprintf("Tunnel Route for network %s in account %s not found", network, accountID))
		d.SetId("")
		return nil
	}

	tunnelRoute := tunnelRoutes[0]

	d.Set("tunnel_id", tunnelRoute.TunnelID)
	d.Set("network", tunnelRoute.Network)
	if len(tunnelRoute.Comment) > 0 {
		d.Set("comment", tunnelRoute.Comment)
	}

	// vnet is optional. Do not set it unless it was specified explicitly
	if _, ok := d.GetOk("virtual_network_id"); ok {
		d.Set("virtual_network_id", tunnelRoute.VirtualNetworkID)
	}

	return nil
}

func resourceCloudflareTunnelRouteCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	resource := cloudflare.TunnelRoutesCreateParams{
		AccountID: d.Get("account_id").(string),
		TunnelID:  d.Get("tunnel_id").(string),
		Network:   d.Get("network").(string),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	if virtualNetworkID, ok := d.Get("virtual_network_id").(string); ok {
		resource.VirtualNetworkID = virtualNetworkID
	}

	newTunnelRoute, err := client.CreateTunnelRoute(ctx, resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Tunnel Route for Network %q: %w", d.Get("network").(string), err))
	}

	if virtualNetworkID, ok := d.Get("virtual_network_id").(string); ok {
		// It's possible to create several routes with the same network but different virtual network ids.
		d.SetId(fmt.Sprintf("%s/%s", newTunnelRoute.Network, virtualNetworkID))
	} else {
		d.SetId(newTunnelRoute.Network)
	}

	return resourceCloudflareTunnelRouteRead(ctx, d, meta)
}

func resourceCloudflareTunnelRouteUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

	if virtualNetworkID, ok := d.Get("virtual_network_id").(string); ok {
		resource.VirtualNetworkID = virtualNetworkID
	}

	_, err := client.UpdateTunnelRoute(ctx, resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Tunnel Route for Network %q: %w", d.Get("network").(string), err))
	}

	return resourceCloudflareTunnelRouteRead(ctx, d, meta)
}

func resourceCloudflareTunnelRouteDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	resource := cloudflare.TunnelRoutesDeleteParams{
		AccountID: d.Get("account_id").(string),
		Network:   d.Get("network").(string),
	}

	if virtualNetworkID, ok := d.Get("virtual_network_id").(string); ok {
		resource.VirtualNetworkID = virtualNetworkID
	}

	err := client.DeleteTunnelRoute(ctx, resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Tunnel Route for Network %q: %w", d.Get("network").(string), err))
	}

	return nil
}

func resourceCloudflareTunnelRouteImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 2 && len(attributes) != 3 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/network" or "accountID/network/virtual_network_id"`, d.Id())
	}

	accountID, network := attributes[0], attributes[1]

	if len(attributes) == 3 {
		// It's possible to create several routes with the same network but different virtual network ids.
		d.SetId(fmt.Sprintf("%s/%s", network, attributes[2]))
		d.Set("virtual_network_id", accountID)
	} else {
		d.SetId(network)
	}

	d.Set("account_id", accountID)
	d.Set("network", network)

	err := resourceCloudflareTunnelRouteRead(ctx, d, meta)
	if err != nil {
		return nil, errors.New("failed to read Tunnel Route state")
	}

	return []*schema.ResourceData{d}, nil
}
