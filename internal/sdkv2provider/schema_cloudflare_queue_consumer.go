package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCloudflareQueueConsumerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The account identifier to target for the resource.",
		},
		"queue_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the queue.",
		},
		"dead_letter_queue": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "ID of the dead letter queue.",
		},
		"environment": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Environment of the consumer.",
		},
		"script_name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "Name of the consumer script.",
		},
		"settings": {
			Type:        schema.TypeSet,
			Optional:    true,
			Description: "Additional settings for the consumer.",
			Elem:        queueConsumerSettings,
		},
	}
}

var queueConsumerSettings = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"max_batch_size": {
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "TODO",
			Default:      10,
			ValidateFunc: validation.IntBetween(1, 100),
		},
		"max_wait_time": &schema.Schema{
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "TODO",
			Default:      5000,
			ValidateFunc: validation.IntBetween(0, 5000),
		},
		"max_retries": &schema.Schema{
			Type:         schema.TypeInt,
			Optional:     true,
			Description:  "TODO",
			Default:      3,
			ValidateFunc: validation.IntBetween(0, 10),
		},
	},
}
