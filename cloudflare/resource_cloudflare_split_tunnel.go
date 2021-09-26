package cloudflare

import (
	"context"
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareSplitTunnel() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCloudflareSplitTunnelRead,
		Create: resourceCloudflareSplitTunnelCreate,
		Update: resourceCloudflareSplitTunnelUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The mode of the split tunnel policy. Either 'include' or 'exclude'.",
				ValidateFunc: validation.StringInSlice([]string{"include", "exclude"}, false),
			},
			"tunnels": {
				Required: true,
				Type:     schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The address for the tunnel.",
						},
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The domain name for the tunnel.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A description for the tunnel.",
						},
					},
				},
			},
		},
	}
}

func resourceCloudflareSplitTunnelRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	mode := d.Get("mode").(string)

	splitTunnel, err := client.ListSplitTunnels(context.Background(), accountID, mode)
	if err != nil {
		return fmt.Errorf("error finding %q Split Tunnels %q", mode, err)
	}

	tunnelList := make([]cloudflare.SplitTunnel, 0)
	for _, t := range splitTunnel {
		tunnelList = append(tunnelList, cloudflare.SplitTunnel{
			Address:     t.Address,
			Host:        t.Host,
			Description: t.Description,
		})
	}

	err = d.Set("tunnels", tunnelList)
	if err != nil {
		return fmt.Errorf("error setting %q tunnels attribute: %q", mode, err)
	}

	return nil
}

func resourceCloudflareSplitTunnelCreate(d *schema.ResourceData, meta interface{}) error {
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

	err := d.Set("tunnels", newTunnel)
	if err != nil {
		return fmt.Errorf("error creating new split tunnel: %q", err)
	}

	return nil
}

func resourceCloudflareSplitTunnelUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	accountID := d.Get("account_id").(string)
	mode := d.Get("mode").(string)

	// get all of the existing split tunnels
	existingTunnels := d.Get("tunnels")
	tunnels := existingTunnels.([]cloudflare.SplitTunnel)
	tunnelList := make([]cloudflare.SplitTunnel, 0)
	tunnelList = append(tunnelList, tunnels...)

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

	tunnelList = append(tunnelList, newTunnel)

	d.SetId(accountID)

	newSplitTunnels, err := client.UpdateSplitTunnel(context.Background(), accountID, mode, tunnelList)
	if err != nil {
		return fmt.Errorf("error updating %q Split Tunnels %q", mode, err)
	}

	d.Set("tunnels", newSplitTunnels)
	if err != nil {
		return fmt.Errorf("error setting %q Split Tunnels: %q", mode, err)
	}

	return nil
}

func resourceCloudflareSplitTunnelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	accountID := d.Get("account_id").(string)
	mode := d.Get("mode").(string)

	d.Set("mode", mode)
	d.Set("account_id", accountID)
	d.SetId(accountID)

	resourceCloudflareSplitTunnelRead(d, meta)

	return []*schema.ResourceData{d}, nil
}
