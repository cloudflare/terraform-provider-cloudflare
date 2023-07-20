package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
		Description: heredoc.Doc(`
			Provides a resource, that manages Cloudflare tunnel virtual networks
			for Zero Trust. Tunnel virtual networks are used for segregation of
			Tunnel IP Routes via Virtualized Networks to handle overlapping
			private IPs in your origins.
		`),
	}
}

func resourceCloudflareTunnelVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tunnelVirtualNetworks, err := client.ListTunnelVirtualNetworks(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.TunnelVirtualNetworksListParams{
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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	resource := cloudflare.TunnelVirtualNetworkCreateParams{
		Name:      name,
		IsDefault: d.Get("is_default_network").(bool),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	newTunnelVirtualNetwork, err := client.CreateTunnelVirtualNetwork(ctx, cloudflare.AccountIdentifier(accountID), resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating Tunnel Virtual Network %q: %w", name, err))
	}

	d.SetId(newTunnelVirtualNetwork.ID)

	return resourceCloudflareTunnelVirtualNetworkRead(ctx, d, meta)
}

func resourceCloudflareTunnelVirtualNetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	resource := cloudflare.TunnelVirtualNetworkUpdateParams{
		Name:             d.Get("name").(string),
		IsDefaultNetwork: cloudflare.BoolPtr(d.Get("is_default_network").(bool)),
		VnetID:           d.Id(),
	}

	if comment, ok := d.Get("comment").(string); ok {
		resource.Comment = comment
	}

	_, err := client.UpdateTunnelVirtualNetwork(ctx, cloudflare.AccountIdentifier(accountID), resource)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating Tunnel Virtual Network %q: %w", d.Id(), err))
	}

	return resourceCloudflareTunnelVirtualNetworkRead(ctx, d, meta)
}

func resourceCloudflareTunnelVirtualNetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	err := client.DeleteTunnelVirtualNetwork(ctx, cloudflare.AccountIdentifier(accountID), d.Id())
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
	d.Set(consts.AccountIDSchemaKey, accountID)

	err := resourceCloudflareTunnelVirtualNetworkRead(ctx, d, meta)
	if err != nil {
		return nil, errors.New("failed to read Tunnel Virtual Network state")
	}

	return []*schema.ResourceData{d}, nil
}
