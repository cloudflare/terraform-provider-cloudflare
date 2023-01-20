package sdkv2provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/MakeNowJust/heredoc/v2"
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
				Description:  fmt.Sprintf("The name of the algorithm used when creating an Origin CA certificate. %s", renderAvailableDocumentationValuesStringSlice([]string{"rsa", "ecc"})),
			},

			"cert_pem": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Origin CA root certificate in PEM format.",
			},
		},
		Description: heredoc.Doc(`
			Use this data source to get the
			[Origin CA root certificate](https://developers.cloudflare.com/ssl/origin-configuration/origin-ca#4-required-for-some-add-cloudflare-origin-ca-root-certificates)
			for a given algorithm."
		`),
	}
}

func dataSourceCloudflareOriginCARootCertificateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	algorithm := strings.ToLower(fmt.Sprintf("%s", d.Get("algorithm")))
	certBytes, err := cloudflare.GetOriginCARootCertificate(algorithm)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch Cloudflare Origin CA root %s certificate: %w", algorithm, err))
	}

	cert := string(certBytes[:])

	d.SetId(stringChecksum(cert))

	if err := d.Set("cert_pem", cert); err != nil {
		return diag.FromErr(fmt.Errorf("error setting cert_pem: %w", err))
	}

	return nil
}
