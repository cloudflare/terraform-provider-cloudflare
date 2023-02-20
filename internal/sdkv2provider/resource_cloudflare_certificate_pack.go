package sdkv2provider

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
)

func resourceCloudflareCertificatePack() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareCertificatePackSchema(),
		CreateContext: resourceCloudflareCertificatePackCreate,
		ReadContext:   resourceCloudflareCertificatePackRead,
		DeleteContext: resourceCloudflareCertificatePackDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceCloudflareCertificatePackImport,
		},
		Description: heredoc.Doc(`
			Provides a Cloudflare Certificate Pack resource that is used to
			provision managed TLS certificates.
		`),
	}
}

func resourceCloudflareCertificatePackCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)
	certificatePackType := d.Get("type").(string)
	certificateHostSet := d.Get("hosts").(*schema.Set)
	certificatePackID := ""

	if certificatePackType == "advanced" {
		validationMethod := d.Get("validation_method").(string)
		validityDays := d.Get("validity_days").(int)
		ca := d.Get("certificate_authority").(string)
		cloudflareBranding := d.Get("cloudflare_branding").(bool)

		cert := cloudflare.CertificatePackRequest{
			Type:                 "advanced",
			Hosts:                expandInterfaceToStringList(certificateHostSet.List()),
			ValidationMethod:     validationMethod,
			ValidityDays:         validityDays,
			CertificateAuthority: ca,
			CloudflareBranding:   cloudflareBranding,
		}
		certPackResponse, err := client.CreateCertificatePack(ctx, zoneID, cert)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create certificate pack: %s", err)))
		}
		certificatePackID = certPackResponse.ID
	} else {
		cert := cloudflare.CertificatePackRequest{
			Type:  certificatePackType,
			Hosts: expandInterfaceToStringList(certificateHostSet.List()),
		}
		certPackResponse, err := client.CreateCertificatePack(ctx, zoneID, cert)
		if err != nil {
			return diag.FromErr(errors.Wrap(err, fmt.Sprintf("failed to create certificate pack: %s", err)))
		}
		certificatePackID = certPackResponse.ID
	}

	if d.Get("wait_for_active_status").(bool) {
		err := resource.RetryContext(ctx, d.Timeout(schema.TimeoutCreate)-time.Minute, func() *resource.RetryError {
			certificatePack, err := client.CertificatePack(ctx, zoneID, certificatePackID)
			if err != nil {
				return resource.NonRetryableError(errors.Wrap(err, "failed to fetch certificate pack"))
			}
			if len(certificatePack.Certificates) == 0 {
				return resource.RetryableError(fmt.Errorf("certificate list in response is empty"))
			}
			for _, certificate := range certificatePack.Certificates {
				if certificate.Status != "active" {
					return resource.RetryableError(fmt.Errorf("expected all certificates in certificate pack to be active state but certificate %s was in state %s", certificate.ID, certificate.Status))
				}
			}
			return nil
		})

		if err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(certificatePackID)

	return resourceCloudflareCertificatePackRead(ctx, d, meta)
}

func resourceCloudflareCertificatePackRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	certificatePack, err := client.CertificatePack(ctx, zoneID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to fetch certificate pack"))
	}

	d.Set("type", certificatePack.Type)
	d.Set("hosts", expandStringListToSet(certificatePack.Hosts))

	if !reflect.ValueOf(certificatePack.ValidationErrors).IsNil() {
		errors := []map[string]interface{}{}
		for _, e := range certificatePack.ValidationErrors {
			errors = append(errors, map[string]interface{}{"message": e.Message})
		}
		d.Set("validation_errors", errors)
	}
	if !reflect.ValueOf(certificatePack.ValidationRecords).IsNil() {
		records := []map[string]interface{}{}
		for _, e := range certificatePack.ValidationRecords {
			records = append(records,
				map[string]interface{}{
					"cname_name":   e.CnameName,
					"cname_target": e.CnameTarget,
					"txt_name":     e.TxtName,
					"txt_value":    e.TxtValue,
					"http_body":    e.HTTPBody,
					"http_url":     e.HTTPUrl,
					"emails":       e.Emails,
				})
		}
		d.Set("validation_records", records)
	}

	return nil
}

func resourceCloudflareCertificatePackDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	zoneID := d.Get(consts.ZoneIDSchemaKey).(string)

	err := client.DeleteCertificatePack(ctx, zoneID, d.Id())
	if err != nil {
		return diag.FromErr(errors.Wrap(err, "failed to delete certificate pack"))
	}

	resourceCloudflareCertificatePackRead(ctx, d, meta)

	return nil
}

func resourceCloudflareCertificatePackImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	attributes := strings.SplitN(d.Id(), "/", 2)

	if len(attributes) != 2 {
		return nil, fmt.Errorf("invalid id (\"%s\") specified, should be in format \"zoneID/certificatePackID\"", d.Id())
	}

	zoneID, certificatePackID := attributes[0], attributes[1]

	tflog.Debug(ctx, fmt.Sprintf("Importing Cloudflare Certificate Pack: id %s for zone %s", certificatePackID, zoneID))

	d.Set(consts.ZoneIDSchemaKey, zoneID)
	d.SetId(certificatePackID)

	resourceCloudflareCertificatePackRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
