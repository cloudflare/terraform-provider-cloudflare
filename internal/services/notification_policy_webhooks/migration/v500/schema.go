package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// SourceCloudflareNotificationPolicyWebhooksSchema returns the legacy
// cloudflare_notification_policy_webhooks schema (schema_version=0).
// This is used by UpgradeFromV4 to parse state from the legacy SDKv2 provider.
// Reference: cloudflare-terraform-v4/internal/sdkv2provider/schema_cloudflare_notification_policy_webhooks.go
//
// This minimal schema includes only the properties needed for state parsing:
// - Required, Optional, Computed flags
// Validators, PlanModifiers, and Descriptions are intentionally omitted.
func SourceCloudflareNotificationPolicyWebhooksSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			// url was Optional in v4 (Required in v5)
			"url": schema.StringAttribute{
				Optional: true,
			},
			"secret": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"last_success": schema.StringAttribute{
				Computed: true,
			},
			"last_failure": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}
