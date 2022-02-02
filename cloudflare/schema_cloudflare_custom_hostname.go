package cloudflare

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomHostnameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"hostname": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringLenBetween(0, 255),
		},
		"custom_origin_server": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ssl": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				SchemaVersion: 1,
				Schema: map[string]*schema.Schema{
					"status": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"method": {
						Type:         schema.TypeString,
						Optional:     true,
						ValidateFunc: validation.StringInSlice([]string{"http", "txt", "email"}, false),
					},
					"type": {
						Type:         schema.TypeString,
						Optional:     true,
						Default:      "dv",
						ValidateFunc: validation.StringInSlice([]string{"dv"}, false),
					},
					"certificate_authority": {
						Type:         schema.TypeString,
						Optional:     true,
						Computed:     true,
						ValidateFunc: validation.StringInSlice([]string{"lets_encrypt", "digicert"}, false),
						Default:      nil,
					},
					"validation_records": {
						Type:     schema.TypeList,
						Computed: true,
						Elem:     sslValidationRecordsSchema(),
					},
					"validation_errors": {
						Type:     schema.TypeList,
						Computed: true,
						Elem:     sslValidationErrorsSchema(),
					},
					"wildcard": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"custom_certificate": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"custom_key": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"settings": {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Resource{
							SchemaVersion: 1,
							Schema: map[string]*schema.Schema{
								"http2": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
								},
								"tls13": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
								},
								"min_tls_version": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"1.0", "1.1", "1.2", "1.3"}, false),
								},
								"ciphers": {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
								},
								"early_hints": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
								},
							},
						},
					},
				},
			},
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ownership_verification": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		"ownership_verification_http": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
