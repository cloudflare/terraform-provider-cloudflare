package sdkv2provider

import (
	"context"
	"fmt"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func datasourceCloudflareTunnel() *schema.Resource {
	return &schema.Resource{
		Description: heredoc.Doc(`
		Use this data source to lookup a single [Cloudflare Tunnel](https://developers.cloudflare.com/api/operations/cloudflare-tunnel-get-a-cloudflare-tunnel).
	`),
		ReadContext: datasourceCloudflareTunnelRead,
		Schema: map[string]*schema.Schema{
			consts.AccountIDSchemaKey: {
				Description: "The account identifier to target for the resource.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"include_token": {
				Description: "Whether to include the tunnel token in the response.",
				Type:        schema.TypeBool,
				Optional:    true,
			},
			"tunnel_id": {
				Description:  "UUID of the tunnel",
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "tunnel_id"},
			},
			"name": {
				Description:  "User-friendly name of the tunnel.",
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"name", "tunnel_id"},
			},
			"status": {
				Description: "Current status of the tunnel. One of: inactive, degraded, healthy, down",
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

func datasourceCloudflareTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accountID := cloudflare.AccountIdentifier(d.Get(consts.AccountIDSchemaKey).(string))

	listParams := cloudflare.TunnelListParams{}
	listParams.UUID = d.Get("tunnel_id").(string)
	listParams.Name = d.Get("name").(string)
	listParams.PerPage = 1

	tunnels, _, err := client.ListTunnels(ctx, accountID, listParams)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(tunnels) == 0 {
		if listParams.UUID != "" {
			return diag.Errorf("no tunnel found with id %q", listParams.UUID)
		} else {
			return diag.Errorf("no tunnel found with name %q", listParams.Name)
		}
	}

	tunnel := tunnels[0]

	d.SetId(tunnel.ID)
	d.Set("name", tunnel.Name)
	d.Set("status", tunnel.Status)
	d.Set("cname", fmt.Sprintf("%s.%s", tunnel.ID, argoTunnelCNAME))
	d.Set("remote_config", tunnel.RemoteConfig)
	d.Set("created_at", tunnel.CreatedAt.Format(time.RFC3339))

	if tunnel.DeletedAt != nil {
		d.Set("deleted_at", tunnel.DeletedAt.Format(time.RFC3339))
	}

	if d.Get("include_token").(bool) {
		token, err := client.GetTunnelToken(ctx, accountID, tunnel.ID)
		if err != nil {
			return diag.FromErr(err)
		}

		d.Set("tunnel_token", token)
	}

	return nil
}
