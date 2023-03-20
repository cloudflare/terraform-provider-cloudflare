package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCustomHostnameFallbackOrigin() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareCustomHostnameFallbackOriginSchema(),
		CreateContext: resourceCloudflareCustomHostnameFallbackOriginCreate,
		ReadContext:   resourceCloudflareCustomHostnameFallbackOriginRead,
		UpdateContext: resourceCloudflareCustomHostnameFallbackOriginUpdate,
		DeleteContext: resourceCloudflareCustomHostnameFallbackOriginDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareCustomHostnameFallbackOriginImport,
		},
		Description: "Provides a Cloudflare custom hostname fallback origin resource.",
	}
}

func resourceCloudflareCustomHostnameFallbackOriginRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	customHostnameFallbackOrigin, err := client.CustomHostnameFallbackOrigin(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error reading custom hostname fallback origin %q: %w", zoneID, err))
	}

	d.Set("origin", customHostnameFallbackOrigin.Origin)
	d.Set("status", customHostnameFallbackOrigin.Status)

	return nil
}

func resourceCloudflareCustomHostnameFallbackOriginDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteCustomHostnameFallbackOrigin(ctx, zoneID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete custom hostname fallback origin: %w", err))
	}

	return nil
}

func resourceCloudflareCustomHostnameFallbackOriginCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	origin := d.Get("origin").(string)

	fallbackOrigin := cloudflare.CustomHostnameFallbackOrigin{
		Origin: origin,
	}

	retry := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDefault), func() *retry.RetryError {
		_, err := client.UpdateCustomHostnameFallbackOrigin(ctx, zoneID, fallbackOrigin)
		if err != nil {
			var requestError *cloudflare.RequestError
			if errors.As(err, &requestError) && sliceContainsInt(requestError.ErrorCodes(), 1414) {
				return retry.RetryableError(fmt.Errorf("expected custom hostname resource to be ready for modification but is still pending"))
			} else {
				return retry.NonRetryableError(fmt.Errorf("failed to create custom hostname fallback origin: %w", err))
			}
		}

		fallbackHostname, err := client.CustomHostnameFallbackOrigin(ctx, zoneID)

		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("failed to fetch custom hostname: %w", err))
		}

		// Address an eventual consistency issue where deleting a fallback hostname
		// and then adding it _may_ cause some issues. It is possible that the status does
		// move into the active state during the retry period.
		if fallbackHostname.Status != "pending_deployment" && fallbackHostname.Status != "active" {
			return retry.RetryableError(fmt.Errorf("expected custom hostname fallback to be created but was %s", fallbackHostname.Status))
		}

		id := stringChecksum(fmt.Sprintf("%s/custom_hostnames_fallback_origin", zoneID))
		d.SetId(id)

		resourceCloudflareCustomHostnameFallbackOriginRead(ctx, d, meta)
		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareCustomHostnameFallbackOriginUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	origin := d.Get("origin").(string)

	fallbackOrigin := cloudflare.CustomHostnameFallbackOrigin{
		Origin: origin,
	}

	retry := retry.RetryContext(ctx, d.Timeout(schema.TimeoutDefault), func() *retry.RetryError {
		_, err := client.UpdateCustomHostnameFallbackOrigin(ctx, zoneID, fallbackOrigin)
		if err != nil {
			var requestError *cloudflare.RequestError
			if errors.As(err, &requestError) && sliceContainsInt(requestError.ErrorCodes(), 1414) {
				return retry.RetryableError(fmt.Errorf("expected custom hostname resource to be ready for modification but is still pending"))
			}
			return retry.NonRetryableError(fmt.Errorf("failed to update custom hostname fallback origin: %w", err))
		}

		resourceCloudflareCustomHostnameFallbackOriginRead(ctx, d, meta)
		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareCustomHostnameFallbackOriginImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)

	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/origin\"", d.Id())
	}

	zoneID, origin := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Custom Hostname Fallback Origin: origin %s for zone %s", origin, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("origin", origin)

	id := stringChecksum(fmt.Sprintf("%s/custom_hostnames_fallback_origin", zoneID))
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}
