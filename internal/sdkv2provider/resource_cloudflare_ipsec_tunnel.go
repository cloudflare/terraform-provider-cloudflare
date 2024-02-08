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

func resourceCloudflareIPsecTunnel() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareIPsecTunnelSchema(),
		CreateContext: resourceCloudflareIPsecTunnelCreate,
		ReadContext:   resourceCloudflareIPsecTunnelRead,
		UpdateContext: resourceCloudflareIPsecTunnelUpdate,
		DeleteContext: resourceCloudflareIPsecTunnelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareIPsecTunnelImport,
		},
		Description: heredoc.Doc(`
			Provides a resource, that manages IPsec tunnels for Magic Transit.
		`),
	}
}

func resourceCloudflareIPsecTunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	newTunnel, err := client.CreateMagicTransitIPsecTunnels(ctx, accountID, []cloudflare.MagicTransitIPsecTunnel{
		IPsecTunnelFromResource(d),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating IPSec tunnel %s: %w", d.Get("name").(string), err))
	}

	d.SetId(newTunnel[0].ID)

	// If PSK is not specified, call generate PSK and populate the field
	psk, pskOk := d.Get("psk").(string)
	if !pskOk || psk == "" {
		psk, _, err = client.GenerateMagicTransitIPsecTunnelPSK(ctx, accountID, d.Id())
		if err != nil {
			defer d.SetId("")
			tflog.Error(ctx, fmt.Sprintf("error creating PSK: %s %s", accountID, d.Id()))
			// Need to delete the tunnel
			return resourceCloudflareIPsecTunnelDelete(ctx, d, meta)
		}
		d.Set("psk", psk)
	}

	return resourceCloudflareIPsecTunnelRead(ctx, d, meta)
}

func resourceCloudflareIPsecTunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid id (\"%s\") specified, should be in format \"accountID/tunnelID\"", d.Id()))
	}

	accountID, tunnelID := attributes[0], attributes[1]
	d.SetId(tunnelID)
	d.Set(consts.AccountIDSchemaKey, accountID)

	readDiags := resourceCloudflareIPsecTunnelRead(ctx, d, meta)
	if readDiags != nil {
		return nil, errors.New("failed to read IPSec Tunnel state")
	}

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareIPsecTunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	tunnel, err := client.GetMagicTransitIPsecTunnel(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "IPsec tunnel not found") {
			tflog.Info(ctx, fmt.Sprintf("IPsec tunnel %s not found", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading IPsec tunnel ID %q: %w", d.Id(), err))
	}

	d.Set("name", tunnel.Name)
	d.Set("customer_endpoint", tunnel.CustomerEndpoint)
	d.Set("cloudflare_endpoint", tunnel.CloudflareEndpoint)
	d.Set("interface_address", tunnel.InterfaceAddress)
	d.Set("health_check_enabled", tunnel.HealthCheck.Enabled)
	d.Set("health_check_target", tunnel.HealthCheck.Target)
	d.Set("health_check_type", tunnel.HealthCheck.Type)
	d.Set("health_check_direction", tunnel.HealthCheck.Direction)
	d.Set("health_check_rate", tunnel.HealthCheck.Rate)
	d.Set("allow_null_cipher", tunnel.AllowNullCipher)

	// Set Remote Identities
	d.Set("hex_id", tunnel.RemoteIdentities.HexID)
	d.Set("fqdn_id", tunnel.RemoteIdentities.FQDNID)
	d.Set("user_id", tunnel.RemoteIdentities.UserID)
	d.Set("remote_id", accountID+"_"+d.Id())

	if len(tunnel.Description) > 0 {
		d.Set("description", tunnel.Description)
	}

	return nil
}

func resourceCloudflareIPsecTunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)
	_, err := client.UpdateMagicTransitIPsecTunnel(ctx, accountID, d.Id(), IPsecTunnelFromResource(d))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error updating IPsec tunnel %q", d.Id())))
	}

	// Note: PSK field is expected to be populated during create. The only reason
	// it can be empty is when the resource wants to regenerate it.
	psk, pskOk := d.Get("psk").(string)
	if !pskOk || psk == "" {
		psk, _, err = client.GenerateMagicTransitIPsecTunnelPSK(ctx, accountID, d.Id())
		if err != nil {
			// Return Update PSK generation failed
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error regenerating PSK: %s %s", accountID, d.Id())))
		} else {
			d.Set("psk", psk)
		}
	}

	return resourceCloudflareIPsecTunnelRead(ctx, d, meta)
}

func resourceCloudflareIPsecTunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	tflog.Info(ctx, fmt.Sprintf("Deleting IPsec tunnel:  %s", d.Id()))

	_, err := client.DeleteMagicTransitIPsecTunnel(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting IPsec tunnel: %w", err))
	}

	return nil
}

func IPsecTunnelFromResource(d *schema.ResourceData) cloudflare.MagicTransitIPsecTunnel {
	tunnel := cloudflare.MagicTransitIPsecTunnel{
		Name:               d.Get("name").(string),
		CustomerEndpoint:   d.Get("customer_endpoint").(string),
		CloudflareEndpoint: d.Get("cloudflare_endpoint").(string),
		InterfaceAddress:   d.Get("interface_address").(string),
	}

	description, descriptionOk := d.GetOk("description")
	if descriptionOk {
		tunnel.Description = description.(string)
	}

	psk, pskOk := d.GetOk("psk")
	if pskOk {
		tunnel.Psk = psk.(string)
	}

	allowNullCipher, allowNullCipherOk := d.GetOk("allow_null_cipher")
	if allowNullCipherOk {
		tunnel.AllowNullCipher = allowNullCipher.(bool)
	}

	healthcheck := IPsecTunnelHealthcheckFromResource(d)
	if healthcheck != nil {
		tunnel.HealthCheck = healthcheck
	}

	return tunnel
}

func IPsecTunnelHealthcheckFromResource(d *schema.ResourceData) *cloudflare.MagicTransitTunnelHealthcheck {
	healthcheck := cloudflare.MagicTransitTunnelHealthcheck{}

	healthcheckEnabled, healthcheckEnabledOk := d.GetOk("health_check_enabled")
	if healthcheckEnabledOk {
		healthcheck.Enabled = healthcheckEnabled.(bool)
	}

	healthcheckTarget, healthcheckTargetOk := d.GetOk("health_check_target")
	if healthcheckTargetOk {
		healthcheck.Target = healthcheckTarget.(string)
	}

	healthcheckType, healthcheckTypeOk := d.GetOk("health_check_type")
	if healthcheckTypeOk {
		healthcheck.Type = healthcheckType.(string)
	}

	healthcheckDirection, healthcheckDirectionOk := d.GetOk("health_check_direction")
	if healthcheckDirectionOk {
		healthcheck.Direction = healthcheckDirection.(string)
	}

	healthcheckRate, healthcheckRateOk := d.GetOk("health_check_rate")
	if healthcheckRateOk {
		healthcheck.Rate = healthcheckRate.(string)
	}

	if healthcheckEnabledOk || healthcheckTargetOk || healthcheckTypeOk || healthcheckDirectionOk || healthcheckRateOk {
		return &healthcheck
	}

	return nil
}
