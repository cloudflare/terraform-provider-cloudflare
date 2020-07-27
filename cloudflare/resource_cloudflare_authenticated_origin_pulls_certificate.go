package cloudflare

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/pkg/errors"
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
			"type": {
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice([]string{"per-zone", "per-hostname"}, false),
				Required:     true,
				ForceNew:     true,
			},
		},
	}
}

func resourceCloudflareAuthenticatedOriginPullsCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		perZoneAOPCert := cloudflare.PerZoneAuthenticatedOriginPullsCertificateParams{
			Certificate: d.Get("certificate").(string),
			PrivateKey:  d.Get("private_key").(string),
		}
		record, err := client.UploadPerZoneAuthenticatedOriginPullsCertificate(zoneID, perZoneAOPCert)
		if err != nil {
			return fmt.Errorf("error creating AOP cert for zone %q: %s", zoneID, err)
		}
		d.SetId(record.ID)
		return resourceCloudflareAuthenticatedOriginPullsCertificateRead(d, meta)
	case aopType == "per-hostname":
		perHostnameAOPCert := cloudflare.PerHostnameAuthenticatedOriginPullsCertificateParams{
			Certificate: d.Get("certificate").(string),
			PrivateKey:  d.Get("private_key").(string),
		}
		record, err := client.UploadPerHostnameAuthenticatedOriginPullsCertificate(zoneID, perHostnameAOPCert)
		if err != nil {
			return fmt.Errorf("error uploading per zone AOP cert")
		}
		d.SetId(record.ID)
		return resourceCloudflareAuthenticatedOriginPullsCertificateRead(d, meta)
	default:
		return errors.New("unknown error")
	}
}

func resourceCloudflareAuthenticatedOriginPullsCertificateRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		record, err := client.GetPerZoneAuthenticatedOriginPullsCertificateDetails(zoneID, certID)
		if err != nil {
			log.Printf("[WARN] Removing record from state because it's not found in API")
			d.SetId("")
			return nil
		}
		d.Set("certificate", record.Certificate)
		return nil

	case aopType == "per-hostname":
		record, err := client.GetPerHostnameAuthenticatedOriginPullsCertificate(zoneID, certID)
		if err != nil {
			log.Printf("[WARN] Removing record from state because it's not found in API")
			d.SetId("")
			return nil
		}
		d.Set("certificate", record.Certificate)
		return nil
	default:
		return errors.New("unknown error")
	}
}

func resourceCloudflareAuthenticatedOriginPullsCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cloudflare.API)
	zoneID := d.Get("zone_id").(string)
	certID := d.Id()
	log.Printf("[DEBUG] zone ID: %s", zoneID)

	switch aopType, ok := d.GetOk("type"); ok {
	case aopType == "per-zone":
		_, err := client.DeletePerZoneAuthenticatedOriginPullsCertificate(zoneID, certID)
		if err != nil {
			return fmt.Errorf("Error deleting: %s", err)
		}

		return nil

	case aopType == "per-hostname":
		_, err := client.DeletePerHostnameAuthenticatedOriginPullsCertificate(zoneID, certID)
		if err != nil {
			log.Printf("[WARN] Removing record from state because it's not found in API")
			d.SetId("")
			return nil
		}
		return nil
	default:
		return errors.New("unknown error")
	}
}
