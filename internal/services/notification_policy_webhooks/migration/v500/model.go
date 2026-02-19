package v500

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x SDKv2)
// ============================================================================

// SourceCloudflareNotificationPolicyWebhooksModel represents the source
// cloudflare_notification_policy_webhooks state structure from the v4 SDKv2 provider.
// This corresponds to schema_version=0 in the raw v4 SDKv2 format.
// Used by UpgradeFromV4 to parse version 0 state.
type SourceCloudflareNotificationPolicyWebhooksModel struct {
	ID          types.String `tfsdk:"id"`
	AccountID   types.String `tfsdk:"account_id"`
	Name        types.String `tfsdk:"name"`
	URL         types.String `tfsdk:"url"`
	Secret      types.String `tfsdk:"secret"`
	Type        types.String `tfsdk:"type"`
	CreatedAt   types.String `tfsdk:"created_at"`
	LastSuccess types.String `tfsdk:"last_success"`
	LastFailure types.String `tfsdk:"last_failure"`
}

// ============================================================================
// Target Models (Current Provider - v5.x+ Plugin Framework)
// ============================================================================

// TargetCloudflareNotificationPolicyWebhooksModel represents the target
// cloudflare_notification_policy_webhooks state structure (version=500).
// This matches the model in the parent package (internal/services/notification_policy_webhooks/model.go).
type TargetCloudflareNotificationPolicyWebhooksModel struct {
	ID          types.String      `tfsdk:"id"`
	AccountID   types.String      `tfsdk:"account_id"`
	Name        types.String      `tfsdk:"name"`
	URL         types.String      `tfsdk:"url"`
	Secret      types.String      `tfsdk:"secret"`
	CreatedAt   timetypes.RFC3339 `tfsdk:"created_at"`
	LastFailure timetypes.RFC3339 `tfsdk:"last_failure"`
	LastSuccess timetypes.RFC3339 `tfsdk:"last_success"`
	Type        types.String      `tfsdk:"type"`
}
