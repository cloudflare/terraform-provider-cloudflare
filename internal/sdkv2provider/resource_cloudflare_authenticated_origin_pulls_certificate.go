package sdkv2provider

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareAuthenticatedOriginPullsCertificate() *schema.Resource {
	return &schema.Resource{
		// You cannot edit AOP certificates, rather, only upload new ones.
		CreateContext: resourceCloudflareAuthenticatedOriginPullsCertificateCreate,
		ReadContext:   resourceCloudflareAuthenticatedOriginPullsCertificateRead,
		DeleteContext: resourceCloudflareAuthenticatedOriginPullsCertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareAuthenticatedOriginPullsCertificateImport,
		},

		Schema: resourceCloudflareAuthenticatedOriginPullsCertificateSchema(),
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Authenticated Origin Pulls certificate
			resource. An uploaded client certificate is required to use Per-Zone
			 or Per-Hostname Authenticated Origin Pulls.
		`),
	}
}

func resourceCloudflareAuthenticatedOriginPullsCertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		perZoneAOPCert := cloudflare.PerZoneAuthenticatedOriginPullsCertificateParams{
			Certificate: d.Get("certificate").(string),
			PrivateKey:  d.Get("private_key").(string),
		}
		record, err := client.UploadPerZoneAuthenticatedOriginPullsCertificate(ctx, zoneID, perZoneAOPCert)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error uploading Per-Zone AOP certificate on zone %q: %w", zoneID, err))
		}
		d.SetId(record.ID)

		perZoneRetryErr := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			resp, err := client.GetPerZoneAuthenticatedOriginPullsCertificateDetails(ctx, zoneID, record.ID)
			if err != nil {
				return resource.NonRetryableError(fmt.Errorf("error reading Per Zone AOP certificate details: %w", err))
			}

			if resp.Status != "active" {
				return resource.RetryableError(fmt.Errorf("expected Per Zone AOP certificate to be active but was in state %s", resp.Status))
			}

			resourceCloudflareAuthenticatedOriginPullsCertificateRead(ctx, d, meta)
			return nil
		})

		if perZoneRetryErr != nil {
			return diag.FromErr(perZoneRetryErr)
		}

		return nil

	case aopType == "per-hostname":
		perHostnameAOPCert := cloudflare.PerHostnameAuthenticatedOriginPullsCertificateParams{
			Certificate: d.Get("certificate").(string),
			PrivateKey:  d.Get("private_key").(string),
		}
		record, err := client.UploadPerHostnameAuthenticatedOriginPullsCertificate(ctx, zoneID, perHostnameAOPCert)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error uploading Per-Hostname AOP certificate on zone %q: %w", zoneID, err))
		}
		d.SetId(record.ID)

		perHostnameRetryErr := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			resp, err := client.GetPerHostnameAuthenticatedOriginPullsCertificate(ctx, zoneID, record.ID)
			if err != nil {
				return resource.NonRetryableError(fmt.Errorf("error reading Per Hostname AOP certificate details: %w", err))
			}

			if resp.Status != "active" {
				return resource.RetryableError(fmt.Errorf("expected Per Hostname AOP certificate to be active but was in state %s", resp.Status))
			}

			resourceCloudflareAuthenticatedOriginPullsCertificateRead(ctx, d, meta)
			return nil
		})

		if perHostnameRetryErr != nil {
			return diag.FromErr(perHostnameRetryErr)
		}

		return nil
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	certID := d.Id()

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		record, err := client.GetPerZoneAuthenticatedOriginPullsCertificateDetails(ctx, zoneID, certID)
		if err != nil {
			var notFoundError *cloudflare.NotFoundError
			if errors.As(err, &notFoundError) {
				tflog.Info(ctx, fmt.Sprintf("Per-Zone Authenticated Origin Pull certificate %s no longer exists", d.Id()))
				d.SetId("")
				return nil
			}
			return diag.FromErr(fmt.Errorf("error finding Per-Zone Authenticated Origin Pull certificate %q: %w", d.Id(), err))
		}
		d.Set("issuer", record.Issuer)
		d.Set("signature", record.Signature)
		d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))
		d.Set("status", record.Status)
		d.Set("uploaded_on", record.UploadedOn.Format(time.RFC3339Nano))
	case aopType == "per-hostname":
		record, err := client.GetPerHostnameAuthenticatedOriginPullsCertificate(ctx, zoneID, certID)
		if err != nil {
			var notFoundError *cloudflare.NotFoundError
			if errors.As(err, &notFoundError) {
				tflog.Info(ctx, fmt.Sprintf("Per-Hostname Authenticated Origin Pull certificate %s no longer exists", d.Id()))
				d.SetId("")
				return nil
			}
			return diag.FromErr(fmt.Errorf("error finding Per-Hostname Authenticated Origin Pull certificate %q: %w", d.Id(), err))
		}
		d.Set("issuer", record.Issuer)
		d.Set("signature", record.Signature)
		d.Set("serial_number", record.SerialNumber)
		d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))
		d.Set("status", record.Status)
		d.Set("uploaded_on", record.UploadedOn.Format(time.RFC3339Nano))
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsCertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	certID := d.Id()

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(ctx, zoneID, certID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error deleting Per-Zone AOP certificate on zone %q: %w", zoneID, err))
		}
	case aopType == "per-hostname":
		_, err := client.DeletePerHostnameAuthenticatedOriginPullsCertificate(ctx, zoneID, certID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error deleting Per-Hostname AOP certificate on zone %q: %w", zoneID, err))
		}
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsCertificateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// split the id so we can lookup
	idAttr := strings.SplitN(d.Id(), "/", 3)

	if len(idAttr) != 3 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/type/certID\"", d.Id())
	}
	zoneID, aopType, certID := idAttr[0], idAttr[1], idAttr[2]
	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.Set("type", aopType)
	d.SetId(certID)

	resourceCloudflareAuthenticatedOriginPullsCertificateRead(ctx, d, meta)
	return []*schema.ResourceData{d}, nil
}
