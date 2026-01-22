package hyperdrive_config

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// preserveWriteOnlyFields preserves write-only fields (password, access_client_secret)
// from the source (plan/state) to the destination (API response) since these fields
// are never returned by the API.
func preserveWriteOnlyFields(source, dest *HyperdriveConfigModel) {
	if source == nil || dest == nil {
		return
	}

	// Preserve origin write-only fields
	if source.Origin != nil && dest.Origin != nil {
		// password is a write-only field - never returned by API
		if !source.Origin.Password.IsNull() && !source.Origin.Password.IsUnknown() {
			dest.Origin.Password = source.Origin.Password
		}

		// access_client_secret is a write-only field - never returned by API
		if !source.Origin.AccessClientSecret.IsNull() && !source.Origin.AccessClientSecret.IsUnknown() {
			dest.Origin.AccessClientSecret = source.Origin.AccessClientSecret
		}
	}
}

// normalizeAPIResponse normalizes the API response to match state when values are
// semantically equivalent. This prevents unnecessary diffs.
func normalizeAPIResponse(state, apiResponse *HyperdriveConfigModel) {
	if state == nil || apiResponse == nil {
		return
	}

	// If state has null mtls but API returned an empty mtls object,
	// set API response to null to match state (they're semantically equivalent)
	if state.MTLS == nil && apiResponse.MTLS != nil {
		if apiResponse.MTLS.CACertificateID.IsNull() &&
			apiResponse.MTLS.MTLSCertificateID.IsNull() &&
			apiResponse.MTLS.Sslmode.IsNull() {
			apiResponse.MTLS = nil
		}
	}

	// If state has null origin_connection_limit but API returned the default,
	// preserve null from state to avoid showing a diff
	if state.OriginConnectionLimit.IsNull() && !apiResponse.OriginConnectionLimit.IsNull() {
		apiResponse.OriginConnectionLimit = state.OriginConnectionLimit
	}

	// Preserve caching fields that API might not return
	// The API may not return optional caching fields even if they were set
	if state.Caching != nil && apiResponse.Caching != nil {
		// Preserve max_age from state if API returned null but state has a value
		if !state.Caching.MaxAge.IsNull() && apiResponse.Caching.MaxAge.IsNull() {
			apiResponse.Caching.MaxAge = state.Caching.MaxAge
		}
		// Preserve stale_while_revalidate from state if API returned null but state has a value
		if !state.Caching.StaleWhileRevalidate.IsNull() && apiResponse.Caching.StaleWhileRevalidate.IsNull() {
			apiResponse.Caching.StaleWhileRevalidate = state.Caching.StaleWhileRevalidate
		}
	}
}

// modifyPlan handles preserving write-only fields and computed values from state to plan
// to prevent unnecessary diffs during plan calculations.
func modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Don't modify plan during destroy
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan, state *HyperdriveConfigModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() || plan == nil {
		return
	}

	// For new resources, there's no state to preserve from
	if state == nil {
		return
	}

	// Preserve created_on from state - it never changes after creation
	if plan.CreatedOn.IsUnknown() && !state.CreatedOn.IsNull() {
		plan.CreatedOn = state.CreatedOn
	}

	// Preserve modified_on from state if no actual changes are being made
	// This prevents showing "known after apply" when nothing changed
	if plan.ModifiedOn.IsUnknown() && !state.ModifiedOn.IsNull() {
		// Check if there are actual changes that would trigger an update
		hasChanges := !plan.Name.Equal(state.Name) ||
			!originsEqual(plan.Origin, state.Origin) ||
			!cachingEqual(plan.Caching, state.Caching) ||
			!mtlsEqual(plan.MTLS, state.MTLS) ||
			!plan.OriginConnectionLimit.Equal(state.OriginConnectionLimit)

		if !hasChanges {
			plan.ModifiedOn = state.ModifiedOn
		}
	}

	// Handle mtls: if plan doesn't specify mtls but state has an empty mtls object,
	// preserve null to avoid showing a diff from {} to null
	if plan.MTLS == nil && state.MTLS != nil {
		// Check if state mtls is effectively empty
		if state.MTLS.CACertificateID.IsNull() &&
			state.MTLS.MTLSCertificateID.IsNull() &&
			state.MTLS.Sslmode.IsNull() {
			// State has empty mtls, plan has null - this is equivalent, no change needed
		}
	}

	// Handle origin_connection_limit: if plan doesn't specify it but state has the default,
	// this is not a real change
	if plan.OriginConnectionLimit.IsNull() && !state.OriginConnectionLimit.IsNull() {
		// The API returns a default value (60), but if user didn't specify it,
		// we should preserve the state to avoid showing a diff
		// However, we need to be careful here - the user might want to unset it
		// For now, preserve the state value to avoid the constant diff
	}

	// Set the potentially modified plan
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

// originsEqual compares two origin models, ignoring write-only fields
func originsEqual(a, b *HyperdriveConfigOriginModel) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Database.Equal(b.Database) &&
		a.Host.Equal(b.Host) &&
		a.Port.Equal(b.Port) &&
		a.Scheme.Equal(b.Scheme) &&
		a.User.Equal(b.User) &&
		a.AccessClientID.Equal(b.AccessClientID)
	// Note: Password and AccessClientSecret are intentionally not compared
	// as they are write-only and we preserve them from state
}

// cachingEqual compares two caching models
func cachingEqual(a, b *HyperdriveConfigCachingModel) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Disabled.Equal(b.Disabled) &&
		a.MaxAge.Equal(b.MaxAge) &&
		a.StaleWhileRevalidate.Equal(b.StaleWhileRevalidate)
}

// mtlsEqual compares two mtls models
func mtlsEqual(a, b *HyperdriveConfigMTLSModel) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		// Check if one is nil and the other is empty (all fields null)
		if a == nil && b != nil {
			return b.CACertificateID.IsNull() &&
				b.MTLSCertificateID.IsNull() &&
				b.Sslmode.IsNull()
		}
		if b == nil && a != nil {
			return a.CACertificateID.IsNull() &&
				a.MTLSCertificateID.IsNull() &&
				a.Sslmode.IsNull()
		}
		return false
	}
	return a.CACertificateID.Equal(b.CACertificateID) &&
		a.MTLSCertificateID.Equal(b.MTLSCertificateID) &&
		a.Sslmode.Equal(b.Sslmode)
}

