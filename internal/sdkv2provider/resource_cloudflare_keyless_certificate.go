package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareKeylessCertificate() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareKeylessCertificateSchema(),
		CreateContext: resourceCloudflareKeylessCertificateCreate,
		ReadContext:   resourceCloudflareKeylessCertificateRead,
		UpdateContext: resourceCloudflareKeylessCertificateUpdate,
		DeleteContext: resourceCloudflareKeylessCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareKeylessCertificateImport,
		},
		Description: heredoc.Doc(`
       			Provides a resource, that manages Keyless certificates.
       		`),
	}
}

func resourceCloudflareKeylessCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	request := cloudflare.KeylessSSLCreateRequest{
		Name:         d.Get("name").(string),
		Host:         d.Get("host").(string),
		Port:         d.Get("port").(int),
		Certificate:  d.Get("certificate").(string),
		BundleMethod: d.Get("bundle_method").(string),
	}

	res, err := client.CreateKeylessSSL(ctx, zoneID, request)

	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create Keyless SSL")))
	}

	retry := retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
		keylessSSL, err := client.KeylessSSL(ctx, zoneID, res.ID)
		if err != nil {
			return retry.NonRetryableError(fmt.Errorf("failed to fetch keyless certificate: %w", err))
		}

		if keylessSSL.Status != "active" {
			return retry.RetryableError(fmt.Errorf("waiting for the keyless certificate to become active"))
		}

		d.SetId(res.ID)

		resourceCloudflareKeylessCertificateRead(ctx, d, meta)
		return nil
	})

	if retry != nil {
		return diag.FromErr(retry)
	}

	return nil
}

func resourceCloudflareKeylessCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	keylessSSL, err := client.KeylessSSL(ctx, zoneID, d.Id())
	if err != nil {
		var notFoundError *cloudflare.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Info(ctx, fmt.Sprintf("Keyless SSL %s no longer exists", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding Keyless SSL %q: %w", d.Id(), err))
	}

	d.Set("name", keylessSSL.Name)
	d.Set("host", keylessSSL.Host)
	d.Set("port", keylessSSL.Port)
	d.Set("status", keylessSSL.Status)
	d.Set("enabled", keylessSSL.Enabled)
	d.Set("port", keylessSSL.Port)
	return nil
}

func resourceCloudflareKeylessCertificateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	request := cloudflare.KeylessSSLUpdateRequest{
		Name:    d.Get("name").(string),
		Host:    d.Get("host").(string),
		Port:    d.Get("port").(int),
		Enabled: cloudflare.BoolPtr(d.Get("enabled").(bool)),
	}

	_, err := client.UpdateKeylessSSL(ctx, zoneID, d.Id(), request)
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to update Keyless SSL")))
	}

	return resourceCloudflareKeylessCertificateRead(ctx, d, meta)
}

func resourceCloudflareKeylessCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteKeylessSSL(ctx, zoneID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to delete Keyless SSL")))
	}

	return nil
}

func resourceCloudflareKeylessCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idAttr := strings.SplitN(d.Id(), "/", 2)
	if len(idAttr) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/KeylessSSLID\"", d.Id())
	}

	zoneID, keylessSslId := idAttr[0], idAttr[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Keyless SSL: id %s for zone %s", keylessSslId, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(keylessSslId)

	resourceCloudflareKeylessCertificateRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
