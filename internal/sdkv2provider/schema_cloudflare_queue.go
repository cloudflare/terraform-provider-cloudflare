package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareQueueSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_id": {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the queue.",
		},
		"consumer": {
			Type:        schema.TypeList,
			Optional:    true,
			Description: "Array of consumers.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"created_on": {
						Description: "Created on date",
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
					},
					"environment": {
						Description: "Environment",
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
					},
					"queue_name": {
						Description: "Queue name",
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
					},
					"script_name": {
						Description: "script_name",
						Type:        schema.TypeString,
						Optional:    true,
						ForceNew:    true,
					},
					"settings": {
						Description: "Settings",
						Type:        schema.TypeList,
						Optional:    true,
						ForceNew:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"batch_size": {
									Description: "Batch size",
									Type:        schema.TypeInt,
									Optional:    true,
									ForceNew:    true,
								},
								"max_retries": {
									Description: "Max retries",
									Type:        schema.TypeInt,
									Optional:    true,
									ForceNew:    true,
								},
								"max_wait_time_ms": {
									Description: "Max wait time in milliseconds",
									Type:        schema.TypeInt,
									Optional:    true,
									ForceNew:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}


