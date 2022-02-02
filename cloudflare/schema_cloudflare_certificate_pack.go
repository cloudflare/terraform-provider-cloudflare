package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCertificatePackSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"custom", "dedicated_custom", "advanced"}, false),
		},
		"hosts": {
			Type:     schema.TypeSet,
			Required: true,
			ForceNew: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"validation_method": {
			Type:         schema.TypeString,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"txt", "http", "email"}, false),
		},
		"validity_days": {
			Type:         schema.TypeInt,
			Optional:     true,
			ForceNew:     true,
			ValidateFunc: validation.IntInSlice([]int{14, 30, 90, 365}),
		},
		"certificate_authority": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"digicert", "lets_encrypt"}, false),
			Default:      nil,
		},
		"validation_records": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem:     sslValidationRecordsSchema(),
		},
		"validation_errors": {
			Type:     schema.TypeList,
			Computed: true,
			Optional: true,
			Elem:     sslValidationErrorsSchema(),
		},
		"cloudflare_branding": {
			Type:     schema.TypeBool,
			Optional: true,
			ForceNew: true,
		},
	}
}
