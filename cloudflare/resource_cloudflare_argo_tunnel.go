package cloudflare

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

const argoTunnelCNAME = "cfargotunnel.com"

func resourceCloudflareArgoTunnel() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareArgoTunnelSchema(),
		CreateContext: resourceCloudflareArgoTunnelCreate,
		ReadContext:   resourceCloudflareArgoTunnelRead,
		DeleteContext: resourceCloudflareArgoTunnelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareArgoTunnelImport,
		},
	}
}

func resourceCloudflareArgoTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accID := d.Get("account_id").(string)
	name := d.Get("name").(string)
	secret := d.Get("secret").(string)

	tunnel, err := client.CreateArgoTunnel(ctx, accID, name, secret)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create Argo Tunnel")))
	}

	d.SetId(tunnel.ID)

	return resourceCloudflareArgoTunnelRead(ctx, d, meta)
}

func resourceCloudflareArgoTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accID := d.Get("account_id").(string)

	tunnel, err := client.ArgoTunnel(ctx, accID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Argo Tunnel: %w", err))
	}

	d.Set("cname", fmt.Sprintf("%s.%s", tunnel.ID, argoTunnelCNAME))

	return nil
}

func resourceCloudflareArgoTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accID := d.Get("account_id").(string)

	cleanupErr := client.CleanupArgoTunnelConnections(ctx, accID, d.Id())
	if cleanupErr != nil {
		return diag.FromErr(errors.Wrap(cleanupErr, fmt.Sprintf("failed to clean up Argo Tunnel connections")))
	}

	deleteErr := client.DeleteArgoTunnel(ctx, accID, d.Id())
	if deleteErr != nil {
		return diag.FromErr(errors.Wrap(deleteErr, fmt.Sprintf("failed to delete Argo Tunnel")))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareArgoTunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	attributes := strings.Split(d.Id(), "/")

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/argoTunnelUUID\"", d.Id())
	}

	accID, tunnelID := attributes[0], attributes[1]

	tunnel, err := client.ArgoTunnel(ctx, accID, tunnelID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to fetch Argo Tunnel %s", tunnelID))
	}

	d.Set("account_id", accID)
	d.Set("name", tunnel.Name)
	d.SetId(tunnel.ID)

	resourceCloudflareArgoTunnelRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
