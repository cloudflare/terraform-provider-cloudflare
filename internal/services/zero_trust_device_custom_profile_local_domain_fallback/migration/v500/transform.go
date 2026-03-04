package v500

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// extractPolicyID extracts the policy_id portion from a potentially composite ID.
// v4 format: "account_id/policy_id" (composite) or just "policy_id"
// v5 format: always just "policy_id"
//
// This handles the case where v4 state stored policy_id as the full composite
// device_profiles.id value ("account_id/policy_id") instead of just the policy portion.
func extractPolicyID(compositeID types.String) types.String {
	if compositeID.IsNull() || compositeID.IsUnknown() {
		return compositeID
	}

	idStr := compositeID.ValueString()

	// Check if it's a composite ID (contains "/")
	if slashIdx := strings.Index(idStr, "/"); slashIdx != -1 && slashIdx < len(idStr)-1 {
		// Extract the part after the slash (policy_id portion)
		policyIDStr := idStr[slashIdx+1:]
		return types.StringValue(policyIDStr)
	}

	// Not composite, return as-is
	return compositeID
}

// TransformToCustomProfile converts source (legacy v4 fallback domain) state to target
// (current v5 custom profile) state.
//
// This function is shared by both UpgradeFromV4 and MoveFallbackDomainToCustomProfile handlers.
//
// Transformation includes:
// - Direct copy: account_id, policy_id
// - Transform ID: "account_id/policy_id" → "policy_id"
// - Transform domains: Set → SetNestedAttribute (with empty dns_server filtering)
func TransformToCustomProfile(ctx context.Context, source SourceCloudflareFallbackDomainModel) (*TargetZeroTrustDeviceCustomProfileLocalDomainFallbackModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Validate required fields
	if source.AccountID.IsNull() || source.AccountID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"account_id is required for fallback domain migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	if source.PolicyID.IsNull() || source.PolicyID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"policy_id is required for custom profile migration. The source state is missing this field, which indicates the resource should have migrated to the default profile instead.",
		)
		return nil, diags
	}

	// Extract policy_id from potentially composite format
	// v4 state may have stored policy_id as the full device_profiles.id value
	// ("account_id/policy_id") instead of just the policy portion
	policyID := extractPolicyID(source.PolicyID)

	// Initialize target with extracted policy_id
	target := &TargetZeroTrustDeviceCustomProfileLocalDomainFallbackModel{
		AccountID: source.AccountID, // Direct copy
		PolicyID:  policyID,          // Extracted policy portion only
	}

	// Transform ID: extract policy_id portion if composite
	// v4 ID format: "account_id/policy_id" (composite)
	// v5 ID format: "policy_id" (just the policy ID portion)
	if !source.ID.IsNull() && !source.ID.IsUnknown() {
		// Extract policy_id from ID field (may also be composite)
		target.ID = extractPolicyID(source.ID)
	} else {
		// ID is null/unknown - use extracted policy_id
		target.ID = policyID
	}

	// Transform domains: source array → target array of pointers
	// Filter out empty dns_server arrays per tf-migrate behavior
	if source.Domains != nil && len(*source.Domains) > 0 {
		targetDomains := make([]*TargetZeroTrustDeviceCustomProfileDomainFallbackItem, 0, len(*source.Domains))

		for _, sourceDomain := range *source.Domains {
			targetDomain := &TargetZeroTrustDeviceCustomProfileDomainFallbackItem{
				Suffix:      sourceDomain.Suffix,
				Description: sourceDomain.Description,
			}

			// Handle dns_server: filter out empty arrays (set to null if empty)
			if sourceDomain.DNSServer != nil && len(*sourceDomain.DNSServer) > 0 {
				targetDomain.DNSServer = sourceDomain.DNSServer
			} else {
				// Empty array or nil → set to nil (will be null in state)
				targetDomain.DNSServer = nil
			}

			targetDomains = append(targetDomains, targetDomain)
		}

		target.Domains = &targetDomains
	} else {
		// No domains or empty array
		target.Domains = &[]*TargetZeroTrustDeviceCustomProfileDomainFallbackItem{}
	}

	return target, diags
}
