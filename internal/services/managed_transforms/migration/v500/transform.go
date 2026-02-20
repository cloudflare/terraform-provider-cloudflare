package v500

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4 managed_headers) state to target (v5 managed_transforms) state.
// This function is shared by both UpgradeFromV0 and MoveState handlers.
func Transform(ctx context.Context, source SourceManagedHeadersModel) (*TargetManagedTransformsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	target := &TargetManagedTransformsModel{
		ID:     source.ID,
		ZoneID: source.ZoneID,
	}

	// Convert managed_request_headers: null → empty set, otherwise convert entries
	if source.ManagedRequestHeaders != nil && len(*source.ManagedRequestHeaders) > 0 {
		entries := make([]TargetHeaderEntryModel, 0, len(*source.ManagedRequestHeaders))
		for _, entry := range *source.ManagedRequestHeaders {
			if entry != nil {
				entries = append(entries, TargetHeaderEntryModel{
					ID:      entry.ID,
					Enabled: entry.Enabled,
				})
			}
		}
		reqHeaders, d := customfield.NewObjectSet(ctx, entries)
		diags.Append(d...)
		target.ManagedRequestHeaders = reqHeaders
	} else {
		target.ManagedRequestHeaders = customfield.NewObjectSetMust(ctx, []TargetHeaderEntryModel{})
	}

	// Convert managed_response_headers: null → empty set, otherwise convert entries
	if source.ManagedResponseHeaders != nil && len(*source.ManagedResponseHeaders) > 0 {
		entries := make([]TargetHeaderEntryModel, 0, len(*source.ManagedResponseHeaders))
		for _, entry := range *source.ManagedResponseHeaders {
			if entry != nil {
				entries = append(entries, TargetHeaderEntryModel{
					ID:      entry.ID,
					Enabled: entry.Enabled,
				})
			}
		}
		respHeaders, d := customfield.NewObjectSet(ctx, entries)
		diags.Append(d...)
		target.ManagedResponseHeaders = respHeaders
	} else {
		target.ManagedResponseHeaders = customfield.NewObjectSetMust(ctx, []TargetHeaderEntryModel{})
	}

	return target, diags
}
