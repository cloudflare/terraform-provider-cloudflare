package sdkv2provider

import (
	"context"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCloudflareTunnel() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc(`
			Use this data source to lookup a single [Cloudflare Tunnel](https://developers.cloudflare.com/api/operations/cloudflare-tunnel-get-a-cloudflare-tunnel).
		`),
		ReadContext: dataSourceCloudflareTunnelRead,
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"tunnel_id": {
				Description: "UUID of the tunnel",
				Type:        schema.TypeString,
				Required:    true,
			},
			"status": {
				Description: "Current status of the tunnel. One of: inactive, degraded, healthy, down",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "User-friendly name of the tunnel.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"remote_config": {
				Description: "Whether the tunnel can be configured remotely from the Zero Trust dashboard.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"cname": {
				Description: "Usable CNAME for accessing the tunnel.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"tunnel_token": {
				Description: "Token used by connector to authenticate and run the tunnel.",
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
			},
			"created_at": {
				Description: "Timestamp of when the tunnel was created.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"deleted_at": {
				Description: "Timestamp of when the tunnel was deleted.",
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
			},
		},
	}
}

func dataSourceCloudflareTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	tunnelID := d.Get("tunnel_id").(string)

	tflog.Debug(ctx, fmt.Sprintf("getting tunnel %s", d.Get("tunnel_id").(string)))

	tunnel, err := client.GetTunnel(ctx, cloudflare.AccountIdentifier(accountID), tunnelID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching tunnel %s: %w", tunnelID, err))
	}

	token, err := client.GetTunnelToken(ctx, cloudflare.AccountIdentifier(accountID), tunnel.ID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error fetching tunnel token %s: %w", tunnelID, err))
	}

	tflog.Debug(ctx, fmt.Sprintf("got token %s", token))

	d.SetId(tunnel.ID)
	d.Set("status", tunnel.Status)
	d.Set("name", tunnel.Name)
	d.Set("remote_config", tunnel.RemoteConfig)
	d.Set("created_at", tunnel.CreatedAt.Format(time.RFC3339))

	if tunnel.DeletedAt != nil {
		d.Set("deleted_at", tunnel.DeletedAt.Format(time.RFC3339))
	}

	d.Set("cname", fmt.Sprintf("%s.%s", tunnel.ID, argoTunnelCNAME))
	d.Set("tunnel_token", token)

	return nil
}
