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

	return diags
}
