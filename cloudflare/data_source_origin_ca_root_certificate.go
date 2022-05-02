package cloudflare

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceCloudflareOriginCARootCertificate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCloudflareOriginCARootCertificateRead,

		Schema: map[string]*schema.Schema{
			"algorithm": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"rsa", "ecc"}, true),
			},

			"cert_pem": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCloudflareOriginCARootCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	algorithm := strings.ToLower(fmt.Sprintf("%s", d.Get("algorithm")))
	certBytes, err := cloudflare.OriginCARootCertificate(algorithm)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Cloudflare Origin CA root %s certificate: %s", algorithm, err))
	}

	cert := string(certBytes[:])

	d.SetId(stringChecksum(cert))

	if err := d.Set("cert_pem", cert); err != nil {
		return diag.FromErr(fmt.Errorf("error setting cert_pem: %s", err))
	}

	return nil
}
