package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareOriginCACertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"csr": {
			Type:         schema.TypeString,
			Optional:     true,
			ValidateFunc: validateCSR,
		},
		"expires_on": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"hostnames": {
			Type:     schema.TypeSet,
			Required: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"request_type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"origin-rsa", "origin-ecc", "keyless-certificate"}, false),
		},
		"requested_validity": {
			Type:         schema.TypeInt,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validation.IntInSlice([]int{7, 30, 90, 365, 730, 1095, 5475}),
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return true
			},
		},
	}
}
