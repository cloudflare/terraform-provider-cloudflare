package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// TransformToDefaultProfile transforms a v4 fallback domain (without policy_id) to v5 default profile structure.
//
// Key transformations:
// - ID: Keep as-is from source (should be account_id)
// - AccountID: Direct copy
// - Domains: Transform from v4 TypeSet to v5 SetNestedAttribute (pointers)
// - PolicyID: Must not be present in source (validate during move)
func TransformToDefaultProfile(ctx context.Context, source SourceCloudflareFallbackDomainModel) (TargetZeroTrustDeviceDefaultProfileLocalDomainFallbackModel, diag.Diagnostics) {
	var diags diag.Diagnostics
	target := TargetZeroTrustDeviceDefaultProfileLocalDomainFallbackModel{}

	// Copy ID and AccountID directly
	target.ID = source.ID
	target.AccountID = source.AccountID

	// Transform domains from v4 TypeSet to v5 SetNestedAttribute
	if source.Domains != nil && len(*source.Domains) > 0 {
		targetDomains := make([]*TargetZeroTrustDeviceDefaultProfileDomainFallbackItem, 0, len(*source.Domains))

		for _, sourceDomain := range *source.Domains {
			targetDomain := &TargetZeroTrustDeviceDefaultProfileDomainFallbackItem{
				Suffix:      sourceDomain.Suffix,
				Description: sourceDomain.Description,
			}

			// Transform dns_server list if present
			if sourceDomain.DNSServer != nil && len(*sourceDomain.DNSServer) > 0 {
				dnsServers := make([]string, 0, len(*sourceDomain.DNSServer))
				for _, server := range *sourceDomain.DNSServer {
					if !server.IsNull() && !server.IsUnknown() {
						dnsServers = append(dnsServers, server.ValueString())
					}
				}

				// Only set dns_server if we have values
				if len(dnsServers) > 0 {
					targetDomain.DNSServer = sourceDomain.DNSServer
				}
			}

			targetDomains = append(targetDomains, targetDomain)
		}

		target.Domains = &targetDomains
	}

	return target, diags
}
