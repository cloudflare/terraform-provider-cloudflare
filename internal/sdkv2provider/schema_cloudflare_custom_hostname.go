package sdkv2provider

import (
	"fmt"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareCustomHostnameSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"hostname": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringLenBetween(0, 255),
			Description:  "Hostname you intend to request a certificate for.",
		},
		"custom_origin_server": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The custom origin server used for certificates.",
		},
		"custom_origin_sni": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The [custom origin SNI](https://developers.cloudflare.com/ssl/ssl-for-saas/hostname-specific-behavior/custom-origin) used for certificates.",
		},
		"ssl": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "SSL configuration of the certificate.",
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
						Description:  fmt.Sprintf("Domain control validation (DCV) method used for this hostname. %s", renderAvailableDocumentationValuesStringSlice([]string{"http", "txt", "email"})),
					},
					"type": {
						Type:         schema.TypeString,
						Optional:     true,
						Default:      "dv",
						ValidateFunc: validation.StringInSlice([]string{"dv"}, false),
						Description:  fmt.Sprintf("Level of validation to be used for this hostname. %s", renderAvailableDocumentationValuesStringSlice([]string{"dv"})),
					},
					"certificate_authority": {
						Type:         schema.TypeString,
						Optional:     true,
						Computed:     true,
						ValidateFunc: validation.StringInSlice([]string{"lets_encrypt", "digicert", "google"}, false),
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
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether the certificate covers a wildcard.",
					},
					"custom_certificate": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "If a custom uploaded certificate is used.",
					},
					"custom_key": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "The key for a custom uploaded certificate.",
					},
					"settings": {
						Type:        schema.TypeList,
						Optional:    true,
						Computed:    true,
						Description: "SSL/TLS settings for the certificate.",
						Elem: &schema.Resource{
							SchemaVersion: 1,
							Schema: map[string]*schema.Schema{
								"http2": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									Description:  fmt.Sprintf("Whether HTTP2 should be supported. %s", renderAvailableDocumentationValuesStringSlice([]string{"on", "off"})),
								},
								"tls13": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									Description:  fmt.Sprintf("Whether TLSv1.3 should be supported. %s", renderAvailableDocumentationValuesStringSlice([]string{"on", "off"})),
								},
								"min_tls_version": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"1.0", "1.1", "1.2", "1.3"}, false),
									Description:  fmt.Sprintf("Lowest version of TLS this certificate should support. %s", renderAvailableDocumentationValuesStringSlice([]string{"1.0", "1.1", "1.2", "1.3"})),
								},
								"ciphers": {
									Type:     schema.TypeSet,
									Optional: true,
									Elem: &schema.Schema{
										Type: schema.TypeString,
									},
									Description: "List of SSL/TLS ciphers to associate with this certificate.",
								},
								"early_hints": {
									Type:         schema.TypeString,
									Optional:     true,
									ValidateFunc: validation.StringInSlice([]string{"on", "off"}, false),
									Description:  fmt.Sprintf("Whether early hints should be supported. %s", renderAvailableDocumentationValuesStringSlice([]string{"on", "off"})),
								},
							},
						},
					},
				},
			},
		},
		"custom_metadata": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "Custom metadata associated with custom hostname. Only supports primitive string values, all other values are accessible via the API directly.",
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"status": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Status of the certificate.",
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
		"wait_for_ssl_pending_validation": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Whether to wait for a custom hostname SSL sub-object to reach status `pending_validation` during creation.",
		},
	}
}
