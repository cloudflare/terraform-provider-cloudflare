package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// This is a near-no-op migration: all fields are direct copies except for the
// timestamp fields (created_at, last_success, last_failure), which change from
// plain TypeString (v4) to timetypes.RFC3339 (v5). These are set to null and
// will be repopulated by the API on the next refresh.
func Transform(_ context.Context, source SourceCloudflareNotificationPolicyWebhooksModel) (*TargetCloudflareNotificationPolicyWebhooksModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetCloudflareNotificationPolicyWebhooksModel{
		// Direct copies
		ID:        source.ID,
		AccountID: source.AccountID,
		Name:      source.Name,
		URL:       source.URL,
		Secret:    source.Secret,
		Type:      source.Type,

		// Timestamp fields: v4 TypeString → v5 timetypes.RFC3339
		// Set to null; the API will repopulate on the next refresh.
		CreatedAt:   timetypes.NewRFC3339Null(),
		LastSuccess: timetypes.NewRFC3339Null(),
		LastFailure: timetypes.NewRFC3339Null(),
	}

	return target, diags
}
