package sdkv2provider

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareOriginCACertificate() *schema.Resource {
	return &schema.Resource{
		Schema:        resourceCloudflareOriginCACertificateSchema(),
		CreateContext: resourceCloudflareOriginCACertificateCreate,
		ReadContext:   resourceCloudflareOriginCACertificateRead,
		UpdateContext: resourceCloudflareOriginCACertificateRead,
		DeleteContext: resourceCloudflareOriginCACertificateDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		CustomizeDiff: customdiff.Sequence(
			customdiff.ForceNewIf("expires_on", mustRenew),
		),
		Description: "Provides a Cloudflare Origin CA certificate used to protect traffic to your origin without involving a third party Certificate Authority.",
	}
}

func mustRenew(ctx context.Context, d *schema.ResourceDiff, meta interface{}) bool {
	// Check when the cert will expire
	expiresonRaw := d.Get("expires_on")
	if (expiresonRaw == nil) || (expiresonRaw == "") {
		return false
	}
	expireson, _ := time.Parse(time.RFC3339, expiresonRaw.(string))

	// Calculate when we should renew
	earlyExpiration := expireson.AddDate(0, 0, -1*d.Get("min_days_for_renewal").(int))

	if time.Now().After(earlyExpiration) {
		tflog.Info(ctx, fmt.Sprintf("We will renew the certificate as we passed the expected date (%s)", earlyExpiration))
		err := d.SetNewComputed("expires_on")
		if err != nil {
			tflog.Warn(ctx, fmt.Sprintf("error setting to renew the certificate: %s", err))
			return false
		}
		return true
	}

	return false
}

func resourceCloudflareOriginCACertificateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)

	hostnames := []string{}
	hostnamesRaw := d.Get("hostnames").(*schema.Set)
	for _, h := range hostnamesRaw.List() {
		hostnames = append(hostnames, h.(string))
	}

	certInput := cloudflare.CreateOriginCertificateParams{
		Hostnames:   hostnames,
		RequestType: d.Get("request_type").(string),
	}

	if csr, ok := d.GetOk("csr"); ok {
		certInput.CSR = csr.(string)
	}

	if requestValidity, ok := d.GetOk("requested_validity"); ok {
		certInput.RequestValidity = requestValidity.(int)
	}

	tflog.Info(ctx, fmt.Sprintf("Creating Cloudflare OriginCACertificate: %#v", certInput))
	cert, err := client.CreateOriginCACertificate(ctx, certInput)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error creating origin certificate: %w", err))
	}

	d.SetId(cert.ID)

	return resourceCloudflareOriginCACertificateRead(ctx, d, meta)
}

func resourceCloudflareOriginCACertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	certID := d.Id()
	cert, err := client.GetOriginCACertificate(ctx, certID)

	tflog.Debug(ctx, fmt.Sprintf("OriginCACertificate: %#v", cert))

	if err != nil {
		if strings.Contains(err.Error(), "Failed to read certificate from Database") {
			tflog.Info(ctx, fmt.Sprintf("OriginCACertificate %s does not exist", certID))
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("error finding OriginCACertificate %q: %w", certID, err))
	}

	if cert.RevokedAt != (time.Time{}) {
		tflog.Info(ctx, fmt.Sprintf("OriginCACertificate %s has been revoked", certID))
		d.SetId("")
		return nil
	}

	hostnames := schema.NewSet(schema.HashString, []interface{}{})
	for _, h := range cert.Hostnames {
		hostnames.Add(h)
	}

	d.Set("certificate", cert.Certificate)
	d.Set("expires_on", cert.ExpiresOn.Format(time.RFC3339))
	d.Set("hostnames", hostnames)
	d.Set("request_type", cert.RequestType)

	certBlock, _ := pem.Decode([]byte(cert.Certificate))
	if certBlock == nil {
		return diag.FromErr(fmt.Errorf("error decoding OriginCACertificate %q: %w", certID, err))
	}

	x509Cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error parsing OriginCACertificate %q: %w", certID, err))
	}
	d.Set("requested_validity", calculateRequestedValidityFromCertificate(x509Cert))

	return nil
}

func resourceCloudflareOriginCACertificateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*cloudflare.API)
	certID := d.Id()

	tflog.Info(ctx, fmt.Sprintf("Revoking Cloudflare OriginCACertificate: id %s", certID))

	_, err := client.RevokeOriginCACertificate(ctx, certID)

	if err != nil {
		return diag.FromErr(fmt.Errorf("error revoking Cloudflare OriginCACertificate: %w", err))
	}

	d.SetId("")
	return nil
}

func validateCSR(v interface{}, k string) (ws []string, errors []error) {
	block, _ := pem.Decode([]byte(v.(string)))
	if block == nil {
		errors = append(errors, fmt.Errorf("%q: invalid PEM data", k))
		return
	}

	_, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q: %w", k, err))
	}
	return
}

func calculateRequestedValidityFromCertificate(cert *x509.Certificate) int {
	diff := cert.NotAfter.UTC().Sub(cert.NotBefore.UTC())
	days := math.Round(diff.Hours() / 24)

	validateDays := []float64{7, 30, 90, 365, 730, 1095, 5475}

	// Find the closest matching requested validity (in days) to avoid possible leap second issue.
	i := 0
	d := math.Abs(days - validateDays[i])
	distanceIdx := i
	distance := d
	for i < len(validateDays) {
		d := math.Abs(validateDays[i] - days)
		if d == 0 {
			return int(days)
		}

		if d < distance {
			distanceIdx = i
			distance = d
		}
		i++
	}
	return int(validateDays[distanceIdx])
}
