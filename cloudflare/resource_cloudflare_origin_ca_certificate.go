package cloudflare

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math"
	"strings"
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareOriginCACertificate() *schema.Resource {
	return &schema.Resource{
		Schema: resourceCloudflareOriginCACertificateSchema(),
		Create: resourceCloudflareOriginCACertificateCreate,
		Update: resourceCloudflareOriginCACertificateCreate,
		Read:   resourceCloudflareOriginCACertificateRead,
		Delete: resourceCloudflareOriginCACertificateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceCloudflareOriginCACertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)

	hostnames := []string{}
	hostnamesRaw := d.Get("hostnames").(*schema.Set)
	for _, h := range hostnamesRaw.List() {
		hostnames = append(hostnames, h.(string))
	}

	certInput := cloudflare.OriginCACertificate{
		Hostnames:   hostnames,
		RequestType: d.Get("request_type").(string),
	}

	if csr, ok := d.GetOk("csr"); ok {
		certInput.CSR = csr.(string)
	}

	if requestValidity, ok := d.GetOk("requested_validity"); ok {
		certInput.RequestValidity = requestValidity.(int)
	}

	log.Printf("[INFO] Creating Cloudflare OriginCACertificate: %#v", certInput)
	cert, err := client.CreateOriginCertificate(context.Background(), certInput)

	if err != nil {
		return fmt.Errorf("error creating origin certificate: %s", err)
	}

	d.SetId(cert.ID)

	return resourceCloudflareOriginCACertificateRead(d, meta)
}

func resourceCloudflareOriginCACertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	certID := d.Id()
	cert, err := client.OriginCertificate(context.Background(), certID)

	log.Printf("[DEBUG] OriginCACertificate: %#v", cert)

	if err != nil {
		if strings.Contains(err.Error(), "Failed to read certificate from Database") {
			log.Printf("[INFO] OriginCACertificate %s does not exist", certID)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error finding OriginCACertificate %q: %s", certID, err)
	}

	if cert.RevokedAt != (time.Time{}) {
		log.Printf("[INFO] OriginCACertificate %s has been revoked", certID)
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

	// lazy approach to extracting the date from a known timestamp in order to
	// `time.Parse` it correctly. Here we are getting the certificate expiry and
	// calculating the validity as the API doesn't return it yet it is present in
	// the schema.
	date := strings.Split(cert.ExpiresOn.Format(time.RFC3339), "T")
	certDate, _ := time.Parse("2006-01-02", date[0])
	now := time.Now()
	duration := certDate.Sub(now)
	var validityDays int
	validityDays = int(math.Ceil(duration.Hours() / 24))

	d.Set("requested_validity", validityDays)

	return nil
}

func resourceCloudflareOriginCACertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	certID := d.Id()

	log.Printf("[INFO] Revoking Cloudflare OriginCACertificate: id %s", certID)

	_, err := client.RevokeOriginCertificate(context.Background(), certID)

	if err != nil {
		return fmt.Errorf("error revoking Cloudflare OriginCACertificate: %s", err)
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
		errors = append(errors, fmt.Errorf("%q: %s", k, err.Error()))
	}
	return
}
