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
			"consumers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Array of consumers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"environment": {
							Description: "Environment",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"queue_name": {
							Description: "Queue name",
							Type:        schema.TypeString,
							Optional:    true,
						},
						"script_name": {
							Description: "script_name",
							Type:        schema.TypeString,
							Required:   true,
						},
						"settings": {
							Description: "Settings",
							Type:        schema.TypeSet,
							Required:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"batch_size": {
										Description: "Batch size",
										Type:        schema.TypeInt,
										Required:    true,
									},
									"max_retries": {
										Description: "Max retries",
										Type:        schema.TypeInt,
										Optional:    true,
										Default:     3,
									},
									"max_wait_time_ms": {
										Description: "Max wait time in milliseconds",
										Type:        schema.TypeInt,
										Optional:    true,
									},
								},
							},
						},
					},
				},
			},
		}
}

