package provider

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareTotalTLS() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareTotalTLSSchema(),
		CreateContext: resourceCloudflareTotalSSLUpdate,
		ReadContext:   resourceCloudflareTotalSSLRead,
		UpdateContext: resourceCloudflareTotalSSLUpdate,
		DeleteContext: resourceCloudflareTotalSSLDelete,
	}
}

func resourceCloudflareTotalSSLUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	settings := cloudflare.TotalTLS{
		Enabled: cloudflare.BoolPtr(d.Get("enabled").(bool)),
	}
	if certificateAuthority, ok := d.GetOk("certificate_authority"); ok {
		settings.CertificateAuthority = certificateAuthority.(string)
	}
	_, err := client.SetTotalTLS(ctx, cloudflare.ZoneIdentifier(zoneID), settings)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating updating total TLS: %w", err))
	}
	d.SetId(zoneID)
	return resourceCloudflareTotalSSLRead(ctx, d, meta)
}

func resourceCloudflareTotalSSLRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	result, err := client.GetTotalTLS(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating updating total TLS: %w", err))
	}
	d.SetId(zoneID)
	d.Set("enabled", result.Enabled)
	d.Set("certificate_authority", result.CertificateAuthority)
	return nil
}

func resourceCloudflareTotalSSLDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	_, err := client.SetTotalTLS(ctx, cloudflare.ZoneIdentifier(zoneID), cloudflare.TotalTLS{Enabled: cloudflare.BoolPtr(false)})
	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating deleting total TLS: %w", err))
	}

	return nil
}
