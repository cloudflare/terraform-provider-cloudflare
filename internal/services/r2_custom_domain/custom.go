package r2_custom_domain

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
)

// preserveStateOnDegradedResponse preserves certain state values when the API
// returns a degraded response due to a transient backend failure (e.g., the
// SSL for SaaS API being temporarily unavailable).
//
// When the R2 custom domain GET API cannot reach SSL4SaaS, it returns:
//   - status.ssl = "unknown" and status.ownership = "unknown"
//   - zone_name, min_tls, and ciphers omitted (null)
//
// These degraded values would overwrite real state values (e.g., "active"),
// causing Terraform to detect false drift and potentially trigger unnecessary
// resource replacement. This function detects the degraded response pattern
// and preserves the previous state values instead.
func preserveStateOnDegradedResponse(ctx context.Context, data *R2CustomDomainModel, previousState *R2CustomDomainModel) {
	if data == nil || previousState == nil {
		return
	}

	// Detect whether this is a degraded response by checking if both status
	// fields are "unknown" — this is the sentinel value the API returns when
	// the SSL4SaaS backend is unreachable.
	isDegraded := isDegradedStatusResponse(ctx, data)

	// Preserve status fields: only restore previous values when the API returned
	// "unknown" (indicating a transient failure), not a legitimate status change.
	preserveStatusFields(ctx, data, previousState, isDegraded)

	// When the response is degraded, the API may also omit fields that are
	// normally present. Preserve these from the previous state.
	if isDegraded {
		preserveOmittedFields(data, previousState)
	}
}

// isDegradedStatusResponse returns true if the API response has both status
// fields set to "unknown", which indicates a transient failure rather than
// a real status value.
func isDegradedStatusResponse(ctx context.Context, data *R2CustomDomainModel) bool {
	if data.Status.IsNull() || data.Status.IsUnknown() {
		return false
	}

	currentStatus, diags := data.Status.Value(ctx)
	if diags.HasError() || currentStatus == nil {
		return false
	}

	return currentStatus.SSL.ValueString() == "unknown" &&
		currentStatus.Ownership.ValueString() == "unknown"
}

// preserveStatusFields restores the previous status.ssl and status.ownership
// values when the API returned "unknown" and we had real values in state.
func preserveStatusFields(ctx context.Context, data *R2CustomDomainModel, previousState *R2CustomDomainModel, isDegraded bool) {
	if !isDegraded {
		return
	}

	if previousState.Status.IsNull() || previousState.Status.IsUnknown() {
		return
	}

	previousStatus, diags := previousState.Status.Value(ctx)
	if diags.HasError() || previousStatus == nil {
		return
	}

	// Only restore if the previous state had real (non-unknown) values.
	// If the previous state was also "unknown", there's nothing better to restore.
	hasPreviousSSL := !previousStatus.SSL.IsNull() && !previousStatus.SSL.IsUnknown() && previousStatus.SSL.ValueString() != "unknown"
	hasPreviousOwnership := !previousStatus.Ownership.IsNull() && !previousStatus.Ownership.IsUnknown() && previousStatus.Ownership.ValueString() != "unknown"

	if !hasPreviousSSL && !hasPreviousOwnership {
		return
	}

	restoredStatus := &R2CustomDomainStatusModel{
		SSL:       previousStatus.SSL,
		Ownership: previousStatus.Ownership,
	}

	newStatus, diags := customfield.NewObject(ctx, restoredStatus)
	if diags.HasError() {
		return
	}
	data.Status = newStatus
}

// preserveOmittedFields restores zone_name, min_tls, and ciphers from the previous state
// when the API omits them in a degraded response. These fields are normally
// present but get dropped when the SSL4SaaS backend is unreachable.
func preserveOmittedFields(data *R2CustomDomainModel, previousState *R2CustomDomainModel) {
	// Preserve zone_name if it went null but was previously set
	if data.ZoneName.IsNull() && !previousState.ZoneName.IsNull() && !previousState.ZoneName.IsUnknown() {
		data.ZoneName = previousState.ZoneName
	}

	// Preserve min_tls if it went null but was previously set
	if data.MinTLS.IsNull() && !previousState.MinTLS.IsNull() && !previousState.MinTLS.IsUnknown() {
		data.MinTLS = previousState.MinTLS
	}

	// Preserve ciphers if it went nil but was previously set
	if data.Ciphers == nil && previousState.Ciphers != nil {
		data.Ciphers = previousState.Ciphers
	}
}

// snapshotState creates a shallow copy of the model so we can compare
// against it after the API response overwrites the fields.
//
// Note: This performs a shallow copy, which means pointer fields like Ciphers
// are shared between the original and snapshot. This is safe because:
// 1. The snapshot is only used for reading values in preserveStateOnDegradedResponse
// 2. The data pointer is completely replaced after API unmarshaling
// 3. No code modifies slice contents through pointers
func snapshotState(data *R2CustomDomainModel) *R2CustomDomainModel {
	if data == nil {
		return nil
	}
	snapshot := *data
	return &snapshot
}
