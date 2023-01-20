package sdkv2provider

import (
	"fmt"
	"strings"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.ZoneIDSchemaKey: {
			Description: "The zone identifier to target for the resource.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},

			Description: "The name of the record.",
		},

		"hostname": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The FQDN of the record.",
		},

		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS"}, false),
			Description:  fmt.Sprintf("The type of the record. %s", renderAvailableDocumentationValuesStringSlice([]string{"A", "AAAA", "CAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS"})),
		},

		"value": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			ConflictsWith:    []string{"data"},
			DiffSuppressFunc: suppressTrailingDots,
			Description:      "The value of the record.",
		},

		"data": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"value"},
			Description:   "Map of attributes that constitute the record value.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					// Properties present in several record types
					"algorithm": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"key_tag": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"flags": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"service": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"certificate": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"type": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"usage": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"selector": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"matching_type": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"weight": {
						Type:     schema.TypeInt,
						Optional: true,
					},

					// SRV record properties
					"proto": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"name": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"priority": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"port": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"target": {
						Type:     schema.TypeString,
						Optional: true,
					},

					// LOC record properties
					"size": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"altitude": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"long_degrees": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"lat_degrees": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"precision_horz": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"precision_vert": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"long_direction": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"long_minutes": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"long_seconds": {
						Type:     schema.TypeFloat,
						Optional: true,
					},
					"lat_direction": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"lat_minutes": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"lat_seconds": {
						Type:     schema.TypeFloat,
						Optional: true,
					},

					// DNSKEY record properties
					"protocol": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"public_key": {
						Type:     schema.TypeString,
						Optional: true,
					},

					// DS record properties
					"digest_type": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"digest": {
						Type:     schema.TypeString,
						Optional: true,
					},

					// NAPTR record properties
					"order": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"preference": {
						Type:     schema.TypeInt,
						Optional: true,
					},
					"regex": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"replacement": {
						Type:     schema.TypeString,
						Optional: true,
					},

					// SSHFP record properties
					"fingerprint": {
						Type:     schema.TypeString,
						Optional: true,
					},

					// URI record properties
					"content": {
						Type:     schema.TypeString,
						Optional: true,
					},

					// CAA record properties
					"tag": {
						Type:     schema.TypeString,
						Optional: true,
					},

					"value": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},

		"ttl": {
			Type:        schema.TypeInt,
			Optional:    true,
			Computed:    true,
			Description: "The TTL of the record.",
		},

		"priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			DiffSuppressFunc: suppressPriority,
			Description:      "The priority of the record.",
		},

		"proxied": {
			Optional:    true,
			Type:        schema.TypeBool,
			Description: "Whether the record gets Cloudflare's origin protection.",
		},

		"created_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the record was created.",
		},

		"metadata": {
			Type:        schema.TypeMap,
			Computed:    true,
			Description: "A key-value map of string metadata Cloudflare associates with the record.",
		},

		"modified_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The RFC3339 timestamp of when the record was last modified.",
		},

		"proxiable": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Shows whether this record can be proxied.",
		},

		"allow_overwrite": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			Description: "Allow creation of this record in Terraform to overwrite an existing record, if any. This does not affect the ability to update the record in Terraform and does not prevent other resources within Terraform or manual changes outside Terraform from overwriting this record. **This configuration is not recommended for most environments**",
		},

		"comment": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Comments or notes about the DNS record. This field has no effect on DNS responses.",
		},

		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Description: "Custom tags for the DNS record.",
		},
	}
}
