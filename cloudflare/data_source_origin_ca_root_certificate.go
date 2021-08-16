package cloudflare

import (
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceCloudflareOriginCARootCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCloudflareOriginCARootCertificateRead,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"rsa", "ecc"}, true),
			},

			"cert_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudflareOriginCARootCertificateRead(d *schema.ResourceData, meta interface{}) error {
	algorithm := strings.ToLower(fmt.Sprintf("%v", d.Get("algorithm")))
	certBytes, err := cloudflare.OriginCARootCertificate(algorithm)
	if err != nil {
		return fmt.Errorf("failed to fetch Cloudflare Origin CA root %s certificate: %s", algorithm, err)
	}

	cert := string(certBytes[:])

	d.SetId(stringChecksum(cert))

	if err := d.Set("cert_pem", cert); err != nil {
		return fmt.Errorf("error setting cert_pem: %s", err)
	}

	return nil
}
