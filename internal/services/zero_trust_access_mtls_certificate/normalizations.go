package zero_trust_access_mtls_certificate

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Normalizing function to ensure consistency between the state/plan and the meaning of the API response.
// Alters the API response before applying it to the state by laxing equalities between null & zero-value
// for some attributes, and nullifies fields that terraform should not be saving in the state.
func normalizeReadZeroTrustAccessMtlsCertificateAPIData(ctx context.Context, data, sourceData *ZeroTrustAccessMTLSCertificateModel) diag.Diagnostics {
	diags := make(diag.Diagnostics, 0)

	if sourceStr := sourceData.Certificate.ValueString(); data.Certificate.IsNull() && sourceStr != "" {
		// Certificate might not be in the API response for certain endpoints. Need to persist it from the state
		data.Certificate = sourceData.Certificate
	}

	// Normalize associated_hostnames to handle empty list vs null differences between provider versions
	// If the API returns nil but source had either nil or empty slice, preserve source to avoid drift
	if data.AssociatedHostnames == nil {
		// If source is nil, keep data as nil (don't apply default)
		// If source is empty slice, preserve the empty slice
		data.AssociatedHostnames = sourceData.AssociatedHostnames
	} else if sourceData.AssociatedHostnames == nil && data.AssociatedHostnames != nil && len(*data.AssociatedHostnames) == 0 {
		// If source was nil but data has empty slice from API, keep it as nil to match source
		data.AssociatedHostnames = nil
	}

	return diags
}
