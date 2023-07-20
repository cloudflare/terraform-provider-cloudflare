package sdkv2provider

import (
	"context"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareURLNormalizationSettings() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareURLNormalizationSettingsSchema(),
		CreateContext: resourceCloudflareURLNormalizationSettingsCreate,
		ReadContext:   resourceCloudflareURLNormalizationSettingsRead,
		UpdateContext: resourceCloudflareURLNormalizationSettingsUpdate,
		DeleteContext: resourceCloudflareURLNormalizationSettingsDelete,
		Importer: &schema.ResourceImporter{
			StateContext: nil,
		},
		Description: heredoc.Doc(`
			Provides a resource to manage URL Normalization Settings.
		`),
	}
}

func resourceCloudflareURLNormalizationSettingsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceCloudflareURLNormalizationSettingsUpdate(ctx, d, meta)
}

func resourceCloudflareURLNormalizationSettingsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	settings, err := client.URLNormalizationSettings(ctx, cloudflare.ZoneIdentifier(zoneID))
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("Failed to create URL Normalization Settings for zone: %q", zoneID)))
	}

	d.SetId(zoneID)
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("type", settings.Type)
	d.Set("scope", settings.Scope)

	return nil
}

func resourceCloudflareURLNormalizationSettingsUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	_type := d.Get("type").(string)
	scope := d.Get("scope").(string)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	params := cloudflare.URLNormalizationSettingsUpdateParams{
		Type:  _type,
		Scope: scope,
	}

	_, err := client.UpdateURLNormalizationSettings(ctx, cloudflare.ZoneIdentifier(zoneID), params)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("Failed to update URL Normalization Settings for zone: %q", zoneID)))
	}

	return resourceCloudflareURLNormalizationSettingsRead(ctx, d, meta)
}

func resourceCloudflareURLNormalizationSettingsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.Set("type", "cloudflare")
	d.Set("scope", "incoming")

	return resourceCloudflareURLNormalizationSettingsUpdate(ctx, d, meta)
}
