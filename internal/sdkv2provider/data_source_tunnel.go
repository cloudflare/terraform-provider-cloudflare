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

func dataSourceCloudflareTunnel() *schema.Resource {
	return &schema.Resource{
		ReadContext:        dataSourceCloudflareTunnelRead,
		DeprecationMessage: "`cloudflare_tunnel` is now deprecated and will be removed in the next major version. Use `cloudflare_zero_trust_tunnel_cloudflared` instead.",

		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the tunnel.",
				ForceNew:    true,
			},
			"is_deleted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, only include deleted tunnels. If false, exclude deleted tunnels. If empty, all tunnels will be included.",
				ForceNew:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the tunnel.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: fmt.Sprintf("The status of the tunnel. %s", renderAvailableDocumentationValuesStringSlice([]string{"inactive", "degraded", "healthy", "down"})),
			},
			"tunnel_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: fmt.Sprintf("The type of the tunnel. %s", renderAvailableDocumentationValuesStringSlice([]string{"cfd_tunnel", "warp_connector"})),
			},
			"remote_config": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the tunnel can be configured remotely from the Zero Trust dashboard.",
			},
		},
		Description: "Use this datasource to lookup a tunnel in an account.",
	}
}

func dataSourceCloudflareZeroTrustTunnelCloudflared() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareTunnelRead,

		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: consts.AccountIDSchemaDescription,
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the tunnel.",
				ForceNew:    true,
			},
			"is_deleted": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, only include deleted tunnels. If false, exclude deleted tunnels. If empty, all tunnels will be included.",
				ForceNew:    true,
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of the tunnel.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: fmt.Sprintf("The status of the tunnel. %s", renderAvailableDocumentationValuesStringSlice([]string{"inactive", "degraded", "healthy", "down"})),
			},
			"tunnel_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: fmt.Sprintf("The type of the tunnel. %s", renderAvailableDocumentationValuesStringSlice([]string{"cfd_tunnel", "warp_connector"})),
			},
			"remote_config": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the tunnel can be configured remotely from the Zero Trust dashboard.",
			},
		},
		Description: "Use this datasource to lookup a tunnel in an account.",
	}
}

func dataSourceCloudflareTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Reading Tunnel")
	client := meta.(*cloudflare.API)
	accID := d.Get(consts.AccountIDSchemaKey).(string)

	name := d.Get("name").(string)
	params := cloudflare.TunnelListParams{Name: name}

	if v, ok := d.GetOkExists("is_deleted"); ok {
		params.IsDeleted = cloudflare.BoolPtr(v.(bool))
	}

	tunnels, _, err := client.ListTunnels(ctx, cloudflare.AccountIdentifier(accID), params)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Tunnel: %w", err))
	}
	if len(tunnels) == 0 {
		return diag.FromErr(fmt.Errorf("No tunnels with name: %s", name))
	}

	tunnel := tunnels[0]

	d.SetId(tunnel.ID)
	d.Set("status", tunnel.Status)
	d.Set("is_deleted", tunnel.DeletedAt != nil)
	d.Set("id", tunnel.ID)
	d.Set("tunnel_type", tunnel.TunnelType)
	d.Set("remote_config", tunnel.RemoteConfig)
	return nil
}
