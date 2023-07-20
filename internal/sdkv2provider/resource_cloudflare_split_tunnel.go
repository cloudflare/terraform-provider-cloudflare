package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
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
		Description: heredoc.Doc(`
			Provides a Cloudflare Split Tunnel resource. Split tunnels are used to either
			include or exclude lists of routes from the WARP client's tunnel.
		`),
	}
}

func resourceCloudflareSplitTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	mode := d.Get("mode").(string)
	_, policyID := parseDevicePolicyID(d.Get("policy_id").(string))

	var splitTunnel []cloudflare.SplitTunnel
	var err error
	if policyID == "" {
		splitTunnel, err = client.ListSplitTunnels(ctx, accountID, mode)
	} else {
		splitTunnel, err = client.ListSplitTunnelsDeviceSettingsPolicy(ctx, accountID, policyID, mode)
	}

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
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	mode := d.Get("mode").(string)
	_, policyID := parseDevicePolicyID(d.Get("policy_id").(string))

	tunnelList, err := expandSplitTunnels(d.Get("tunnels").(*schema.Set).List())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating %q Split Tunnels: %w", mode, err))
	}

	var newSplitTunnels []cloudflare.SplitTunnel
	if policyID == "" {
		d.SetId(accountID)
		newSplitTunnels, err = client.UpdateSplitTunnel(ctx, accountID, mode, tunnelList)
	} else {
		d.SetId(fmt.Sprintf("%s/%s", accountID, policyID))
		newSplitTunnels, err = client.UpdateSplitTunnelDeviceSettingsPolicy(ctx, accountID, policyID, mode, tunnelList)
	}
	if err != nil {
		return diag.FromErr(fmt.Errorf("error updating %q Split Tunnels: %w", mode, err))
	}

	if err := d.Set("tunnels", flattenSplitTunnels(newSplitTunnels)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting %q tunnels attribute: %w", mode, err))
	}

	return resourceCloudflareSplitTunnelRead(ctx, d, meta)
}

func resourceCloudflareSplitTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	mode := d.Get("mode").(string)
	_, policyID := parseDevicePolicyID(d.Get("policy_id").(string))

	var err error
	if policyID == "" {
		_, err = client.UpdateSplitTunnel(ctx, accountID, mode, nil)
	} else {
		_, err = client.UpdateSplitTunnelDeviceSettingsPolicy(ctx, accountID, policyID, mode, nil)
	}
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func resourceCloudflareSplitTunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 3)

	if len(attributes) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/policyID/mode\"", d.Id())
	}

	accountID, policyID, mode := attributes[0], attributes[1], attributes[2]

	d.Set(consts.AccountIDSchemaKey, accountID)
	d.Set("mode", mode)
	if policyID == "default" {
		d.Set("policy_id", accountID)
		d.SetId(accountID)
	} else {
		d.Set("policy_id", fmt.Sprintf("%s/%s", accountID, policyID))
		d.SetId(fmt.Sprintf("%s/%s", accountID, policyID))
	}

	resourceCloudflareSplitTunnelRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

// flattenSplitTunnels accepts the cloudflare.SplitTunnel struct and returns the
// schema representation for use in Terraform state.
func flattenSplitTunnels(tunnels []cloudflare.SplitTunnel) *schema.Set {
	schemaTunnels := &schema.Set{F: schema.HashResource(tunnelSetResource)}
	for _, t := range tunnels {
		schemaTunnels.Add(map[string]interface{}{
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
