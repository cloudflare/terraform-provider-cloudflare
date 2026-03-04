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
// This source model is specifically for resources WITHOUT policy_id (default profile path).
// Resources with policy_id migrate to the custom profile resource instead.
type SourceCloudflareFallbackDomainModel struct {
	ID        types.String                `tfsdk:"id"` // Format: "account_id" for default profile
	AccountID types.String                `tfsdk:"account_id"`
	Domains   *[]SourceFallbackDomainItem `tfsdk:"domains"` // TypeSet stored as array
	PolicyID  types.String                `tfsdk:"policy_id"` // Must be null/absent for default profile
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

// TargetZeroTrustDeviceDefaultProfileLocalDomainFallbackModel represents the current
// cloudflare_zero_trust_device_default_profile_local_domain_fallback resource state from v5.x+ provider.
//
// Schema version: 500 (when TF_MIG_TEST=1, otherwise 1)
// Resource type: cloudflare_zero_trust_device_default_profile_local_domain_fallback
//
// Note: This matches the model structure in the parent package's model.go file.
type TargetZeroTrustDeviceDefaultProfileLocalDomainFallbackModel struct {
	ID        types.String                                             `tfsdk:"id"` // API-generated ID
	AccountID types.String                                             `tfsdk:"account_id"`
	Domains   *[]*TargetZeroTrustDeviceDefaultProfileDomainFallbackItem `tfsdk:"domains"` // SetNestedAttribute (array of pointers)
}

// TargetZeroTrustDeviceDefaultProfileDomainFallbackItem represents a domain entry from v5.x+ provider.
// In v5, this is part of a SetNestedAttribute.
type TargetZeroTrustDeviceDefaultProfileDomainFallbackItem struct {
	Suffix      types.String    `tfsdk:"suffix"`      // Required in v5
	Description types.String    `tfsdk:"description"` // Optional
	DNSServer   *[]types.String `tfsdk:"dns_server"`  // ListAttribute of strings
}
