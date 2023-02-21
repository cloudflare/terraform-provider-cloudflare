package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

const argoTunnelCNAME = "cfargotunnel.com"

func resourceCloudflareTunnel() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTunnelSchema(),
		CreateContext: resourceCloudflareTunnelCreate,
		ReadContext:   resourceCloudflareTunnelRead,
		DeleteContext: resourceCloudflareTunnelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareTunnelImport,
		},
		Description: heredoc.Doc(`
			Tunnel exposes applications running on your local web server on any
			network with an internet connection without manually adding DNS
			records or configuring a firewall or router.
		`),
	}
}

func resourceCloudflareTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accID := d.Get(consts.AccountIDSchemaKey).(string)
	name := d.Get("name").(string)
	secret := d.Get("secret").(string)

	tunnel, err := client.CreateTunnel(ctx, cloudflare.AccountIdentifier(accID), cloudflare.TunnelCreateParams{Name: name, Secret: secret})
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create Argo Tunnel")))
	}

	d.SetId(tunnel.ID)

	return resourceCloudflareTunnelRead(ctx, d, meta)
}

func resourceCloudflareTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accID := d.Get(consts.AccountIDSchemaKey).(string)

	tunnel, err := client.Tunnel(ctx, cloudflare.AccountIdentifier(accID), d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Argo Tunnel: %w", err))
	}

	token, err := client.TunnelToken(ctx, cloudflare.AccountIdentifier(accID), tunnel.ID)

	if err != nil {
		tflog.Warn(ctx, "unable to set the tunnel_token in state because it's not found in API")
		d.Set("tunnel_token", "")
		return nil
	}

	d.Set("cname", fmt.Sprintf("%s.%s", tunnel.ID, argoTunnelCNAME))
	d.Set("tunnel_token", token)

	return nil
}

func resourceCloudflareTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	accID := d.Get(consts.AccountIDSchemaKey).(string)

	cleanupErr := client.CleanupTunnelConnections(ctx, cloudflare.AccountIdentifier(accID), d.Id())
	if cleanupErr != nil {
		return diag.FromErr(errors.Wrap(cleanupErr, fmt.Sprintf("failed to clean up Argo Tunnel connections")))
	}

	deleteErr := client.DeleteTunnel(ctx, cloudflare.AccountIdentifier(accID), d.Id())
	if deleteErr != nil {
		return diag.FromErr(errors.Wrap(deleteErr, fmt.Sprintf("failed to delete Argo Tunnel")))
	}

	d.SetId("")

	return nil
}

func resourceCloudflareTunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client := meta.(*cloudflare.API)
	attributes := strings.Split(d.Id(), "/")

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"accountID/argoTunnelUUID\"", d.Id())
	}

	accID, tunnelID := attributes[0], attributes[1]

	tunnel, err := client.Tunnel(ctx, cloudflare.AccountIdentifier(accID), tunnelID)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to fetch Argo Tunnel %s", tunnelID))
	}

	d.Set(consts.AccountIDSchemaKey, accID)
	d.Set("name", tunnel.Name)
	d.SetId(tunnel.ID)

	resourceCloudflareTunnelRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
