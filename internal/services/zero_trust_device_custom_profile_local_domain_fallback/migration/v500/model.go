package v500

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ============================================================================
// Source Models (Legacy Provider - v4.x)
// ============================================================================

// SourceCloudflareFallbackDomainModel represents the legacy cloudflare_zero_trust_local_fallback_domain
// (or cloudflare_fallback_domain) resource state from v4.x SDKv2 provider.
//
// Schema version: 0 (SDKv2 default)
// Resource types:
//   - cloudflare_zero_trust_local_fallback_domain
//   - cloudflare_fallback_domain (deprecated alias)
//
// This source model is specifically for resources WITH policy_id (custom profile path).
// Resources without policy_id migrate to the default profile resource instead.
type SourceCloudflareFallbackDomainModel struct {
	ID        types.String                `tfsdk:"id"` // Format: "account_id/policy_id" for custom profile
	AccountID types.String                `tfsdk:"account_id"`
	Domains   *[]SourceFallbackDomainItem `tfsdk:"domains"` // TypeSet stored as array
	PolicyID  types.String                `tfsdk:"policy_id"` // Must be present for custom profile
}

// SourceFallbackDomainItem represents a domain entry from v4.x provider.
// In v4, this was part of a TypeSet of objects.
type SourceFallbackDomainItem struct {
	Suffix      types.String    `tfsdk:"suffix"`      // Optional in v4, Required in v5
	Description types.String    `tfsdk:"description"` // Optional
	DNSServer   *[]types.String `tfsdk:"dns_server"`  // TypeList of strings
}

// ============================================================================
// Target Models (Current Provider - v5.x+)
// ============================================================================

// TargetZeroTrustDeviceCustomProfileLocalDomainFallbackModel represents the current
// cloudflare_zero_trust_device_custom_profile_local_domain_fallback resource state from v5.x+ provider.
//
// Schema version: 500 (when TF_MIG_TEST=1, otherwise 1)
// Resource type: cloudflare_zero_trust_device_custom_profile_local_domain_fallback
//
// Note: This matches the model structure in the parent package's model.go file.
type TargetZeroTrustDeviceCustomProfileLocalDomainFallbackModel struct {
	ID        types.String                                             `tfsdk:"id"` // Format: "policy_id" only (NOT "account_id/policy_id")
	PolicyID  types.String                                             `tfsdk:"policy_id"` // Required in v5
	AccountID types.String                                             `tfsdk:"account_id"`
	Domains   *[]*TargetZeroTrustDeviceCustomProfileDomainFallbackItem `tfsdk:"domains"` // SetNestedAttribute (array of pointers)
}

// TargetZeroTrustDeviceCustomProfileDomainFallbackItem represents a domain entry from v5.x+ provider.
// In v5, this is part of a SetNestedAttribute.
type TargetZeroTrustDeviceCustomProfileDomainFallbackItem struct {
	Suffix      types.String    `tfsdk:"suffix"`      // Required in v5
	Description types.String    `tfsdk:"description"` // Optional
	DNSServer   *[]types.String `tfsdk:"dns_server"`  // ListAttribute of strings
}
