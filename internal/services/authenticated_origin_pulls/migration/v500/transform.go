package v500

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// Transform converts source (legacy v4) state to target (current v5) state.
//
// v4 structure (flat):
//   - zone_id
//   - hostname
//   - authenticated_origin_pulls_certificate
//   - enabled
//
// v5 structure (nested config):
//   - zone_id
//   - config[0].hostname
//   - config[0].cert_id
//   - config[0].enabled
//
// This function handles the restructuring from flat fields to nested config list.
func Transform(ctx context.Context, source SourceCloudflareAuthenticatedOriginPullsModel) (*TargetAuthenticatedOriginPullsModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Step 1: Validate required fields
	if source.ZoneID.IsNull() || source.ZoneID.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"zone_id is required for authenticated_origin_pulls migration. The source state is missing this field, which indicates corrupted state.",
		)
		return nil, diags
	}

	// Validate Per-Hostname mode fields (this resource only handles Per-Hostname AOP)
	if source.Hostname.IsNull() || source.Hostname.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"hostname is required for authenticated_origin_pulls migration. Resources without hostname should migrate to cloudflare_authenticated_origin_pulls_settings instead.",
		)
		return nil, diags
	}

	if source.AuthenticatedOriginPullsCertificate.IsNull() || source.AuthenticatedOriginPullsCertificate.IsUnknown() {
		diags.AddError(
			"Missing required field",
			"authenticated_origin_pulls_certificate is required for Per-Hostname authenticated origin pulls migration.",
		)
		return nil, diags
	}

	// Step 2: Initialize target model
	target := &TargetAuthenticatedOriginPullsModel{
		ZoneID:   source.ZoneID,
		Hostname: source.Hostname, // Set top-level hostname (required for Read() API calls)
		ID:       source.Hostname, // ID matches hostname in v5
		// All other computed fields will be populated by API refresh
	}

	// Step 3: Restructure flat fields → nested config list
	// The v5 resource requires exactly one config item (enforced by resource.go)
	config := &TargetAuthenticatedOriginPullsConfigModel{
		Hostname: source.Hostname,
		CERTID:   source.AuthenticatedOriginPullsCertificate,
		Enabled:  source.Enabled,
	}

	// Create config list with single item
	configList := []*TargetAuthenticatedOriginPullsConfigModel{config}
	target.Config = &configList

	return target, diags
}
