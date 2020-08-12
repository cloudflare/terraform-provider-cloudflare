package cloudflare

import (
	"fmt"
	"log"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCloudflareAuthenticatedOriginPullsCertificate() *schema.Resource {
	return &schema.Resource{
		// You cannot edit AOP certificates, rather, only upload new ones.
		Create: resourceCloudflareAuthenticatedOriginPullsCertificateCreate,
		Read:   resourceCloudflareAuthenticatedOriginPullsCertificateRead,
		Delete: resourceCloudflareAuthenticatedOriginPullsCertificateDelete,
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"certificate": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				ForceNew:  true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"expires_on": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"uploaded_on": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"per-zone", "per-hostname"}, false),
				Required:     true,
				ForceNew:     true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
		},
	}
}

func resourceCloudflareAuthenticatedOriginPullsCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		perZoneAOPCert := cloudflare.PerZoneAuthenticatedOriginPullsCertificateParams{
			Certificate: d.Get("certificate").(string),
			PrivateKey:  d.Get("private_key").(string),
		}
		record, err := client.UploadPerZoneAuthenticatedOriginPullsCertificate(zoneID, perZoneAOPCert)
		if err != nil {
			return fmt.Errorf("Error uploading Per-Zone AOP certificate on zone %q: %s", zoneID, err)
		}
		d.SetId(record.ID)

		return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			resp, err := client.GetPerZoneAuthenticatedOriginPullsCertificateDetails(zoneID, record.ID)
			if err != nil {
				return resource.NonRetryableError(fmt.Errorf("Error reading Per Zone AOP certificate details: %s", err))
			}

			if resp.Status != "active" {
				return resource.RetryableError(fmt.Errorf("Expected Per Zone AOP certificate to be active but was in state %s", resp.Status))
			}

			return resource.NonRetryableError(resourceCloudflareAuthenticatedOriginPullsCertificateRead(d, meta))
		})
	case aopType == "per-hostname":
		perHostnameAOPCert := cloudflare.PerHostnameAuthenticatedOriginPullsCertificateParams{
			Certificate: d.Get("certificate").(string),
			PrivateKey:  d.Get("private_key").(string),
		}
		record, err := client.UploadPerHostnameAuthenticatedOriginPullsCertificate(zoneID, perHostnameAOPCert)
		if err != nil {
			return fmt.Errorf("Error uploading Per-Hostname AOP certificate on zone %q: %s", zoneID, err)
		}
		d.SetId(record.ID)

		return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
			resp, err := client.GetPerHostnameAuthenticatedOriginPullsCertificate(zoneID, record.ID)
			if err != nil {
				return resource.NonRetryableError(fmt.Errorf("Error reading Per Hostname AOP certificate details: %s", err))
			}

			if resp.Status != "active" {
				return resource.RetryableError(fmt.Errorf("Expected Per Hostname AOP certificate to be active but was in state %s", resp.Status))
			}

			return resource.NonRetryableError(resourceCloudflareAuthenticatedOriginPullsCertificateRead(d, meta))
		})
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		record, err := client.GetPerZoneAuthenticatedOriginPullsCertificateDetails(zoneID, certID)
		if err != nil {
			log.Printf("[WARN] Removing certificate from state because it was not found in API.")
			d.SetId("")
			return nil
		}
		d.Set("issuer", record.Issuer)
		d.Set("signature", record.Signature)
		d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))
		d.Set("status", record.Status)
		d.Set("uploaded_on", record.UploadedOn)
	case aopType == "per-hostname":
		record, err := client.GetPerHostnameAuthenticatedOriginPullsCertificate(zoneID, certID)
		if err != nil {
			log.Printf("[WARN] Removing certificate from state because it was not found in API.")
			d.SetId("")
			return nil
		}
		d.Set("issuer", record.Issuer)
		d.Set("signature", record.Signature)
		d.Set("serial_number", record.SerialNumber)
		d.Set("expires_on", record.ExpiresOn.Format(time.RFC3339Nano))
		d.Set("status", record.Status)
		d.Set("uploaded_on", record.UploadedOn)
	}
	return nil
}

func resourceCloudflareAuthenticatedOriginPullsCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(zoneID, certID)
		if err != nil {
			return fmt.Errorf("Error deleting Per-Zone AOP certificate on zone %q: %s", zoneID, err)
		}
	case aopType == "per-hostname":
		_, err := client.DeletePerHostnameAuthenticatedOriginPullsCertificate(zoneID, certID)
		if err != nil {
			return fmt.Errorf("Error deleting Per-Hostname AOP certificate on zone %q: %s", zoneID, err)
		}
	}
	return nil
}
