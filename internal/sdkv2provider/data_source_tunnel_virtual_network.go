package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareTunnelVirtualNetwork() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareTunnelVirtualNetworkRead,

		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				Description: "The Virtual Network Name.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"comment": {
				Description: "The Virtual Network Comment.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_default": {
				Description: "If true, only include deleted virtual networks. If false, exclude deleted virtual networks. If empty, all virtual networks will be included.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
		Description: "Use this datasource to lookup a tunnel virtual network in an account.",
	}
}

func dataSourceCloudflareTunnelVirtualNetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)

	tflog.Debug(ctx, "reading virtual networks")

	tvn, err := client.ListTunnelVirtualNetworks(ctx, cloudflare.AccountIdentifier(accountID), cloudflare.TunnelVirtualNetworksListParams{
		Name: d.Get("name").(string),
	})
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Tunnel Virtual Networks: %w", err))
	}
	if len(tvn) == 0 {
		return diag.FromErr(fmt.Errorf("No Tunnel Virtual Networks found with name: %s", d.Get("name").(string)))
	}

	network := tvn[0]

	d.SetId(network.ID)
	d.Set("name", network.Name)
	d.Set("comment", network.Comment)
	d.Set("is_default", network.IsDefaultNetwork)
	return nil
}
