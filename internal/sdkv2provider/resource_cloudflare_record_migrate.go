package sdkv2provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareRecordV1() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"data": {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"value"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"protocol": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"public_key": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"digest_type": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"digest": {
							Type:     schema.TypeString,
							Optional: true,
						},
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
						"fingerprint": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"content": {
							Type:     schema.TypeString,
							Optional: true,
						},
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
		},
	}
}

func resourceCloudflareRecordStateUpgradeV2(_ context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	rawState["data"] = []interface{}{rawState["data"]}
	return rawState, nil
}
