package sdkv2provider

import (
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCloudflareNotificationPolicyWebhookSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		consts.AccountIDSchemaKey: {
			Description: consts.AccountIDSchemaDescription,
			Type:        schema.TypeString,
			Required:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the webhook destination.",
		},
		"url": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Description: "The URL of the webhook destinations.",
		},
		"secret": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An optional secret can be provided that will be passed in the `cf-webhook-auth` header when dispatching a webhook notification. Secrets are not returned in any API response body. Refer to the [documentation](https://api.cloudflare.com/#notification-webhooks-create-webhook) for more details.",
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the notification webhook was created.",
		},
		"last_success": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the notification webhook was last successful.",
		},
		"last_failure": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Timestamp of when the notification webhook last faiuled.",
		},
	}
}
