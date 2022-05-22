package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareRecordSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"zone_id": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},

		"name": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
			StateFunc: func(i interface{}) string {
				return strings.ToLower(i.(string))
			},
		},

		"hostname": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"type": {
			Type:         schema.TypeString,
			Required:     true,
			ForceNew:     true,
			ValidateFunc: validation.StringInSlice([]string{"A", "AAAA", "CAA", "CNAME", "TXT", "SRV", "LOC", "MX", "NS", "SPF", "CERT", "DNSKEY", "DS", "NAPTR", "SMIMEA", "SSHFP", "TLSA", "URI", "PTR"}, false),
		},

		"value": {
			Type:          schema.TypeString,
			Optional:      true,
			Computed:      true,
			ConflictsWith: []string{"data"},
		},

		"data": {
			Type:          schema.TypeList,
			MaxItems:      1,
			Optional:      true,
			ConflictsWith: []string{"value"},
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
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},

		"priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			DiffSuppressFunc: suppressPriority,
		},

		"proxied": {
			Optional: true,
			Type:     schema.TypeBool,
		},

		"created_on": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"metadata": {
			Type:     schema.TypeMap,
			Computed: true,
		},

		"modified_on": {
			Type:     schema.TypeString,
			Computed: true,
		},

		"proxiable": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"allow_overwrite": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
	}
}
