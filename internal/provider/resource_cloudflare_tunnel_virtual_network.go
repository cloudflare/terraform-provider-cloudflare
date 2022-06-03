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

func resourceCloudflareTunnelVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTunnelVirtualNetworkSchema(),
		CreateContext: resourceCloudflareTunnelVirtualNetworkCreate,
		ReadContext:   resourceCloudflareTunnelVirtualNetworkRead,
		UpdateContext: resourceCloudflareTunnelVirtualNetworkUpdate,
		DeleteContext: resourceCloudflareTunnelVirtualNetworkDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTunnelVirtualNetworkImport,
		},
	}
}

func resourceCloudflareTunnelVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	tunnelVirtualNetworks, err := client.ListTunnelVirtualNetworks(ctx, cloudflare.TunnelVirtualNetworksListParams{
		AccountID: accountID,
		IsDeleted: cloudflare.BoolPtr(false),
		ID:        d.Id(),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Tunnel Virtual Network: %w", err))
	}

	if len(tunnelVirtualNetworks) < 1 {
		tflog.Info(ctx, fmt.Sprintf("Tunnel Virtual Network for ID %s in account %s not found", d.Id(), accountID))
		d.SetId("")
		return nil
	}

	tunnelVirtualNetwork := tunnelVirtualNetworks[0]

	d.Set("name", tunnelVirtualNetwork.Name)
	d.Set("is_default_network", tunnelVirtualNetwork.IsDefaultNetwork)

	if len(tunnelVirtualNetwork.Comment) > 0 {
		d.Set("comment", tunnelVirtualNetwork.Comment)
	}

	return nil
}

func resourceCloudflareTunnelVirtualNetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	name := d.Get("name").(string)

	resource := cloudflare.TunnelVirtualNetworkCreateParams{
		AccountID: d.Get("account_id").(string),
		Name:      name,
		IsDefault: d.Get("is_default_network").(bool),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	newTunnelVirtualNetwork, err := client.CreateTunnelVirtualNetwork(ctx, resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Tunnel Virtual Network %q: %w", name, err))
	}

	d.SetId(newTunnelVirtualNetwork.ID)

	return resourceCloudflareTunnelVirtualNetworkRead(ctx, d, meta)
}

func resourceCloudflareTunnelVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	resource := cloudflare.TunnelVirtualNetworkUpdateParams{
		AccountID:        d.Get("account_id").(string),
		Name:             d.Get("name").(string),
		IsDefaultNetwork: cloudflare.BoolPtr(d.Get("is_default_network").(bool)),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	_, err := client.UpdateTunnelVirtualNetwork(ctx, resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Tunnel Virtual Network %q: %w", d.Id(), err))
	}

	return resourceCloudflareTunnelVirtualNetworkRead(ctx, d, meta)
}

func resourceCloudflareTunnelVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	err := client.DeleteTunnelVirtualNetwork(ctx, cloudflare.TunnelVirtualNetworkDeleteParams{
		AccountID: d.Get("account_id").(string),
		VnetID:    d.Id(),
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting Tunnel Virtual Network %q: %w", d.Id(), err))
	}

	return nil
}

func resourceCloudflareTunnelVirtualNetworkImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf(`invalid id (%q) specified, should be in format "accountID/vnetID"`, d.Id())
	}

	accountID, vnetID := attributes[0], attributes[1]

	d.SetId(vnetID)
	d.Set("account_id", accountID)

	err := resourceCloudflareTunnelVirtualNetworkRead(ctx, d, meta)
	if err != nil {
		return nil, errors.New("failed to read Tunnel Virtual Network state")
	}

	return []*schema.ResourceData{d}, nil
}
