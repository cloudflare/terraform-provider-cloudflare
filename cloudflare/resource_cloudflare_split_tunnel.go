package cloudflare

import (
	"context"
	"fmt"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareSplitTunnel() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareSplitTunnelSchema(),
		ReadContext:   resourceCloudflareSplitTunnelRead,
		CreateContext: resourceCloudflareSplitTunnelUpdate, // Intentionally identical to Update as the resource is always present
		UpdateContext: resourceCloudflareSplitTunnelUpdate,
		DeleteContext: resourceCloudflareSplitTunnelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareSplitTunnelImport,
		},
	}
}

func resourceCloudflareSplitTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	mode := d.Get("mode").(string)

	splitTunnel, err := client.ListSplitTunnels(ctx, accountID, mode)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error finding %q Split Tunnels: %w", mode, err))
	}

	if err := d.Set("tunnels", flattenSplitTunnels(splitTunnel)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting %q tunnels attribute: %w", mode, err))
	}

	return nil
}

func resourceCloudflareSplitTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	mode := d.Get("mode").(string)

	tunnelList, err := expandSplitTunnels(d.Get("tunnels").([]interface{}))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating %q Split Tunnels: %w", mode, err))
	}

	newSplitTunnels, err := client.UpdateSplitTunnel(ctx, accountID, mode, tunnelList)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating %q Split Tunnels: %w", mode, err))
	}

	if err := d.Set("tunnels", flattenSplitTunnels(newSplitTunnels)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting %q tunnels attribute: %w", mode, err))
	}

	d.SetId(accountID)

	return resourceCloudflareSplitTunnelRead(ctx, d, meta)
}

func resourceCloudflareSplitTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	mode := d.Get("mode").(string)

	_, err := client.UpdateSplitTunnel(ctx, accountID, mode, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func resourceCloudflareSplitTunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/mode\"", d.Id())
	}

	accountID, mode := attributes[0], attributes[1]

	d.Set("mode", mode)
	d.Set("account_id", accountID)
	d.SetId(accountID)

	resourceCloudflareSplitTunnelRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

// flattenSplitTunnels accepts the cloudflare.SplitTunnel struct and returns the
// schema representation for use in Terraform state.
func flattenSplitTunnels(tunnels []cloudflare.SplitTunnel) []interface{} {
	schemaTunnels := make([]interface{}, 0)

	for _, t := range tunnels {
		schemaTunnels = append(schemaTunnels, map[string]interface{}{
			"address":     t.Address,
			"host":        t.Host,
			"description": t.Description,
		})
	}

	return schemaTunnels
}

// expandSplitTunnels accepts the schema representation of Split Tunnels and
// returns a fully qualified struct.
func expandSplitTunnels(tunnels []interface{}) ([]cloudflare.SplitTunnel, error) {
	tunnelList := make([]cloudflare.SplitTunnel, 0)

	for _, tunnel := range tunnels {
		if tunnel.(map[string]interface{})["address"].(string) != "" && tunnel.(map[string]interface{})["host"].(string) != "" {
			return []cloudflare.SplitTunnel{}, errors.New("address and host are mutually exclusive and cannot be applied together in the same block")
		}

		tunnelList = append(tunnelList, cloudflare.SplitTunnel{
			Address:     tunnel.(map[string]interface{})["address"].(string),
			Host:        tunnel.(map[string]interface{})["host"].(string),
			Description: tunnel.(map[string]interface{})["description"].(string),
		})
	}

	return tunnelList, nil
}
