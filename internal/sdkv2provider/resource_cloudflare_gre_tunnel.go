package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareGRETunnel() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareGRETunnelSchema(),
		CreateContext: resourceCloudflareGRETunnelCreate,
		ReadContext:   resourceCloudflareGRETunnelRead,
		UpdateContext: resourceCloudflareGRETunnelUpdate,
		DeleteContext: resourceCloudflareGRETunnelDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareGRETunnelImport,
		},
		Description: "Provides a resource, that manages GRE tunnels for Magic Transit.",
	}
}

func resourceCloudflareGRETunnelCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	newTunnel, err := client.CreateMagicTransitGRETunnels(ctx, accountID, []cloudflare.MagicTransitGRETunnel{
		GRETunnelFromResource(d),
	})

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating GRE tunnel %s: %w", d.Get("name").(string), err))
	}

	d.SetId(newTunnel[0].ID)

	return resourceCloudflareGRETunnelRead(ctx, d, meta)
}

func resourceCloudflareGRETunnelImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, errors.New(fmt.Sprintf("invalid id (\"%s\") specified, should be in format \"accountID/tunnelID\"", d.Id()))
	}

	accountID, tunnelID := attributes[0], attributes[1]
	d.SetId(tunnelID)
	d.Set(consts.AccountIDSchemaKey, accountID)

	readErr := resourceCloudflareGRETunnelRead(ctx, d, meta)
	if readErr != nil {
		return nil, errors.New("failed to read GRE Tunnel state")
	}

	return []*schema.ResourceData{d}, nil
}

func resourceCloudflareGRETunnelRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	tunnel, err := client.GetMagicTransitGRETunnel(ctx, accountID, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "GRE tunnel not found") {
			tflog.Info(ctx, fmt.Sprintf("GRE tunnel %s not found", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error reading GRE tunnel ID %q: %w", d.Id(), err))
	}

	d.Set("name", tunnel.Name)
	d.Set("customer_gre_endpoint", tunnel.CustomerGREEndpoint)
	d.Set("cloudflare_gre_endpoint", tunnel.CloudflareGREEndpoint)
	d.Set("interface_address", tunnel.InterfaceAddress)
	d.Set("ttl", int(tunnel.TTL))
	d.Set("mtu", int(tunnel.MTU))
	d.Set("health_check_enabled", tunnel.HealthCheck.Enabled)
	d.Set("health_check_target", tunnel.HealthCheck.Target)
	d.Set("health_check_type", tunnel.HealthCheck.Type)

	if len(tunnel.Description) > 0 {
		d.Set("description", tunnel.Description)
	}

	return nil
}

func resourceCloudflareGRETunnelUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	_, err := client.UpdateMagicTransitGRETunnel(ctx, accountID, d.Id(), GRETunnelFromResource(d))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("error updating GRE tunnel %q", d.Id())))
	}

	return resourceCloudflareGRETunnelRead(ctx, d, meta)
}

func resourceCloudflareGRETunnelDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	accountID := d.Get(consts.AccountIDSchemaKey).(string)
	client := meta.(*cloudflare.API)

	tflog.Info(ctx, fmt.Sprintf("Deleting GRE tunnel:  %s", d.Id()))

	_, err := client.DeleteMagicTransitGRETunnel(ctx, accountID, d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("error deleting GRE tunnel: %w", err))
	}

	return nil
}

func GRETunnelFromResource(d *schema.ResourceData) cloudflare.MagicTransitGRETunnel {
	tunnel := cloudflare.MagicTransitGRETunnel{
		Name:                  d.Get("name").(string),
		CustomerGREEndpoint:   d.Get("customer_gre_endpoint").(string),
		CloudflareGREEndpoint: d.Get("cloudflare_gre_endpoint").(string),
		InterfaceAddress:      d.Get("interface_address").(string),
	}

	description, descriptionOk := d.GetOk("description")
	if descriptionOk {
		tunnel.Description = description.(string)
	}

	ttl, ttlOk := d.GetOk("ttl")
	if ttlOk {
		tunnel.TTL = uint8(ttl.(int))
	}

	mtu, mtuOk := d.GetOk("mtu")
	if mtuOk {
		tunnel.MTU = uint16(mtu.(int))
	}

	healthcheck := GRETunnelHealthcheckFromResource(d)
	if healthcheck != nil {
		tunnel.HealthCheck = healthcheck
	}

	return tunnel
}

func GRETunnelHealthcheckFromResource(d *schema.ResourceData) *cloudflare.MagicTransitGRETunnelHealthcheck {
	healthcheck := cloudflare.MagicTransitGRETunnelHealthcheck{}

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

	if healthcheckEnabledOk || healthcheckTargetOk || healthcheckTypeOk {
		return &healthcheck
	}

	return nil
}
