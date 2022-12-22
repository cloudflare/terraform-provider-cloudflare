package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
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
			Description: "DNS record name (or @ for the zone apex).",
		},

		"hostname": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR", "HTTPS"}, false),
			Description:  "DNS record type.",
		},

		"value": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			ConflictsWith:    []string{"data"},
			DiffSuppressFunc: suppressTrailingDots,
			Description:      "DNS record value.",
		},

		"data": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"value"},
			Description:   "Metadata about the record.",
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
			Description: "Time to live, in seconds, of the DNS record. Must be between 60 and 86400, or 1 for 'automatic'.",
		},

		"priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			DiffSuppressFunc: suppressPriority,
			Description:      "Required for MX, SRV and URI records; unused by other record types. Records with lower priorities are preferred.",
		},

		"proxied": {
			Optional:    true,
			Type:        schema.TypeBool,
			Description: "Whether the record is receiving the performance and security benefits of Cloudflare's network.",
		},

		"created_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "When the record was created.",
		},

		"metadata": {
			Type:     schema.TypeMap,
			Computed: true,
		},

		"modified_on": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "The date and time the record was last modified.",
		},

		"proxiable": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "Whether the record can be proxied by Cloudflare or not.",
		},

		"allow_overwrite": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
