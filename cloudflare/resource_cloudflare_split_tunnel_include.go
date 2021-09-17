package cloudflare

import (
	"context"
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCloudflareSplitTunnelInclude() *schema.Resource {
	return &schema.Resource{
		Read: resourceCloudflareSplitTunnelIncludeRead,
		Update: resourceCloudflareSplitTunnelIncludeUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnels": {
				Computed: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The address for the tunnel.",
						},
						"host": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "The domain name for the tunnel.",
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "A description for the tunnel.",
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareSplitTunnelIncludeRead(d *schema.ResourceData, meta interface{}) error {
  client := meta.(*cloudflare.API)
  accountID := d.Get("account_id").(string)

  splitTunnel, err := client.SplitTunnelInclude(context.Background(), accountID)
  if err != nil {
    return fmt.Errorf("Error finding Include Split Tunnels %q", err)
  }

  tunnelList := make(map[string]interface{}, 0)
  for _, t := range splitTunnel {
    tunnelList = append(tunnelList, map[string]interface{}{
      "address":      t.Address,
      "host":         t.Host,
      "description":  t.Description,
    })
  }

  err = d.Set("tunnels", tunnelList)
  if err != nil {
    return fmt.Errorf("error setting Include Split Tunnels: %s", err)
  }
  return nil
}

func resourceCloudflareSplitTunnelIncludeUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)

	// get all of the existing split tunnels
  existingSplitTunnel, err := client.SplitTunnelInclude(context.Background(), accountID)
	if err != nil {
		return fmt.Errorf("Error finding Include Split Tunnels %q", err)
	}

	tunnelList := make([]interface{}, 0)
	for _, t := range existingSplitTunnel {
		tun := cloudflare.SplitTunnel{}
		if inputAddress, ok := t.GetOk("address"); ok {
			tun.Address = inputAddress.(string)
		}
		if inputHost, ok := t.GetOk("host"); ok {
			tun.Host = inputHost.(string)
		}
		if inputDescription, ok := t.GetOk("description"); ok {
			tun.Description = inputDescription.(string)
		}

		tunnelList = append(tunnelList, tun)
	}

	// add the new split tunnel
	newTunnel := cloudflare.SplitTunnel{}
	if inputAddress, ok := d.GetOk("address"); ok {
		newTunnel.Address = inputAddress.(string)
	}
	if inputHost, ok := d.GetOk("host"); ok {
		newTunnel.Host = inputHost.(string)
	}
	if inputDescription, ok := d.GetOk("description"); ok {
		newTunnel.Description = inputDescription.(string)
	}

	tunnelList.append(tunnelList, newTunnel)

	err = d.Set("tunnels", tunnelList)
	if err != nil {
		return fmt.Errorf("error setting Include Split Tunnels: %s", err)
	}

	newSplitTunnel, err := client.UpdateSplitTunnelInclude(context.Background(), accountID, tunnelList)
	if err != nil {
		return fmt.Errorf("Error updating Include Split Tunnels %q", err)
	}

	return resourceCloudflareSplitTunnelIncludeRead(d, meta)
}